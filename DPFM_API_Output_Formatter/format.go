package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Processing_Formatter"
)

func ConvertToGeneral(SDC *dpfm_api_input_reader.SDC) *General {
	data := SDC.General

	general := &General{
		Product:                       data.Product,
		ProductType:                   data.ProductType,
		BaseUnit:                      data.BaseUnit,
		ValidityStartDate:             data.ValidityStartDate,
		ProductGroup:                  data.ProductGroup,
		GrossWeight:                   data.GrossWeight,
		NetWeight:                     data.NetWeight,
		WeightUnit:                    data.WeightUnit,
		InternalCapacityQuantity:      data.InternalCapacityQuantity,
		InternalCapacityQuantityUnit:  data.InternalCapacityQuantityUnit,
		SizeOrDimensionText:           data.SizeOrDimensionText,
		ProductStandardID:             data.ProductStandardID,
		IndustryStandardName:          data.IndustryStandardName,
		ItemCategory:                  data.ItemCategory,
		BarcodeType:                   data.BarcodeType,
		CountryOfOrigin:               data.CountryOfOrigin,
		CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
		ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
		CreationDate:                  data.CreationDate,
		LastChangeDate:                data.LastChangeDate,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}

	return general
}

func ConvertToGeneralUpdates(generalUpdates *dpfm_api_processing_formatter.GeneralUpdates) *General {
	data := generalUpdates

	general := &General{
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

	return general
}

func ConvertToBusinessPartner(SDC *dpfm_api_input_reader.SDC) *[]BusinessPartner {
	var businessPartner []BusinessPartner

	for _, data := range SDC.General.BusinessPartner {
		businessPartner = append(businessPartner, BusinessPartner{
			Product:                data.Product,
			BusinessPartner:        data.BusinessPartner,
			ValidityEndDate:        data.ValidityEndDate,
			ValidityStartDate:      data.ValidityStartDate,
			BusinessPartnerProduct: data.BusinessPartnerProduct,
			IsMarkedForDeletion:    data.IsMarkedForDeletion,
		})
	}

	return &businessPartner
}

func ConvertToBusinessPartnerUpdates(businessPartnerUpdates *[]dpfm_api_processing_formatter.BusinessPartnerUpdates) *[]BusinessPartner {
	var businessPartner []BusinessPartner

	for _, data := range *businessPartnerUpdates {
		businessPartner = append(businessPartner, BusinessPartner{
			BusinessPartnerProduct: data.BusinessPartnerProduct,
			IsMarkedForDeletion:    data.IsMarkedForDeletion,
		})
	}

	return &businessPartner
}

func ConvertToAllergen(SDC *dpfm_api_input_reader.SDC) *[]Allergen {
	var allergen []Allergen

	for _, data := range SDC.General.Allergen {
		allergen = append(allergen, Allergen{
			Product:             data.Product,
			BusinessPartner:     data.BusinessPartner,
			Allergen:            data.Allergen,
			AllergenIsContained: data.AllergenIsContained,
		})
	}
	return &allergen
}

func ConvertToAllergenUpdates(allergenUpdates *[]dpfm_api_processing_formatter.AllergenUpdates) *[]Allergen {
	var allergen []Allergen

	for _, data := range *allergenUpdates {
		allergen = append(allergen, Allergen{
			Allergen:            data.Allergen,
			AllergenIsContained: data.AllergenIsContained,
		})
	}

	return &allergen
}

func ConvertToNutritionalInfo(SDC *dpfm_api_input_reader.SDC) *[]NutritionalInfo {
	var nutritionalInfo []NutritionalInfo

	for _, data := range SDC.General.NutritionalInfo {
		nutritionalInfo = append(nutritionalInfo, NutritionalInfo{
			Product:             data.Product,
			BusinessPartner:     data.BusinessPartner,
			Nutrient:            data.Nutrient,
			NutrientContent:     data.NutrientContent,
			NutrientContentUnit: data.NutrientContentUnit,
		})
	}
	return &nutritionalInfo
}

func ConvertToNutritionalInfoUpdates(NutritionalInfoUpdates *[]dpfm_api_processing_formatter.NutritionalInfoUpdates) *[]NutritionalInfo {
	var nutritionalInfo []NutritionalInfo

	for _, data := range *NutritionalInfoUpdates {
		nutritionalInfo = append(nutritionalInfo, NutritionalInfo{
			Nutrient:            data.Nutrient,
			NutrientContent:     data.NutrientContent,
			NutrientContentUnit: data.NutrientContentUnit,
		})
	}

	return &nutritionalInfo
}

func ConvertToCalories(SDC *dpfm_api_input_reader.SDC) *[]Calories {
	var calories []Calories

	for _, data := range SDC.General.Calories {
		calories = append(calories, Calories{
			Product:            data.Product,
			BusinessPartner:    data.BusinessPartner,
			CaloryUnitQuantity: data.CaloryUnitQuantity,
			Calories:           data.Calories,
			CaloryUnit:         data.CaloryUnit,
		})
	}
	return &calories
}

func ConvertToCaloriesUpdates(CaloriesUpdates *[]dpfm_api_processing_formatter.CaloriesUpdates) *[]Calories {
	var calories []Calories

	for _, data := range *CaloriesUpdates {
		calories = append(calories, Calories{
			CaloryUnitQuantity: data.CaloryUnitQuantity,
			Calories:           data.Calories,
			CaloryUnit:         data.CaloryUnit,
		})
	}

	return &calories
}

func ConvertToBPPlant(SDC *dpfm_api_input_reader.SDC) *[]BPPlant {
	var bpPlant []BPPlant

	for _, businessPartnerData := range SDC.General.BusinessPartner {
		for _, data := range businessPartnerData.BPPlant {
			bpPlant = append(bpPlant, BPPlant{
				Product:                                   data.Product,
				BusinessPartner:                           data.BusinessPartner,
				Plant:                                     data.Plant,
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
				ProfitCenter:                              data.ProfitCenter,
				IsMarkedForDeletion:                       data.IsMarkedForDeletion,
			})
		}
	}
	return &bpPlant
}

func ConvertToBPPlantUpdates(bpPlantUpdates *[]dpfm_api_processing_formatter.BPPlantUpdates) *[]BPPlant {
	var bpPlant []BPPlant

	for _, data := range *bpPlantUpdates {
		bpPlant = append(bpPlant, BPPlant{
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
		})
	}

	return &bpPlant
}

func ConvertToProductDescByBP(SDC *dpfm_api_input_reader.SDC) *[]ProductDescByBP {
	var productDescByBP []ProductDescByBP

	for _, data := range SDC.General.ProductDescByBP {
		productDescByBP = append(productDescByBP, ProductDescByBP{
			Product:            data.Product,
			BusinessPartner:    data.BusinessPartner,
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}
	return &productDescByBP
}

func ConvertToProductDescByBPUpdates(productDescByBPUpdates *[]dpfm_api_processing_formatter.ProductDescByBPUpdates) *[]ProductDescByBP {
	var productDescByBP []ProductDescByBP

	for _, data := range *productDescByBPUpdates {
		productDescByBP = append(productDescByBP, ProductDescByBP{
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}

	return &productDescByBP
}

func ConvertToStorageLocation(SDC *dpfm_api_input_reader.SDC) *[]StorageLocation {
	var storageLocation []StorageLocation

	for _, businessPartnerData := range SDC.General.BusinessPartner {
		for _, bpPlantData := range businessPartnerData.BPPlant {
			for _, data := range bpPlantData.StorageLocation {
				storageLocation = append(storageLocation, StorageLocation{
					Product:              data.Product,
					BusinessPartner:      data.BusinessPartner,
					Plant:                data.Plant,
					StorageLocation:      data.StorageLocation,
					CreationDate:         data.CreationDate,
					InventoryBlockStatus: data.InventoryBlockStatus,
					IsMarkedForDeletion:  data.IsMarkedForDeletion,
				})
			}
		}
	}

	return &storageLocation
}

func ConvertToStorageLocationUpdates(storageLocationUpdates *[]dpfm_api_processing_formatter.StorageLocationUpdates) *[]StorageLocation {
	var storageLocation []StorageLocation

	for _, data := range *storageLocationUpdates {
		storageLocation = append(storageLocation, StorageLocation{
			InventoryBlockStatus: data.InventoryBlockStatus,
			IsMarkedForDeletion:  data.IsMarkedForDeletion,
		})
	}

	return &storageLocation
}

func ConvertToMRPArea(SDC *dpfm_api_input_reader.SDC) *[]MRPArea {
	var mRPArea []MRPArea

	for _, businessPartnerData := range SDC.General.BusinessPartner {
		for _, bpPlantData := range businessPartnerData.BPPlant {
			for _, data := range bpPlantData.MRPArea {
				mRPArea = append(mRPArea, MRPArea{
					Product:                                  data.Product,
					BusinessPartner:                          data.BusinessPartner,
					Plant:                                    data.Plant,
					MRPArea:                                  data.MRPArea,
					StorageLocationForMRP:                    data.StorageLocationForMRP,
					MRPType:                                  data.MRPType,
					MRPController:                            data.MRPController,
					ReorderThresholdQuantity:                 data.ReorderThresholdQuantity,
					PlanningTimeFence:                        data.PlanningTimeFence,
					MRPPlanningCalendar:                      data.MRPPlanningCalendar,
					SafetyStockQuantityInBaseUnit:            data.SafetyStockQuantityInBaseUnit,
					SafetyDuration:                           data.SafetyDuration,
					MaximumStockQuantityInBaseUnit:           data.MaximumStockQuantityInBaseUnit,
					MinumumDeliveryQuantityInBaseUnit:        data.MinumumDeliveryQuantityInBaseUnit,
					MinumumDeliveryLotSizeQuantityInBaseUnit: data.MinumumDeliveryLotSizeQuantityInBaseUnit,
					StandardDeliveryLotSizeQuantityInBaseUnit: data.StandardDeliveryLotSizeQuantityInBaseUnit,
					DeliveryLotSizeRoundingQuantityInBaseUnit: data.DeliveryLotSizeRoundingQuantityInBaseUnit,
					MaximumDeliveryLotSizeQuantityInBaseUnit:  data.MaximumDeliveryLotSizeQuantityInBaseUnit,
					MaximumDeliveryQuantityInBaseUnit:         data.MaximumDeliveryQuantityInBaseUnit,
					DeliveryLotSizeIsFixed:                    data.DeliveryLotSizeIsFixed,
					StandardDeliveryDurationInDays:            data.StandardDeliveryDurationInDays,
					IsMarkedForDeletion:                       data.IsMarkedForDeletion,
				})
			}
		}
	}
	return &mRPArea
}

func ConvertToMRPAreaUpdates(mrpAreaUpdates *[]dpfm_api_processing_formatter.MRPAreaUpdates) *[]MRPArea {
	var mrpArea []MRPArea

	for _, data := range *mrpAreaUpdates {
		mrpArea = append(mrpArea, MRPArea{
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
		})
	}

	return &mrpArea
}

func ConvertToWorkScheduling(SDC *dpfm_api_input_reader.SDC) *[]WorkScheduling {
	var workScheduling []WorkScheduling

	for _, businessPartnerData := range SDC.General.BusinessPartner {
		for _, bpPlantData := range businessPartnerData.BPPlant {
			for _, data := range bpPlantData.WorkScheduling {
				workScheduling = append(workScheduling, WorkScheduling{
					Product:                       data.Product,
					BusinessPartner:               data.BusinessPartner,
					Plant:                         data.Plant,
					ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
					ProductProcessingTime:         data.ProductProcessingTime,
					ProductionSupervisor:          data.ProductionSupervisor,
					ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
					ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
					PDTCompIsMarkedForBackflush:   data.PDTCompIsMarkedForBackflush,
					ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
					IsMarkedForDeletion:           data.IsMarkedForDeletion,
				})
			}
		}
	}
	return &workScheduling
}

func ConvertToWorkSchedulingUpdates(workSchedulingUpdates *[]dpfm_api_processing_formatter.WorkSchedulingUpdates) *[]WorkScheduling {
	var workScheduling []WorkScheduling

	for _, data := range *workSchedulingUpdates {
		workScheduling = append(workScheduling, WorkScheduling{
			ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
			ProductProcessingTime:         data.ProductProcessingTime,
			ProductionSupervisor:          data.ProductionSupervisor,
			ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
			ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
			PDTCompIsMarkedForBackflush:   data.PDTCompIsMarkedForBackflush,
			ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}

	return &workScheduling
}

func ConvertToAccounting(SDC *dpfm_api_input_reader.SDC) *[]Accounting {
	var accounting []Accounting

	for _, businessPartnerData := range SDC.General.BusinessPartner {
		for _, bpPlantData := range businessPartnerData.BPPlant {
			for _, data := range bpPlantData.Accounting {
				accounting = append(accounting, Accounting{
					Product:             data.Product,
					BusinessPartner:     data.BusinessPartner,
					Plant:               data.Plant,
					ValuationClass:      data.ValuationClass,
					CostingPolicy:       data.CostingPolicy,
					PriceUnitQty:        data.PriceUnitQty,
					StandardPrice:       data.StandardPrice,
					MovingAveragePrice:  data.MovingAveragePrice,
					PriceLastChangeDate: data.PriceLastChangeDate,
					IsMarkedForDeletion: data.IsMarkedForDeletion,
				})
			}
		}
	}

	return &accounting
}

func ConvertToAccountingUpdates(accountingUpdates *[]dpfm_api_processing_formatter.AccountingUpdates) *[]Accounting {
	var accounting []Accounting

	for _, data := range *accountingUpdates {
		accounting = append(accounting, Accounting{
			ValuationClass:      data.ValuationClass,
			CostingPolicy:       data.CostingPolicy,
			PriceUnitQty:        data.PriceUnitQty,
			StandardPrice:       data.StandardPrice,
			MovingAveragePrice:  data.MovingAveragePrice,
			PriceLastChangeDate: data.PriceLastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}

	return &accounting
}

func ConvertToTax(SDC *dpfm_api_input_reader.SDC) *[]Tax {
	var tax []Tax

	for _, data := range SDC.General.Tax {
		tax = append(tax, Tax{
			Product:                  data.Product,
			Country:                  data.Country,
			ProductTaxCategory:       data.ProductTaxCategory,
			ProductTaxClassification: data.ProductTaxClassification,
		})
	}

	return &tax
}

func ConvertToTaxUpdates(taxUpdates *[]dpfm_api_processing_formatter.TaxUpdates) *[]Tax {
	var tax []Tax

	for _, data := range *taxUpdates {
		tax = append(tax, Tax{
			ProductTaxClassification: data.ProductTaxClassification,
		})
	}

	return &tax
}

func ConvertToProductDescription(SDC *dpfm_api_input_reader.SDC) *[]ProductDescription {
	var productDescription []ProductDescription

	for _, data := range SDC.General.ProductDescription {
		productDescription = append(productDescription, ProductDescription{
			Product:            data.Product,
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}

	return &productDescription
}

func ConvertToProductDescriptionUpdates(productDescriptionUpdates *[]dpfm_api_processing_formatter.ProductDescriptionUpdates) *[]ProductDescription {
	var productDescription []ProductDescription

	for _, data := range *productDescriptionUpdates {
		productDescription = append(productDescription, ProductDescription{
			ProductDescription: data.ProductDescription,
		})
	}

	return &productDescription
}
