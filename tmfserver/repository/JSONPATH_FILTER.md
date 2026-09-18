# TM Forum REST API: JSONPath `filter` Query Parameter Guide

This document describes how to use the `filter` URL query parameter to perform rich JSONPath-based filtering on TM Forum (TMF) Open API resources.

---

## 1. Overview

In addition to standard attribute equality filters (e.g. `?lifecycleStatus=Launched`), the TMF API supports the `filter` query parameter. This allows clients to query deeply nested properties, search across arrays, filter by sub-document attributes, and combine conditions with Boolean **AND** and **OR** logic:

```http
GET /tmf-api/{apiFamily}/{apiVersion}/{resource}?filter={jsonpath_expression}
```

### URL Encoding Requirement
Because JSONPath expressions contain characters with special meaning in URLs (such as `[`, `]`, `?`, `@`, `'`, `"`, `=`, `;`, and spaces), the entire expression **must be URL-encoded** when sent in an HTTP request.

| Format | Example |
| :--- | :--- |
| **Raw Expression** | `relatedParty[?(@.role == 'seller')].id == 'did:elsi:VATES-12345678'` |
| **URL-Encoded** | `?filter=relatedParty%5B%3F(%40.role%20%3D%3D%20'seller')%5D.id%20%3D%3D%20'did%3Aelsi%3AVATES-12345678'` |

---

## 2. Boolean Logic & Combining Expressions (AND & OR)

You can combine multiple JSONPath expressions to create expressive Boolean queries using semicolons (`;`) and repeated query parameters (`&filter=`).

### 2.1 OR Logic: Semicolon (`;`)
Separate multiple expressions with semicolons within a single `filter` parameter to **OR** them:

```http
?filter=expression1;expression2;...;expressionN
```

The database query matches rows satisfying **any** of the expressions:
```sql
(expression1 OR expression2 OR ... OR expressionN)
```

**Example:** Match offerings that are either `Launched` OR `Active`:
```http
GET /tmf-api/productCatalogManagement/v4/productOffering?filter=lifecycleStatus%20%3D%3D%20'Launched'%3BlifecycleStatus%20%3D%3D%20'Active'
```
*Unencoded:* `?filter=lifecycleStatus == 'Launched';lifecycleStatus == 'Active'`

> **String Literal Safety**: Semicolons inside quoted string literals (e.g. `'Broadband; High-Speed'`) are treated as literal text and will **not** be treated as OR separators.

---

### 2.2 AND Logic: Repeated Parameters (`&filter=`)
Repeat the `filter` query parameter to **AND** multiple expressions:

```http
?filter=expression1&filter=expression2
```

The database query matches rows satisfying **all** of the expressions:
```sql
(expression1) AND (expression2)
```

**Example:** Match offerings with status `Launched` **AND** category ID `cat-fiber`:
```http
GET /tmf-api/productCatalogManagement/v4/productOffering?filter=lifecycleStatus%20%3D%3D%20'Launched'&filter=category%5B*%5D.id%20%3D%3D%20'cat-fiber'
```

---

### 2.3 Mixed AND & OR (Conjunctive Normal Form)
Each `filter` parameter forms an **OR group**, and all `filter` parameters are joined by **AND**:

```http
?filter=A;B&filter=C;D
```

This maps directly to:
```sql
((A OR B) AND (C OR D))
```

**Example:** Match offerings that are (`Launched` OR `Active`) **AND** belong to (`cat-fiber` OR `cat-wireless`):
```http
GET /tmf-api/productCatalogManagement/v4/productOffering?filter=lifecycleStatus%20%3D%3D%20'Launched'%3BlifecycleStatus%20%3D%3D%20'Active'&filter=category%5B*%5D.id%20%3D%3D%20'cat-fiber'%3Bcategory%5B*%5D.id%20%3D%3D%20'cat-wireless'
```

---

## 3. Supported JSONPath Subset

The API supports a high-performance subset of JSONPath optimized for SQLite JSONB storage.

### 3.1 Object Property Navigation
Access top-level and nested object properties using standard dot-notation:

| Syntax | Description | Example |
| :--- | :--- | :--- |
| `name` or `$.name` | Top-level property | `name == 'Broadband 500Mbps'` |
| `place.address.city` | Nested object property | `place.address.city == 'Madrid'` |
| `$['lifecycleStatus']` | Bracket notation for property | `$['lifecycleStatus'] == 'Launched'` |

> **Note**: The leading `$` or `$.` root identifier is optional. Both `name` and `$.name` are supported.

---

### 3.2 Array Indexing
Access a specific element in an array by its zero-based index:

| Syntax | Description | Example |
| :--- | :--- | :--- |
| `attachment[0].url` | First element of array | `attachment[0].url == 'https://example.com/spec.pdf'` |
| `organizationIdentification[0].identificationId` | First identification entry | `organizationIdentification[0].identificationId == 'did:elsi:123'` |

