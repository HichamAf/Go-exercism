package thefarm

import (
    "fmt"
    "errors"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
    r1, err := fc.FodderAmount(numberOfCows); if err != nil {
		return 0.0, err //ors.New("amount could not be determined")
	}
    r2, err := fc.FatteningFactor(); if err != nil {
		return 0.0, err //ors.New("factor could not be determined")
	}

    var result = r1 * r2 / float64(numberOfCows)
    return result, nil
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
    if numberOfCows > 0 {
    	result, err := DivideFood(fc, numberOfCows)
    	return result, err
    }
    return 0.0, errors.New("invalid number of cows")
}
// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
    NumberOfCows int
    Message string
    Err error
}

func (i *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", i.NumberOfCows, i.Message)
}

func ValidateNumberOfCows(numberOfCows int) error {
    if numberOfCows == 0 {
    	return &InvalidCowsError {
            NumberOfCows: numberOfCows,
            Message: "no cows don't need food",
        }
    } else if numberOfCows < 0 {
    	return &InvalidCowsError {
            NumberOfCows: numberOfCows,
            Message: "there are no negative cows",
        }
    }
    return nil
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
