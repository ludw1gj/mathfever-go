package model

// Calculation holds information about a category of calculations.
type Category struct {
	Name        string `json:"name"`        // name of category
	ImageURL    string `json:"image_url"`   // url of image for category
	Description string `json:"description"` // describes the category
}

// CategoryWithCalculations hold a category and it's calculations.
type CategoryWithCalculations struct {
	Category     Category      `json:"category"`     // the category
	Calculations []Calculation `json:"calculations"` // the calculations belonging to the category
}
