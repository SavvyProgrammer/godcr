package materialplus

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"

	"github.com/raedahgroup/godcr-gio/ui/values"
)

// ProgressBar indicates the progress of a process. Height defines the thickness of the progressbar,
// BackgroundColor defines the color of the track, ProgressColor defines the color of the moving progress.
type ProgressBar struct {
	Height          int
	BackgroundColor color.RGBA
	ProgressColor   color.RGBA
}

// track lays out a rectangle to represent the level of progress yet to be completed.
func (p *ProgressBar) track(gtx *layout.Context) {
	borderedRectangle(gtx, values.ProgressBarGray, gtx.Constraints.Width.Max, p.Height)
}

// value lays out a rectangle to represent the level of progress that has been completed.
func (p *ProgressBar) value(gtx *layout.Context, progress float64) {
	width := progress / 100 * float64(gtx.Constraints.Width.Max)
	if width > float64(gtx.Constraints.Width.Max) {
		width = float64(gtx.Constraints.Width.Max)
	}
	borderedRectangle(gtx, p.ProgressColor, int(width), p.Height)
}

// borderedRectangle defines the dimensions of the rectangle, draws it and adds color it using the Fill method.
func borderedRectangle(gtx *layout.Context, color color.RGBA, x, y int) {
	br := float32(y / 5)
	rect := f32.Rectangle{
		Max: f32.Point{
			X: float32(x),
			Y: float32(y),
		},
	}
	clip.Rect{
		Rect: rect,
		NE:   br, NW: br, SE: br, SW: br,
	}.Op(gtx.Ops).Add(gtx.Ops)

	Fill(gtx, color, x, y)
}

// Layout lays out the track and level of progress on each other.
func (p *ProgressBar) Layout(gtx *layout.Context, progress float64) {
	layout.Stack{}.Layout(gtx,
		layout.Stacked(func() {
			p.track(gtx)
			p.value(gtx, progress)
		}),
	)
}

// ProgressBar returns a new ProgressBar instance.
func (t *Theme) ProgressBar() *ProgressBar {
	return &ProgressBar{
		Height:          values.DefaultProgressBarHeight,
		BackgroundColor: values.ProgressBarGray,
		ProgressColor:   values.ProgressBarGreen,
	}
}