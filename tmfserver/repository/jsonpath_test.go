package repository

import (
	"database/sql"
	"encoding/json"
	"os"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestJSONPath_SimpleProperty(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantPath   string
		wantSelect string
		wantWhere  string
	}{
		{
			name:       "bare name",
			path:       "name",
			wantPath:   "$.name",
			wantSelect: "tmf_object.content ->> '$.name'",
			wantWhere:  "tmf_object.content ->> '$.name' = ?",
		},
		{
			name:       "dollar name",
			path:       "$.name",
			wantPath:   "$.name",
			wantSelect: "tmf_object.content ->> '$.name'",
			wantWhere:  "tmf_object.content ->> '$.name' = ?",
		},
		{
			name:       "nested object path",
			path:       "place.address.city",
			wantPath:   "$.place.address.city",
			wantSelect: "tmf_object.content ->> '$.place.address.city'",
			wantWhere:  "tmf_object.content ->> '$.place.address.city' = ?",
		},
		{
			name:       "bracket property notation",
			path:       "$['lifecycleStatus']",
			wantPath:   "$.lifecycleStatus",
			wantSelect: "tmf_object.content ->> '$.lifecycleStatus'",
			wantWhere:  "tmf_object.content ->> '$.lifecycleStatus' = ?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seg, err := MapJSONPathToSQL("tmf_object", "content", tt.path, "=", "Launched")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if seg.NormalizedPath != tt.wantPath {
				t.Errorf("NormalizedPath: got %q, want %q", seg.NormalizedPath, tt.wantPath)
			}
			if seg.SelectExpr != tt.wantSelect {
				t.Errorf("SelectExpr: got %q, want %q", seg.SelectExpr, tt.wantSelect)
			}
			if seg.WhereExpr != tt.wantWhere {
				t.Errorf("WhereExpr: got %q, want %q", seg.WhereExpr, tt.wantWhere)
			}
			if len(seg.Args) != 1 || seg.Args[0] != "Launched" {
				t.Errorf("Args: got %v, want ['Launched']", seg.Args)
			}
			if seg.IsArray {
				t.Errorf("expected IsArray to be false")
			}
		})
	}
}

func TestJSONPath_ArrayIndex(t *testing.T) {
	seg, err := MapJSONPathToSQL("tmf_object", "content", "attachment[0].url", "=", "https://example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if seg.NormalizedPath != "$.attachment[0].url" {
		t.Errorf("NormalizedPath: got %q, want '$.attachment[0].url'", seg.NormalizedPath)
	}
	if seg.SelectExpr != "tmf_object.content ->> '$.attachment[0].url'" {
		t.Errorf("SelectExpr: got %q", seg.SelectExpr)
	}
	if seg.WhereExpr != "tmf_object.content ->> '$.attachment[0].url' = ?" {
		t.Errorf("WhereExpr: got %q", seg.WhereExpr)
	}
	if len(seg.Args) != 1 || seg.Args[0] != "https://example.com" {
		t.Errorf("Args: got %v", seg.Args)
	}
}

func TestJSONPath_ArrayWildcard(t *testing.T) {
	seg, err := MapJSONPathToSQL("tmf_object", "content", "category[*].id", "=", "cat-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !seg.IsArray {
		t.Errorf("expected IsArray to be true")
	}
	wantSelect := "(SELECT json_group_array(json_each.value ->> '$.id') FROM json_each(tmf_object.content, '$.category'))"
	if seg.SelectExpr != wantSelect {
		t.Errorf("SelectExpr: got %q, want %q", seg.SelectExpr, wantSelect)
	}
	wantWhere := "EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.category') WHERE json_each.value ->> '$.id' = ?)"
	if seg.WhereExpr != wantWhere {
		t.Errorf("WhereExpr: got %q, want %q", seg.WhereExpr, wantWhere)
	}
	if len(seg.Args) != 1 || seg.Args[0] != "cat-123" {
		t.Errorf("Args: got %v, want ['cat-123']", seg.Args)
	}
}

