package types

import (
	_ "embed"
	"slices"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/hesusruiz/tmforum/internal/errl"
)

type Action struct {
	Resource string   // The resource name, e.g., "ProductOffering", "ProductSpecification"
	Action   string   // The action (synonim of the HTTP verb), e.g., "CREATE", "UPDATE"
	Required []string // A list of the required fields in the body of the request
	Fields   []string // A list of all the fields in the body of the request
}

func (a *Action) HasField(field string) bool {
	return slices.Contains(a.Fields, field)
}

type Resource struct {
	BasePath string
	Public   bool
	Actions  map[string]*Action
}

type Resources map[string]*Resource

var tmf_resource_requirements Resources

//go:embed tmf_operations.yaml
var tmfOperationsYAML []byte

// ParseActionDefinitions parses the YAML definition of the TMF resources and actions.
func ParseActionDefinitions() {
	// Parse the YAML definition
	if err := yaml.Unmarshal(tmfOperationsYAML, &tmf_resource_requirements); err != nil {
		panic(errl.Errorf("failed to unmarshal tmf_operations.yaml: %w", err))
	}
}

func GetResourceDefinition(resource string) *Resource {
	res, ok := tmf_resource_requirements[resource]
	if !ok {
		return nil
	}
	return res
}

func GetActionDefinition(resource string, action string) *Action {
	res := GetResourceDefinition(resource)
	if res == nil {
		return nil
	}

	act, ok := res.Actions[action]
	if !ok {
		return nil
	}

	return act

}

// The names of some special objects in the DOME ecosystem
const ProductOffering = "productOffering"
const ProductSpecification = "productSpecification"
const ProductOfferingPrice = "productOfferingPrice"
const ServiceSpecification = "serviceSpecification"
const ResourceSpecification = "resourceSpecification"
const Category = "category"
const Catalog = "catalog"
const Organization = "organization"
const Individual = "individual"

var withoutRelyingParty = []string{Category, Organization, Individual}

func IsWithoutRelyingParty(resourceName string) bool {
	return slices.Contains(withoutRelyingParty, strings.ToLower(resourceName))
}

var RequiredFieldsForAllObjects = []string{
	"id", "href",
}

var RecommendedFieldsForAllObjects = []string{
	"name", "version", "lastUpdate",
}

var DoNotRequireRelatedParties = []string{
	"category",
	"individual",
	"organization",
}

var DoNotRequireBuyerInfo = []string{
	"category",
	"individual",
	"organization",
	"catalog",
	"productoffering",
	"productspecification",
	"productofferingprice",
	"resourcespecification",
	"servicespecification",
}
