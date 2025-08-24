package barout

type Data interface {
	Short() string
	Long() string
	BackgroundColor() string
	ForegroundColor() string
	Label() string
}
