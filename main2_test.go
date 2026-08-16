// Copyright 2025 Jesus Ruiz
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/hesusruiz/tmforum/config"
	"github.com/hesusruiz/tmforum/tmfserver/repository"
	"github.com/hesusruiz/tmforum/types"
)

const (
	serverURL = "https://tmf.mycredential.eu/tmf-api/productCatalogManagement/v4"
	// serverURL = "https://tmf.evidenceledger.eu/tmf-api/productCatalogManagement/v4"
)

var (
	apiToken = `eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL2NhdGFsb2cuaXNiZW9uYm9hcmQuY29tIiwiZXhwIjoxNzYxMjI5NDE3LCJpYXQiOjE3NjEyMjU4MTcsImlzcyI6Imh0dHBzOi8vY2VydGF1dGguZXZpZGVuY2VsZWRnZXIuZXUiLCJqdGkiOiJUQ1RJVktTQjZQQTNTRlJYNVlCVUZPUkdZNyIsIm5vbmNlIjoiZDA4Mjg1N2MwNDIxYmViZTBkYTU5YTA5ZmYwYWVhM2NlOW1SczZVN1AiLCJzY29wZSI6Im9wZW5pZCBlaWRhcyIsInN1YiI6Imh0dHBzOi8vY2F0YWxvZy5pc2Jlb25ib2FyZC5jb20iLCJ2YyI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvbnMvY3JlZGVudGlhbHMvIiwiaHR0cHM6Ly9jcmVkZW50aWFscy5ldWRpc3RhY2suZXUvLndlbGwta25vd24vY3JlZGVudGlhbHMvbGVhcl9jcmVkZW50aWFsX2VtcGxveWVlL3czYy92MyJdLCJjcmVkZW50aWFsU3RhdHVzIjp7ImlkIjoiaHR0cHM6Ly9pc3N1ZXIuZG9tZS1tYXJrZXRwbGFjZS1zYngub3JnL2JhY2tvZmZpY2UvdjEvY3JlZGVudGlhbHMvc3RhdHVzLzEjU1lDOTA4UklRUXFlVVhSMTluaDNFUSIsInN0YXR1c0xpc3RDcmVkZW50aWFsIjoiaHR0cHM6Ly9pc3N1ZXIuZG9tZS1tYXJrZXRwbGFjZS1zYngub3JnL2JhY2tvZmZpY2UvdjEvY3JlZGVudGlhbHMvc3RhdHVzLzEiLCJzdGF0dXNMaXN0SW5kZXgiOiJTWUM5MDhSSVFRcWVVWFIxOW5oM0VRIiwic3RhdHVzUHVycG9zZSI6InJldm9jYXRpb24iLCJ0eXBlIjoiUGxhaW5MaXN0RW50aXR5In0sImNyZWRlbnRpYWxTdWJqZWN0Ijp7Im1hbmRhdGUiOnsibWFuZGF0ZWUiOnsiZW1haWwiOiJhbGJhLmxvcGV6QGluMi5lcyIsImVtcGxveWVlSWQiOiIxMjM0NTY3OEEiLCJmaXJzdE5hbWUiOiJKb2huIiwiaWQiOiIxMjM0NTY3OEEiLCJsYXN0TmFtZSI6IkRvZSJ9LCJtYW5kYXRvciI6eyJjb21tb25OYW1lIjoiSm9obiBEb2UiLCJjb3VudHJ5IjoiRVMiLCJlbWFpbCI6ImFsYmEubG9wZXpAaW4yLmVzIiwiaWQiOiJkaWQ6ZWxzaTpWQVRFUy0xMTExMTExMUsiLCJvcmdhbml6YXRpb24iOiJJU0JFIEZvdW5kYXRpb24iLCJvcmdhbml6YXRpb25JZGVudGlmaWVyIjoiVkFURVMtMTExMTExMTFLIiwic2VyaWFsTnVtYmVyIjoiMTIzNDU2NzhBIn0sInBvd2VyIjpbeyJhY3Rpb24iOlsiRXhlY3V0ZSJdLCJkb21haW4iOiJET01FIiwiZnVuY3Rpb24iOiJPbmJvYXJkaW5nIiwidHlwZSI6ImRvbWFpbiJ9LHsiYWN0aW9uIjpbIkNyZWF0ZSIsIlVwZGF0ZSIsIkRlbGV0ZSJdLCJkb21haW4iOiJET01FIiwiZnVuY3Rpb24iOiJQcm9kdWN0T2ZmZXJpbmciLCJ0eXBlIjoiZG9tYWluIn1dfX0sImRlc2NyaXB0aW9uIjoiVmVyaWZpYWJsZSBDcmVkZW50aWFsIGZvciBlbXBsb3llZXMgb2YgYW4gb3JnYW5pemF0aW9uIiwiaWQiOiJ1cm46dXVpZENXWFhTQzZDS1M3RElQV1hVNkM3NkFFSFpJIiwiaXNzdWVyIjp7ImNvbW1vbk5hbWUiOiJDZXJ0QXV0aCBJZGVudGl0eSBQcm92aWRlciBmb3IgSVNCRSIsImNvdW50cnkiOiJFUyIsImlkIjoiZGlkOmVsc2k6VkFURVMtQjYwNjQ1OTAwIiwib3JnYW5pemF0aW9uIjoiSU4yIiwib3JnYW5pemF0aW9uSWRlbnRpZmllciI6IlZBVEVTLUI2MDY0NTkwMCIsInNlcmlhbE51bWJlciI6IkI0NzQ0NzU2MCJ9LCJ0eXBlIjpbIkxFQVJDcmVkZW50aWFsRW1wbG95ZWUiLCJWZXJpZmlhYmxlQ3JlZGVudGlhbCJdLCJ2YWxpZEZyb20iOiIyMDI1LTEwLTIzVDEzOjA2OjIwWiIsInZhbGlkVW50aWwiOiIyMDI2LTEwLTIzVDEzOjA2OjIwWiJ9fQ.LSIItTupThpBDG08pWbAdIk0_qaG-U6w7SpbOPyjOUn-0wybXPwl-dyv8uRiEnbkLb0Mwgch-zkQGav3ImF4UA`
)

