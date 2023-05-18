package existence_conf

type Returns struct {
	ConnectionKey         string                `json:"connection_key"`
	Result                bool                  `json:"result"`
	RedisKey              string                `json:"redis_key"`
	RuntimeSessionID      string                `json:"runtime_session_id"`
	BusinessPartnerID     *int                  `json:"business_partner"`
	Filepath              string                `json:"filepath"`
	ServiceLabel          string                `json:"service_label"`
	ProductTypeReturn     ProductTypeReturn     `json:"ProductType"`
	QuantityUnitReturn    QuantityUnitReturn    `json:"QuantityUnit"`
	ProductGroupReturn    ProductGroupReturn    `json:"ProductGroup"`
	CountryReturn         CountryReturn         `json:"Country"`
	LanguageReturn        LanguageReturn        `json:"Language"`
	BPGeneralReturn       BPGeneralReturn       `json:"BusinessPartnerGeneral"`
	AllergenReturn        AllergenReturn        `json:"Allergen"`
	PlantGeneralReturn    PlantGeneralReturn    `json:"PlantGeneral"`
	StorageLocationReturn StorageLocationReturn `json:"StorageLocation"`
	StorageBinReturn      StorageBinReturn      `json:"StorageBin"`
	APISchema             string                `json:"api_schema"`
	Accepter              []string              `json:"accepter"`
	Deleted               bool                  `json:"deleted"`
}

type ProductTypeReturn struct {
	ProductType string `json:"ProductType"`
}

type QuantityUnitReturn struct {
	QuantityUnit string `json:"QuantityUnit"`
}

type ProductGroupReturn struct {
	ProductGroup string `json:"ProductGroup"`
}

type CountryReturn struct {
	Country string `json:"Country"`
}

type LanguageReturn struct {
	Language string `json:"Language"`
}

type BPGeneralReturn struct {
	BusinessPartner int `json:"BusinessPartner"`
}

type AllergenReturn struct {
	Allergen string `json:"Allergen"`
}

type PlantGeneralReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type StorageLocationReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	StorageLocation string `json:"StorageLocation"`
}

type StorageBinReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	StorageLocation string `json:"StorageLocation"`
	StorageBin      string `json:"StorageBin"`
}
