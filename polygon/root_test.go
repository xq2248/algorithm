package polygon

import (
	"testing"
)

func TestWalk(t *testing.T) {
	var sixPoints []Point = []Point{
		{1.0, 1.0},
		{1.0, 2.0},
		{3.0, 2.0},
		{3.0, 1.0},
		//{6.0, 4.0},
		//{6.0, 1.0},
	}
	Walk(sixPoints, 3)
}
