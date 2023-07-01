package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToGeneralCreates(sdc *dpfm_api_input_reader.SDC) (*General, error) {
	data := sdc.General

	general, err := TypeConverter[*General](data)
	if err != nil {
		return nil, err
	}

	return general, nil
}

func ConvertToBusinessPartnerCreates(sdc *dpfm_api_input_reader.SDC) (*[]BusinessPartner, error) {
	businessPartners := make([]BusinessPartner, 0)

	for _, data := range sdc.General.BusinessPartner {
		businessPartner, err := TypeConverter[*BusinessPartner](data)
		if err != nil {
			return nil, err
		}

		businessPartners = append(businessPartners, *businessPartner)
	}

	return &businessPartners, nil
}

func ConvertToAllergenCreates(sdc *dpfm_api_input_reader.SDC) (*[]Allergen, error) {
	allergens := make([]Allergen, 0)

	for _, businessPartner := range sdc.General.BusinessPartner {
		for _, data := range businessPartner.Allergen {
			allergen, err := TypeConverter[*Allergen](data)
			if err != nil {
				return nil, err
			}

			allergens = append(allergens, *allergen)
		}
	}

	return &allergens, nil
}

func ConvertToCaloriesCreates(sdc *dpfm_api_input_reader.SDC) (*[]Calories, error) {
	calorieses := make([]Calories, 0)

	for _, businessPartner := range sdc.General.BusinessPartner {
		for _, data := range businessPartner.Calories {
			calories, err := TypeConverter[*Calories](data)
			if err != nil {
				return nil, err
			}

			calorieses = append(calorieses, *calories)
		}
	}

	return &calorieses, nil
}

func ConvertToNutritionalInfoCreates(sdc *dpfm_api_input_reader.SDC) (*[]NutritionalInfo, error) {
	nutritionalInfos := make([]NutritionalInfo, 0)

	for _, businessPartner := range sdc.General.BusinessPartner {
		for _, data := range businessPartner.NutritionalInfo {
			nutritionalInfo, err := TypeConverter[*NutritionalInfo](data)
			if err != nil {
				return nil, err
			}

			nutritionalInfos = append(nutritionalInfos, *nutritionalInfo)
		}
	}

	return &nutritionalInfos, nil
}

func ConvertToProductDescriptionCreates(sdc *dpfm_api_input_reader.SDC) (*[]ProductDescription, error) {
	productDescriptions := make([]ProductDescription, 0)

	for _, businessPartner := range sdc.General.BusinessPartner {
		for _, data := range businessPartner.ProductDescription {
			productDescription, err := TypeConverter[*ProductDescription](data)
			if err != nil {
				return nil, err
			}

			productDescriptions = append(productDescriptions, *productDescription)
		}
	}

	return &productDescriptions, nil
}

func ConvertToProductDescByBPCreates(sdc *dpfm_api_input_reader.SDC) (*[]ProductDescByBP, error) {
	productDescByBPs := make([]ProductDescByBP, 0)

	for _, businessPartner := range sdc.General.BusinessPartner {
		for _, productDescByBP := range businessPartner.ProductDescription {
			for _, data := range productDescByBP.ProductDescByBP {
				productDescByBP, err := TypeConverter[*ProductDescByBP](data)
				if err != nil {
					return nil, err
				}

				productDescByBPs = append(productDescByBPs, *productDescByBP)
			}
		}
	}

	return &productDescByBPs, nil
}

func ConvertToBPPlantCreates(sdc *dpfm_api_input_reader.SDC) (*[]BPPlant, error) {
	bpPlants := make([]BPPlant, 0)

	for _, data := range sdc.General.BPPlant {
		bpPlant, err := TypeConverter[*BPPlant](data)
		if err != nil {
			return nil, err
		}

		bpPlants = append(bpPlants, *bpPlant)
	}

	return &bpPlants, nil
}

func ConvertToAccountingCreates(sdc *dpfm_api_input_reader.SDC) (*[]Accounting, error) {
	accountings := make([]Accounting, 0)

	for _, bpPlant := range sdc.General.BPPlant {
		for _, data := range bpPlant.Accounting {
			accounting, err := TypeConverter[*Accounting](data)
			if err != nil {
				return nil, err
			}

			accountings = append(accountings, *accounting)
		}
	}

	return &accountings, nil
}

func ConvertToMRPAreaCreates(sdc *dpfm_api_input_reader.SDC) (*[]MRPArea, error) {
	mrpAreas := make([]MRPArea, 0)

	for _, bpPlant := range sdc.General.BPPlant {
		for _, data := range bpPlant.MRPArea {
			mrpArea, err := TypeConverter[*MRPArea](data)
			if err != nil {
				return nil, err
			}

			mrpAreas = append(mrpAreas, *mrpArea)
		}
	}

	return &mrpAreas, nil
}

