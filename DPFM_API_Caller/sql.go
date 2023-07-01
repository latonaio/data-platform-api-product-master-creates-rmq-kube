package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *dpfm_api_output_formatter.General
	var businessPartner *[]dpfm_api_output_formatter.BusinessPartner
	var allergen *[]dpfm_api_output_formatter.Allergen
	var calories *[]dpfm_api_output_formatter.Calories
	var nutritionalInfo *[]dpfm_api_output_formatter.NutritionalInfo
	var productDescription *[]dpfm_api_output_formatter.ProductDescription
	var productDescByBP *[]dpfm_api_output_formatter.ProductDescByBP
	var bPPlant *[]dpfm_api_output_formatter.BPPlant
	var accounting *[]dpfm_api_output_formatter.Accounting
	var mRPArea *[]dpfm_api_output_formatter.MRPArea
	var quality *[]dpfm_api_output_formatter.Quality
	var storageLocation *[]dpfm_api_output_formatter.StorageLocation
	var storageBin *[]dpfm_api_output_formatter.StorageBin
	var production *[]dpfm_api_output_formatter.Production
	var tax *[]dpfm_api_output_formatter.Tax
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalCreateSql(nil, mtx, input, output, errs, log)
		case "BusinessPartner":
			businessPartner = c.businessPartnerCreateSql(nil, mtx, input, output, errs, log)
		case "Allergen":
			allergen = c.allergenCreateSql(nil, mtx, input, output, errs, log)
		case "Calories":
			calories = c.caloriesCreateSql(nil, mtx, input, output, errs, log)
		case "NutritionalInfo":
			nutritionalInfo = c.nutritionalInfoCreateSql(nil, mtx, input, output, errs, log)
		case "ProductDescription":
			productDescription = c.productDescriptionCreateSql(nil, mtx, input, output, errs, log)
		case "ProductDescByBP":
			productDescByBP = c.productDescByBPCreateSql(nil, mtx, input, output, errs, log)
		case "BPPlant":
			bPPlant = c.bpPlantCreateSql(nil, mtx, input, output, errs, log)
		case "Accounting":
			accounting = c.accountingCreateSql(nil, mtx, input, output, errs, log)
		case "MRPArea":
			mRPArea = c.mrpAreaCreateSql(nil, mtx, input, output, errs, log)
		case "Quality":
			quality = c.qualityCreateSql(nil, mtx, input, output, errs, log)
		case "StorageLocation":
			storageLocation = c.storageLocationCreateSql(nil, mtx, input, output, errs, log)
		case "StorageBin":
			storageBin = c.storageBinCreateSql(nil, mtx, input, output, errs, log)
		case "Production":
			production = c.productionCreateSql(nil, mtx, input, output, errs, log)
		case "Tax":
			tax = c.taxCreateSql(nil, mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:            general,
		BusinessPartner:    businessPartner,
		Allergen:           allergen,
		Calories:           calories,
		NutritionalInfo:    nutritionalInfo,
		ProductDescription: productDescription,
		ProductDescByBP:    productDescByBP,
		BPPlant:            bPPlant,
		Accounting:         accounting,
		MRPArea:            mRPArea,
		Quality:            quality,
		StorageLocation:    storageLocation,
		StorageBin:         storageBin,
		Production:     	production,
		Tax:                tax,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *dpfm_api_output_formatter.General
	var businessPartner *[]dpfm_api_output_formatter.BusinessPartner
	var allergen *[]dpfm_api_output_formatter.Allergen
	var calories *[]dpfm_api_output_formatter.Calories
	var nutritionalInfo *[]dpfm_api_output_formatter.NutritionalInfo
	var productDescription *[]dpfm_api_output_formatter.ProductDescription
	var productDescByBP *[]dpfm_api_output_formatter.ProductDescByBP
	var bPPlant *[]dpfm_api_output_formatter.BPPlant
	var accounting *[]dpfm_api_output_formatter.Accounting
	var mRPArea *[]dpfm_api_output_formatter.MRPArea
	var quality *[]dpfm_api_output_formatter.Quality
	var storageLocation *[]dpfm_api_output_formatter.StorageLocation
	var storageBin *[]dpfm_api_output_formatter.StorageBin
	var production *[]dpfm_api_output_formatter.Production
	var tax *[]dpfm_api_output_formatter.Tax
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalUpdateSql(mtx, input, output, errs, log)
		case "BusinessPartner":
			businessPartner = c.businessPartnerUpdateSql(mtx, input, output, errs, log)
		case "Allergen":
			allergen = c.allergenUpdateSql(mtx, input, output, errs, log)
		case "Calories":
			calories = c.caloriesUpdateSql(mtx, input, output, errs, log)
		case "NutritionalInfo":
			nutritionalInfo = c.nutritionalInfoUpdateSql(mtx, input, output, errs, log)
		case "ProductDescription":
			productDescription = c.productDescriptionUpdateSql(mtx, input, output, errs, log)
		case "ProductDescByBP":
			productDescByBP = c.productDescByBPUpdateSql(mtx, input, output, errs, log)
		case "BPPlant":
			bPPlant = c.bpPlantUpdateSql(mtx, input, output, errs, log)
		case "Accounting":
			accounting = c.accountingUpdateSql(mtx, input, output, errs, log)
		case "MRPArea":
			mRPArea = c.mrpAreaUpdateSql(mtx, input, output, errs, log)
		case "Quality":
			quality = c.qualityUpdateSql(mtx, input, output, errs, log)
		case "StorageLocation":
			storageLocation = c.storageLocationUpdateSql(mtx, input, output, errs, log)
		case "StorageBin":
			storageBin = c.storageBinUpdateSql(mtx, input, output, errs, log)
		case "Production":
			production = c.productionUpdateSql(mtx, input, output, errs, log)
		case "Tax":
			tax = c.taxUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:            general,
		BusinessPartner:    businessPartner,
		Allergen:           allergen,
		Calories:           calories,
		NutritionalInfo:    nutritionalInfo,
		ProductDescription: productDescription,
		ProductDescByBP:    productDescByBP,
		BPPlant:            bPPlant,
		Accounting:         accounting,
		MRPArea:            mRPArea,
		Quality:            quality,
		StorageLocation:    storageLocation,
		StorageBin:         storageBin,
		Production:     	production,
		Tax:                tax,
	}

	return data
}

