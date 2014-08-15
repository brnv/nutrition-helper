package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type (
	NutritionFacts struct {
		Protein      float64
		Fat          float64
		Carbohydrate float64
		Calories     float64
		Weight       float64
	}
)

func (f NutritionFacts) String() string {
	return fmt.Sprintf("%2.2f/%2.2f/%2.2f/%2.2f\n", f.Protein, f.Fat,
		f.Carbohydrate, f.Calories)
}

var (
	weight         string
	nutritionFacts string
	dailyNorm      string
	facts          NutritionFacts
	norm           NutritionFacts
	reFacts        = regexp.MustCompile("([0-9\\.]*)/([0-9\\.]*)/([0-9\\.]*)/([0-9\\.]*)")
)

func main() {
	flag.StringVar(&weight, "w", "100",
		"weight of eaten meal in grams")
	flag.StringVar(&nutritionFacts, "f", "0/0/0/0",
		"proteins/fats/carbohydrate/calories")
	flag.StringVar(&dailyNorm, "d", "90/90/340/2400",
		"proteins/fats/carbohydrate/calories")
	flag.Parse()

	facts = parseFacts(reFacts.FindStringSubmatch(nutritionFacts))
	norm = parseFacts(reFacts.FindStringSubmatch(dailyNorm))
	facts.Weight = parseFloat(weight)

	fmt.Printf("%s", calculatePercentage(facts, norm))
}

func calculatePercentage(part NutritionFacts, from NutritionFacts) NutritionFacts {
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

func parseFacts(from []string) NutritionFacts {
	return NutritionFacts{
		Protein:      parseFloat(from[1]),
		Fat:          parseFloat(from[2]),
		Carbohydrate: parseFloat(from[3]),
		Calories:     parseFloat(from[4]),
	}
}

func parseFloat(str string) float64 {
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Println("[error]", err)
	}
	return float
}
