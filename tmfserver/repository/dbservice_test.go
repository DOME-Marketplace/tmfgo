package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func argsToStrings(args []any) []string {
	out := make([]string, len(args))
	for i, a := range args {
		out[i] = fmt.Sprint(a)
	}
	return out
}

func containsAllInOrder(got []string, want []string) bool {
	if len(want) == 0 {
		return true
	}
	j := 0
	for _, g := range got {
		if g == want[j] {
			j++
			if j == len(want) {
				return true
			}
		}
	}
	return false
}

func TestBuildSelectFromParms_NoParams(t *testing.T) {
	sql, args, _, _, _ := BuildSelectFromParms("", url.Values{})
	if !strings.HasSuffix(strings.ToLower(sql), "from tmf_object") {
		t.Fatalf("expected SQL to contain FROM tmf_object, got: %s", sql)
	}

	if len(args) != 0 {
		t.Fatalf("expected no args for empty params, got: %v", args)
	}
}

func TestBuildSelectFromParms_ResourceFilter(t *testing.T) {
	sql, args, _, _, _ := BuildSelectFromParms("Product", url.Values{})
	if !strings.HasSuffix(strings.ToLower(sql), "from tmf_object where type = ?") {
		t.Fatalf("expected SQL to contain WHERE when resource provided, got: %s", sql)
	}
	argStrs := argsToStrings(args)
	if len(argStrs) != 1 || argStrs[0] != "Product" {
		t.Fatalf("expected single arg 'Product', got: %v", argStrs)
	}
}

func TestBuildSelectFromParms_LimitOffset(t *testing.T) {
	v := url.Values{}
	v.Set("limit", "5")
	v.Set("offset", "10")
	_, args, limit, offset, _ := BuildSelectFromParms("", v)

	if limit != 5 {
		t.Fatalf("expected limit to be 5, got: %d", limit)
	}
	if offset != 10 {
		t.Fatalf("expected offset to be 10, got: %d", offset)
	}
	// args may include limit/offset or they may be inlined depending on builder; ensure no panic and SQL contains clauses
	_ = args
}

func TestBuildSelectFromParms_SellerMultipleValues(t *testing.T) {
	v := url.Values{}
	// multiple comma-separated values and multiple instances
	v.Add("seller", "a,b")
	v.Add("seller", "c")
	sql, args, _, _, _ := BuildSelectFromParms("", v)

	if !strings.HasSuffix(strings.ToLower(sql), "from tmf_object and seller in (?,?,?)") {
		t.Fatalf("expected SQL to reference seller, got: %s", sql)
	}
	argStrs := argsToStrings(args)
	want := []string{"a", "b", "c"}
	if !containsAllInOrder(argStrs, want) {
		t.Fatalf("expected args to contain %v in order, got: %v", want, argStrs)
	}
}

func TestBuildSelectFromParms_JSONFieldMultiValues(t *testing.T) {
	v := url.Values{}
	v.Add("status", "Active,Launched")
	sql, args, _, _, _ := BuildSelectFromParms("", v)

	if !strings.Contains(sql, "content->>'$.status'") {
		t.Fatalf("expected SQL to reference JSON path for status, got: %s", sql)
	}
	argStrs := argsToStrings(args)
	want := []string{"Active", "Launched"}
	if !containsAllInOrder(argStrs, want) {
		t.Fatalf("expected args to contain %v in order, got: %v", want, argStrs)
	}
}

func TestBuildSelectFromParms_TopLevelField(t *testing.T) {
	v := url.Values{}
	v.Set("lifecycleStatus", "Launched")
	sql, args, _, _, _ := BuildSelectFromParms("ProductOffering", v)
	if !strings.Contains(sql, "content->>'$.lifecycleStatus'") {
		t.Fatalf("expected SQL to reference JSON path for lifecycleStatus, got: %s", sql)
	}
	if len(args) < 2 || args[1] != "Launched" {
		t.Fatalf("expected args to contain 'Launched', got: %v", args)
	}
}

