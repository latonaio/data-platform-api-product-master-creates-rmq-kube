package dpfm_api_output_formatter

import (
	"data-platform-api-product-master-creates-rmq-kube/sub_func_complementer"
)

func ConvertToGeneral(subfuncSDC *sub_func_complementer.SDC) *General {
	data := subfuncSDC.Message.General

	general := &General{
		Product:                       data.Product,
		ProductType:                   data.ProductType,
		BaseUnit:                      data.BaseUnit,
		ValidityStartDate:             data.ValidityStartDate,
		ProductGroup:                  data.ProductGroup,
		Division:                      data.Division,
		GrossWeight:                   data.GrossWeight,
		WeightUnit:                    data.WeightUnit,
		SizeOrDimensionText:           data.SizeOrDimensionText,
		IndustryStandardName:          data.IndustryStandardName,
		ProductStandardID:             data.ProductStandardID,
		CreationDate:                  data.CreationDate,
		LastChangeDate:                data.LastChangeDate,
		NetWeight:                     data.NetWeight,
		CountryOfOrigin:               data.CountryOfOrigin,
		CountryOfOriginLanguage:       data.CountryOfOriginLanguage,
		ItemCategory:                  data.ItemCategory,
		ProductAccountAssignmentGroup: data.ProductAccountAssignmentGroup,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}

	return general
}

func ConvertToBusinessPartner(subfuncSDC *sub_func_complementer.SDC) *[]BusinessPartner {
	var businessPartner []BusinessPartner

	for _, data := range *subfuncSDC.Message.BusinessPartner {
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

func ConvertToBPPlant(subfuncSDC *sub_func_complementer.SDC) *[]BPPlant {
	var bpPlant []BPPlant

	for _, data := range *subfuncSDC.Message.BPPlant {
		bpPlant = append(bpPlant, BPPlant{
			Product:                                   data.Product,
			BusinessPartner:                           data.BusinessPartner,
			Plant:                                     data.Plant,
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
			ProfitCenter:                              data.ProfitCenter,
			IsMarkedForDeletion:                       data.IsMarkedForDeletion,
		})
	}
	return &bpPlant
}

func ConvertToProductDescByBP(subfuncSDC *sub_func_complementer.SDC) *[]ProductDescByBP {
	var productDescByBP []ProductDescByBP

	for _, data := range *subfuncSDC.Message.ProductDescByBP {
		productDescByBP = append(productDescByBP, ProductDescByBP{
			Product:            data.Product,
			BusinessPartner:    data.BusinessPartner,
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}
	return &productDescByBP
}

func ConvertToStorageLocation(subfuncSDC *sub_func_complementer.SDC) *[]StorageLocation {
	var storageLocation []StorageLocation

	for _, data := range *subfuncSDC.Message.StorageLocation {
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

	return &storageLocation
}

func ConvertToMRPArea(subfuncSDC *sub_func_complementer.SDC) *[]MRPArea {
	var mRPArea []MRPArea

	for _, data := range *subfuncSDC.Message.MRPArea {
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
	return &mRPArea
}

func ConvertToWorkScheduling(subfuncSDC *sub_func_complementer.SDC) *[]WorkScheduling {
	var workScheduling []WorkScheduling

	for _, data := range *subfuncSDC.Message.WorkScheduling {
		workScheduling = append(workScheduling, WorkScheduling{
			Product:                       data.Product,
			BusinessPartner:               data.BusinessPartner,
			Plant:                         data.Plant,
			ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
			ProductProcessingTime:         data.ProductProcessingTime,
			ProductionSupervisor:          data.ProductionSupervisor,
			ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
			ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
			MatlCompIsMarkedForBackflush:  data.MatlCompIsMarkedForBackflush,
			ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}
	return &workScheduling
}

func ConvertToAccounting(subfuncSDC *sub_func_complementer.SDC) *[]Accounting {
	var accounting []Accounting

	for _, data := range *subfuncSDC.Message.Accounting {
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

	return &accounting
}

func ConvertToTax(subfuncSDC *sub_func_complementer.SDC) *[]Tax {
	var tax []Tax

	for _, data := range *subfuncSDC.Message.Tax {
		tax = append(tax, Tax{
			Product:                  data.Product,
			Country:                  data.Country,
			ProductTaxCategory:       data.ProductTaxCategory,
			ProductTaxClassification: data.ProductTaxClassification,
		})
	}

	return &tax
}

func ConvertToProductDescription(subfuncSDC *sub_func_complementer.SDC) []ProductDescription {
	var productDescription []ProductDescription

	for _, data := range *subfuncSDC.Message.ProductDescription {
		productDescription = append(productDescription, ProductDescription{
			Product:            data.Product,
			Language:           data.Language,
			ProductDescription: data.ProductDescription,
		})
	}

	return productDescription
}
