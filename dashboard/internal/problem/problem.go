package problem

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type Service interface {
	Create(context.Context, *models.Problem) (*models.Problem, models.AppError)
	Fetch(context.Context, int) (*models.Problem, models.AppError)
	GetInitCode(context.Context, int, string) (*models.InitProblemCode, models.AppError)
	GetProblemsForSubject(context.Context, int, int) ([]*models.Problem, models.AppError)
}

type service struct {
	repo repository.ProblemRepository
	fs   repository.FileSystem
}

func New(repo repository.ProblemRepository, fs repository.FileSystem) *service {
	return &service{
		repo: repo,
		fs:   fs,
	}
}

func (s *service) Create(ctx context.Context, problem *models.Problem) (*models.Problem, models.AppError) {
	var err models.AppError
	if problem.ProblemLink, err = s.fs.StoreFile(ctx, []byte(problem.ProblemFile), fmt.Sprintf("%d", problem.ID), models.PROBLEM, models.Text.GetExtension()); err != models.NoError {
		return nil, err
	}

	testLinks := make([]models.TestLinkType, len(problem.TestFiles))
	for i, testFile := range problem.TestFiles {
		testLinks[i].Language = testFile.Language
		if testLinks[i].Link, err = s.fs.StoreFile(ctx, []byte(testFile.File), uuid.New().String(), models.TESTFILE, testFile.Language.GetExtension()); err != models.NoError {
			return nil, err
		}
	}
	problem.TestLinks = testLinks

	if err := s.repo.CreateProblem(ctx, *problem); err != models.NoError {
		return nil, err
	}

	return problem, models.NoError
}

func (s *service) Fetch(ctx context.Context, id int) (*models.Problem, models.AppError) {
	var problem models.Problem
	var err models.AppError

	if problem, err = s.repo.GetProblemByID(ctx, id); err != models.NoError {
		return nil, err
	}
	return &problem, err
}

func (s *service) GetInitCode(ctx context.Context, id int, lang string) (*models.InitProblemCode, models.AppError) {
	problem, err := s.repo.GetProblemByID(ctx, id)
	if err != models.NoError {
		return nil, err
	}

	var link string
	for _, test := range problem.TestLinks {
		if test.Language == models.ProgrammingLanguageEnum(lang) {
			link = test.Link
			break
		}
	}

	code, err := s.fs.GetFile(ctx, link)
	return models.NewInitProblemCode(code), err
}

func (s *service) GetProblemsForSubject(ctx context.Context, subjectID, collegeID int) ([]*models.Problem, models.AppError) {
	return s.repo.GetProblemsForSubject(ctx, subjectID, collegeID)
}