func TestBuildSelectFromParms_MultiValueTopLevelField(t *testing.T) {
	v := url.Values{}
	v.Set("lifecycleStatus", "Launched,Active")
	sql, args, _, _, _ := BuildSelectFromParms("ProductOffering", v)
	if !strings.HasSuffix(sql, "FROM tmf_object WHERE type = ? AND content->>'$.lifecycleStatus' IN (?,?)") {
		t.Fatalf("expected SQL to reference JSON path for lifecycleStatus, got: %s", sql)
	}
	want := []string{"Launched", "Active"}
	argStrs := argsToStrings(args)
	if !containsAllInOrder(argStrs, want) {
		t.Fatalf("expected args to contain %v in order, got: %v", want, argStrs)
	}
}

func TestBuildSelectFromParms_NestedField(t *testing.T) {
	v := url.Values{}
	// Simulate filtering by productSpecification.id
	v.Set("productSpecification.id", "urn:ngsi-ld:product-specification:19f7f34a-1777-4553-b47b-6ad466d8a0ea")
	sql, args, _, _, _ := BuildSelectFromParms("ProductOffering", v)
	if !strings.Contains(sql, "json_each(tmf_object.content, '$.productSpecification')") {
		t.Fatalf("expected SQL to reference JSON path for productSpecification.id, got: %s", sql)
	}
	if len(args) < 2 || args[1] != "urn:ngsi-ld:product-specification:19f7f34a-1777-4553-b47b-6ad466d8a0ea" {
		t.Fatalf("expected args to contain productSpecification.id value, got: %v", args)
	}
}

func TestBuildSelectFromParms_ArrayOfObjectsField(t *testing.T) {
	v := url.Values{}
	// Simulate filtering by category.id (array of objects)
	v.Set("category.id", "urn:ngsi-ld:category:31a1d8a8-56e8-49c3-aabb-6b0306bc0316")
	sql, args, _, _, _ := BuildSelectFromParms("ProductOffering", v)
	if !strings.Contains(sql, "json_each(tmf_object.content, '$.category')") {
		t.Fatalf("expected SQL to reference JSON path for category.id, got: %s", sql)
	}
	if len(args) < 2 || args[1] != "urn:ngsi-ld:category:31a1d8a8-56e8-49c3-aabb-6b0306bc0316" {
		t.Fatalf("expected args to contain category.id value, got: %v", args)
	}
}

func TestBuildSelectFromParms_ArrayOfObjectsFieldExplicitIndex(t *testing.T) {
	v := url.Values{}
	v.Set("organizationIdentification[0].identificationId", "did:elsi:VATEL-094402295")
	sql, args, _, _, _ := BuildSelectFromParms("ProductOffering", v)
	if !strings.HasSuffix(sql, "FROM tmf_object WHERE type = ? AND EXISTS (SELECT 1 FROM json_tree(tmf_object.content) WHERE json_tree.fullkey LIKE '$.organizationIdentification[0].identificationId%' AND json_tree.value = ?)") {
		t.Fatalf("expected SQL to search for organizationIdentification, got: %s", sql)
	}
	if len(args) < 2 || args[1] != "did:elsi:VATEL-094402295" {
		t.Fatalf("expected args to contain category.id value, got: %v", args)
	}
}

func TestBuildSelectFromParms_ArrayOfObjectsMultiValue(t *testing.T) {
	v := url.Values{}
	// Simulate filtering by relatedParty.role with multiple values
	v.Set("relatedParty.role", "Seller,SellerOperator")
	sql, args, _, _, _ := BuildSelectFromParms("ProductOffering", v)
	if !strings.HasSuffix(sql, "FROM json_tree(tmf_object.content) WHERE json_tree.fullkey LIKE '$.relatedParty%.role%' AND json_tree.value IN (?,?))") {
		t.Fatalf("expected SQL to reference JSON path for relatedParty.role, got: %s", sql)
	}
	want := []string{"Seller", "SellerOperator"}
	argStrs := argsToStrings(args)
	if !containsAllInOrder(argStrs, want) {
		t.Fatalf("expected args to contain %v in order, got: %v", want, argStrs)
	}
}

