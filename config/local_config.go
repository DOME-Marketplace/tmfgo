package config

import "github.com/DOME-Marketplace/tmfgo/types"

// ******************************************************************
// Local developmentconfiguration
// ******************************************************************

// It is used to test against the DOME remote TMF server
var lclConfig = &Config{
	Environment:  LOCAL,
	ProxyEnabled: false,

	// The operator is Altia
	ServerOperatorOrganizationIdentifier: "VATES-A15456585",
	ServerOperatorDid:                    "did:elsi:VATES-A15456585",
	ServerOperatorName:                   "ALTIA CONSULTORES SA",
	ServerOperatorCountry:                "ES",

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
	Dbname:          "data/tmf.lcl.db",
	ClonePeriod:     DefaultClonePeriod,
	Features: Features{
		OfferingLaunchOnlyByAdmin: false,
		GenerateIDOnCreate:        false,
		AllowIDInBody:             false,
		VerifyJWTSignature:        false,
	},
}
