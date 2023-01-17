package dpfm_api_processing_formatter

type GeneralUpdates struct {
	BaseUnit                      *string  `json:"BaseUnit"`
	ValidityStartDate             *string  `json:"ValidityStartDate"`
	Division                      *string  `json:"Division"`
	GrossWeight                   *float32 `json:"GrossWeight"`
	WeightUnit                    *string  `json:"WeightUnit"`
	SizeOrDimensionText           *string  `json:"SizeOrDimensionText"`
	IndustryStandardName          *string  `json:"IndustryStandardName"`
	ProductStandardID             *string  `json:"ProductStandardID"`
	NetWeight                     *float32 `json:"NetWeight"`
	CountryOfOrigin               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage       *string  `json:"CountryOfOriginLanguage"`
	ProductAccountAssignmentGroup *string  `json:"ProductAccountAssignmentGroup"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}

type BusinessPartnerUpdates struct {
	BusinessPartnerProduct *string `json:"BusinessPartnerProduct"`
	IsMarkedForDeletion    *bool   `json:"IsMarkedForDeletion"`
}

type BPPlantUpdates struct {
	AvailabilityCheckType                     *string  `json:"AvailabilityCheckType"`
	MRPType                                   *string  `json:"MRPType"`
	MRPController                             *string  `json:"MRPController"`
	ReorderThresholdQuantity                  *float32 `json:"ReorderThresholdQuantity"`
	PlanningTimeFence                         *int     `json:"PlanningTimeFence"`
	MRPPlanningCalender                       *string  `json:"MRPPlanningCalender"`
	SafetyStockQuantityInBaseUnit             *float32 `json:"SafetyStockQuantityInBaseUnit"`
	SafetyDuration                            *int     `json:"SafetyDuration"`
	MaximumStockQuantityInBaseUnit            *float32 `json:"MaximumStockQuantityInBaseUnit"`
	MinumumDeliveryQuantityInBaseUnit         *float32 `json:"MinumumDeliveryQuantityInBaseUnit"`
	MinumumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MinumumDeliveryLotSizeQuantityInBaseUnit"`
	DeliveryLotSizeRoundingQuantityInBaseUnit *float32 `json:"DeliveryLotSizeRoundingQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit  *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	MaximumDeliveryQuantityInBaseUnit         *float32 `json:"MaximumDeliveryQuantityInBaseUnit"`
	DeliveryLotSizeIsFixed                    *bool    `json:"DeliveryLotSizeIsFixed"`
	StandardDeliveryDurationInDays            *int     `json:"StandardDeliveryDurationInDays"`
	IsBatchManagementRequired                 *bool    `json:"IsBatchManagementRequired"`
	BatchManagementPolicy                     *string  `json:"BatchManagementPolicy"`
	InventoryUnit                             *string  `json:"InventoryUnit"`
	IsMarkedForDeletion                       *bool    `json:"IsMarkedForDeletion"`
}

type StorageLocationUpdates struct {
	InventoryBlockStatus *bool `json:"InventoryBlockStatus"`
	IsMarkedForDeletion  *bool `json:"IsMarkedForDeletion"`
}

type MRPAreaUpdates struct {
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

type WorkSchedulingUpdates struct {
	ProductionInvtryManagedLoc    *string `json:"ProductionInvtryManagedLoc"`
	ProductProcessingTime         *int    `json:"ProductProcessingTime"`
	ProductionSupervisor          *string `json:"ProductionSupervisor"`
	ProductProductionQuantityUnit *string `json:"ProductProductionQuantityUnit"`
	ProdnOrderIsBatchRequired     *bool   `json:"ProdnOrderIsBatchRequired"`
	MatlCompIsMarkedForBackflush  *bool   `json:"MatlCompIsMarkedForBackflush"`
	ProductionSchedulingProfile   *string `json:"ProductionSchedulingProfile"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}

type AccountingUpdates struct {
	ValuationClass      *string  `json:"ValuationClass"`
	CostingPolicy       *string  `json:"CostingPolicy"`
	PriceUnitQty        *string  `json:"PriceUnitQty"`
	StandardPrice       *float32 `json:"StandardPrice"`
	MovingAveragePrice  *float32 `json:"MovingAveragePrice"`
	PriceLastChangeDate *string  `json:"PriceLastChangeDate"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type ProductDescriptionUpdates struct {
	ProductDescription *string `json:"ProductDescription"`
}

type ProductDescByBPUpdates struct {
	Language           string  `json:"Language"`
	ProductDescription *string `json:"ProductDescription"`
}

type TaxUpdates struct {
	ProductTaxClassification *string `json:"ProductTaxClassification"`
}
