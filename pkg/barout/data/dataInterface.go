package data


// Data is the interface that wraps methods to retrieve information about a blocklet's output.
// Client code can implement this interface to provide custom data for different blocklets.
type Data interface {
	Short() string
	Long() string
	BackgroundColor() string
	ForegroundColor() string
	Label() string
}
