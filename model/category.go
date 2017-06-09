package model

type Category struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	URL         string `json:"url"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

var (
	networking = Category{
		"Networking",
		"networking",
		"/networking",
		"/assets/resources/images/category/networking.jpg",
		"Calculations that you might use for Networking/Computer Science: Binary to Decimal, Binary " +
			"to Hexadecimal, Decimal to Binary, Decimal to Hexadecimal, Hexadecimal to Binary, and " +
			"Hexadecimal to Decimal.",
	}
	numbers = Category{
		"Primes and Factors",
		"numbers",
		"/numbers",
		"/assets/resources/images/category/addition.jpg",
		"Calculations about numbers! Find Highest Common Factor, find Lowest Common Multiple, and " +
			"figuring out Prime Numbers.",
	}
	percentages = Category{
		"Percentages",
		"percentages",
		"/percentages",
		"/assets/resources/images/category/algebra.jpg",
		"Calculations for percentages! Find the value from a percentage, find a percentage from a " +
			"value, or find the percentage change between two values.",
	}
	tsa = Category{
		"Total Surface Area",
		"tsa",
		"/tsa",
		"/assets/resources/images/category/geometry.jpg",
		"Calculations that you might use for Total Surface Area: Pythagorean Theorem (also known as " +
			"Pythagoras's Theorem), Total Surface Area of Cone, Total Surface Area of Cube, Total Surface " +
			"Area of Cylinder, Total Surface Area of Rectangular Prism, Total Surface Area of Sphere, and " +
			"Total Surface Area of Square Based Triangle.",
	}
	CategoryData = []Category{
		networking,
		numbers,
		percentages,
		tsa,
	}
)
