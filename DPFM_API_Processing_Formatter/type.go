package dpfm_api_processing_formatter

type GeneralUpdates struct {
	Product                       string   `json:"Product"`
	ValidityStartDate             string   `json:"ValidityStartDate"`
	ValidityEndDate	              string   `json:"ValidityEndDate"`
	ProductGroup	              *string  `json:"ProductGroup"`
	GrossWeight                   *float32 `json:"GrossWeight"`
	NetWeight                     *float32 `json:"NetWeight"`
	WeightUnit                    *string  `json:"WeightUnit"`
	InternalCapacityQuantity      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit  *string  `json:"InternalCapacityQuantityUnit"`
	SizeOrDimensionText           *string  `json:"SizeOrDimensionText"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	IndustryStandardName          *string  `json:"IndustryStandardName"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string  `json:"CountryOfOriginLanguage"`
	LocalRegionOfOrigin           *string  `json:"LocalRegionOfOrigin"`
	LocalSubRegionOfOrigin        *string  `json:"LocalSubRegionOfOrigin"`
	BarcodeType                   *string  `json:"BarcodeType"`
	ProductAccountAssignmentGroup *string  `json:"ProductAccountAssignmentGroup"`
}

type ProductDescriptionUpdates struct {
	Product            string  `json:"Product"`
	Language           string  `json:"Language"`
	ProductDescription *string `json:"ProductDescription"`
}

type BusinessPartnerUpdates struct {
	Product                string  `json:"Product"`
	BusinessPartner        int 	   `json:"BusinessPartner"`
	ValidityStartDate      string  `json:"ValidityStartDate"`
	ValidityEndDate        string  `json:"ValidityEndDate"`
	BusinessPartnerProduct *string `json:"BusinessPartnerProduct"`
}

type ProductDescByBPUpdates struct {
	Product            string  `json:"Product"`
	BusinessPartner    int     `json:"BusinessPartner"`
	Language           string  `json:"Language"`
	ProductDescription *string `json:"ProductDescription"`
}

type BPPlantUpdates struct {
	Product                                   string   `json:"Product"`
	BusinessPartner                           int      `json:"BusinessPartner"`
	Plant                                     string   `json:"Plant"`
	MRPType                                   string   `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantityInBaseUnit        *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int     `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32 `json:"SafetyDuration"`
	SafetyDurationUnit                        *string  `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinimumDeliveryQuantityInBaseUnit         *float32 `json:"MinimumDeliveryQuantityInBaseUnit"`
	MinimumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinimumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryQuantityInBaseUnit        *float32 `json:"StandardDeliveryQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration	              *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit	          *string  `json:"StandardDeliveryDurationUnit"`
	IsBatchManagementRequired                 *bool    `json:"IsBatchManagementRequired"`
	BatchManagementPolicy                     *string  `json:"BatchManagementPolicy"`
	ProfitCenter                              *string  `json:"ProfitCenter"`
}

