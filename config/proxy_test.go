package config

import (
	"fmt"
	"testing"
)

func TestUpstreamURL(t *testing.T) {
	tests := []struct {
		name         string
		resourceName string
		wantURL      string
		wantErr      bool
	}{
		{
			name:         "catalog resource",
			resourceName: "catalog",
			wantURL:      "http://tm-forum-api-product-catalog:8080/tmf-api/productCatalogManagement/v4",
			wantErr:      false,
		},
		{
			name:         "inventory resource",
			resourceName: "inventory",
			wantURL:      "http://tm-forum-api-product-inventory:8080/tmf-api/productInventory/v4",
			wantErr:      false,
		},
		{
			name:         "ordering resource",
			resourceName: "ordering",
			wantURL:      "http://tm-forum-api-product-ordering-management:8080/tmf-api/productOrderingManagement/v4",
			wantErr:      false,
		},
		{
			name:         "billing resource",
			resourceName: "billing",
			wantURL:      "http://tm-forum-api-account:8080/tmf-api/accountManagement/v4",
			wantErr:      false,
		},
		{
			name:         "usage resource",
			resourceName: "usage",
			wantURL:      "http://tm-forum-api-usage-management:8080/tmf-api/usageManagement/v4",
			wantErr:      false,
		},
		{
			name:         "party resource",
			resourceName: "party",
			wantURL:      "http://tm-forum-api-party-catalog:8080/tmf-api/party/v4",
			wantErr:      false,
		},
		{
			name:         "customer resource",
			resourceName: "customer",
			wantURL:      "http://tm-forum-api-customer-management:8080/tmf-api/customerManagement/v4",
			wantErr:      false,
		},
		{
			name:         "resources resource",
			resourceName: "resources",
			wantURL:      "http://tm-forum-api-resource-catalog:8080/tmf-api/resourceCatalog/v4",
			wantErr:      false,
		},
		{
			name:         "services resource",
			resourceName: "services",
			wantURL:      "http://tm-forum-api-service-catalog:8080/tmf-api/serviceCatalogManagement/v4",
			wantErr:      false,
		},
		{
			name:         "resourceInventory resource",
			resourceName: "resourceInventory",
			wantURL:      "http://tm-forum-api-resource-inventory:8080/tmf-api/resourceInventoryManagement/v4",
			wantErr:      false,
		},
		{
			name:         "serviceInventory resource",
			resourceName: "serviceInventory",
			wantURL:      "http://tm-forum-api-service-inventory:8080/tmf-api/serviceInventory/v4",
			wantErr:      false,
		},
		{
			name:         "unknown resource returns error",
			resourceName: "nonExistentResource",
			wantURL:      "",
			wantErr:      true,
		},
		{
			name:         "empty resource name returns error",
			resourceName: "",
			wantURL:      "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, err := InternalUpstreamURL(tt.resourceName)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpstreamURL(%q) error = %v, wantErr %v", tt.resourceName, err, tt.wantErr)
				return
			}

			if gotURL != tt.wantURL {
				t.Errorf("UpstreamURL(%q) = %q, want %q", tt.resourceName, gotURL, tt.wantURL)
			}
		})
	}
}

func TestUpstreamURL_URLFormat(t *testing.T) {
	url, err := InternalUpstreamURL("catalog")
	if err != nil {
		t.Fatalf("UpstreamURL(\"catalog\") unexpected error: %v", err)
	}

	// Verify the URL follows the expected http://host:port/path pattern
	want := fmt.Sprintf("http://%s:%d%s", "tm-forum-api-product-catalog", 8080, "/tmf-api/productCatalogManagement/v4")
	if url != want {
		t.Errorf("URL format mismatch: got %q, want %q", url, want)
	}
}
