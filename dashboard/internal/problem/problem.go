package problem

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type ProblemService interface {
	Create(context.Context, *models.Problem) (*models.Problem, models.AppError)
	Fetch(context.Context, string) (*models.Problem, models.AppError)
}

type problemService struct {
	repo repository.ProblemRepository
}

func NewProblemService(repo repository.ProblemRepository) *problemService {
	return &problemService{
		repo: repo,
	}
}

const (
	TESTFILE = "testfile"
	PROBLEM  = "problem"
)

func (s *problemService) Create(ctx context.Context, problem *models.Problem) (*models.Problem, models.AppError) {
	problem.ID = uuid.New()

	problemFileID, err := s.storeFile(ctx, []byte(problem.ProblemFile), PROBLEM, problem.ID, problem.Title)
	if err != models.NoError {
		return nil, err
	}
	problem.ProblemLink = generateFileLink(problemFileID, PROBLEM)

	testLinks := make([]models.TestLinkType, len(problem.TestFiles))
	for i, testFile := range problem.TestFiles {
		testFileID, err := s.storeFile(ctx, []byte(testFile.File), TESTFILE, problem.ID, testFile.Title)
		if err != models.NoError {
			return nil, err
		}
		testLinks[i] = models.TestLinkType{
			Language: testFile.Language,
			Link:     generateFileLink(testFileID, TESTFILE),
		}
	}
	problem.TestLinks = testLinks

	if err := s.repo.CreateProblem(ctx, *problem); err != models.NoError {
		return nil, err
	}

	return problem, models.NoError
}

func (s *problemService) Fetch(ctx context.Context, id string) (*models.Problem, models.AppError) {
	if problemID, err := uuid.Parse(id); err != nil {
		return nil, models.InternalError.Add(err)
	} else {
		if problem, err := s.repo.GetProblemByID(ctx, problemID); err != models.NoError {
			return nil, err
		} else {
			return &problem, models.NoError
		}
	}
}
