package nutrition

type Facts struct {
	Proteins     float64 `json:protein`
	Fats         float64 `json:fat`
	Carbohydrate float64 `json:carbohydrates`
	Calories     float64 `json:calories`
}
