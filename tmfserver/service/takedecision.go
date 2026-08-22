package service

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/hesusruiz/tmforum/internal/errl"
	"github.com/hesusruiz/tmforum/internal/jpath"
	pdp "github.com/hesusruiz/tmforum/pdp"
	repo "github.com/hesusruiz/tmforum/tmfserver/repository"
	"github.com/hesusruiz/tmforum/types"
)

// checkAuthorization evaluates access authorization for a TMF request against both hardcoded and user-defined policies.
// It first checks hardcoded policies, and if they pass, proceeds to evaluate user policies in the PDP engine.
//
// Parameters:
//   - ruleEngine: PDP engine instance for evaluating user policies
//   - req: The request details containing auth and operation information
//   - tokenClaims: JWT claims from the authentication token
//   - objectMap: Map of TMF objects involved in the request
//
// Returns:
//   - authorized: boolean indicating if access is authorized
//   - err: error if any occurred during policy evaluation
//
// The function works in two stages:
// 1. Evaluates hardcoded policies first - if these fail, denies access immediately
// 2. If hardcoded policies pass, evaluates user-defined policies through the PDP engine
func (svc *Service) checkAuthorization(
	ruleEngine Authorizer,
	req *Request,
	objectMap repo.TMFObjectMap,
) (authorized bool, err error) {

	// Evaluate the hardcoded policies, if they fail return immediately.
	// Otherwise, continue to see if the user policies allow access
	_, err = svc.hardcodedPolicies(req, objectMap)
	if err != nil {
		return false, errl.Error(err)
	}

	// The caller is the owner, at least according to hardcoded policies.
	// The user policies will determine the final decision.
	req.AuthUser.IsOwner = true

	if err := svc.userPolicies(ruleEngine, req, req.AuthUser.TokenMap, objectMap); err != nil {
		return false, errl.Error(err)
	}

	return true, nil

}

func (svc *Service) isServerOperator(user types.AuthUser) bool {
	return repo.SameOrganizations(user.OrganizationIdentifier, svc.ServerOperatorDid)
}

func (svc *Service) isTrustedParty(user types.AuthUser) bool {
	for _, trustedParty := range svc.AdditionalTrustedparties {
		if repo.SameOrganizations(user.OrganizationIdentifier, trustedParty.Did) {
			return true
		}
	}

	return false
}

// hardcodedPolicies is the main entry point for pre-PDP policy enforcement.
// It matches the top-down, operation-first structure and logic of hardcodedPolicies.md:
// 1. CREATE
// 2. READ / LIST
// 3. UPDATE
// 4. DELETE
func (svc *Service) hardcodedPolicies(req *Request, obj repo.TMFObjectMap) (reason string, err error) {

	// Attribute integrity check (Seller/Buyer info pairs must be consistent)
	if err := validateAttributeIntegrity(obj); err != nil {
		return "", err
	}

	// Dispatch by operation top-down
	switch req.Action {
	case ActionCREATE:
		return svc.evalCreatePolicy(req, obj)
	case ActionREAD, ActionLIST:
		return svc.evalReadListPolicy(req, obj)
	case ActionUPDATE:
		return svc.evalUpdatePolicy(req, obj)
	case ActionDELETE:
		return svc.evalDeletePolicy(req, obj)
	default:
		return "", errl.Errorf("unsupported action %s", req.Action)
	}
}

// -----------------------------------------------------------------------------
// Operation Policy Evaluators matching hardcodedPolicies.md
// -----------------------------------------------------------------------------

