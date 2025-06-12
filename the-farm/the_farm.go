package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	numCows int
}

var ErrNonScale = errors.New("non-scale error")
var ErrDivByZero = errors.New("division by zero")
var ErrNegativeFodder = errors.New("negative fodder")

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %v cows", e.numCows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, ErrDivByZero
	}

	// weight, err := weightFodder.FodderAmount()
	// if weight < 0 {
	// 	switch err {
	// 	case nil, ErrScaleMalfunction:
	// 		return 0, ErrNegativeFodder
	// 	default:
	// 		return 0, ErrNonScale
	// 	}
	// }

	// if cows < 0 {
	// 	return 0.0, &SillyNephewError{cows}
	// }
	// switch err {
	// case ErrScaleMalfunction:
	// 	return 2.0 * weight / float64(cows), nil
	// case nil:
	// 	return weight / float64(cows), nil
	// default:
	// 	return 0, ErrNonScale
	// }

	fodder, err := weightFodder.FodderAmount()
	cowNegative := cows < 0
	fodNegative := fodder < 0
	switch err {
	case ErrScaleMalfunction:
		switch {
		case fodNegative:
			return 0, ErrNegativeFodder
		case cowNegative:
			return 0, &SillyNephewError{cows}
		default:
			return 2.0 * fodder / float64(cows), nil
		}
	case nil:
		switch {
		case fodNegative:
			return 0, ErrNegativeFodder
		case cowNegative:
			return 0, &SillyNephewError{cows}
		default:
			return fodder / float64(cows), nil
		}
	default:
		switch {
		case fodNegative:
			return 0, ErrNonScale
		case cowNegative:
			return 0, &SillyNephewError{cows}
		default:
			return 0, ErrNonScale
		}
	}
}
