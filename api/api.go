// Package api contains the MathAPI interface and types that implement it. It is used for the execution of
// calculation math functions.
package api

// MathAPI is an interface which encompasses calculation-api types which implement the ExecuteMath method.
type MathAPI interface {
	// ExecuteMath returns the math function's string output and an error if validation
	// or template error occurs.
	ExecuteMath() (string, error)
}