// checkServerOperatorAccess checks if the caller is the server operator, and if it has the required powers for the action.
//
// Parameters:
//   - req: The request details
//
// Returns:
//   - authorized: boolean indicating if access is authorized
//   - reason: string indicating the reason for authorization
//   - err: error if any occurred during policy evaluation
func (svc *Service) checkServerOperatorAccess(req *Request) (bool, string, error) {
	caller := req.AuthUser
	if !svc.isServerOperator(caller) {
		return false, "", nil
	}

	if caller.IsLEAR {
		return true, fmt.Sprintf("caller %s is server operator and LEAR", caller.OrganizationIdentifier), nil
	}

	switch req.Action {
	case ActionREAD, ActionLIST:
		return true, fmt.Sprintf("caller %s is server operator (read/list)", caller.OrganizationIdentifier), nil
	case ActionCREATE:
		if svc.IsDOME() && !caller.ProductCreatePower {
			return false, "", errl.Errorf("caller %s is server operator but lacks create power", caller.OrganizationIdentifier)
		}
		return true, fmt.Sprintf("caller %s is server operator with create power", caller.OrganizationIdentifier), nil
	case ActionUPDATE:
		if !caller.ProductUpdatePower {
			return false, "", errl.Errorf("caller %s is server operator but lacks update power", caller.OrganizationIdentifier)
		}
		return true, fmt.Sprintf("caller %s is server operator with update power", caller.OrganizationIdentifier), nil
	case ActionDELETE:
		if !caller.ProductDeletePower {
			return false, "", errl.Errorf("caller %s is server operator but lacks delete power", caller.OrganizationIdentifier)
		}
		return true, fmt.Sprintf("caller %s is server operator with delete power", caller.OrganizationIdentifier), nil
	}
	return false, "", errl.Errorf("unsupported action %s", req.Action)
}

// CREATE policy implementation matching hardcodedPolicies.md
func (svc *Service) evalCreatePolicy(req *Request, obj repo.TMFObjectMap) (string, error) {
	caller := req.AuthUser
	if !caller.IsAuthenticated {
		return "", errl.Errorf("user not authenticated")
	}

	// Superuser Exception
	if isSO, reason, err := svc.checkServerOperatorAccess(req); isSO || err != nil {
		return reason, err
	}

	// Non-ServerOperator callers require ProductCreatePower
	if !caller.IsLEAR && !caller.ProductCreatePower {
		return "", errl.Errorf("caller %s lacks create power", caller.OrganizationIdentifier)
	}

	// Special Objects (Category, Organization, Individual)
	if isSpecialObject(obj) {
		switch strings.ToLower(obj.Type()) {
		case "category":
			return "", errl.Errorf("Category objects can only be created by server operator")
		case "organization":
			return "", errl.Errorf("Organization objects can only be created by server operator")
		case "individual":
			mandatorOrg := getIndividualIssuingAuthority(obj)
			if repo.SameOrganizations(mandatorOrg, caller.OrganizationIdentifier) {
				return fmt.Sprintf("caller %s is mandator in Individual object", caller.OrganizationIdentifier), nil
			}
			return "", errl.Errorf("caller %s is not mandator (%s) in Individual object", caller.OrganizationIdentifier, mandatorOrg)
		}
	}

	// On CREATE, auto-populate missing Seller/SellerOperator info
	if err := svc.ensureSellerInfoOnCreate(req, obj); err != nil {
		return "", err
	}

	objSeller, objSellerOperator, _ := obj.GetSellerInfo("")
	objBuyer, objBuyerOperator, _ := obj.GetBuyerInfo("")

	// Public objects: MUST be Seller or SellerOperator, and SellerOperator MUST be ServerOperator
	if obj.IsPotentiallyPublic() && objBuyer == "" {
		isSeller := repo.SameOrganizations(objSeller, caller.OrganizationIdentifier)
		isSellerOp := repo.SameOrganizations(objSellerOperator, caller.OrganizationIdentifier)
		if !isSeller && !isSellerOp {
			return "", errl.Errorf("caller %s is neither Seller %s nor SellerOperator %s for public object", caller.OrganizationIdentifier, objSeller, objSellerOperator)
		}
		if !repo.SameOrganizations(objSellerOperator, svc.ServerOperatorDid) {
			return "", errl.Errorf("SellerOperator %s must be server operator %s", objSellerOperator, svc.ServerOperatorDid)
		}
		return "caller is Seller/SellerOperator and server operator matches", nil
	}

	// Private objects: MUST be Seller, SellerOperator, Buyer, or BuyerOperator
	isSeller := repo.SameOrganizations(objSeller, caller.OrganizationIdentifier)
	isSellerOp := repo.SameOrganizations(objSellerOperator, caller.OrganizationIdentifier)
	isBuyer := repo.SameOrganizations(objBuyer, caller.OrganizationIdentifier)
	isBuyerOp := repo.SameOrganizations(objBuyerOperator, caller.OrganizationIdentifier)

	if !isSeller && !isSellerOp && !isBuyer && !isBuyerOp {
		return "", errl.Errorf("caller %s is not party involved in private object (Seller=%s, Buyer=%s)", caller.OrganizationIdentifier, objSeller, objBuyer)
	}

	// Operator associated with caller's role MUST be ServerOperator
	if (isSeller || isSellerOp) && repo.SameOrganizations(objSellerOperator, svc.ServerOperatorDid) {
		return "caller is Seller/SellerOperator on server operator node", nil
	}
	if (isBuyer || isBuyerOp) && repo.SameOrganizations(objBuyerOperator, svc.ServerOperatorDid) {
		return "caller is Buyer/BuyerOperator on server operator node", nil
	}

	return "", errl.Errorf("operator for caller's role does not match server operator %s", svc.ServerOperatorDid)
}

