package nutrition

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/brnv/nutrition-helper/calculator"
)

var (
	weight         string
	nutritionFacts string
	facts          calculator.NutritionFacts
	dailyNorm      = calculator.NutritionFacts{
		Protein:      90,
		Fat:          90,
		Carbohydrate: 340,
		Calories:     2400,
	}
	reFacts = regexp.MustCompile("([0-9\\.]*)/([0-9\\.]*)/([0-9\\.]*)/([0-9\\.]*)")
)

func Eat(form url.Values) {
	t := time.Now().Unix()

	f, _ := os.OpenFile("diary", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)

	facts.Protein, _ = strconv.ParseFloat(form.Get("proteinInput"), 32)
	facts.Fat, _ = strconv.ParseFloat(form.Get("fatsInput"), 32)
	facts.Carbohydrate, _ = strconv.ParseFloat(form.Get("carbohydratesInput"), 32)
	facts.Calories, _ = strconv.ParseFloat(form.Get("caloriesInput"), 32)
	facts.Weight, _ = strconv.ParseFloat(form.Get("weightInput"), 32)

	f.Write([]byte(strconv.FormatInt(t, 10) + "\t" + form.Get("typeInput") +
		"\t" + form.Get("productInput") + "\t" + form.Get("weightInput") +
		"\t" + fmt.Sprintf("%s", facts) + "\t" + fmt.Sprintf("%s", dailyNorm) +
		"\t" + fmt.Sprintf("%s",
		calculator.CalculatePercentage(facts, dailyNorm)) + "\n"))

}