// ProductSpecification represents the structure for a product specification.
// This is a simplified version based on the TMF620 swagger file.
type ProductSpecification struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Href            string `json:"href,omitempty"`
	Brand           string `json:"brand,omitempty"`
	Description     string `json:"description,omitempty"`
	LastUpdate      string `json:"lastUpdate,omitempty"`
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	Version         string `json:"version,omitempty"`
}

var LocalConfig = &config.Config{
	Environment:  config.LOCAL,
	ProxyEnabled: false,

	// The operator is Altia
	ServerOperatorOrganizationIdentifier: "VATES-A15456585",
	ServerOperatorDid:                    "did:elsi:VATES-A15456585",
	ServerOperatorName:                   "ALTIA CONSULTORES SA",
	ServerOperatorCountry:                "ES",

	AdminToken: "eyJhdWQiOiJodHRwczovL2NhdGFsb2cuaX",

	LEARPower: types.OnePower{
		Type:     "domain",
		Domain:   "DOME",
		Function: "Onboarding",
		Action:   []string{"execute"},
	},
	ProductCreatePower: types.OnePower{
		Type:     "domain",
		Domain:   "DOME",
		Function: "ProductOffering",
		Action:   []string{"Create"},
	},
	ProductUpdatePower: types.OnePower{
		Type:     "domain",
		Domain:   "DOME",
		Function: "ProductOffering",
		Action:   []string{"Update"},
	},
	ProductDeletePower: types.OnePower{
		Type:     "domain",
		Domain:   "DOME",
		Function: "ProductOffering",
		Action:   []string{"Delete"},
	},

	PolicyFileName:  "auth_policies.star",
	RemoteTMFServer: "https://tmf.dome-marketplace-sbx.org",
	VerifierServer:  "https://verifier.dome-marketplace-sbx.org",
	Dbname:          ":memory:",
	ClonePeriod:     config.DefaultClonePeriod,
	RestartHour:     -1,
	Features: config.Features{
		OfferingLaunchOnlyByAdmin: false,
		GenerateIDOnCreate:        true,
		AllowIDInBody:             false,
		VerifyJWTSignature:        false,
	},
}

// Create and start the web server only once even if the tests can run in parallel
func init() {

	// Set the nocolor option for logs
	os.Setenv("ISBETMF_LOGS_NOCOLOR", "true")

	// // Generate a default configuration suitable for the environment
	// // The approach is that instead of many configurable parameters, we have a set of profiles, with "hardcoded"
	// // parameters for each environment, but that can be easity extended for other purposes.
	// configuration, err := config.LoadConfig("local", true)
	// if err != nil {
	// 	panic(err)
	// }

	go runNormalProcess(LocalConfig)

	time.Sleep(1 * time.Second)

}