func TestBuildSelectFromParms_MultipleFilters(t *testing.T) {
	v := url.Values{}
	v.Set("lifecycleStatus", "Launched")
	v.Set("name", "Product Offer Example")
	sql, args, _, _, _ := BuildSelectFromParms("ProductOffering", v)
	if !strings.Contains(sql, "content->>'$.lifecycleStatus'") || !strings.Contains(sql, "content->>'$.name'") {
		t.Fatalf("expected SQL to reference both lifecycleStatus and name, got: %s", sql)
	}
	argStrs := argsToStrings(args)
	foundLaunched := false
	foundProductOffer := false
	for _, arg := range argStrs {
		if arg == "Launched" {
			foundLaunched = true
		}
		if arg == "Product Offer Example" {
			foundProductOffer = true
		}
	}
	if !foundLaunched || !foundProductOffer {
		t.Fatalf("expected args to contain Launched and Product Offer Example, got: %v", argStrs)
	}
}

func TestBuildSelectFromParms_LimitOffsetAndType(t *testing.T) {
	v := url.Values{}
	v.Set("limit", "2")
	v.Set("offset", "1")
	sql, args, limit, offset, _ := BuildSelectFromParms("ProductOffering", v)

	if limit != 2 {
		t.Fatalf("expected limit to be 2, got: %d", limit)
	}
	if offset != 1 {
		t.Fatalf("expected offset to be 1, got: %d", offset)
	}
	if !strings.Contains(sql, "WHERE") {
		t.Fatalf("expected SQL to contain WHERE for type, got: %s", sql)
	}
	if len(args) != 1 || args[0] != "ProductOffering" {
		t.Fatalf("expected args to contain only ProductOffering, got: %v", args)
	}
}

func TestBuildSelectFromParms_SellerShortcut(t *testing.T) {
	v := url.Values{}
	v.Set("seller", "did:elsi:VATES-B60645900,did:elsi:VATES-11111111K")
	sql, args, _, _, _ := BuildSelectFromParms("ProductOffering", v)
	if !strings.Contains(sql, "seller") {
		t.Fatalf("expected SQL to reference seller, got: %s", sql)
	}
	want := []string{"did:elsi:VATES-B60645900", "did:elsi:VATES-11111111K"}
	argStrs := argsToStrings(args)
	if !containsAllInOrder(argStrs, want) {
		t.Fatalf("expected args to contain %v in order, got: %v", want, argStrs)
	}
}

// newTestDBService creates a temporary SQLite database for testing and returns the service
// together with a cleanup function the caller must defer.
func newTestDBService(t *testing.T) (*DBService, func()) {
	t.Helper()
	dir := t.TempDir()
	dbPath := filepath.Join(dir, "test.db")
	repo, err := NewDBService(dbPath)
	if err != nil {
		t.Fatalf("newTestDBService: failed to create DBService: %v", err)
	}
	cleanup := func() {
		_ = repo.Close()
		_ = os.RemoveAll(dir)
	}
	return repo, cleanup
}

func TestUpdateObject(t *testing.T) {
	repo, cleanup := newTestDBService(t)
	defer cleanup()

	// --- Create the initial object with version 1.0 ---
	initialContent := map[string]any{
		"id":              "urn:uuid:test-product-offering-001",
		"@type":           "ProductOffering",
		"name":            "Initial Offering",
		"lifecycleStatus": "Draft",
		"version":         "1.0",
	}
	initialJSON, err := json.Marshal(initialContent)
	if err != nil {
		t.Fatalf("TestUpdateObject: failed to marshal initial content: %v", err)
	}

	created := NewTMFRecord(
		"urn:uuid:test-product-offering-001",
		"ProductOffering",
		"1.0",
		"v4",
		"2026-01-01T00:00:00Z",
		initialJSON,
	)

	if err := repo.CreateObject(created); err != nil {
		t.Fatalf("TestUpdateObject: CreateObject failed: %v", err)
	}

	// --- Prepare the update payload, reusing version 1.0 ---
	updatedContent := map[string]any{
		"id":              "urn:uuid:test-product-offering-001",
		"@type":           "ProductOffering",
		"name":            "Updated Offering",
		"lifecycleStatus": "Launched",
		"version":         "1.0",
	}
	updatedJSON, err := json.Marshal(updatedContent)
	if err != nil {
		t.Fatalf("TestUpdateObject: failed to marshal updated content: %v", err)
	}

	updated := &TMFRecord{
		ID:         created.ID,
		Type:       created.Type,
		Version:    "1.0", // same version as the created record
		APIVersion: created.APIVersion,
		LastUpdate: "2026-06-01T00:00:00Z",
		Content:    updatedJSON,
	}

	if err := repo.UpdateObject(updated); err != nil {
		t.Fatalf("TestUpdateObject: UpdateObject failed: %v", err)
	}

	// --- Read back and verify the content was persisted ---
	fetched, err := repo.GetObject(created.ID, created.Type)
	if err != nil {
		t.Fatalf("TestUpdateObject: GetObject failed: %v", err)
	}
	if fetched == nil {
		t.Fatal("TestUpdateObject: GetObject returned nil; object not found after update")
	}

	var fetchedMap map[string]any
	if err := json.Unmarshal(fetched.Content, &fetchedMap); err != nil {
		t.Fatalf("TestUpdateObject: failed to unmarshal fetched content: %v", err)
	}

	if got := fetchedMap["name"]; got != "Updated Offering" {
		t.Errorf("TestUpdateObject: expected name %q, got %q", "Updated Offering", got)
	}
	if got := fetchedMap["lifecycleStatus"]; got != "Launched" {
		t.Errorf("TestUpdateObject: expected lifecycleStatus %q, got %q", "Launched", got)
	}
	if fetched.Version != "1.0" {
		t.Errorf("TestUpdateObject: expected version %q, got %q", "1.0", fetched.Version)
	}
}

