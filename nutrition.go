package main

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/brnv/nutrition-helper/calculator"
)

var (
	weight         string
	nutritionFacts string
	dailyNorm      string
	facts          calculator.NutritionFacts
	norm           calculator.NutritionFacts
	reFacts        = regexp.MustCompile("([0-9\\.]*)/([0-9\\.]*)/([0-9\\.]*)/([0-9\\.]*)")
)

func Eat() {
	flag.StringVar(&weight, "w", "100",
		"weight of eaten meal in grams")
	flag.StringVar(&nutritionFacts, "f", "0/0/0/0",
		"proteins/fats/carbohydrate/calories")
	flag.StringVar(&dailyNorm, "d", "90/90/340/2400",
		"proteins/fats/carbohydrate/calories")
	flag.Parse()

	facts = calculator.ParseFacts(reFacts.FindStringSubmatch(nutritionFacts))
	norm = calculator.ParseFacts(reFacts.FindStringSubmatch(dailyNorm))
	facts.Weight = calculator.ParseFloat(weight)

	fmt.Printf("%s", calculator.CalculatePercentage(facts, norm))
}
