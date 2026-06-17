package models

// RenderOptions is the interface that wraps methods to retrieve rendering options for a blocklet's output.
type RenderOptions interface {
	Label() string
	Format() string
	ForegroundColor() string
	BackgroundColor() string
}
