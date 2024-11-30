package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	_ "github.com/joho/godotenv/autoload"
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
	tmpl := template.Must(template.New("day").ParseFiles("generate/templates/day.tmpl", "generate/templates/day_test.tmpl"))
	day_f, err := createFileIfNotExists(fmt.Sprintf("day%02d/main.go", day))
	if err != nil {
		log.Fatalln("Error creating file:", err)
	}
	defer day_f.Close()
	err = tmpl.ExecuteTemplate(day_f, "day.tmpl", templateData{day})
	if err != nil {
		log.Fatalln("Error executing template:", err)
	}

	test_f, err := createFileIfNotExists(fmt.Sprintf("day%02d/main_test.go", day))
	if err != nil {
		log.Fatalln("Error creating file:", err)
	}
	defer test_f.Close()
	err = tmpl.ExecuteTemplate(test_f, "day_test.tmpl", templateData{day})
	if err != nil {
		log.Fatalln("Error executing template:", err)
	}
	fmt.Println("Successfully generated day!")
}

func getInput(day int, cookie string) []byte {
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Fatal request: %s", err)
	}

	sessionCookie := http.Cookie{
		Name:  "session",
		Value: cookie,
	}
	req.AddCookie(&sessionCookie)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("Fatal request:", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Response body:", err)
	}

	return body
}

func main() {
	day := os.Args[1]

	intDay, err := strconv.Atoi(day)
	if err != nil {
		log.Fatal(err)
	}
	generateDay(intDay)

	os.WriteFile(fmt.Sprintf("day%02d/input/inp.txt", intDay), getInput(intDay, os.Getenv("AOC_SESSION")), fs.ModeDevice)
}
