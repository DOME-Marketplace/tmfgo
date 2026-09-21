package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	repository "github.com/DOME-Marketplace/tmfgo/tmfserver/repository"
	"github.com/gavv/httpexpect/v2"
)

func TestPerformance(t *testing.T) {
	baseURL := os.Getenv("TMF_BASE_URL")
	if baseURL == "" {
		baseURL = serverURL
	}

	// Create a new httpexpect instance.
	// We reuse the serverURL and apiToken from main_test.go since we are in the same package.
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  baseURL,
		Reporter: httpexpect.NewAssertReporter(t),
		// Disable printers for performance test to avoid noise and slowdown
		Printers: []httpexpect.Printer{},
	})

	iterations := 50 // Start with a smaller number for the first run
	start := time.Now()

	for i := range iterations {
		// 1. Create (POST)
		ps := repository.TMFObjectMap{
			"name":            fmt.Sprintf("Perf Test Spec %d", i),
			"brand":           "TestBrand",
			"description":     "Performance test description.",
			"lifecycleStatus": "Active",
		}

		resp := e.POST("/productSpecification").
			WithHeader("Authorization", "Bearer "+apiToken).
			WithJSON(ps).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		id := resp.Value("id").String().Raw()

		// 2. Get (GET)
		e.GET("/productSpecification/{id}", id).
			WithHeader("Authorization", "Bearer "+apiToken).
			Expect().
			Status(http.StatusOK)

		// 3. Update (PATCH)
		updatePayload := map[string]any{"description": "Updated description."}
		e.PATCH("/productSpecification/{id}", id).
			WithHeader("Authorization", "Bearer "+apiToken).
			WithJSON(updatePayload).
			Expect().
			Status(http.StatusOK)

		// 4. Delete (DELETE)
		e.DELETE("/productSpecification/{id}", id).
			WithHeader("Authorization", "Bearer "+apiToken).
			Expect().
			Status(http.StatusNoContent)
	}

	elapsed := time.Since(start)
	totalOps := iterations * 4
	opsPerSec := float64(totalOps) / elapsed.Seconds()

	t.Logf("Performance test finished. %d iterations took %s", iterations, elapsed)
	t.Logf("Total HTTP operations: %d", totalOps)
	t.Logf("Operations per second: %.2f ops/s", opsPerSec)
}

func TestPerformanceISBEDev(t *testing.T) {
	baseURL := os.Getenv("TMF_BASE_URL")
	if baseURL == "" {
		baseURL = serverURL
	}

	// Create a new httpexpect instance with the specific server URL
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  baseURL,
		Reporter: httpexpect.NewAssertReporter(t),
		// Disable printers for performance test to avoid noise and slowdown
		Printers: []httpexpect.Printer{},
	})

	iterations := 50 // Start with a smaller number for the first run
	start := time.Now()

	for i := 0; i < iterations; i++ {
		// 1. Create (POST)
		ps := repository.TMFObjectMap{
			"name":            fmt.Sprintf("Perf Test Spec %d", i),
			"brand":           "TestBrand",
			"description":     "Performance test description.",
			"lifecycleStatus": "Active",
		}

		resp := e.POST("/productSpecification").
			WithHeader("Authorization", "Bearer "+apiToken).
			WithJSON(ps).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		id := resp.Value("id").String().Raw()

		// 2. Get (GET)
		e.GET("/productSpecification/{id}", id).
			WithHeader("Authorization", "Bearer "+apiToken).
			Expect().
			Status(http.StatusOK)

		// 3. Update (PATCH)
		updatePayload := map[string]any{"description": "Updated description."}
		e.PATCH("/productSpecification/{id}", id).
			WithHeader("Authorization", "Bearer "+apiToken).
			WithJSON(updatePayload).
			Expect().
			Status(http.StatusOK)

		// 4. Delete (DELETE)
		e.DELETE("/productSpecification/{id}", id).
			WithHeader("Authorization", "Bearer "+apiToken).
			Expect().
			Status(http.StatusNoContent)
	}

	elapsed := time.Since(start)
	totalOps := iterations * 4
	opsPerSec := float64(totalOps) / elapsed.Seconds()

	t.Logf("Performance test finished. %d iterations took %s", iterations, elapsed)
	t.Logf("Total HTTP operations: %d", totalOps)
	t.Logf("Operations per second: %.2f ops/s", opsPerSec)
}

func TestPerformanceDetailed(t *testing.T) {
	baseURL := os.Getenv("TMF_BASE_URL")
	if baseURL == "" {
		baseURL = serverURL
	}

	// Create a new httpexpect instance.
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  baseURL,
		Reporter: httpexpect.NewAssertReporter(t),
		// Disable printers for performance test to avoid noise and slowdown
		Printers: []httpexpect.Printer{},
	})

	iterations := 50
	ids := make([]string, 0, iterations)

	// 1. Creation Loop
	startCreate := time.Now()
	for i := 0; i < iterations; i++ {
		ps := repository.TMFObjectMap{
			"name":            fmt.Sprintf("Perf Test Spec %d", i),
			"brand":           "TestBrand",
			"description":     "Performance test description.",
			"lifecycleStatus": "Active",
		}

		resp := e.POST("/productSpecification").
			WithHeader("Authorization", "Bearer "+apiToken).
			WithJSON(ps).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		id := resp.Value("id").String().Raw()
		ids = append(ids, id)
	}
	elapsedCreate := time.Since(startCreate)
	createOpsPerSec := float64(iterations) / elapsedCreate.Seconds()

	// 2. Retrieval Loop
	startGet := time.Now()
	for _, id := range ids {
		e.GET("/productSpecification/{id}", id).
			WithHeader("Authorization", "Bearer "+apiToken).
			Expect().
			Status(http.StatusOK)
	}
	elapsedGet := time.Since(startGet)
	getOpsPerSec := float64(iterations) / elapsedGet.Seconds()

	// 3. Update Loop
	startUpdate := time.Now()
	updatePayload := map[string]any{"description": "Updated description."}
	for _, id := range ids {
		e.PATCH("/productSpecification/{id}", id).
			WithHeader("Authorization", "Bearer "+apiToken).
			WithJSON(updatePayload).
			Expect().
			Status(http.StatusOK)
	}
	elapsedUpdate := time.Since(startUpdate)
	updateOpsPerSec := float64(iterations) / elapsedUpdate.Seconds()

	// 4. Deletion Loop
	startDelete := time.Now()
	for _, id := range ids {
		e.DELETE("/productSpecification/{id}", id).
			WithHeader("Authorization", "Bearer "+apiToken).
			Expect().
			Status(http.StatusNoContent)
	}
	elapsedDelete := time.Since(startDelete)
	deleteOpsPerSec := float64(iterations) / elapsedDelete.Seconds()

	t.Logf("Performance Test Detailed Results (%d iterations):", iterations)
	t.Logf("Creation:  %s total, %.2f ops/s", elapsedCreate, createOpsPerSec)
	t.Logf("Retrieval: %s total, %.2f ops/s", elapsedGet, getOpsPerSec)
	t.Logf("Update:    %s total, %.2f ops/s", elapsedUpdate, updateOpsPerSec)
	t.Logf("Deletion:  %s total, %.2f ops/s", elapsedDelete, deleteOpsPerSec)
}
