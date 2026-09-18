package repository

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Valid SQL identifier regex (tables, columns)
var sqlIdentifierRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

// JSONPathQueryOptions defines the parameters for mapping a JSONPath expression
// to SQLite SQL query segments.
type JSONPathQueryOptions struct {
	// TableName is the name of the SQLite table (e.g. "tmf_object"). Defaults to "tmf_object".
	TableName string

	// ColumnName is the column containing the JSON/JSONB data (e.g. "content"). Defaults to "content".
	ColumnName string

	// Path is the JSONPath query expression (e.g. "category[*].id", "relatedParty[?(@.role == 'seller')].id").
	Path string

	// Operator is the SQL comparison operator for the WHERE clause (e.g. "=", "!=", "IN", "LIKE", ">", "<", ">=", "<=", "IS NOT NULL", "EXISTS").
	// Defaults to "=" (or "IN" if multiple values are provided).
	Operator string

	// Values are the target values to match in the WHERE clause.
	Values []any

	// RawJSON specifies whether SelectExpr should extract JSON objects/arrays directly using '->'
	// instead of text/scalar values using '->>'.
	RawJSON bool
}

// JSONPathSegments contains modular SQL expressions and arguments produced from a JSONPath query.
// Callers can integrate these segments into customized SELECT and WHERE statements alongside other columns.
type JSONPathSegments struct {
	// SelectExpr is a SQL expression suitable for use in the SELECT clause.
	// For simple paths:  "tmf_object.content ->> '$.name'"
	// For array paths:   "(SELECT json_group_array(json_each.value ->> '$.id') FROM json_each(tmf_object.content, '$.category'))"
	SelectExpr string

	// SelectArgs contains arguments bound to placeholders in SelectExpr (if any, e.g. from filter predicates).
	SelectArgs []any

	// WhereExpr is a boolean SQL expression suitable for use in a WHERE or AND clause.
	// For simple paths:  "tmf_object.content ->> '$.name' = ?"
	// For array paths:   "EXISTS (SELECT 1 FROM json_each(tmf_object.content, '$.category') WHERE json_each.value ->> '$.id' = ?)"
	WhereExpr string

	// WhereArgs contains arguments bound to placeholders in WhereExpr.
	WhereArgs []any

	// FromClause is an optional JOIN fragment if the caller chooses a JOIN over a subquery.
	FromClause string

	// Args is an alias for WhereArgs, for convenient binding in WHERE clauses.
	Args []any

	// IsArray indicates whether the path targets multiple elements (e.g. via array wildcard [*]).
	IsArray bool

	// IsPredicate indicates whether a filter predicate like [?(@.field == 'val')] was present.
	IsPredicate bool

	// NormalizedPath is the standard SQLite json path (e.g. "$.category[0].id").
	NormalizedPath string
}

// SegmentType represents the type of a parsed JSONPath segment.
type SegmentType int

const (
	SegProperty  SegmentType = iota // Property access: .field or ['field']
	SegIndex                        // Array index: [0]
	SegWildcard                     // Array wildcard: [*]
	SegFilter                       // Array filter predicate: [?(@.k == 'v')]
	SegRecursive                    // Recursive descent: ..field
)

// FilterPredicate represents a parsed array filter predicate: [?(@.property <op> <literal>)].
type FilterPredicate struct {
	Field string
	Op    string // "=", "!=", ">", "<", ">=", "<=", "LIKE"
	Value any
}

// ParsedSegment is an element in the parsed JSONPath.
type ParsedSegment struct {
	Type      SegmentType
	Key       string
	Index     int
	Predicate *FilterPredicate
}

// MapJSONPathToSQL is a convenience helper for GenerateJSONPathQuerySegments.
func MapJSONPathToSQL(tableName, columnName, path string, operator string, values ...any) (*JSONPathSegments, error) {
	return GenerateJSONPathQuerySegments(JSONPathQueryOptions{
		TableName:  tableName,
		ColumnName: columnName,
		Path:       path,
		Operator:   operator,
		Values:     values,
	})
}

