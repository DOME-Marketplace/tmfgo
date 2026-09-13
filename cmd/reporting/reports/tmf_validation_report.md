# TMForum Object Validation Report

**Generated:** 2025-11-21 06:42:55 UTC

**Configuration:**
- Base URL: `https://tmf.dome-marketplace-sbx.org`
- Object Types: agreement, appliedCustomerBillingRate, resourceOrder, serviceCatalog, serviceCategory, serviceSpecification, cancelResourceOrder, category, partyRole, productOfferingPrice, serviceOrder, heal, monitor, partyAccount, product, productSpecification, resourceCategory, agreementSpecification, billingAccount, cancelProductOrder, catalog, customer, financialAccount, resourceCatalog, service, cancelServiceOrder, customerBill, customerBillOnDemand, individual, productOrder, resource, scale, usage, billPresentationMedia, migrate, billFormat, billingCycleSpecification, quote, resourceCandidate, resourceSpecification, settlementAccount, usageSpecification, organization, productOffering, resourceFunction, serviceCandidate
- Timeout: 30 seconds
- Validate Required Fields: true
- Validate Related Party: true

## Summary Statistics

| Metric | Value |
|--------|-------|
| Total Objects | 1082 |
| Valid Objects | 255 |
| Invalid Objects | 827 |
| Total Errors | 1378 |
| Total Warnings | 2413 |
| Total Errors Fixed | 0 |
| Total Warnings Fixed | 0 |
| Processing Time | 224.587µs |

## Statistics by Object Type

| Object Type | Count | Valid | Invalid | Errors | Warnings | Errors Fixed | Warnings Fixed |
|-------------|-------|-------|---------|--------|----------|--------------|----------------|
| agreement | 2 | 0 | 2 | 2 | 4 | 0 | 0 |
| billingAccount | 17 | 0 | 17 | 33 | 51 | 0 | 0 |
| catalog | 48 | 0 | 48 | 93 | 98 | 0 | 0 |
| category | 111 | 111 | 0 | 0 | 221 | 0 | 0 |
| customerBill | 41 | 0 | 41 | 118 | 123 | 0 | 0 |
| individual | 90 | 90 | 0 | 0 | 360 | 0 | 0 |
| organization | 26 | 26 | 0 | 0 | 93 | 0 | 0 |
| product | 43 | 0 | 43 | 106 | 129 | 0 | 0 |
| productOffering | 61 | 0 | 61 | 105 | 62 | 0 | 0 |
| productOfferingPrice | 363 | 0 | 363 | 363 | 482 | 0 | 0 |
| productOrder | 83 | 0 | 83 | 249 | 332 | 0 | 0 |
| productSpecification | 48 | 25 | 23 | 46 | 46 | 0 | 0 |
| resource | 28 | 0 | 28 | 56 | 84 | 0 | 0 |
| resourceSpecification | 16 | 0 | 16 | 16 | 32 | 0 | 0 |
| service | 82 | 0 | 82 | 164 | 246 | 0 | 0 |
| serviceSpecification | 16 | 3 | 13 | 13 | 29 | 0 | 0 |
| usageSpecification | 7 | 0 | 7 | 14 | 21 | 0 | 0 |

## Error Summary

| Error Code | Count |
|-------------|-------|
| MISSING_RELATED_PARTY_INFO | 1378 |

## Warning Summary

| Warning Code | Count |
|---------------|-------|
| MISSING_RECOMMENDED_FIELD | 2413 |

## Detailed Validation Results

### agreement Objects

#### Object: [urn:ngsi-ld:agreement:adc8a095-b1bd-471c-87f0-a49487618c7c](https://tmf.dome-marketplace-sbx.org/tmf-api/agreementManagement/v4/agreement/urn:ngsi-ld:agreement:adc8a095-b1bd-471c-87f0-a49487618c7c)