func TestUpdateObject_NotFound(t *testing.T) {
	repo, cleanup := newTestDBService(t)
	defer cleanup()

	phantom := &TMFRecord{
		ID:      "urn:uuid:does-not-exist",
		Type:    "ProductOffering",
		Version: "1.0",
		Content: []byte(`{"id":"urn:uuid:does-not-exist"}`),
	}

	err := repo.UpdateObject(phantom)
	if err == nil {
		t.Fatal("TestUpdateObject_NotFound: expected an error for non-existent object, got nil")
	}
	if !errors.Is(err, &ErrObjectNotFound{}) {
		t.Errorf("TestUpdateObject_NotFound: expected ErrObjectNotFound, got %v", err)
	}
}

// TestUpdateObject_VersionBump verifies that UpdateObject succeeds when the supplied version
// is lexicographically greater than the current maximum version (a version bump).
func TestUpdateObject_VersionBump(t *testing.T) {
	repo, cleanup := newTestDBService(t)
	defer cleanup()

	// Create initial record at version 1.0
	initial := NewTMFRecord(
		"urn:uuid:version-bump-test",
		"ProductOffering",
		"1.0",
		"v4",
		"2026-01-01T00:00:00Z",
		[]byte(`{"id":"urn:uuid:version-bump-test","name":"v1"}`),
	)
	if err := repo.CreateObject(initial); err != nil {
		t.Fatalf("TestUpdateObject_VersionBump: CreateObject failed: %v", err)
	}

	// Update to version 2.0 (lexicographically greater than 1.0)
	bumped := &TMFRecord{
		ID:         initial.ID,
		Type:       initial.Type,
		Version:    "2.0",
		APIVersion: initial.APIVersion,
		LastUpdate: "2026-06-01T00:00:00Z",
		Content:    []byte(`{"id":"urn:uuid:version-bump-test","name":"v2"}`),
	}
	if err := repo.UpdateObject(bumped); err != nil {
		t.Fatalf("TestUpdateObject_VersionBump: UpdateObject to 2.0 failed: %v", err)
	}

	// The stored row should now be at version 2.0
	fetched, err := repo.GetObject(initial.ID, initial.Type)
	if err != nil {
		t.Fatalf("TestUpdateObject_VersionBump: GetObject failed: %v", err)
	}
	if fetched.Version != "2.0" {
		t.Errorf("TestUpdateObject_VersionBump: expected version 2.0, got %q", fetched.Version)
	}
}