// GenerateJSONPathQuerySegments maps a JSONPath query to SQLite SQL segments for SELECT and WHERE.
func GenerateJSONPathQuerySegments(opts JSONPathQueryOptions) (*JSONPathSegments, error) {
	tableName := opts.TableName
	if tableName == "" {
		tableName = "tmf_object"
	}
	columnName := opts.ColumnName
	if columnName == "" {
		columnName = "content"
	}

	if !sqlIdentifierRegex.MatchString(tableName) {
		return nil, fmt.Errorf("invalid table name: %q", tableName)
	}
	if !sqlIdentifierRegex.MatchString(columnName) {
		return nil, fmt.Errorf("invalid column name: %q", columnName)
	}

	rawPath := strings.TrimSpace(opts.Path)
	if rawPath == "" {
		return nil, fmt.Errorf("empty JSONPath")
	}

	segments, err := parseJSONPath(rawPath)
	if err != nil {
		return nil, fmt.Errorf("invalid JSONPath %q: %w", rawPath, err)
	}
	if len(segments) == 0 {
		return nil, fmt.Errorf("empty JSONPath after parsing: %q", rawPath)
	}

	operator := strings.ToUpper(strings.TrimSpace(opts.Operator))
	switch operator {
	case "":
		if len(opts.Values) > 1 {
			operator = "IN"
		} else {
			operator = "="
		}
	case "==":
		operator = "="
	}

	// Count wildcards and filter predicates to select execution strategy
	wildcardCount := 0
	filterCount := 0
	hasRecursive := false
	for _, seg := range segments {
		switch seg.Type {
		case SegWildcard:
			wildcardCount++
		case SegFilter:
			filterCount++
		case SegRecursive:
			hasRecursive = true
		}
	}

	tableCol := tableName + "." + columnName

	// Strategy 1: Simple scalar/object path (no wildcards, filters, or recursive descent)
	if wildcardCount == 0 && filterCount == 0 && !hasRecursive {
		return buildSimplePathSegments(tableCol, segments, operator, opts.Values, opts.RawJSON)
	}

	// Strategy 2: Single array wildcard or single filter predicate (uses json_each)
	if (wildcardCount == 1 && filterCount == 0 && !hasRecursive) ||
		(wildcardCount == 0 && filterCount == 1 && !hasRecursive) {
		return buildJsonEachSegments(tableCol, segments, operator, opts.Values, opts.RawJSON)
	}

	// Strategy 3: Multi-level wildcards or recursive descent (uses json_tree)
	return buildJsonTreeSegments(tableCol, segments, operator, opts.Values, opts.RawJSON)
}

// buildSimplePathSegments generates SQL for paths without array iteration (e.g. $.name, $.place[0].id).
func buildSimplePathSegments(tableCol string, segments []ParsedSegment, operator string, values []any, rawJSON bool) (*JSONPathSegments, error) {
	sqlitePath := buildSQLitePath(segments)
	arrow := "->>"
	if rawJSON {
		arrow = "->"
	}

	selectExpr := fmt.Sprintf("%s %s '%s'", tableCol, arrow, sqlitePath)
	var whereExpr string
	var args []any

	switch operator {
	case "EXISTS", "IS NOT NULL":
		whereExpr = fmt.Sprintf("%s ->> '%s' IS NOT NULL", tableCol, sqlitePath)
	case "IS NULL":
		whereExpr = fmt.Sprintf("%s ->> '%s' IS NULL", tableCol, sqlitePath)
	case "IN":
		if len(values) == 0 {
			whereExpr = "0 = 1"
		} else {
			placeholders := make([]string, len(values))
			for i := range values {
				placeholders[i] = "?"
				args = append(args, values[i])
			}
			whereExpr = fmt.Sprintf("%s ->> '%s' IN (%s)", tableCol, sqlitePath, strings.Join(placeholders, ", "))
		}
	default:
		if len(values) == 0 {
			whereExpr = fmt.Sprintf("%s ->> '%s' IS NOT NULL", tableCol, sqlitePath)
		} else {
			whereExpr = fmt.Sprintf("%s ->> '%s' %s ?", tableCol, sqlitePath, operator)
			args = append(args, values[0])
		}
	}

	return &JSONPathSegments{
		SelectExpr:     selectExpr,
		SelectArgs:     nil,
		WhereExpr:      whereExpr,
		WhereArgs:      args,
		Args:           args,
		IsArray:        false,
		IsPredicate:    false,
		NormalizedPath: sqlitePath,
	}, nil
}