- **Type:** agreement
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to agreement (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:agreement:efffa430-8b52-4e9c-9e83-e1849ec9b162](https://tmf.dome-marketplace-sbx.org/tmf-api/agreementManagement/v4/agreement/urn:ngsi-ld:agreement:efffa430-8b52-4e9c-9e83-e1849ec9b162)

- **Type:** agreement
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to agreement (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

### billingAccount Objects

#### Object: [urn:ngsi-ld:billing-account:3b49919b-ab08-4969-ab53-cfd06cc21206](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:3b49919b-ab08-4969-ab53-cfd06cc21206)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:d2c224df-b007-4524-9029-7e7e1b021d35](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:d2c224df-b007-4524-9029-7e7e1b021d35)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:c84d03ff-fc74-435c-a54c-fed6e95ff80a](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:c84d03ff-fc74-435c-a54c-fed6e95ff80a)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:f98a4654-612b-474a-875b-107243e814bb](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:f98a4654-612b-474a-875b-107243e814bb)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:5a7a0a3f-61f8-4b62-bdf9-5212cbfc129c](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:5a7a0a3f-61f8-4b62-bdf9-5212cbfc129c)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:f727fb85-51f5-4dca-822f-4f4bb2775549](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:f727fb85-51f5-4dca-822f-4f4bb2775549)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:4da08929-a8a0-4362-aed8-a4d0d25032f2](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:4da08929-a8a0-4362-aed8-a4d0d25032f2)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:b62afbb3-0991-4c3d-85be-d96f95c55435](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:b62afbb3-0991-4c3d-85be-d96f95c55435)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:d8f73230-fe44-4594-9cc0-6c2b76107187](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:d8f73230-fe44-4594-9cc0-6c2b76107187)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:1011429f-061e-481c-9a1d-49226061b1c3](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:1011429f-061e-481c-9a1d-49226061b1c3)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:1ff83842-9917-4ac7-8f43-8e4e53a147ee](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:1ff83842-9917-4ac7-8f43-8e4e53a147ee)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:c2b41658-aabb-41da-ba36-d618ab076327](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:c2b41658-aabb-41da-ba36-d618ab076327)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:804ca8ab-75ec-469a-8bc9-8722c78a67db](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:804ca8ab-75ec-469a-8bc9-8722c78a67db)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:beb15d55-3a13-423f-bcb6-85066c45b984](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:beb15d55-3a13-423f-bcb6-85066c45b984)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:0ffd0fc2-4524-4f2c-881a-b5f4b936f33f](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:0ffd0fc2-4524-4f2c-881a-b5f4b936f33f)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:449b4aac-a10c-475a-8f9e-320bf2a0a4b7](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:449b4aac-a10c-475a-8f9e-320bf2a0a4b7)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:billing-account:84257bf5-e3b4-4305-926a-fc55e6f992ef](https://tmf.dome-marketplace-sbx.org/tmf-api/accountManagement/v4/billingAccount/urn:ngsi-ld:billing-account:84257bf5-e3b4-4305-926a-fc55e6f992ef)

- **Type:** billingAccount
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to billingAccount (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

### catalog Objects

#### Object: [urn:ngsi-ld:catalog:1d535a9b-212c-4e8b-aaff-8a412e61dd0d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:1d535a9b-212c-4e8b-aaff-8a412e61dd0d)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:e40feb38-04b6-485b-9edc-789704d3cd85](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:e40feb38-04b6-485b-9edc-789704d3cd85)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:03159bce-35ca-4938-bc2b-b8239e1008ca](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:03159bce-35ca-4938-bc2b-b8239e1008ca)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:c016bd48-8d84-48c9-82f1-8cb1d0c1cddd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:c016bd48-8d84-48c9-82f1-8cb1d0c1cddd)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:4a86c5a6-610f-47b0-8c0f-6bbf472409da](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:4a86c5a6-610f-47b0-8c0f-6bbf472409da)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:44e13ee8-c1ee-4c7a-a0f0-a308cd73a0e3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:44e13ee8-c1ee-4c7a-a0f0-a308cd73a0e3)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:ec45d9d1-50a1-445d-848e-09f50ca7862e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:ec45d9d1-50a1-445d-848e-09f50ca7862e)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:ad73d40f-11b6-4275-ab8a-6d1983e7b3ee](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:ad73d40f-11b6-4275-ab8a-6d1983e7b3ee)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:fb58a661-0341-4b61-9dac-610e228ad6bd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:fb58a661-0341-4b61-9dac-610e228ad6bd)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:2df6d831-26d8-430f-9ac7-79cf447c47fe](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:2df6d831-26d8-430f-9ac7-79cf447c47fe)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:721d9e67-0a46-4126-a12e-8f91670ceaf7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:721d9e67-0a46-4126-a12e-8f91670ceaf7)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:dfd8bf69-cb7e-4d0b-a3de-f407b2849580](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:dfd8bf69-cb7e-4d0b-a3de-f407b2849580)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:b2a1b435-d234-4f58-a23f-92c6ba5798b5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:b2a1b435-d234-4f58-a23f-92c6ba5798b5)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:134206dc-e658-4f12-a59b-ad2e17f4ede4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:134206dc-e658-4f12-a59b-ad2e17f4ede4)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:cd93ad8f-0052-4938-9843-cca9dc17156e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:cd93ad8f-0052-4938-9843-cca9dc17156e)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:5593c5c0-c62f-4d98-a0ef-fc596ce73f63](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:5593c5c0-c62f-4d98-a0ef-fc596ce73f63)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:c55cf33d-57ad-4d16-83ae-b0ac3e127472](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:c55cf33d-57ad-4d16-83ae-b0ac3e127472)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:18de6414-9deb-4a48-9100-c368629599d6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:18de6414-9deb-4a48-9100-c368629599d6)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:1bd42c14-9654-42ca-bb7c-2fad73861073](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:1bd42c14-9654-42ca-bb7c-2fad73861073)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:446acf1f-1303-4785-87bf-c17e481e4839](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:446acf1f-1303-4785-87bf-c17e481e4839)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:b212f3a4-209e-4574-b259-82a5729fec99](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:b212f3a4-209e-4574-b259-82a5729fec99)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:1e2f3493-1c39-41a4-bdee-65384c07e6e9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:1e2f3493-1c39-41a4-bdee-65384c07e6e9)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:e8e80643-495c-4a0a-bd4d-a75eef718f8b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:e8e80643-495c-4a0a-bd4d-a75eef718f8b)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:e5366fc2-509d-400c-ac2a-f7f473a9af59](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:e5366fc2-509d-400c-ac2a-f7f473a9af59)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:dcf716e9-ddbb-45a6-9aa5-9e6c29c7b2ee](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:dcf716e9-ddbb-45a6-9aa5-9e6c29c7b2ee)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:698132af-c9d6-4d45-8fd2-ccdee83f4c0f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:698132af-c9d6-4d45-8fd2-ccdee83f4c0f)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:18e99625-e4d3-4588-a670-98564d71c927](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:18e99625-e4d3-4588-a670-98564d71c927)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:a2fdefa0-43c8-4735-909e-6caaa4326a65](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:a2fdefa0-43c8-4735-909e-6caaa4326a65)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:062ab97e-3e5b-405f-afaa-5c06b6cccfc7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:062ab97e-3e5b-405f-afaa-5c06b6cccfc7)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:2e2beb9f-735e-4489-8b31-8be7c52f9522](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:2e2beb9f-735e-4489-8b31-8be7c52f9522)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:5b15d259-b2af-4b96-a281-d549f8c82fcb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:5b15d259-b2af-4b96-a281-d549f8c82fcb)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:4d20c693-17f8-4452-89cb-bdd9c15203a0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:4d20c693-17f8-4452-89cb-bdd9c15203a0)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:acb04f6a-06b0-4ced-9252-70fd524aa818](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:acb04f6a-06b0-4ced-9252-70fd524aa818)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:998b7410-341b-4aa5-acb5-063b2194c8c2](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:998b7410-341b-4aa5-acb5-063b2194c8c2)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:e10b5325-120d-4e9b-8e68-2fedf68a0fb5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:e10b5325-120d-4e9b-8e68-2fedf68a0fb5)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:0eb552fc-e2f5-4583-8d25-6f554dbd7a86](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:0eb552fc-e2f5-4583-8d25-6f554dbd7a86)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:d993a67a-750f-4e77-a0dc-f163940c25ac](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:d993a67a-750f-4e77-a0dc-f163940c25ac)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:cf620bc3-cf54-42c4-94d3-1e17c9b1ffc7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:cf620bc3-cf54-42c4-94d3-1e17c9b1ffc7)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:e6ab5c99-47a6-4870-b22b-648518e59a42](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:e6ab5c99-47a6-4870-b22b-648518e59a42)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:33b0e425-c1b9-4f14-9bb9-7d4303b1171e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:33b0e425-c1b9-4f14-9bb9-7d4303b1171e)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:57e71578-05b2-4d2b-9435-874665b0e955](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:57e71578-05b2-4d2b-9435-874665b0e955)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:74dbdcc0-f828-4e19-9dd0-3764d2ec3da5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:74dbdcc0-f828-4e19-9dd0-3764d2ec3da5)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:1b9bbf14-51f2-46b9-aeba-98b5e1446e21](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:1b9bbf14-51f2-46b9-aeba-98b5e1446e21)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:612e2b17-d4f8-47c3-bebe-bff91290a179](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:612e2b17-d4f8-47c3-bebe-bff91290a179)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:35bd03d5-b5b5-49ed-b0e1-8ad1282a193f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:35bd03d5-b5b5-49ed-b0e1-8ad1282a193f)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:e5f817da-cc9b-488b-beae-04a7e586b845](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:e5f817da-cc9b-488b-beae-04a7e586b845)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:d14f13f2-e96f-4584-aacf-a4d769bddf9d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:d14f13f2-e96f-4584-aacf-a4d769bddf9d)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:catalog:30aef942-e9ea-4b5f-bfc2-2d3ccf8a6e9d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/catalog/urn:ngsi-ld:catalog:30aef942-e9ea-4b5f-bfc2-2d3ccf8a6e9d)

- **Type:** catalog
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to catalog (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

### category Objects

#### Object: [urn:ngsi-ld:category:73dd1ca0-defd-4642-bbec-f44d52273973](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:73dd1ca0-defd-4642-bbec-f44d52273973)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:57158a09-c62b-4160-a49e-2295ca068682](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:57158a09-c62b-4160-a49e-2295ca068682)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:3cad4d15-bbe6-44d6-a82a-a6478b8e53ce](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:3cad4d15-bbe6-44d6-a82a-a6478b8e53ce)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:35182b9c-be04-4a40-ae67-f9433f4d36e1](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:35182b9c-be04-4a40-ae67-f9433f4d36e1)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:b332fe31-ec4d-4883-b051-e663a3d40830](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:b332fe31-ec4d-4883-b051-e663a3d40830)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:3ae62a2d-a194-4250-9b3f-8c9349ef4799](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:3ae62a2d-a194-4250-9b3f-8c9349ef4799)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:bb19f802-c0dd-4686-9049-cbd4665e7a04](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:bb19f802-c0dd-4686-9049-cbd4665e7a04)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:44370eaa-8828-4e1d-b9be-03b64eb11309](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:44370eaa-8828-4e1d-b9be-03b64eb11309)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:840bf629-6e64-4e49-94c0-740d7b312dd4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:840bf629-6e64-4e49-94c0-740d7b312dd4)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:7c796f0f-8e89-4578-bd6f-5d6c56ec912d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:7c796f0f-8e89-4578-bd6f-5d6c56ec912d)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:44e33c19-dd19-4ee9-a816-9dea2d34f02c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:44e33c19-dd19-4ee9-a816-9dea2d34f02c)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:d154f150-2980-4b5d-9aef-2fdc123dc4d9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:d154f150-2980-4b5d-9aef-2fdc123dc4d9)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:9bfdb312-ad44-41c5-9c5a-4d58ad861f63](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:9bfdb312-ad44-41c5-9c5a-4d58ad861f63)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:7aa7ae47-4c3d-447e-848b-60b3f015cb49](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:7aa7ae47-4c3d-447e-848b-60b3f015cb49)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:2b015a9f-50e6-453f-a251-e8770e2260bb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:2b015a9f-50e6-453f-a251-e8770e2260bb)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:b51bdeb5-0cf0-41a3-b0af-3ef48a88c988](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:b51bdeb5-0cf0-41a3-b0af-3ef48a88c988)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:15f52622-9097-4fba-a66b-435fb651a393](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:15f52622-9097-4fba-a66b-435fb651a393)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:bd99e5e4-5143-470e-9c05-d7839dfc71bc](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:bd99e5e4-5143-470e-9c05-d7839dfc71bc)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:ed12cd4a-3a23-46fe-8135-ea5bdfbb2b4b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:ed12cd4a-3a23-46fe-8135-ea5bdfbb2b4b)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:17c3aafe-2c2d-49be-8dce-dd4981e22e28](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:17c3aafe-2c2d-49be-8dce-dd4981e22e28)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:abe28b41-cc3a-4624-9c23-2ed3e021dbb5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:abe28b41-cc3a-4624-9c23-2ed3e021dbb5)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:08892d05-2e7a-41f9-8e8b-2dcc28f86b8a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:08892d05-2e7a-41f9-8e8b-2dcc28f86b8a)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:8337bace-8228-4925-a75a-e2654f0bac41](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:8337bace-8228-4925-a75a-e2654f0bac41)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:96778468-41b5-4245-842e-f3c481a240b0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:96778468-41b5-4245-842e-f3c481a240b0)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:c053c314-6696-49bd-a10b-052e6f3dbddb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:c053c314-6696-49bd-a10b-052e6f3dbddb)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:94bd8c25-4183-4124-a6ed-0516d4ced4a6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:94bd8c25-4183-4124-a6ed-0516d4ced4a6)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:630359eb-7d66-4eb2-bf07-2590841049a0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:630359eb-7d66-4eb2-bf07-2590841049a0)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:af87fc59-aafb-4301-8ba9-3594558a9eff](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:af87fc59-aafb-4301-8ba9-3594558a9eff)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:9c02c10e-1cdf-4de3-a99e-f0b689881d2b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:9c02c10e-1cdf-4de3-a99e-f0b689881d2b)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:8795e633-8a64-438f-8cad-a2d5175cc9cb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:8795e633-8a64-438f-8cad-a2d5175cc9cb)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:47eeb384-4400-49c1-b0e5-718ef17860ba](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:47eeb384-4400-49c1-b0e5-718ef17860ba)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:228fb947-48df-4281-b95d-c675bdc31b7a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:228fb947-48df-4281-b95d-c675bdc31b7a)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:03e76d62-a8c6-4571-ad7a-c31e754d3d3f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:03e76d62-a8c6-4571-ad7a-c31e754d3d3f)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:6ff41b67-bf97-4772-86ef-52d1d4998a3e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:6ff41b67-bf97-4772-86ef-52d1d4998a3e)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:419a8fac-d55a-4ae7-867f-2a3a338844cb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:419a8fac-d55a-4ae7-867f-2a3a338844cb)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:9a7f4ac4-e481-4303-98fc-b4f9911f6529](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:9a7f4ac4-e481-4303-98fc-b4f9911f6529)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:0b5cb3a5-b08e-4fca-8fb4-f99afaf71a2a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:0b5cb3a5-b08e-4fca-8fb4-f99afaf71a2a)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:061ba18d-4263-42f3-9d6e-17d90ad25832](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:061ba18d-4263-42f3-9d6e-17d90ad25832)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:b1f8c6b7-e73e-415a-9f28-51c01adc36c1](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:b1f8c6b7-e73e-415a-9f28-51c01adc36c1)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:55ff6485-3bcf-4009-815a-fca65727c64d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:55ff6485-3bcf-4009-815a-fca65727c64d)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:cfb68599-fde0-457e-bd8d-f6f6db20025f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:cfb68599-fde0-457e-bd8d-f6f6db20025f)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:ef80f5f1-c480-48f8-b96b-b893efec3e4d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:ef80f5f1-c480-48f8-b96b-b893efec3e4d)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:13837a4a-4eee-4879-be8b-da8bc88f8086](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:13837a4a-4eee-4879-be8b-da8bc88f8086)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:10ba35dd-4906-4434-adac-9bb9268e1127](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:10ba35dd-4906-4434-adac-9bb9268e1127)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:b4998d10-d9c2-46ad-a08d-63df7cd8e911](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:b4998d10-d9c2-46ad-a08d-63df7cd8e911)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:3ebbe833-9f85-42d2-b76d-7429d156eba6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:3ebbe833-9f85-42d2-b76d-7429d156eba6)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:3beea627-9691-4ed8-9bbf-7c6fe06e278a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:3beea627-9691-4ed8-9bbf-7c6fe06e278a)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:818597c8-6de7-4961-8ce3-61fb30d7fbb8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:818597c8-6de7-4961-8ce3-61fb30d7fbb8)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:a456f7a7-cbb0-4b71-9c25-f8cb46555016](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:a456f7a7-cbb0-4b71-9c25-f8cb46555016)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:2cc758fa-80b3-4704-9c38-26918c25b535](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:2cc758fa-80b3-4704-9c38-26918c25b535)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:6df17fbc-465d-4e0e-93af-ecc0ccdf34b1](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:6df17fbc-465d-4e0e-93af-ecc0ccdf34b1)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:9c12beda-8e3b-4b90-ba96-191e0129a9d4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:9c12beda-8e3b-4b90-ba96-191e0129a9d4)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:f0c23dff-8f09-474a-8306-607ca835d8ce](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:f0c23dff-8f09-474a-8306-607ca835d8ce)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:c9f5f9b8-2991-4cb5-aad9-eede449c88e8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:c9f5f9b8-2991-4cb5-aad9-eede449c88e8)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:917fa857-f354-49f1-84f3-38540e07e434](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:917fa857-f354-49f1-84f3-38540e07e434)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:4c9b41bb-2eeb-4229-8442-b0e8d8127271](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:4c9b41bb-2eeb-4229-8442-b0e8d8127271)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:465a29be-b24c-448e-80ed-c9a1e5aeb3da](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:465a29be-b24c-448e-80ed-c9a1e5aeb3da)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:ac6eefc6-1e58-432f-a9b9-c8349c3a3d68](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:ac6eefc6-1e58-432f-a9b9-c8349c3a3d68)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:1d731f47-e2ce-4f2a-899b-f7f1c449cd98](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:1d731f47-e2ce-4f2a-899b-f7f1c449cd98)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:63ed0da1-bc76-4e8d-af4d-f551d8148394](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:63ed0da1-bc76-4e8d-af4d-f551d8148394)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:2fffafeb-64ca-4d72-9e4c-ed7d2973b3a9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:2fffafeb-64ca-4d72-9e4c-ed7d2973b3a9)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:fcf925b0-bb66-4860-8c2a-db17861fe105](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:fcf925b0-bb66-4860-8c2a-db17861fe105)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:8ec9aa71-afa0-4ae4-93d2-61eef79fe9c6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:8ec9aa71-afa0-4ae4-93d2-61eef79fe9c6)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:b664a377-e12c-465b-ae2c-0c5a53438e65](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:b664a377-e12c-465b-ae2c-0c5a53438e65)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:a6f7ec89-e2e3-429d-a16c-30b0a1ab0ff6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:a6f7ec89-e2e3-429d-a16c-30b0a1ab0ff6)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:4b11ef5c-50a6-4a20-8817-8152b372cc95](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:4b11ef5c-50a6-4a20-8817-8152b372cc95)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:69261f7b-f228-4aa5-a180-02f2e4c192bf](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:69261f7b-f228-4aa5-a180-02f2e4c192bf)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:9508e28c-dca7-4cdb-b49a-8ad9b473a106](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:9508e28c-dca7-4cdb-b49a-8ad9b473a106)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:2fe2b111-636e-4ccc-96fc-4a2bde912e31](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:2fe2b111-636e-4ccc-96fc-4a2bde912e31)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:e0c4ad97-585d-44c7-99b7-f55e31be1ce3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:e0c4ad97-585d-44c7-99b7-f55e31be1ce3)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:c8286912-7ad4-431c-a112-f3c0298c1806](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:c8286912-7ad4-431c-a112-f3c0298c1806)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:e464a506-ef2c-468c-9033-ec5572d5076e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:e464a506-ef2c-468c-9033-ec5572d5076e)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:91764e83-85e2-48c8-9311-b19ee59219cc](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:91764e83-85e2-48c8-9311-b19ee59219cc)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:0b4a0d9e-0a4e-4489-9da0-6ffd18e1cbc4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:0b4a0d9e-0a4e-4489-9da0-6ffd18e1cbc4)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:0ed1e073-6168-4bcf-b249-221c8f386507](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:0ed1e073-6168-4bcf-b249-221c8f386507)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:0d52ebea-3b4c-45c0-9dc7-c1379135950d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:0d52ebea-3b4c-45c0-9dc7-c1379135950d)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:821c1437-f982-446e-866c-c43a680bf0f0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:821c1437-f982-446e-866c-c43a680bf0f0)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:0b74e44c-c05b-4bbe-8530-a73a4d9cf928](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:0b74e44c-c05b-4bbe-8530-a73a4d9cf928)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:b334a7a2-3e7d-473e-a042-09ba2ff6cff6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:b334a7a2-3e7d-473e-a042-09ba2ff6cff6)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:31ca6777-cc13-4d74-bbc7-b8bc10e6fcd7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:31ca6777-cc13-4d74-bbc7-b8bc10e6fcd7)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:8cdb4483-2185-4dc6-afd0-a70d30010849](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:8cdb4483-2185-4dc6-afd0-a70d30010849)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:39a55d2f-6014-4cab-8f40-cc564e2e253b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:39a55d2f-6014-4cab-8f40-cc564e2e253b)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:607af425-d44e-4274-be73-9493297ec4a5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:607af425-d44e-4274-be73-9493297ec4a5)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:4cb199d8-4bbb-4ea7-b811-03a54e6fcbc9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:4cb199d8-4bbb-4ea7-b811-03a54e6fcbc9)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:424eca22-053f-458e-8bd2-51318b733f7e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:424eca22-053f-458e-8bd2-51318b733f7e)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:0186f2e7-869f-4118-83c5-7de931733af0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:0186f2e7-869f-4118-83c5-7de931733af0)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:b523a55b-29fe-431c-809b-4339ae34d944](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:b523a55b-29fe-431c-809b-4339ae34d944)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:f5e856f7-a619-4fd4-b9d2-99fe8a778db9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:f5e856f7-a619-4fd4-b9d2-99fe8a778db9)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:8d222bda-159e-4957-8b76-0fb06b4449dd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:8d222bda-159e-4957-8b76-0fb06b4449dd)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:d6aa17c8-a0d4-4312-9ab1-ae78a0aace1d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:d6aa17c8-a0d4-4312-9ab1-ae78a0aace1d)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:814efafb-b2be-4d7e-b83d-0d447d9f68ac](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:814efafb-b2be-4d7e-b83d-0d447d9f68ac)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:91abfd48-4870-4af1-b64f-69cb144d1dc6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:91abfd48-4870-4af1-b64f-69cb144d1dc6)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:a14a4067-a5c9-4122-b2a7-28ee2dd85036](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:a14a4067-a5c9-4122-b2a7-28ee2dd85036)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:dc159e2d-d0a1-49e2-a2ec-b7c4c97e4dd7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:dc159e2d-d0a1-49e2-a2ec-b7c4c97e4dd7)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:55b2cb9f-e13b-499e-bf76-fd1eb56d2bde](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:55b2cb9f-e13b-499e-bf76-fd1eb56d2bde)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:405c922f-0de1-4e7f-b478-713240b9b381](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:405c922f-0de1-4e7f-b478-713240b9b381)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:466cc605-e806-43e1-9291-84deac0028f1](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:466cc605-e806-43e1-9291-84deac0028f1)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:b96911a6-e09e-4776-b4fc-4faf7004a162](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:b96911a6-e09e-4776-b4fc-4faf7004a162)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:ffeae64f-edc2-4754-906e-a4d35d10c806](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:ffeae64f-edc2-4754-906e-a4d35d10c806)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:b0cd7c24-e3c1-49a7-8fc4-8e8d492d513b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:b0cd7c24-e3c1-49a7-8fc4-8e8d492d513b)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:d64a7aff-f49c-4cbf-84d8-4b883e59e392](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:d64a7aff-f49c-4cbf-84d8-4b883e59e392)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:a7779dea-c0c0-47af-b078-3098c309cc23](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:a7779dea-c0c0-47af-b078-3098c309cc23)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:ef149b50-08ea-4ee2-bf41-a90b04a1ce79](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:ef149b50-08ea-4ee2-bf41-a90b04a1ce79)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:3867db98-1eb3-4bff-8a23-715082f404ff](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:3867db98-1eb3-4bff-8a23-715082f404ff)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:1d359cb4-dc15-4c87-afd2-da18a2940b5b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:1d359cb4-dc15-4c87-afd2-da18a2940b5b)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:9b57eaca-bc73-4283-8226-8c8666498cfa](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:9b57eaca-bc73-4283-8226-8c8666498cfa)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:5b55bcbc-4776-4702-b400-bf32a1821fb4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:5b55bcbc-4776-4702-b400-bf32a1821fb4)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:961389ac-7e86-4e80-864a-0f130cc9863b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:961389ac-7e86-4e80-864a-0f130cc9863b)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:515f7838-331f-4164-bb9f-8d843e04a101](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:515f7838-331f-4164-bb9f-8d843e04a101)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:0c4373c1-9a3b-4797-9880-24d170a350ee](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:0c4373c1-9a3b-4797-9880-24d170a350ee)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:category:645f0c9a-816c-4cf8-a307-2503c4c75742](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/category/urn:ngsi-ld:category:645f0c9a-816c-4cf8-a307-2503c4c75742)

- **Type:** category
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to category (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

### customerBill Objects

#### Object: [urn:ngsi-ld:customer-bill:87fe1707-c5b3-4a74-9d67-3113c2b07045](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:87fe1707-c5b3-4a74-9d67-3113c2b07045)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:5efe1990-db82-4eef-99e3-d3d4671ca625](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:5efe1990-db82-4eef-99e3-d3d4671ca625)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:beb50ce0-3ab9-403f-84e6-ee03c44be13b](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:beb50ce0-3ab9-403f-84e6-ee03c44be13b)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:f59a5928-230f-4cd1-a354-c993c565a4a6](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:f59a5928-230f-4cd1-a354-c993c565a4a6)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:5f53afa7-a1c6-458a-8dd7-45da0b8ef3e7](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:5f53afa7-a1c6-458a-8dd7-45da0b8ef3e7)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:b7b8bf50-300d-494b-a406-0f2a3778fcbc](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:b7b8bf50-300d-494b-a406-0f2a3778fcbc)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:80397d05-2196-4e39-9265-4610e37fde50](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:80397d05-2196-4e39-9265-4610e37fde50)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:61c3057a-e565-4bb4-940a-61dd30fd6de3](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:61c3057a-e565-4bb4-940a-61dd30fd6de3)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:885f82ce-116d-4271-95ac-5b1150f210a4](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:885f82ce-116d-4271-95ac-5b1150f210a4)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:e642addf-48e2-4627-8489-a1bad7d09e10](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:e642addf-48e2-4627-8489-a1bad7d09e10)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:7753b831-d482-44b1-a3de-9790c7df2f40](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:7753b831-d482-44b1-a3de-9790c7df2f40)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:f25072b8-572e-4d53-898d-3fef1c7bd3d6](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:f25072b8-572e-4d53-898d-3fef1c7bd3d6)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:b542fc2a-975e-446e-b18a-2a6d37957f0b](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:b542fc2a-975e-446e-b18a-2a6d37957f0b)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:e1b3be4f-dcc6-4133-b2b3-47154665444a](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:e1b3be4f-dcc6-4133-b2b3-47154665444a)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:82bc0490-7868-4f71-ae26-94dbb314eb17](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:82bc0490-7868-4f71-ae26-94dbb314eb17)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:d45f9686-e3b6-433d-980f-3261aaafaa74](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:d45f9686-e3b6-433d-980f-3261aaafaa74)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:f7555d3f-c980-4e60-baa5-bacaeaf318df](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:f7555d3f-c980-4e60-baa5-bacaeaf318df)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:77c737cf-3e7e-4efd-bc15-d4e55d4de86e](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:77c737cf-3e7e-4efd-bc15-d4e55d4de86e)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:b5bf019c-b42e-47e5-9fa5-e3cb042c1612](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:b5bf019c-b42e-47e5-9fa5-e3cb042c1612)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:c84035b8-0282-4893-bf76-41d026e9d795](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:c84035b8-0282-4893-bf76-41d026e9d795)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:28cf30a1-92af-4b50-99c2-83ef2e6ab12c](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:28cf30a1-92af-4b50-99c2-83ef2e6ab12c)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:9143c42b-8429-45e6-812e-ba8a29b840b9](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:9143c42b-8429-45e6-812e-ba8a29b840b9)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:f07eebc8-435b-4abe-ad98-a17117031588](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:f07eebc8-435b-4abe-ad98-a17117031588)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:9d17558b-d31a-40cd-9358-f7d671577beb](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:9d17558b-d31a-40cd-9358-f7d671577beb)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:4defc8f0-46a5-4fdc-8e27-b608b80578aa](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:4defc8f0-46a5-4fdc-8e27-b608b80578aa)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:0b94eac0-1103-4fa3-8e2b-e75bb78e45bc](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:0b94eac0-1103-4fa3-8e2b-e75bb78e45bc)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:77a43a73-397d-4aea-90f2-1d1bd5e37a2a](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:77a43a73-397d-4aea-90f2-1d1bd5e37a2a)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:0c2aa78f-4c1b-44e0-b31a-b753066bba65](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:0c2aa78f-4c1b-44e0-b31a-b753066bba65)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:48c5e7eb-006d-48c4-b7db-d43602e42c65](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:48c5e7eb-006d-48c4-b7db-d43602e42c65)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Buyer' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:385c62cf-2d3b-403b-9d27-9e2297f99b0e](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:385c62cf-2d3b-403b-9d27-9e2297f99b0e)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Buyer' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:4f0d1fdd-78fc-4256-8edf-4569dc2bb68c](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:4f0d1fdd-78fc-4256-8edf-4569dc2bb68c)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Buyer' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:b9119a39-c32e-40fc-a574-2805d1b14789](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:b9119a39-c32e-40fc-a574-2805d1b14789)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Buyer' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:4d1ed9c5-1287-496a-a7c0-6262a0c96940](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:4d1ed9c5-1287-496a-a7c0-6262a0c96940)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Buyer' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:2099f627-3bdd-4d5e-8e18-c5dcdc799555](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:2099f627-3bdd-4d5e-8e18-c5dcdc799555)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:58dc4ab8-62bf-4a20-837c-d28fdaf714e9](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:58dc4ab8-62bf-4a20-837c-d28fdaf714e9)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Buyer' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:25329f76-abae-441d-84b3-8a8f800fed09](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:25329f76-abae-441d-84b3-8a8f800fed09)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Buyer' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:9fca8c85-c7cb-44be-ab3d-1912a32fae6e](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:9fca8c85-c7cb-44be-ab3d-1912a32fae6e)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:db0368e8-9e77-4691-8618-d6514dd0a6d6](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:db0368e8-9e77-4691-8618-d6514dd0a6d6)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:ec2bd9a5-94f5-4eaf-9e77-34b6fc43de39](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:ec2bd9a5-94f5-4eaf-9e77-34b6fc43de39)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:db07e589-adf2-409a-a196-705b011f69d0](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:db07e589-adf2-409a-a196-705b011f69d0)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:customer-bill:cc3e86c0-61aa-4655-91bb-0e881e6c7df6](https://tmf.dome-marketplace-sbx.org/tmf-api/customerBillManagement/v4/customerBill/urn:ngsi-ld:customer-bill:cc3e86c0-61aa-4655-91bb-0e881e6c7df6)

- **Type:** customerBill
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to customerBill (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

### individual Objects

#### Object: [urn:ngsi-ld:individual:97790bad-066f-41fa-bb5a-745d783e044e](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:97790bad-066f-41fa-bb5a-745d783e044e)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:038e6cf4-8805-4905-a151-42012ddcdfbf](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:038e6cf4-8805-4905-a151-42012ddcdfbf)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:c1046c38-844f-471f-873c-78b35c270c27](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:c1046c38-844f-471f-873c-78b35c270c27)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:afa0bf6f-0b1e-4d9c-8529-e57bd28740f8](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:afa0bf6f-0b1e-4d9c-8529-e57bd28740f8)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:911f60f5-9a54-4d1c-ae99-46079d5c6143](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:911f60f5-9a54-4d1c-ae99-46079d5c6143)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:9506e609-cf28-4302-9592-7a49521d0412](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:9506e609-cf28-4302-9592-7a49521d0412)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:c5b57368-f3cc-46ed-838d-25153f471c5d](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:c5b57368-f3cc-46ed-838d-25153f471c5d)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:d6cccbbb-3fae-413f-89b9-c3d41ea25362](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:d6cccbbb-3fae-413f-89b9-c3d41ea25362)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:73ecc3c1-41b9-48a3-8c99-b29e6e02ad2d](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:73ecc3c1-41b9-48a3-8c99-b29e6e02ad2d)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:8f307e88-4208-4484-9a62-b06b9b4d4a29](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:8f307e88-4208-4484-9a62-b06b9b4d4a29)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:9e2e4859-c2f4-491c-a787-5f3aeec93ef6](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:9e2e4859-c2f4-491c-a787-5f3aeec93ef6)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:b8d107f5-5888-413c-a334-df87468398ad](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:b8d107f5-5888-413c-a334-df87468398ad)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:d3b2d0a8-d490-43df-8f00-f76b45cae5af](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:d3b2d0a8-d490-43df-8f00-f76b45cae5af)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:68f857d5-4de3-4f22-8daf-0b46d9699a8d](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:68f857d5-4de3-4f22-8daf-0b46d9699a8d)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:5c9d40b0-5628-4dec-9228-08fbee4bcce2](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:5c9d40b0-5628-4dec-9228-08fbee4bcce2)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:46b73e37-a453-4b19-8df0-85356a52cd8a](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:46b73e37-a453-4b19-8df0-85356a52cd8a)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:b266d7da-0852-4659-a003-820b391f7927](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:b266d7da-0852-4659-a003-820b391f7927)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:98d935c4-5f25-42de-a000-844950f1b766](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:98d935c4-5f25-42de-a000-844950f1b766)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:0774f8c5-073e-4334-97b2-5b71167dfc2b](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:0774f8c5-073e-4334-97b2-5b71167dfc2b)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:a0e4422b-1464-42a5-ac76-2901329a5c8b](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:a0e4422b-1464-42a5-ac76-2901329a5c8b)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:b068358a-d4c0-4f20-a522-ac6286181d49](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:b068358a-d4c0-4f20-a522-ac6286181d49)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:c475f14f-97b3-45ec-8a34-4f22f4931f04](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:c475f14f-97b3-45ec-8a34-4f22f4931f04)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:f366c06b-38b2-4039-b66a-8dde9a88f1f2](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:f366c06b-38b2-4039-b66a-8dde9a88f1f2)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:80bb5420-a8c4-47a6-aa73-8c96ccc9a525](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:80bb5420-a8c4-47a6-aa73-8c96ccc9a525)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:6aebe06a-54d2-4f17-87c4-940c1e22ef91](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:6aebe06a-54d2-4f17-87c4-940c1e22ef91)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:bb26a135-8281-430c-b2d0-652b752b0c15](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:bb26a135-8281-430c-b2d0-652b752b0c15)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:5e7c0d3f-235e-4507-9875-a166fb41fb35](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:5e7c0d3f-235e-4507-9875-a166fb41fb35)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:551f1882-63c9-44fb-b001-d9c2648ba6bc](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:551f1882-63c9-44fb-b001-d9c2648ba6bc)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:16a87ab8-2d62-4298-97dd-6c96f2556397](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:16a87ab8-2d62-4298-97dd-6c96f2556397)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:521e32a5-bdc3-4530-a878-644299792594](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:521e32a5-bdc3-4530-a878-644299792594)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:2419897b-0da3-41f2-a34f-d92ccff3a9e2](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:2419897b-0da3-41f2-a34f-d92ccff3a9e2)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:67c4985f-2ed1-43d2-a631-6261a66f7c8e](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:67c4985f-2ed1-43d2-a631-6261a66f7c8e)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:6176fae2-ef4e-4bb5-aae6-f7d6e41efc3f](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:6176fae2-ef4e-4bb5-aae6-f7d6e41efc3f)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:c6abe161-5c9a-4d56-9817-917162f1e453](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:c6abe161-5c9a-4d56-9817-917162f1e453)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:b8be3518-b3ff-43a0-b8a4-5310a9ada50d](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:b8be3518-b3ff-43a0-b8a4-5310a9ada50d)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:b1256dde-1380-4ed5-95bb-8ebdcce08774](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:b1256dde-1380-4ed5-95bb-8ebdcce08774)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:50b2b804-bd99-4d30-951a-b536d3e56bf9](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:50b2b804-bd99-4d30-951a-b536d3e56bf9)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:c4177a93-1966-4fe3-bb02-ec5212b11a17](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:c4177a93-1966-4fe3-bb02-ec5212b11a17)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:775c5cc0-ae6d-4871-a8bb-c6919ef6b8c2](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:775c5cc0-ae6d-4871-a8bb-c6919ef6b8c2)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:c66bb2f7-e174-4795-a4d8-b71861080d73](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:c66bb2f7-e174-4795-a4d8-b71861080d73)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:01e063e2-505e-4e88-a571-487cba934da0](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:01e063e2-505e-4e88-a571-487cba934da0)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:1ad08cb8-8a19-4ee7-8090-a7c41c87f3d6](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:1ad08cb8-8a19-4ee7-8090-a7c41c87f3d6)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:47e67fa7-7cc6-42df-a455-b4c76c53ae56](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:47e67fa7-7cc6-42df-a455-b4c76c53ae56)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:16151a07-a931-4de5-acde-fb95b4473d1f](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:16151a07-a931-4de5-acde-fb95b4473d1f)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:4f276f57-6429-4830-938a-ed8f5f80e668](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:4f276f57-6429-4830-938a-ed8f5f80e668)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:fbddcd79-2096-4d2e-9884-fff079fc4bb2](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:fbddcd79-2096-4d2e-9884-fff079fc4bb2)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:49f59271-fe50-4917-9133-0b52d2e485cf](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:49f59271-fe50-4917-9133-0b52d2e485cf)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:32224a86-369a-46bb-8824-c56873b8df1d](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:32224a86-369a-46bb-8824-c56873b8df1d)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:936d68b8-3f6c-4913-adbe-cbde3db52e32](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:936d68b8-3f6c-4913-adbe-cbde3db52e32)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:7d0c44d0-2898-432d-bebb-58aeb0b9c5db](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:7d0c44d0-2898-432d-bebb-58aeb0b9c5db)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:561cb679-d7d3-4bed-b4e8-11f6f602bc81](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:561cb679-d7d3-4bed-b4e8-11f6f602bc81)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:2e573a9c-1c0e-4c43-af43-8f46245b06fb](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:2e573a9c-1c0e-4c43-af43-8f46245b06fb)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:5ec0e661-b7e2-4f16-bf40-010350a9e16c](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:5ec0e661-b7e2-4f16-bf40-010350a9e16c)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:99258028-9b1b-40ef-8f1a-3c7cd8c0bb27](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:99258028-9b1b-40ef-8f1a-3c7cd8c0bb27)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:442c5ac4-a40d-4fc6-90d8-db4594cde0cd](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:442c5ac4-a40d-4fc6-90d8-db4594cde0cd)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:97f2fc00-b629-4189-9925-73063b98aabc](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:97f2fc00-b629-4189-9925-73063b98aabc)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:5c2c34cb-ecfa-4c8e-b1a8-ec7487f708f4](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:5c2c34cb-ecfa-4c8e-b1a8-ec7487f708f4)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:1992e134-bb50-444e-9ee4-d1b3089f9088](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:1992e134-bb50-444e-9ee4-d1b3089f9088)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:fa4043a7-393d-4e8d-9eee-1766fde861dd](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:fa4043a7-393d-4e8d-9eee-1766fde861dd)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:5cee7312-69f8-4a15-9f48-8aa4833bd628](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:5cee7312-69f8-4a15-9f48-8aa4833bd628)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:eec0f329-1f83-4c16-bd65-aeae879c5566](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:eec0f329-1f83-4c16-bd65-aeae879c5566)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:3be1d80a-6cdf-410a-abce-36dc51dabbc3](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:3be1d80a-6cdf-410a-abce-36dc51dabbc3)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:c60bca70-2dc9-448f-8031-5f3274b17f42](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:c60bca70-2dc9-448f-8031-5f3274b17f42)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:503c45b1-c5c2-4675-a186-c8f5a57d122b](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:503c45b1-c5c2-4675-a186-c8f5a57d122b)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:7e3ea768-5c13-4d2c-93e7-664e83a9c553](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:7e3ea768-5c13-4d2c-93e7-664e83a9c553)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:613a6f20-c5f2-4678-9d9c-9c948d60d9e5](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:613a6f20-c5f2-4678-9d9c-9c948d60d9e5)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:262d5b68-4371-408f-8750-646c8f938a33](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:262d5b68-4371-408f-8750-646c8f938a33)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:4b66482b-8e2f-4f76-b4fa-44b1ba91fa98](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:4b66482b-8e2f-4f76-b4fa-44b1ba91fa98)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:1dd8899c-b9bd-4646-87a7-25701cb3bb7c](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:1dd8899c-b9bd-4646-87a7-25701cb3bb7c)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:9097b22c-5ec7-4073-8e8b-3ebda5cb1c19](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:9097b22c-5ec7-4073-8e8b-3ebda5cb1c19)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:f8aa794f-a54b-4025-8c0c-8162a79745ee](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:f8aa794f-a54b-4025-8c0c-8162a79745ee)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:093c3e41-5b33-4558-86d4-c3bd879e3778](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:093c3e41-5b33-4558-86d4-c3bd879e3778)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:8c8e8670-b1f4-4a70-a916-c69ae780fc88](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:8c8e8670-b1f4-4a70-a916-c69ae780fc88)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:e94a5bda-e810-4bbb-aaff-538ab4846ba0](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:e94a5bda-e810-4bbb-aaff-538ab4846ba0)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:0e24e8d0-2898-4afa-a8c3-15f563ade93f](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:0e24e8d0-2898-4afa-a8c3-15f563ade93f)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:410cc56f-445e-45fc-98fb-2d8998fc0b18](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:410cc56f-445e-45fc-98fb-2d8998fc0b18)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:49e20e20-0616-41d4-a4b0-0d2450307b34](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:49e20e20-0616-41d4-a4b0-0d2450307b34)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:aab1c921-cb9b-4bb2-b198-213e5d51f382](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:aab1c921-cb9b-4bb2-b198-213e5d51f382)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:e305e366-84ab-4ce0-938b-96b3fa6e372a](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:e305e366-84ab-4ce0-938b-96b3fa6e372a)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:65d60c41-5e5e-47a3-aa57-2eaa5265d451](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:65d60c41-5e5e-47a3-aa57-2eaa5265d451)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:e76d5820-65e7-4f65-94cc-48b9b646557e](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:e76d5820-65e7-4f65-94cc-48b9b646557e)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:f2d2460c-e9d0-4a02-ae3d-06a65e608e69](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:f2d2460c-e9d0-4a02-ae3d-06a65e608e69)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:df2c1ef3-1d77-411c-bdda-ba0f0ad20649](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:df2c1ef3-1d77-411c-bdda-ba0f0ad20649)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:d7dfbc4e-fae0-4f0e-98aa-3ef3c9a2a893](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:d7dfbc4e-fae0-4f0e-98aa-3ef3c9a2a893)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:1dbfcd6b-9fae-4df1-97be-cc6a331840b6](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:1dbfcd6b-9fae-4df1-97be-cc6a331840b6)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:0ce6cb59-1f98-4a2b-82fa-6819f217866f](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:0ce6cb59-1f98-4a2b-82fa-6819f217866f)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:79d65e28-ea8a-40dd-89a6-32ee266af46d](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:79d65e28-ea8a-40dd-89a6-32ee266af46d)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:a6db3ab3-d6dd-41a3-bc76-3fc7edd440b2](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:a6db3ab3-d6dd-41a3-bc76-3fc7edd440b2)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:ff3a3436-fdc4-4214-a2f8-eb2ea0d80e26](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:ff3a3436-fdc4-4214-a2f8-eb2ea0d80e26)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:individual:622547ef-076a-4aba-90cc-f899b75bd282](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/individual/urn:ngsi-ld:individual:622547ef-076a-4aba-90cc-f899b75bd282)

- **Type:** individual
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to individual (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

### organization Objects

#### Object: [urn:ngsi-ld:organization:00026ead-d16c-40ee-9218-48defc7ce749](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:00026ead-d16c-40ee-9218-48defc7ce749)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:a2f5ebea-49c9-4015-a9d6-56f2c566f6c9](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:a2f5ebea-49c9-4015-a9d6-56f2c566f6c9)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:d8dcc7a3-0774-4824-b552-d05d91986565](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:d8dcc7a3-0774-4824-b552-d05d91986565)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:c2c39ab0-1f15-40f9-89ce-fc131953d33e](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:c2c39ab0-1f15-40f9-89ce-fc131953d33e)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:4eb54a0c-a916-499a-bf5f-6bba76101e1e](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:4eb54a0c-a916-499a-bf5f-6bba76101e1e)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:fb47dd79-6497-4ec8-8456-e6e06d3c698b](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:fb47dd79-6497-4ec8-8456-e6e06d3c698b)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:44c33614-df98-459e-b710-064b1a7c6f65](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:44c33614-df98-459e-b710-064b1a7c6f65)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:6c53e937-212b-4e8b-997c-4d8695f789d1](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:6c53e937-212b-4e8b-997c-4d8695f789d1)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:338282a4-3c06-41d0-8c35-3fe5cecc38db](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:338282a4-3c06-41d0-8c35-3fe5cecc38db)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:3c9beb13-1187-4d67-9ecf-5e6825ffb2fd](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:3c9beb13-1187-4d67-9ecf-5e6825ffb2fd)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:5ca475ec-52d3-4a0c-a1d5-0071ef09b59d](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:5ca475ec-52d3-4a0c-a1d5-0071ef09b59d)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:e4fa0e9f-1779-49c0-9e8a-66a02bf1fe4e](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:e4fa0e9f-1779-49c0-9e8a-66a02bf1fe4e)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:5f4a3634-efba-4e25-bf15-569492850093](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:5f4a3634-efba-4e25-bf15-569492850093)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:ae659146-0bab-434f-820c-6fb46701f553](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:ae659146-0bab-434f-820c-6fb46701f553)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:df924e5d-e8c8-4ea4-aca8-edaf5acdc109](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:df924e5d-e8c8-4ea4-aca8-edaf5acdc109)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:95fdc12e-6889-4f08-8ff8-296b10e8e781](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:95fdc12e-6889-4f08-8ff8-296b10e8e781)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:eb6647da-84f2-4645-8d9f-c2905775b561](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:eb6647da-84f2-4645-8d9f-c2905775b561)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:3d660d74-8f2a-444f-9451-9a6c8a0b3b66](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:3d660d74-8f2a-444f-9451-9a6c8a0b3b66)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:2bcbe859-e316-42f2-919c-f470cff9e235](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:2bcbe859-e316-42f2-919c-f470cff9e235)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:339fb0e7-3d61-4022-8d64-84a8dbe610f8](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:339fb0e7-3d61-4022-8d64-84a8dbe610f8)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:add54e87-7742-4072-a88c-789b24b92d03](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:add54e87-7742-4072-a88c-789b24b92d03)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:649e20ba-8cdc-4090-ba5e-a835b63123fc](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:649e20ba-8cdc-4090-ba5e-a835b63123fc)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:f465f0c1-b735-49f1-a1e1-c02200f250e9](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:f465f0c1-b735-49f1-a1e1-c02200f250e9)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:3af4da45-59b1-4271-84b4-d3e87bc907ae](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:3af4da45-59b1-4271-84b4-d3e87bc907ae)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:1fc74f3a-d918-4116-8d64-163b11cf8170](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:1fc74f3a-d918-4116-8d64-163b11cf8170)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:organization:63ded74e-f2b5-4f73-a34b-3422f5ff3fee](https://tmf.dome-marketplace-sbx.org/tmf-api/party/v4/organization/urn:ngsi-ld:organization:63ded74e-f2b5-4f73-a34b-3422f5ff3fee)

- **Type:** organization
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:48 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to organization (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

### product Objects

#### Object: [urn:ngsi-ld:product:45eaa84b-122f-4e6d-9797-31ab1ab16134](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:45eaa84b-122f-4e6d-9797-31ab1ab16134)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:b0a89e11-6fde-4bd0-948b-508088655e63](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:b0a89e11-6fde-4bd0-948b-508088655e63)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:34c6acaf-e895-4747-b371-6410d05e6a73](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:34c6acaf-e895-4747-b371-6410d05e6a73)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Buyer' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:0226d089-0f29-4ac5-a78e-9ca7fc8e4488](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:0226d089-0f29-4ac5-a78e-9ca7fc8e4488)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Buyer' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:b669d51d-5fc0-4f3a-aa13-532f5b45bbbc](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:b669d51d-5fc0-4f3a-aa13-532f5b45bbbc)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:3fddee64-1256-496a-aa15-cdcf69b26642](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:3fddee64-1256-496a-aa15-cdcf69b26642)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:eb8c00a0-f013-4d7a-9b3c-477a72c0cd01](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:eb8c00a0-f013-4d7a-9b3c-477a72c0cd01)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:e308688a-28bd-479c-87c5-ee3ea307dcde](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:e308688a-28bd-479c-87c5-ee3ea307dcde)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:61540da5-526b-4063-b783-cd4e2c0c7d2f](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:61540da5-526b-4063-b783-cd4e2c0c7d2f)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:76c8df70-fea8-496b-88b5-90a1ed1b6430](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:76c8df70-fea8-496b-88b5-90a1ed1b6430)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:c34b7125-ee54-460a-87aa-e542d047aa50](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:c34b7125-ee54-460a-87aa-e542d047aa50)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:1993a294-196e-495b-948a-eec015d895d4](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:1993a294-196e-495b-948a-eec015d895d4)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:71963e8a-b738-48e7-9fed-2191b904dc44](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:71963e8a-b738-48e7-9fed-2191b904dc44)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:67103076-a396-45bc-ad7b-f58065e5e002](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:67103076-a396-45bc-ad7b-f58065e5e002)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:6a86f9b0-b00f-4fd9-b117-48304c543964](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:6a86f9b0-b00f-4fd9-b117-48304c543964)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:a0e4e826-71e4-40f6-a322-f9fd7f4521e4](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:a0e4e826-71e4-40f6-a322-f9fd7f4521e4)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:da4ddcd9-5f05-4599-a064-bcf26ad2a195](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:da4ddcd9-5f05-4599-a064-bcf26ad2a195)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:e7537033-063f-4394-add9-41fed7b2c736](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:e7537033-063f-4394-add9-41fed7b2c736)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:30187daf-064b-4c00-a00d-ffad86e8eb8c](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:30187daf-064b-4c00-a00d-ffad86e8eb8c)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:02e5916d-b80f-4847-863d-1e888a968a92](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:02e5916d-b80f-4847-863d-1e888a968a92)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:fe1d79a5-c1d5-4d60-8324-573326e3a65d](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:fe1d79a5-c1d5-4d60-8324-573326e3a65d)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:fd806e4d-33ca-4933-bed7-3e79657ce80f](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:fd806e4d-33ca-4933-bed7-3e79657ce80f)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:f3a29112-f559-4c68-b5af-ec9821732c44](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:f3a29112-f559-4c68-b5af-ec9821732c44)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:1ceebe5a-6583-4062-8906-0ae87d2cb9b2](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:1ceebe5a-6583-4062-8906-0ae87d2cb9b2)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:a3c90d3d-524b-40c6-839a-22b53792dbc2](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:a3c90d3d-524b-40c6-839a-22b53792dbc2)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:20e796a3-a53f-4f09-a27c-8c1c38d94a9a](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:20e796a3-a53f-4f09-a27c-8c1c38d94a9a)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:8f4823ca-1b5c-4625-92b0-45c852a0ded8](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:8f4823ca-1b5c-4625-92b0-45c852a0ded8)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:7f639423-63e1-4e1a-bcc9-477eaf234501](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:7f639423-63e1-4e1a-bcc9-477eaf234501)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:09b6a975-f510-4ef5-97bd-ac5d00d9c01b](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:09b6a975-f510-4ef5-97bd-ac5d00d9c01b)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:7575d6d5-4a32-4a99-8f72-5715c005e23d](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:7575d6d5-4a32-4a99-8f72-5715c005e23d)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:0b3acbf0-c65c-4749-b86b-6d5fdc196b6b](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:0b3acbf0-c65c-4749-b86b-6d5fdc196b6b)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:4c97c536-79bd-433c-8ab2-3fcaac888bdb](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:4c97c536-79bd-433c-8ab2-3fcaac888bdb)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:d0dc9008-e514-40aa-bcb4-761b79c28707](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:d0dc9008-e514-40aa-bcb4-761b79c28707)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:92c65cc5-bd77-4182-92b3-4d521ae1dfe2](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:92c65cc5-bd77-4182-92b3-4d521ae1dfe2)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:79b01644-7cb5-4162-8971-a6b03626ad97](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:79b01644-7cb5-4162-8971-a6b03626ad97)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:8cae331e-5901-4fb3-9ce9-a122f0039a92](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:8cae331e-5901-4fb3-9ce9-a122f0039a92)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:6e0784ed-5028-44f2-9207-a861bf6d0500](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:6e0784ed-5028-44f2-9207-a861bf6d0500)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:6d6bd534-9782-4e5d-9ad0-268c6abce924](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:6d6bd534-9782-4e5d-9ad0-268c6abce924)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:a5680a3a-6953-4a34-a8ea-56ca3f6e99cd](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:a5680a3a-6953-4a34-a8ea-56ca3f6e99cd)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:23ac0ab2-9c03-43ed-9d0c-0688f84f7f00](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:23ac0ab2-9c03-43ed-9d0c-0688f84f7f00)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:1e642558-761a-4155-abf7-fc47f27c47b5](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:1e642558-761a-4155-abf7-fc47f27c47b5)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:8a0e52f2-3cb2-4389-b6ce-501e0169da5f](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:8a0e52f2-3cb2-4389-b6ce-501e0169da5f)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product:374535f1-5e66-4b22-9e13-ff89c097413c](https://tmf.dome-marketplace-sbx.org/tmf-api/productInventory/v4/product/urn:ngsi-ld:product:374535f1-5e66-4b22-9e13-ff89c097413c)

- **Type:** product
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:43 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to product (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

### productOffering Objects

#### Object: [urn:ngsi-ld:product-offering:005974a1-f327-47bd-96fc-2c263f2818c4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:005974a1-f327-47bd-96fc-2c263f2818c4)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:4174966c-aca0-4787-88fe-96dd235fa2df](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:4174966c-aca0-4787-88fe-96dd235fa2df)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:17f86010-d4c4-4120-99f5-25d25f376bbb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:17f86010-d4c4-4120-99f5-25d25f376bbb)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:9bbc3d54-daae-414e-a63a-52fa06dfcd0f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:9bbc3d54-daae-414e-a63a-52fa06dfcd0f)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:dcdd91b1-22ee-41e2-968b-24289c077e18](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:dcdd91b1-22ee-41e2-968b-24289c077e18)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:7bbf2620-43fe-4afe-b2ff-cfe83f78484a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:7bbf2620-43fe-4afe-b2ff-cfe83f78484a)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:1c9bdbba-e2bf-4b24-8586-9d01f8900cf1](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:1c9bdbba-e2bf-4b24-8586-9d01f8900cf1)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:66b48359-1d47-42af-a974-1c40ee50f3dd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:66b48359-1d47-42af-a974-1c40ee50f3dd)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:54a9af6f-f353-48a5-b599-84a535f0bc74](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:54a9af6f-f353-48a5-b599-84a535f0bc74)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:4699bdc1-592a-42b0-ad2d-93b0a06b044a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:4699bdc1-592a-42b0-ad2d-93b0a06b044a)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:2135b6c4-6de8-4ec9-9790-a3a322650a8a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:2135b6c4-6de8-4ec9-9790-a3a322650a8a)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:fd2bd8d2-f504-4254-9d2b-f248534d4351](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:fd2bd8d2-f504-4254-9d2b-f248534d4351)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:887a7ba9-2626-4f88-9252-5bc20239f61a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:887a7ba9-2626-4f88-9252-5bc20239f61a)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:98f94a51-6c8b-4778-a39b-10f3428428b7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:98f94a51-6c8b-4778-a39b-10f3428428b7)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:f85cc2c5-36e3-40b9-bf3d-4de468571c83](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:f85cc2c5-36e3-40b9-bf3d-4de468571c83)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:046b7413-46e5-404d-b6f1-e8d7f0db20e1](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:046b7413-46e5-404d-b6f1-e8d7f0db20e1)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)

#### Object: [urn:ngsi-ld:product-offering:71d584b8-486b-4f47-ba95-e328474a8236](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:71d584b8-486b-4f47-ba95-e328474a8236)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:af3e6eb5-6fac-4e6c-8675-75807b938d88](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:af3e6eb5-6fac-4e6c-8675-75807b938d88)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:8683c5f5-220c-4341-81df-dd29be9acd78](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:8683c5f5-220c-4341-81df-dd29be9acd78)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:686f78fc-4b21-4c4c-9f1f-0cb8ee9fd1c9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:686f78fc-4b21-4c4c-9f1f-0cb8ee9fd1c9)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:dcb09973-7df6-42fd-86c4-6b5d79e00374](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:dcb09973-7df6-42fd-86c4-6b5d79e00374)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:a1f60a8c-010c-42f6-a0fd-f5e1884a8212](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:a1f60a8c-010c-42f6-a0fd-f5e1884a8212)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:e7dacb3f-7332-428f-8d1a-59a86ffc5985](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:e7dacb3f-7332-428f-8d1a-59a86ffc5985)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:7f4b87dc-ef28-4579-99ad-6680d48947c3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:7f4b87dc-ef28-4579-99ad-6680d48947c3)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:87a8fa81-4b13-446b-9bdb-abfdc9e20dff](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:87a8fa81-4b13-446b-9bdb-abfdc9e20dff)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:75400f3b-452f-4b76-b98e-134179eb56bd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:75400f3b-452f-4b76-b98e-134179eb56bd)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:19121875-d09e-4de0-bf2d-5f07c8b65d2f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:19121875-d09e-4de0-bf2d-5f07c8b65d2f)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:9c8d762d-691f-4a91-a8ad-e5d479353756](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:9c8d762d-691f-4a91-a8ad-e5d479353756)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:d61bda3f-7b40-4417-b140-057178b2c6ae](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:d61bda3f-7b40-4417-b140-057178b2c6ae)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:b60af935-1059-41cf-a789-e8593500eaec](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:b60af935-1059-41cf-a789-e8593500eaec)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:b51ce551-5844-4f1d-8f73-b420b799347d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:b51ce551-5844-4f1d-8f73-b420b799347d)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:f1646ae0-a1c6-4655-b24a-681b1114f88b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:f1646ae0-a1c6-4655-b24a-681b1114f88b)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:620c434b-f829-44f3-9161-21c97c733225](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:620c434b-f829-44f3-9161-21c97c733225)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:208c28e6-23af-4f00-8608-bafb5db10ce7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:208c28e6-23af-4f00-8608-bafb5db10ce7)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:bd9bb9be-79b6-44e0-a3f7-b4e1eb1400ac](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:bd9bb9be-79b6-44e0-a3f7-b4e1eb1400ac)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:0de55201-a30a-4ba8-b92c-5729037d5773](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:0de55201-a30a-4ba8-b92c-5729037d5773)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:6caa329a-713f-4816-aab0-daad3536fe7a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:6caa329a-713f-4816-aab0-daad3536fe7a)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:afd19732-3509-4e52-9fe5-1bd47ab311d8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:afd19732-3509-4e52-9fe5-1bd47ab311d8)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:7c38669c-9b54-4e03-a0ac-862e8d6850ca](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:7c38669c-9b54-4e03-a0ac-862e8d6850ca)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:f5353acd-ee1d-41b6-80d6-d3eb6ac86640](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:f5353acd-ee1d-41b6-80d6-d3eb6ac86640)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:055a8b41-f80e-4835-80f1-b8cb10d27109](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:055a8b41-f80e-4835-80f1-b8cb10d27109)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:abadbea7-1c87-46b4-b391-9110ccd7519c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:abadbea7-1c87-46b4-b391-9110ccd7519c)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:7b2fc7ad-1835-4688-9e8e-c88fc47ec179](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:7b2fc7ad-1835-4688-9e8e-c88fc47ec179)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:6b357653-81fe-4cfc-8423-19d6752822b0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:6b357653-81fe-4cfc-8423-19d6752822b0)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:41bab3d8-d030-41fb-aa74-a9ce18d80899](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:41bab3d8-d030-41fb-aa74-a9ce18d80899)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:5b9eecc9-6005-4ac2-a334-c2351e89f001](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:5b9eecc9-6005-4ac2-a334-c2351e89f001)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:fc6d2f30-ec21-4d77-a1bd-d705dae868d8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:fc6d2f30-ec21-4d77-a1bd-d705dae868d8)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:30e2552b-0773-4453-8809-d6d573d2f88f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:30e2552b-0773-4453-8809-d6d573d2f88f)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:4c2f6cc2-0a19-49f3-a949-5438a7193429](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:4c2f6cc2-0a19-49f3-a949-5438a7193429)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:fe3c1b4b-655e-4b0e-b854-2241636a4682](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:fe3c1b4b-655e-4b0e-b854-2241636a4682)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:97488a56-6158-404e-ba9a-4975b7384924](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:97488a56-6158-404e-ba9a-4975b7384924)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:ac394906-585e-41c2-a0f8-21ef00454eb5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:ac394906-585e-41c2-a0f8-21ef00454eb5)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:e6429790-8432-4bb1-9e2b-e6c010f7ca39](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:e6429790-8432-4bb1-9e2b-e6c010f7ca39)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:5cac5209-a67f-4f98-bfea-d4444569ab47](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:5cac5209-a67f-4f98-bfea-d4444569ab47)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:d554fb90-a2ab-4521-b86c-66e19d798532](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:d554fb90-a2ab-4521-b86c-66e19d798532)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:409bfcc3-9af1-41c1-9834-8aa35280819d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:409bfcc3-9af1-41c1-9834-8aa35280819d)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:ebc0dc5a-4d40-4d1b-8c38-ec3b7b3c30ab](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:ebc0dc5a-4d40-4d1b-8c38-ec3b7b3c30ab)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:1a4ec325-fc21-4385-886f-25da9c978f63](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:1a4ec325-fc21-4385-886f-25da9c978f63)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:e70c66c4-3581-46da-be85-1d655b5ce31c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:e70c66c4-3581-46da-be85-1d655b5ce31c)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:c1790fda-d88a-4303-bc4b-fb99ea5df960](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:c1790fda-d88a-4303-bc4b-fb99ea5df960)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering:b84d9559-b141-4d2a-804b-04569dd3e974](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOffering/urn:ngsi-ld:product-offering:b84d9559-b141-4d2a-804b-04569dd3e974)

- **Type:** productOffering
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:55 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOffering (Code: MISSING_RECOMMENDED_FIELD)

### productOfferingPrice Objects

#### Object: [urn:ngsi-ld:product-offering-price:24934d22-46b7-44b7-a250-c74e8ea443d1](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:24934d22-46b7-44b7-a250-c74e8ea443d1)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d989cd04-4d33-41a9-a5b2-f5e0839db9f9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d989cd04-4d33-41a9-a5b2-f5e0839db9f9)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:43afdc36-6981-4019-8f58-d69f56f40f6e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:43afdc36-6981-4019-8f58-d69f56f40f6e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5cce5f34-9a42-4a9d-991e-b72845ded1a9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5cce5f34-9a42-4a9d-991e-b72845ded1a9)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:16a1dc71-7e91-40d2-b07e-d8e5e275cf8a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:16a1dc71-7e91-40d2-b07e-d8e5e275cf8a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:a0f8fe96-f382-4f56-aae6-6d90cec6797e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:a0f8fe96-f382-4f56-aae6-6d90cec6797e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:a824e4b5-b235-434b-acc5-5db0199432f3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:a824e4b5-b235-434b-acc5-5db0199432f3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b3eb4c12-803f-4fec-8925-9a333d6f587a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b3eb4c12-803f-4fec-8925-9a333d6f587a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:fa293088-a161-4640-9f5d-869072853b84](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:fa293088-a161-4640-9f5d-869072853b84)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)

#### Object: [urn:ngsi-ld:product-offering-price:3ae20fb5-4c93-4e28-8803-fb08cdb495c6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3ae20fb5-4c93-4e28-8803-fb08cdb495c6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:1f66f58d-6a2b-4014-aa89-be103f992dea](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:1f66f58d-6a2b-4014-aa89-be103f992dea)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:dc88b199-207f-4f21-8f92-e16470a10b02](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:dc88b199-207f-4f21-8f92-e16470a10b02)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:fba7585e-2fbb-4c29-a54d-902340ad158a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:fba7585e-2fbb-4c29-a54d-902340ad158a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3a0fbcbb-d241-4692-8a4a-045f8c9a459c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3a0fbcbb-d241-4692-8a4a-045f8c9a459c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4bffd57e-6868-44d2-a750-15eb3ef19a94](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4bffd57e-6868-44d2-a750-15eb3ef19a94)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:61d36718-e1af-40af-9dd4-d46ba5f39ad4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:61d36718-e1af-40af-9dd4-d46ba5f39ad4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:9f9a4fbe-b6cc-4cdf-be9e-418fa25ff7f6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:9f9a4fbe-b6cc-4cdf-be9e-418fa25ff7f6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6e13c06d-c2a4-49c9-99f5-5d5b7fb2c3fd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6e13c06d-c2a4-49c9-99f5-5d5b7fb2c3fd)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:50124c9d-4cd8-4a4d-ba7e-2a49dd08e951](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:50124c9d-4cd8-4a4d-ba7e-2a49dd08e951)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d4cc9f2f-73d5-417d-9a54-2da4a45281cd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d4cc9f2f-73d5-417d-9a54-2da4a45281cd)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4f74eb07-202c-40f9-bbb8-c7d276e71df0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4f74eb07-202c-40f9-bbb8-c7d276e71df0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:0fe122dd-0dbd-4502-8cc9-de9d50359fd0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:0fe122dd-0dbd-4502-8cc9-de9d50359fd0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:bae86ed8-38d4-4cd4-98fa-193d47a7c39d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:bae86ed8-38d4-4cd4-98fa-193d47a7c39d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5764025c-0b6a-4634-91fc-54e139397ccd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5764025c-0b6a-4634-91fc-54e139397ccd)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:2d533789-dca5-4158-bb5b-2bf48c61c738](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:2d533789-dca5-4158-bb5b-2bf48c61c738)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c50488b4-bd60-40ad-af06-837d76065d28](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c50488b4-bd60-40ad-af06-837d76065d28)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3da6a9ac-6d05-4225-936c-95ff20d6314a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3da6a9ac-6d05-4225-936c-95ff20d6314a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d56ef6d2-5450-4699-b487-f62582c4e005](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d56ef6d2-5450-4699-b487-f62582c4e005)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:a935eb24-c8a2-4b85-844c-8d72c89f4448](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:a935eb24-c8a2-4b85-844c-8d72c89f4448)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:fd1c14f0-6a1f-4119-88f1-ae2b16a31b27](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:fd1c14f0-6a1f-4119-88f1-ae2b16a31b27)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7159ac41-6541-41af-9bde-7a8f1f4b6504](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7159ac41-6541-41af-9bde-7a8f1f4b6504)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:f3b33b3e-d6de-4ff0-a8f4-8c049a4c4b83](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:f3b33b3e-d6de-4ff0-a8f4-8c049a4c4b83)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:af5f1f67-c4ab-4ddf-bf42-f9ace2975b05](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:af5f1f67-c4ab-4ddf-bf42-f9ace2975b05)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:21de9359-96c9-4274-bfde-2351324ed426](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:21de9359-96c9-4274-bfde-2351324ed426)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6523ecb7-b6c4-47ef-ab86-1f827c9e99d0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6523ecb7-b6c4-47ef-ab86-1f827c9e99d0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ac33232f-d8f5-4388-a705-b4fbbdc96c48](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ac33232f-d8f5-4388-a705-b4fbbdc96c48)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:889ae01f-a833-44ce-aa08-ecc1b0158931](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:889ae01f-a833-44ce-aa08-ecc1b0158931)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b6e9523f-f259-4574-aa2f-1630b120a171](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b6e9523f-f259-4574-aa2f-1630b120a171)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:646a7334-2ae2-4182-8ae6-4e24f8ddb21d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:646a7334-2ae2-4182-8ae6-4e24f8ddb21d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:258cbab9-c6b8-4435-8bfe-0288ead01108](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:258cbab9-c6b8-4435-8bfe-0288ead01108)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c4d80434-af55-4483-916f-7df7be8eac3c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c4d80434-af55-4483-916f-7df7be8eac3c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:49648693-b36b-4eb2-954d-75438ed8fff8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:49648693-b36b-4eb2-954d-75438ed8fff8)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6ccf2ee9-6ae2-4019-b5c9-8d7b64025f98](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6ccf2ee9-6ae2-4019-b5c9-8d7b64025f98)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:012db45a-5e82-4738-aff0-ff3d065ee194](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:012db45a-5e82-4738-aff0-ff3d065ee194)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ddd90524-e862-49af-b01e-22e906a18e27](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ddd90524-e862-49af-b01e-22e906a18e27)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:21e2b57f-1512-4053-b08e-c3d309580c24](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:21e2b57f-1512-4053-b08e-c3d309580c24)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:86f7d405-fada-4bf3-97c3-1afa8b671d8e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:86f7d405-fada-4bf3-97c3-1afa8b671d8e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d46a935e-ff10-4b16-95c2-4c00897c1c89](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d46a935e-ff10-4b16-95c2-4c00897c1c89)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7c337e51-20da-4083-8241-28d5be7a5887](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7c337e51-20da-4083-8241-28d5be7a5887)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:775bc3c7-e63d-4ee8-8c90-d0d0c18dadc7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:775bc3c7-e63d-4ee8-8c90-d0d0c18dadc7)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d65d6d85-bb29-4d10-88b0-bcfdb98db24a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d65d6d85-bb29-4d10-88b0-bcfdb98db24a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:19a4f8a2-fc72-4cf0-90b8-b272eaf1a098](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:19a4f8a2-fc72-4cf0-90b8-b272eaf1a098)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:1c049cbd-4307-447f-8f12-277f43f14d34](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:1c049cbd-4307-447f-8f12-277f43f14d34)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6584eca7-5d04-4234-8e27-1754286b7cf4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6584eca7-5d04-4234-8e27-1754286b7cf4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:75773dfb-3159-4463-858d-2f46be394770](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:75773dfb-3159-4463-858d-2f46be394770)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7e103d66-34d8-42b9-bb2f-ed7d581756b2](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7e103d66-34d8-42b9-bb2f-ed7d581756b2)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5df8bd48-3dd9-4dbf-a343-ab6f6e10617f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5df8bd48-3dd9-4dbf-a343-ab6f6e10617f)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:dddd5496-be8a-4d7d-8c0f-2d1be7f2bd84](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:dddd5496-be8a-4d7d-8c0f-2d1be7f2bd84)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c1948bd6-2559-4146-ad61-681a83068283](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c1948bd6-2559-4146-ad61-681a83068283)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6055443d-785b-4644-bbda-4dd25c4108bb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6055443d-785b-4644-bbda-4dd25c4108bb)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:efa9c9a1-3a7e-4e87-ae8d-fea9ea2ba5e0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:efa9c9a1-3a7e-4e87-ae8d-fea9ea2ba5e0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b4aa8615-1743-4260-b334-aec163ce799d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b4aa8615-1743-4260-b334-aec163ce799d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:8d9338c8-6a43-4ae4-a57f-96deaba612db](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:8d9338c8-6a43-4ae4-a57f-96deaba612db)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ce65f509-1d37-4f93-bde1-f897df1cc250](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ce65f509-1d37-4f93-bde1-f897df1cc250)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6eed089a-0fcb-4e07-8443-87500cc4086b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6eed089a-0fcb-4e07-8443-87500cc4086b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:27c7220b-33b4-4f9a-b540-75a17d4a321a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:27c7220b-33b4-4f9a-b540-75a17d4a321a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:aa5dddd0-44bc-4edd-a07c-8a7e8a5c8219](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:aa5dddd0-44bc-4edd-a07c-8a7e8a5c8219)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4a6382b9-fc5f-4b72-85ef-853f20bc941a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4a6382b9-fc5f-4b72-85ef-853f20bc941a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3e563b0c-cc92-422b-a530-87927f5342d6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3e563b0c-cc92-422b-a530-87927f5342d6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c798d017-8e6e-40be-9f09-ce2ae7f28ba6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c798d017-8e6e-40be-9f09-ce2ae7f28ba6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5bfe47e3-2466-4381-a22c-da072b5a333c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5bfe47e3-2466-4381-a22c-da072b5a333c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:80d5fe58-79a7-4691-a011-5d6d7efd0681](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:80d5fe58-79a7-4691-a011-5d6d7efd0681)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4a65a9f9-037d-4694-9bcc-1206487485ff](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4a65a9f9-037d-4694-9bcc-1206487485ff)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5456d2c3-bd4d-4283-b49d-b5286ed5dacf](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5456d2c3-bd4d-4283-b49d-b5286ed5dacf)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c40d71ec-4a8e-4ecf-a1dc-50ed537cb915](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c40d71ec-4a8e-4ecf-a1dc-50ed537cb915)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:08cc66ce-1df6-4a70-98fc-10f94336a732](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:08cc66ce-1df6-4a70-98fc-10f94336a732)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:0ddc0dd1-467d-4292-8325-bcd727db0ce4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:0ddc0dd1-467d-4292-8325-bcd727db0ce4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3bb60c49-65de-4a77-bdbe-4a6ded9637f5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3bb60c49-65de-4a77-bdbe-4a6ded9637f5)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:df04f82d-08d6-44ab-a8c5-65a3108b0b6b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:df04f82d-08d6-44ab-a8c5-65a3108b0b6b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b7663735-496f-42d0-ba20-e792e14eacd4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b7663735-496f-42d0-ba20-e792e14eacd4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:acd98be5-13b5-4ddf-9dc6-d6331bcab0b8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:acd98be5-13b5-4ddf-9dc6-d6331bcab0b8)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:37b0a2bc-daa5-4f04-9b0d-5eb1ee2c7034](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:37b0a2bc-daa5-4f04-9b0d-5eb1ee2c7034)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7b534ee7-0e93-4f33-b34c-0c6502c382b7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7b534ee7-0e93-4f33-b34c-0c6502c382b7)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:9b27d33c-0856-400c-ba91-36ca0669c0c7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:9b27d33c-0856-400c-ba91-36ca0669c0c7)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:9c17d6a9-264b-4245-a0dd-96b7b791d41f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:9c17d6a9-264b-4245-a0dd-96b7b791d41f)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:13a1d807-012c-4e7a-bd46-1ed1422e057d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:13a1d807-012c-4e7a-bd46-1ed1422e057d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c4b5bbf8-cac3-4477-b6a5-9a5cb50d5210](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c4b5bbf8-cac3-4477-b6a5-9a5cb50d5210)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:0b2dcaa4-2c1d-4d50-9bca-ca69e79facf0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:0b2dcaa4-2c1d-4d50-9bca-ca69e79facf0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7b60dfbc-dffe-40da-8c62-3dc410402a51](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7b60dfbc-dffe-40da-8c62-3dc410402a51)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:f346907c-e792-4ca0-830d-21772439bfaa](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:f346907c-e792-4ca0-830d-21772439bfaa)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c0642c81-b162-48f1-9d08-42cffe7769ce](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c0642c81-b162-48f1-9d08-42cffe7769ce)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:f026d149-b29e-4d15-b710-5cc785808c98](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:f026d149-b29e-4d15-b710-5cc785808c98)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:33358d97-bc25-45ce-96ed-6aabd2ed6fa0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:33358d97-bc25-45ce-96ed-6aabd2ed6fa0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:9541b92a-8422-43b8-8bc3-2191a80c4db0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:9541b92a-8422-43b8-8bc3-2191a80c4db0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:e2416ee0-7778-4708-b905-5146d5376c8e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:e2416ee0-7778-4708-b905-5146d5376c8e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:30d168c5-9537-490d-8dfa-eb352d4d8c9b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:30d168c5-9537-490d-8dfa-eb352d4d8c9b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:96b1e4cb-b8b3-4fa6-acd7-9e1f47794ba0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:96b1e4cb-b8b3-4fa6-acd7-9e1f47794ba0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6f4fd197-816f-4ab8-988d-1e89c866d445](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6f4fd197-816f-4ab8-988d-1e89c866d445)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:63578d82-7280-4b2d-9e85-6379f7b17d58](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:63578d82-7280-4b2d-9e85-6379f7b17d58)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:653a42f3-a031-400d-88eb-a56fa03015d8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:653a42f3-a031-400d-88eb-a56fa03015d8)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:13 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:dcb20903-1d01-4b74-8c0e-46a2f2b920dd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:dcb20903-1d01-4b74-8c0e-46a2f2b920dd)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:acbd99be-2f4e-4dea-9a81-6cb797d4e1ae](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:acbd99be-2f4e-4dea-9a81-6cb797d4e1ae)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b55cc65d-2c12-436c-ba02-d2c64687bdd8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b55cc65d-2c12-436c-ba02-d2c64687bdd8)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c0263817-5961-4164-b8f4-95042817975f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c0263817-5961-4164-b8f4-95042817975f)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d4eb9f12-39f3-4f20-b2ec-8b1ec9fd288f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d4eb9f12-39f3-4f20-b2ec-8b1ec9fd288f)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:a6760111-3b5b-4b03-a2b3-1af9cde04cf0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:a6760111-3b5b-4b03-a2b3-1af9cde04cf0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:00c8b914-9b16-4b2f-8558-021ed2938c8e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:00c8b914-9b16-4b2f-8558-021ed2938c8e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:9f8c70a3-dde9-4925-b1f4-66baacaf5c91](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:9f8c70a3-dde9-4925-b1f4-66baacaf5c91)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:fbaa4ff7-ad07-4151-a66e-5404df28e77d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:fbaa4ff7-ad07-4151-a66e-5404df28e77d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ed158cbe-5c29-4fa9-85c5-9171b85336cb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ed158cbe-5c29-4fa9-85c5-9171b85336cb)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:588267d9-c68d-4754-bca7-1169e97e0df6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:588267d9-c68d-4754-bca7-1169e97e0df6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:9c9ae134-4f97-4b9c-9fd2-b3fd8e96da5f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:9c9ae134-4f97-4b9c-9fd2-b3fd8e96da5f)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:1a8bb8de-96a6-4ab3-a1fb-4153aa6ab44f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:1a8bb8de-96a6-4ab3-a1fb-4153aa6ab44f)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7cb21085-01c0-4b51-9c4b-f3ff6f510eba](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7cb21085-01c0-4b51-9c4b-f3ff6f510eba)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:499e5825-32f8-41a7-9f49-e5dba28094ac](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:499e5825-32f8-41a7-9f49-e5dba28094ac)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:a7fcc974-0754-43ac-a3ca-b9d083d81649](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:a7fcc974-0754-43ac-a3ca-b9d083d81649)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5a4968ff-053c-4040-a28c-14871c1cf8e6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5a4968ff-053c-4040-a28c-14871c1cf8e6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:f4513107-2697-4e21-9d8f-89d8cacd4a55](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:f4513107-2697-4e21-9d8f-89d8cacd4a55)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:de400ec1-6932-42c8-806a-a789618054bd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:de400ec1-6932-42c8-806a-a789618054bd)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:2c064d03-5860-4cd2-8607-9a85450a5826](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:2c064d03-5860-4cd2-8607-9a85450a5826)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b206f02d-50cd-4100-b722-bbc7db0daa35](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b206f02d-50cd-4100-b722-bbc7db0daa35)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3a29cc75-5792-4222-a7a2-38acc237658d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3a29cc75-5792-4222-a7a2-38acc237658d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d671dce4-1960-4470-b994-fe35cca4efa5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d671dce4-1960-4470-b994-fe35cca4efa5)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3f8d055d-cec8-47f5-a116-dfaa90c08598](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3f8d055d-cec8-47f5-a116-dfaa90c08598)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:010ed217-8db2-404b-8aa8-6e3b9cd60f61](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:010ed217-8db2-404b-8aa8-6e3b9cd60f61)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ea94085a-b883-4ec3-a3ed-005a4afb333e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ea94085a-b883-4ec3-a3ed-005a4afb333e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6862a634-b40d-4f9e-8719-1b57513b954f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6862a634-b40d-4f9e-8719-1b57513b954f)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:f0aac682-a141-4f55-8fb4-8502d6da40ce](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:f0aac682-a141-4f55-8fb4-8502d6da40ce)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:beae51b7-9c13-45ec-ad04-f750f5b1eefb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:beae51b7-9c13-45ec-ad04-f750f5b1eefb)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:e11d0edf-7ec3-4bd9-aab0-3942931f5e63](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:e11d0edf-7ec3-4bd9-aab0-3942931f5e63)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:e4357db7-2339-4051-afc1-a4e4e95ac3db](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:e4357db7-2339-4051-afc1-a4e4e95ac3db)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:abefd473-1f0a-4244-9a69-c0568504ab7b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:abefd473-1f0a-4244-9a69-c0568504ab7b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:00eb3279-aa7b-414c-9201-8fb455342e46](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:00eb3279-aa7b-414c-9201-8fb455342e46)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b54bc88e-8104-44e7-a7ec-051881852093](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b54bc88e-8104-44e7-a7ec-051881852093)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:9c957909-3a01-473e-a7e0-cbd873ebfb2a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:9c957909-3a01-473e-a7e0-cbd873ebfb2a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:229eead4-7103-400d-989f-ab5a7fb4b3d3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:229eead4-7103-400d-989f-ab5a7fb4b3d3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5cf12673-a33f-4516-ad69-70a102b50bf0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5cf12673-a33f-4516-ad69-70a102b50bf0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:75534c79-5493-47ce-99e9-a06125a61322](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:75534c79-5493-47ce-99e9-a06125a61322)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:05099d5c-29f3-4929-baa3-505df5ca68e0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:05099d5c-29f3-4929-baa3-505df5ca68e0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:143b311e-bfa2-4612-8974-75b29c991d45](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:143b311e-bfa2-4612-8974-75b29c991d45)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:bee90d28-fc11-4bd1-af98-ba8757b532e4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:bee90d28-fc11-4bd1-af98-ba8757b532e4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:62085d6a-ab60-4d4e-8514-5120f69ac3b6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:62085d6a-ab60-4d4e-8514-5120f69ac3b6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:036cd75a-908a-42ef-830b-14d18dc65f87](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:036cd75a-908a-42ef-830b-14d18dc65f87)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:2bf6fc82-6423-429c-a963-94e137e55e65](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:2bf6fc82-6423-429c-a963-94e137e55e65)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:dccf3108-14bc-4e7e-bd26-27b3c9f4abdc](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:dccf3108-14bc-4e7e-bd26-27b3c9f4abdc)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:75759183-6a6f-43b0-a2d0-d5be4f4aa3d2](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:75759183-6a6f-43b0-a2d0-d5be4f4aa3d2)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:853011f4-ca90-41e3-95c9-1141836723eb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:853011f4-ca90-41e3-95c9-1141836723eb)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:8c8a046b-15a8-4a8c-a616-f186fab53714](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:8c8a046b-15a8-4a8c-a616-f186fab53714)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:f65e274c-76bc-4d52-a090-d0c0b13ed15b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:f65e274c-76bc-4d52-a090-d0c0b13ed15b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4dd1b451-8366-4be9-b473-568f1ff704bd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4dd1b451-8366-4be9-b473-568f1ff704bd)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6baf7b40-a203-4aa1-8c52-e19d34ebae55](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6baf7b40-a203-4aa1-8c52-e19d34ebae55)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:2f4f669e-8365-4f7f-b6c4-547c0cd58924](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:2f4f669e-8365-4f7f-b6c4-547c0cd58924)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7a45ace3-f8d4-4c28-9c9b-7191a5fe28ed](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7a45ace3-f8d4-4c28-9c9b-7191a5fe28ed)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:1f6b7299-89e1-42aa-80fa-282aba42504c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:1f6b7299-89e1-42aa-80fa-282aba42504c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4bc35f20-74dd-4f2e-9fd5-839bcf779790](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4bc35f20-74dd-4f2e-9fd5-839bcf779790)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:dc2edec2-ae18-413e-918d-eb0ff81dc5bd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:dc2edec2-ae18-413e-918d-eb0ff81dc5bd)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:91de3726-4d0d-4782-99a7-0a5c39d8619d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:91de3726-4d0d-4782-99a7-0a5c39d8619d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3587a720-0e0f-4569-9674-62a647e137de](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3587a720-0e0f-4569-9674-62a647e137de)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4e3181d7-22ac-45fe-998e-d8e89c0d2f3d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4e3181d7-22ac-45fe-998e-d8e89c0d2f3d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4e4d1da0-009a-4d76-9bad-444331a2197d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4e4d1da0-009a-4d76-9bad-444331a2197d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5174d7f1-8b7f-4fac-abb9-c2ba78f799f4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5174d7f1-8b7f-4fac-abb9-c2ba78f799f4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:12a5c9e7-99e1-4b68-8512-0665c6a42027](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:12a5c9e7-99e1-4b68-8512-0665c6a42027)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:a4d492e6-bfe1-4d0b-8aec-fea3233e846e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:a4d492e6-bfe1-4d0b-8aec-fea3233e846e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:70e027f1-931e-4854-92dc-ca2ec86bb3f8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:70e027f1-931e-4854-92dc-ca2ec86bb3f8)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:789a3618-a0fa-4736-a6b4-fc538a23df93](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:789a3618-a0fa-4736-a6b4-fc538a23df93)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5dbda5f9-1024-4123-bd45-b0919aa7420c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5dbda5f9-1024-4123-bd45-b0919aa7420c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:9351cabe-b828-4469-aa96-53eabcdd33e6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:9351cabe-b828-4469-aa96-53eabcdd33e6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b89eba0f-1b8f-4583-8617-9079a0ace4cf](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b89eba0f-1b8f-4583-8617-9079a0ace4cf)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:67d11a7e-85d1-4603-808c-374265c87df5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:67d11a7e-85d1-4603-808c-374265c87df5)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:028604b4-6a02-4151-b6b4-9a0d40819ec3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:028604b4-6a02-4151-b6b4-9a0d40819ec3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:9fc44f02-809c-4ad9-8c02-01739369f050](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:9fc44f02-809c-4ad9-8c02-01739369f050)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ae15b3b0-6650-4bfa-aacd-595ad247b604](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ae15b3b0-6650-4bfa-aacd-595ad247b604)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:cb1c522d-6992-47ee-bd3f-06310f0a2e04](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:cb1c522d-6992-47ee-bd3f-06310f0a2e04)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:14c6e0ab-14ba-4684-9592-e053dfaeb9f3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:14c6e0ab-14ba-4684-9592-e053dfaeb9f3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:8331f85c-2521-4a89-8ba5-1fc507badafa](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:8331f85c-2521-4a89-8ba5-1fc507badafa)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:89992021-0f73-416c-a553-7c9625950ca8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:89992021-0f73-416c-a553-7c9625950ca8)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:dda6aea9-4127-4591-b746-25ab03b8896e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:dda6aea9-4127-4591-b746-25ab03b8896e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:2e49cb26-de04-4ba8-b014-78d7c44b79bd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:2e49cb26-de04-4ba8-b014-78d7c44b79bd)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:58e60366-e54c-4bfb-b74a-d379d1b5e810](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:58e60366-e54c-4bfb-b74a-d379d1b5e810)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:fdb99f48-6b79-4e40-9a79-4eaf27dec313](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:fdb99f48-6b79-4e40-9a79-4eaf27dec313)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c469000b-40b7-4cf5-a343-9727427127f3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c469000b-40b7-4cf5-a343-9727427127f3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:64847008-03a7-49f6-8fc8-e8bb36c07d56](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:64847008-03a7-49f6-8fc8-e8bb36c07d56)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b2bcb36b-af5c-40d4-85b4-0fe414db6511](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b2bcb36b-af5c-40d4-85b4-0fe414db6511)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:8366f679-1b36-47b1-98d1-8ec7419c6f49](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:8366f679-1b36-47b1-98d1-8ec7419c6f49)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:73a3c6e8-5767-473f-b152-5520cd3475d2](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:73a3c6e8-5767-473f-b152-5520cd3475d2)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:cf7ba21c-c3ba-4260-a91c-a1188fdfe648](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:cf7ba21c-c3ba-4260-a91c-a1188fdfe648)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:831fadd7-86b2-49e8-b852-d04d97eb7cb2](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:831fadd7-86b2-49e8-b852-d04d97eb7cb2)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c4f98a8a-01ff-4637-82aa-b1414858230a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c4f98a8a-01ff-4637-82aa-b1414858230a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:36753374-ad79-4dee-98b4-e05d3986174d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:36753374-ad79-4dee-98b4-e05d3986174d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c33530a8-f56f-45cf-8387-2efcd5aec3f3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c33530a8-f56f-45cf-8387-2efcd5aec3f3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c2d6e61d-eb87-4942-9f44-6af9e8877e31](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c2d6e61d-eb87-4942-9f44-6af9e8877e31)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:1d431284-ad03-458a-aba1-d7a0950dc72f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:1d431284-ad03-458a-aba1-d7a0950dc72f)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:2364d6c1-21e8-436b-b54d-289d67f9c198](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:2364d6c1-21e8-436b-b54d-289d67f9c198)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3d904149-53b1-49eb-9bf3-ac2f4ae7ac0b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3d904149-53b1-49eb-9bf3-ac2f4ae7ac0b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:aa5580ce-dd29-4305-8560-9416d8ecd8e3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:aa5580ce-dd29-4305-8560-9416d8ecd8e3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:81ba5637-95b9-4c08-a21b-fa4238b05d68](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:81ba5637-95b9-4c08-a21b-fa4238b05d68)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4fe1bc96-cb11-4dfc-a523-abdfefd49fe8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4fe1bc96-cb11-4dfc-a523-abdfefd49fe8)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ea057aa3-dfa3-4d90-aaf8-f1cb8cdd3966](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ea057aa3-dfa3-4d90-aaf8-f1cb8cdd3966)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:46a9dd25-cba6-4b26-b7b3-1e75048697a6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:46a9dd25-cba6-4b26-b7b3-1e75048697a6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:36b33c3a-2923-48b0-ad9a-82ca7ab8790d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:36b33c3a-2923-48b0-ad9a-82ca7ab8790d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:24 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:30a13850-3ea2-4e5c-a79b-0d88543fef64](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:30a13850-3ea2-4e5c-a79b-0d88543fef64)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:77d25781-f15a-4864-965e-3206c79d4740](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:77d25781-f15a-4864-965e-3206c79d4740)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d9ebcd5c-7e5a-4fb5-8283-8a1a6f9326f9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d9ebcd5c-7e5a-4fb5-8283-8a1a6f9326f9)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d789ab42-e949-4c96-b287-a140f3528267](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d789ab42-e949-4c96-b287-a140f3528267)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b886312f-a488-440b-999e-3c1cb5dcc0b2](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b886312f-a488-440b-999e-3c1cb5dcc0b2)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:885084a0-900d-4956-9cc3-f6fd44d6b9d7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:885084a0-900d-4956-9cc3-f6fd44d6b9d7)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:66a73399-85cc-47b0-8675-301552ba725c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:66a73399-85cc-47b0-8675-301552ba725c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:8c96201e-dd51-41c9-af24-96d210a4f156](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:8c96201e-dd51-41c9-af24-96d210a4f156)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:a3b49d8f-ed46-448a-a159-4a451c2198ad](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:a3b49d8f-ed46-448a-a159-4a451c2198ad)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:8c7e0dc8-a03d-40b3-b2e6-7ba8ef1a4b5c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:8c7e0dc8-a03d-40b3-b2e6-7ba8ef1a4b5c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d2593254-0e54-4cfe-820f-9dc727aaa629](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d2593254-0e54-4cfe-820f-9dc727aaa629)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7d78b71d-262e-4a83-ac22-8e1768f5f4a4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7d78b71d-262e-4a83-ac22-8e1768f5f4a4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:011d51c8-4108-43c5-be2e-95c22cbb4ce8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:011d51c8-4108-43c5-be2e-95c22cbb4ce8)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:220f95c1-b20c-405e-8351-a34746330b42](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:220f95c1-b20c-405e-8351-a34746330b42)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3afd750c-1547-466a-ae2a-1cefec930ccf](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3afd750c-1547-466a-ae2a-1cefec930ccf)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:f074c366-8286-4000-a6bd-0796fdfd7fd4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:f074c366-8286-4000-a6bd-0796fdfd7fd4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:539b3e98-35ad-4b9d-a514-48f9618e5c5e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:539b3e98-35ad-4b9d-a514-48f9618e5c5e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4bce7b40-75d9-4d17-8c52-9711e823a2ef](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4bce7b40-75d9-4d17-8c52-9711e823a2ef)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:fc105ca2-9d7e-46fc-ae7b-0241f2bd2751](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:fc105ca2-9d7e-46fc-ae7b-0241f2bd2751)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:bb486136-cc13-4dc3-af66-23680fe32ac0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:bb486136-cc13-4dc3-af66-23680fe32ac0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:17c99d15-dd9e-4d1c-ae3a-8605ba01c1ef](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:17c99d15-dd9e-4d1c-ae3a-8605ba01c1ef)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d7d1003e-d1c8-4582-9b32-17f417ddcb25](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d7d1003e-d1c8-4582-9b32-17f417ddcb25)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c91597ef-6c7e-4fae-a128-6924d5890ec4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c91597ef-6c7e-4fae-a128-6924d5890ec4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b242ca11-4d68-4308-bbf8-47b1ca1d2526](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b242ca11-4d68-4308-bbf8-47b1ca1d2526)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6b51a48c-2731-421d-b3a7-05c00f67fe12](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6b51a48c-2731-421d-b3a7-05c00f67fe12)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b8c20bcd-b227-4983-b96f-d58ebc7284b6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b8c20bcd-b227-4983-b96f-d58ebc7284b6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7dab9686-4c52-4951-b420-d4aff9a19cb2](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7dab9686-4c52-4951-b420-d4aff9a19cb2)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:a3b3de69-2903-4ac9-8bd4-15be81cb0691](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:a3b3de69-2903-4ac9-8bd4-15be81cb0691)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:79f5acff-9d87-4476-8ee8-47768585a82f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:79f5acff-9d87-4476-8ee8-47768585a82f)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:226c388e-69e5-45d1-8171-baef2a29336c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:226c388e-69e5-45d1-8171-baef2a29336c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b9008c32-0aa8-4d2c-a64d-4230d01f1e9b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b9008c32-0aa8-4d2c-a64d-4230d01f1e9b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7116e4c4-3c97-4652-8824-62371a77e9ac](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7116e4c4-3c97-4652-8824-62371a77e9ac)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6b64a300-8c2c-464d-974e-28bdd06994b3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6b64a300-8c2c-464d-974e-28bdd06994b3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ca721fb0-507f-42d4-80a9-395337766d6c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ca721fb0-507f-42d4-80a9-395337766d6c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:74d0de69-f40c-4345-87bc-2c44f8ea286e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:74d0de69-f40c-4345-87bc-2c44f8ea286e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6d73cf20-dec9-49d9-9be5-31cfd868fe86](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6d73cf20-dec9-49d9-9be5-31cfd868fe86)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:bb1e78a5-a7b4-4162-884f-170f10e6cecb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:bb1e78a5-a7b4-4162-884f-170f10e6cecb)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:34d2613e-ed4f-47dd-ade7-593c048392c5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:34d2613e-ed4f-47dd-ade7-593c048392c5)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:bbfffe02-466a-4f9c-bee2-fedc07c39bb7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:bbfffe02-466a-4f9c-bee2-fedc07c39bb7)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:87fd14f7-cb6e-4bd2-9149-45ccae1c4a9c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:87fd14f7-cb6e-4bd2-9149-45ccae1c4a9c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:85684d70-d96d-4f6c-8e80-008ae5792611](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:85684d70-d96d-4f6c-8e80-008ae5792611)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:89acb00d-42bd-4582-a945-94d800a77889](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:89acb00d-42bd-4582-a945-94d800a77889)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:a19fb68d-7896-49f7-bb97-a10fd3d8c289](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:a19fb68d-7896-49f7-bb97-a10fd3d8c289)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b5299589-96f8-4540-ac7f-3316adf3fb41](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b5299589-96f8-4540-ac7f-3316adf3fb41)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d4568a25-a296-4154-8a88-394117e891b0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d4568a25-a296-4154-8a88-394117e891b0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:821f7740-e067-47ed-99b0-d8e0af013d46](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:821f7740-e067-47ed-99b0-d8e0af013d46)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:799962bc-f628-48ff-ae11-60a8cd6ab85b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:799962bc-f628-48ff-ae11-60a8cd6ab85b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:88428c8f-1e6c-48d8-98f5-9f7d07f23f42](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:88428c8f-1e6c-48d8-98f5-9f7d07f23f42)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:dd616652-c77b-49cb-92d9-dbbf3b022c1f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:dd616652-c77b-49cb-92d9-dbbf3b022c1f)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c8a0d287-b104-4dd4-bec1-e4c38108f042](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c8a0d287-b104-4dd4-bec1-e4c38108f042)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ee651a53-68bc-49d6-a00c-3505a785da63](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ee651a53-68bc-49d6-a00c-3505a785da63)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ba526e3b-e4ef-41af-9280-7deb3ae8d3ef](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ba526e3b-e4ef-41af-9280-7deb3ae8d3ef)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:2b6dfba2-f197-4ca2-848e-cbc8a9fec26c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:2b6dfba2-f197-4ca2-848e-cbc8a9fec26c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3560cd20-758b-4667-b17e-71e8b254cd26](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3560cd20-758b-4667-b17e-71e8b254cd26)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:da29fd02-e98a-4a2c-b745-d25546f5b2f3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:da29fd02-e98a-4a2c-b745-d25546f5b2f3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:dde62bae-5bb9-4a86-8a72-d045ac056745](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:dde62bae-5bb9-4a86-8a72-d045ac056745)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:debb531a-58cd-4967-a5e3-fb506e535f63](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:debb531a-58cd-4967-a5e3-fb506e535f63)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3bd80ac0-df29-422f-9514-3f391e2a5386](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3bd80ac0-df29-422f-9514-3f391e2a5386)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:32f94c7c-95d1-47e2-a901-93f251fa1183](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:32f94c7c-95d1-47e2-a901-93f251fa1183)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:fbaf08b9-3bc0-43d2-a6b0-2947de5ebb4a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:fbaf08b9-3bc0-43d2-a6b0-2947de5ebb4a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:eb095bb1-a0b8-4a1f-9ffb-cecc857f3494](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:eb095bb1-a0b8-4a1f-9ffb-cecc857f3494)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:56e098be-5276-4567-8735-045358d040e6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:56e098be-5276-4567-8735-045358d040e6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:192eb6b4-cc67-47ba-bc84-17083ce43c16](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:192eb6b4-cc67-47ba-bc84-17083ce43c16)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:41278603-c6b1-441a-98a3-762186d84f44](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:41278603-c6b1-441a-98a3-762186d84f44)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:06d54d38-6839-4c89-95e6-af204f1b9bd3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:06d54d38-6839-4c89-95e6-af204f1b9bd3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5f2f1480-8101-4621-a31d-84b51b016e2b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5f2f1480-8101-4621-a31d-84b51b016e2b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3282cc40-b169-4536-9e79-e2f6cc73e72e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3282cc40-b169-4536-9e79-e2f6cc73e72e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5feb54eb-33f8-43a3-aa2f-68f5bc8f2585](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5feb54eb-33f8-43a3-aa2f-68f5bc8f2585)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:2bdea539-1470-449e-a79e-4c9e8243497c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:2bdea539-1470-449e-a79e-4c9e8243497c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:469c81eb-b9d5-4347-8913-d1225505696a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:469c81eb-b9d5-4347-8913-d1225505696a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4033a990-8900-48a8-8b94-feaf092961e4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4033a990-8900-48a8-8b94-feaf092961e4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:f8363712-ecba-446d-b453-678dde3a2ef6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:f8363712-ecba-446d-b453-678dde3a2ef6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:2462c62c-4b5b-494a-8c32-1dcdfae17a7a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:2462c62c-4b5b-494a-8c32-1dcdfae17a7a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:43f872ae-9de5-4927-a119-139675a03236](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:43f872ae-9de5-4927-a119-139675a03236)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:96fd27d6-c0f4-4c98-8f17-e900b217fe46](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:96fd27d6-c0f4-4c98-8f17-e900b217fe46)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6fbeef02-49b5-4685-86d4-58bcac75d701](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6fbeef02-49b5-4685-86d4-58bcac75d701)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:de1ae86d-438f-406e-80c2-858fe200daaa](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:de1ae86d-438f-406e-80c2-858fe200daaa)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:8711a720-3648-485d-81f6-7c0bf81afb02](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:8711a720-3648-485d-81f6-7c0bf81afb02)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:98861972-fdf3-4cb6-a38f-9f851185168d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:98861972-fdf3-4cb6-a38f-9f851185168d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7976a0eb-60be-42a0-99f3-fa5efdb4be7a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7976a0eb-60be-42a0-99f3-fa5efdb4be7a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:62581ec3-cf51-47ca-8ba0-8f82b96d6868](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:62581ec3-cf51-47ca-8ba0-8f82b96d6868)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:fc4c0e5c-2751-48af-aa0b-71cc6dc03e6b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:fc4c0e5c-2751-48af-aa0b-71cc6dc03e6b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3e3b6d90-fe3f-46aa-b971-2ebdf4d3ce9a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3e3b6d90-fe3f-46aa-b971-2ebdf4d3ce9a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4fefddf8-a6a9-4d62-8e51-c655a67dbabd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4fefddf8-a6a9-4d62-8e51-c655a67dbabd)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:1f5f2891-5ff4-4751-b420-1068af6ac0ef](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:1f5f2891-5ff4-4751-b420-1068af6ac0ef)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:db22b038-e309-47ce-a38e-8abd4de4f4e5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:db22b038-e309-47ce-a38e-8abd4de4f4e5)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:e499e2c0-882e-4a10-94e0-84959f009904](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:e499e2c0-882e-4a10-94e0-84959f009904)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c9124c8c-ace8-42f1-a42c-2156dd2b506c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c9124c8c-ace8-42f1-a42c-2156dd2b506c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:76ad71ac-d40d-4e62-bfed-9faf4d224e59](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:76ad71ac-d40d-4e62-bfed-9faf4d224e59)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6e5dc4e3-4ad0-48f6-a92e-ec5b48324c03](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6e5dc4e3-4ad0-48f6-a92e-ec5b48324c03)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7d259a70-bb35-4988-b131-4da3335c7c5c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7d259a70-bb35-4988-b131-4da3335c7c5c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:27994a95-751f-4e15-81af-36fc5cc5cdf1](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:27994a95-751f-4e15-81af-36fc5cc5cdf1)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:752275bd-c88b-4a7e-aaa3-ec566d0dc197](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:752275bd-c88b-4a7e-aaa3-ec566d0dc197)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3cba3d9d-507e-4bae-a632-59698c4950ed](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3cba3d9d-507e-4bae-a632-59698c4950ed)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:552e3f5d-2964-445b-9fb2-eea3d1168252](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:552e3f5d-2964-445b-9fb2-eea3d1168252)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:0bad5ee8-4cfc-493c-9fc5-3707cc823f33](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:0bad5ee8-4cfc-493c-9fc5-3707cc823f33)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d7af253c-1e3e-41ad-b4b5-dc7cfabf621b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d7af253c-1e3e-41ad-b4b5-dc7cfabf621b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:1027d4a1-958f-49b4-ba5f-9afad1d2716a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:1027d4a1-958f-49b4-ba5f-9afad1d2716a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:1bfde910-3ec3-4a51-bb45-6a21c50bb095](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:1bfde910-3ec3-4a51-bb45-6a21c50bb095)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:25d69a0c-df35-49f0-8317-8ac2e05bd128](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:25d69a0c-df35-49f0-8317-8ac2e05bd128)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:35 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5abc072e-6017-471e-819e-ded12cf4f0eb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5abc072e-6017-471e-819e-ded12cf4f0eb)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:189a9f03-27b0-4e38-9a04-7ee3c458be0a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:189a9f03-27b0-4e38-9a04-7ee3c458be0a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7e30fd40-4a31-4221-b4de-c9250213b110](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7e30fd40-4a31-4221-b4de-c9250213b110)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c13fdaae-051b-45ca-86eb-654f5d22ccd4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c13fdaae-051b-45ca-86eb-654f5d22ccd4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:196dae1b-c441-487a-8f4d-9f5490b8f050](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:196dae1b-c441-487a-8f4d-9f5490b8f050)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:98ac2e31-249f-4ce4-8e91-a01888b0d1f7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:98ac2e31-249f-4ce4-8e91-a01888b0d1f7)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3432c068-6a90-430a-afa9-ebf451f2c1c1](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3432c068-6a90-430a-afa9-ebf451f2c1c1)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:0a68e165-fcff-4f83-92df-70f24faad4f4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:0a68e165-fcff-4f83-92df-70f24faad4f4)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:65e0d9c8-d710-43d4-8576-2bc3e7e7bf66](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:65e0d9c8-d710-43d4-8576-2bc3e7e7bf66)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:86c7df56-f19d-4c34-b6e4-7b60d7b964bb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:86c7df56-f19d-4c34-b6e4-7b60d7b964bb)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:89c03787-cfca-419c-8c91-bf739e40cb88](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:89c03787-cfca-419c-8c91-bf739e40cb88)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:1171ebfc-cc87-4ef3-9d3e-1c3f89469aea](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:1171ebfc-cc87-4ef3-9d3e-1c3f89469aea)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:4fb4c894-e2ee-4f1b-b2cd-93e75a838910](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:4fb4c894-e2ee-4f1b-b2cd-93e75a838910)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:318c8bc8-9cf7-41dd-bc3f-f7477684c582](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:318c8bc8-9cf7-41dd-bc3f-f7477684c582)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:560d8e3d-4e92-409f-b9f6-da86a710f732](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:560d8e3d-4e92-409f-b9f6-da86a710f732)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:860c45cd-4c02-4263-b41c-275c5c1b273d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:860c45cd-4c02-4263-b41c-275c5c1b273d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:cfab7960-1c8d-4465-b0a2-c1fc5f28b8d9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:cfab7960-1c8d-4465-b0a2-c1fc5f28b8d9)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:f1c77f18-b7c5-4351-ba96-c34c83dc9126](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:f1c77f18-b7c5-4351-ba96-c34c83dc9126)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:afe5d8e5-b4e6-41a6-8d33-4adbb5daf62d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:afe5d8e5-b4e6-41a6-8d33-4adbb5daf62d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:9a486c59-49db-42a1-8e9b-a0db598512d3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:9a486c59-49db-42a1-8e9b-a0db598512d3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:bd278d1e-088f-4ec9-adbc-1faeded7f13b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:bd278d1e-088f-4ec9-adbc-1faeded7f13b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:388bcb19-f02d-4279-8488-908dfe189f3b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:388bcb19-f02d-4279-8488-908dfe189f3b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3e20cb40-b498-4796-9978-985308efa639](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3e20cb40-b498-4796-9978-985308efa639)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:625238e6-8798-4ae5-ac74-2f872f755f23](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:625238e6-8798-4ae5-ac74-2f872f755f23)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d4a7647c-87d4-4357-9637-67063052acd0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d4a7647c-87d4-4357-9637-67063052acd0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:def44115-6fd4-4768-bd86-e80643a35064](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:def44115-6fd4-4768-bd86-e80643a35064)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d4368b6e-539b-4f84-8923-1bd3042212fc](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d4368b6e-539b-4f84-8923-1bd3042212fc)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:0d02b372-430b-4bb6-9f12-c67214edd8a3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:0d02b372-430b-4bb6-9f12-c67214edd8a3)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:bbf124dd-587d-47ed-9e77-c16855021cad](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:bbf124dd-587d-47ed-9e77-c16855021cad)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c294ece1-c65b-46a6-8729-bee5b5e4c00c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c294ece1-c65b-46a6-8729-bee5b5e4c00c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:7331aad3-2f71-451d-910e-f63435761866](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:7331aad3-2f71-451d-910e-f63435761866)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:5f1910c3-1ad2-4d44-9e71-a3ab91e600d6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:5f1910c3-1ad2-4d44-9e71-a3ab91e600d6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:20576f40-8804-4254-b4bf-609bcddeb264](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:20576f40-8804-4254-b4bf-609bcddeb264)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:72de7ae2-3f8d-4066-af3e-949ace677424](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:72de7ae2-3f8d-4066-af3e-949ace677424)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c713edc2-c783-4b76-a506-cb51df0f0ad5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c713edc2-c783-4b76-a506-cb51df0f0ad5)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:cdb90c94-fe6c-412f-8a65-330e569d90fd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:cdb90c94-fe6c-412f-8a65-330e569d90fd)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:eca21ec0-bf62-4c32-a6f2-5a5f97c16922](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:eca21ec0-bf62-4c32-a6f2-5a5f97c16922)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:34a6cf3d-d331-41a2-a5e8-69879cde023a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:34a6cf3d-d331-41a2-a5e8-69879cde023a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:f293e4ef-9ade-45c4-8786-677c1b09bcd5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:f293e4ef-9ade-45c4-8786-677c1b09bcd5)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:c95a6853-808a-4381-97a8-da6194ddd03c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:c95a6853-808a-4381-97a8-da6194ddd03c)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:dc96b1ce-f691-4f59-ba82-67354c2ed89b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:dc96b1ce-f691-4f59-ba82-67354c2ed89b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6a75671a-1f82-42ad-9259-e2c13f73c0df](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6a75671a-1f82-42ad-9259-e2c13f73c0df)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:61291e91-b0dd-4a38-9e1d-c5aaf9579338](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:61291e91-b0dd-4a38-9e1d-c5aaf9579338)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:246c1588-8926-46ae-a2ca-ebc5faf2e746](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:246c1588-8926-46ae-a2ca-ebc5faf2e746)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:8ebd8bc7-f348-417e-ac95-59adb3ed2933](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:8ebd8bc7-f348-417e-ac95-59adb3ed2933)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:3385c995-513d-47aa-b081-25315672dbb8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:3385c995-513d-47aa-b081-25315672dbb8)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:203b9afb-c71f-4d83-8656-73e1a703b1c5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:203b9afb-c71f-4d83-8656-73e1a703b1c5)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:9fd82253-4897-427c-a215-50f6b4a5febd](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:9fd82253-4897-427c-a215-50f6b4a5febd)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:b77c51c4-a060-4a54-8663-568080ff7ce7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:b77c51c4-a060-4a54-8663-568080ff7ce7)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:aeecebd6-d2e7-4fdc-9712-9827e520daeb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:aeecebd6-d2e7-4fdc-9712-9827e520daeb)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:68236147-8812-4bc1-97e9-badbb4922d8d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:68236147-8812-4bc1-97e9-badbb4922d8d)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:19ced729-426f-4387-86ba-f56436dfc31a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:19ced729-426f-4387-86ba-f56436dfc31a)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:88606fdd-f145-4a73-9210-a74efd98771e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:88606fdd-f145-4a73-9210-a74efd98771e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ae1e70b7-837d-4812-a493-9ec565cde1e6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ae1e70b7-837d-4812-a493-9ec565cde1e6)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6a0f74f1-aebe-4306-b7ab-ef726c921ee7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6a0f74f1-aebe-4306-b7ab-ef726c921ee7)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:6380fdea-968a-4c6e-8916-0f34d921074b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:6380fdea-968a-4c6e-8916-0f34d921074b)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:0cf7e3a1-c348-4f9b-a998-d910dfa17da8](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:0cf7e3a1-c348-4f9b-a998-d910dfa17da8)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:0b3d79b7-a884-498c-b3a2-9cf85a188e2e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:0b3d79b7-a884-498c-b3a2-9cf85a188e2e)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:da57686e-c6c9-4ec4-ad2a-5a01d4f97151](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:da57686e-c6c9-4ec4-ad2a-5a01d4f97151)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:a4e41091-a17b-4132-8037-ba3b2d277911](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:a4e41091-a17b-4132-8037-ba3b2d277911)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' object (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:ce3d9d28-7d35-42c9-87e4-686342e0aab0](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:ce3d9d28-7d35-42c9-87e4-686342e0aab0)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:d5770a00-12ee-4c43-a38d-95b13a8a75ad](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:d5770a00-12ee-4c43-a38d-95b13a8a75ad)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-offering-price:e5ec47b1-7bf9-4822-bad7-3cc3b55dc3a2](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productOfferingPrice/urn:ngsi-ld:product-offering-price:e5ec47b1-7bf9-4822-bad7-3cc3b55dc3a2)

- **Type:** productOfferingPrice
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:42 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOfferingPrice (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

### productOrder Objects

#### Object: [urn:ngsi-ld:product-order:fdc85cc5-6e74-4e67-a9d0-1f3d7a79dce3](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:fdc85cc5-6e74-4e67-a9d0-1f3d7a79dce3)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:4902aa01-8949-4a46-9be4-1cf98aa4a4f7](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:4902aa01-8949-4a46-9be4-1cf98aa4a4f7)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:117780df-ea0f-45b9-bcb3-1cbe5b17b1f9](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:117780df-ea0f-45b9-bcb3-1cbe5b17b1f9)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:06350c8a-4238-4284-a9b4-f5c047b91292](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:06350c8a-4238-4284-a9b4-f5c047b91292)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:9dee5ca2-947e-4443-aca8-f82ec936d2b2](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:9dee5ca2-947e-4443-aca8-f82ec936d2b2)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:fa5e91ba-6418-4354-b840-1f9e31ca1e74](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:fa5e91ba-6418-4354-b840-1f9e31ca1e74)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:abee6b6d-f3e0-4465-bdac-bea275ddecc2](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:abee6b6d-f3e0-4465-bdac-bea275ddecc2)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:1c4401c3-4770-403e-af95-6b9737bc087a](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:1c4401c3-4770-403e-af95-6b9737bc087a)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:525f2198-45a7-44a9-a0d3-413972bb5f56](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:525f2198-45a7-44a9-a0d3-413972bb5f56)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:bfdd1be4-a32b-4017-855d-d4fdaa95818a](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:bfdd1be4-a32b-4017-855d-d4fdaa95818a)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:0c6cc039-21bd-462b-9a45-b6ec77245e96](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:0c6cc039-21bd-462b-9a45-b6ec77245e96)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:367cc7f9-2ea7-4fc8-b9d4-1f2aff2e57d0](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:367cc7f9-2ea7-4fc8-b9d4-1f2aff2e57d0)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:4ffa78f5-2b85-4c9e-b09f-00d0ecad3a44](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:4ffa78f5-2b85-4c9e-b09f-00d0ecad3a44)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:d7971e80-dd40-42c3-b33d-1a5485dc38a4](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:d7971e80-dd40-42c3-b33d-1a5485dc38a4)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:18777b58-f7d2-4570-ab6d-2917f86b3c55](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:18777b58-f7d2-4570-ab6d-2917f86b3c55)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:740e7b80-672d-4e6a-9f29-67c16469a4e1](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:740e7b80-672d-4e6a-9f29-67c16469a4e1)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:95f1cf2c-48b1-4f6c-a844-accbb3c8816c](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:95f1cf2c-48b1-4f6c-a844-accbb3c8816c)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:fe85000e-ad28-441c-94ac-e5648df09280](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:fe85000e-ad28-441c-94ac-e5648df09280)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:898c190e-94d6-486d-a9d5-d19efb3d0c07](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:898c190e-94d6-486d-a9d5-d19efb3d0c07)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:803f018a-e174-4871-b87f-0c348a0a5b94](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:803f018a-e174-4871-b87f-0c348a0a5b94)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:88870dfa-14be-4a22-98d1-3186b78f1778](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:88870dfa-14be-4a22-98d1-3186b78f1778)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:d4277dff-1a6d-4907-9862-5fe0f1479469](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:d4277dff-1a6d-4907-9862-5fe0f1479469)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:dbd270d5-355d-4ad0-b3ab-1224b31e6445](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:dbd270d5-355d-4ad0-b3ab-1224b31e6445)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:f0cb34b1-3d90-4b40-8f75-d02481114b77](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:f0cb34b1-3d90-4b40-8f75-d02481114b77)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:eff32662-fe9c-4790-8453-6854fae46f1c](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:eff32662-fe9c-4790-8453-6854fae46f1c)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:6b541f51-1e2c-4b94-82e4-ce3f84291b12](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:6b541f51-1e2c-4b94-82e4-ce3f84291b12)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:4706b999-549d-43c8-ad67-0993854a77c8](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:4706b999-549d-43c8-ad67-0993854a77c8)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:38b7c870-ac61-4f83-b318-a50974141de6](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:38b7c870-ac61-4f83-b318-a50974141de6)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:93e1a2cf-59fa-4894-a4ef-553ee0cdc9b7](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:93e1a2cf-59fa-4894-a4ef-553ee0cdc9b7)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:46f10f36-8e34-42be-9d0d-9222545375bd](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:46f10f36-8e34-42be-9d0d-9222545375bd)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:e3270bdb-eb80-4786-a31b-8d8e15f8e9da](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:e3270bdb-eb80-4786-a31b-8d8e15f8e9da)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:75c36489-3505-4720-954c-85978149052c](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:75c36489-3505-4720-954c-85978149052c)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:2513eb2e-4019-4d21-a6cb-01e6195a4391](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:2513eb2e-4019-4d21-a6cb-01e6195a4391)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:8de7dd6b-2b3b-475d-942f-37948e44381c](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:8de7dd6b-2b3b-475d-942f-37948e44381c)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:a1435e0b-d17f-4397-9872-62cb1efdadc3](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:a1435e0b-d17f-4397-9872-62cb1efdadc3)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:83d3cfe4-8e37-458a-b8d3-422aaa608e69](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:83d3cfe4-8e37-458a-b8d3-422aaa608e69)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:c081d3a4-d6c2-425e-ad7a-d83584aeecff](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:c081d3a4-d6c2-425e-ad7a-d83584aeecff)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:f20d87a8-9ba6-46af-b1e4-53fef5e2299a](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:f20d87a8-9ba6-46af-b1e4-53fef5e2299a)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:b31a12fa-bd5f-47eb-a78b-85fd69a879df](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:b31a12fa-bd5f-47eb-a78b-85fd69a879df)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:686aced4-03f6-4f7d-838b-e5813fbbb6fc](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:686aced4-03f6-4f7d-838b-e5813fbbb6fc)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:a5af6327-3557-4527-a4b3-eecd23c9601b](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:a5af6327-3557-4527-a4b3-eecd23c9601b)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:ba655feb-9f9b-4d45-91a5-7cf4fefe8439](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:ba655feb-9f9b-4d45-91a5-7cf4fefe8439)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:9167c4c4-c9d7-4dd1-bab8-bdaa1e97232b](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:9167c4c4-c9d7-4dd1-bab8-bdaa1e97232b)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:86dbcf46-379e-454a-a77a-ed8e903baafe](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:86dbcf46-379e-454a-a77a-ed8e903baafe)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:822acf8c-4b70-4608-8289-47de3f55cf40](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:822acf8c-4b70-4608-8289-47de3f55cf40)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:932d989e-9129-40d0-8d7f-3acc042312cb](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:932d989e-9129-40d0-8d7f-3acc042312cb)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:9a9f16d8-2b81-4ef8-ba4e-36228f33f36c](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:9a9f16d8-2b81-4ef8-ba4e-36228f33f36c)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:c1214f15-923a-4e41-bbf7-e7bcad11216a](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:c1214f15-923a-4e41-bbf7-e7bcad11216a)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:11c88d4f-2b28-470c-b057-3e1b8e4dcf5f](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:11c88d4f-2b28-470c-b057-3e1b8e4dcf5f)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:7b558e5e-fd00-43ca-a672-ed5d4c917038](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:7b558e5e-fd00-43ca-a672-ed5d4c917038)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:51678d2f-c987-4b7c-840c-785a8aecce55](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:51678d2f-c987-4b7c-840c-785a8aecce55)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:3b9bf611-738e-4d28-9fba-5d32dc6c6bb1](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:3b9bf611-738e-4d28-9fba-5d32dc6c6bb1)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:1567676a-d3c5-4ca9-b464-abed2f0742a5](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:1567676a-d3c5-4ca9-b464-abed2f0742a5)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:1a745f6d-bcdb-40b1-8532-b0fd638e2ec7](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:1a745f6d-bcdb-40b1-8532-b0fd638e2ec7)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:ab616fcd-0b62-4aed-9138-2c7aad0b2d42](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:ab616fcd-0b62-4aed-9138-2c7aad0b2d42)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:1c3655bc-df51-4d65-9508-7d075ffa3d33](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:1c3655bc-df51-4d65-9508-7d075ffa3d33)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:fb5f93a2-99e0-4802-8ee1-be2c5631b487](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:fb5f93a2-99e0-4802-8ee1-be2c5631b487)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:ab734040-739c-4ead-9a4b-1b3dd7c92e7c](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:ab734040-739c-4ead-9a4b-1b3dd7c92e7c)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:045ae968-e1d7-4ad6-9815-ef52444b556e](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:045ae968-e1d7-4ad6-9815-ef52444b556e)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:244f947f-a211-45e0-92d9-181cf52e22c5](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:244f947f-a211-45e0-92d9-181cf52e22c5)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:57b2c9c4-90bd-4b65-9259-3245c5412557](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:57b2c9c4-90bd-4b65-9259-3245c5412557)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:47811485-15a5-4a32-83d2-1e045d66418d](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:47811485-15a5-4a32-83d2-1e045d66418d)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:65046acc-e7a6-42a8-b9b4-037f6864ea78](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:65046acc-e7a6-42a8-b9b4-037f6864ea78)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:9bbbb640-f6c2-4ffa-9ffa-2ccbfb4960cd](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:9bbbb640-f6c2-4ffa-9ffa-2ccbfb4960cd)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:ffa828ba-1f6a-43ac-b6fc-4bb0f3c91b6b](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:ffa828ba-1f6a-43ac-b6fc-4bb0f3c91b6b)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:456add68-c16d-4360-9f82-c4e44f07585d](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:456add68-c16d-4360-9f82-c4e44f07585d)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:44dd81dc-7b13-45bf-9939-3148b5b3cea7](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:44dd81dc-7b13-45bf-9939-3148b5b3cea7)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:90cd0e3b-5c96-4490-96fc-cf242fcb517c](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:90cd0e3b-5c96-4490-96fc-cf242fcb517c)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:d5c8d11b-b638-4240-a554-357183250d92](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:d5c8d11b-b638-4240-a554-357183250d92)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:d0d52d66-ff36-49cd-bb59-017771e03395](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:d0d52d66-ff36-49cd-bb59-017771e03395)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:33b43e7a-cff2-4e69-b879-9ed9e0904838](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:33b43e7a-cff2-4e69-b879-9ed9e0904838)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:667f766c-9f7e-4699-af59-edc6537ba4a7](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:667f766c-9f7e-4699-af59-edc6537ba4a7)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:56081401-5159-432c-b59b-1da51815fba1](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:56081401-5159-432c-b59b-1da51815fba1)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:f9583989-3f7f-4a63-a7d5-e1e29de5e259](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:f9583989-3f7f-4a63-a7d5-e1e29de5e259)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:d9ba0c19-68d2-4bdc-ae3e-e578e579d84b](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:d9ba0c19-68d2-4bdc-ae3e-e578e579d84b)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:9b00eecc-c2f6-4544-8e31-748d61867c06](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:9b00eecc-c2f6-4544-8e31-748d61867c06)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:3c5cf5e3-3d86-4469-a9ac-23f2d28dbda6](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:3c5cf5e3-3d86-4469-a9ac-23f2d28dbda6)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:304f90fc-42e7-4fce-a8c3-60d031548c81](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:304f90fc-42e7-4fce-a8c3-60d031548c81)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:90f607ca-4c54-401b-9976-36f85d1308bb](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:90f607ca-4c54-401b-9976-36f85d1308bb)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:e85bc0fd-de26-4052-a32e-4d0278ebdd57](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:e85bc0fd-de26-4052-a32e-4d0278ebdd57)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:f67c89d7-0d37-4824-b53e-1f89dacb2396](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:f67c89d7-0d37-4824-b53e-1f89dacb2396)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:84126d35-3052-4250-9510-be0f5bc976cc](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:84126d35-3052-4250-9510-be0f5bc976cc)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-order:b443cf29-005d-483e-ae3a-76a4b5af1319](https://tmf.dome-marketplace-sbx.org/tmf-api/productOrderingManagement/v4/productOrder/urn:ngsi-ld:product-order:b443cf29-005d-483e-ae3a-76a4b5af1319)

- **Type:** productOrder
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productOrder (Code: MISSING_RECOMMENDED_FIELD)
  - name: Recommended field 'name' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

### productSpecification Objects

#### Object: [urn:ngsi-ld:product-specification:bf72e349-03f0-4e88-965a-73815d8881b4](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:bf72e349-03f0-4e88-965a-73815d8881b4)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:55158c3d-ef8f-4f1c-b9d0-82dd3138b2ae](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:55158c3d-ef8f-4f1c-b9d0-82dd3138b2ae)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:5e2ced54-45f1-4687-b9f9-ee13cb318a66](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:5e2ced54-45f1-4687-b9f9-ee13cb318a66)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:ba4123f5-77f2-4b80-8c8d-13c01d3a6e72](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:ba4123f5-77f2-4b80-8c8d-13c01d3a6e72)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:534697ce-c774-4cc2-a8f4-d4f2ae6cf4f9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:534697ce-c774-4cc2-a8f4-d4f2ae6cf4f9)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:bfe7418d-2614-4e6a-920c-6d6711171f10](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:bfe7418d-2614-4e6a-920c-6d6711171f10)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:351bbc6e-515f-4159-9082-e60563c927da](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:351bbc6e-515f-4159-9082-e60563c927da)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:1cb811e6-4661-41ff-935c-26176e209dab](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:1cb811e6-4661-41ff-935c-26176e209dab)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:32096043-6ce0-49c0-acb6-4ffde182b254](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:32096043-6ce0-49c0-acb6-4ffde182b254)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:36f45313-8f8d-49c4-bbf2-bd2816934e6d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:36f45313-8f8d-49c4-bbf2-bd2816934e6d)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:1226966e-795f-4508-a876-6c67c5cb1059](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:1226966e-795f-4508-a876-6c67c5cb1059)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:8e867c80-2a58-49c7-9502-30f8373802ad](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:8e867c80-2a58-49c7-9502-30f8373802ad)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:9f3eebce-ddfd-40a4-9227-a62f98f94344](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:9f3eebce-ddfd-40a4-9227-a62f98f94344)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:839e37f8-5f9d-48fc-b2c9-3e5a32d90424](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:839e37f8-5f9d-48fc-b2c9-3e5a32d90424)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC

#### Object: [urn:ngsi-ld:product-specification:2adb577d-ec29-4406-b2d5-86b295c2c434](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:2adb577d-ec29-4406-b2d5-86b295c2c434)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC

#### Object: [urn:ngsi-ld:product-specification:64c19187-4f24-4dcb-bfa2-a06afd3b5b11](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:64c19187-4f24-4dcb-bfa2-a06afd3b5b11)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:6a8476a3-88e5-49de-873b-4ad36ebf6db5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:6a8476a3-88e5-49de-873b-4ad36ebf6db5)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:cb3d1a1a-802a-4463-bc2d-056014051201](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:cb3d1a1a-802a-4463-bc2d-056014051201)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:9f141090-01ef-4988-a639-97d84f2d76ba](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:9f141090-01ef-4988-a639-97d84f2d76ba)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:1fd76f44-79c6-46d6-9fb1-11899181e905](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:1fd76f44-79c6-46d6-9fb1-11899181e905)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:e4794956-d42e-46c9-a2ce-3df53e2df26d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:e4794956-d42e-46c9-a2ce-3df53e2df26d)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:411816fd-3f8b-4e6f-8c35-02e8cada5ce9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:411816fd-3f8b-4e6f-8c35-02e8cada5ce9)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:3c2ea42e-77de-4012-8d12-2e760facca23](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:3c2ea42e-77de-4012-8d12-2e760facca23)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:a45a592d-480d-429d-8ad8-f81b1bf528bb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:a45a592d-480d-429d-8ad8-f81b1bf528bb)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:ace61c97-0f51-4601-8359-59b3987ff3f7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:ace61c97-0f51-4601-8359-59b3987ff3f7)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:ccaec925-9fad-40c0-8015-5b9087f92aff](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:ccaec925-9fad-40c0-8015-5b9087f92aff)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:2532e67f-0a84-4aed-b31f-7654984daac9](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:2532e67f-0a84-4aed-b31f-7654984daac9)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:10ebf2c1-fe79-4191-81a3-b58207307c5a](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:10ebf2c1-fe79-4191-81a3-b58207307c5a)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:7b61b736-d15c-411c-90fa-649508d2a9f5](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:7b61b736-d15c-411c-90fa-649508d2a9f5)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:021517fe-650f-4834-92f1-4bf29487c5e3](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:021517fe-650f-4834-92f1-4bf29487c5e3)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:90962671-9c9d-4328-b78b-b551b74fdc92](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:90962671-9c9d-4328-b78b-b551b74fdc92)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:e84e6c6c-7791-40f1-8904-b14134e3cbcb](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:e84e6c6c-7791-40f1-8904-b14134e3cbcb)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:e9ac6275-bb21-47d7-89bd-187f71f55e3c](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:e9ac6275-bb21-47d7-89bd-187f71f55e3c)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:95332c87-ccb7-4ab2-bce9-78b21295e06f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:95332c87-ccb7-4ab2-bce9-78b21295e06f)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:7f3a773f-969e-479c-bbbf-d9a11292351f](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:7f3a773f-969e-479c-bbbf-d9a11292351f)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:8569bbcc-aca6-4aa7-b26e-290e9191a352](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:8569bbcc-aca6-4aa7-b26e-290e9191a352)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:3e0e7cc0-77ef-4546-9c71-0254b53c147d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:3e0e7cc0-77ef-4546-9c71-0254b53c147d)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:21d73795-9213-4ab8-a721-c8973fa7acce](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:21d73795-9213-4ab8-a721-c8973fa7acce)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:e3fb95ab-8cc0-4ad1-8b83-881ab14b8d36](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:e3fb95ab-8cc0-4ad1-8b83-881ab14b8d36)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:ec990ad1-fc40-43b8-8681-5f327ee32c3e](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:ec990ad1-fc40-43b8-8681-5f327ee32c3e)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:21c93fdf-e74b-4815-84cb-7efa7932fa40](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:21c93fdf-e74b-4815-84cb-7efa7932fa40)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:70c8489b-2116-49d2-95ef-ff12e0a3fb8d](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:70c8489b-2116-49d2-95ef-ff12e0a3fb8d)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:6bf3eab8-25fd-42f7-a809-170835e9bfc7](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:6bf3eab8-25fd-42f7-a809-170835e9bfc7)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:dc9dea35-41ab-4a3c-b4c2-76a80f6a9aca](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:dc9dea35-41ab-4a3c-b4c2-76a80f6a9aca)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:5d32ba32-c7d5-43fc-b64a-2c00a21a470b](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:5d32ba32-c7d5-43fc-b64a-2c00a21a470b)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:f046a9e9-781f-4352-a1ce-6f7b1728fc94](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:f046a9e9-781f-4352-a1ce-6f7b1728fc94)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:d881ea40-8069-4973-9bea-b99f4d0a2a07](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:d881ea40-8069-4973-9bea-b99f4d0a2a07)

- **Type:** productSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:product-specification:8561a66d-9875-4bf9-8c92-6819bc6bb3d6](https://tmf.dome-marketplace-sbx.org/tmf-api/productCatalogManagement/v4/productSpecification/urn:ngsi-ld:product-specification:8561a66d-9875-4bf9-8c92-6819bc6bb3d6)

- **Type:** productSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing or invalid 'name' field in 'relatedParty' object for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Seller' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to productSpecification (Code: MISSING_RECOMMENDED_FIELD)

### resource Objects

#### Object: [urn:ngsi-ld:resource:8e58861a-5f65-4873-b0d4-d74296160870](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:8e58861a-5f65-4873-b0d4-d74296160870)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:2abad639-11bc-4dff-a554-d5c4a8a175b5](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:2abad639-11bc-4dff-a554-d5c4a8a175b5)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:1c06d20b-302f-43d2-8928-3df67e96eb7f](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:1c06d20b-302f-43d2-8928-3df67e96eb7f)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:e21b848b-74ca-4b93-9127-e1f043da1660](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:e21b848b-74ca-4b93-9127-e1f043da1660)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:94e4c43b-9a49-4880-9144-53db38b69083](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:94e4c43b-9a49-4880-9144-53db38b69083)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:c56a4d6c-43a0-45b6-b61e-b9ff13aed81f](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:c56a4d6c-43a0-45b6-b61e-b9ff13aed81f)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:bb198be7-01b5-4c40-837b-b1b97b2ab5bb](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:bb198be7-01b5-4c40-837b-b1b97b2ab5bb)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:80fe9ef6-c5df-4403-89f9-fcea27e8157b](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:80fe9ef6-c5df-4403-89f9-fcea27e8157b)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:b706d1c7-0ec5-4e6f-a11d-0ca1e3f5ceea](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:b706d1c7-0ec5-4e6f-a11d-0ca1e3f5ceea)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:0a3bdcff-f809-45de-8abd-72e197478580](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:0a3bdcff-f809-45de-8abd-72e197478580)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:7465525d-582d-4b7b-928a-c389de3c5808](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:7465525d-582d-4b7b-928a-c389de3c5808)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:a7bcfbf8-7692-4ea7-ae20-03f01a2568aa](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:a7bcfbf8-7692-4ea7-ae20-03f01a2568aa)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:0541b853-2bdf-4774-9ecd-29bdab473475](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:0541b853-2bdf-4774-9ecd-29bdab473475)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:fe949971-f5eb-4f6d-bd83-37208e808cad](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:fe949971-f5eb-4f6d-bd83-37208e808cad)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:1d6a9d07-b4b1-48ec-8a72-efe9fbc476ff](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:1d6a9d07-b4b1-48ec-8a72-efe9fbc476ff)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:876f576f-c686-48f9-a5ce-d24bf2ff6343](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:876f576f-c686-48f9-a5ce-d24bf2ff6343)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:8e4d2ad7-fc63-4d4a-9b2e-5a3514c3a56a](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:8e4d2ad7-fc63-4d4a-9b2e-5a3514c3a56a)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:c714add4-d818-4a23-ae1d-f770cf56ccb3](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:c714add4-d818-4a23-ae1d-f770cf56ccb3)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:d41ca912-91ae-4bab-be0f-e2ce3419d33f](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:d41ca912-91ae-4bab-be0f-e2ce3419d33f)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:a6f30dbd-8a74-4427-b5d5-abe19dc61d79](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:a6f30dbd-8a74-4427-b5d5-abe19dc61d79)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:a700af5a-f59e-4ba0-b801-4ba09c3d2115](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:a700af5a-f59e-4ba0-b801-4ba09c3d2115)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:1cb497c3-6212-4fee-a042-11af0e15dda7](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:1cb497c3-6212-4fee-a042-11af0e15dda7)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:47494568-4ca8-4fbd-9a02-484f368b9558](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:47494568-4ca8-4fbd-9a02-484f368b9558)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:27d89a1a-fbac-4c1d-b4c6-e585b49f2d2f](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:27d89a1a-fbac-4c1d-b4c6-e585b49f2d2f)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:3a6ab23e-a3c1-45ab-91cc-a4f4dedeac24](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:3a6ab23e-a3c1-45ab-91cc-a4f4dedeac24)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:ce133ff0-0eb5-46db-a837-63c0af515aba](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:ce133ff0-0eb5-46db-a837-63c0af515aba)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:9accc80d-1490-4994-a1fa-8624bcf968c8](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:9accc80d-1490-4994-a1fa-8624bcf968c8)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource:8d70ccb9-5f54-42b4-aeb2-14a15348ba51](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceInventoryManagement/v4/resource/urn:ngsi-ld:resource:8d70ccb9-5f54-42b4-aeb2-14a15348ba51)

- **Type:** resource
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:46 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resource (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

### resourceSpecification Objects

#### Object: [urn:ngsi-ld:resource-specification:9bd04c66-97a5-489a-ae93-7dddb4cd341b](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:9bd04c66-97a5-489a-ae93-7dddb4cd341b)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:f0153609-77d6-4464-9352-679f8d0e015f](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:f0153609-77d6-4464-9352-679f8d0e015f)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:e84e77c9-55e1-4c2f-953a-09dd52003f92](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:e84e77c9-55e1-4c2f-953a-09dd52003f92)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:59cf3608-2879-4f38-a8f7-5baf61653c92](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:59cf3608-2879-4f38-a8f7-5baf61653c92)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:46e91850-687d-4454-a53a-b2211bc8f4b8](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:46e91850-687d-4454-a53a-b2211bc8f4b8)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:996d8528-7988-4698-a0c7-613e7e646de5](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:996d8528-7988-4698-a0c7-613e7e646de5)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:ff1e8458-12e5-4c2f-9b10-27f0111c76eb](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:ff1e8458-12e5-4c2f-9b10-27f0111c76eb)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:31c96384-4fca-49ac-b67b-e37e9ba17ad5](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:31c96384-4fca-49ac-b67b-e37e9ba17ad5)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:cac71f92-58a7-4eeb-b7ef-a65f017de86d](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:cac71f92-58a7-4eeb-b7ef-a65f017de86d)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:d3d44ad6-691d-4456-8e87-014cb9cceceb](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:d3d44ad6-691d-4456-8e87-014cb9cceceb)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:5338b9dd-5faf-4085-b63d-df2885e87bd4](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:5338b9dd-5faf-4085-b63d-df2885e87bd4)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:cea066f1-a03b-401e-81ca-17c4e34b57da](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:cea066f1-a03b-401e-81ca-17c4e34b57da)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:98e53881-0968-455f-8ca2-ca22bab21cc5](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:98e53881-0968-455f-8ca2-ca22bab21cc5)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:167a0638-cdd8-4535-abe8-07d903e673e9](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:167a0638-cdd8-4535-abe8-07d903e673e9)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:2411f1cb-f15e-4c7c-a391-55d1377bf602](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:2411f1cb-f15e-4c7c-a391-55d1377bf602)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:resource-specification:0d389ed9-44c7-4d11-9afd-4ccf467943bd](https://tmf.dome-marketplace-sbx.org/tmf-api/resourceCatalog/v4/resourceSpecification/urn:ngsi-ld:resource-specification:0d389ed9-44c7-4d11-9afd-4ccf467943bd)

- **Type:** resourceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to resourceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

### service Objects

#### Object: [urn:ngsi-ld:service:7c4c95cf-0b7d-40d0-890a-4dc72833bb93](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:7c4c95cf-0b7d-40d0-890a-4dc72833bb93)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:8da6974f-43e3-4cb8-9dd0-16c7235845be](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:8da6974f-43e3-4cb8-9dd0-16c7235845be)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:46da6440-ab04-40a5-829d-1e0733e20dce](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:46da6440-ab04-40a5-829d-1e0733e20dce)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:3d8a197f-d247-4a51-a93b-ebfe97589258](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:3d8a197f-d247-4a51-a93b-ebfe97589258)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:f0477ba9-7d87-4ed9-a7e1-2835b31196dc](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:f0477ba9-7d87-4ed9-a7e1-2835b31196dc)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:1f25b49d-b116-4348-837c-85b6dffaa323](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:1f25b49d-b116-4348-837c-85b6dffaa323)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:4913d3de-fa4f-4e69-9e09-d8cb7bf1ca03](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:4913d3de-fa4f-4e69-9e09-d8cb7bf1ca03)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:c16510b6-291f-4b8a-a1a7-5d4c614568e2](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:c16510b6-291f-4b8a-a1a7-5d4c614568e2)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:e5b86a59-e684-4063-8db2-92e08a003f16](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:e5b86a59-e684-4063-8db2-92e08a003f16)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:11ada031-6eed-4385-814b-0d0e18a2fab0](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:11ada031-6eed-4385-814b-0d0e18a2fab0)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:5410fce1-2c30-4e04-8849-8b171ff80250](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:5410fce1-2c30-4e04-8849-8b171ff80250)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:1659addd-fd42-41fc-92ad-0216619b0d38](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:1659addd-fd42-41fc-92ad-0216619b0d38)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:d4a44923-72b3-4181-92b6-cdf8cf77b097](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:d4a44923-72b3-4181-92b6-cdf8cf77b097)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:c7af8f17-aaa6-4016-b3ec-37b05013bbeb](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:c7af8f17-aaa6-4016-b3ec-37b05013bbeb)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:4c633ebc-f5e4-470b-8232-2690d1299f6e](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:4c633ebc-f5e4-470b-8232-2690d1299f6e)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:89d65206-4e9e-44d2-b599-5f421bd27035](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:89d65206-4e9e-44d2-b599-5f421bd27035)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:b0fe3d04-5383-4eb4-a2a7-21567bd04a8e](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:b0fe3d04-5383-4eb4-a2a7-21567bd04a8e)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:4624e43f-9d3f-432f-9276-a24876a4343d](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:4624e43f-9d3f-432f-9276-a24876a4343d)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:6d0509ff-d999-4c18-b789-695f585a26d9](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:6d0509ff-d999-4c18-b789-695f585a26d9)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:a8b86ff8-c985-4c14-a511-d28014b5f619](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:a8b86ff8-c985-4c14-a511-d28014b5f619)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:694594f9-c1eb-404b-bc88-6eb730e3721e](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:694594f9-c1eb-404b-bc88-6eb730e3721e)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:a362767d-c4c1-421c-88ee-bc30d2789fa8](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:a362767d-c4c1-421c-88ee-bc30d2789fa8)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:9d2e3dad-680a-4eed-8522-1872a7d6cef8](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:9d2e3dad-680a-4eed-8522-1872a7d6cef8)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:cdd2bfc2-4358-44f4-af9f-22f1557761c0](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:cdd2bfc2-4358-44f4-af9f-22f1557761c0)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:68c7911b-7d29-447f-80d3-b8a7d893d671](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:68c7911b-7d29-447f-80d3-b8a7d893d671)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:e29b32ad-0cc5-4c0b-b82b-d064777fd69d](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:e29b32ad-0cc5-4c0b-b82b-d064777fd69d)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:421ceae5-3635-4797-a28a-92d2f337aab6](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:421ceae5-3635-4797-a28a-92d2f337aab6)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:efe10698-a942-415d-af3e-bce1fdf02562](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:efe10698-a942-415d-af3e-bce1fdf02562)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:f843d6ae-b369-4913-a85c-6d996de7d5cd](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:f843d6ae-b369-4913-a85c-6d996de7d5cd)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:9e80e0cb-2388-433a-81e9-5e62e5086293](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:9e80e0cb-2388-433a-81e9-5e62e5086293)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:30812ea8-8da9-4563-be81-515e5d732d1c](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:30812ea8-8da9-4563-be81-515e5d732d1c)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:43b03329-9231-45af-8c34-9595f2db5749](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:43b03329-9231-45af-8c34-9595f2db5749)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:6477e357-1175-4892-89ef-d59ca4c8d129](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:6477e357-1175-4892-89ef-d59ca4c8d129)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:0a5a2a4f-622b-4800-b85a-de07cb45f1da](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:0a5a2a4f-622b-4800-b85a-de07cb45f1da)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:00a9d53d-c2b8-413c-a5df-dad0753fccee](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:00a9d53d-c2b8-413c-a5df-dad0753fccee)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:66b48b0e-9f8e-48c1-b26d-4363e423ee70](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:66b48b0e-9f8e-48c1-b26d-4363e423ee70)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:0598499d-ae4e-479e-8a37-cb20390bf47f](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:0598499d-ae4e-479e-8a37-cb20390bf47f)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:76b41627-731b-40a6-917c-98143156aeec](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:76b41627-731b-40a6-917c-98143156aeec)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:e1848935-7aaa-4894-a14c-f0ae5497763f](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:e1848935-7aaa-4894-a14c-f0ae5497763f)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:3bb1adf7-f7df-40b8-a961-48fd8920628b](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:3bb1adf7-f7df-40b8-a961-48fd8920628b)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:e69da896-6874-4431-a6f0-0f011cc7a3d3](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:e69da896-6874-4431-a6f0-0f011cc7a3d3)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:c90151d4-d176-466e-b2e8-d9248f65ecbe](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:c90151d4-d176-466e-b2e8-d9248f65ecbe)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:d1e316ed-6f50-42f1-8f94-7cd22bdc3ae4](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:d1e316ed-6f50-42f1-8f94-7cd22bdc3ae4)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:7378ccbd-4cf1-4536-ac90-a8b651eee300](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:7378ccbd-4cf1-4536-ac90-a8b651eee300)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:798ad965-6955-4f10-9113-3ef76efa9ace](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:798ad965-6955-4f10-9113-3ef76efa9ace)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:d4d981b3-925b-42af-b060-f7fbaee178d0](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:d4d981b3-925b-42af-b060-f7fbaee178d0)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:73fad827-19d7-4aeb-afad-bb07e6fc1d3e](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:73fad827-19d7-4aeb-afad-bb07e6fc1d3e)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:b2db24c7-c520-412b-9881-0b5b130195e0](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:b2db24c7-c520-412b-9881-0b5b130195e0)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:c6b30737-1bd9-41af-9b35-70d1d2fdead0](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:c6b30737-1bd9-41af-9b35-70d1d2fdead0)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:25bd5bdf-c46c-4af9-92fe-bbf8b8570295](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:25bd5bdf-c46c-4af9-92fe-bbf8b8570295)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:7dbeef12-eb5b-409e-9acf-d9800198974f](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:7dbeef12-eb5b-409e-9acf-d9800198974f)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:f4b83327-b60d-45be-9853-ae347dd7965f](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:f4b83327-b60d-45be-9853-ae347dd7965f)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:d09e3723-7471-43ce-99be-92473b1b8e23](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:d09e3723-7471-43ce-99be-92473b1b8e23)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:4e6b7e95-f622-4cd3-965b-87b67c18fa7b](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:4e6b7e95-f622-4cd3-965b-87b67c18fa7b)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:d8be7dcd-6585-42c5-8104-eadbde259c02](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:d8be7dcd-6585-42c5-8104-eadbde259c02)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:673a1a4b-8fa6-4b6e-8b75-eb960103aafe](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:673a1a4b-8fa6-4b6e-8b75-eb960103aafe)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:47a0eeb6-92d8-4ddf-a060-f6a4c9ed5fc1](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:47a0eeb6-92d8-4ddf-a060-f6a4c9ed5fc1)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:718dc9cb-1b02-411f-8fa8-63a7961bb563](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:718dc9cb-1b02-411f-8fa8-63a7961bb563)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:d257497c-f703-484d-8dec-82e01321072d](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:d257497c-f703-484d-8dec-82e01321072d)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:69254f62-2af0-4931-8d5a-f949708b3974](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:69254f62-2af0-4931-8d5a-f949708b3974)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:ed6eed98-3f29-4f18-a93f-6db2ca4735c8](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:ed6eed98-3f29-4f18-a93f-6db2ca4735c8)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:f67e1b72-72ad-4a84-8603-322aaf2fb55a](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:f67e1b72-72ad-4a84-8603-322aaf2fb55a)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:db2d6070-3c20-4c66-85f1-a4ce41377823](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:db2d6070-3c20-4c66-85f1-a4ce41377823)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:2c5f61f4-5faf-45eb-a9c1-234d8615c9a5](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:2c5f61f4-5faf-45eb-a9c1-234d8615c9a5)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:f62c0a94-3314-4f93-a1a5-297d55e20b5f](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:f62c0a94-3314-4f93-a1a5-297d55e20b5f)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:b90a8a01-2886-4113-996c-ba398b210c39](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:b90a8a01-2886-4113-996c-ba398b210c39)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:e98f8a5e-7574-4e39-9610-ae150027557d](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:e98f8a5e-7574-4e39-9610-ae150027557d)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:abe6375c-7732-4470-b471-5da39ff214ed](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:abe6375c-7732-4470-b471-5da39ff214ed)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:e1a9680f-f415-49c4-aad4-ca00d87cce83](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:e1a9680f-f415-49c4-aad4-ca00d87cce83)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:f410a722-b15f-4553-8a56-fedb8b306921](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:f410a722-b15f-4553-8a56-fedb8b306921)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:0139f0a7-6696-4de9-b257-e6f13cd9137a](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:0139f0a7-6696-4de9-b257-e6f13cd9137a)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:e8f0dd0b-9c3d-445e-8b96-c2fd16ab12e4](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:e8f0dd0b-9c3d-445e-8b96-c2fd16ab12e4)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:f7a59980-663f-4607-9b8a-faa91e1300ed](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:f7a59980-663f-4607-9b8a-faa91e1300ed)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:3d3648e6-635c-4371-9a6b-81e16643e53b](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:3d3648e6-635c-4371-9a6b-81e16643e53b)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:007ea371-828c-4e46-9941-d9486fa44b7a](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:007ea371-828c-4e46-9941-d9486fa44b7a)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:86d17dbe-2075-42d5-bb30-41f4bd555d98](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:86d17dbe-2075-42d5-bb30-41f4bd555d98)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:169aa5c7-e738-468f-96ce-2b15569890d9](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:169aa5c7-e738-468f-96ce-2b15569890d9)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:f99033ce-1b9a-4e75-bf3c-c2a4c93dc7f5](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:f99033ce-1b9a-4e75-bf3c-c2a4c93dc7f5)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:ba0ce9fc-ade9-47ed-8a9b-f70760339707](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:ba0ce9fc-ade9-47ed-8a9b-f70760339707)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:6af744e0-c4fa-4153-85a5-cb60a6bd6551](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:6af744e0-c4fa-4153-85a5-cb60a6bd6551)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:4c83d495-e0da-4aa7-94c3-8c4d184f5c56](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:4c83d495-e0da-4aa7-94c3-8c4d184f5c56)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service:80609802-f167-4fa9-b8f0-657f887a2dec](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceInventory/v4/service/urn:ngsi-ld:service:80609802-f167-4fa9-b8f0-657f887a2dec)

- **Type:** service
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:45 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to service (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

### serviceSpecification Objects

#### Object: [urn:ngsi-ld:service-specification:ab1f9684-e04f-4692-bb2d-20827f1bb759](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:ab1f9684-e04f-4692-bb2d-20827f1bb759)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:8f5f2d0c-9af4-47ad-a932-387455fc11df](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:8f5f2d0c-9af4-47ad-a932-387455fc11df)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:69a374ff-97c5-417b-8b31-2bd36798006b](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:69a374ff-97c5-417b-8b31-2bd36798006b)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:80b2ba93-5d5d-4753-a29f-b80114a01333](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:80b2ba93-5d5d-4753-a29f-b80114a01333)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:5ff52024-ea17-4f17-aa39-0499f57fc7d1](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:5ff52024-ea17-4f17-aa39-0499f57fc7d1)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:8a8741bb-c559-46a8-9ae8-d2f00c7504ca](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:8a8741bb-c559-46a8-9ae8-d2f00c7504ca)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:6ea578e6-ff2e-418b-993f-5169d271d78e](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:6ea578e6-ff2e-418b-993f-5169d271d78e)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:8ac3638e-1800-4530-b0a4-61bc73ae869a](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:8ac3638e-1800-4530-b0a4-61bc73ae869a)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:79ac0a54-ab0e-4018-a892-be1e54e8b2be](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:79ac0a54-ab0e-4018-a892-be1e54e8b2be)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:b8e690b3-3a24-4273-b569-7270128e1bc2](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:b8e690b3-3a24-4273-b569-7270128e1bc2)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:1e0f2ebf-e89a-4996-8ca8-e24eebbe3a7a](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:1e0f2ebf-e89a-4996-8ca8-e24eebbe3a7a)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:18a977ba-160a-4aef-a55f-ecff6a86e9a2](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:18a977ba-160a-4aef-a55f-ecff6a86e9a2)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:d612a81a-b154-4ca4-8ed0-28c13233e268](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:d612a81a-b154-4ca4-8ed0-28c13233e268)

- **Type:** serviceSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:89d6bddb-7805-4042-a60a-030ef09ed816](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:89d6bddb-7805-4042-a60a-030ef09ed816)

- **Type:** serviceSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:f70f5f90-14d2-4442-a459-dae184f233e4](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:f70f5f90-14d2-4442-a459-dae184f233e4)

- **Type:** serviceSpecification
- **Valid:** true
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:service-specification:7cda71b5-24b0-4c71-b650-749ca7615531](https://tmf.dome-marketplace-sbx.org/tmf-api/serviceCatalogManagement/v4/serviceSpecification/urn:ngsi-ld:service-specification:7cda71b5-24b0-4c71-b650-749ca7615531)

- **Type:** serviceSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:12 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to serviceSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)

### usageSpecification Objects

#### Object: [urn:ngsi-ld:usageSpecification:21333496-4652-4bbf-be85-a897278d4ee9](https://tmf.dome-marketplace-sbx.org/tmf-api/usageManagement/v4/usageSpecification/urn:ngsi-ld:usageSpecification:21333496-4652-4bbf-be85-a897278d4ee9)

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to usageSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:usageSpecification:24abecd5-bf1f-42e0-a34f-1657f39dffe1](https://tmf.dome-marketplace-sbx.org/tmf-api/usageManagement/v4/usageSpecification/urn:ngsi-ld:usageSpecification:24abecd5-bf1f-42e0-a34f-1657f39dffe1)

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to usageSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:usageSpecification:156985fe-69c9-4f2f-a906-cbdc01d4d427](https://tmf.dome-marketplace-sbx.org/tmf-api/usageManagement/v4/usageSpecification/urn:ngsi-ld:usageSpecification:156985fe-69c9-4f2f-a906-cbdc01d4d427)

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to usageSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:usageSpecification:3aacf05a-838b-4160-a810-4447fc58695e](https://tmf.dome-marketplace-sbx.org/tmf-api/usageManagement/v4/usageSpecification/urn:ngsi-ld:usageSpecification:3aacf05a-838b-4160-a810-4447fc58695e)

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to usageSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:usageSpecification:fd5da80e-fb29-45e8-8636-46b46dc2973a](https://tmf.dome-marketplace-sbx.org/tmf-api/usageManagement/v4/usageSpecification/urn:ngsi-ld:usageSpecification:fd5da80e-fb29-45e8-8636-46b46dc2973a)

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to usageSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:usageSpecification:e16ff3ad-0bf6-4b8c-95d9-0d2e69b4c12a](https://tmf.dome-marketplace-sbx.org/tmf-api/usageManagement/v4/usageSpecification/urn:ngsi-ld:usageSpecification:e16ff3ad-0bf6-4b8c-95d9-0d2e69b4c12a)

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to usageSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

#### Object: [urn:ngsi-ld:usageSpecification:7f9a65f1-3a08-43a0-bbc4-5ca8ec9f5ce5](https://tmf.dome-marketplace-sbx.org/tmf-api/usageManagement/v4/usageSpecification/urn:ngsi-ld:usageSpecification:7f9a65f1-3a08-43a0-bbc4-5ca8ec9f5ce5)

- **Type:** usageSpecification
- **Valid:** false
- **Timestamp:** 2025-11-21 06:42:47 UTC
- **Errors:**
  - relatedParty: Missing 'relatedParty' info for 'Seller' and 'SellerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
  - relatedParty: Missing 'relatedParty' info for 'Buyer' and 'BuyerOperator' role (Code: MISSING_RELATED_PARTY_INFO)
- **Warnings:**
  - @type: Recommended field '@type' is missing, setting it to usageSpecification (Code: MISSING_RECOMMENDED_FIELD)
  - version: Recommended field 'version' is missing (Code: MISSING_RECOMMENDED_FIELD)
  - lastUpdate: Recommended field 'lastUpdate' is missing (Code: MISSING_RECOMMENDED_FIELD)

---

*Report generated by TMForum Proxy Validator*
*Generated at: 2025-11-21 06:42:55 UTC*
