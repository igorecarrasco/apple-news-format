package components

import (
	"fmt"
)

type HeadingLevel string

const (
	HEADINGLEVELNONE HeadingLevel = ""
	HEADINGLEVEL1    HeadingLevel = "1"
	HEADINGLEVEL2    HeadingLevel = "2"
	HEADINGLEVEL3    HeadingLevel = "3"
	HEADINGLEVEL4    HeadingLevel = "4"
	HEADINGLEVEL5    HeadingLevel = "5"
	HEADINGLEVEL6    HeadingLevel = "6"
)

type HeadingStruct struct {
	TextStruct
}

func NewHeadingWithLevel(level HeadingLevel) *HeadingStruct {
	h := HeadingStruct{
		TextStruct: TextStruct{
			componentStruct: componentStruct{},
		},
	}
	h.SetRole(fmt.Sprintf("heading%s", level))
	return &h
}
