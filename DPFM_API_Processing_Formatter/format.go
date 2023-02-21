package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToGeneralUpdates(general dpfm_api_input_reader.General) *GeneralUpdates {
	data := general

	return &GeneralUpdates{
		BaseUnit:                      data.BaseUnit,
		ValidityStartDate:             data.ValidityStartDate,
		GrossWeight:                   data.GrossWeight,
		NetWeight:                     data.NetWeight,
		WeightUnit:                    data.WeightUnit,
		InternalCapacityQuantity:      data.InternalCapacityQuantity,
		InternalCapacityQuantityUnit:  data.InternalCapacityQuantityUnit,
		SizeOrDimensionText:           data.SizeOrDimensionText,
		ProductStandardID:             data.ProductStandardID,
		IndustryStandardName:          data.IndustryStandardName,
		BarcodeType:                   data.BarcodeType,
		CountryOfOrigin:               data.CountryOfOrigin,
		CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
		ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func ConvertToBusinessPartnerUpdates(businessPartnerUpdates dpfm_api_input_reader.BusinessPartner) *BusinessPartnerUpdates {
	data := businessPartnerUpdates

	return &BusinessPartnerUpdates{
		BusinessPartnerProduct: data.BusinessPartnerProduct,
		IsMarkedForDeletion:    data.IsMarkedForDeletion,
	}
}

func ConvertToAllergenUpdates(allergenUpdates dpfm_api_input_reader.Allergen) *AllergenUpdates {
	data := allergenUpdates

	return &AllergenUpdates{
		Allergen:            data.Allergen,
		AllergenIsContained: data.AllergenIsContained,
	}
}

func ConvertToNutritionalInfoUpdates(nutritionalInfoUpdates dpfm_api_input_reader.NutritionalInfo) *NutritionalInfoUpdates {
	data := nutritionalInfoUpdates

	return &NutritionalInfoUpdates{
		Nutrient:            data.Nutrient,
		NutrientContent:     data.NutrientContent,
		NutrientContentUnit: data.NutrientContentUnit,
	}
}

func ConvertToCaloriesUpdates(caloriesUpdates dpfm_api_input_reader.Calories) *CaloriesUpdates {
	data := caloriesUpdates

	return &CaloriesUpdates{
		CaloryUnitQuantity: data.CaloryUnitQuantity,
		Calories:           data.Calories,
		CaloryUnit:         data.CaloryUnit,
	}
}

func ConvertToBPPlantUpdates(bpPlantUpdates dpfm_api_input_reader.BPPlant) *BPPlantUpdates {
	data := bpPlantUpdates

	return &BPPlantUpdates{
		AvailabilityCheckType:                     data.AvailabilityCheckType,
		MRPType:                                   data.MRPType,
		MRPController:                             data.MRPController,
		ReorderThresholdQuantity:                  data.ReorderThresholdQuantity,
		PlanningTimeFence:                         data.PlanningTimeFence,
		MRPPlanningCalendar:                       data.MRPPlanningCalendar,
		SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
		SafetyDuration:                            data.SafetyDuration,
		MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
		MinimumDeliveryQuantityInBaseUnit:         data.MinimumDeliveryQuantityInBaseUnit,
		MinimumDeliveryLotSizeQuantityInBaseUnit:  data.MinimumDeliveryLotSizeQuantityInBaseUnit,
		StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
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

func ConvertToStorageLocationUpdates(storageLocationUpdates dpfm_api_input_reader.StorageLocation) *StorageLocationUpdates {
	data := storageLocationUpdates

	return &StorageLocationUpdates{
		InventoryBlockStatus: data.InventoryBlockStatus,
		IsMarkedForDeletion:  data.IsMarkedForDeletion,
	}
}

func ConvertToMRPAreaUpdates(mRPAreaUpdates dpfm_api_input_reader.MRPArea) *MRPAreaUpdates {
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

func ConvertToWorkSchedulingUpdates(workSchedulingUpdates dpfm_api_input_reader.WorkScheduling) *WorkSchedulingUpdates {
	data := workSchedulingUpdates

	return &WorkSchedulingUpdates{
		ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
		ProductProcessingTime:         data.ProductProcessingTime,
		ProductionSupervisor:          data.ProductionSupervisor,
		ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
		ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
		PDTCompIsMarkedForBackflush:   data.PDTCompIsMarkedForBackflush,
		ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func ConvertToAccountingUpdates(accountingUpdates dpfm_api_input_reader.Accounting) *AccountingUpdates {
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

func ConvertToProductDescriptionUpdates(productDescriptionUpdates dpfm_api_input_reader.ProductDescription) *ProductDescriptionUpdates {
	data := productDescriptionUpdates

	return &ProductDescriptionUpdates{
		ProductDescription: data.ProductDescription,
	}
}

func ConvertToProductDescByBPUpdates(productDescByBPUpdates dpfm_api_input_reader.ProductDescByBP) *ProductDescByBPUpdates {
	data := productDescByBPUpdates

	return &ProductDescByBPUpdates{
		Language:           data.Language,
		ProductDescription: data.ProductDescription,
	}
}

func ConvertToTaxUpdates(taxUpdates dpfm_api_input_reader.Tax) *TaxUpdates {
	data := taxUpdates

	return &TaxUpdates{
		ProductTaxClassification: data.ProductTaxClassification,
	}
}
