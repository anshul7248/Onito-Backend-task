package usecase

import (
	"context"
	models "testTask/model"
	testTask "testTask/testTask"
)

type TestTaskUsecase struct {
	testTaskRepos testTask.TestTaskRepoInterface
}

func NewTestTaskUsecase(repo testTask.TestTaskRepoInterface) testTask.TestTaskUsecaseInterface {
	return &TestTaskUsecase{
		testTaskRepos: repo,
	}
}
func (r *TestTaskUsecase) LongestDurationMovies(ctx context.Context) (*models.ApiResponse, error) {
	return r.testTaskRepos.LongestDurationMovies(ctx)
}
func (r *TestTaskUsecase) NewMovie(ctx context.Context, tconst, titletype, primarytitle, runtimeminutes, genres string) (*models.ApiResponse, error) {
	return r.testTaskRepos.NewMovie(ctx, tconst, titletype, primarytitle, runtimeminutes, genres)
}
func (r *TestTaskUsecase) MoviesByAverageRating(ctx context.Context) (*models.ApiResponse, error) {
	return r.testTaskRepos.MoviesByAverageRating(ctx)
}
func (r *TestTaskUsecase) UpdateRunTimes(ctx context.Context) (*models.ApiResponse, error) {
	return r.testTaskRepos.UpdateRunTimes(ctx)
}

func (r *TestTaskUsecase) MovieBySubtotal(ctx context.Context) (*models.ApiResponse, error) {
	return r.testTaskRepos.MovieBySubtotal(ctx)
}
