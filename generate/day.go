package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
)

func createFileIfNotExists(path string) (*os.File, error) {
	_, err := os.Stat(path)
	if err == nil {
		return nil, fmt.Errorf("Exists")
	}
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type templateData struct {
	Day int
}

func generateDay(day int) {
	fmt.Println("Generating day", day)
	tmpl := template.Must(template.New("day").ParseFiles("gen/templates/day.tmpl", "gen/templates/day_test.tmpl"))
	day_f, err := createFileIfNotExists(fmt.Sprintf("day%02d/main.go", day))
	defer day_f.Close()
	if err != nil {
		log.Fatalln("Error creating file:", err)
	}
	err = tmpl.ExecuteTemplate(day_f, "day.tmpl", templateData{day})
	if err != nil {
		log.Fatalln("Error executing template:", err)
	}

	test_f, err := createFileIfNotExists(fmt.Sprintf("day%02d/main_test.go", day))
	defer test_f.Close()
	if err != nil {
		log.Fatalln("Error creating file:", err)
	}
	err = tmpl.ExecuteTemplate(test_f, "day_test.tmpl", templateData{day})
	if err != nil {
		log.Fatalln("Error executing template:", err)
	}
	fmt.Println("Successfully generated day!")
}

func main() {
	day := os.Args[1]
	intDay, err := strconv.Atoi(day)
	if err != nil {
		log.Fatal(err)
	}
	generateDay(intDay)
}