// buildJsonEachSegments generates SQL using json_each for single-level arrays.
func buildJsonEachSegments(tableCol string, segments []ParsedSegment, operator string, values []any, rawJSON bool) (*JSONPathSegments, error) {
	var preSegments []ParsedSegment
	var pivotSegment ParsedSegment
	var postSegments []ParsedSegment

	foundPivot := false
	for _, seg := range segments {
		if !foundPivot && (seg.Type == SegWildcard || seg.Type == SegFilter) {
			pivotSegment = seg
			foundPivot = true
			continue
		}
		if !foundPivot {
			preSegments = append(preSegments, seg)
		} else {
			postSegments = append(postSegments, seg)
		}
	}

	arrayPath := buildSQLitePath(preSegments)
	elemPath := buildSQLitePath(postSegments)
	// For elements inside json_each: if elemPath is "$", we extract json_each.value directly;
	// otherwise, json_each.value ->> '$.subfield'
	elemExpr := "json_each.value"
	if elemPath != "$" {
		arrow := "->>"
		if rawJSON {
			arrow = "->"
		}
		elemExpr = fmt.Sprintf("json_each.value %s '%s'", arrow, elemPath)
	}

	elemWhereExpr := "json_each.value"
	if elemPath != "$" {
		elemWhereExpr = fmt.Sprintf("json_each.value ->> '%s'", elemPath)
	}

	var selectConds []string
	var selectArgs []any
	var whereConds []string
	var whereArgs []any

	if pivotSegment.Type == SegFilter && pivotSegment.Predicate != nil {
		pred := pivotSegment.Predicate
		filterFieldPath := "$." + strings.TrimPrefix(pred.Field, "$.")
		predExpr := fmt.Sprintf("json_each.value ->> '%s' %s ?", filterFieldPath, pred.Op)
		selectConds = append(selectConds, predExpr)
		selectArgs = append(selectArgs, pred.Value)
		whereConds = append(whereConds, predExpr)
		whereArgs = append(whereArgs, pred.Value)
	}

	// Build SELECT expression: (SELECT json_group_array(...) FROM json_each(...) WHERE ...)
	var selectWhereClause string
	if len(selectConds) > 0 {
		selectWhereClause = " WHERE " + strings.Join(selectConds, " AND ")
	}
	selectExpr := fmt.Sprintf("(SELECT json_group_array(%s) FROM json_each(%s, '%s')%s)", elemExpr, tableCol, arrayPath, selectWhereClause)

	// Build WHERE condition
	switch operator {
	case "EXISTS", "IS NOT NULL":
		if elemPath != "$" {
			whereConds = append(whereConds, fmt.Sprintf("%s IS NOT NULL", elemWhereExpr))
		}
	case "IS NULL":
		whereConds = append(whereConds, fmt.Sprintf("%s IS NULL", elemWhereExpr))
	case "IN":
		if len(values) == 0 {
			whereConds = append(whereConds, "0 = 1")
		} else {
			placeholders := make([]string, len(values))
			for i := range values {
				placeholders[i] = "?"
				whereArgs = append(whereArgs, values[i])
			}
			whereConds = append(whereConds, fmt.Sprintf("%s IN (%s)", elemWhereExpr, strings.Join(placeholders, ", ")))
		}
	default:
		if len(values) > 0 {
			whereConds = append(whereConds, fmt.Sprintf("%s %s ?", elemWhereExpr, operator))
			whereArgs = append(whereArgs, values[0])
		}
	}

	whereClause := ""
	if len(whereConds) > 0 {
		whereClause = " WHERE " + strings.Join(whereConds, " AND ")
	}
	whereExpr := fmt.Sprintf("EXISTS (SELECT 1 FROM json_each(%s, '%s')%s)", tableCol, arrayPath, whereClause)

	// Optional FromClause for callers wanting direct JOINs
	fromClause := fmt.Sprintf("LEFT JOIN json_each(%s, '%s')", tableCol, arrayPath)

	return &JSONPathSegments{
		SelectExpr:     selectExpr,
		SelectArgs:     selectArgs,
		WhereExpr:      whereExpr,
		WhereArgs:      whereArgs,
		FromClause:     fromClause,
		Args:           whereArgs,
		IsArray:        true,
		IsPredicate:    pivotSegment.Type == SegFilter,
		NormalizedPath: arrayPath + "[*]" + strings.TrimPrefix(elemPath, "$"),
	}, nil
}