---

### 3.3 Array Wildcard (`[*]`)
Matches **any** element in an array. If any element satisfies the condition, the resource is returned:

| Syntax | Description | Example |
| :--- | :--- | :--- |
| `category[*].id` | Matches if any category in the array has this ID | `category[*].id == 'urn:ngsi-ld:category:broadband'` |
| `relatedParty[*].role` | Matches if any related party has this role | `relatedParty[*].role == 'seller'` |

---

### 3.4 Array Filter Predicates (`[?(@.field <op> <value>)]`)
Filter elements in an array based on an attribute within each element before checking the target property:

```
{array}[?(@.{attribute} {operator} {literal})].{targetProperty} {operator} {value}
```

#### Examples:
- **Filter by related party role:**
  ```text
  relatedParty[?(@.role == 'seller')].id == 'did:elsi:VATES-12345678'
  ```
- **Filter by characteristic name and numeric value:**
  ```text
  productCharacteristic[?(@.name == 'downloadSpeed')].value >= 500
  ```
- **Filter by attachment type using pattern matching:**
  ```text
  attachment[?(@.attachmentType == 'Picture')].name LIKE '%.png'
  ```

---

### 3.5 Multi-Level Wildcards & Recursive Descent
For complex schemas with nested arrays or arbitrary depths:

| Syntax | Description | Example |
| :--- | :--- | :--- |
| `items[*].components[*].id` | Multi-level array traversal | `items[*].components[*].id == 'part-99'` |
| `..identificationId` | Recursive descent (any node matching key) | `..identificationId == 'VATES-12345678'` |

---

## 4. Supported Operators & Literals

### 4.1 Comparison Operators

| Operator | Description | Example |
| :--- | :--- | :--- |
| `==` or `=` | Equal to | `lifecycleStatus == 'Launched'` |
| `!=` or `<>` | Not equal to | `lifecycleStatus != 'Retired'` |
| `>` | Greater than | `price.amount > 100` |
| `>=` | Greater than or equal to | `productCharacteristic[?(@.name == 'bandwidth')].value >= 1000` |
| `<` | Less than | `price.amount < 50` |
| `<=` | Less than or equal to | `price.amount <= 25.5` |
| `LIKE` | SQL pattern match (`%` = any characters, `_` = one character) | `name LIKE 'Fiber%'` |
| `IN` | Set inclusion with comma-separated values in parentheses | `lifecycleStatus IN ('Launched', 'Active')` |
| `IS NOT NULL` | Property existence check | `validFor.endDateTime IS NOT NULL` |
| `IS NULL` | Property absence / null check | `validFor.endDateTime IS NULL` |

> **Shorthand Existence Check**: A bare path without an operator (e.g. `?filter=validFor.endDateTime`) is interpreted as `validFor.endDateTime IS NOT NULL`.

### 4.2 Supported Literal Types
- **Strings**: Single (`'...'`) or double (`"..."`) quotes (e.g. `'seller'`, `"Active"`).
- **Integers**: Unquoted integer numbers (e.g. `100`, `-5`).
- **Floats**: Unquoted decimal numbers (e.g. `19.99`, `0.5`).
- **Booleans**: Unquoted `true` or `false`.
- **Null**: Unquoted `null`.

---

## 5. Complete Request Examples

### 5.1 Query by Category ID (Array Wildcard)
Find offerings belonging to the broadband category:
```http
GET /tmf-api/productCatalogManagement/v4/productOffering?filter=category%5B*%5D.id%20%3D%3D%20'urn:ngsi-ld:category:broadband'
```
*Unencoded:* `category[*].id == 'urn:ngsi-ld:category:broadband'`

---

### 5.2 Query by Specific Seller Identity (Array Filter Predicate)
Find offerings where `relatedParty` has an entry with `role == 'seller'` and ID `did:elsi:VATES-B12345678`:
```http
GET /tmf-api/productCatalogManagement/v4/productOffering?filter=relatedParty%5B%3F(%40.role%20%3D%3D%20'seller')%5D.id%20%3D%3D%20'did%3Aelsi%3AVATES-B12345678'
```
*Unencoded:* `relatedParty[?(@.role == 'seller')].id == 'did:elsi:VATES-B12345678'`

---

### 5.3 Query with Multiple Categories (Set Inclusion `IN`)
Find offerings in any of several categories:
```http
GET /tmf-api/productCatalogManagement/v4/productOffering?filter=category%5B*%5D.id%20IN%20('cat-fiber'%2C%20'cat-broadband')
```
*Unencoded:* `category[*].id IN ('cat-fiber', 'cat-broadband')`

---

