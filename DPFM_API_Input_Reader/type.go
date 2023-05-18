package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey    string   `json:"connection_key"`
	Result           bool     `json:"result"`
	RedisKey         string   `json:"redis_key"`
	Filepath         string   `json:"filepath"`
	APIStatusCode    int      `json:"api_status_code"`
	RuntimeSessionID string   `json:"runtime_session_id"`
	BusinessPartner  *int     `json:"business_partner"`
	ServiceLabel     string   `json:"service_label"`
	APIType          string   `json:"api_type"`
	General          General  `json:"ProductMaster"`
	APISchema        string   `json:"api_schema"`
	Accepter         []string `json:"accepter"`
	Deleted          bool     `json:"deleted"`
}

type General struct {
	Product                       string            `json:"Product"`
	ProductType                   *string           `json:"ProductType"`
	BaseUnit                      *string           `json:"BaseUnit"`
	ValidityStartDate             *string           `json:"ValidityStartDate"`
	ProductGroup                  *string           `json:"ProductGroup"`
	GrossWeight                   *float32          `json:"GrossWeight"`
	NetWeight                     *float32          `json:"NetWeight"`
	WeightUnit                    *string           `json:"WeightUnit"`
	InternalCapacityQuantity      *float32          `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit  *string           `json:"InternalCapacityQuantityUnit"`
	SizeOrDimensionText           *string           `json:"SizeOrDimensionText"`
	ProductStandardID             *string           `json:"ProductStandardID"`
	IndustryStandardName          *string           `json:"IndustryStandardName"`
	ItemCategory                  *string           `json:"ItemCategory"`
	CountryOfOrigin               *string           `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string           `json:"CountryOfOriginLanguage"`
	BarcodeType                   *string           `json:"BarcodeType"`
	ProductAccountAssignmentGroup *string           `json:"ProductAccountAssignmentGroup"`
	CreationDate                  *string           `json:"CreationDate"`
	LastChangeDate                *string           `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool             `json:"IsMarkedForDeletion"`
	BusinessPartner               []BusinessPartner `json:"BusinessPartner"`
	BPPlant                       []BPPlant         `json:"BPPlant"`
	GeneralDoc                    []GeneralDoc      `json:"GeneralDoc"`
	Tax                           []Tax             `json:"Tax"`
}

type BusinessPartner struct {
	Product                string               `json:"Product"`
	BusinessPartner        int                  `json:"BusinessPartner"`
	ValidityEndDate        string               `json:"ValidityEndDate"`
	ValidityStartDate      string               `json:"ValidityStartDate"`
	BusinessPartnerProduct *string              `json:"BusinessPartnerProduct"`
	IsMarkedForDeletion    *bool                `json:"IsMarkedForDeletion"`
	Allergen               []Allergen           `json:"Allergen"`
	Calories               []Calories           `json:"Calories"`
	NutritionalInfo        []NutritionalInfo    `json:"NutritionalInfo"`
	ProductDescription     []ProductDescription `json:"ProductDescription"`
}

type Allergen struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Allergen            string `json:"Allergen"`
	AllergenIsContained *bool  `json:"AllergenIsContained"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type Calories struct {
	Product            string  `json:"Product"`
	BusinessPartner    int     `json:"BusinessPartner"`
	CaloryUnitQuantity int     `json:"CaloryUnitQuantity"`
	Calories           *int    `json:"Calories"`
	CaloryUnit         *string `json:"CaloryUnit"`
}

type NutritionalInfo struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Nutrient            string  `json:"Nutrient"`
	NutrientContent     *int    `json:"NutrientContent"`
	NutrientContentUnit *string `json:"NutrientContentUnit"`
}

type ProductDescription struct {
	Product            string            `json:"Product"`
	Language           string            `json:"Language"`
	ProductDescription *string           `json:"ProductDescription"`
	ProductDescByBP    []ProductDescByBP `json:"ProductDescByBP"`
}

type ProductDescByBP struct {
	Product            string  `json:"Product"`
	BusinessPartner    int     `json:"BusinessPartner"`
	Language           string  `json:"Language"`
	ProductDescription *string `json:"ProductDescription"`
}

type BPPlant struct {
	Product                                   string            `json:"Product"`
	BusinessPartner                           int               `json:"BusinessPartner"`
	Plant                                     string            `json:"Plant"`
	AvailabilityCheckType                     *string           `json:"AvailabilityCheckType"`
	MRPType                                   *string           `json:"MRPType"`
	MRPController                             *string           `json:"MRPController"`
	ReorderThresholdQuantity                  *float32          `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int              `json:"PlanningTimeFence"`
	MRPPlanningCalendar                       *string           `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32          `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int              `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32          `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32          `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32          `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32          `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32          `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32          `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32          `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool             `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int              `json:"StandardDeliveryDurationInDays"`
	IsBatchManagementRequired                 *bool             `json:"IsBatchManagementRequired"`
	BatchManagementPolicy                     *string           `json:"BatchManagementPolicy"`
	InventoryUnit                             *string           `json:"InventoryUnit"`
	ProfitCenter                              *string           `json:"ProfitCenter"`
	IsMarkedForDeletion                       *bool             `json:"IsMarkedForDeletion"`
	Accounting                                []Accounting      `json:"Accounting"`
	BPPlantDoc                                []BPPlantDoc      `json:"BPPlantDoc"`
	MRPArea                                   []MRPArea         `json:"MRPArea"`
	Quality                                   []Quality         `json:"Quality"`
	StorageLocation                           []StorageLocation `json:"StorageLocation"`
	WorkScheduling                            []WorkScheduling  `json:"WorkScheduling"`
}

