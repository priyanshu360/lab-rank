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
	fs   repository.FileSystem
}

func NewProblemService(repo repository.ProblemRepository, fs repository.FileSystem) *problemService {
	return &problemService{
		repo: repo,
		fs:   fs,
	}
}

func (s *problemService) Create(ctx context.Context, problem *models.Problem) (*models.Problem, models.AppError) {
	problem.ID = uuid.New()

	var err models.AppError
	if problem.ProblemLink, err = s.fs.StoreFile(ctx, []byte(problem.ProblemFile), problem.ID, models.PROBLEM, models.Text.GetExtension()); err != models.NoError {
		return nil, err
	}

	testLinks := make([]models.TestLinkType, len(problem.TestFiles))
	for i, testFile := range problem.TestFiles {
		testLinks[i].Language = testFile.Language
		if testLinks[i].Link, err = s.fs.StoreFile(ctx, []byte(testFile.File), problem.ID, models.TESTFILE, testFile.Language.GetExtension()); err != models.NoError {
			return nil, err
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
