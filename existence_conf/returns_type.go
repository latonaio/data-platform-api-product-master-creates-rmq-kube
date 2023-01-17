package existence_conf

type Returns struct {
	ConnectionKey      string              `json:"connection_key"`
	Result             bool                `json:"result"`
	RedisKey           string              `json:"redis_key"`
	RuntimeSessionID   string              `json:"runtime_session_id"`
	BusinessPartnerID  *int                `json:"business_partner"`
	Filepath           string              `json:"filepath"`
	ServiceLabel       string              `json:"service_label"`
	ProductTypeReturn  ProductTypeReturn   `json:"ProductType"`
	QuantityUnitReturn QuantityUnitReturn  `json:"QuantityUnit"`
	ProductGroupReturn *ProductGroupReturn `json:"ProductGroup,omitempty"`
	LanguageReturn     *LanguageReturn     `json:"Language,omitempty"`
	CountryReturn      *CountryReturn      `json:"Country,omitempty"`
	APISchema          string              `json:"api_schema"`
	Accepter           []string            `json:"accepter"`
	Deleted            bool                `json:"deleted"`
}

type ProductTypeReturn struct {
	ProductType string `json:"BusinessPartner"`
}

type QuantityUnitReturn struct {
	QuantityUnit string `json:"QuantityUnit"`
}

type ProductGroupReturn struct {
	ProductGroup string `json:"ProductGroup"`
}

type LanguageReturn struct {
	Language string `json:"Language"`
}

type CountryReturn struct {
	Country string `json:"Country"`
}
