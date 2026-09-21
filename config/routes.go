// Copyright 2023 Jesus Ruiz. All rights reserved.
// Use of this source code is governed by an Apache 2.0
// license that can be found in the LICENSE file.

package config

import "github.com/DOME-Marketplace/tmfgo/internal/errl"

var GeneratedDefaultResourceToPathPrefix = map[string]string{
	"agreement":                  "/tmf-api/agreementManagement/v4/agreement",
	"agreementSpecification":     "/tmf-api/agreementManagement/v4/agreementSpecification",
	"appliedCustomerBillingRate": "/tmf-api/customerBillManagement/v4/appliedCustomerBillingRate",
	"billFormat":                 "/tmf-api/accountManagement/v4/billFormat",
	"billPresentationMedia":      "/tmf-api/accountManagement/v4/billPresentationMedia",
	"billingAccount":             "/tmf-api/accountManagement/v4/billingAccount",
	"billingCycleSpecification":  "/tmf-api/accountManagement/v4/billingCycleSpecification",
	"cancelProductOrder":         "/tmf-api/productOrderingManagement/v4/cancelProductOrder",
	"cancelResourceOrder":        "/tmf-api/resourceOrderingManagement/v4/cancelResourceOrder",
	"cancelServiceOrder":         "/tmf-api/serviceOrdering/v4/cancelServiceOrder",
	"catalog":                    "/tmf-api/productCatalogManagement/v4/catalog",
	"category":                   "/tmf-api/productCatalogManagement/v4/category",
	"customer":                   "/tmf-api/customerManagement/v4/customer",
	"customerBill":               "/tmf-api/customerBillManagement/v4/customerBill",
	"customerBillOnDemand":       "/tmf-api/customerBillManagement/v4/customerBillOnDemand",
	"financialAccount":           "/tmf-api/accountManagement/v4/financialAccount",
	"heal":                       "/tmf-api/resourceFunctionActivation/v4/heal",
	"individual":                 "/tmf-api/party/v4/individual",
	"migrate":                    "/tmf-api/resourceFunctionActivation/v4/migrate",
	"monitor":                    "/tmf-api/resourceFunctionActivation/v4/monitor",
	"organization":               "/tmf-api/party/v4/organization",
	"partyAccount":               "/tmf-api/accountManagement/v4/partyAccount",
	"partyRole":                  "/tmf-api/partyRoleManagement/v4/partyRole",
	"product":                    "/tmf-api/productInventory/v4/product",
	"productOffering":            "/tmf-api/productCatalogManagement/v4/productOffering",
	"productOfferingPrice":       "/tmf-api/productCatalogManagement/v4/productOfferingPrice",
	"productOrder":               "/tmf-api/productOrderingManagement/v4/productOrder",
	"productSpecification":       "/tmf-api/productCatalogManagement/v4/productSpecification",
	"quote":                      "/tmf-api/quoteManagement/v4/quote",
	"resource":                   "/tmf-api/resourceInventoryManagement/v4/resource",
	"resourceCandidate":          "/tmf-api/resourceCatalog/v4/resourceCandidate",
	"resourceCatalog":            "/tmf-api/resourceCatalog/v4/resourceCatalog",
	"resourceCategory":           "/tmf-api/resourceCatalog/v4/resourceCategory",
	"resourceFunction":           "/tmf-api/resourceFunctionActivation/v4/resourceFunction",
	"resourceOrder":              "/tmf-api/resourceOrderingManagement/v4/resourceOrder",
	"resourceSpecification":      "/tmf-api/resourceCatalog/v4/resourceSpecification",
	"scale":                      "/tmf-api/resourceFunctionActivation/v4/scale",
	"service":                    "/tmf-api/serviceInventory/v4/service",
	"serviceCandidate":           "/tmf-api/serviceCatalogManagement/v4/serviceCandidate",
	"serviceCatalog":             "/tmf-api/serviceCatalogManagement/v4/serviceCatalog",
	"serviceCategory":            "/tmf-api/serviceCatalogManagement/v4/serviceCategory",
	"serviceOrder":               "/tmf-api/serviceOrdering/v4/serviceOrder",
	"serviceSpecification":       "/tmf-api/serviceCatalogManagement/v4/serviceSpecification",
	"settlementAccount":          "/tmf-api/accountManagement/v4/settlementAccount",
	"usage":                      "/tmf-api/usageManagement/v4/usage",
	"usageSpecification":         "/tmf-api/usageManagement/v4/usageSpecification",
}

func ExternalUpstreamTMFPath(resourceName string) (string, error) {
	if pathPrefix, exists := GeneratedDefaultResourceToPathPrefix[resourceName]; exists {
		return pathPrefix, nil
	}
	return "", errl.Errorf("unknown object type: %s", resourceName)
}
