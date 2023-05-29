package model

type ApiResponse struct {
	Status           string                   `json:"status"`
	Msg              string                   `json:"msg"`
	ResponseCode     int64                    `json:"Response_Code"`
	Movies           *[]Movies                `json:"Get_Movies,omitempty"`
	AverageRating    *[]MoviesByAverageRating `json:"AverageRating,omitempty"`
	MoviesBySubtotal *[]MovieBySubtotal       `json:"movie_by_subtotal,omitempty"`
}
type Movies struct {
	Tconst         string `json:"tconst"`
	Titletype      string `json:"titletype"`
	Primarytitle   string `json:"primarytitle"`
	Runtimeminutes string `json:"runtimeminutes"`
	Genres         string `json:"genres"`
}
type MoviesByAverageRating struct {
	Tconst        string `json:"tconst"`
	Primarytitle  string `json:"PrimaryTitle"`
	Genres        string `json:"genres"`
	Averagerating string `json:"averageRating"`
}

type MovieBySubtotal struct {
	Genres       string `json:"genres"`
	Primarytitle string `json:"primaryTitle"`
	Numvotes     int64  `json:"numVotes"`
}