// READ / LIST policy implementation matching hardcodedPolicies.md
func (svc *Service) evalReadListPolicy(req *Request, obj repo.TMFObjectMap) (string, error) {
	caller := req.AuthUser

	// Special objects READ/LIST rules
	if isSpecialObject(obj) {
		switch strings.ToLower(obj.Type()) {
		case "category", "organization":
			// Public object types: readable by any caller (auth or unauth)
			return "public special object (category/organization)", nil
		case "individual":
			// Private object type: requires authentication & org match
			if !caller.IsAuthenticated {
				return "", errl.Errorf("user not authenticated")
			}
			if svc.isServerOperator(caller) {
				return "caller is server operator", nil
			}
			mandatorOrg := getIndividualIssuingAuthority(obj)
			if repo.SameOrganizations(mandatorOrg, caller.OrganizationIdentifier) {
				return fmt.Sprintf("caller %s is mandator in Individual object", caller.OrganizationIdentifier), nil
			}
			return "", errl.Errorf("caller %s is not mandator in Individual object", caller.OrganizationIdentifier)
		}
	}

	objSeller, objSellerOperator, _ := obj.GetSellerInfo("")
	objBuyer, objBuyerOperator, _ := obj.GetBuyerInfo("")

	// Public objects: If object is public and has no buyer, allow access, even to unauthenticated users.
	// But only if its lifecycleStatus is "Launched", which means the Seller wants it to be public
	if obj.IsPotentiallyPublic() && objBuyer == "" && strings.ToLower(obj.LifecycleStatus()) == "launched" {
		return "access to launched public resource", nil
	}

	// From now on, the object is either:
	// - a public aobject not yet launched and without buyer information
	// - a public object with buyer information (i.e. buyer has a buyer role)
	// - a private object (like an agreement) that should be accessible only by the parties involved

	// Private objects require authentication
	if !caller.IsAuthenticated {
		return "", errl.Errorf("user not authenticated")
	}

	// If caller is a server operator with proper powers, allow access
	if isSO, reason, err := svc.checkServerOperatorAccess(req); isSO || err != nil {
		return reason, err
	}

	// Private objects: caller's organization must match Seller, SellerOperator, Buyer, or BuyerOperator
	isParty := repo.SameOrganizations(objSeller, caller.OrganizationIdentifier) ||
		repo.SameOrganizations(objSellerOperator, caller.OrganizationIdentifier) ||
		repo.SameOrganizations(objBuyer, caller.OrganizationIdentifier) ||
		repo.SameOrganizations(objBuyerOperator, caller.OrganizationIdentifier)

	if isParty {
		return "caller is party involved in private object", nil
	}

	return "", errl.Errorf("caller %s is not party involved in private object", caller.OrganizationIdentifier)
}