type MRPAreaUpdates struct {
	Product                                   string   `json:"Product"`
	BusinessPartner                           int      `json:"BusinessPartner"`
	Plant                                     string   `json:"Plant"`
	MRPArea                                   string   `json:"MRPArea"`
	MRPType                                   string   `json:"MRPType"`
	MRPController                             string   `json:"MRPController"`
	StorageLocationForMRP                     string   `json:"StorageLocationForMRP"`
	ReorderThresholdQuantityInBaseUnit        *float32 `json:"ReorderThresholdQuantityInBaseUnit"`
	PlanningTimeFenceInDays                   *int     `json:"PlanningTimeFenceInDays"`
	MRPPlanningCalendar                       *string  `json:"MRPPlanningCalendar"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *float32 `json:"SafetyDuration"`
	SafetyDurationUnit                        *string  `json:"SafetyDurationUnit"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinumumDeliveryQuantityInBaseUnit         *float32 `json:"MinumumDeliveryQuantityInBaseUnit"`
	MinumumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinumumDeliveryLotSizeQuantityInBaseUnit"`
	StandardDeliveryQuantityInBaseUnit        *float32 `json:"StandardDeliveryQuantityInBaseUnit"`
	StandardDeliveryLotSizeQuantityInBaseUnit *float32 `json:"StandardDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDuration	              *float32 `json:"StandardDeliveryDuration"`
	StandardDeliveryDurationUnit              *string  `json:"StandardDeliveryDurationUnit"`
}

type ProductionUpdates struct {
	Product												string		`json:"Product"`
	BusinessPartner										int			`json:"BusinessPartner"`
	Plant												string		`json:"Plant"`
	ProductionStorageLocation							string		`json:"ProductionStorageLocation"`
	ProductProcessingDuration							float32		`json:"ProductProcessingDuration"`
	ProductProductionQuantityUnit						string		`json:"ProductProductionQuantityUnit"`
	MinimumProductionQuantityInBaseUnit					float32		`json:"MinimumProductionQuantityInBaseUnit"`
	MinimumProductionLotSizeQuantityInBaseUnit			float32		`json:"MinimumProductionLotSizeQuantityInBaseUnit"`
	StandardProductionQuantityInBaseUnit				float32		`json:"StandardProductionQuantityInBaseUnit"`
	StandardProductionLotSizeQuantityInBaseUnit			float32		`json:"StandardProductionLotSizeQuantityInBaseUnit"`
	MaximumProductionQuantityInBaseUnit					float32		`json:"MaximumProductionQuantityInBaseUnit"`
	MaximumProductionLotSizeQuantityInBaseUnit			float32		`json:"MaximumProductionLotSizeQuantityInBaseUnit"`
	ProductionLotSizeRoundingQuantityInBaseUnit			*float32	`json:"ProductionLotSizeRoundingQuantityInBaseUnit"`
	MinimumProductionQuantityInProductionUnit			float32		`json:"MinimumProductionQuantityInProductionUnit"`
	MinimumProductionLotSizeQuantityInProductionUnit	float32		`json:"MinimumProductionLotSizeQuantityInProductionUnit"`
	StandardProductionQuantityInProductionUnit			float32		`json:"StandardProductionQuantityInProductionUnit"`
	StandardProductionLotSizeQuantityInProductionUnit	float32		`json:"StandardProductionLotSizeQuantityInProductionUnit"`
	MaximumProductionLotSizeQuantityInProductionUnit	float32		`json:"MaximumProductionLotSizeQuantityInProductionUnit"`
	MaximumProductionQuantityInProductionUnit			float32		`json:"MaximumProductionQuantityInProductionUnit"`
	ProductionLotSizeRoundingQuantityInProductionUnit	*float32	`json:"ProductionLotSizeRoundingQuantityInProductionUnit"`
	ProductionLotSizeIsFixed							*bool		`json:"ProductionLotSizeIsFixed"`
	ProductIsBatchManagedInProductionPlant				*bool		`json:"ProductIsBatchManagedInProductionPlant"`
	ProductIsMarkedForBackflush							*bool		`json:"ProductIsMarkedForBackflush"`
	ProductionSchedulingProfile							*string		`json:"ProductionSchedulingProfile"`
}

type AccountingUpdates struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Plant               string   `json:"Plant"`
	ValuationClass      *string  `json:"ValuationClass"`
}

type QualityUpdates struct {
	Product                        string  `json:"Product"`
	BusinessPartner                int     `json:"BusinessPartner"`
	Plant                          string  `json:"Plant"`
	QualityMgmtCtrlKey             *string `json:"QualityMgmtCtrlKey"`
	ProductSpecification		   *string `json:"ProductSpecification"`
	MaximumStoragePeriodInDays     *int	   `json:"MaximumStoragePeriodInDays"`
	RecrrgInspIntervalTimeInDays   *int    `json:"RecrrgInspIntervalTimeInDays"`
	ProductQualityCertificateType  *string `json:"ProductQualityCertificateType"`
}

type StorageLocationUpdates struct {
	Product              string  `json:"Product"`
	BusinessPartner      int     `json:"BusinessPartner"`
	Plant                string  `json:"Plant"`
	StorageLocation      string  `json:"StorageLocation"`
	BlockStatus 		 *bool   `json:"BlockStatus"`
}

type StorageBinUpdates struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Plant               string  `json:"Plant"`
	StorageLocation     string  `json:"StorageLocation"`
	StorageBin          string  `json:"StorageBin"`
	BlockStatus         *bool   `json:"BlockStatus"`
}

type TaxUpdates struct {
	Product                  string  `json:"Product"`
	Country                  string  `json:"Country"`
	ProductTaxClassification *string `json:"ProductTaxClassification"`
}

type AllergenUpdates struct {
	Product             string `json:"Product"`
	BusinessPartner     int    `json:"BusinessPartner"`
	Allergen            string `json:"Allergen"`
	AllergenIsContained *bool  `json:"AllergenIsContained"`
}

type NutritionalInfoUpdates struct {
	Product             string   `json:"Product"`
	BusinessPartner     int      `json:"BusinessPartner"`
	Nutrient            string   `json:"Nutrient"`
	NutrientContent     *float32 `json:"NutrientContent"`
	NutrientContentUnit *string  `json:"NutrientContentUnit"`
}

type CaloriesUpdates struct {
	Product            	string   `json:"Product"`
	BusinessPartner    	int      `json:"BusinessPartner"`
	Calories           	*float32 `json:"Calories"`
	CaloryUnit         	*string  `json:"CaloryUnit"`
	CaloryUnitQuantity 	*int     `json:"CaloryUnitQuantity"`
}
