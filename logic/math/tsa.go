package math

import (
	"math"
	"path/filepath"
)

var tsaTplDir = filepath.Join(tplDir, "tsa")

// TSAPythagoreanTheorem outputs the proof and answer of finding the hypotenuse using the pythagorean theorem.
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
		round(a, .5, 2),
		round(b, .5, 2),
		round(aSqr, .5, 2),
		round(bSqr, .5, 2),
		round(ab, .5, 2),
		round(answer, .5, 2),
	}
	tplFile := filepath.Join(tsaTplDir, "pythagorean_theorem.gohtml")
	return parseTemplate(tplFile, data)
}

// TsaCone outputs the proof and answer of finding the total surface area of a cone.
func TsaCone(radius float64, slantHeight float64) (string, error) {
	/*
		Find the Total Surface Area of a cone, with the equation:
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
		round(radius, .5, 2),
		round(slantHeight, .5, 2),
		round(baseArea, .5, 2),
		round(curvedSurfaceArea, .5, 2),
		round(tsa, .5, 2),
	}
	tplFile := filepath.Join(tsaTplDir, "cone.gohtml")
	return parseTemplate(tplFile, data)
}

// TsaCube outputs the proof and answer of finding the total surface area of a cube.
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
		round(length, .5, 2),
		round(lengthSqr, .5, 2),
		round(tsa, .5, 2),
	}
	tplFile := filepath.Join(tsaTplDir, "cube.gohtml")
	return parseTemplate(tplFile, data)
}

// TsaCylinder outputs the proof and answer of finding the total surface area of a cylinder.
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
		round(radius, .5, 2),
		round(height, .5, 2),
		round(radiusSqr, .5, 2),
		round(circlesArea, .5, 2),
		round(rh, .5, 2),
		round(curvedSurfaceArea, .5, 2),
		round(tsa, .5, 2),
	}
	tplFile := filepath.Join(tsaTplDir, "cylinder.gohtml")
	return parseTemplate(tplFile, data)
}

// TsaRectangularPrism outputs the proof and answer of finding the total surface area of a rectangular prism.
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
		round(height, .5, 2),
		round(length, .5, 2),
		round(width, .5, 2),
		round(add, .5, 2),
		round(tsa, .5, 2),
	}
	tplFile := filepath.Join(tsaTplDir, "rectangular_prism.gohtml")
	return parseTemplate(tplFile, data)
}

// TsaSphere outputs the proof and answer of finding the total surface area of a sphere.
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
		round(radius, .5, 2),
		round(tsa, .5, 2),
	}
	tplFile := filepath.Join(tsaTplDir, "sphere.gohtml")
	return parseTemplate(tplFile, data)
}

// TsaSquareBaseTriangle outputs the proof and answer of finding the total surface area of a square base triangle.
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
		round(baseLength, .5, 2),
		round(height, .5, 2),
		round(squareArea, .5, 2),
		round(fourTrianglesArea, .5, 2),
		round(tsa, .5, 2),
	}
	tplFile := filepath.Join(tsaTplDir, "square_base_triangle.gohtml")
	return parseTemplate(tplFile, data)
}
