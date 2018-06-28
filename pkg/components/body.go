package components


type bodyStruct struct {
	textStruct
}

func NewBody() *bodyStruct {
	b := bodyStruct{
		textStruct: textStruct{
			componentStruct: componentStruct{
			},
		},
	}
	b.SetRole("body")
	return &b
}