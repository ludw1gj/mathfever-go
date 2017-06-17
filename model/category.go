package model

// Calculation holds information about a category of calculations.
type Category struct {
	Name        string `json:"name"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

// CategoryWithCalculations hold a category and it's calculations.
type CategoryWithCalculations struct {
	Category     Category      `json:"category"`
	Calculations []Calculation `json:"calculations"`
}
