package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

type ProjectOptions struct {
	CSSFramework string
}

var newCmd = &cobra.Command{
	Use:   "new  [project-name]",
	Short: "Create a new HTMX project",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]

		// picking a css framwork
		opts, err := promptProjectOptions()
		if err != nil {
			fmt.Println("[ship-htmx-app]: Prompt cancelled by user")
		}

		// project structure
		if err := scaffoldProject(projectName, opts); err != nil {
			fmt.Println("[ship-htmx-app]: Failed to scaffold project")
		}
		return nil
	},
}

func promptProjectOptions() (ProjectOptions, error) {
	var opts ProjectOptions
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a CSS framwork").
				Options(
					huh.NewOption("Plain", "none(I will do it myself"),
					huh.NewOption("Pico CSS (minimal)", "pico"),
					huh.NewOption("Tailwind CSS", "tailwind"),
				).
				Value(&opts.CSSFramework),
		),
	)

	// will cancel if user presses ctrl+c
	if err := form.Run(); err != nil {
		return ProjectOptions{}, err
	}

	return opts, nil
}

func scaffoldProject(projectName string, opts ProjectOptions) error {
	if err := createProjectStructure(projectName); err != nil {
		fmt.Errorf("[ship-htmx-app]: failed to create project structure", err)
	}

	if err := applyCSSFramework(projectName, opts.CSSFramework); err != nil {
		fmt.Errorf("[ship-htmx-app]: failed to add CSS framwork", err)
	}
	return nil
}

func createProjectStructure(projectName string) error {
	dirs := []string{
		projectName,
		projectName + "/src",
		projectName + "/src/public",
		projectName + "/src/public/css",
		projectName + "/src/public/js",
		projectName + "/src/public/partials",
	}

	for _, dir := range dirs {
		// @dev: 0755 is needed for user permission on unix
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Errorf("[ship-htmx-app]: Failed to create directory %s : %w", dir, err)
		}
	}

	files := map[string]string{
		projectName + "/package.json":             makePackageJSON(projectName),
		projectName + "/server.js":                setUpServer(),
		projectName + "/src/public/indes.html":    makeIndexHTML(projectName),
		projectName + "/src/public/css/style.css": "",
	}

	for path, content := range files {
		//@dev : 0644 is needed for user permission on unix
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Errorf("[ship-htmx-app]: failed to create files %s: %n", path, err)
		}
	}
	return nil
}

func makeIndexHTML(projectName string) string {
	return `
		<!DOCTYPE html>
		<html lang="en">
			<head>
  			<meta charset="UTF-8" />
  			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
  			<title>` + projectName + `</title>
  			<script src="https://unpkg.com/htmx.org@1.9.10"></script>
			</head>
		<body>
  		<h1>` + projectName + `</h1>
  			<button hx-get="/hello" hx-target="#result">Click me</button>
  			<div id="result"></div>
				</body>
		</html>
	`
}

func makePackageJSON(projectName string) string {
	return `
	{
  	"name": "` + projectName + `",
  	"version": "0.1.0",
  	"scripts": {
    	"start": "node server.js",
    	"dev": "nodemon server.js"
  	},
  	"dependencies": {
    	"express": "^4.18.2"
  	},
  	"devDependencies": {
    	"nodemon": "^3.0.1"
  	}
	}
	`
}

func setUpServer() string {
	return `
					const express = require('express')
					
					const app = express()
					const port = process.env.PORT || 3000

					app.use(express.static('public'))
					app.use(express.urlencoded({ extended: true }))

					app.get('/', (req, res) => {
  					res.sendFile(__dirname + '/views/index.html')
					})

					app.listen(port, () => {
  					console.log('Server running at http://localhost:' + port)
					}) 
	`
}

func cssLink(framwork string) string {
	switch framwork {
	case "pico":
		return `<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css">`
	case "tailwind":
		return `<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/tailwindcss@4/tailwind.min.css">`
	default:
		return `<link rel="stylesheet" href="/css/style.css">`
	}
}

func applyCSSFramework(projectName string, framework string) error {
	path := projectName + "/src/public/index.html"

	content, err := os.ReadFile(path)

	if err != nil {
		fmt.Errorf("Failed to read file : index.html")
	}

	updatedContent := strings.Replace(string(content), "</head>", cssLink(framework)+"\n</head>", 1)

	if err := os.WriteFile(path, []byte(updatedContent), 0644); err != nil {
		fmt.Errorf("Failed to update index.html")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)
}
