package model

type Home struct {
	Name string `json:"name" form:"name" query:"name"`
	Date string `json:"date" form:"date" query:"date"`
}

type Rating struct {
	Score        float32
	TotalReviews int32
}
