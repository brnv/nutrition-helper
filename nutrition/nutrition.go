package nutrition

type Facts struct {
	Proteins      float64 `json:proteins`
	Fats          float64 `json:fats`
	Carbohydrates float64 `json:carbohydrates`
	Calories      float64 `json:calories`
}
