package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
)

type (
	NutritionFacts struct {
		Protein      float64
		Fat          float64
		Carbohydrate float64
		Calories     float64
	}
	Meal struct {
		Weight float64
		Name   string
		NutritionFacts
	}
)

func (f *NutritionFacts) String() string {
	return fmt.Sprintf("%2.2f/%2.2f/%2.2f/%2.2f\n", f.Protein, f.Fat,
		f.Carbohydrate, f.Calories)
}

func (m *Meal) String() string {
	return fmt.Sprintf("%s/%2.2f/%s", m.Name, m.Weight, &m.NutritionFacts)
}

func main() {
	var mealNameInput string
	flag.StringVar(&mealNameInput, "n", "meal",
		"Name of product you eaten")

	var mealWeightInput string
	flag.StringVar(&mealWeightInput, "w", "0",
		"weight of eaten meal in grams")

	var mealFactsInput string
	flag.StringVar(&mealFactsInput, "f", "0/0/0/0",
		"protein/fat/carbohydrate/calories")

	var dailyNormInput string
	flag.StringVar(&dailyNormInput, "d", "90/90/340/2400",
		"protein/fat/carbohydrate/calories")

	flag.Parse()

	re := regexp.MustCompile("([0-9\\.]*)/([0-9\\.]*)/([0-9\\.]*)/([0-9\\.]*)")
	parsedFacts := re.FindStringSubmatch(mealFactsInput)
	parsedDailyNorm := re.FindStringSubmatch(dailyNormInput)

	meal := Meal{}
	meal.Weight, _ = strconv.ParseFloat(mealWeightInput, 1)
	meal.Name = mealNameInput
	parseFacts(&meal.NutritionFacts, parsedFacts)

	var dailyNorm = NutritionFacts{}
	parseFacts(&dailyNorm, parsedDailyNorm)

	calculated := calculatePercentage(meal, dailyNorm)
	meal.NutritionFacts = calculated
	fmt.Print(&meal)
}

func calculatePercentage(meal Meal, dailyNorm NutritionFacts) NutritionFacts {
	percentage := func(weight float64, fact float64, dailyFact float64) float64 {
		return weight * fact / 100 / dailyFact * 100
	}
	calculated := NutritionFacts{}
	calculated.Protein = percentage(meal.Weight, meal.Protein, dailyNorm.Protein)
	calculated.Fat = percentage(meal.Weight, meal.Fat, dailyNorm.Fat)
	calculated.Carbohydrate = percentage(meal.Weight, meal.Carbohydrate, dailyNorm.Carbohydrate)
	calculated.Calories = percentage(meal.Weight, meal.Calories, dailyNorm.Calories)

	return calculated
}

func parseFacts(facts *NutritionFacts, from []string) {
	facts.Protein, _ = strconv.ParseFloat(from[1], 1)
	facts.Fat, _ = strconv.ParseFloat(from[2], 1)
	facts.Carbohydrate, _ = strconv.ParseFloat(from[3], 1)
	facts.Calories, _ = strconv.ParseFloat(from[4], 1)

}
