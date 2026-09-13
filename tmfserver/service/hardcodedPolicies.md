# Rules for the hardcoded policies.

## Introduction, concepts and definitions

The Access Control Component protects the TM Forum APIs that manage the objects in a TM Forum database. The combination of the TM Forum database, the TM Forum APIs and Access Control Component are typically deployed and operated together by the same organization and the resulting system is called an Access Node (AN). We say that the AN exposes the TM Forum APIs, which can be consumed by the same organization or by other organizations than the one operating the AN.

The DOME ecosystem is composed of a set of ANs operated by different organizations that replicate the TM Forum objects among themselves according to some replication policies (which are not the subject of this document).

### Participating entities

There are different types of organizations and Access Nodes in the DOME ecosystem, depending on the role of the organization operating the AN.

The types of organizations are Marketplace Operators, Sellers and Buyers.

- **Marketplace Operators**: each Marketplace Operator owns its Sellers and Buyers and has its own onboarding process and mechanism to identify and authenticate its Seller and Buyer organizations. To connect to the DOME ecosystem, each Marketplace Operator runs one or more ANs that replicate TMF Objects with other ANs in the ecosystem. In DOME, since the TMF Objects correspond to the products sold or bought by the Sellers and Buyers, we say that the Marketplace operates the AN "on behalf" of the Sellers and Buyers of the Marketplace Operator. Typically, the Marketplace Operator is the one using the TMF APIs provided by its AN, via integration with the IT systems of the Marketplace Operator.

- **Sellers and Buyers**: they are the entities that buy and sell products and services to each other, via a Marketplace Operator that they use.

For the purposes of this document, we distinguish between two types of Marketplace Operators which we will call colloquially the "classical" and the "DOME way".

A "classical" Marketplace Operator uses whatever marketplace software they want to provide the functionality of a marketplace to its Sellers and Buyers, and operates one or more ANs to connect to the DOME ecosystem. The TM Forum APIs exposed by the ANs are used by the IT systems of the Marketplace Operator in a transparent way to the Sellers and Buyers. The Marketplace Operator manages the identities of the organizations in the way that they wish, and those will be translated to the ones used in the DOME ecosystem in a way that will be described later in this document.