// UPDATE policy implementation matching hardcodedPolicies.md
func (svc *Service) evalUpdatePolicy(req *Request, obj repo.TMFObjectMap) (string, error) {
	caller := req.AuthUser
	if !caller.IsAuthenticated {
		return "", errl.Errorf("user not authenticated")
	}

	// Superuser Exception
	if isSO, reason, err := svc.checkServerOperatorAccess(req); isSO || err != nil {
		return reason, err
	}

	// Non-ServerOperator callers require ProductUpdatePower
	if !caller.IsLEAR && !caller.ProductUpdatePower {
		return "", errl.Errorf("caller %s lacks update power", caller.OrganizationIdentifier)
	}

	// Special Objects
	if isSpecialObject(obj) {
		switch strings.ToLower(obj.Type()) {
		case "category":
			return "", errl.Errorf("Category objects can only be updated by server operator")
		case "organization":
			return "", errl.Errorf("Organization objects can only be updated by server operator")
		case "individual":
			mandatorOrg := getIndividualIssuingAuthority(obj)
			if repo.SameOrganizations(mandatorOrg, caller.OrganizationIdentifier) {
				return fmt.Sprintf("caller %s is mandator in Individual object", caller.OrganizationIdentifier), nil
			}
			return "", errl.Errorf("caller %s is not mandator in Individual object", caller.OrganizationIdentifier)
		}
	}

	objSeller, objSellerOperator, _ := obj.GetSellerInfo("")
	objBuyer, objBuyerOperator, _ := obj.GetBuyerInfo("")

	// Public objects: MUST be Seller or SellerOperator, and SellerOperator MUST be ServerOperator
	if obj.IsPotentiallyPublic() && objBuyer == "" {
		isSeller := repo.SameOrganizations(objSeller, caller.OrganizationIdentifier)
		isSellerOp := repo.SameOrganizations(objSellerOperator, caller.OrganizationIdentifier)
		if !isSeller && !isSellerOp {
			return "", errl.Errorf("caller %s is neither Seller %s nor SellerOperator %s for public object", caller.OrganizationIdentifier, objSeller, objSellerOperator)
		}
		if !repo.SameOrganizations(objSellerOperator, svc.ServerOperatorDid) {
			return "", errl.Errorf("SellerOperator %s must be server operator %s", objSellerOperator, svc.ServerOperatorDid)
		}
		return "caller is Seller/SellerOperator and server operator matches", nil
	}

	// Private objects: MUST be Seller, SellerOperator, Buyer, or BuyerOperator
	isSeller := repo.SameOrganizations(objSeller, caller.OrganizationIdentifier)
	isSellerOp := repo.SameOrganizations(objSellerOperator, caller.OrganizationIdentifier)
	isBuyer := repo.SameOrganizations(objBuyer, caller.OrganizationIdentifier)
	isBuyerOp := repo.SameOrganizations(objBuyerOperator, caller.OrganizationIdentifier)

	if !isSeller && !isSellerOp && !isBuyer && !isBuyerOp {
		return "", errl.Errorf("caller %s is not party involved in private object", caller.OrganizationIdentifier)
	}

	if (isSeller || isSellerOp) && repo.SameOrganizations(objSellerOperator, svc.ServerOperatorDid) {
		return "caller is Seller/SellerOperator on server operator node", nil
	}
	if (isBuyer || isBuyerOp) && repo.SameOrganizations(objBuyerOperator, svc.ServerOperatorDid) {
		return "caller is Buyer/BuyerOperator on server operator node", nil
	}

	return "", errl.Errorf("operator for caller's role does not match server operator %s", svc.ServerOperatorDid)
}

