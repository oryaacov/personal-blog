package models

type Article struct {
	Id          string `bson:"_id"`
	Href        string `bson:"href"`
	Title       string `bson:"title"`
	Image       string `bson:"image"`
	Summary     string `bson:"summary"`
	PublishedAt uint   `bson:"publishDate"`
	Body        string `bson:"body,omitempty"`
}
