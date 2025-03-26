package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*.go.tmpl
var templateFS embed.FS

// cd api/ && go generate ./...
func main() {
	var (
		entityName   = flag.String("entity", "", "Entity type name")
		templateCode = flag.String("template", "service", "Template basename")
	)
	flag.Parse()
	if *entityName == "" {
		fmt.Println("Required flags: -entity")
		flag.PrintDefaults()
		os.Exit(1)
	}
	// fill data
	pkgEntityName := strings.ToLower(*entityName)
	data := struct {
		PkgEntity string
		Entity    string
	}{
		Entity:    *entityName,
		PkgEntity: pkgEntityName,
	}
	// get template from file
	templateFile := fmt.Sprintf("templates/%s.go.tmpl", *templateCode)
	tmplContent, err := templateFS.ReadFile(templateFile)
	if err != nil {
		fmt.Printf("Error reading template file: %v\n", err)
		os.Exit(1)
	}
	tmpl, err := template.New("service").Parse(string(tmplContent))
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		os.Exit(1)
	}
	// create dir if it does not exist
	moduleRoot, err := findGoModRoot()
	if err != nil {
		fmt.Printf("Error getting root directory: %v\n", err)
		os.Exit(1)
	}
	output := filepath.Join(moduleRoot, "../", "api", "services", pkgEntityName+"ss", *templateCode+".go")
	err = os.MkdirAll(filepath.Dir(output), 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}
	// open file to write
	file, err := os.Create(output)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// execute template
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("generated service for %s at %s\n", *entityName, output)
}

func findGoModRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found")
		}
		dir = parent
	}
}
