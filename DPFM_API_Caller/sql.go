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
	var bPPlant *[]dpfm_api_output_formatter.BPPlant
	var allergen *[]dpfm_api_output_formatter.Allergen
	var calories *[]dpfm_api_output_formatter.Calories
	var nutritionalInfo *[]dpfm_api_output_formatter.NutritionalInfo
	var productDescByBP *[]dpfm_api_output_formatter.ProductDescByBP
	var productDescription *[]dpfm_api_output_formatter.ProductDescription
	var storageLocation *[]dpfm_api_output_formatter.StorageLocation
	var workScheduling *[]dpfm_api_output_formatter.WorkScheduling
	var mRPArea *[]dpfm_api_output_formatter.MRPArea
	var accounting *[]dpfm_api_output_formatter.Accounting
	var tax *[]dpfm_api_output_formatter.Tax
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalCreateSql(nil, mtx, input, output, errs, log)
		case "BusinessPartner":
			businessPartner = c.businessPartnerCreateSql(nil, mtx, input, output, errs, log)
		case "BPPlant":
			bPPlant = c.bPPlantCreateSql(nil, mtx, input, output, errs, log)
		case "Allergen":
			allergen = c.allergenCreateSql(nil, mtx, input, output, errs, log)
		case "Calories":
			calories = c.caloriesCreateSql(nil, mtx, input, output, errs, log)
		case "NutritionalInfo":
			nutritionalInfo = c.nutritionalInfoCreateSql(nil, mtx, input, output, errs, log)
		case "ProductDescByBP":
			productDescByBP = c.productDescByBPCreateSql(nil, mtx, input, output, errs, log)
		case "ProductDescription":
			productDescription = c.productDescriptionCreateSql(nil, mtx, input, output, errs, log)
		case "StorageLocation":
			storageLocation = c.storageLocationCreateSql(nil, mtx, input, output, errs, log)
		case "WorkScheduling":
			workScheduling = c.workSchedulingCreateSql(nil, mtx, input, output, errs, log)
		case "MRPArea":
			mRPArea = c.mRPAreaCreateSql(nil, mtx, input, output, errs, log)
		case "Accounting":
			accounting = c.accountingCreateSql(nil, mtx, input, output, errs, log)
		case "Tax":
			tax = c.taxCreateSql(nil, mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:            general,
		BusinessPartner:    businessPartner,
		BPPlant:            bPPlant,
		Allergen:           allergen,
		Calories:           calories,
		NutritionalInfo:    nutritionalInfo,
		ProductDescByBP:    productDescByBP,
		ProductDescription: productDescription,
		StorageLocation:    storageLocation,
		WorkScheduling:     workScheduling,
		MRPArea:            mRPArea,
		Accounting:         accounting,
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
	var bPPlant *[]dpfm_api_output_formatter.BPPlant
	var allergen *[]dpfm_api_output_formatter.Allergen
	var calories *[]dpfm_api_output_formatter.Calories
	var nutritionalInfo *[]dpfm_api_output_formatter.NutritionalInfo
	var productDescByBP *[]dpfm_api_output_formatter.ProductDescByBP
	var productDescription *[]dpfm_api_output_formatter.ProductDescription
	var storageLocation *[]dpfm_api_output_formatter.StorageLocation
	var workScheduling *[]dpfm_api_output_formatter.WorkScheduling
	var mRPArea *[]dpfm_api_output_formatter.MRPArea
	var accounting *[]dpfm_api_output_formatter.Accounting
	var tax *[]dpfm_api_output_formatter.Tax
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalUpdateSql(mtx, input, output, errs, log)
		case "BusinessPartner":
			businessPartner = c.businessPartnerUpdateSql(mtx, input, output, errs, log)
		case "BPPlant":
			bPPlant = c.bPPlantUpdateSql(mtx, input, output, errs, log)
		case "Allergen":
			allergen = c.allergenUpdateSql(mtx, input, output, errs, log)
		case "Calories":
			calories = c.caloriesUpdateSql(mtx, input, output, errs, log)
		case "NutritionalInfo":
			nutritionalInfo = c.nutritionalInfoUpdateSql(mtx, input, output, errs, log)
		case "ProductDescByBP":
			productDescByBP = c.productDescByBPUpdateSql(mtx, input, output, errs, log)
		case "ProductDescription":
			productDescription = c.productDescriptionUpdateSql(mtx, input, output, errs, log)
		case "StorageLocation":
			storageLocation = c.storageLocationUpdateSql(mtx, input, output, errs, log)
		case "WorkScheduling":
			workScheduling = c.workSchedulingUpdateSql(mtx, input, output, errs, log)
		case "MRPArea":
			mRPArea = c.mRPAreaUpdateSql(mtx, input, output, errs, log)
		case "Accounting":
			accounting = c.accountingUpdateSql(mtx, input, output, errs, log)
		case "Tax":
			tax = c.taxUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:            general,
		BusinessPartner:    businessPartner,
		BPPlant:            bPPlant,
		Allergen:           allergen,
		Calories:           calories,
		NutritionalInfo:    nutritionalInfo,
		ProductDescByBP:    productDescByBP,
		ProductDescription: productDescription,
		StorageLocation:    storageLocation,
		WorkScheduling:     workScheduling,
		MRPArea:            mRPArea,
		Accounting:         accounting,
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

	data := dpfm_api_output_formatter.ConvertToGeneral(input)

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
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToBusinessPartner(input)

	return data
}

