package problem

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type ProblemService interface {
	Create(context.Context, *models.Problem) (*models.Problem, models.AppError)
	Fetch(context.Context, string, string) ([]*models.Problem, models.AppError)
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

func (s *problemService) Fetch(ctx context.Context, id, limit string) ([]*models.Problem, models.AppError) {
	var problems []*models.Problem
	switch {
	case id != "":
		if problemID, err := uuid.Parse(id); err != nil {
			return problems, models.InternalError.Add(err)
		} else {
			if problem, err := s.repo.GetProblemByID(ctx, problemID); err != models.NoError {
				return nil, err
			} else {
				problems = append(problems, &problem)
				return problems, models.NoError
			}
		}

	case limit != "":
		if limit, err := strconv.ParseInt(limit, 10, 64); err != nil {
			return s.repo.GetProblemsListByLimit(ctx, 1, 10)

		} else {
			return s.repo.GetProblemsListByLimit(ctx, 1, int(limit))
		}
	default:

		return s.repo.GetProblemsListByLimit(ctx, 1, 10)
	}
}
