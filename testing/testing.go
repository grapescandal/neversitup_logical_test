package testing

import (
	"fmt"
	"sort"
	"strconv"
)

func LengthCheck(input string) (isPass bool) {
	isPass = false

	if len(input) >= 6 {
		isPass = true
	}
	return isPass
}

func AdjacentDuplicateCheck(input string) (isPass bool) {
	isPass = true
	duplicateCounter := 0
	nextIndexToCheck := 0

	for i, c := range input {
		if i == 0 {
			continue
		}

		if nextIndexToCheck != 0 {
			if i < nextIndexToCheck {
				continue
			}
		}

		character := string(c)

		if character == string(input[i-1]) {
			duplicateCounter++
			nextIndexToCheck = i + 2
		}

		if duplicateCounter >= 2 {
			isPass = false
			break
		} else if duplicateCounter == 1 && nextIndexToCheck == i {
			duplicateCounter = 0
			nextIndexToCheck = 0
		}
	}

	return isPass
}

func SortingCheck(input string) (isPass bool) {
	isPass = true
	inputInt := []int{}

	for _, c := range input {
		intCharacter, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println("Error can't cast character to int")
			return isPass
		}

		inputInt = append(inputInt, intCharacter)
	}

	inputAscending := []int{}
	inputAscending = append(inputAscending, inputInt...)
	sort.SliceStable(inputAscending, func(i, j int) bool {
		return input[i] < input[j]
	})

	ascNextIndexToCheck := 0
	acsCounter := 0
	for i := range inputAscending {

		if i == len(inputAscending)-1 {
			break
		}

		if inputAscending[i]+1 == inputAscending[i+1] {
			acsCounter++
			ascNextIndexToCheck = i + 1
		} else if inputAscending[i]+1 != inputAscending[i+1] && i == ascNextIndexToCheck {
			acsCounter = 0
			ascNextIndexToCheck = 0
		}

		if acsCounter == 2 {
			isPass = false
			return isPass
		}
	}

	inputDescending := []int{}
	inputDescending = append(inputDescending, inputInt...)
	sort.SliceStable(inputDescending, func(i, j int) bool {
		return input[i] > input[j]
	})

	dscNextIndexToCheck := 0
	dscCounter := 0
	for i := range inputDescending {

		if i == len(inputDescending)-1 {
			break
		}

		if inputDescending[i]+1 == inputDescending[i+1] {
			dscCounter++
			dscNextIndexToCheck = i + 1
		} else if inputDescending[i]+1 != inputDescending[i+1] && i == dscNextIndexToCheck {
			dscCounter = 0
			dscNextIndexToCheck = 0
		}

		if dscCounter == 2 {
			isPass = false
			return isPass
		}
	}
	return isPass
}

func DuplicateCheck(input string) (isPass bool) {
	isPass = true
	duplicateCounter := 0
	nextIndexToCheck := 0

	for i, c := range input {
		if i == 0 {
			continue
		}

		if nextIndexToCheck != 0 {
			if i < nextIndexToCheck {
				continue
			}
		}

		character := string(c)

		if character == string(input[i-1]) {
			duplicateCounter++
			nextIndexToCheck = i + 2
		}

		if duplicateCounter > 2 {
			isPass = false
			break
		}
	}

	return isPass
}