### 5.4 OR Combination via Semicolon
Find offerings that are either `Launched` OR `Active`:
```http
GET /tmf-api/productCatalogManagement/v4/productOffering?filter=lifecycleStatus%20%3D%3D%20'Launched'%3BlifecycleStatus%20%3D%3D%20'Active'
```
*Unencoded:* `lifecycleStatus == 'Launched';lifecycleStatus == 'Active'`

---

### 5.5 AND Combination via Multiple Parameters
Find offerings that are `Launched` AND belong to category `cat-fiber`:
```http
GET /tmf-api/productCatalogManagement/v4/productOffering?filter=lifecycleStatus%20%3D%3D%20'Launched'&filter=category%5B*%5D.id%20%3D%3D%20'cat-fiber'
```

---

### 5.6 Name Prefix Search (LIKE)
Find all offerings whose name starts with `"Fiber"`:
```http
GET /tmf-api/productCatalogManagement/v4/productOffering?filter=name%20LIKE%20'Fiber%25'
```
*Unencoded:* `name LIKE 'Fiber%'` *(Note: `%` is encoded as `%25`)*

---

### 5.7 Query Organization by DID Identity
Find organizations where `organizationIdentification` has `identificationType == 'did:elsi'` matching a specific value:
```http
GET /tmf-api/partyManagement/v4/organization?filter=organizationIdentification%5B%3F(%40.identificationType%20%3D%3D%20'did%3Aelsi')%5D.identificationId%20%3D%3D%20'did%3Aelsi%3AVATES-12345678'
```
*Unencoded:* `organizationIdentification[?(@.identificationType == 'did:elsi')].identificationId == 'did:elsi:VATES-12345678'`

---

## 6. cURL Usage Examples

When using `curl`, you can let cURL handle URL-encoding with `--data-urlencode` via a GET request (`-G`):

```bash
# 1. Array wildcard filter
curl -G "https://example.com/tmf-api/productCatalogManagement/v4/productOffering" \
  --data-urlencode "filter=category[*].id == 'cat-fiber'"

# 2. OR query using semicolons
curl -G "https://example.com/tmf-api/productCatalogManagement/v4/productOffering" \
  --data-urlencode "filter=lifecycleStatus == 'Launched';lifecycleStatus == 'Active'"

# 3. AND query with multiple filters
curl -G "https://example.com/tmf-api/productCatalogManagement/v4/productOffering" \
  --data-urlencode "filter=lifecycleStatus == 'Launched'" \
  --data-urlencode "filter=category[*].id == 'cat-fiber'"

# 4. Combined with pagination
curl -G "https://example.com/tmf-api/productCatalogManagement/v4/productOffering" \
  --data-urlencode "filter=category[*].id == 'cat-fiber'" \
  -d "limit=20" \
  -d "offset=0"
```

---

## 7. Combining with Standard Parameters

The `filter` parameter integrates directly with standard TMF parameters:
- `limit`: Limits the number of returned objects (e.g. `&limit=20`).
- `offset`: Starting index for pagination (e.g. `&offset=40`).
- Direct indexed column filters (e.g. `&seller=did:elsi:...`).

**Combined Example:**
```http
GET /tmf-api/productCatalogManagement/v4/productOffering?filter=category%5B*%5D.id%20%3D%3D%20'cat-100'&seller=did:elsi:VATES-12345678&limit=10&offset=0
```

---

## 8. URL Encoding Quick Reference

| Character | URL-Encoded | Description |
| :---: | :---: | :--- |
| ` ` (space) | `%20` or `+` | Word separation |
| `[` | `%5B` | Array index / filter open |
| `]` | `%5D` | Array index / filter close |
| `?` | `%3F` | Filter predicate prefix |
| `@` | `%40` | Current element selector |
| `;` | `%3B` | **OR** expression separator |
| `&` | `&` (URL delimiter) | **AND** parameter separator |
| `=` | `%3D` | Equality operator |
| `'` | `%27` | Single quote delimiter |
| `"` | `%22` | Double quote delimiter |
| `%` | `%25` | Wildcard for `LIKE` operator |
| `:` | `%3A` | Colon in URIs / DIDs |
| `,` | `%2C` | Comma in `IN (...)` lists |
| `(` | `%28` | Opening parenthesis for `IN` |
| `)` | `%29` | Closing parenthesis for `IN` |

---

## 9. Unsupported Features & Best Practices

To ensure predictable performance and safety, the following constructs are **not supported**:
- Arbitrary script execution (e.g. JavaScript expressions `[?(@.length > 5)]`).
- Array slicing (e.g. `[0:2]` or `[-2:]`).
- Regular expression literals (e.g. `=~ /regex/`). Use the `LIKE` operator instead.
- Multiple logical conditions joined by `&&` or `||` inside a single predicate (e.g. `[?(@.a == 1 && @.b == 2)]`). Filter on the primary attribute or chain multiple `filter` parameters.
