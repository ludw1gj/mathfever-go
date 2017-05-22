package model

type Category struct {
	Name         string        `json:"name"`
	URL          string        `json:"url"`
	Image        string        `json:"image"`
	Calculations []Calculation `json:"calculation"`
}

var (
	CategoryData []Category

	networkingCategory = Category{
		"Networking",
		"/networking",
		"/public/img/category-card/networking.jpg",
		networkingCalculations(),
	}
	numbersCategory = Category{
		"Numbers",
		"/numbers",
		"/public/img/category-card/addition.jpg",
		numbersCalculations(),
	}
	percentagesCategory = Category{
		"Percentages",
		"/percentages",
		"/public/img/category-card/algebra.jpg",
		percentagesCalculations(),
	}
	tsaCategory = Category{
		"Total Surface Area",
		"/tsa",
		"/public/img/category-card/geometry-1044090.jpg",
		tsaCalculations(),
	}
)

func init() {
	CategoryData = append(CategoryData, networkingCategory, numbersCategory, percentagesCategory, tsaCategory)
}
