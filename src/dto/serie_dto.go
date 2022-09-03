package dto

type SerieDTO struct {
	Title      string `json:"title"`
	Category   string `json:"category"`
	LaunchYear string `json:"launchYear"`
	IsFinished bool   `json:"isFinished"`
}
