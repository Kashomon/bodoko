package core

import (
	"github.com/Kashomon/bodoko/core"
)

func FromMovetree(mt *core.Movetree) *Cropping {
	ints := mt.getIntersections()
	middle := ints / 2
	return FromPreset(ALL)
}
