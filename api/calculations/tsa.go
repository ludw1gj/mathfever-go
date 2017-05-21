package calculations

import "math"

func TSAPythagoreanTheorem(a float64, b float64) (string, error) {
	/*
		Use the Pythagorean Theorem to find side lengths (a or b), or the hypotenuse (c) of a right-angle triangle.

		Pythagorean Theorem: right-angle triangle with side lengths (a, b) and hypotenuse (c).
		a^2 + b^2 = c^2

		Example:
		When a = 25 and b = 17, find c.
		c^2 	= 25^2 + 17^2
			= 625 + 289
			= 914
		c	= √914
			≈ 30.2
	*/
	aSqr := math.Pow(a, 2)
	bSqr := math.Pow(b, 2)
	ab := aSqr + bSqr
	answer := math.Sqrt(ab)

	data := struct {
		A      float64
		B      float64
		ASqr   float64
		BSqr   float64
		AB     float64
		Answer float64
	}{
		Round(a, .5, 2),
		Round(b, .5, 2),
		Round(aSqr, .5, 2),
		Round(bSqr, .5, 2),
		Round(ab, .5, 2),
		Round(answer, .5, 2),
	}
	return parseTemplate("./templates/calculations/tsa/pythagoreanTheorem.gohtml", data)
}

func TsaCone(radius float64, slantHeight float64) (string, error) {
	/*
		Find the Total Surface Area of a sphere, with the equation:
		Where S is Slant Height, r is radius
		TSA     = tsa of base (circle) + tsa of curved surface
			= πr^2 + πrS
			= πr(r + S)
	*/
	baseArea := math.Pi * math.Pow(radius, 2)
	curvedSurfaceArea := math.Pi * radius * slantHeight
	tsa := baseArea + curvedSurfaceArea

	data := struct {
		Radius            float64
		SlantHeight       float64
		BaseArea          float64
		CurvedSurfaceArea float64
		TSA               float64
	}{
		Round(radius, .5, 2),
		Round(slantHeight, .5, 2),
		Round(baseArea, .5, 2),
		Round(curvedSurfaceArea, .5, 2),
		Round(tsa, .5, 2),
	}
	return parseTemplate("./templates/calculations/tsa/tsaCone.gohtml", data)
}

func TsaCube(length float64) (string, error) {
	/*
		Find the Total Surface Area of a cube, with the equation:
		TSA = 6L^2
	*/
	lengthSqr := math.Pow(length, 2)
	tsa := 6 * lengthSqr

	data := struct {
		Length    float64
		LengthSqr float64
		TSA       float64
	}{
		Round(length, .5, 2),
		Round(lengthSqr, .5, 2),
		Round(tsa, .5, 2),
	}
	return parseTemplate("./templates/calculations/tsa/tsaCube.gohtml", data)
}

func TsaCylinder(radius float64, height float64) (string, error) {
	/*
		Find the Total Surface Area of a cylinder, with the equation:
		TSA = tsa of 2 circles + curved surface
		    = 2πr^2 + 2πrh
		    = 2πr(r + h)
	*/
	radiusSqr := math.Pow(radius, 2)
	circlesArea := 2 * math.Pi * radiusSqr
	rh := radius * height
	curvedSurfaceArea := 2 * math.Pi * rh
	tsa := circlesArea + curvedSurfaceArea

	data := struct {
		Radius            float64
		Height            float64
		RadiusSqr         float64
		CirclesArea       float64
		RH                float64
		CurvedSurfaceArea float64
		TSA               float64
	}{
		Round(radius, .5, 2),
		Round(height, .5, 2),
		Round(radiusSqr, .5, 2),
		Round(circlesArea, .5, 2),
		Round(rh, .5, 2),
		Round(curvedSurfaceArea, .5, 2),
		Round(tsa, .5, 2),
	}
	return parseTemplate("./templates/calculations/tsa/tsaCylinder.gohtml", data)
}

func TsaRectangularPrism(height float64, length float64, width float64) (string, error) {
	/*
		Find the Total Surface Area of a rectangular prism, with the equation:
	    	TSA = 2(wh + lw + lh)
	*/
	add := (width * height) + (length * width) + (length * height)
	tsa := 2 * add

	data := struct {
		Height float64
		Length float64
		Width  float64
		Add    float64
		TSA    float64
	}{
		Round(height, .5, 2),
		Round(length, .5, 2),
		Round(width, .5, 2),
		Round(add, .5, 2),
		Round(tsa, .5, 2),
	}
	return parseTemplate("./templates/calculations/tsa/tsaRectangularPrism.gohtml", data)
}

func TsaSphere(radius float64) (string, error) {
	/*
		Find the Total Surface Area of a sphere, with the equation:
	    	TSA = 4πr^2
	*/
	tsa := 4 * math.Pi * math.Pow(radius, 2)

	data := struct {
		Radius float64
		TSA    float64
	}{
		Round(radius, .5, 2),
		Round(tsa, .5, 2),
	}
	return parseTemplate("./templates/calculations/tsa/tsaSphere.gohtml", data)
}

func TsaSquareBaseTriangle(baseLength float64, height float64) (string, error) {
	/*
		Find the Total Surface Area of a square based triangle, with the equation:
		TSA = tsa of square + tsa of 4 triangles
		    = b^2 + 4 * (1/2bh)
		    = b^2 + 2bh
	*/
	squareArea := math.Pow(baseLength, 2)
	fourTrianglesArea := 2 * (baseLength * height)
	tsa := squareArea + fourTrianglesArea

	data := struct {
		BaseLength        float64
		Height            float64
		SquareArea        float64
		FourTrianglesArea float64
		TSA               float64
	}{
		Round(baseLength, .5, 2),
		Round(height, .5, 2),
		Round(squareArea, .5, 2),
		Round(fourTrianglesArea, .5, 2),
		Round(tsa, .5, 2),
	}
	return parseTemplate("./templates/calculations/tsa/tsaSquareBaseTriangle.gohtml", data)
}
