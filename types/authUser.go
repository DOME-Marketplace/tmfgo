package types

// AuthUser represents the authenticated user's information from the JWT mandator object.
type AuthUser struct {
	CommonName             string `json:"commonName"`
	Country                string `json:"country"`
	EmailAddress           string `json:"emailAddress"`
	Organization           string `json:"organization"`
	OrganizationIdentifier string `json:"organizationIdentifier"`
	SerialNumber           string `json:"serialNumber"`
	IsAuthenticated        bool
	IsLEAR                 bool
	IsOwner                bool
	ProductCreatePower     bool
	ProductUpdatePower     bool
	ProductDeletePower     bool
	AccessToken            string
	TokenMap               map[string]any
}

func (u *AuthUser) ToMap() map[string]any {
	return map[string]any{
		"commonName":             u.CommonName,
		"country":                u.Country,
		"emailAddress":           u.EmailAddress,
		"organization":           u.Organization,
		"organizationIdentifier": u.OrganizationIdentifier,
		"serialNumber":           u.SerialNumber,
		"isAuthenticated":        u.IsAuthenticated,
		"isLEAR":                 u.IsLEAR,
		"isOwner":                u.IsOwner,
	}
}
