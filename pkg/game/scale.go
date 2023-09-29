package game

import (
	"github.com/GodsBoss/cyber-revolution-2789/pkg/scale"
)

// TODO: This stuff should be moved into gggg somehow.

type scaler interface {
	Scale() int
	Recalculate(availableWidth, availableHeight int)
	RealSize() (realWidth, realHeight int)
}

func createScaler() scaler {
	return &scale.ByInteger{
		UnscaledWidth:    320,
		UnscaledHeight:   200,
		HorizontalMargin: 20,
		VerticalMargin:   20,
	}
}
