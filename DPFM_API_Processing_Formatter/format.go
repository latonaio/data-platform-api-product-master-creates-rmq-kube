package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToGeneralUpdates(general dpfm_api_input_reader.General) *GeneralUpdates {
	data := general

	return &GeneralUpdates{
		Product:                       data.Product,
		ValidityStartDate:             data.ValidityStartDate,
		ValidityEndDate: 	           data.ValidityEndDate,
		ProductGroup: 	       		   data.ProductGroup,		
		GrossWeight:                   data.GrossWeight,
		NetWeight:                     data.NetWeight,
		WeightUnit:                    data.WeightUnit,
		InternalCapacityQuantity:      data.InternalCapacityQuantity,
		InternalCapacityQuantityUnit:  data.InternalCapacityQuantityUnit,
		SizeOrDimensionText:           data.SizeOrDimensionText,
		ProductStandardID:             data.ProductStandardID,
		IndustryStandardName:          data.IndustryStandardName,
		MarkingOfMaterial:             data.MarkingOfMaterial,
		BarcodeType:                   data.BarcodeType,
		CountryOfOrigin:               data.CountryOfOrigin,
		CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
		LocalRegionOfOrigin:           data.LocalRegionOfOrigin,
		LocalSubRegionOfOrigin:        data.LocalSubRegionOfOrigin,
		ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
	}
}

func ConvertToBusinessPartnerUpdates(businessPartnerUpdates dpfm_api_input_reader.BusinessPartner) *BusinessPartnerUpdates {
	data := businessPartnerUpdates

	return &BusinessPartnerUpdates{
		Product:                data.Product,
		BusinessPartner:        data.BusinessPartner,
		ValidityEndDate:        data.ValidityEndDate,
		ValidityStartDate:      data.ValidityStartDate,
		BusinessPartnerProduct: data.BusinessPartnerProduct,
	}
}

func ConvertToBPPlantUpdates(bpPlantUpdates dpfm_api_input_reader.BPPlant) *BPPlantUpdates {
	data := bpPlantUpdates

	return &BPPlantUpdates{
		Product:								   data.Product,
		BusinessPartner:						   data.BusinessPartner,
		Plant:						   			   data.Plant,
		MRPType:                                   data.MRPType,
		MRPController:                             data.MRPController,
		ReorderThresholdQuantityInBaseUnit:        data.ReorderThresholdQuantityInBaseUnit,
		PlanningTimeFenceInDays:                   data.PlanningTimeFenceInDays,
		MRPPlanningCalendar:                       data.MRPPlanningCalendar,
		SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
		SafetyDuration:                            data.SafetyDuration,
		SafetyDurationUnit:                        data.SafetyDurationUnit,
		MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
		MinimumDeliveryQuantityInBaseUnit:         data.MinimumDeliveryQuantityInBaseUnit,
		MinimumDeliveryLotSizeQuantityInBaseUnit:  data.MinimumDeliveryLotSizeQuantityInBaseUnit,
		StandardDeliveryQuantityInBaseUnit:        data.StandardDeliveryQuantityInBaseUnit,
		StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
		MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
		MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
		DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
		DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
		StandardDeliveryDuration:                  data.StandardDeliveryDuration,
		StandardDeliveryDurationUnit:              data.StandardDeliveryDurationUnit,
		IsBatchManagementRequired:                 data.IsBatchManagementRequired,
		BatchManagementPolicy:                     data.BatchManagementPolicy,
		ProfitCenter:                              data.ProfitCenter,
	}
}

func ConvertToStorageLocationUpdates(storageLocationUpdates dpfm_api_input_reader.StorageLocation) *StorageLocationUpdates {
	data := storageLocationUpdates

	return &StorageLocationUpdates{
		Product:				data.Product,
		BusinessPartner:		data.BusinessPartner,
		Plant:					data.Plant,
		StorageLocation:		data.StorageLocation,
		BlockStatus:			data.BlockStatus,
	}
}