The other type of Marketplace Operator works in the "DOME way". One of these marketplaces is the DOME Marketplace (https://dome-marketplace.eu) operated by the DOME Foundation, but there may be other independent marketplaces operating in the same way.

The DOME Marketplace operated by the DOME Foundation is special in the sense that the DOME Foundation provides the governance of the whole decentralized ecosystem, but for the purposes of this document it is equal to any other marketplace operating the "DOME way".

In this mode, Sellers and Buyers manage the products and services they sell or buy in several ways, depending on their choice:

- Manually, via the user interface (screens) provided by the Marketplace Operator where they have onboarded. The Marketplace Operator operates the ANs that connect to the DOME ecosystem and its Seller and Buyer customers do not know anything about the TMF APIs and the ANs.

- Programatically, via the TM Forum APIs. They have integrated their IT systems with the TM Forum APIs of an Access Node.

In this last programmatic model, there are two possibilities:

- The Seller or Buyer can operate its own instance of AN, and the AN replicates objects with the other ANs in the ecosystem.

- The Seller/Buyer uses the TM Forum APIs provided by the AN operated by its Marketplace Operator. This is the easiest technical model for a Seller/Buyer, but it implies a higher level of dependency from the Marketplace Operator.

### The hardcoded policies and the PDP

The Access Control Component of the AN enforces a set of authorization policies. Some of them are "hardcoded" in the sense that they can not be modified by the user and are the same for all instances of the AN. In addition to the hardcoded policies, the AN also allows the definition and enforcement of user-defined policies.

This document covers only the hardcoded policies for the TMF server. They are checked before the user-defined policies.

The core component of the Access Control Component which evaluates the policies and makes the decision to allow or deny access is called the PDP (Policy Decision Point). The Access Control Component also implements the PEP (Policy Enforcement Point), which intercepts all requests before they reach the TM Forum APIs and asks the PDP for a decision.

The PDP uses the following information to evaluate the policies (received by the `hardcodedPolicies()` function in the code):

- The incoming request received from the user. It contains information about the authenticated user, the action to perform, the resource name, the ID, the query parameters, and the body.
- The TMF object that is being accessed, for all requests except a CREATE. When creating a new object the PDP receives the object to create in the body of the request.

### The Request object

The request object (`req`) received by the PDP contains the following information:
- `method`: HTTP method (READ, CREATE, UPDATE, DELETE, LIST)
- `action`: TMF action, which is a more semantic representation of the action to be performed (READ, CREATE, UPDATE, DELETE, LIST)
- `apiFamily`: API family (TMF620, TMF629, etc.)
- `apiVersion`: API version (v4, v5, etc.)
- `resourceName`: name of the resource (productOffering, productSpecification, etc.)
- `id`: identifier of the resource, for the actions that include it (CREATE does not carry an identifier as it is created by the server)
- `queryParams`: query parameters
- `body`: request body in raw format, for the actions that require it.
- `authUser`: information about the authenticated user, or empty if not authenticated.

### The authenticated User

The `req` object includes the `authUser` field which contains information about the authenticated user which makes the request. We will call it the `caller`. A caller can be a human or a machine, acting on behalf of an organization. We support both types.

The relevant fields in `authUser` are as follows.

`IsAuthenticated` indicates if the caller is authenticated. If false, the rest of the fields are not relevant.

`OrganizationIdentifier` is the unique identifier of the organization on behalf of which the caller is acting.

The caller acts on behalf of the organization, with a subset of the powers that the organization has, in the form of delegated powers. The relevant powers for TMF APIs are `ProductCreatePower`, `ProductUpdatePower` and `ProductDeletePower`.

A caller that has full powers to act in DOME on behalf of an organization is called a LEAR (Legal Entity Authorised Representative), and its powers include the `ProductCreatePower`, `ProductUpdatePower` and `ProductDeletePower` powers (among others).

The `IsOwner` flag is a precomputed field that indicates that the organization owns the object being modified. It is there to facilitate the logic when enforcing the policies.

### The Object Map

The `objectMap` object is a pointer to a TMFObjectMap struct, which is a map of TMF objects.

For a CREATE action, `objectMap` is the incoming object map that the user sent with the request.
For the other actions, `objectMap` contains the existing object in the repository on which the action is to be performed.

### Types of TMF Objects

There are two types of objects according to their default visibility property:
- Public type: these objects are readable/listable by all users, even unauthenticated ones.
- Non-public type: these objects are only accessible to authenticated users, and according to the policies on ownership of the objects. This can happen if the seller or buyer information is present in the object.

By default, the TMF objects which are of a public type are the ones related to the publishing of the commercial offers of products and services, like Catalog, Category, ProductOffering, ProductSpecification, ProductOfferingPrice, etc.

However, sometimes a Seller may create a ProductOffering specific to a given Buyer, as it may happen in negotiated deals. In this case, the Buyer information is present in the object, and as such, the object is not public anymore.

The Organization object is also a public object type.

Most other object types in DOME are private, because they are associated to the commercial activity of the participants in the marketplace. Private objects are always associated to a Seller and a Buyer, which are the participants in the transaction.

There are three types of objects which are special:
- Category: is a public object which can never be private. In DOME, categories are defined by the DOME Foundation and are common to all participants in the ecosystem.
- Organization: is a public object which can never be private. They are created by the Marketplace Operator and represent the participant organizations. The information in the Organization object is public, as it appears in public business registries or other public sources.
- Individual: is a private object which represents a natural person as an employee of a given Organization. The Individual object is always associated to an Organization, the employer, and is private for obvious reasons.

### Ownership information inside each object

Most objects have attributes that determine its relationship with one or more organizations, and these relationships are relevant for policy enforcement.

**Seller** and **SellerOperator**: They must exist in all non-special objects, both public and private. For public objects the Seller attribute is the organization that created the object. When a Seller organization is using the infrastructure of a Marketplace Operator, the SellerOperator attribute is the organization that is providing the infrastructure. This is important because the SellerOperator operates on behalf of the Seller, and the actual requests that are received by the Access Control component carry the authentication data of the SellerOperator.

**Buyer** and **BuyerOperator**: These fields must exist only in private objects (except in Individual objects). Private objects are the ones created for (or created by) a specific customer (Buyer). For example, when creating and managing an order for a given ProductOffering. The private objects are only accessible by the parties involved in the transaction: the Buyer and the Seller and the respective Marketplace Operators (BuyerOperator and SellerOperator). The actual actions that can be performed by the user sending the request are defined in the sections below.

If a Seller operates itself the infrastructure that is making the requests, the Seller and SellerOperator fields are the same and contain the OrganizationIdentifier of the Seller.

Similarly, if a Buyer operates itself the infrastructure that is making the requests, the Buyer and BuyerOperator fields are the same and contain the OrganizationIdentifier of the Buyer.

### Public and private object instances

A public object is one which is an instance of a public type (eg. ProductOffering) and does not have Buyer info. The public object types are those identified in tmf_operations.yaml with the attribute 'public: true'. There is also the method `IsPotentiallyPublic()` method in each object to get that information at runtime.

Examples of public types are:
- Catalog
- ProductOffering
- ProductSpecification
- ResourceSpecification
- ServiceSpecification
- ProductOfferingPrice
- ProductOfferingTerm

Category and Organization are types with public visibility but with access control rules that are special and are described in its own section.

Another way to say look at a public object instance: An object can be of a public type but not be public if it has Buyer or BuyerOperator attributes set. In that case, it is a private object.
An object which is an instance of a non-public type is always private, even if it does not include Buyer information.

### The ServerOperator

DOME is a federated ecosystem, composed of multiple Marketplace operators interacting with each other, and possibly Sellers and Buyers operating their own infrastructure and making requests to us. The words "making requests to us" are critical: we need to identify not only the user and the organization making the requests but also who are we. The Access Control component is assumed to be operated by an organization that may operate on behalf of itself or on behalf of other organizations, like the DOME Marketplace instance which is operated by the DOME Foundation. The Access Control component has a parameter called the ServerOperator which is the OrganizationIdentifier of the organization operating the component and the TM Forum database.

The ServerOperator has by definition full access to all objects in its server. However, the entities (employees or machines) acting on behalf of the ServerOperator can only perform the actions that the ServerOperator organization has delegated to them. This is elaborated in the policies definition below.

### Marketplace Operators

The organizations that can operate on behalf of other organizations (Marketplace Operators), have to be onboarded in DOME with a special process and are registered in a specific Trusted List. This list is accessible by the Access Control component with the `isTrustedParty()` function.

When the Seller and SellerOperator attributes in an object are different, it means that the object is managed by the SellerOperator on behalf of the Seller, and we have to apply special policies to requests that target that object. The same logic applies to the Buyer and BuyerOperator attributes. In both cases, the SellerOperator and BuyerOperator must be in the Trusted Participants List.

When the Seller and SellerOperator attributes are the same, it means that the Seller operates its own infrastructure, and it is not using a Marketplace Operator. In this case, the Seller does not need to be in the Trusted Participants List. The same logic applies to the Buyer and BuyerOperator: if the Buyer and BuyerOperator are the same, it means that the Buyer operates its own infrastructure, and it is not using a Marketplace Operator. In this case, the Buyer does not need to be in the Trusted Participants List.

### The "home" of an object

Objects have the concept of "home". The "home" of an object is the server or servers that created the object. For example, if a Seller creates a ProductOffering in a server operated by a Marketplace Operator, then the Marketplace Operator is the "home" of the object. The ProductOffering can be replicated and be read in many other servers, but the ProductOffering is "owned" by the home server. Policies treat the home server as having special privileges on the object.

## Policies Definition

This section describes the hardcoded policies organized from an operation-first (top-down) perspective.

The policies are structured into four main operations, in this order:
1. **CREATE**
2. **READ / LIST**
3. **UPDATE**
4. **DELETE**

Under each operation, the policies are grouped by object scope: **Public objects**, **Private objects**, and **Special objects** (`Category`, `Organization`, and `Individual`).

We deliberately keep some redundancy in the policy definitions across the different operations and object types, to make the policies easier to understand and so easier to maintain and keep correct.

**Superuser Exception:** If the caller is the `ServerOperator`, it can perform any action on any object on that server with any non-empty values for `Seller` and `SellerOperator`, or `Buyer` and `BuyerOperator`. The caller must also have the corresponding power delegated by its organization to perform the action, except that READ/LIST do not require powers.

---

### CREATE

**Caller Authentication & Powers:** The caller must be authenticated and have at least the `ProductCreatePower` power delegated by its organization.

#### Public objects

- The caller MUST be one of: `Seller` or `SellerOperator`. Any caller outside these roles is rejected.

- The `SellerOperator` MUST be the same as the `ServerOperator`. More explicitly:
  - If the caller is `Seller`, the `SellerOperator` must be the `ServerOperator`.
  - If the caller is `SellerOperator`, it must be the same as the `ServerOperator`.

- **Special case when Seller info is omitted:** If the incoming object does not include `Seller` and `SellerOperator` attributes, the Access Control component automatically sets `Seller` to the caller and `SellerOperator` to the `ServerOperator`.

#### Private objects

As described above, a private object is one of a private type, or one of a public type which includes Buyer information.

- The caller MUST be one of: `Seller`, `SellerOperator`, `Buyer`, or `BuyerOperator`. Any caller outside these roles is rejected. If buyer information is not available, it is assumed that the caller matches neither the `Buyer` nor the `BuyerOperator`.

- The operator associated with the caller's role (`SellerOperator` for Seller-side callers, or `BuyerOperator` for Buyer-side callers) MUST be the same as the `ServerOperator`. More explicitly:
  - If the caller is `Seller`, the `SellerOperator` must be the `ServerOperator`.
  - If the caller is `SellerOperator`, it must be the same as the `ServerOperator`.
  - If the caller is `Buyer`, the `BuyerOperator` must be the `ServerOperator`.
  - If the caller is `BuyerOperator`, it must be the same as the `ServerOperator`.

#### Special objects (Category, Organization, Individual)

- **Category:** Can only be created by the `ServerOperator`. The caller must be the `ServerOperator`.

- **Organization:** Can only be created by the `ServerOperator`. The caller must be the `ServerOperator`.

- **Individual:** The caller must be the same as the organization of the Individual.

---

### READ / LIST

#### Public objects

**Caller Authentication & Powers:** There is no need for authentication. Any caller can perform a READ or LIST operation on a public object.
But the lifecycleStatus of the object must be in the "Launched" state to be publicly accessible.

#### Private objects

**Caller Authentication & Powers:** The caller MUST be authenticated.

The caller's organization must match at least one of the parties involved in the object: `Seller`, `SellerOperator`, `Buyer`, or `BuyerOperator`.

#### Special objects (Category, Organization, Individual)

- **Category:** Being a public object type, `Category` objects can be read or listed by any caller (authenticated or unauthenticated).

- **Organization:** Being a public object type, `Organization` objects can be read or listed by any caller (authenticated or unauthenticated).

- **Individual:** Being a private object type, `Individual` objects require authentication. The caller must belong to the organization associated to the Individual.

---

### UPDATE

**Caller Authentication & Powers:** The caller must be authenticated and have at least the `ProductUpdatePower` power delegated by its organization.

#### Public objects

- The caller MUST be one of: `Seller` or `SellerOperator`. Any caller outside these roles is rejected.

- The `SellerOperator` MUST be the same as the `ServerOperator`. More explicitly:
  - If the caller is `Seller`, the `SellerOperator` must be the `ServerOperator`.
  - If the caller is `SellerOperator`, it must be the same as the `ServerOperator`.

#### Private objects

As described above, a private object is one of a private type, or one of a public type which includes Buyer information.

- The caller MUST be one of: `Seller`, `SellerOperator`, `Buyer`, or `BuyerOperator`. Any caller outside these roles is rejected. If buyer information is not available, it is assumed that the caller matches neither the `Buyer` nor the `BuyerOperator`.

- The operator associated with the caller's role (`SellerOperator` for Seller-side callers, or `BuyerOperator` for Buyer-side callers) MUST be the same as the `ServerOperator`. More explicitly:
  - If the caller is `Seller`, the `SellerOperator` must be the `ServerOperator`.
  - If the caller is `SellerOperator`, it must be the same as the `ServerOperator`.
  - If the caller is `Buyer`, the `BuyerOperator` must be the `ServerOperator`.
  - If the caller is `BuyerOperator`, it must be the same as the `ServerOperator`.

#### Special objects (Category, Organization, Individual)

- **Category:** Can only be updated by the `ServerOperator`. The caller must be the `ServerOperator`.

- **Organization:** Can only be updated by the `ServerOperator`. The caller must be the `ServerOperator`.

- **Individual:** The caller must be the same as the organization of the Individual.

---

### DELETE

**Caller Authentication & Powers:** The caller must be authenticated and have at least the `ProductDeletePower` power delegated by its organization.

#### Public objects

- The caller MUST be one of: `Seller` or `SellerOperator`. Any caller outside these roles is rejected.

- The `SellerOperator` MUST be the same as the `ServerOperator`. More explicitly:
  - If the caller is `Seller`, the `SellerOperator` must be the `ServerOperator`.
  - If the caller is `SellerOperator`, it must be the same as the `ServerOperator`.

#### Private objects

As described above, a private object is one of a private type, or one of a public type which includes Buyer information.

- The caller MUST be one of: `Seller`, `SellerOperator`, or `BuyerOperator`. Any caller outside these roles (including the `Buyer` directly) is rejected. If buyer information is not available, it is assumed that the caller matches neither the `Buyer` nor the `BuyerOperator`.

- The operator associated with the caller's role (`SellerOperator` for Seller-side callers, or `BuyerOperator` for Buyer-side callers) MUST be the same as the `ServerOperator`. More explicitly:
  - If the caller is `Seller`, the `SellerOperator` must be the `ServerOperator`.
  - If the caller is `SellerOperator`, it must be the same as the `ServerOperator`.
  - If the caller is `BuyerOperator`, it must be the same as the `ServerOperator`.

*(Note: The `Buyer` cannot delete private objects directly, but can do so through its `BuyerOperator`, the `Seller`, or the `SellerOperator`).*

#### Special objects (Category, Organization, Individual)

- **Category:** Can only be deleted by the `ServerOperator`. The caller must be the `ServerOperator`.

- **Organization:** Can only be deleted by the `ServerOperator`. The caller must be the `ServerOperator`.

- **Individual:** The caller must be the same as the organization of the Individual.
