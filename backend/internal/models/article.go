package models

import "time"

type Article struct {
	id          string    `db:"id"`
	title       string    `db:"title"`
	image       string    `db:"image"`
	summary     string    `db:"summary"`
	publishedAt time.Time `db:"created_at"`
	body        string    `db:"body"`
}
