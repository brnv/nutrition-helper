package nutrition

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/brnv/nutrition-helper/calculator"
	"github.com/fuxiaohei/jx"
)

var (
	storage        = initStorage()
	weight         string
	nutritionFacts string
	facts          calculator.NutritionFacts
	dailyNeed      = calculator.NutritionFacts{
		Protein:      90,
		Fat:          90,
		Carbohydrate: 340,
		Calories:     2400,
	}
	reFacts = regexp.MustCompile("([0-9\\.]*)/([0-9\\.]*)/([0-9\\.]*)/([0-9\\.]*)")
)

type StorageRecord struct {
	Id              int64 `jx:"pk-auto"`
	Time            int64
	Type            string
	Name            string
	Weight          string
	OriginalFacts   calculator.NutritionFacts
	DailyNeedFacts  calculator.NutritionFacts
	DailyPercentage calculator.NutritionFacts
}

func Eat(form url.Values) {
	meal := StorageRecord{}

	meal.Time = time.Now().Unix()
	meal.Type = form.Get("typeInput")
	meal.Name = form.Get("productInput")
	meal.Weight = form.Get("weightInput")

	meal.OriginalFacts.Protein, _ = strconv.ParseFloat(form.Get("proteinInput"), 2)
	meal.OriginalFacts.Fat, _ = strconv.ParseFloat(form.Get("fatsInput"), 2)
	meal.OriginalFacts.Carbohydrate, _ = strconv.ParseFloat(form.Get("carbohydratesInput"), 2)
	meal.OriginalFacts.Calories, _ = strconv.ParseFloat(form.Get("caloriesInput"), 2)
	meal.OriginalFacts.Weight, _ = strconv.ParseFloat(form.Get("weightInput"), 2)

	meal.DailyNeedFacts = dailyNeed
	meal.DailyPercentage = calculator.CalculatePercentage(facts, dailyNeed)

	//insert(meal)

	filename := fmt.Sprintf("%v%v%v", time.Now().Year(), time.Now().Month(),
		time.Now().Day())

	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)

	f.Write([]byte(strconv.FormatInt(meal.Time, 10) + "\t" + form.Get("typeInput") +
		"\t" + form.Get("productInput") + "\t" + form.Get("weightInput") +
		"\t" + fmt.Sprintf("%s", meal.OriginalFacts) + "\t" + fmt.Sprintf("%s", dailyNeed) +
		"\t" + fmt.Sprintf("%s",
		calculator.CalculatePercentage(meal.OriginalFacts, dailyNeed)) + "\n"))

}

func Get(rec *StorageRecord) {
	r := storage.Get(rec)
	log.Println(r)

}

func insert(record StorageRecord) error {
	storage.Sync(new(StorageRecord))
	err := storage.Insert(record)
	if err != nil {
		log.Println("[error]", err)
	}
	return nil
}

func update() {

}

func initStorage() *jx.Storage {
	log.SetFlags(log.Lshortfile)
	storage, err := jx.NewStorage(".data")
	if err != nil {
		log.Println("[error]", err)
	}
	return storage
}