func (c *DPFMAPICaller) bPPlantCreateSql(
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
	for _, businessPartnerData := range input.General.BusinessPartner {
		bPPlantData := businessPartnerData.BPPlant
		for _, bPPlantData := range bPPlantData {
			res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": bPPlantData, "function": "ProductMasterBPPlant", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Item Data cannot insert"
				return nil
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToBPPlant(input)

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
	for _, allergenData := range input.General.Allergen {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": allergenData, "function": "ProductMasterAllergen", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToAllergen(input)

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
	for _, caloriesData := range input.General.Calories {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": caloriesData, "function": "ProductMasterCalories", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToCalories(input)

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
	for _, nutritionalInfoData := range input.General.NutritionalInfo {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": nutritionalInfoData, "function": "ProductMasterNutritionalInfo", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToNutritionalInfo(input)

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
	for _, productDescriptionData := range input.General.ProductDescription {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productDescriptionData, "function": "ProductMasterProductDescription", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToProductDescription(input)

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
	for _, productDescByBPData := range input.General.ProductDescByBP {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productDescByBPData, "function": "ProductMasterProductDescByBP", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToProductDescByBP(input)

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
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToTax(input)

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
	for _, businessPartnerData := range input.General.BusinessPartner {
		bPPlantData := businessPartnerData.BPPlant
		for _, bPPlantData := range bPPlantData {
			for _, storageLocationData := range bPPlantData.StorageLocation {
				res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": storageLocationData, "function": "ProductMasterStorageLocation", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Item Data cannot insert"
					return nil
				}
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToStorageLocation(input)

	return data
}

func (c *DPFMAPICaller) workSchedulingCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.WorkScheduling {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, businessPartnerData := range input.General.BusinessPartner {
		bPPlantData := businessPartnerData.BPPlant
		for _, bPPlantData := range bPPlantData {
			for _, workSchedulingData := range bPPlantData.WorkScheduling {
				res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": workSchedulingData, "function": "ProductMasterWorkScheduling", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Item Data cannot insert"
					return nil
				}
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToWorkScheduling(input)

	return data
}

func (c *DPFMAPICaller) mRPAreaCreateSql(
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
	for _, businessPartnerData := range input.General.BusinessPartner {
		bPPlantData := businessPartnerData.BPPlant
		for _, bPPlantData := range bPPlantData {
			for _, mRPAreaData := range bPPlantData.MRPArea {
				res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": mRPAreaData, "function": "ProductMasterMRPArea", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Item Data cannot insert"
					return nil
				}
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToMRPArea(input)

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
	for _, businessPartnerData := range input.General.BusinessPartner {
		bPPlantData := businessPartnerData.BPPlant
		for _, bPPlantData := range bPPlantData {
			for _, accountingData := range bPPlantData.Accounting {
				res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": accountingData, "function": "ProductMasterAccounting", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Item Data cannot insert"
					return nil
				}
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToAccounting(input)

	return data
}

func (c *DPFMAPICaller) generalUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	req := dpfm_api_processing_formatter.ConvertToGeneralUpdates(input.General)

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterGeneral", "runtime_session_id": 123})
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

	data := dpfm_api_output_formatter.ConvertToGeneralUpdates(req)

	return data
}

func (c *DPFMAPICaller) businessPartnerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BusinessPartner {
	var req *[]dpfm_api_processing_formatter.BusinessPartnerUpdates
	for _, v := range input.General.BusinessPartner {
		*req = append(*req, *dpfm_api_processing_formatter.ConvertToBusinessPartnerUpdates(v))
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterBusinessPartner", "runtime_session_id": 123})
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
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToBusinessPartnerUpdates(req)

	return data
}

func (c *DPFMAPICaller) bPPlantUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BPPlant {
	var req *[]dpfm_api_processing_formatter.BPPlantUpdates
	for _, businessPartner := range input.General.BusinessPartner {
		for _, v := range businessPartner.BPPlant {
			*req = append(*req, *dpfm_api_processing_formatter.ConvertToBPPlantUpdates(v))
		}
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterBPPlant", "runtime_session_id": 123})
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
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToBPPlantUpdates(req)

	return data
}

func (c *DPFMAPICaller) allergenUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Allergen {
	var req *[]dpfm_api_processing_formatter.AllergenUpdates
	for _, v := range input.General.Allergen {
		*req = append(*req, *dpfm_api_processing_formatter.ConvertToAllergenUpdates(v))
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterAllergen", "runtime_session_id": 123})
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
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToAllergenUpdates(req)

	return data
}

func (c *DPFMAPICaller) caloriesUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Calories {
	var req *[]dpfm_api_processing_formatter.CaloriesUpdates
	for _, v := range input.General.Calories {
		*req = append(*req, *dpfm_api_processing_formatter.ConvertToCaloriesUpdates(v))
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterCalories", "runtime_session_id": 123})
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
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToCaloriesUpdates(req)

	return data
}

func (c *DPFMAPICaller) nutritionalInfoUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.NutritionalInfo {
	var req *[]dpfm_api_processing_formatter.NutritionalInfoUpdates
	for _, v := range input.General.NutritionalInfo {
		*req = append(*req, *dpfm_api_processing_formatter.ConvertToNutritionalInfoUpdates(v))
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterNutritionalInfo", "runtime_session_id": 123})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "NutritionalInfo Data cannot insert"
		return nil
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToNutritionalInfoUpdates(req)

	return data
}

func (c *DPFMAPICaller) productDescriptionUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductDescription {
	var req *[]dpfm_api_processing_formatter.ProductDescriptionUpdates
	for _, v := range input.General.ProductDescription {
		*req = append(*req, *dpfm_api_processing_formatter.ConvertToProductDescriptionUpdates(v))
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterProductDescription", "runtime_session_id": 123})
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
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToProductDescriptionUpdates(req)

	return data
}

func (c *DPFMAPICaller) productDescByBPUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ProductDescByBP {
	var req *[]dpfm_api_processing_formatter.ProductDescByBPUpdates
	for _, v := range input.General.ProductDescByBP {
		*req = append(*req, *dpfm_api_processing_formatter.ConvertToProductDescByBPUpdates(v))
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterProductDescByBP", "runtime_session_id": 123})
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
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToProductDescByBPUpdates(req)

	return data
}

func (c *DPFMAPICaller) taxUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Tax {
	var req *[]dpfm_api_processing_formatter.TaxUpdates
	for _, v := range input.General.Tax {
		*req = append(*req, *dpfm_api_processing_formatter.ConvertToTaxUpdates(v))
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterTax", "runtime_session_id": 123})
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
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToTaxUpdates(req)

	return data
}

func (c *DPFMAPICaller) storageLocationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StorageLocation {
	var req *[]dpfm_api_processing_formatter.StorageLocationUpdates
	for _, businessPartner := range input.General.BusinessPartner {
		for _, bPPlant := range businessPartner.BPPlant {
			for _, v := range bPPlant.StorageLocation {
				*req = append(*req, *dpfm_api_processing_formatter.ConvertToStorageLocationUpdates(v))
			}
		}
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterStorageLocation", "runtime_session_id": 123})
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
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToStorageLocationUpdates(req)

	return data
}

func (c *DPFMAPICaller) workSchedulingUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.WorkScheduling {
	var req *[]dpfm_api_processing_formatter.WorkSchedulingUpdates
	for _, businessPartner := range input.General.BusinessPartner {
		for _, bPPlant := range businessPartner.BPPlant {
			for _, v := range bPPlant.WorkScheduling {
				*req = append(*req, *dpfm_api_processing_formatter.ConvertToWorkSchedulingUpdates(v))
			}
		}
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterWorkScheduling", "runtime_session_id": 123})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "WorkScheduling Data cannot insert"
		return nil
	}
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToWorkSchedulingUpdates(req)

	return data
}

func (c *DPFMAPICaller) accountingUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Accounting {
	var req *[]dpfm_api_processing_formatter.AccountingUpdates
	for _, businessPartner := range input.General.BusinessPartner {
		for _, bPPlant := range businessPartner.BPPlant {
			for _, v := range bPPlant.Accounting {
				*req = append(*req, *dpfm_api_processing_formatter.ConvertToAccountingUpdates(v))
			}
		}
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterAccounting", "runtime_session_id": 123})
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
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToAccountingUpdates(req)

	return data
}

func (c *DPFMAPICaller) mRPAreaUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.MRPArea {
	var req *[]dpfm_api_processing_formatter.MRPAreaUpdates
	for _, businessPartner := range input.General.BusinessPartner {
		for _, bPPlant := range businessPartner.BPPlant {
			for _, v := range bPPlant.MRPArea {
				*req = append(*req, *dpfm_api_processing_formatter.ConvertToMRPAreaUpdates(v))
			}
		}
	}

	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "ProductMasterMRPArea", "runtime_session_id": 123})
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
	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data := dpfm_api_output_formatter.ConvertToMRPAreaUpdates(req)

	return data
}
