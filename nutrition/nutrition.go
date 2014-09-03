package nutrition

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/brnv/nutrition-helper/calculator"
	"github.com/fuxiaohei/jx"
	"github.com/gocraft/web"
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
	Facts           calculator.NutritionFacts
	DailyNeed       calculator.NutritionFacts
	DailyPercentage calculator.NutritionFacts
}

func Eat(request *web.Request) {
	meal := StorageRecord{}

	request.ParseForm()
	form := request.Form

	meal.Time = time.Now().Unix()
	meal.Type = form.Get("typeInput")
	meal.Name = form.Get("productInput")
	meal.Weight = form.Get("weightInput")

	meal.Facts.Protein, _ = strconv.ParseFloat(form.Get("proteinInput"), 2)
	meal.Facts.Fat, _ = strconv.ParseFloat(form.Get("fatsInput"), 2)
	meal.Facts.Carbohydrate, _ = strconv.ParseFloat(form.Get("carbohydratesInput"), 2)
	meal.Facts.Calories, _ = strconv.ParseFloat(form.Get("caloriesInput"), 2)
	meal.Facts.Weight, _ = strconv.ParseFloat(form.Get("weightInput"), 2)

	meal.DailyNeed = dailyNeed
	meal.DailyPercentage = calculator.CalculatePercentage(facts, dailyNeed)

	//insert(meal)

	dirname := strings.TrimRight(request.RemoteAddr, ":1234567890")
	os.Mkdir(dirname, 0755)

	filename := fmt.Sprintf("./%v/%v%v%v", dirname, time.Now().Year(),
		time.Now().Month(), time.Now().Day())

	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)

	ttime := fmt.Sprintf("%02v:%02v:%02v", time.Now().Hour(), time.Now().Minute(),
		time.Now().Second())

	f.Write([]byte(ttime + "\t" + form.Get("typeInput") +
		"\t" + form.Get("productInput") + "\t" + form.Get("weightInput") +
		"\t" + fmt.Sprintf("%s", meal.Facts) + "\t" + fmt.Sprintf("%s", dailyNeed) +
		"\t" + fmt.Sprintf("%s",
		calculator.CalculatePercentage(meal.Facts, dailyNeed)) + "\n"))
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
