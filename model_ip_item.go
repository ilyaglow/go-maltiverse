/*
 * Maltiverse API
 *
 * Maltiverse API
 *
 * API version: 1.0.0
 * Contact: root@maltiverse.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package maltiverse

type IPItem struct {
	IpAddr         string   `json:"ip_addr"`
	Type_          string   `json:"type"`
	Classification string   `json:"classification"`
	Tag            []string `json:"tag,omitempty"`
	Blacklist      []BlItem `json:"blacklist,omitempty"`
	// The the date when the indicator is created. If no specified takes the current date value
	CreationTime string `json:"creation_time,omitempty"`
	// The the date when the indicator got its last modification. If no specified takes the current date value
	ModificationTime string `json:"modification_time,omitempty"`
	// Country code related to the IP address.
	CountryCode string `json:"country_code,omitempty"`
	// City related to the IP address.
	City string `json:"city,omitempty"`
	// State related to the IP address.
	State    string        `json:"state,omitempty"`
	Location *LocationItem `json:"location,omitempty"`
	Email    []string      `json:"email,omitempty"`
	// Address related to the IP address.
	Address string `json:"address,omitempty"`
	// Autonomous system related to the IP address.
	AsName string `json:"as_name,omitempty"`
	// Autonomous system CIDR related to the IP address.
	AsnCidr string `json:"asn_cidr,omitempty"`
	// Country Code related to the IP address.
	AsnCountryCode string `json:"asn_country_code,omitempty"`
	// ASN date related to the IP address.
	AsnDate string `json:"asn_date,omitempty"`
	// ASN registry related to the IP address.
	AsnRegistry string `json:"asn_registry,omitempty"`
	// CIDR related to the IP address.
	Cidr string `json:"cidr,omitempty"`
	// Registrant name related to the IP address.
	RegistrantName string `json:"registrant_name,omitempty"`
	// Postal code to the IP address.
	PostalCode string `json:"postal_code,omitempty"`
	// Last time the whois reocord related to the IP address was updated.
	LastUpdated string `json:"last_updated,omitempty"`
}