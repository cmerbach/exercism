package thefarm

import (
	"errors";
	"fmt"
)


func DivideFood(calculator FodderCalculator, numCows int) (float64, error) {

	totalFodder, err := calculator.FodderAmount(numCows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	foodPerCow := (totalFodder * fatteningFactor) / float64(numCows)

	return foodPerCow, nil
	
}

func ValidateInputAndDivideFood(calculator FodderCalculator, numCows int) (float64, error) {
 
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(calculator, numCows)

}


type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (e InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.message)
}

func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{
			numberOfCows: numCows,
			message:      "there are no negative cows",
		}
	}
	
	if numCows == 0 {
		return &InvalidCowsError{
			numberOfCows: numCows,
			message:      "no cows don't need food",
		}
	}
	
	return nil
}

