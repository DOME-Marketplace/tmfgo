package config

import "github.com/hesusruiz/tmforum/types"

// ******************************************************************
// DOME configurations
// ******************************************************************
//
// As this PDP is designed for DOME and ISBE environments, many config data items are configured in code.
// This avoids many configuration errors and simplifies deployment, at the expense of some flexibility.
// However, this flexibility is not really needed in practice, as the DOME environments are well defined and stable.
// Minimizing errors is here much more important than the ease to configure these parameters.

// The DOME SBX environment is used for development and testing, so we call it DOME_DEV
var domedevConfig = &Config{
	Environment: DOME_DEV,

	// In DOME, we act as a smart PDP and caching proxy, using the DOME TMF server as the real TMF server
	ProxyEnabled: true,

	// The server operator and admin is Altia
	// TODO: change this to the DOME Foundation, as soon as we have its VAT ID
	ServerOperatorOrganizationIdentifier: "VATES-A15456585",
	ServerOperatorDid:                    "did:elsi:VATES-A15456585",
	ServerOperatorName:                   "ALTIA CONSULTORES SA",
	ServerOperatorCountry:                "ES",

	// The LEAR has the power to Onboard
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
	Dbname:          "data/tmf.dome.sbx.db",
	ClonePeriod:     DefaultClonePeriod,
	Features: Features{
		OfferingLaunchOnlyByAdmin: false,
		GenerateIDOnCreate:        false,
		AllowIDInBody:             false,
		RetrieveLocalFirst:        true,
	},
}

// The DOME DEV2 environment is the pre-production environment
var domepreConfig = &Config{
	Environment:  DOME_PRE,
	ProxyEnabled: true,

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
	RemoteTMFServer: "https://tmf.dome-marketplace-dev2.org",
	VerifierServer:  "https://verifier.dome-marketplace-dev2.org",
	Dbname:          "data/tmf.dome.dev2.db",
	ClonePeriod:     DefaultClonePeriod,
	Features: Features{
		OfferingLaunchOnlyByAdmin: false,
		GenerateIDOnCreate:        false,
		AllowIDInBody:             false,
		RetrieveLocalFirst:        true,
	},
}

// The DOME PRO environment is the production environment
var domeproConfig = &Config{
	Environment:  DOME_PRO,
	ProxyEnabled: true,

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
	RemoteTMFServer: "https://tmf.dome-marketplace.eu",
	VerifierServer:  "https://verifier.dome-marketplace.eu",
	Dbname:          "data/tmf.dome.pro.db",
	ClonePeriod:     DefaultClonePeriod,
	Features: Features{
		OfferingLaunchOnlyByAdmin: false,
		GenerateIDOnCreate:        false,
		AllowIDInBody:             false,
		RetrieveLocalFirst:        true,
	},
}
