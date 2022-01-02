package models

type Article struct {
	Id          string `json:"id" bson:"_id"`
	Href        string `json:"sumhrefmary" bson:"href"`
	Title       string `json:"title" bson:"title"`
	Image       string `json:"img" bson:"image"`
	Summary     string `json:"summary" bson:"summary"`
	PublishedAt uint   `json:"date" bson:"publishDate"`
	Body        string `json:"body" bson:"body,omitempty"`
}
