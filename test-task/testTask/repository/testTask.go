package repository

import (
	"context"
	"fmt"
	"strconv"
	models "testTask/model"
	testTask "testTask/testTask"

	"github.com/jinzhu/gorm"
)

type TestTaskRepo struct {
	DBConn *gorm.DB
}

func NewTestTaskRepo(Conn *gorm.DB) testTask.TestTaskRepoInterface {
	return &TestTaskRepo{
		DBConn: Conn,
	}
}
func (r *TestTaskRepo) LongestDurationMovies(ctx context.Context) (*models.ApiResponse, error) {
	st := []models.Movies{}

	query := fmt.Sprintf("select tconst,primarytitle, runtimeminutes, genres from movies order by " +
		"runtimeminutes DESC limit 10")

	if err := r.DBConn.Raw(query).Scan(&st); err.Error != nil {
		return &models.ApiResponse{Status: "0", Msg: "No data Found", ResponseCode: 404}, nil
	}
	return &models.ApiResponse{Status: "1", Msg: "Success", ResponseCode: 200, Movies: &st}, nil

}

func (r *TestTaskRepo) NewMovie(ctx context.Context, tconst, titletype, primarytitle, runtimeminutes, genres string) (*models.ApiResponse, error) {
	b := models.Movies{
		Tconst:         tconst,
		Titletype:      titletype,
		Primarytitle:   primarytitle,
		Runtimeminutes: runtimeminutes,
		Genres:         genres,
	}
	strconv.ParseInt(runtimeminutes, 10, 64)
	if err := r.DBConn.Create(&b); err.Error != nil {
		return &models.ApiResponse{Status: "0", Msg: "Error in creating a new movie", ResponseCode: 404}, nil
	}
	return &models.ApiResponse{Status: "1", Msg: "Success", ResponseCode: 200}, nil

}
func (r *TestTaskRepo) MoviesByAverageRating(ctx context.Context) (*models.ApiResponse, error) {
	mvr := []models.MoviesByAverageRating{}

	query := fmt.Sprintf("select mv.tconst,mv.primarytitle,mv.genres,ra.averagerating from movies as mv " +
		" left join ratings as ra on mv.tconst=ra.tconst where ra.averagerating > '6.0' order by averagerating desc")

	if err := r.DBConn.Raw(query).Scan(&mvr); err.Error != nil {
		return &models.ApiResponse{Status: "0", Msg: "No data Found", ResponseCode: 404}, nil
	}
	return &models.ApiResponse{Status: "1", Msg: "Success", ResponseCode: 200, AverageRating: &mvr}, nil

}
func (r *TestTaskRepo) UpdateRunTimes(ctx context.Context) (*models.ApiResponse, error) {
	// urt := []models.GetMovies{}

	if err := r.DBConn.Exec(`UPDATE movies
	SET runtimeMinutes = 
	  CASE
		WHEN genres = 'Documentary' THEN runtimeMinutes + 15
		WHEN genres = 'Animation' THEN runtimeMinutes + 30
		ELSE runtimeMinutes + 45
	  END;`); err.Error != nil {
		return &models.ApiResponse{Status: "0", Msg: "Sorry! Data Can't be Updated", ResponseCode: 404}, nil
	}
	return &models.ApiResponse{Status: "1", Msg: "Runtime Updated Successfully", ResponseCode: 200}, nil

}

func (r *TestTaskRepo) MovieBySubtotal(ctx context.Context) (*models.ApiResponse, error) {
	mbs := []models.MovieBySubtotal{}

	query := fmt.Sprintf("SELECT genres, primaryTitle, SUM(numVotes::integer) AS numVotes " +
		"FROM movies mv left join ratings rt on mv.tconst=rt.tconst " +
		"GROUP BY ROLLUP (genres, primaryTitle) " +
		"ORDER BY genres ASC, primaryTitle ASC NULLS LAST")

	if err := r.DBConn.Raw(query).Scan(&mbs); err.Error != nil {
		return &models.ApiResponse{Status: "0", Msg: "No data Found", ResponseCode: 404}, nil
	}
	return &models.ApiResponse{Status: "1", Msg: "Success", ResponseCode: 200, MoviesBySubtotal: &mbs}, nil

}