func ConvertToStorageBinUpdates(storageBinUpdates dpfm_api_input_reader.StorageBin) *StorageBinUpdates {
	data := storageBinUpdates

	return &StorageBinUpdates{
		Product:				data.Product,
		BusinessPartner:		data.BusinessPartner,
		Plant:					data.Plant,
		StorageLocation:		data.StorageLocation,
		StorageBin:				data.StorageBin,
		BlockStatus:			data.BlockStatus,
	}
}

func ConvertToMRPAreaUpdates(mRPAreaUpdates dpfm_api_input_reader.MRPArea) *MRPAreaUpdates {
	data := mRPAreaUpdates

	return &MRPAreaUpdates{
		Product:								   data.Product,
		BusinessPartner:						   data.BusinessPartner,
		Plant:									   data.Plant,
		MRPArea:								   data.MRPArea,
		MRPType:                                   data.MRPType,
		MRPController:                             data.MRPController,
		StorageLocationForMRP:                     data.StorageLocationForMRP,
		ReorderThresholdQuantityInBaseUnit:        data.ReorderThresholdQuantityInBaseUnit,
		PlanningTimeFenceInDays:                   data.PlanningTimeFenceInDays,
		MRPPlanningCalendar:                       data.MRPPlanningCalendar,
		SafetyStockQuantityInBaseUnit:             data.SafetyStockQuantityInBaseUnit,
		SafetyDuration:                            data.SafetyDuration,
		SafetyDurationUnit:                        data.SafetyDurationUnit,
		MaximumStockQuantityInBaseUnit:            data.MaximumStockQuantityInBaseUnit,
		MinumumDeliveryQuantityInBaseUnit:         data.MinumumDeliveryQuantityInBaseUnit,
		MinumumDeliveryLotSizeQuantityInBaseUnit:  data.MinumumDeliveryLotSizeQuantityInBaseUnit,
		StandardDeliveryQuantityInBaseUnit:        data.StandardDeliveryQuantityInBaseUnit,
		StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
		MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
		MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
		DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
		DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
		StandardDeliveryDuration:                  data.StandardDeliveryDuration,
		StandardDeliveryDurationUnit:              data.StandardDeliveryDurationUnit,
	}
}

func ConvertToProductionUpdates(productionUpdates dpfm_api_input_reader.Production) *ProductionUpdates {
	data := productionUpdates

	return &ProductionUpdates{
			Product:					   						data.Product,
			BusinessPartner:			   						data.BusinessPartner,
			Plant:						   						data.Plant,
			ProductionStorageLocation:							data.ProductionStorageLocation,
			ProductProcessingDuration:							data.ProductProcessingDuration,
			ProductProductionQuantityUnit:						data.ProductProductionQuantityUnit,
			MinimumProductionQuantityInBaseUnit:				data.MinimumProductionQuantityInBaseUnit,
			MinimumProductionLotSizeQuantityInBaseUnit:			data.MinimumProductionLotSizeQuantityInBaseUnit,
			StandardProductionQuantityInBaseUnit:				data.StandardProductionQuantityInBaseUnit,
			StandardProductionLotSizeQuantityInBaseUnit:		data.StandardProductionLotSizeQuantityInBaseUnit,
			MaximumProductionQuantityInBaseUnit:				data.MaximumProductionQuantityInBaseUnit,
			MaximumProductionLotSizeQuantityInBaseUnit:			data.MaximumProductionLotSizeQuantityInBaseUnit,
			ProductionLotSizeRoundingQuantityInBaseUnit:		data.ProductionLotSizeRoundingQuantityInBaseUnit,
			MinimumProductionQuantityInProductionUnit:			data.MinimumProductionQuantityInProductionUnit,
			MinimumProductionLotSizeQuantityInProductionUnit:	data.MinimumProductionLotSizeQuantityInProductionUnit,
			StandardProductionQuantityInProductionUnit:			data.StandardProductionQuantityInProductionUnit,
			StandardProductionLotSizeQuantityInProductionUnit:	data.StandardProductionLotSizeQuantityInProductionUnit,
			MaximumProductionLotSizeQuantityInProductionUnit:	data.MaximumProductionLotSizeQuantityInProductionUnit,
			MaximumProductionQuantityInProductionUnit:			data.MaximumProductionQuantityInProductionUnit,
			ProductionLotSizeRoundingQuantityInProductionUnit:	data.ProductionLotSizeRoundingQuantityInProductionUnit,
			ProductionLotSizeIsFixed:							data.ProductionLotSizeIsFixed,
			ProductIsBatchManagedInProductionPlant:				data.ProductIsBatchManagedInProductionPlant,
			ProductIsMarkedForBackflush:						data.ProductIsMarkedForBackflush,
			ProductionSchedulingProfile:						data.ProductionSchedulingProfile,
	}
}

