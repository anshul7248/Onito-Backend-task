package testTask

import (
	"context"
	models "testTask/model"
)

type TestTaskUsecaseInterface interface {
	LongestDurationMovies(ctx context.Context) (*models.ApiResponse, error)
	NewMovie(ctx context.Context, tconst, titletype, primarytitle, runtimeminutes, genres string) (*models.ApiResponse, error)
	MoviesByAverageRating(ctx context.Context) (*models.ApiResponse, error)
	UpdateRunTimes(ctx context.Context) (*models.ApiResponse, error)
	MovieBySubtotal(ctx context.Context) (*models.ApiResponse, error)
}