func ConvertToQualityCreates(sdc *dpfm_api_input_reader.SDC) (*[]Quality, error) {
	qualities := make([]Quality, 0)

	for _, bpPlant := range sdc.General.BPPlant {
		for _, data := range bpPlant.Quality {
			quality, err := TypeConverter[*Quality](data)
			if err != nil {
				return nil, err
			}

			qualities = append(qualities, *quality)
		}
	}

	return &qualities, nil
}

func ConvertToStorageLocationCreates(sdc *dpfm_api_input_reader.SDC) (*[]StorageLocation, error) {
	storageLocations := make([]StorageLocation, 0)

	for _, bpPlant := range sdc.General.BPPlant {
		for _, data := range bpPlant.StorageLocation {
			storageLocation, err := TypeConverter[*StorageLocation](data)
			if err != nil {
				return nil, err
			}

			storageLocations = append(storageLocations, *storageLocation)
		}
	}

	return &storageLocations, nil
}

func ConvertToStorageBinCreates(sdc *dpfm_api_input_reader.SDC) (*[]StorageBin, error) {
	storageBins := make([]StorageBin, 0)

	for _, bpPlant := range sdc.General.BPPlant {
		for _, storageLocation := range bpPlant.StorageLocation {
			for _, data := range storageLocation.StorageBin {
				storageBin, err := TypeConverter[*StorageBin](data)
				if err != nil {
					return nil, err
				}

				storageBins = append(storageBins, *storageBin)
			}
		}
	}

	return &storageBins, nil
}

func ConvertToProductionCreates(sdc *dpfm_api_input_reader.SDC) (*[]Production, error) {
	productions := make([]Production, 0)

	for _, bpPlant := range sdc.General.BPPlant {
		for _, data := range bpPlant.Production {
			production, err := TypeConverter[*Production](data)
			if err != nil {
				return nil, err
			}

			productions = append(productions, *production)
		}
	}

	return &productions, nil
}

func ConvertToTaxCreates(sdc *dpfm_api_input_reader.SDC) (*[]Tax, error) {
	taxes := make([]Tax, 0)

	for _, data := range sdc.General.Tax {
		tax, err := TypeConverter[*Tax](data)
		if err != nil {
			return nil, err
		}

		taxes = append(taxes, *tax)
	}

	return &taxes, nil
}

func ConvertToGeneralUpdates(generalData dpfm_api_input_reader.General) (*General, error) {
	data := generalData

	general, err := TypeConverter[*General](data)
	if err != nil {
		return nil, err
	}

	return general, nil
}

func ConvertToBusinessPartnerUpdates(businessPartnerUpdates *[]dpfm_api_processing_formatter.BusinessPartnerUpdates) (*[]BusinessPartner, error) {
	businessPartners := make([]BusinessPartner, 0)

	for _, data := range *businessPartnerUpdates {
		businessPartner, err := TypeConverter[*BusinessPartner](data)
		if err != nil {
			return nil, err
		}

		businessPartners = append(businessPartners, *businessPartner)
	}

	return &businessPartners, nil
}

func ConvertToAllergenUpdates(allergenUpdates *[]dpfm_api_processing_formatter.AllergenUpdates) (*[]Allergen, error) {
	allergens := make([]Allergen, 0)

	for _, data := range *allergenUpdates {
		allergen, err := TypeConverter[*Allergen](data)
		if err != nil {
			return nil, err
		}

		allergens = append(allergens, *allergen)
	}

	return &allergens, nil
}

func ConvertToCaloriesUpdates(caloriesUpdates *[]dpfm_api_processing_formatter.CaloriesUpdates) (*[]Calories, error) {
	calorieses := make([]Calories, 0)

	for _, data := range *caloriesUpdates {
		calories, err := TypeConverter[*Calories](data)
		if err != nil {
			return nil, err
		}

		calorieses = append(calorieses, *calories)
	}

	return &calorieses, nil
}

func ConvertToNutritionalInfoUpdates(nutritionalInfoUpdates *[]dpfm_api_processing_formatter.NutritionalInfoUpdates) (*[]NutritionalInfo, error) {
	nutritionalInfos := make([]NutritionalInfo, 0)

	for _, data := range *nutritionalInfoUpdates {
		nutritionalInfo, err := TypeConverter[*NutritionalInfo](data)
		if err != nil {
			return nil, err
		}

		nutritionalInfos = append(nutritionalInfos, *nutritionalInfo)
	}

	return &nutritionalInfos, nil
}

