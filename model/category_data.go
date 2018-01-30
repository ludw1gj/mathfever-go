package model

var (
	categories = []Category{
		{
			"Networking",
			"/static/resource/images/category/networking.jpg",
			"Calculations that you might use for Networking/Computer Science: Binary to Decimal, Binary " +
				"to Hexadecimal, Decimal to Binary, Decimal to Hexadecimal, Hexadecimal to Binary, and " +
				"Hexadecimal to Decimal.",
		},
		{
			"Primes and Factors",
			"/static/resource/images/category/numbers.jpg",
			"Calculations about numbers! Find Highest Common Factor, find Lowest Common Multiple, and " +
				"figuring out Prime Numbers.",
		},
		{
			"Percentages",
			"/static/resource/images/category/percentages.jpg",
			"Calculations for percentages! Find the value from a percentage, find a percentage from a " +
				"value, or find the percentage change between two values.",
		},
		{
			"Total Surface Area",
			"/static/resource/images/category/tsa.jpg",
			"Calculations that you might use for Total Surface Area: Pythagorean Theorem (also known as " +
				"Pythagoras's Theorem), Total Surface Area of Cone, Total Surface Area of Cube, Total Surface " +
				"Area of Cylinder, Total Surface Area of Rectangular Prism, Total Surface Area of Sphere, and " +
				"Total Surface Area of Square Based Triangle.",
		},
	}
	categoriesWithCalculations []CategoryWithCalculations
)

func init() {
	for _, category := range GetAllCategories() {
		calculations, _ := GetCalculationsByCategoryName(category.Name)

		categoriesWithCalculations = append(categoriesWithCalculations, CategoryWithCalculations{
			category,
			calculations,
		})
	}
}