// buildJsonTreeSegments handles multi-level wildcards or recursive descent using json_tree.
func buildJsonTreeSegments(tableCol string, segments []ParsedSegment, operator string, values []any, rawJSON bool) (*JSONPathSegments, error) {
	likePattern := buildLikePattern(segments)

	valueExpr := "json_tree.value"
	if rawJSON {
		valueExpr = "json_tree.atom"
	}

	selectExpr := fmt.Sprintf("(SELECT json_group_array(%s) FROM json_tree(%s) WHERE json_tree.fullkey LIKE '%s')", valueExpr, tableCol, likePattern)

	var whereConds []string
	var args []any
	whereConds = append(whereConds, fmt.Sprintf("json_tree.fullkey LIKE '%s'", likePattern))

	switch operator {
	case "EXISTS", "IS NOT NULL":
		whereConds = append(whereConds, "json_tree.value IS NOT NULL")
	case "IS NULL":
		whereConds = append(whereConds, "json_tree.value IS NULL")
	case "IN":
		if len(values) == 0 {
			whereConds = append(whereConds, "0 = 1")
		} else {
			placeholders := make([]string, len(values))
			for i := range values {
				placeholders[i] = "?"
				args = append(args, values[i])
			}
			whereConds = append(whereConds, fmt.Sprintf("json_tree.value IN (%s)", strings.Join(placeholders, ", ")))
		}
	default:
		if len(values) > 0 {
			whereConds = append(whereConds, fmt.Sprintf("json_tree.value %s ?", operator))
			args = append(args, values[0])
		}
	}

	whereExpr := fmt.Sprintf("EXISTS (SELECT 1 FROM json_tree(%s) WHERE %s)", tableCol, strings.Join(whereConds, " AND "))

	return &JSONPathSegments{
		SelectExpr:     selectExpr,
		SelectArgs:     nil,
		WhereExpr:      whereExpr,
		WhereArgs:      args,
		Args:           args,
		IsArray:        true,
		IsPredicate:    false,
		NormalizedPath: likePattern,
	}, nil
}

// buildSQLitePath converts a slice of parsed segments into a SQLite json path string (e.g. $.a.b[0].c).
func buildSQLitePath(segments []ParsedSegment) string {
	if len(segments) == 0 {
		return "$"
	}
	var sb strings.Builder
	sb.WriteString("$")
	for _, seg := range segments {
		switch seg.Type {
		case SegProperty:
			sb.WriteString(".")
			sb.WriteString(seg.Key)
		case SegIndex:
			sb.WriteString("[")
			sb.WriteString(strconv.Itoa(seg.Index))
			sb.WriteString("]")
		case SegWildcard:
			sb.WriteString("[*]")
		case SegRecursive:
			sb.WriteString("..")
			sb.WriteString(seg.Key)
		}
	}
	return sb.String()
}

// buildLikePattern converts parsed segments into a SQL LIKE pattern matching fullkey in json_tree.
func buildLikePattern(segments []ParsedSegment) string {
	var sb strings.Builder
	sb.WriteString("$")
	for _, seg := range segments {
		switch seg.Type {
		case SegProperty:
			sb.WriteString(".")
			sb.WriteString(seg.Key)
		case SegIndex:
			sb.WriteString("[")
			sb.WriteString(strconv.Itoa(seg.Index))
			sb.WriteString("]")
		case SegWildcard:
			sb.WriteString("[%]")
		case SegFilter:
			sb.WriteString("[%]")
		case SegRecursive:
			sb.WriteString("%.")
			sb.WriteString(seg.Key)
		}
	}
	return sb.String()
}

