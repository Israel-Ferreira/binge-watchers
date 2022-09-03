package models

import "time"

type Model struct {
	ID        uint64     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type TvShow struct {
	Model
	Title      string `json:"title"`
	IsFinished bool   `json:"is_finished"`
	Category   string `json:"category"`
	LaunchYear string `json:"launch_year"`
	ShowRunner string `json:"showrunner"`
}