func ConvertToQualityUpdates(qualityUpdates dpfm_api_input_reader.Quality) *QualityUpdates {

	data := qualityUpdates

	return &QualityUpdates{
		Product:						data.Product,
		BusinessPartner:				data.BusinessPartner,
		Plant:							data.Plant,
		QualityMgmtCtrlKey:             data.QualityMgmtCtrlKey,
		ProductSpecification:			data.ProductSpecification,
		MaximumStoragePeriodInDays:     data.MaximumStoragePeriodInDays,
		RecrrgInspIntervalTimeInDays:   data.RecrrgInspIntervalTimeInDays,
		ProductQualityCertificateType:  data.ProductQualityCertificateType,
	}
}

func ConvertToAccountingUpdates(accountingUpdates dpfm_api_input_reader.Accounting) *AccountingUpdates {
	data := accountingUpdates

	return &AccountingUpdates{
		Product:			 data.Product,
		BusinessPartner:	 data.BusinessPartner,
		Plant:				 data.Plant,
		ValuationClass:      data.ValuationClass,
	}
}

func ConvertToProductDescriptionUpdates(productDescriptionUpdates dpfm_api_input_reader.ProductDescription) *ProductDescriptionUpdates {
	data := productDescriptionUpdates

	return &ProductDescriptionUpdates{
		Product:			data.Product,
		ProductDescription: data.ProductDescription,
	}
}

func ConvertToProductDescByBPUpdates(productDescByBPUpdates dpfm_api_input_reader.ProductDescByBP) *ProductDescByBPUpdates {
	data := productDescByBPUpdates

	return &ProductDescByBPUpdates{
		Product:			data.Product,
		BusinessPartner:	data.BusinessPartner,
		Language:           data.Language,
		ProductDescription: data.ProductDescription,
	}
}

func ConvertToTaxUpdates(taxUpdates dpfm_api_input_reader.Tax) *TaxUpdates {
	data := taxUpdates

	return &TaxUpdates{
		Product:				  data.Product,
		Country:				  data.Country,
		ProductTaxCategory:		  data.ProductTaxCategory,
		ProductTaxClassification: data.ProductTaxClassification,
	}
}

func ConvertToAllergenUpdates(allergenUpdates dpfm_api_input_reader.Allergen) *AllergenUpdates {
	data := allergenUpdates

	return &AllergenUpdates{
		Product:				data.Product,
		BusinessPartner:		data.BusinessPartner,
		Allergen:				data.Allergen,
		AllergenIsContained:	data.AllergenIsContained,
	}
}

func ConvertToNutritionalInfoUpdates(nutritionalInfoUpdates dpfm_api_input_reader.NutritionalInfo) *NutritionalInfoUpdates {
	data := nutritionalInfoUpdates

	return &NutritionalInfoUpdates{
		Product:                data.Product,
		BusinessPartner:        data.BusinessPartner,
		Nutrient:				data.Nutrient,
		NutrientContent:		data.NutrientContent,
		NutrientContentUnit:	data.NutrientContentUnit,
	}
}

func ConvertToCaloriesUpdates(caloriesUpdates dpfm_api_input_reader.Calories) *CaloriesUpdates {
	data := caloriesUpdates

	return &CaloriesUpdates{
		Product:				data.Product,
		BusinessPartner:		data.BusinessPartner,
		Calories:				data.Calories,
		CaloryUnit:				data.CaloryUnit,
		CaloryUnitQuantity:		data.CaloryUnitQuantity,
	}
}
