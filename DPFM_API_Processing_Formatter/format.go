package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToGeneralUpdates(general dpfm_api_input_reader.General) *GeneralUpdates {
	data := general

	return &GeneralUpdates{
		BaseUnit:                      data.BaseUnit,
		ValidityStartDate:             data.ValidityStartDate,
		Division:                      data.Division,
		GrossWeight:                   data.GrossWeight,
		WeightUnit:                    data.WeightUnit,
		SizeOrDimensionText:           data.SizeOrDimensionText,
		IndustryStandardName:          data.IndustryStandardName,
		ProductStandardID:             data.ProductStandardID,
		NetWeight:                     data.NetWeight,
		CountryOfOrigin:               data.CountryOfOrigin,
		CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
		ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func ConvertToBusinessPartnerUpdates(businessPartnerUpdates BusinessPartnerUpdates) *BusinessPartnerUpdates {
	data := businessPartnerUpdates

	return &BusinessPartnerUpdates{
		BusinessPartnerProduct: data.BusinessPartnerProduct,
		IsMarkedForDeletion:    data.IsMarkedForDeletion,
	}
}

func ConvertToBPPlantUpdates(bpPlantUpdates BPPlantUpdates) *BPPlantUpdates {
	data := bpPlantUpdates

	return &BPPlantUpdates{
		AvailabilityCheckType:                     data.AvailabilityCheckType,
		MRPType:                                   data.MRPType,
		MRPController:                             data.MRPController,
		ReorderThresholdQuantity:                  data.ReorderThresholdQuantity,
		PlanningTimeFence:                         data.PlanningTimeFence,
		MRPPlanningCalender:                       data.MRPPlanningCalender,
		SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
		SafetyDuration:                            data.SafetyDuration,
		MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
		MinumumDeliveryQuantityInBaseUnit:         data.MinumumDeliveryQuantityInBaseUnit,
		MinumumDeliveryLotSizeQuantityInBaseUnit:  data.MinumumDeliveryLotSizeQuantityInBaseUnit,
		DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
		MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
		MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
		DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
		StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
		IsBatchManagementRequired:                 data.IsBatchManagementRequired,
		BatchManagementPolicy:                     data.BatchManagementPolicy,
		InventoryUnit:                             data.InventoryUnit,
		IsMarkedForDeletion:                       data.IsMarkedForDeletion,
	}
}

func ConvertToStorageLocationUpdates(storageLocationUpdates StorageLocationUpdates) *StorageLocationUpdates {
	data := storageLocationUpdates

	return &StorageLocationUpdates{
		InventoryBlockStatus: data.InventoryBlockStatus,
		IsMarkedForDeletion:  data.IsMarkedForDeletion,
	}
}

func ConvertToMRPAreaUpdates(mRPAreaUpdates MRPAreaUpdates) *MRPAreaUpdates {
	data := mRPAreaUpdates

	return &MRPAreaUpdates{
		StorageLocationForMRP:                     data.StorageLocationForMRP,
		MRPType:                                   data.MRPType,
		MRPController:                             data.MRPController,
		ReorderThresholdQuantity:                  data.ReorderThresholdQuantity,
		PlanningTimeFence:                         data.PlanningTimeFence,
		MRPPlanningCalendar:                       data.MRPPlanningCalendar,
		SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
		SafetyDuration:                            data.SafetyDuration,
		MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
		MinumumDeliveryQuantityInBaseUnit:         data.MinumumDeliveryQuantityInBaseUnit,
		MinumumDeliveryLotSizeQuantityInBaseUnit:  data.MinumumDeliveryLotSizeQuantityInBaseUnit,
		StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
		DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
		MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
		MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
		DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
		StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
		IsMarkedForDeletion:                       data.IsMarkedForDeletion,
	}
}

func ConvertToWorkSchedulingUpdates(workSchedulingUpdates WorkSchedulingUpdates) *WorkSchedulingUpdates {
	data := workSchedulingUpdates

	return &WorkSchedulingUpdates{
		ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
		ProductProcessingTime:         data.ProductProcessingTime,
		ProductionSupervisor:          data.ProductionSupervisor,
		ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
		ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
		MatlCompIsMarkedForBackflush:  data.MatlCompIsMarkedForBackflush,
		ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func ConvertToAccountingUpdates(accountingUpdates AccountingUpdates) *AccountingUpdates {
	data := accountingUpdates

	return &AccountingUpdates{
		ValuationClass:      data.ValuationClass,
		CostingPolicy:       data.CostingPolicy,
		PriceUnitQty:        data.PriceUnitQty,
		StandardPrice:       data.StandardPrice,
		MovingAveragePrice:  data.MovingAveragePrice,
		PriceLastChangeDate: data.PriceLastChangeDate,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}

func ConvertToProductDescriptionUpdates(productDescriptionUpdates ProductDescriptionUpdates) *ProductDescriptionUpdates {
	data := productDescriptionUpdates

	return &ProductDescriptionUpdates{
		ProductDescription: data.ProductDescription,
	}
}

func ConvertToProductDescByBPUpdates(productDescByBPUpdates ProductDescByBPUpdates) *ProductDescByBPUpdates {
	data := productDescByBPUpdates

	return &ProductDescByBPUpdates{
		Language:           data.Language,
		ProductDescription: data.ProductDescription,
	}
}

func ConvertToTaxUpdates(taxUpdates TaxUpdates) *TaxUpdates {
	data := taxUpdates

	return &TaxUpdates{
		ProductTaxClassification: data.ProductTaxClassification,
	}
}