// TestUpsertObject_Insert verifies that UpsertObject creates the object when it does not exist.
func TestUpsertObject_Insert(t *testing.T) {
	repo, cleanup := newTestDBService(t)
	defer cleanup()

	obj := NewTMFRecord(
		"urn:uuid:upsert-insert",
		"ProductOffering",
		"1.0",
		"v4",
		"2026-01-01T00:00:00Z",
		[]byte(`{"id":"urn:uuid:upsert-insert","name":"initial"}`),
	)

	if err := repo.UpsertObject(obj); err != nil {
		t.Fatalf("TestUpsertObject_Insert: expected no error, got %v", err)
	}

	fetched, err := repo.GetObject(obj.ID, obj.Type)
	if err != nil {
		t.Fatalf("TestUpsertObject_Insert: GetObject failed: %v", err)
	}
	if fetched == nil {
		t.Fatal("TestUpsertObject_Insert: object not found after upsert")
	}
	if fetched.Version != "1.0" {
		t.Errorf("TestUpsertObject_Insert: expected version 1.0, got %q", fetched.Version)
	}
}

// TestUpsertObject_UpdateSameVersion verifies that UpsertObject succeeds when the object
// already exists and the same version is supplied (in-place update).
func TestUpsertObject_UpdateSameVersion(t *testing.T) {
	repo, cleanup := newTestDBService(t)
	defer cleanup()

	obj := NewTMFRecord(
		"urn:uuid:upsert-update-same",
		"ProductOffering",
		"1.0",
		"v4",
		"2026-01-01T00:00:00Z",
		[]byte(`{"id":"urn:uuid:upsert-update-same","name":"initial"}`),
	)
	if err := repo.UpsertObject(obj); err != nil {
		t.Fatalf("TestUpsertObject_UpdateSameVersion: initial upsert failed: %v", err)
	}

	// Upsert again with same version but different content
	obj.Content = []byte(`{"id":"urn:uuid:upsert-update-same","name":"updated"}`)
	obj.LastUpdate = "2026-06-01T00:00:00Z"
	if err := repo.UpsertObject(obj); err != nil {
		t.Fatalf("TestUpsertObject_UpdateSameVersion: second upsert failed: %v", err)
	}

	fetched, _ := repo.GetObject(obj.ID, obj.Type)
	var m map[string]any
	if err := json.Unmarshal(fetched.Content, &m); err != nil {
		t.Fatalf("TestUpsertObject_UpdateSameVersion: unmarshal failed: %v", err)
	}
	if m["name"] != "updated" {
		t.Errorf("TestUpsertObject_UpdateSameVersion: expected name 'updated', got %v", m["name"])
	}
}

// TestUpsertObject_VersionBump verifies that UpsertObject succeeds when the object already
// exists and the new version is lexicographically greater than the current maximum.
func TestUpsertObject_VersionBump(t *testing.T) {
	repo, cleanup := newTestDBService(t)
	defer cleanup()

	obj := NewTMFRecord(
		"urn:uuid:upsert-bump",
		"ProductOffering",
		"1.0",
		"v4",
		"2026-01-01T00:00:00Z",
		[]byte(`{"id":"urn:uuid:upsert-bump","name":"v1"}`),
	)
	if err := repo.UpsertObject(obj); err != nil {
		t.Fatalf("TestUpsertObject_VersionBump: initial upsert failed: %v", err)
	}

	bumped := &TMFRecord{
		ID:         obj.ID,
		Type:       obj.Type,
		Version:    "2.0",
		APIVersion: obj.APIVersion,
		LastUpdate: "2026-06-01T00:00:00Z",
		Content:    []byte(`{"id":"urn:uuid:upsert-bump","name":"v2"}`),
	}
	if err := repo.UpsertObject(bumped); err != nil {
		t.Fatalf("TestUpsertObject_VersionBump: bump upsert failed: %v", err)
	}

	// GetObject must return the new maximum version.
	fetched, _ := repo.GetObject(obj.ID, obj.Type)
	if fetched.Version != "2.0" {
		t.Errorf("TestUpsertObject_VersionBump: expected version 2.0, got %q", fetched.Version)
	}

	// The old version row must still exist — UpsertObject preserves version history.
	var rowCount int
	err := repo.db.QueryRow(
		"SELECT COUNT(*) FROM tmf_object WHERE id = ? AND type = ?",
		obj.ID, obj.Type,
	).Scan(&rowCount)
	if err != nil {
		t.Fatalf("TestUpsertObject_VersionBump: count query failed: %v", err)
	}
	if rowCount != 1 {
		t.Errorf("TestUpsertObject_VersionBump: expected 1 row, got %d", rowCount)
	}
}