func ConvertToProductDescriptionUpdates(productDescriptionUpdates *[]dpfm_api_processing_formatter.ProductDescriptionUpdates) (*[]ProductDescription, error) {
	productDescriptions := make([]ProductDescription, 0)

	for _, data := range *productDescriptionUpdates {
		productDescription, err := TypeConverter[*ProductDescription](data)
		if err != nil {
			return nil, err
		}

		productDescriptions = append(productDescriptions, *productDescription)
	}

	return &productDescriptions, nil
}

func ConvertToProductDescByBPUpdates(productDescByBPUpdates *[]dpfm_api_processing_formatter.ProductDescByBPUpdates) (*[]ProductDescByBP, error) {
	productDescByBPs := make([]ProductDescByBP, 0)

	for _, data := range *productDescByBPUpdates {
		productDescByBP, err := TypeConverter[*ProductDescByBP](data)
		if err != nil {
			return nil, err
		}

		productDescByBPs = append(productDescByBPs, *productDescByBP)
	}

	return &productDescByBPs, nil
}

func ConvertToBPPlantUpdates(bpPlantUpdates *[]dpfm_api_processing_formatter.BPPlantUpdates) (*[]BPPlant, error) {
	bpPlants := make([]BPPlant, 0)

	for _, data := range *bpPlantUpdates {
		bpPlant, err := TypeConverter[*BPPlant](data)
		if err != nil {
			return nil, err
		}

		bpPlants = append(bpPlants, *bpPlant)
	}

	return &bpPlants, nil
}

func ConvertToAccountingUpdates(accountingUpdates *[]dpfm_api_processing_formatter.AccountingUpdates) (*[]Accounting, error) {
	accountings := make([]Accounting, 0)

	for _, data := range *accountingUpdates {
		accounting, err := TypeConverter[*Accounting](data)
		if err != nil {
			return nil, err
		}

		accountings = append(accountings, *accounting)
	}

	return &accountings, nil
}

func ConvertToMRPAreaUpdates(mrpAreaUpdates *[]dpfm_api_processing_formatter.MRPAreaUpdates) (*[]MRPArea, error) {
	mrpAreas := make([]MRPArea, 0)

	for _, data := range *mrpAreaUpdates {
		mrpArea, err := TypeConverter[*MRPArea](data)
		if err != nil {
			return nil, err
		}

		mrpAreas = append(mrpAreas, *mrpArea)
	}

	return &mrpAreas, nil
}

func ConvertToQualityUpdates(qualityUpdates *[]dpfm_api_processing_formatter.QualityUpdates) (*[]Quality, error) {
	qualities := make([]Quality, 0)

	for _, data := range *qualityUpdates {
		quality, err := TypeConverter[*Quality](data)
		if err != nil {
			return nil, err
		}

		qualities = append(qualities, *quality)
	}

	return &qualities, nil
}

func ConvertToStorageLocationUpdates(storageLocationUpdates *[]dpfm_api_processing_formatter.StorageLocationUpdates) (*[]StorageLocation, error) {
	storageLocations := make([]StorageLocation, 0)

	for _, data := range *storageLocationUpdates {
		storageLocation, err := TypeConverter[*StorageLocation](data)
		if err != nil {
			return nil, err
		}

		storageLocations = append(storageLocations, *storageLocation)
	}

	return &storageLocations, nil
}

func ConvertToStorageBinUpdates(storageBinUpdates *[]dpfm_api_processing_formatter.StorageBinUpdates) (*[]StorageBin, error) {
	storageBins := make([]StorageBin, 0)

	for _, data := range *storageBinUpdates {
		storageBin, err := TypeConverter[*StorageBin](data)
		if err != nil {
			return nil, err
		}

		storageBins = append(storageBins, *storageBin)
	}

	return &storageBins, nil
}

func ConvertToProductionUpdates(productionUpdates *[]dpfm_api_processing_formatter.ProductionUpdates) (*[]Production, error) {
	productions := make([]Production, 0)

	for _, data := range *productionUpdates {
		production, err := TypeConverter[*Production](data)
		if err != nil {
			return nil, err
		}

		productions = append(productions, *production)
	}

	return &productions, nil
}

func ConvertToTaxUpdates(taxUpdates *[]dpfm_api_processing_formatter.TaxUpdates) (*[]Tax, error) {
	taxes := make([]Tax, 0)

	for _, data := range *taxUpdates {
		tax, err := TypeConverter[*Tax](data)
		if err != nil {
			return nil, err
		}

		taxes = append(taxes, *tax)
	}

	return &taxes, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