// doHTTPRequestStd is a helper function to make HTTP requests using standard net/http package.
func doHTTPRequestStd(t *testing.T, client *http.Client, method, url string, token string, payload any) (*http.Response, []byte) {
	t.Helper()
	var bodyReader io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("failed to marshal payload: %v", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("HTTP request %s %s failed: %v", method, url, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	return resp, respBody
}

func TestNoAuthorization_StdHTTP(t *testing.T) {
	client := &http.Client{}

	// 1. Create (POST) without authentication
	ps := repository.TMFObjectMap{
		"name":            "My Test Product Specification",
		"brand":           "TestBrand",
		"description":     "A detailed description of my test product specification.",
		"lifecycleStatus": "Active",
	}

	resp, _ := doHTTPRequestStd(t, client, http.MethodPost, serverURL+"/productSpecification", "", ps)
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status %d %s, got %d %s",
			http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized),
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}
}

func TestProductSpecificationHappy_StdHTTP(t *testing.T) {
	client := &http.Client{}

	// 1. Create (POST)
	ps := repository.TMFObjectMap{
		"name":            "My Test Product Specification",
		"brand":           "TestBrand",
		"description":     "A detailed description of my test product specification.",
		"lifecycleStatus": "Active",
	}

	resp, body := doHTTPRequestStd(t, client, http.MethodPost, serverURL+"/productSpecification", apiToken, ps)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d: %s", http.StatusCreated, resp.StatusCode, string(body))
	}

	var createdSpecObj map[string]any
	if err := json.Unmarshal(body, &createdSpecObj); err != nil {
		t.Fatalf("failed to unmarshal JSON response: %v", err)
	}

	id, _ := createdSpecObj["id"].(string)
	if id == "" {
		t.Errorf("expected non-empty id")
	}

	name, _ := createdSpecObj["name"].(string)
	if name != ps.Name() {
		t.Errorf("expected name %q, got %q", ps.Name(), name)
	}

	brand, _ := createdSpecObj["brand"].(string)
	if brand != ps.GetStringField("brand") {
		t.Errorf("expected brand %q, got %q", ps.GetStringField("brand"), brand)
	}

	createdSpecID := id

	// 2. Get (GET)
	resp, body = doHTTPRequestStd(t, client, http.MethodGet, serverURL+"/productSpecification/"+createdSpecID, apiToken, nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, resp.StatusCode, string(body))
	}

	var getSpecObj map[string]any
	if err := json.Unmarshal(body, &getSpecObj); err != nil {
		t.Fatalf("failed to unmarshal JSON response: %v", err)
	}
	if getSpecID, _ := getSpecObj["id"].(string); getSpecID != createdSpecID {
		t.Errorf("expected id %q, got %q", createdSpecID, getSpecID)
	}

	// 3. Update (PATCH)
	updatePayload := map[string]any{"description": "An updated description."}
	resp, body = doHTTPRequestStd(t, client, http.MethodPatch, serverURL+"/productSpecification/"+createdSpecID, apiToken, updatePayload)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, resp.StatusCode, string(body))
	}

	var patchSpecObj map[string]any
	if err := json.Unmarshal(body, &patchSpecObj); err != nil {
		t.Fatalf("failed to unmarshal JSON response: %v", err)
	}
	if desc, _ := patchSpecObj["description"].(string); desc != "An updated description." {
		t.Errorf("expected description %q, got %q", "An updated description.", desc)
	}

	// 4. Delete (DELETE)
	resp, body = doHTTPRequestStd(t, client, http.MethodDelete, serverURL+"/productSpecification/"+createdSpecID, apiToken, nil)
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected status %d, got %d: %s", http.StatusNoContent, resp.StatusCode, string(body))
	}

	// 5. List (GET)
	// Create two new product specifications to test the list functionality
	spec1 := ProductSpecification{
		Name:  "List Spec 1",
		Brand: "TestBrand",
	}
	spec2 := ProductSpecification{
		Name:  "List Spec 2",
		Brand: "TestBrand",
	}

	resp, body = doHTTPRequestStd(t, client, http.MethodPost, serverURL+"/productSpecification", apiToken, spec1)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d: %s", http.StatusCreated, resp.StatusCode, string(body))
	}
	var spec1Obj map[string]any
	_ = json.Unmarshal(body, &spec1Obj)
	createdSpec1ID, _ := spec1Obj["id"].(string)

	resp, body = doHTTPRequestStd(t, client, http.MethodPost, serverURL+"/productSpecification", apiToken, spec2)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d: %s", http.StatusCreated, resp.StatusCode, string(body))
	}
	var spec2Obj map[string]any
	_ = json.Unmarshal(body, &spec2Obj)
	createdSpec2ID, _ := spec2Obj["id"].(string)

	// Retrieve list and assert it contains at least two items
	resp, body = doHTTPRequestStd(t, client, http.MethodGet, serverURL+"/productSpecification", apiToken, nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, resp.StatusCode, string(body))
	}

	var list []any
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatalf("failed to unmarshal JSON list response: %v", err)
	}

	if len(list) < 2 {
		t.Errorf("expected at least 2 items, got %d", len(list))
	}

	// Cleanup
	resp, body = doHTTPRequestStd(t, client, http.MethodDelete, serverURL+"/productSpecification/"+createdSpec1ID, apiToken, nil)
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected status %d, got %d: %s", http.StatusNoContent, resp.StatusCode, string(body))
	}
	resp, body = doHTTPRequestStd(t, client, http.MethodDelete, serverURL+"/productSpecification/"+createdSpec2ID, apiToken, nil)
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected status %d, got %d: %s", http.StatusNoContent, resp.StatusCode, string(body))
	}
}