func (c *DPFMAPICaller) generalCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID

	generalData := input.General
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": generalData, "function": "ProductMasterGeneral", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "General Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToGeneralCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) businessPartnerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BusinessPartner {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, businessPartnerData := range input.General.BusinessPartner {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": businessPartnerData, "function": "ProductMasterBusinessPartner", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "BusinessPartner Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToBusinessPartnerCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) allergenCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Allergen {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, businessPartner := range input.General.BusinessPartner {
		for _, allergenData := range businessPartner.Allergen {
			res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": allergenData, "function": "ProductMasterAllergen", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Allergen Data cannot insert"
				return nil
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAllergenCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) caloriesCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Calories {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, businessPartner := range input.General.BusinessPartner {
		for _, caloriesData := range businessPartner.Calories {
			res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": caloriesData, "function": "ProductMasterCalories", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Calories Data cannot insert"
				return nil
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToCaloriesCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) nutritionalInfoCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.NutritionalInfo {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, businessPartner := range input.General.BusinessPartner {
		for _, nutritionalInfoData := range businessPartner.NutritionalInfo {
			res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": nutritionalInfoData, "function": "ProductMasterNutritionalInfo", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Nutritional Info Data cannot insert"
				return nil
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToNutritionalInfoCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) productDescriptionCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductDescription {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, businessPartner := range input.General.BusinessPartner {
		for _, productDescriptionData := range businessPartner.ProductDescription {
			res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productDescriptionData, "function": "ProductMasterProductDescription", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "ProductDescription Data cannot insert"
				return nil
			}

		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToProductDescriptionCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) productDescByBPCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductDescByBP {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, businessPartner := range input.General.BusinessPartner {
		for _, productDescription := range businessPartner.ProductDescription {
			for _, productDescByBPData := range productDescription.ProductDescByBP {
				res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productDescByBPData, "function": "ProductMasterProductDescByBP", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "ProductDescByBP Data cannot insert"
					return nil
				}
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToProductDescByBPCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) bpPlantCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BPPlant {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, bPPlantData := range input.General.BPPlant {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": bPPlantData, "function": "ProductMasterBPPlant", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "BPPlant Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToBPPlantCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) accountingCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Accounting {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, bPPlantData := range input.General.BPPlant {
		for _, accountingData := range bPPlantData.Accounting {
			res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": accountingData, "function": "ProductMasterAccounting", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Accounting Data cannot insert"
				return nil
			}

		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAccountingCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) mrpAreaCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.MRPArea {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, bPPlantData := range input.General.BPPlant {
		for _, mRPAreaData := range bPPlantData.MRPArea {
			res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": mRPAreaData, "function": "ProductMasterMRPArea", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "MRPArea Data cannot insert"
				return nil
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToMRPAreaCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) qualityCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Quality {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, bPPlantData := range input.General.BPPlant {
		for _, qualityData := range bPPlantData.Quality {
			res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": qualityData, "function": "ProductMasterQuality", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Quality Data cannot insert"
				return nil
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToQualityCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) storageLocationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StorageLocation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, bPPlantData := range input.General.BPPlant {
		for _, storageLocationData := range bPPlantData.StorageLocation {
			res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": storageLocationData, "function": "ProductMasterStorageLocation", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "StorageLocation Data cannot insert"
				return nil
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToStorageLocationCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) storageBinCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StorageBin {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, bPPlantData := range input.General.BPPlant {
		for _, storageLocationData := range bPPlantData.StorageLocation {
			for _, storageBinData := range storageLocationData.StorageBin {
				res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": storageBinData, "function": "ProductMasterStorageBin", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "StorageBin Data cannot insert"
					return nil
				}
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToStorageBinCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) productionCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Production {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, bPPlantData := range input.General.BPPlant {
		for _, productionData := range bPPlantData.Production {
			res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productionData, "function": "ProductMasterProduction", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Production Data cannot insert"
				return nil
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToProductionCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) taxCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Tax {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, taxData := range input.General.Tax {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": taxData, "function": "ProductMasterTax", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Tax Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToTaxCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) generalUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	general := input.General
	generalData := dpfm_api_processing_formatter.ConvertToGeneralUpdates(general)

	sessionID := input.RuntimeSessionID
	if generalIsUpdate(generalData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": generalData, "function": "ProductMasterGeneral", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "General Data cannot insert"
			return nil
		}
		if output.SQLUpdateResult == nil {
			output.SQLUpdateResult = getBoolPtr(true)
		}
	}

	data, err := dpfm_api_output_formatter.ConvertToGeneralUpdates(general)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) businessPartnerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BusinessPartner {
	req := make([]dpfm_api_processing_formatter.BusinessPartnerUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, businessPartner := range general.BusinessPartner {
		businessPartnerData := *dpfm_api_processing_formatter.ConvertToBusinessPartnerUpdates(businessPartner)

		if businessPartnerIsUpdate(&businessPartnerData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": businessPartnerData, "function": "ProductMasterBusinessPartner", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "BusinessPartner Data cannot update"
				return nil
			}
		}
		req = append(req, businessPartnerData)
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToBusinessPartnerUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) allergenUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Allergen {
	req := make([]dpfm_api_processing_formatter.AllergenUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, businessPartner := range general.BusinessPartner {
		for _, allergen := range businessPartner.Allergen {
			allergenData := *dpfm_api_processing_formatter.ConvertToAllergenUpdates(allergen)

			if allergenIsUpdate(&allergenData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": allergenData, "function": "ProductMasterAllergen", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Allergen Data cannot update"
					return nil
				}
			}
			req = append(req, allergenData)
		}
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAllergenUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) caloriesUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Calories {
	req := make([]dpfm_api_processing_formatter.CaloriesUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, businessPartner := range general.BusinessPartner {
		for _, calories := range businessPartner.Calories {
			caloriesData := *dpfm_api_processing_formatter.ConvertToCaloriesUpdates(calories)

			if caloriesIsUpdate(&caloriesData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": caloriesData, "function": "ProductMasterCalories", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Calories Data cannot update"
					return nil
				}
			}
			req = append(req, caloriesData)
		}
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToCaloriesUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) nutritionalInfoUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.NutritionalInfo {
	req := make([]dpfm_api_processing_formatter.NutritionalInfoUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, businessPartner := range general.BusinessPartner {
		for _, nutritionalInfo := range businessPartner.NutritionalInfo {
			nutritionalInfoData := *dpfm_api_processing_formatter.ConvertToNutritionalInfoUpdates(nutritionalInfo)

			if nutritionalInfoIsUpdate(&nutritionalInfoData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": nutritionalInfoData, "function": "ProductMasterNutritionalInfo", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "NutritionalInfo Data cannot update"
					return nil
				}
			}
			req = append(req, nutritionalInfoData)
		}
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToNutritionalInfoUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) productDescriptionUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductDescription {
	req := make([]dpfm_api_processing_formatter.ProductDescriptionUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, businessPartner := range general.BusinessPartner {
		for _, productDescription := range businessPartner.ProductDescription {
			productDescriptionData := *dpfm_api_processing_formatter.ConvertToProductDescriptionUpdates(productDescription)

			if productDescriptionIsUpdate(&productDescriptionData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productDescriptionData, "function": "ProductMasterProductDescription", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "ProductDescription Data cannot update"
					return nil
				}
			}
			req = append(req, productDescriptionData)
		}
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToProductDescriptionUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) productDescByBPUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductDescByBP {
	req := make([]dpfm_api_processing_formatter.ProductDescByBPUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, businessPartner := range general.BusinessPartner {
		for _, productDescription := range businessPartner.ProductDescription {
			for _, productDescByBP := range productDescription.ProductDescByBP {
				productDescByBPData := *dpfm_api_processing_formatter.ConvertToProductDescByBPUpdates(productDescByBP)

				if productDescByBPIsUpdate(&productDescByBPData) {
					res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productDescByBPData, "function": "ProductMasterProductDescByBP", "runtime_session_id": sessionID})
					if err != nil {
						err = xerrors.Errorf("rmq error: %w", err)
						*errs = append(*errs, err)
						return nil
					}
					res.Success()
					if !checkResult(res) {
						output.SQLUpdateResult = getBoolPtr(false)
						output.SQLUpdateError = "ProductDescByBP Data cannot update"
						return nil
					}
				}
				req = append(req, productDescByBPData)
			}
		}
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToProductDescByBPUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) bpPlantUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BPPlant {
	req := make([]dpfm_api_processing_formatter.BPPlantUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, bpPlant := range general.BPPlant {
		bpPlantData := *dpfm_api_processing_formatter.ConvertToBPPlantUpdates(bpPlant)

		if bpPlantIsUpdate(&bpPlantData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": bpPlantData, "function": "ProductMasterBPPlant", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "BPPlant Data cannot update"
				return nil
			}
		}
		req = append(req, bpPlantData)
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToBPPlantUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) accountingUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Accounting {
	req := make([]dpfm_api_processing_formatter.AccountingUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, bpPlant := range general.BPPlant {
		for _, accounting := range bpPlant.Accounting {
			accountingData := *dpfm_api_processing_formatter.ConvertToAccountingUpdates(accounting)

			if accountingIsUpdate(&accountingData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": accountingData, "function": "ProductMasterAccounting", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Accounting Data cannot update"
					return nil
				}
			}
			req = append(req, accountingData)
		}
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToAccountingUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) mrpAreaUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.MRPArea {
	req := make([]dpfm_api_processing_formatter.MRPAreaUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, bpPlant := range general.BPPlant {
		for _, mrpArea := range bpPlant.MRPArea {
			mrpAreaData := *dpfm_api_processing_formatter.ConvertToMRPAreaUpdates(mrpArea)

			if mrpAreaIsUpdate(&mrpAreaData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": mrpAreaData, "function": "ProductMasterMRPArea", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "MRPArea Data cannot update"
					return nil
				}
			}
			req = append(req, mrpAreaData)
		}
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToMRPAreaUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) qualityUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Quality {
	req := make([]dpfm_api_processing_formatter.QualityUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, bpPlant := range general.BPPlant {
		for _, quality := range bpPlant.Quality {
			qualityData := *dpfm_api_processing_formatter.ConvertToQualityUpdates(quality)

			if qualityIsUpdate(&qualityData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": qualityData, "function": "ProductMasterQuality", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Quality Data cannot update"
					return nil
				}
			}
			req = append(req, qualityData)
		}
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToQualityUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) storageLocationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StorageLocation {
	req := make([]dpfm_api_processing_formatter.StorageLocationUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, bpPlant := range general.BPPlant {
		for _, storageLocation := range bpPlant.StorageLocation {
			storageLocationData := *dpfm_api_processing_formatter.ConvertToStorageLocationUpdates(storageLocation)

			if storageLocationIsUpdate(&storageLocationData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": storageLocationData, "function": "ProductMasterStorageLocation", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "StorageLocation Data cannot update"
					return nil
				}
			}
			req = append(req, storageLocationData)
		}
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToStorageLocationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) storageBinUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StorageBin {
	req := make([]dpfm_api_processing_formatter.StorageBinUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, bpPlant := range general.BPPlant {
		for _, storageLocation := range bpPlant.StorageLocation {
			for _, storageBin := range storageLocation.StorageBin {
				storageBinData := *dpfm_api_processing_formatter.ConvertToStorageBinUpdates(storageBin)

				if storageBinIsUpdate(&storageBinData) {
					res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": storageBinData, "function": "ProductMasterStorageBin", "runtime_session_id": sessionID})
					if err != nil {
						err = xerrors.Errorf("rmq error: %w", err)
						*errs = append(*errs, err)
						return nil
					}
					res.Success()
					if !checkResult(res) {
						output.SQLUpdateResult = getBoolPtr(false)
						output.SQLUpdateError = "StorageBin Data cannot update"
						return nil
					}
				}
				req = append(req, storageBinData)
			}
		}
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToStorageBinUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) productionUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Production {
	req := make([]dpfm_api_processing_formatter.ProductionUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, bpPlant := range general.BPPlant {
		for _, production := range bpPlant.Production {
			productionData := *dpfm_api_processing_formatter.ConvertToProductionUpdates(production)

			if productionIsUpdate(&productionData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productionData, "function": "ProductMasterProduction", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Production Data cannot update"
					return nil
				}
			}
			req = append(req, productionData)
		}
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToProductionUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) taxUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Tax {
	req := make([]dpfm_api_processing_formatter.TaxUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, tax := range general.Tax {
		taxData := *dpfm_api_processing_formatter.ConvertToTaxUpdates(tax)

		if taxIsUpdate(&taxData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": taxData, "function": "ProductMasterTax", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Tax Data cannot update"
				return nil
			}
		}
		req = append(req, taxData)
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToTaxUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func generalIsUpdate(general *dpfm_api_processing_formatter.GeneralUpdates) bool {
	product := general.Product

	return !(len(product) == 0)
}

func businessPartnerIsUpdate(businessPartner *dpfm_api_processing_formatter.BusinessPartnerUpdates) bool {
	product := businessPartner.Product
	bp := businessPartner.BusinessPartner

	return !(len(product) == 0 || bp == 0)
}

func allergenIsUpdate(allergen *dpfm_api_processing_formatter.AllergenUpdates) bool {
	product := allergen.Product
	bp := allergen.BusinessPartner
	al := allergen.Allergen

	return !(len(product) == 0 || bp == 0 || len(al) == 0)
}

func caloriesIsUpdate(calories *dpfm_api_processing_formatter.CaloriesUpdates) bool {
	product := calories.Product
	bp := calories.BusinessPartner
	caloryUnitQuantity := calories.CaloryUnitQuantity

	return !(len(product) == 0 || bp == 0 || caloryUnitQuantity == 0)
}

func nutritionalInfoIsUpdate(nutritionalInfo *dpfm_api_processing_formatter.NutritionalInfoUpdates) bool {
	product := nutritionalInfo.Product
	bp := nutritionalInfo.BusinessPartner
	nutrient := nutritionalInfo.Nutrient

	return !(len(product) == 0 || bp == 0 || len(nutrient) == 0)
}

func productDescriptionIsUpdate(productDescription *dpfm_api_processing_formatter.ProductDescriptionUpdates) bool {
	product := productDescription.Product
	language := productDescription.Language

	return !(len(product) == 0 || len(language) == 0)
}

func productDescByBPIsUpdate(productDescByBP *dpfm_api_processing_formatter.ProductDescByBPUpdates) bool {
	product := productDescByBP.Product
	bp := productDescByBP.BusinessPartner
	language := productDescByBP.Language

	return !(len(product) == 0 || bp == 0 || len(language) == 0)
}

func bpPlantIsUpdate(bpPlant *dpfm_api_processing_formatter.BPPlantUpdates) bool {
	product := bpPlant.Product
	bp := bpPlant.BusinessPartner
	plant := bpPlant.Plant

	return !(len(product) == 0 || bp == 0 || len(plant) == 0)
}

func accountingIsUpdate(accounting *dpfm_api_processing_formatter.AccountingUpdates) bool {
	product := accounting.Product
	bp := accounting.BusinessPartner
	plant := accounting.Plant

	return !(len(product) == 0 || bp == 0 || len(plant) == 0)
}

func mrpAreaIsUpdate(mrpArea *dpfm_api_processing_formatter.MRPAreaUpdates) bool {
	product := mrpArea.Product
	bp := mrpArea.BusinessPartner
	plant := mrpArea.Plant

	return !(len(product) == 0 || bp == 0 || len(plant) == 0)
}

func qualityIsUpdate(quality *dpfm_api_processing_formatter.QualityUpdates) bool {
	product := quality.Product
	bp := quality.BusinessPartner
	plant := quality.Plant

	return !(len(product) == 0 || bp == 0 || len(plant) == 0)
}

func storageLocationIsUpdate(storageLocation *dpfm_api_processing_formatter.StorageLocationUpdates) bool {
	product := storageLocation.Product
	bp := storageLocation.BusinessPartner
	plant := storageLocation.Plant

	return !(len(product) == 0 || bp == 0 || len(plant) == 0)
}

func storageBinIsUpdate(storageBin *dpfm_api_processing_formatter.StorageBinUpdates) bool {
	product := storageBin.Product
	bp := storageBin.BusinessPartner
	plant := storageBin.Plant

	return !(len(product) == 0 || bp == 0 || len(plant) == 0)
}

func workSchedulingIsUpdate(WorkScheduling *dpfm_api_processing_formatter.WorkSchedulingUpdates) bool {
	product := WorkScheduling.Product
	bp := WorkScheduling.BusinessPartner
	plant := WorkScheduling.Plant

	return !(len(product) == 0 || bp == 0 || len(plant) == 0)
}

func taxIsUpdate(tax *dpfm_api_processing_formatter.TaxUpdates) bool {
	product := tax.Product
	country := tax.Country
	productTaxCategory := tax.ProductTaxCategory

	return !(len(product) == 0 || len(country) == 0 || len(productTaxCategory) == 0)
}
