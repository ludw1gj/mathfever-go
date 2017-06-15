package model

type Category struct {
	Name        string `json:"name"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

type CategoryWithCalculations struct {
	Category     Category      `json:"category"`
	Calculations []Calculation `json:"calculations"`
}
