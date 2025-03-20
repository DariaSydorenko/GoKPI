package utils

import (
	"math"
)

const EFFICIENCY_OF_FLUE_GAS_PURIFICATION = 0.985
const EMISSION_INDEX = 0.0

func CalculateEmission(lowerCalorificValue, fractionOfFlyAsh, ash, combustibleSubstances float64) float64 {
	return (math.Pow(10, 6) / lowerCalorificValue) * fractionOfFlyAsh *
		(ash / (100 - combustibleSubstances)) *
		(1 - EFFICIENCY_OF_FLUE_GAS_PURIFICATION) + EMISSION_INDEX
}

func CalculateGrossEmission(emission, mass, lowerCalorificValue float64) float64 {
	return math.Pow(10, -6) * emission * lowerCalorificValue * mass
}
