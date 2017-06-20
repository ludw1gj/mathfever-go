package database

var (
	networking = category{
		"Networking",
		"/static/resource/image/category/networking.jpg",
		"Calculations that you might use for Networking/Computer Science: Binary to Decimal, Binary " +
			"to Hexadecimal, Decimal to Binary, Decimal to Hexadecimal, Hexadecimal to Binary, and " +
			"Hexadecimal to Decimal.",
	}
	numbers = category{
		"Primes and Factors",
		"/static/resource/image/category/numbers.jpg",
		"Calculations about numbers! Find Highest Common Factor, find Lowest Common Multiple, and " +
			"figuring out Prime Numbers.",
	}
	percentages = category{
		"Percentages",
		"/static/resource/image/category/percentages.jpg",
		"Calculations for percentages! Find the value from a percentage, find a percentage from a " +
			"value, or find the percentage change between two values.",
	}
	tsa = category{
		"Total Surface Area",
		"/static/resource/image/category/tsa.jpg",
		"Calculations that you might use for Total Surface Area: Pythagorean Theorem (also known as " +
			"Pythagoras's Theorem), Total Surface Area of Cone, Total Surface Area of Cube, Total Surface " +
			"Area of Cylinder, Total Surface Area of Rectangular Prism, Total Surface Area of Sphere, and " +
			"Total Surface Area of Square Based Triangle.",
	}
	categoryData = []category{
		networking,
		numbers,
		percentages,
		tsa,
	}
	categoriesData []categoryWithCalculations
)

// populate categoriesData
func init() {
	for _, categ := range GetAllCategories() {
		calcs, _ := GetCalculationsByCategoryName(categ.Name)

		categoriesData = append(categoriesData, categoryWithCalculations{
			categ,
			calcs,
		})
	}
}
