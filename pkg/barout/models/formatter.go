package models

// Formatter is the interface that wraps methods to format a blocklet's output.
type Formatter interface {
	Render(s ...string) string
	Validate(f string) bool
}