// DELETE policy implementation matching hardcodedPolicies.md
func (svc *Service) evalDeletePolicy(req *Request, obj repo.TMFObjectMap) (string, error) {
	caller := req.AuthUser
	if !caller.IsAuthenticated {
		return "", errl.Errorf("user not authenticated")
	}

	// Superuser Exception
	if isSO, reason, err := svc.checkServerOperatorAccess(req); isSO || err != nil {
		return reason, err
	}

	// Non-ServerOperator callers require ProductDeletePower
	if !caller.IsLEAR && !caller.ProductDeletePower {
		return "", errl.Errorf("caller %s lacks delete power", caller.OrganizationIdentifier)
	}

	// Special Objects
	if isSpecialObject(obj) {
		switch strings.ToLower(obj.Type()) {
		case "category":
			return "", errl.Errorf("Category objects can only be deleted by server operator")
		case "organization":
			return "", errl.Errorf("Organization objects can only be deleted by server operator")
		case "individual":
			mandatorOrg := getIndividualIssuingAuthority(obj)
			if repo.SameOrganizations(mandatorOrg, caller.OrganizationIdentifier) {
				return fmt.Sprintf("caller %s is mandator in Individual object", caller.OrganizationIdentifier), nil
			}
			return "", errl.Errorf("caller %s is not mandator in Individual object", caller.OrganizationIdentifier)
		}
	}

	objSeller, objSellerOperator, _ := obj.GetSellerInfo("")
	objBuyer, objBuyerOperator, _ := obj.GetBuyerInfo("")

	// Public objects: MUST be Seller or SellerOperator, and SellerOperator MUST be ServerOperator
	if obj.IsPotentiallyPublic() && objBuyer == "" {
		isSeller := repo.SameOrganizations(objSeller, caller.OrganizationIdentifier)
		isSellerOp := repo.SameOrganizations(objSellerOperator, caller.OrganizationIdentifier)
		if !isSeller && !isSellerOp {
			return "", errl.Errorf("caller %s is neither Seller %s nor SellerOperator %s for public object", caller.OrganizationIdentifier, objSeller, objSellerOperator)
		}
		if !repo.SameOrganizations(objSellerOperator, svc.ServerOperatorDid) {
			return "", errl.Errorf("SellerOperator %s must be server operator %s", objSellerOperator, svc.ServerOperatorDid)
		}
		return "caller is Seller/SellerOperator and server operator matches", nil
	}

	// Private objects: MUST be Seller, SellerOperator, or BuyerOperator (Buyer directly is REJECTED)
	isSeller := repo.SameOrganizations(objSeller, caller.OrganizationIdentifier)
	isSellerOp := repo.SameOrganizations(objSellerOperator, caller.OrganizationIdentifier)
	isBuyerOp := repo.SameOrganizations(objBuyerOperator, caller.OrganizationIdentifier)

	if !isSeller && !isSellerOp && !isBuyerOp {
		return "", errl.Errorf("caller %s is not Seller, SellerOperator, or BuyerOperator for private object delete", caller.OrganizationIdentifier)
	}

	if (isSeller || isSellerOp) && repo.SameOrganizations(objSellerOperator, svc.ServerOperatorDid) {
		return "caller is Seller/SellerOperator on server operator node", nil
	}
	if isBuyerOp && repo.SameOrganizations(objBuyerOperator, svc.ServerOperatorDid) {
		return "caller is BuyerOperator on server operator node", nil
	}

	return "", errl.Errorf("operator for caller's role does not match server operator %s", svc.ServerOperatorDid)
}

// -----------------------------------------------------------------------------
// Pipeline Step Helpers
// -----------------------------------------------------------------------------

// Step 1: Ensure Seller/SellerOperator and Buyer/BuyerOperator pairs are not partially set.
// Also, reject when only buyer info is set, as we need always the seller.
func validateAttributeIntegrity(obj repo.TMFObjectMap) error {
	objSeller, objSellerOperator, _ := obj.GetSellerInfo("")
	objBuyer, objBuyerOperator, _ := obj.GetBuyerInfo("")

	if isPartiallySet(objSeller, objSellerOperator) {
		return errl.Errorf("objSeller and objSellerOperator must both be set or both be empty, got objSeller='%s', objSellerOperator='%s'", objSeller, objSellerOperator)
	}
	if isPartiallySet(objBuyer, objBuyerOperator) {
		return errl.Errorf("objBuyer and objBuyerOperator must both be set or both be empty, got objBuyer='%s', objBuyerOperator='%s'", objBuyer, objBuyerOperator)
	}
	if objBuyer != "" && objSeller == "" {
		return errl.Errorf("objBuyer and objBuyerOperator must both be set or both be empty, got objBuyer='%s', objBuyerOperator='%s'", objBuyer, objBuyerOperator)
	}
	return nil
}

