package lbs

import "github.com/hjd919/gom/lbs/gaode"

func NewLbs(key string) *gaode.Handle {
	return &gaode.Handle{
		Key: key,
	}
}
