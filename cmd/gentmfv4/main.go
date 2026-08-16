package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/goccy/go-yaml"
	"github.com/hesusruiz/tmforum/internal/jpath"
	"golang.org/x/tools/imports"
)

// This is a simple tool to process the Swagger files in the "swagger" directory
// and extract the mapping of last path part to management system and the routes.
// It assumes the Swagger files are in the format used by the TMForum APIs.
// It will print the mapping and the routes to the standard output in JSON format.

//go:embed routes.hbs
var routesTemplate string

var (
	managementToUpstream = map[string]string{}
	resourceToManagement = map[string]string{}
	resourceToFullPath   = map[string]string{}
)

func main() {

	// Visit all the JSON files in the "swagger" directory
	swaggerDir := "./www/oapiv4/oapiv4"

	// Read the directory entries
	dirEntries, err := os.ReadDir(swaggerDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory %s: %v\n", swaggerDir, err)
		os.Exit(1)
	}

	// Process each file in the directory
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			// Process the file
			filePath := path.Join(swaggerDir, dirEntry.Name())
			if !strings.HasSuffix(filePath, ".json") {
				// Skip non-JSON files
				continue
			}
			processOneFile(filePath)
		}
	}

	if true {
		os.Exit(0)
	}

	tmpl, err := template.New("routes").Parse(routesTemplate)
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, map[string]any{
		"ResourceToManagement":   resourceToManagement,
		"ResourceToStandardPath": resourceToFullPath,
		"ManagementToUpstream":   managementToUpstream,
	})
	if err != nil {
		panic(err)
	}

	// Adjust imports before writing the file
	out, err := imports.Process("perico.go", b.Bytes(), nil)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./perico.go", out, 0644)
	if err != nil {
		panic(err)
	}

}

func processOneFile(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Unmarshal the JSON content into a map
	var swagMap map[string]any
	err = yaml.Unmarshal(content, &swagMap)
	if err != nil {
		panic(err)
	}

	title := jpath.GetString(swagMap, "info.title")
	if len(title) == 0 {
		panic("info.title key not found or empty")
	}

	basePath := jpath.GetString(swagMap, "basePath")
	if len(basePath) == 0 {
		panic("basePath key not found or empty")
	}

	basePathTrimmed := strings.Trim(basePath, "/")

	basePathParts := strings.Split(basePathTrimmed, "/")
	if len(basePathParts) < 2 {
		panic("basePath does not contain at least 2 parts")
	}

	apiFamily := basePathParts[1]
	fmt.Println("API family:", apiFamily)

	// Get the "paths" key from the map, which has this form:
	// "paths": {
	//     "/catalog": {
	//         "get": {
	// ... (and other HTTP methods)

	paths := jpath.GetMap(swagMap, "paths")

	localResourceNames := map[string]bool{}

	// Iterate over the keys in the "paths" map
	for thePath := range paths {

		thePath = strings.Trim(thePath, "/")

		pathParts := strings.Split(thePath, "/")
		resourceName := pathParts[0]

		// Skip importJob and exportJob
		if resourceName == "importJob" || resourceName == "exportJob" {
			continue
		}

		// Skip hub and listener
		if resourceName == "hub" || resourceName == "listener" {
			continue
		}

		// Add to the set of canonicalized resource names
		localResourceNames[strings.ToLower(resourceName)] = true

		// Map of resource name to management system
		resourceToManagement[resourceName] = apiFamily

		// Map of resource name to full collection path
		resourceToFullPath[resourceName] = path.Join("/tmf-api", basePath, resourceName)

	}

	// ******************************
	// Get the definitions
	definitions := jpath.GetMap(swagMap, "definitions")

	for definitionName, v := range definitions {

		// Skip Event definitions
		if strings.Contains(definitionName, "ChangeEvent") || strings.Contains(definitionName, "DeleteEvent") || strings.Contains(definitionName, "CreateEvent") {
			continue
		}

		// Get all properties
		properties := jpath.GetMap(v, "properties")

		var found bool
		var propsSlice = []string{}
		for propertyName := range properties {
			propsSlice = append(propsSlice, propertyName)
			if propertyName == "name" {
				found = true
				break
			}
		}

		if !found {
			// fmt.Println("   NAME NOT REQUIRED FOR", definitionName, ">>>>", propsSlice)
		}

		if strings.HasSuffix(definitionName, "_Create") || strings.HasSuffix(definitionName, "_Update") {
			requiredAttributes := jpath.GetListString(v, "required")
			if len(requiredAttributes) > 0 {
				fmt.Println(" ", definitionName, "->", requiredAttributes)
			}
		}

	}

	// // fmt.Println(description)
	// fmt.Printf("- **%s**\n", apiFamily)

	// for resourceName := range localResourceNames {
	// 	fmt.Printf("  - %s\n", resourceName)
	// }

	// fmt.Println()

}