// parseJSONPath parses a JSONPath string into a sequence of ParsedSegments.
func parseJSONPath(input string) ([]ParsedSegment, error) {
	s := strings.TrimSpace(input)
	if after, ok := strings.CutPrefix(s, "$"); ok {
		s = after
		if strings.HasPrefix(s, ".") && !strings.HasPrefix(s, "..") {
			s = s[1:]
		}
	}

	var segments []ParsedSegment
	i := 0
	n := len(s)

	for i < n {
		c := s[i]

		// Skip dots
		if c == '.' {
			if i+1 < n && s[i+1] == '.' {
				// Recursive descent: ..key
				i += 2
				keyStart := i
				for i < n && (isIdentRune(rune(s[i])) || s[i] == '_') {
					i++
				}
				key := s[keyStart:i]
				if key == "" {
					return nil, fmt.Errorf("expected property name after '..' at pos %d", keyStart)
				}
				segments = append(segments, ParsedSegment{Type: SegRecursive, Key: key})
				continue
			}
			i++
			continue
		}

		// Bracket expression: [0], [*], ['prop'], or [?(...)]
		if c == '[' {
			closeIdx := findMatchingBracket(s, i)
			if closeIdx == -1 {
				return nil, fmt.Errorf("unclosed bracket at pos %d", i)
			}
			inside := strings.TrimSpace(s[i+1 : closeIdx])
			i = closeIdx + 1

			if inside == "*" {
				segments = append(segments, ParsedSegment{Type: SegWildcard})
			} else if strings.HasPrefix(inside, "?") {
				// Filter expression: ?(...)
				filterPred, err := parseFilterPredicate(inside)
				if err != nil {
					return nil, err
				}
				segments = append(segments, ParsedSegment{Type: SegFilter, Predicate: filterPred})
			} else if (strings.HasPrefix(inside, "'") && strings.HasSuffix(inside, "'")) ||
				(strings.HasPrefix(inside, "\"") && strings.HasSuffix(inside, "\"")) {
				// Quoted property name: ['foo']
				key := inside[1 : len(inside)-1]
				segments = append(segments, ParsedSegment{Type: SegProperty, Key: key})
			} else if idx, err := strconv.Atoi(inside); err == nil {
				// Array index: [0]
				segments = append(segments, ParsedSegment{Type: SegIndex, Index: idx})
			} else {
				return nil, fmt.Errorf("unsupported bracket expression: [%s]", inside)
			}
			continue
		}

		// Property identifier: foo
		start := i
		for i < n && s[i] != '.' && s[i] != '[' {
			i++
		}
		prop := strings.TrimSpace(s[start:i])
		if prop != "" {
			segments = append(segments, ParsedSegment{Type: SegProperty, Key: prop})
		}
	}

	return segments, nil
}

func findMatchingBracket(s string, startIdx int) int {
	depth := 0
	inQuote := rune(0)
	for i := startIdx; i < len(s); i++ {
		r := rune(s[i])
		if inQuote != 0 {
			if r == inQuote && (i == 0 || s[i-1] != '\\') {
				inQuote = 0
			}
			continue
		}
		if r == '\'' || r == '"' {
			inQuote = r
			continue
		}
		switch r {
		case '[':
			depth++
		case ']':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}

// parseFilterPredicate parses the contents inside ?(...) into a FilterPredicate.
// Examples:
// ?(@.role == 'seller')
// ?(@.status != "closed")
// ?(@.price > 100)
func parseFilterPredicate(inside string) (*FilterPredicate, error) {
	expr := strings.TrimSpace(strings.TrimPrefix(inside, "?"))
	if strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") {
		expr = strings.TrimSpace(expr[1 : len(expr)-1])
	}

	if !strings.HasPrefix(expr, "@") {
		return nil, fmt.Errorf("filter predicate must start with '@': %q", expr)
	}

	expr = strings.TrimPrefix(expr, "@")
	expr = strings.TrimPrefix(expr, ".")

	opIndex, opLen, foundOp := findOuterOperator(expr)
	if opIndex == -1 {
		return nil, fmt.Errorf("no supported comparison operator found in filter: %q", expr)
	}

	field := strings.TrimSpace(expr[:opIndex])
	valStr := strings.TrimSpace(expr[opIndex+opLen:])

	if field == "" {
		return nil, fmt.Errorf("missing field in filter: %q", expr)
	}

	val, err := parseLiteralValue(valStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse literal %q in filter: %w", valStr, err)
	}

	return &FilterPredicate{
		Field: field,
		Op:    foundOp,
		Value: val,
	}, nil
}

// parseLiteralValue parses literal values (string, number, bool, null) from predicate expressions.
func parseLiteralValue(s string) (any, error) {
	s = strings.TrimSpace(s)
	if (strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'")) ||
		(strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"")) {
		return s[1 : len(s)-1], nil
	}
	if s == "true" {
		return true, nil
	}
	if s == "false" {
		return false, nil
	}
	if s == "null" {
		return nil, nil
	}
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i, nil
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f, nil
	}
	// Fallback to string literal
	return s, nil
}

func isIdentRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// ParseFilterExpression parses a complete filter expression string, e.g.:
// "category[*].id == 'cat-123'"
// "name LIKE 'Fiber%'"
// "status IN ('Active', 'Launched')"
// "validFor.endDateTime IS NOT NULL"
// "lifecycleStatus"
// into its target JSONPath, comparison operator, and comparison value(s).
func ParseFilterExpression(raw string) (path string, op string, values []any, err error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", "", nil, fmt.Errorf("empty filter expression")
	}

	opIndex, opLen, foundOp := findOuterOperator(raw)
	if opIndex == -1 {
		// No outer operator: treat as bare path with "IS NOT NULL"
		return raw, "IS NOT NULL", nil, nil
	}

	path = strings.TrimSpace(raw[:opIndex])
	if path == "" {
		return "", "", nil, fmt.Errorf("missing JSONPath before operator in: %q", raw)
	}

	valStr := strings.TrimSpace(raw[opIndex+opLen:])
	op = strings.ToUpper(foundOp)
	switch op {
	case "==":
		op = "="
	case "<>":
		op = "!="
	}

	switch op {
	case "IS NOT NULL", "IS NULL", "EXISTS":
		return path, op, nil, nil
	case "IN":
		vals, err := parseInClauseValues(valStr)
		if err != nil {
			return "", "", nil, fmt.Errorf("failed to parse IN values in %q: %w", raw, err)
		}
		return path, op, vals, nil
	default:
		if valStr == "" {
			return "", "", nil, fmt.Errorf("missing value after operator %q in: %q", foundOp, raw)
		}
		val, err := parseLiteralValue(valStr)
		if err != nil {
			return "", "", nil, fmt.Errorf("failed to parse value %q in %q: %w", valStr, raw, err)
		}
		return path, op, []any{val}, nil
	}
}

// GenerateCombinedFilterQuery processes filter parameters where:
// - Each element in filterParams (from separate &filter= parameters) is combined with AND.
// - Within each element, expressions separated by semicolons ';' are combined with OR.
// Example: ["a;b", "c"] produces "((a OR b) AND (c))".
func GenerateCombinedFilterQuery(tableName, columnName string, filterParams []string) (string, []any, error) {
	if len(filterParams) == 0 {
		return "", nil, nil
	}

	var andClauses []string
	var allArgs []any

	for _, param := range filterParams {
		trimmedParam := strings.TrimSpace(param)
		if trimmedParam == "" {
			continue
		}

		orParts := splitFilterExpressions(trimmedParam)
		if len(orParts) == 0 {
			continue
		}

		var orClauses []string
		var orArgs []any

		for _, expr := range orParts {
			path, op, vals, err := ParseFilterExpression(expr)
			if err != nil {
				return "", nil, err
			}

			seg, err := GenerateJSONPathQuerySegments(JSONPathQueryOptions{
				TableName:  tableName,
				ColumnName: columnName,
				Path:       path,
				Operator:   op,
				Values:     vals,
			})
			if err != nil {
				return "", nil, err
			}

			orClauses = append(orClauses, seg.WhereExpr)
			orArgs = append(orArgs, seg.WhereArgs...)
		}

		if len(orClauses) == 1 {
			andClauses = append(andClauses, orClauses[0])
			allArgs = append(allArgs, orArgs...)
		} else if len(orClauses) > 1 {
			andClauses = append(andClauses, "("+strings.Join(orClauses, " OR ")+")")
			allArgs = append(allArgs, orArgs...)
		}
	}

	if len(andClauses) == 0 {
		return "", nil, nil
	}

	if len(andClauses) == 1 {
		return andClauses[0], allArgs, nil
	}

	return "(" + strings.Join(andClauses, " AND ") + ")", allArgs, nil
}

