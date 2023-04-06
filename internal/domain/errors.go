package domain

import "errors"

var (
	ErrNoRates    = errors.New("no rates to add")
	ErrDateFormat = errors.New("bad format - must be like '21.03.2021'")
)