// Helper to check if object is a special type (organization, individual, category).
func isSpecialObject(obj repo.TMFObjectMap) bool {
	objType := obj.Type()
	switch strings.ToLower(objType) {
	case "organization", "individual", "category":
		return true
	default:
		return false
	}
}

// Helper to extract issuingAuthority from Individual object.
func getIndividualIssuingAuthority(obj repo.TMFObjectMap) string {
	for _, item := range jpath.GetList(obj, "individualIdentification") {
		m, ok := item.(map[string]any)
		if ok && m["identificationType"] == "learcredentialemployee" {
			issuingAuth, ok := m["issuingAuthority"].(string)
			if ok {
				return issuingAuth
			}
		}
	}
	return ""
}

// Auto-assign Seller info if omitted on CREATE.
func (svc *Service) ensureSellerInfoOnCreate(req *Request, obj repo.TMFObjectMap) error {
	if isSpecialObject(obj) {
		return nil
	}
	objSeller, _, _ := obj.GetSellerInfo("")
	if objSeller != "" {
		return nil
	}
	caller := req.AuthUser
	if err := obj.SetSellerInfo(svc.ServerOperatorDid, caller.OrganizationIdentifier, "v4"); err != nil {
		return errl.Errorf("error setting seller info: %w", err)
	}
	return nil
}

// userPolicies constructs the policy input for a PDP (policy decision point)
// from the provided request, token claims, TMF object map and authenticated user,
// then evaluates the assembled input against the rules engine to determine
// whether the request is authorized.
//
// The assembled input contains the following top-level keys:
//   - "request": a dereferenced map representation of req (req.ToMap())
//   - "token": a dereferenced map representation of tokenClaims
//   - "tmf": a dereferenced map representation of the incoming TMF object map
//   - "user": a dereferenced map representation of the authenticated user
//
// The function will pretty-print the assembled input as JSON when the default
// slog logger is at debug level to aid debugging.
//
// Decision & error semantics:
//   - If ruleEngine is nil the function defaults to allowing the request (returns nil).
//   - If ruleEngine is present, ruleEngine.Authorize(input) is invoked.
//   - Any error returned by the PDP is treated as a rejection and returned to the caller.
//   - If the PDP returns false the request is considered denied and a non-nil error is returned.
//   - If the PDP returns true the request is authorized and the function returns nil.
//
// Side effects:
//   - Logs an informational message on authorization and prints debug JSON when enabled.
//   - Constructs a single "input" object (OPA-style) so policies can access request,
//     token, tmf and user together and support callbacks.
//
// Return value:
//   - nil when the request is authorized.
//   - non-nil error when the rules engine rejects the request or an internal error occurs.
func (svc *Service) userPolicies(
	ruleEngine Authorizer,
	req *Request,
	tokenClaims map[string]any,
	objectMap repo.TMFObjectMap,
) (err error) {

	userArgument := pdp.StarTMFMap(req.AuthUser.ToMap())
	incomingObjectArgument := pdp.StarTMFMap(objectMap)
	requestArgument := pdp.StarTMFMap(req.ToMap())
	tokenArgument := pdp.StarTMFMap(tokenClaims)

	// Assemble all data in a single input argument, to the style of OPA.
	input := map[string]any{
		"request": requestArgument,
		"token":   tokenArgument,
		"tmf":     incomingObjectArgument,
		"user":    userArgument,
	}

	decision := true
	if ruleEngine != nil {
		decision, err = ruleEngine.Authorize(input)

		// An error is considered a rejection
		if err != nil {
			return errl.Errorf("rules engine rejected request due to an error: %w", err)
		}
	}

	// The rules engine rejected the request
	if !decision {
		return errl.Errorf("rules engine rejected request due to policy")
	}

	// The rules engine accepted the request
	slog.Debug("rules engine accepted request")
	return nil
}

// isPartiallySet returns true if exactly one of the fields is empty.
func isPartiallySet(a, b string) bool {
	return (a == "") != (b == "")
}
