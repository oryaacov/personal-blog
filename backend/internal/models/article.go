package models

type Article struct {
	Id          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Body        string `json:"body" bson:"body"`
	PublishedAt uint   `json:"date" bson:"publish_date"`
}
