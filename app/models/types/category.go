package types

// Category holds information about a category of calculations.
type Category struct {
	Name         string        `json:"name"`         // name of the category
	Slug         string        `json:"slug"`         // slug of the category
	URL          string        `json:"url"`          // url of the category
	ImageURL     string        `json:"image_url"`    // url of the image for category
	Description  string        `json:"description"`  // description of the category
	Calculations []Calculation `json:"calculations"` // the calculations belonging to the category
}
