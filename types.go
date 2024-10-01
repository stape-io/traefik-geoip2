package incsw

const (
	dataTypeExtended = iota
	dataTypePointer
	dataTypeString
	dataTypeFloat64
	dataTypeBytes
	dataTypeUint16
	dataTypeUint32
	dataTypeMap
	dataTypeInt32
	dataTypeUint64
	dataTypeUint128
	dataTypeSlice
	dataTypeDataCacheContainer
	dataTypeEndMarker
	dataTypeBool
	dataTypeFloat32

	dataSectionSeparatorSize
)

type Continent struct {
	GeoNameID uint32
	Code      string
	Names     map[string]string
}

type Country struct {
	ISOCode           string
	Names             map[string]string
	Type              string // [RepresentedCountry]
	GeoNameID         uint32
	Confidence        uint16 // Enterprise [Country, RegisteredCountry]
	IsInEuropeanUnion bool
}

type Subdivision struct {
	ISOCode    string
	Names      map[string]string
	GeoNameID  uint32
	Confidence uint16 // Enterprise
}

type City struct {
	Names      map[string]string
	GeoNameID  uint32
	Confidence uint16 // Enterprise
}

type Location struct {
	Latitude       float64
	Longitude      float64
	TimeZone       string
	AccuracyRadius uint16
	MetroCode      uint16
}

type Postal struct {
	Code       string
	Confidence uint16 // Enterprise
}

type Traits struct {
	StaticIPScore                float64 // Enterprise
	ISP                          string  // Enterprise
	Organization                 string  // Enterprise
	ConnectionType               string  // Enterprise
	Domain                       string  // Enterprise
	UserType                     string  // Enterprise
	AutonomousSystemOrganization string  // Enterprise
	AutonomousSystemNumber       uint32  // Enterprise
	IsLegitimateProxy            bool    // Enterprise
	MobileCountryCode            string  // Enterprise
	MobileNetworkCode            string  // Enterprise
	IsAnonymousProxy             bool
	IsSatelliteProvider          bool
}

type CountryResult struct {
	Continent          Continent
	Country            Country
	RegisteredCountry  Country
	RepresentedCountry Country
	Traits             Traits
}

type CityResult struct {
	Continent          Continent
	Country            Country
	Subdivisions       []Subdivision
	City               City
	Location           Location
	Postal             Postal
	RegisteredCountry  Country
	RepresentedCountry Country
	Traits             Traits
}

type ISP struct {
	AutonomousSystemNumber       uint32
	AutonomousSystemOrganization string
	ISP                          string
	Organization                 string
	MobileCountryCode            string
	MobileNetworkCode            string
}

type ConnectionType struct {
	ConnectionType string
}

type AnonymousIP struct {
	IsAnonymous        bool
	IsAnonymousVPN     bool
	IsHostingProvider  bool
	IsPublicProxy      bool
	IsTorExitNode      bool
	IsResidentialProxy bool
}

type ASN struct {
	AutonomousSystemNumber       uint32
	AutonomousSystemOrganization string
}

type Domain struct {
	Domain string
}