// splitFilterExpressions splits a filter string by semicolons ';' that are not inside quotes or brackets.
func splitFilterExpressions(input string) []string {
	var parts []string
	depth := 0
	inQuote := rune(0)
	start := 0
	n := len(input)

	for i := 0; i < n; i++ {
		r := rune(input[i])
		if inQuote != 0 {
			if r == inQuote && (i == 0 || input[i-1] != '\\') {
				inQuote = 0
			}
			continue
		}
		if r == '\'' || r == '"' {
			inQuote = r
			continue
		}
		if r == '[' || r == '(' {
			depth++
			continue
		}
		if r == ']' || r == ')' {
			if depth > 0 {
				depth--
			}
			continue
		}

		if depth == 0 && r == ';' {
			part := strings.TrimSpace(input[start:i])
			if part != "" {
				parts = append(parts, part)
			}
			start = i + 1
		}
	}

	if start < n {
		part := strings.TrimSpace(input[start:])
		if part != "" {
			parts = append(parts, part)
		}
	}

	return parts
}

// findOuterOperator locates the outermost comparison operator in an expression.
func findOuterOperator(s string) (int, int, string) {
	depth := 0
	parenDepth := 0
	inQuote := rune(0)
	n := len(s)

	type opPattern struct {
		op     string
		isWord bool
		canon  string
	}

	candidates := []opPattern{
		{op: "IS NOT NULL", isWord: true, canon: "IS NOT NULL"},
		{op: "IS NULL", isWord: true, canon: "IS NULL"},
		{op: "LIKE", isWord: true, canon: "LIKE"},
		{op: "IN", isWord: true, canon: "IN"},
		{op: "EXISTS", isWord: true, canon: "EXISTS"},
		{op: "==", isWord: false, canon: "="},
		{op: "!=", isWord: false, canon: "!="},
		{op: "<>", isWord: false, canon: "!="},
		{op: ">=", isWord: false, canon: ">="},
		{op: "<=", isWord: false, canon: "<="},
		{op: "=", isWord: false, canon: "="},
		{op: ">", isWord: false, canon: ">"},
		{op: "<", isWord: false, canon: "<"},
	}

	for i := range n {
		r := rune(s[i])
		if inQuote != 0 {
			if r == inQuote && (i == 0 || s[i-1] != '\\') {
				inQuote = 0
			}
			continue
		}
		if r == '\'' || r == '"' {
			inQuote = r
			continue
		}
		if r == '[' {
			depth++
			continue
		}
		if r == ']' {
			if depth > 0 {
				depth--
			}
			continue
		}
		if r == '(' {
			parenDepth++
			continue
		}
		if r == ')' {
			if parenDepth > 0 {
				parenDepth--
			}
			continue
		}

		if depth == 0 && parenDepth == 0 {
			for _, cand := range candidates {
				if i+len(cand.op) <= n {
					sub := s[i : i+len(cand.op)]
					if strings.EqualFold(sub, cand.op) {
						if cand.isWord {
							leftOk := i == 0 || unicode.IsSpace(rune(s[i-1])) || s[i-1] == ']' || s[i-1] == ')'
							rightIdx := i + len(cand.op)
							rightOk := rightIdx == n || unicode.IsSpace(rune(s[rightIdx])) || s[rightIdx] == '('
							if leftOk && rightOk {
								return i, len(cand.op), cand.canon
							}
						} else {
							return i, len(cand.op), cand.canon
						}
					}
				}
			}
		}
	}
	return -1, 0, ""
}

// parseInClauseValues parses a parenthesized comma-separated list of literals for an IN clause: ('a', 'b', 123).
func parseInClauseValues(s string) ([]any, error) {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "(") || !strings.HasSuffix(s, ")") {
		return nil, fmt.Errorf("IN clause values must be enclosed in parentheses: %q", s)
	}
	inside := strings.TrimSpace(s[1 : len(s)-1])
	if inside == "" {
		return nil, nil
	}

	// Split by comma outside quotes
	var items []string
	inQuote := rune(0)
	start := 0
	n := len(inside)

	for i := range n {
		r := rune(inside[i])
		if inQuote != 0 {
			if r == inQuote && (i == 0 || inside[i-1] != '\\') {
				inQuote = 0
			}
			continue
		}
		if r == '\'' || r == '"' {
			inQuote = r
			continue
		}
		if r == ',' {
			item := strings.TrimSpace(inside[start:i])
			if item != "" {
				items = append(items, item)
			}
			start = i + 1
		}
	}
	if start < n {
		item := strings.TrimSpace(inside[start:])
		if item != "" {
			items = append(items, item)
		}
	}

	var values []any
	for _, item := range items {
		val, err := parseLiteralValue(item)
		if err != nil {
			return nil, err
		}
		values = append(values, val)
	}
	return values, nil
}
