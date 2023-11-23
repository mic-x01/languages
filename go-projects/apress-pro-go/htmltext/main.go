package main

import (
	"html/template"
	"os"
	// "os"
)

func GetCategories(products []Product) (categories []template.HTML) {
	catMap := map[string]string{}
	for _, p := range products {
		if catMap[p.Category] == "" {
			catMap[p.Category] = p.Category
			categories = append(categories, "<b>p.Category</b>")
		}
	}
	return
}

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}

func main() {
	allTemplates := template.New("allTemplates")
	allTemplates.Funcs(map[string]interface{}{
		"getCats": GetCategories,
	})
	allTemplates, err := allTemplates.ParseGlob("templates/*.html")
	if err == nil {
		selectedTemplated := allTemplates.Lookup("mainTemplate")
		err = Exec(selectedTemplated)
	}
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
