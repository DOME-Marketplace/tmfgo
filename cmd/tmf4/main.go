package main

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/hesusruiz/tmforum/types"
	"github.com/pb33f/libopenapi"
	v2 "github.com/pb33f/libopenapi/datamodel/high/v2"
	"go.yaml.in/yaml/v4"
)

// This is a simple tool to process the Swagger files in the "swagger" directory
// and extract the mapping of last path part to management system and the routes.
// It assumes the Swagger files are in the format used by the TMForum APIs.
// It will print the mapping and the routes to the standard output in JSON format.

const updateFile = true

func main() {

	// Visit all the JSON files in the "swagger" directory
	swaggerDir := "./www/oapiv4/oapiv4"

	// Read the directory entries
	dirEntries, err := os.ReadDir(swaggerDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory %s: %v\n", swaggerDir, err)
		os.Exit(1)
	}

	// This is the map where we will accumulate the information for generation of code
	resources := make(types.Resources)

	// Process each file in the directory (we assume all files are in the same directory)
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			// Process the file
			filePath := path.Join(swaggerDir, dirEntry.Name())
			if !strings.HasSuffix(filePath, ".json") {
				// Skip non-JSON files
				continue
			}
			resources = processOneFile(filePath, resources)
		}
	}

	// printStatistics(resources)

	generateDefinitions(resources)

}

func processOneFile(filePath string, resources types.Resources) types.Resources {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	document, err := libopenapi.NewDocument(content)
	if err != nil {
		panic(err)
	}

	// Build a Swagger 2.0 model from the document
	v2Model, err := document.BuildV2Model()
	if err != nil {
		panic(err)
	}

	// We will iterate through the paths, which are the routes and actions in our server
	pathItems := v2Model.Model.Paths.PathItems
	basePath := v2Model.Model.BasePath

	for path, pathItem := range pathItems.FromNewest() {

		if pathItem == nil {
			fmt.Println("pathItem is nil")
			continue
		}

		// For each path, we are interested only in the CREATE and UPDATE operations,
		// which are the ones sending some data in the body of the request.

		op := pathItem.Post
		if op != nil {
			resources = processCREATEorUPDATE("CREATE", path, op, resources, basePath)
		}

		op = pathItem.Patch
		if op != nil {
			resources = processCREATEorUPDATE("UPDATE", path, op, resources, basePath)
		}

	}

	return resources

}

func processCREATEorUPDATE(action string, path string, op *v2.Operation, resources types.Resources, basePath string) types.Resources {
	if op != nil {

		// We skip the operations that we do not have to implement
		id := op.OperationId
		if strings.HasPrefix(id, "listen") || strings.HasSuffix(id, "Job") || strings.HasSuffix(id, "Listener") {
			return resources
		}

		// The Swagger parameters definition for the operation
		params := op.Parameters

		requiredProperties := []string{}
		allProperties := []string{}
		var resourceName string

		for _, param := range params {
			paramName := param.Name
			paramIn := param.In

			// We only process the data in the body of the request
			if paramIn == "body" {

				resourceName = paramName

				bodySchemaProxy := param.Schema
				if bodySchemaProxy != nil {
					bodySchema := bodySchemaProxy.Schema()

					requiredProperties = append(requiredProperties, bodySchema.Required...)

					bodyProperties := bodySchema.Properties
					if bodyProperties != nil {
						for key, bodyPropProxy := range bodyProperties.FromOldest() {
							bodyPropSchema := bodyPropProxy.Schema()
							if bodyPropSchema != nil {
								allProperties = append(allProperties, key)
							}
						}
					}

				}
			}
		}

		// // Add to required some fields in DOME and ISBE
		// // All objects except "category", "individual" and "organization", require the "relatedParty"
		// exceptions := []string{"category", "individual", "organization"}
		// if action == "CREATE" {
		// 	if !slices.Contains(exceptions, resourceName) {
		// 		// Add to required if it is not yet
		// 		if !slices.Contains(requiredProperties, "relatedParty") {
		// 			requiredProperties = append(requiredProperties, "relatedParty")
		// 		}
		// 	}
		// }

		thisOperation := &types.Action{
			Resource: resourceName,
			Action:   action,
			Required: requiredProperties,
			Fields:   allProperties,
		}

		if resources[resourceName] == nil {
			resources[resourceName] = &types.Resource{
				BasePath: basePath,
				Actions:  make(map[string]*types.Action),
				Public:   IsPublicResource(resourceName),
			}
		}

		resources[resourceName].Actions[action] = thisOperation

	}
	return resources

}

func printStatistics(resources types.Resources) {

	fmt.Println("Resources with VERSION")
	for resourceName, resource := range resources {
		for actionName, action := range resource.Actions {
			if actionName == "UPDATE" {
				continue
			}

			var publicStr string
			if resource.Public {
				publicStr = "Public YES"
			} else {
				publicStr = "Public NO "
			}

			hasVersion := slices.Contains(action.Fields, "version")
			var versionStr string
			if hasVersion {
				versionStr = "Version YES"
			} else {
				versionStr = "Version NO "
			}

			hasLifecycleStatus := slices.Contains(action.Fields, "lifecycleStatus")
			var lifecycleStatusStr string
			if hasLifecycleStatus {
				lifecycleStatusStr = "LifecycleStatus YES"
			} else {
				lifecycleStatusStr = "LifecycleStatus NO "
			}

			hasLastUpdate := slices.Contains(action.Fields, "lastUpdate")
			var lastUpdateStr string
			if hasLastUpdate {
				lastUpdateStr = "LastUpdate YES"
			} else {
				lastUpdateStr = "LastUpdate NO "
			}

			fmt.Printf("Resource: %s (%s)\n", resourceName, resource.BasePath)
			fmt.Printf("    %s %s %s %s\n", publicStr, versionStr, lifecycleStatusStr, lastUpdateStr)
		}
	}

}

func generateDefinitions(resources types.Resources) {

	b, err := yaml.Marshal(resources)
	if err != nil {
		panic(err)
	}

	if updateFile {

		err = os.WriteFile("./tmf_operations.yaml", b, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println(string(b))
	}
}

// Public TMF resources are accessible by all users, even unauthenticated ones
func IsPublicResource(resourceName string) bool {
	resourceName = strings.ToLower(resourceName)
	_, ok := PublicResources[resourceName]
	return ok
}

// The public objects are the following:
var PublicResources = map[string]bool{
	// TMF620 Product Catalog Management
	"category":             true,
	"catalog":              true,
	"productoffering":      true,
	"productofferingprice": true,
	"productspecification": true,

	// TMF633 Service Catalog Management
	"servicecatalog":       true,
	"servicecategory":      true,
	"servicecandidate":     true,
	"servicespecification": true,

	// TMF634 Resource Catalog Management
	"resourcecatalog":       true,
	"resourcecategory":      true,
	"resourcecandidate":     true,
	"resourcespecification": true,

	// Organization from TMF632 Party Management. But Individual is private.
	"organization": true,

	// TMF669 Party Role Management
	// "partyrole": true,
}
