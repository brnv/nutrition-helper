package calculator

import "github.com/brnv/nutrition-helper/nutrition"

var weight float64

func DailyPercentage(consumedWeight float64, consumed nutrition.Facts,
	needed nutrition.Facts) nutrition.Facts {

	setWeight(consumedWeight)

	return nutrition.Facts{
		Proteins:      getPercentage(consumed.Proteins, needed.Proteins),
		Fats:          getPercentage(consumed.Fats, needed.Fats),
		Carbohydrates: getPercentage(consumed.Carbohydrates, needed.Carbohydrates),
		Calories:      getPercentage(consumed.Calories, needed.Calories),
	}
}

func getPercentage(consumed float64, needed float64) float64 {
	return consumed / 100 * weight / needed * 100
}

func setWeight(value float64) {
	weight = value
}