func TestJSONPath_ArrayWildcard_MultipleValues(t *testing.T) {
	seg, err := MapJSONPathToSQL("tmf_object", "content", "category[*].id", "IN", "cat-1", "cat-2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantWhere := "EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.category') WHERE json_each.value ->> '$.id' IN (?, ?))"
	if seg.WhereExpr != wantWhere {
		t.Errorf("WhereExpr: got %q, want %q", seg.WhereExpr, wantWhere)
	}
	if len(seg.Args) != 2 || seg.Args[0] != "cat-1" || seg.Args[1] != "cat-2" {
		t.Errorf("Args: got %v, want ['cat-1', 'cat-2']", seg.Args)
	}
}

func TestJSONPath_ArrayFilterPredicate(t *testing.T) {
	seg, err := MapJSONPathToSQL("tmf_object", "content", "relatedParty[?(@.role == 'seller')].id", "=", "urn:org:123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !seg.IsPredicate {
		t.Errorf("expected IsPredicate to be true")
	}

	wantSelect := "(SELECT json_group_array(json_each.value ->> '$.id') FROM json_each(tmf_object.content, '$.relatedParty') WHERE json_each.value ->> '$.role' = ?)"
	if seg.SelectExpr != wantSelect {
		t.Errorf("SelectExpr: got %q, want %q", seg.SelectExpr, wantSelect)
	}

	wantWhere := "EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.relatedParty') WHERE json_each.value ->> '$.role' = ? AND json_each.value ->> '$.id' = ?)"
	if seg.WhereExpr != wantWhere {
		t.Errorf("WhereExpr: got %q, want %q", seg.WhereExpr, wantWhere)
	}

	if len(seg.Args) != 2 || seg.Args[0] != "seller" || seg.Args[1] != "urn:org:123" {
		t.Errorf("Args: got %v, want ['seller', 'urn:org:123']", seg.Args)
	}
}

func TestJSONPath_ArrayFilterPredicate_NumericAndOperators(t *testing.T) {
	seg, err := MapJSONPathToSQL("tmf_object", "content", "productCharacteristic[?(@.value > 100)].name", "LIKE", "%speed%")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantWhere := "EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.productCharacteristic') WHERE json_each.value ->> '$.value' > ? AND json_each.value ->> '$.name' LIKE ?)"
	if seg.WhereExpr != wantWhere {
		t.Errorf("WhereExpr: got %q, want %q", seg.WhereExpr, wantWhere)
	}

	if len(seg.Args) != 2 || seg.Args[0] != int64(100) || seg.Args[1] != "%speed%" {
		t.Errorf("Args: got %v, want [100, '%%speed%%']", seg.Args)
	}
}

func TestJSONPath_MultiLevelWildcards(t *testing.T) {
	seg, err := MapJSONPathToSQL("tmf_object", "content", "items[*].components[*].id", "=", "comp-99")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantWhere := "EXISTS (SELECT 1 FROM json_tree(tmf_object.content) WHERE json_tree.fullkey LIKE '$.items[%].components[%].id' AND json_tree.value = ?)"
	if seg.WhereExpr != wantWhere {
		t.Errorf("WhereExpr: got %q, want %q", seg.WhereExpr, wantWhere)
	}
	if len(seg.Args) != 1 || seg.Args[0] != "comp-99" {
		t.Errorf("Args: got %v, want ['comp-99']", seg.Args)
	}
}

func TestJSONPath_RecursiveDescent(t *testing.T) {
	seg, err := MapJSONPathToSQL("tmf_object", "content", "..identificationId", "=", "id-456")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantWhere := "EXISTS (SELECT 1 FROM json_tree(tmf_object.content) WHERE json_tree.fullkey LIKE '$%.identificationId' AND json_tree.value = ?)"
	if seg.WhereExpr != wantWhere {
		t.Errorf("WhereExpr: got %q, want %q", seg.WhereExpr, wantWhere)
	}
}

func TestJSONPath_RawJSONOption(t *testing.T) {
	seg, err := GenerateJSONPathQuerySegments(JSONPathQueryOptions{
		TableName:  "tmf_object",
		ColumnName: "content",
		Path:       "category[*]",
		RawJSON:    true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantSelect := "(SELECT json_group_array(json_each.value) FROM json_each(tmf_object.content, '$.category'))"
	if seg.SelectExpr != wantSelect {
		t.Errorf("SelectExpr: got %q, want %q", seg.SelectExpr, wantSelect)
	}
}

func TestJSONPath_InvalidInput(t *testing.T) {
	badInputs := []struct {
		name       string
		tableName  string
		columnName string
		path       string
	}{
		{"empty path", "tmf_object", "content", ""},
		{"invalid table name", "tmf;drop table", "content", "name"},
		{"invalid column name", "tmf_object", "content--", "name"},
		{"unclosed bracket", "tmf_object", "content", "items[0"},
		{"malformed filter", "tmf_object", "content", "items[?(@.foo)]"},
	}

	for _, tt := range badInputs {
		t.Run(tt.name, func(t *testing.T) {
			_, err := MapJSONPathToSQL(tt.tableName, tt.columnName, tt.path, "=")
			if err == nil {
				t.Errorf("expected error for input %v, got nil", tt)
			}
		})
	}
}

func TestJSONPath_SQLiteExecution(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open sqlite in-memory db: %v", err)
	}
	defer db.Close()

	// Create table and insert test row
	_, err = db.Exec(`CREATE TABLE tmf_object (
		id TEXT PRIMARY KEY,
		type TEXT,
		seller TEXT,
		content BLOB
	)`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	testJSON := `{
		"name": "Fiber Internet 500Mbps",
		"lifecycleStatus": "Launched",
		"category": [
			{"id": "cat-broadband", "name": "Broadband"},
			{"id": "cat-fiber", "name": "Fiber"}
		],
		"relatedParty": [
			{"role": "seller", "id": "urn:org:telecom"},
			{"role": "buyer", "id": "urn:org:customer"}
		],
		"productCharacteristic": [
			{"name": "bandwidth", "value": 500}
		]
	}`

	// Store as JSONB if supported, or JSON
	_, err = db.Exec(`INSERT INTO tmf_object (id, type, seller, content) VALUES (?, ?, ?, jsonb(?))`,
		"po-100", "ProductOffering", "urn:org:telecom", testJSON)
	if err != nil {
		// Fallback to json() if jsonb() is not compiled into this sqlite driver build
		_, err = db.Exec(`INSERT INTO tmf_object (id, type, seller, content) VALUES (?, ?, ?, json(?))`,
			"po-100", "ProductOffering", "urn:org:telecom", testJSON)
		if err != nil {
			t.Fatalf("failed to insert test data: %v", err)
		}
	}

	// Test 1: Query scalar property
	seg1, err := MapJSONPathToSQL("tmf_object", "content", "$.lifecycleStatus", "=", "Launched")
	if err != nil {
		t.Fatalf("failed to generate segments: %v", err)
	}
	query1 := "SELECT id, " + seg1.SelectExpr + " FROM tmf_object WHERE type = ? AND " + seg1.WhereExpr
	args1 := append([]any{"ProductOffering"}, seg1.Args...)
	var id1, val1 string
	err = db.QueryRow(query1, args1...).Scan(&id1, &val1)
	if err != nil {
		t.Fatalf("query 1 execution failed: %v\nQuery: %s\nArgs: %v", err, query1, args1)
	}
	if id1 != "po-100" || val1 != "Launched" {
		t.Errorf("query 1 unexpected result: id=%q, val=%q", id1, val1)
	}

	// Test 2: Array wildcard filter (category[*].id == 'cat-fiber')
	seg2, err := MapJSONPathToSQL("tmf_object", "content", "category[*].id", "=", "cat-fiber")
	if err != nil {
		t.Fatalf("failed to generate segments: %v", err)
	}
	query2 := "SELECT id, " + seg2.SelectExpr + " FROM tmf_object WHERE " + seg2.WhereExpr
	var id2, categoriesJSON string
	err = db.QueryRow(query2, seg2.Args...).Scan(&id2, &categoriesJSON)
	if err != nil {
		t.Fatalf("query 2 execution failed: %v\nQuery: %s\nArgs: %v", err, query2, seg2.Args)
	}
	if id2 != "po-100" || !strings.Contains(categoriesJSON, "cat-fiber") {
		t.Errorf("query 2 unexpected result: id=%q, categories=%q", id2, categoriesJSON)
	}

	// Test 3: Array predicate filter (relatedParty[?(@.role == 'seller')].id)
	seg3, err := MapJSONPathToSQL("tmf_object", "content", "relatedParty[?(@.role == 'seller')].id", "=", "urn:org:telecom")
	if err != nil {
		t.Fatalf("failed to generate segments: %v", err)
	}
	query3 := "SELECT id, " + seg3.SelectExpr + " FROM tmf_object WHERE " + seg3.WhereExpr
	args3 := append([]any{}, seg3.SelectArgs...)
	args3 = append(args3, seg3.WhereArgs...)
	var id3, sellerIDs string
	err = db.QueryRow(query3, args3...).Scan(&id3, &sellerIDs)
	if err != nil {
		t.Fatalf("query 3 execution failed: %v\nQuery: %s\nArgs: %v", err, query3, args3)
	}
	if id3 != "po-100" || !strings.Contains(sellerIDs, "urn:org:telecom") {
		t.Errorf("query 3 unexpected result: id=%q, sellerIDs=%q", id3, sellerIDs)
	}
}

func TestParseFilterExpression(t *testing.T) {
	tests := []struct {
		input      string
		wantPath   string
		wantOp     string
		wantValues []any
	}{
		{
			input:      "category[*].id == 'cat-123'",
			wantPath:   "category[*].id",
			wantOp:     "=",
			wantValues: []any{"cat-123"},
		},
		{
			input:      "name LIKE 'Fiber%'",
			wantPath:   "name",
			wantOp:     "LIKE",
			wantValues: []any{"Fiber%"},
		},
		{
			input:      "productCharacteristic[?(@.name == 'bandwidth')].value >= 500",
			wantPath:   "productCharacteristic[?(@.name == 'bandwidth')].value",
			wantOp:     ">=",
			wantValues: []any{int64(500)},
		},
		{
			input:      "status IN ('Active', 'Launched')",
			wantPath:   "status",
			wantOp:     "IN",
			wantValues: []any{"Active", "Launched"},
		},
		{
			input:      "validFor.endDateTime IS NOT NULL",
			wantPath:   "validFor.endDateTime",
			wantOp:     "IS NOT NULL",
			wantValues: nil,
		},
		{
			input:      "lifecycleStatus",
			wantPath:   "lifecycleStatus",
			wantOp:     "IS NOT NULL",
			wantValues: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			path, op, vals, err := ParseFilterExpression(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if path != tt.wantPath {
				t.Errorf("path: got %q, want %q", path, tt.wantPath)
			}
			if op != tt.wantOp {
				t.Errorf("op: got %q, want %q", op, tt.wantOp)
			}
			if len(vals) != len(tt.wantValues) {
				t.Fatalf("values length: got %d, want %d", len(vals), len(tt.wantValues))
			}
			for i := range vals {
				if vals[i] != tt.wantValues[i] {
					t.Errorf("val[%d]: got %v, want %v", i, vals[i], tt.wantValues[i])
				}
			}
		})
	}
}

func TestGenerateCombinedFilterQuery(t *testing.T) {
	t.Run("single expression", func(t *testing.T) {
		where, args, err := GenerateCombinedFilterQuery("tmf_object", "content", []string{"name == 'Fiber'"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if where != "tmf_object.content ->> '$.name' = ?" {
			t.Errorf("got where: %q", where)
		}
		if len(args) != 1 || args[0] != "Fiber" {
			t.Errorf("got args: %v", args)
		}
	})

	t.Run("OR expressions separated by semicolon", func(t *testing.T) {
		where, args, err := GenerateCombinedFilterQuery("tmf_object", "content", []string{
			"lifecycleStatus == 'Launched';lifecycleStatus == 'Active'",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "(tmf_object.content ->> '$.lifecycleStatus' = ? OR tmf_object.content ->> '$.lifecycleStatus' = ?)"
		if where != want {
			t.Errorf("got where: %q, want: %q", where, want)
		}
		if len(args) != 2 || args[0] != "Launched" || args[1] != "Active" {
			t.Errorf("got args: %v", args)
		}
	})

	t.Run("AND expressions via multiple parameters", func(t *testing.T) {
		where, args, err := GenerateCombinedFilterQuery("tmf_object", "content", []string{
			"lifecycleStatus == 'Launched'",
			"category[*].id == 'cat-123'",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "(tmf_object.content ->> '$.lifecycleStatus' = ? AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.category') WHERE json_each.value ->> '$.id' = ?))"
		if where != want {
			t.Errorf("got where: %q, want: %q", where, want)
		}
		if len(args) != 2 || args[0] != "Launched" || args[1] != "cat-123" {
			t.Errorf("got args: %v", args)
		}
	})

	t.Run("AND of OR groups", func(t *testing.T) {
		where, args, err := GenerateCombinedFilterQuery("tmf_object", "content", []string{
			"lifecycleStatus == 'Launched';lifecycleStatus == 'Active'",
			"category[*].id == 'cat-1';category[*].id == 'cat-2'",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := "((tmf_object.content ->> '$.lifecycleStatus' = ? OR tmf_object.content ->> '$.lifecycleStatus' = ?) AND (EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.category') WHERE json_each.value ->> '$.id' = ?) OR EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.category') WHERE json_each.value ->> '$.id' = ?)))"
		if where != want {
			t.Errorf("got where: %q, want: %q", where, want)
		}
		if len(args) != 4 || args[0] != "Launched" || args[1] != "Active" || args[2] != "cat-1" || args[3] != "cat-2" {
			t.Errorf("got args: %v", args)
		}
	})

	t.Run("semicolon inside quoted string is preserved", func(t *testing.T) {
		where, args, err := GenerateCombinedFilterQuery("tmf_object", "content", []string{
			"name == 'Fiber; 100Mbps'",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if where != "tmf_object.content ->> '$.name' = ?" {
			t.Errorf("got where: %q", where)
		}
		if len(args) != 1 || args[0] != "Fiber; 100Mbps" {
			t.Errorf("got args: %v", args)
		}
	})
}

func TestBuildSelectFromParms_FilterParam(t *testing.T) {
	// Test BuildSelectFromParms with ?filter=lifecycleStatus == 'Launched';lifecycleStatus == 'Active'&filter=category[*].id == 'cat-1'
	v := make(map[string][]string)
	v["filter"] = []string{
		"lifecycleStatus == 'Launched';lifecycleStatus == 'Active'",
		"category[*].id == 'cat-1'",
	}
	sql, args, _, _, err := BuildSelectFromParms("ProductOffering", v)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(sql, "FROM tmf_object WHERE type = ? AND ((tmf_object.content ->> '$.lifecycleStatus' = ? OR tmf_object.content ->> '$.lifecycleStatus' = ?) AND EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.category') WHERE json_each.value ->> '$.id' = ?))") {
		t.Errorf("unexpected SQL: %s", sql)
	}

	if len(args) != 4 || args[0] != "ProductOffering" || args[1] != "Launched" || args[2] != "Active" || args[3] != "cat-1" {
		t.Errorf("unexpected args: %v", args)
	}
}

// loadTestOfferingsDB creates an in-memory SQLite database with the tmf_object schema
// and inserts all 10 product offerings from testdata/product_offerings.json as binary jsonb.
func loadTestOfferingsDB(t *testing.T) *sql.DB {
	t.Helper()

	data, err := os.ReadFile("testdata/product_offerings.json")
	if err != nil {
		t.Fatalf("failed to read testdata/product_offerings.json: %v", err)
	}

	var offerings []map[string]any
	if err := json.Unmarshal(data, &offerings); err != nil {
		t.Fatalf("failed to unmarshal testdata/product_offerings.json: %v", err)
	}

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory sqlite db: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tmf_object (
		"id" TEXT NOT NULL,
		"type" TEXT NOT NULL,
		"version" TEXT DEFAULT '',
		"api_version" TEXT DEFAULT '',
		"seller" TEXT DEFAULT '',
		"seller_operator" TEXT DEFAULT '',
		"buyer" TEXT DEFAULT '',
		"buyer_operator" TEXT DEFAULT '',
		"last_update" TEXT DEFAULT '',
		"content" BLOB NOT NULL,
		"random" INTEGER DEFAULT 0,
		"created_at" INTEGER,
		"updated_at" INTEGER,
		PRIMARY KEY ("id", "type")
	);`)
	if err != nil {
		t.Fatalf("failed to create tmf_object table: %v", err)
	}

	for _, off := range offerings {
		id, _ := off["id"].(string)
		version, _ := off["version"].(string)
		lastUpdate, _ := off["lastUpdate"].(string)

		var seller, sellerOperator string
		if rps, ok := off["relatedParty"].([]any); ok {
			for _, rpRaw := range rps {
				if rp, ok := rpRaw.(map[string]any); ok {
					role, _ := rp["role"].(string)
					name, _ := rp["name"].(string)
					switch role {
					case "Seller":
						seller = name
					case "SellerOperator":
						sellerOperator = name
					}
				}
			}
		}

		contentBytes, err := json.Marshal(off)
		if err != nil {
			t.Fatalf("failed to marshal offering %s: %v", id, err)
		}

		// Insert content as jsonb(), matching production storage
		_, err = db.Exec(`INSERT INTO tmf_object (
			id, type, version, seller, seller_operator, last_update, content, created_at, updated_at
		) VALUES (?, 'ProductOffering', ?, ?, ?, ?, jsonb(?), unixepoch(), unixepoch())`,
			id, version, seller, sellerOperator, lastUpdate, contentBytes,
		)
		if err != nil {
			t.Fatalf("failed to insert offering %s with jsonb(): %v", id, err)
		}
	}

	return db
}

func TestJSONPath_ProductOfferings_RealDB(t *testing.T) {
	db := loadTestOfferingsDB(t)
	defer db.Close()

	// Verify all 10 offerings were inserted
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM tmf_object WHERE type = 'ProductOffering'").Scan(&count)
	if err != nil {
		t.Fatalf("failed to count offerings: %v", err)
	}
	if count != 10 {
		t.Fatalf("expected 10 offerings inserted, got: %d", count)
	}

	tests := []struct {
		name          string
		filters       []string // expressions passed as &filter=
		wantCount     int
		mustIncludeID string // optional ID expected in result
		mustExcludeID string // optional ID expected NOT in result
	}{
		// 1. Top-Level Scalar Properties
		{
			name:          "exact name equality",
			filters:       []string{"name == 'Gas Emission Monitoring IoT'"},
			wantCount:     1,
			mustIncludeID: "urn:ngsi-ld:product-offering:cd390626-e57e-41e7-b575-7c6d9f71a2ef",
		},
		{
			name:          "lifecycleStatus Obsolete",
			filters:       []string{"lifecycleStatus == 'Obsolete'"},
			wantCount:     1,
			mustIncludeID: "urn:ngsi-ld:product-offering:1f835fe9-2527-4579-9609-24858cdb4b8e", // Google Workspace
		},
		{
			name:          "lifecycleStatus Launched",
			filters:       []string{"lifecycleStatus == 'Launched'"},
			wantCount:     9,
			mustExcludeID: "urn:ngsi-ld:product-offering:1f835fe9-2527-4579-9609-24858cdb4b8e",
		},
		{
			name:          "dollar-prefixed path",
			filters:       []string{"$.version == '5.0'"},
			wantCount:     1,
			mustIncludeID: "urn:ngsi-ld:product-offering:f4b1fd69-eb0e-41b2-a20a-d1047ba14ddb", // Voting Platform
		},
		{
			name:          "bracket property notation",
			filters:       []string{"$['name'] == 'CloudFerro Cloud'"},
			wantCount:     1,
			mustIncludeID: "urn:ngsi-ld:product-offering:17e4bfb5-8c36-4cc3-bcb1-b4b02c33359d",
		},
		{
			name:      "boolean property equality (isBundle == false)",
			filters:   []string{"isBundle == false"},
			wantCount: 10,
		},

		// 2. Nested Object Properties
		{
			name:          "nested productSpecification.name",
			filters:       []string{"productSpecification.name == 'Voting'"},
			wantCount:     1,
			mustIncludeID: "urn:ngsi-ld:product-offering:f4b1fd69-eb0e-41b2-a20a-d1047ba14ddb",
		},
		{
			name:          "nested productSpecification ID",
			filters:       []string{"productSpecification.id == 'urn:ngsi-ld:product-specification:51c70fb6-514e-48d3-bd70-84f0ab379358'"},
			wantCount:     1,
			mustIncludeID: "urn:ngsi-ld:product-offering:1f835fe9-2527-4579-9609-24858cdb4b8e", // Google Workspace
		},

		// 3. Array Indexing
		{
			name:          "first category name Infrastructure",
			filters:       []string{"category[0].name == 'Infrastructure'"},
			wantCount:     2, // Gas Emission, Nivola PostgreSQL
			mustIncludeID: "urn:ngsi-ld:product-offering:cd390626-e57e-41e7-b575-7c6d9f71a2ef",
		},
		{
			name:      "first relatedParty role SellerOperator",
			filters:   []string{"relatedParty[0].role == 'SellerOperator'"},
			wantCount: 10,
		},
		{
			name:          "second relatedParty name VATRO-1572582",
			filters:       []string{"relatedParty[1].name == 'VATRO-1572582'"},
			wantCount:     4, // 4 BEIA offerings
			mustIncludeID: "urn:ngsi-ld:product-offering:cd390626-e57e-41e7-b575-7c6d9f71a2ef",
		},

		// 4. Array Wildcard ([*])
		{
			name:          "wildcard category name Databases",
			filters:       []string{"category[*].name == 'Databases'"},
			wantCount:     2, // Nivola PostgreSQL & Nivola Oracle
			mustIncludeID: "urn:ngsi-ld:product-offering:73a949ac-a492-4de4-a106-35de37261cd2",
		},
		{
			name:          "wildcard category name Agriculture",
			filters:       []string{"category[*].name == 'Agriculture'"},
			wantCount:     2, // IoT Coverage & CloudFerro Cloud
			mustIncludeID: "urn:ngsi-ld:product-offering:48c7f488-6aa2-4c8c-a2f9-5b8aba8fb616",
		},
		{
			name:          "wildcard category name Software Services",
			filters:       []string{"category[*].name == 'Software Services'"},
			wantCount:     4, // 4 BEIA offerings
			mustIncludeID: "urn:ngsi-ld:product-offering:4de370e0-fba8-4d49-95eb-013ebc5eb7b5",
		},
		{
			name:          "wildcard productOfferingPrice name Recurring Payment",
			filters:       []string{"productOfferingPrice[*].name == 'Recurring Payment'"},
			wantCount:     2, // Gas Emission & Radon
			mustIncludeID: "urn:ngsi-ld:product-offering:2825967d-0565-442e-a5d9-df9d3df2f5df",
		},

		// 5. Array Wildcard with IN operator
		{
			name:          "category wildcard IN (Databases, Agriculture)",
			filters:       []string{"category[*].name IN ('Databases', 'Agriculture')"},
			wantCount:     4, // 2 databases + 2 agriculture
			mustIncludeID: "urn:ngsi-ld:product-offering:73a949ac-a492-4de4-a106-35de37261cd2",
		},
		{
			name:          "scalar version IN (5.0, 1.0)",
			filters:       []string{"version IN ('5.0', '1.0')"},
			wantCount:     3, // Voting (5.0), CloudFerro (1.0), Smart Parking (1.0)
			mustIncludeID: "urn:ngsi-ld:product-offering:f4b1fd69-eb0e-41b2-a20a-d1047ba14ddb",
		},

		// 6. Array Filter Predicates ([?(@.field == 'val')])
		{
			name:          "relatedParty filter by role Seller matching VATRO-1572582",
			filters:       []string{"relatedParty[?(@.role == 'Seller')].name == 'VATRO-1572582'"},
			wantCount:     4, // All BEIA offerings
			mustIncludeID: "urn:ngsi-ld:product-offering:cd390626-e57e-41e7-b575-7c6d9f71a2ef",
			mustExcludeID: "urn:ngsi-ld:product-offering:1f835fe9-2527-4579-9609-24858cdb4b8e",
		},
		{
			name:          "relatedParty filter by role Seller matching CSI Piemonte (VATIT-01995120019)",
			filters:       []string{"relatedParty[?(@.role == 'Seller')].name == 'VATIT-01995120019'"},
			wantCount:     2, // Nivola PostgreSQL & Nivola Oracle
			mustIncludeID: "urn:ngsi-ld:product-offering:73a949ac-a492-4de4-a106-35de37261cd2",
		},
		{
			name:          "relatedParty filter by role Seller matching Engineering (VATIT-12622480155)",
			filters:       []string{"relatedParty[?(@.role == 'Seller')].name == 'VATIT-12622480155'"},
			wantCount:     1, // Google Workspace
			mustIncludeID: "urn:ngsi-ld:product-offering:1f835fe9-2527-4579-9609-24858cdb4b8e",
		},
		{
			name:          "category filter by name Infrastructure matching specific ID",
			filters:       []string{"category[?(@.name == 'Infrastructure')].id == 'urn:ngsi-ld:category:86836789-af84-4ca0-98f6-ee923f22d98d'"},
			wantCount:     7,
			mustIncludeID: "urn:ngsi-ld:product-offering:cd390626-e57e-41e7-b575-7c6d9f71a2ef",
		},

		// 7. Pattern Matching with LIKE
		{
			name:          "name LIKE prefix Nivola PaaS%",
			filters:       []string{"name LIKE 'Nivola PaaS%'"},
			wantCount:     2, // PostgreSQL & Oracle
			mustIncludeID: "urn:ngsi-ld:product-offering:13a0a99b-c75b-4502-8352-804ba7f95475",
		},
		{
			name:          "name LIKE substring %IoT%",
			filters:       []string{"name LIKE '%IoT%'"},
			wantCount:     3, // Gas Emission, Radon, IoT Coverage
			mustIncludeID: "urn:ngsi-ld:product-offering:48c7f488-6aa2-4c8c-a2f9-5b8aba8fb616",
		},
		{
			name:          "description LIKE %GRAFANA%",
			filters:       []string{"description LIKE '%GRAFANA%'"},
			wantCount:     2, // Gas Emission & Radon
			mustIncludeID: "urn:ngsi-ld:product-offering:2825967d-0565-442e-a5d9-df9d3df2f5df",
		},

		// 8. Existence and Nullability
		{
			name:      "validFor.startDateTime IS NOT NULL",
			filters:   []string{"validFor.startDateTime IS NOT NULL"},
			wantCount: 10,
		},
		{
			name:      "validFor.startDateTime bare path shorthand",
			filters:   []string{"validFor.startDateTime"},
			wantCount: 10,
		},
		{
			name:      "validFor.endDateTime IS NULL",
			filters:   []string{"validFor.endDateTime IS NULL"},
			wantCount: 10, // None of the offerings define endDateTime
		},

		// 9. Boolean OR Combinations via Semicolon
		{
			name:          "OR via semicolon: Voting Platform or CloudFerro Cloud",
			filters:       []string{"name == 'Voting Platform';name == 'CloudFerro Cloud'"},
			wantCount:     2,
			mustIncludeID: "urn:ngsi-ld:product-offering:f4b1fd69-eb0e-41b2-a20a-d1047ba14ddb",
		},
		{
			name:          "OR via semicolon across categories: Databases or Agriculture",
			filters:       []string{"category[*].name == 'Databases';category[*].name == 'Agriculture'"},
			wantCount:     4,
			mustIncludeID: "urn:ngsi-ld:product-offering:48c7f488-6aa2-4c8c-a2f9-5b8aba8fb616",
		},

		// 10. Boolean AND Combinations via Multiple Parameters
		{
			name: "AND via multiple parameters: Launched AND Databases",
			filters: []string{
				"lifecycleStatus == 'Launched'",
				"category[*].name == 'Databases'",
			},
			wantCount:     2,
			mustIncludeID: "urn:ngsi-ld:product-offering:73a949ac-a492-4de4-a106-35de37261cd2",
		},
		{
			name: "AND via multiple parameters: Obsolete AND Infrastructure",
			filters: []string{
				"lifecycleStatus == 'Obsolete'",
				"category[*].name == 'Infrastructure'",
			},
			wantCount:     1,
			mustIncludeID: "urn:ngsi-ld:product-offering:1f835fe9-2527-4579-9609-24858cdb4b8e", // Google Workspace
		},
		{
			name: "AND returning 0 results: Obsolete AND Databases",
			filters: []string{
				"lifecycleStatus == 'Obsolete'",
				"category[*].name == 'Databases'",
			},
			wantCount: 0,
		},

		// 11. Conjunctive Normal Form (AND of OR groups)
		{
			name: "AND of OR groups: (Databases OR Professional) AND Launched",
			filters: []string{
				"category[*].name == 'Databases';category[*].name == 'Professional'",
				"lifecycleStatus == 'Launched'",
			},
			wantCount: 4, // 2 Databases (PostgreSQL, Oracle) + 2 Professional (Voting, User license)
		},
		{
			name: "AND of OR groups: (Infrastructure OR Public Sector) AND Seller VATIT-01995120019",
			filters: []string{
				"category[*].name == 'Infrastructure';category[*].name == 'Public Sector'",
				"relatedParty[?(@.role == 'Seller')].name == 'VATIT-01995120019'",
			},
			wantCount:     2, // Both Nivola offerings match
			mustIncludeID: "urn:ngsi-ld:product-offering:73a949ac-a492-4de4-a106-35de37261cd2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 1. Generate combined WHERE query and arguments
			whereSQL, args, err := GenerateCombinedFilterQuery("tmf_object", "content", tt.filters)
			if err != nil {
				t.Fatalf("GenerateCombinedFilterQuery failed: %v", err)
			}

			// 2. Build full SELECT query
			query := "SELECT id FROM tmf_object WHERE type = 'ProductOffering'"
			if whereSQL != "" {
				query += " AND " + whereSQL
			}

			// 3. Execute against in-memory SQLite database
			rows, err := db.Query(query, args...)
			if err != nil {
				t.Fatalf("query execution failed: %v\nSQL: %s\nArgs: %v", err, query, args)
			}
			defer rows.Close()

			var matchedIDs []string
			foundMustInclude := false
			foundMustExclude := false

			for rows.Next() {
				var id string
				if err := rows.Scan(&id); err != nil {
					t.Fatalf("rows.Scan failed: %v", err)
				}
				matchedIDs = append(matchedIDs, id)
				if tt.mustIncludeID != "" && id == tt.mustIncludeID {
					foundMustInclude = true
				}
				if tt.mustExcludeID != "" && id == tt.mustExcludeID {
					foundMustExclude = true
				}
			}
			if err := rows.Err(); err != nil {
				t.Fatalf("rows error: %v", err)
			}

			if len(matchedIDs) != tt.wantCount {
				t.Errorf("match count mismatch: got %d, want %d\nMatched IDs: %v\nSQL: %s\nArgs: %v",
					len(matchedIDs), tt.wantCount, matchedIDs, query, args)
			}

			if tt.mustIncludeID != "" && !foundMustInclude {
				t.Errorf("expected result to include ID %q, but was not found in: %v", tt.mustIncludeID, matchedIDs)
			}
			if tt.mustExcludeID != "" && foundMustExclude {
				t.Errorf("expected result to NOT include ID %q, but it was found", tt.mustExcludeID)
			}
		})
	}
}

// TestJSONPath_BuildSelectFromParms_EndToEnd verifies that BuildSelectFromParms correctly
// integrates with the in-memory SQLite database using real JSONB content.
func TestJSONPath_BuildSelectFromParms_EndToEnd(t *testing.T) {
	db := loadTestOfferingsDB(t)
	defer db.Close()

	// Query: ProductOffering with filter=category[*].name == 'Databases'
	v := make(map[string][]string)
	v["filter"] = []string{"category[*].name == 'Databases'"}

	sqlQuery, args, _, _, err := BuildSelectFromParms("ProductOffering", v)
	if err != nil {
		t.Fatalf("BuildSelectFromParms failed: %v", err)
	}

	rows, err := db.Query(sqlQuery, args...)
	if err != nil {
		t.Fatalf("query failed: %v\nSQL: %s\nArgs: %v", err, sqlQuery, args)
	}
	defer rows.Close()

	type resultRecord struct {
		id         string
		resType    string
		version    string
		apiVersion string
		seller     string
		buyer      string
		lastUpdate string
		jsonText   string
		createdAt  int64
		updatedAt  int64
		totalCount int
	}

	var results []resultRecord
	for rows.Next() {
		var r resultRecord
		err := rows.Scan(
			&r.id, &r.resType, &r.version, &r.apiVersion,
			&r.seller, &r.buyer, &r.lastUpdate, &r.jsonText,
			&r.createdAt, &r.updatedAt, &r.totalCount,
		)
		if err != nil {
			t.Fatalf("scan failed: %v", err)
		}
		results = append(results, r)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("rows error: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
	if results[0].totalCount != 2 {
		t.Errorf("expected totalCount 2, got %d", results[0].totalCount)
	}

	// Verify that json(content) extracted valid JSON containing the category
	for _, r := range results {
		if !strings.Contains(r.jsonText, "Databases") {
			t.Errorf("expected json content to contain 'Databases', got: %s", r.jsonText)
		}
	}
}

// TestJSONPath_SelectExpr_Projection verifies that generated SelectExpr expressions
// can be used to extract JSON arrays and fields from the real JSONB data.
func TestJSONPath_SelectExpr_Projection(t *testing.T) {
	db := loadTestOfferingsDB(t)
	defer db.Close()

	// 1. Array projection: extract all category names for Gas Emission Monitoring
	seg, err := MapJSONPathToSQL("tmf_object", "content", "category[*].name", "=")
	if err != nil {
		t.Fatalf("MapJSONPathToSQL failed: %v", err)
	}

	gasEmissionID := "urn:ngsi-ld:product-offering:cd390626-e57e-41e7-b575-7c6d9f71a2ef"
	query := "SELECT id, " + seg.SelectExpr + " FROM tmf_object WHERE id = ?"
	var id, categoriesJSON string
	err = db.QueryRow(query, gasEmissionID).Scan(&id, &categoriesJSON)
	if err != nil {
		t.Fatalf("query failed: %v\nQuery: %s", err, query)
	}

	var catNames []string
	if err := json.Unmarshal([]byte(categoriesJSON), &catNames); err != nil {
		t.Fatalf("failed to unmarshal extracted categories %q: %v", categoriesJSON, err)
	}

	expectedCategories := []string{"Infrastructure", "Networking and Content Delivery", "Data (DaaS)", "Software Services"}
	if len(catNames) != len(expectedCategories) {
		t.Fatalf("expected %d categories, got %d: %v", len(expectedCategories), len(catNames), catNames)
	}
	for i, expected := range expectedCategories {
		if catNames[i] != expected {
			t.Errorf("cat[%d]: got %q, want %q", i, catNames[i], expected)
		}
	}

	// 2. Predicate filter projection: extract the name of the Seller relatedParty
	segPred, err := MapJSONPathToSQL("tmf_object", "content", "relatedParty[?(@.role == 'Seller')].name", "=")
	if err != nil {
		t.Fatalf("MapJSONPathToSQL failed: %v", err)
	}

	queryPred := "SELECT id, " + segPred.SelectExpr + " FROM tmf_object WHERE id = ?"
	argsPred := append([]any{}, segPred.SelectArgs...)
	argsPred = append(argsPred, gasEmissionID)

	var predID, sellerNamesJSON string
	err = db.QueryRow(queryPred, argsPred...).Scan(&predID, &sellerNamesJSON)
	if err != nil {
		t.Fatalf("queryPred failed: %v\nQuery: %s\nArgs: %v", err, queryPred, argsPred)
	}

	var sellers []string
	if err := json.Unmarshal([]byte(sellerNamesJSON), &sellers); err != nil {
		t.Fatalf("failed to unmarshal extracted sellers %q: %v", sellerNamesJSON, err)
	}
	if len(sellers) != 1 || sellers[0] != "VATRO-1572582" {
		t.Errorf("expected ['VATRO-1572582'], got: %v", sellers)
	}
}
