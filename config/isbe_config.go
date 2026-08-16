package config

import "github.com/hesusruiz/tmforum/types"

// ******************************************************************
// ISBE configurations
// ******************************************************************
//
// As this PDP is designed for DOME and ISBE environments, many config data items are configured in code.
// This avoids many configuration errors and simplifies deployment, at the expense of some flexibility.
// However, this flexibility is not really needed in practice, as the DOME environments are well defined and stable.
// Minimizing errors is here much more important than the ease to configure these parameters.

var isbedevConfig = &Config{
	Environment:  ISBE_DEV,
	ProxyEnabled: false,

	// The operator is Alastria
	ServerOperatorOrganizationIdentifier: "VATES-G87936159",
	ServerOperatorDid:                    "did:elsi:VATES-G87936159",
	ServerOperatorName:                   "Alastria",
	ServerOperatorCountry:                "ES",

	LEARPower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "Onboarding",
		Action:   []string{"execute"},
	},
	ProductCreatePower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "ProductOffering",
		Action:   []string{"Create"},
	},
	ProductUpdatePower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "ProductOffering",
		Action:   []string{"Update"},
	},
	ProductDeletePower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "ProductOffering",
		Action:   []string{"Delete"},
	},

	PolicyFileName: "auth_policies.star",
	VerifierServer: "https://idp.dev.cloud-w.envs.redisbe.com/auth/realms/dev-isbe",

	Dbname:      "data/isbetmf.db",
	ClonePeriod: DefaultClonePeriod,
	Features: Features{
		OfferingLaunchOnlyByAdmin: true,
		GenerateIDOnCreate:        true,
		AllowIDInBody:             true,
		VerifyJWTSignature:        true,
	},
}

var isbepreConfig = &Config{
	Environment:  ISBE_PRE,
	ProxyEnabled: false,

	// The operator is Alastria
	ServerOperatorOrganizationIdentifier: "VATES-G87936159",
	ServerOperatorDid:                    "did:elsi:VATES-G87936159",
	ServerOperatorName:                   "Alastria",
	ServerOperatorCountry:                "ES",

	LEARPower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "Onboarding",
		Action:   []string{"execute"},
	},
	ProductCreatePower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "ProductOffering",
		Action:   []string{"Create"},
	},
	ProductUpdatePower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "ProductOffering",
		Action:   []string{"Update"},
	},
	ProductDeletePower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "ProductOffering",
		Action:   []string{"Delete"},
	},

	PolicyFileName: "auth_policies.star",
	VerifierServer: "https://idp.pre.portal.redisbe.com/auth/realms/pre-isbe",

	Dbname:      "data/isbetmf.db",
	ClonePeriod: DefaultClonePeriod,
	Features: Features{
		OfferingLaunchOnlyByAdmin: true,
		GenerateIDOnCreate:        true,
		AllowIDInBody:             true,
		VerifyJWTSignature:        true,
	},
}

var isbeproConfig = &Config{
	Environment:  ISBE_PRO,
	ProxyEnabled: false,

	// The operator is Alastria
	ServerOperatorOrganizationIdentifier: "VATES-G87936159",
	ServerOperatorDid:                    "did:elsi:VATES-G87936159",
	ServerOperatorName:                   "Alastria",
	ServerOperatorCountry:                "ES",

	LEARPower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "Onboarding",
		Action:   []string{"execute"},
	},
	ProductCreatePower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "ProductOffering",
		Action:   []string{"Create"},
	},
	ProductUpdatePower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "ProductOffering",
		Action:   []string{"Update"},
	},
	ProductDeletePower: types.OnePower{
		Type:     "organization",
		Domain:   "ISBE",
		Function: "ProductOffering",
		Action:   []string{"Delete"},
	},

	PolicyFileName: "auth_policies.star",
	VerifierServer: "https://idp.portal.redisbe.com/auth/realms/pro-isbe",

	Dbname:      "data/isbetmf.db",
	ClonePeriod: DefaultClonePeriod,
	Features: Features{
		OfferingLaunchOnlyByAdmin: true,
		GenerateIDOnCreate:        true,
		AllowIDInBody:             true,
		VerifyJWTSignature:        true,
	},
}
