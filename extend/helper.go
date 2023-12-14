package extend

import "errors"

var Helper HelperTools

type HelperTools struct{}

func (h *HelperTools) err() error {
	return errors.New("ffff")
}
