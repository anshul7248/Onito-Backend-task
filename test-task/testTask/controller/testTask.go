package controller

import (
	"context"
	"net/http"
	testTask "testTask/testTask"

	"github.com/labstack/echo"
)

type TestTaskController struct {
	testTaskUsecase testTask.TestTaskUsecaseInterface
}

func (r *TestTaskController) LongestDurationMovies(c echo.Context) error {
	var t map[string]interface{}
	c.Bind(&t)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	authresp, _ := r.testTaskUsecase.LongestDurationMovies(ctx)
	if authresp == nil {
		return c.JSON(http.StatusUnauthorized, authresp)
	}
	return c.JSON(http.StatusOK, authresp)
}

func (r *TestTaskController) NewMovie(c echo.Context) error {
	var t map[string]interface{}
	c.Bind(&t)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	authresp, _ := r.testTaskUsecase.NewMovie(ctx, t["tconst"].(string), t["titletype"].(string), t["primarytitle"].(string), t["runtimeminutes"].(string), t["genres"].(string))
	if authresp == nil {
		return c.JSON(http.StatusUnauthorized, authresp)
	}
	return c.JSON(http.StatusOK, authresp)
}
func (r *TestTaskController) MoviesByAverageRating(c echo.Context) error {
	var t map[string]interface{}
	c.Bind(&t)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	authresp, _ := r.testTaskUsecase.MoviesByAverageRating(ctx)
	if authresp == nil {
		return c.JSON(http.StatusUnauthorized, authresp)
	}
	return c.JSON(http.StatusOK, authresp)
}

func (r *TestTaskController) UpdateRunTimes(c echo.Context) error {
	var t map[string]interface{}
	c.Bind(&t)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	authResponse, _ := r.testTaskUsecase.UpdateRunTimes(ctx)

	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}

	return c.JSON(http.StatusOK, authResponse)

}
func (r *TestTaskController) MovieBySubtotal(c echo.Context) error {
	var t map[string]interface{}
	c.Bind(&t)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	authresp, _ := r.testTaskUsecase.MovieBySubtotal(ctx)
	if authresp == nil {
		return c.JSON(http.StatusUnauthorized, authresp)
	}
	return c.JSON(http.StatusOK, authresp)
}

func NewTestTaskController(e *echo.Echo, testTaskcase testTask.TestTaskRepoInterface) {
	handler := &TestTaskController{
		testTaskUsecase: testTaskcase,
	}
	e.GET("/api/v1/longest-duration-movies", handler.LongestDurationMovies)
	e.POST("/api/v1/new-movie", handler.NewMovie)
	e.GET("/api/v1/top-rated-movies", handler.MoviesByAverageRating)
	e.POST("/api/v1/update-runtime-minutes", handler.UpdateRunTimes)
	e.GET("/api/v1/genre-movies-with-subtotals", handler.MovieBySubtotal)

}
