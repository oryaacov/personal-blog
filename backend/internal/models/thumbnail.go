package models

type Thumbnail struct {
	Id          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Image       string `json:"img" bson:"image"`
	Summary     string `json:"summary" bson:"summary"`
	PublishedAt uint   `json:"date" bson:"publish_date"`
	Priority    uint   `json:"priority" bson:"priority"`
}
