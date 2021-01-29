package test

import (
	"testing"

	"github.com/hjd919/gom/lbs"
)

func TestLbs(t *testing.T) {
	b := lbs.NewLbs("aaa")
	b.GeocodeRegeo(1, 1)
}
