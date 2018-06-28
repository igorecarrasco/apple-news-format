package styles

type Color string

type StrokeStyleOption string

const (
	StrokeStyleValueSolid  StrokeStyleOption = "solid"
	StrokeStyleValueDashed StrokeStyleOption = "dashed"
	StrokeStyleValueDotted StrokeStyleOption = "dotted"
)

type Border struct {
	All    StrokeStyle `json:"strokeStyle,omitempty"`
	Bottom bool        `json:"bottom,omitempty"`
	Left   bool        `json:"left,omitempty"`
	Right  bool        `json:"right,omitempty"`
	Top    bool        `json:"top,omitempty"`
}

type StrokeStyle struct {
	Color Color             `json:"color,omitempty"`
	Style StrokeStyleOption `json:"style,omitempty"`
	Width string            `json:"width"`
}

type ComponentStyle struct {
	BackgroundColor Color  `json:"backgroundColor,omitempty"`
	Border          *Border `json:"border,omitempty"`
	Fill            *Fill   `json:"fill,omitempty"`
}