func TestInvalidSeller_StdHTTP(t *testing.T) {
	client := &http.Client{}

	// 1. Create (POST)
	ps := repository.TMFObjectMap{
		"name":            "My Test Product Specification",
		"brand":           "TestBrand",
		"description":     "A detailed description of my test product specification.",
		"lifecycleStatus": "Active",
	}

	ps.SetSellerInfo("pepe", "juan", "v4")

	resp, _ := doHTTPRequestStd(t, client, http.MethodPost, serverURL+"/productSpecification", apiToken, ps)
	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("expected status %d, got %d", http.StatusForbidden, resp.StatusCode)
	}
}

func TestCategoryCreateUpdateName_StdHTTP(t *testing.T) {
	client := &http.Client{}

	// 1. Create (POST)
	ps := repository.TMFObjectMap{
		"@type":           "category",
		"name":            "Identidad Digital",
		"brand":           "TestBrand",
		"description":     "A detailed description of my test product specification.",
		"lifecycleStatus": "Active",
	}

	resp, body := doHTTPRequestStd(t, client, http.MethodPost, serverURL+"/category", LocalConfig.AdminToken, ps)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status %d, got %d: %s", http.StatusCreated, resp.StatusCode, string(body))
	}

	var createdObj map[string]any
	if err := json.Unmarshal(body, &createdObj); err != nil {
		t.Fatalf("failed to unmarshal JSON response: %v", err)
	}

	id, _ := createdObj["id"].(string)
	if id == "" {
		t.Errorf("expected non-empty id")
	}

	name, _ := createdObj["name"].(string)
	if name != ps.Name() {
		t.Errorf("expected name %q, got %q", ps.Name(), name)
	}

	brand, _ := createdObj["brand"].(string)
	if brand != ps.GetStringField("brand") {
		t.Errorf("expected brand %q, got %q", ps.GetStringField("brand"), brand)
	}

	createdSpecID := id

	// 2. Get (GET)
	resp, body = doHTTPRequestStd(t, client, http.MethodGet, serverURL+"/category/"+createdSpecID, apiToken, nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, resp.StatusCode, string(body))
	}

	var getObj map[string]any
	if err := json.Unmarshal(body, &getObj); err != nil {
		t.Fatalf("failed to unmarshal JSON response: %v", err)
	}
	if getID, _ := getObj["id"].(string); getID != createdSpecID {
		t.Errorf("expected id %q, got %q", createdSpecID, getID)
	}

	// 3. Update (PATCH)
	updatePayload := map[string]any{
		"description": "An updated description.",
		"name":        "Gestión de Identidad Digital y Confianza",
	}

	resp, body = doHTTPRequestStd(t, client, http.MethodPatch, serverURL+"/category/"+createdSpecID, LocalConfig.AdminToken, updatePayload)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, resp.StatusCode, string(body))
	}

	// 4. Get again the updated object (GET)
	resp, body = doHTTPRequestStd(t, client, http.MethodGet, serverURL+"/category/"+createdSpecID, apiToken, nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, resp.StatusCode, string(body))
	}

	var updatedObject map[string]any
	if err := json.Unmarshal(body, &updatedObject); err != nil {
		t.Fatalf("failed to unmarshal JSON response: %v", err)
	}

	if updatedName, _ := updatedObject["name"].(string); updatedName != updatePayload["name"].(string) {
		t.Errorf("expected name %q, got %q", updatePayload["name"].(string), updatedName)
	}
	if updatedDesc, _ := updatedObject["description"].(string); updatedDesc != "An updated description." {
		t.Errorf("expected description %q, got %q", "An updated description.", updatedDesc)
	}
}
