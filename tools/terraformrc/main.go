package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	// TODO make configurable
	output := ".terraformrc"
	templatePath := "../../.terraformrc.tmpl"
	binPath, err := filepath.Abs("../../build/")
	if err != nil {
		panic(err)
	}

	// Set template replacement
	vars := make(map[string]interface{})
	vars["ProviderPath"] = binPath

	// parse the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	// create a new file
	file, _ := os.Create(output)
	defer file.Close()

	// apply the template to the vars map and write the result to file.
	tmpl.Execute(file, vars)

	log.Printf("Generated %s with %s", output, binPath)
}
