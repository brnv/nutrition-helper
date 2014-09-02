package calculator

import (
	"fmt"
	"log"
	"strconv"
)

type NutritionFacts struct {
	Protein      float64 `json:protein`
	Fat          float64 `json:fat`
	Carbohydrate float64 `json:carbohydrates`
	Calories     float64 `json:calories`
	Weight       float64 `json:weight`
}

func (f NutritionFacts) String() string {
	return fmt.Sprintf("%2.2f/%2.2f/%2.2f/%2.2f/", f.Protein, f.Fat,
		f.Carbohydrate, f.Calories)
}

func CalculatePercentage(part NutritionFacts, from NutritionFacts) NutritionFacts {
	percentage := func(weight float64, fact float64, dailyFact float64) float64 {
		return fact / 100 * weight / dailyFact * 100
	}
	calculated := NutritionFacts{}
	calculated.Protein = percentage(part.Weight, part.Protein, from.Protein)
	calculated.Fat = percentage(part.Weight, part.Fat, from.Fat)
	calculated.Carbohydrate = percentage(part.Weight, part.Carbohydrate, from.Carbohydrate)
	calculated.Calories = percentage(part.Weight, part.Calories, from.Calories)
	return calculated
}

func ParseFacts(from []string) NutritionFacts {
	return NutritionFacts{
		Protein:      ParseFloat(from[1]),
		Fat:          ParseFloat(from[2]),
		Carbohydrate: ParseFloat(from[3]),
		Calories:     ParseFloat(from[4]),
	}
}

func ParseFloat(str string) float64 {
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Println("[error]", err)
	}
	return float
}