type Accounting struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Plant               string   `json:"Plant"`
	ValuationClass      *string  `json:"ValuationClass"`
	CostingPolicy       *string  `json:"CostingPolicy"`
	PriceUnitQty        *string  `json:"PriceUnitQty"`
	StandardPrice       *float32 `json:"StandardPrice"`
	MovingAveragePrice  *float32 `json:"MovingAveragePrice"`
	PriceLastChangeDate *string  `json:"PriceLastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type BPPlantDoc struct {
	Product                  string  `json:"Product"`
	BusinessPartner          int     `json:"BusinessPartner"`
	Plant                    string  `json:"Plant"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type MRPArea struct {
	Product                                   string   `json:"Product"`
	BusinessPartner                           int      `json:"BusinessPartner"`
	Plant                                     string   `json:"Plant"`
	MRPArea                                   string   `json:"MRPArea"`
	StorageLocationForMRP                     *string  `json:"StorageLocationForMRP"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantity                  *float32 `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int     `json:"PlanningTimeFence"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int     `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinumumDeliveryQuantityInBaseUnit         *float32 `json:"MinumumDeliveryQuantityInBaseUnit"`
	MinumumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinumumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int     `json:"StandardDeliveryDurationInDays"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type Quality struct {
	Product                        string  `json:"Product"`
	BusinessPartner                int     `json:"BusinessPartner"`
	Plant                          string  `json:"Plant"`
	MaximumStoragePeriod           *string `json:"MaximumStoragePeriod"`
	QualityMgmtCtrlKey             *string `json:"QualityMgmtCtrlKey"`
	MatlQualityAuthorizationGroup  *string `json:"MatlQualityAuthorizationGroup"`
	HasPostToInspectionStock       *bool   `json:"HasPostToInspectionStock"`
	InspLotDocumentationIsRequired *bool   `json:"InspLotDocumentationIsRequired"`
	SuplrQualityManagementSystem   *string `json:"SuplrQualityManagementSystem"`
	RecrrgInspIntervalTimeInDays   *int    `json:"RecrrgInspIntervalTimeInDays"`
	ProductQualityCertificateType  *string `json:"ProductQualityCertificateType"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}

type StorageLocation struct {
	Product              string       `json:"Product"`
	BusinessPartner      int          `json:"BusinessPartner"`
	Plant                string       `json:"Plant"`
	StorageLocation      string       `json:"StorageLocation"`
	CreationDate         *string      `json:"CreationDate"`
	InventoryBlockStatus *bool        `json:"InventoryBlockStatus"`
	IsMarkedForDeletion  *bool        `json:"IsMarkedForDeletion"`
	StorageBin           []StorageBin `json:"StorageBin"`
}

type StorageBin struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Plant               string  `json:"Plant"`
	StorageLocation     string  `json:"StorageLocation"`
	StorageBin          string  `json:"StorageBin"`
	CreationDate        *string `json:"CreationDate"`
	BlockStatus         *bool   `json:"BlockStatus"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type WorkScheduling struct {
	Product                       string   `json:"Product"`
	BusinessPartner               int      `json:"BusinessPartner"`
	Plant                         string   `json:"Plant"`
	ProductionInvtryManagedLoc    *string  `json:"ProductionInvtryManagedLoc"`
	ProductProcessingTime         *int     `json:"ProductProcessingTime"`
	ProductionSupervisor          *string  `json:"ProductionSupervisor"`
	ProductProductionQuantityUnit *string  `json:"ProductProductionQuantityUnit"`
	ProdnOrderIsBatchRequired     *bool    `json:"ProdnOrderIsBatchRequired"`
	PDTCompIsMarkedForBackflush   *bool    `json:"PDTCompIsMarkedForBackflush"`
	ProductionSchedulingProfile   *string  `json:"ProductionSchedulingProfile"`
	MinimumLotSizeQuantity        *float32 `json:"MinimumLotSizeQuantity"`
	StandardLotSizeQuantity       *float32 `json:"StandardLotSizeQuantity"`
	LotSizeRoundingQuantity       *float32 `json:"LotSizeRoundingQuantity"`
	MaximumLotSizeQuantity        *float32 `json:"MaximumLotSizeQuantity"`
	LotSizeIsFixed                *bool    `json:"LotSizeIsFixed"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}

type GeneralDoc struct {
	Product                  string  `json:"Product"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type Tax struct {
	Product                  string  `json:"Product"`
	Country                  string  `json:"Country"`
	ProductTaxCategory       string  `json:"ProductTaxCategory"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
}
