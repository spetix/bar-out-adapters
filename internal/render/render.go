package render

type RenderOptionsImpl struct {
	label           *string
	format          *string
	foregroundColor *string
	backgroundColor *string
}

func (r *RenderOptionsImpl) Label() string {
	return *r.label
}

func (r *RenderOptionsImpl) SetLabel(l *string) {
	r.label = l
}

func (r *RenderOptionsImpl) Format() string {
	return *r.format
}

func (r *RenderOptionsImpl) SetFormat(f *string) {
	r.format = f
}

func (r *RenderOptionsImpl) ForegroundColor() string {
	return *r.foregroundColor
}

func (r *RenderOptionsImpl) SetForegroundColor(f *string) {
	r.foregroundColor = f
}

func (r *RenderOptionsImpl) BackgroundColor() string {
	return *r.backgroundColor
}

func (r *RenderOptionsImpl) SetBackgroundColor(b *string) {
	r.backgroundColor = b
}

func NewRenderOptionsImpl() *RenderOptionsImpl {
	label := "🎄"
	format := ""
	foregroundColor := "#ffffff"
	backgroundColor := "#000000"
	return &RenderOptionsImpl{
		label:           &label,
		format:          &format,
		foregroundColor: &foregroundColor,
		backgroundColor: &backgroundColor,
	}
}
