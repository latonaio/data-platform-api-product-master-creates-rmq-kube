package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) generalQuantityUnitExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	generals := make([]dpfm_api_input_reader.General, 0, 1)
	generals = append(generals, input.General)
	for _, general := range generals {
		quantityUnit := getGeneralQuantityUnitExistenceConfKey(mapper, &general, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(quantityUnit) {
				wg2.Done()
				return
			}
			res, err := c.quantityUnitExistenceConfRequest(quantityUnit, mapper, input, existenceMap, mtx, log)
			if err != nil {
				mtx.Lock()
				*errs = append(*errs, err)
				mtx.Unlock()
			}
			if res != "" {
				*exconfErrMsg = res
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) bpPlantQuantityUnitExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	bpPlants := input.General.BPPlant
	for _, bpPlant := range bpPlants {
		quantityUnit := getBPPlantQuantityUnitExistenceConfKey(mapper, &bpPlant, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(quantityUnit) {
				wg2.Done()
				return
			}
			res, err := c.quantityUnitExistenceConfRequest(quantityUnit, mapper, input, existenceMap, mtx, log)
			if err != nil {
				mtx.Lock()
				*errs = append(*errs, err)
				mtx.Unlock()
			}
			if res != "" {
				*exconfErrMsg = res
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) caloriesQuantityUnitExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	businessPartners := input.General.BusinessPartner
	for _, businessPartner := range businessPartners {
		calories := businessPartner.Calories
		for _, calory := range calories {
			quantityUnit := getCaloriesQuantityUnitExistenceConfKey(mapper, &calory, exconfErrMsg)
			wg2.Add(1)
			exReqTimes++
			go func() {
				if isZero(quantityUnit) {
					wg2.Done()
					return
				}
				res, err := c.quantityUnitExistenceConfRequest(quantityUnit, mapper, input, existenceMap, mtx, log)
				if err != nil {
					mtx.Lock()
					*errs = append(*errs, err)
					mtx.Unlock()
				}
				if res != "" {
					*exconfErrMsg = res
				}
				wg2.Done()
			}()
		}
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) nutritionalInfoQuantityUnitExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	businessPartners := input.General.BusinessPartner
	for _, businessPartner := range businessPartners {
		nutritionalInfos := businessPartner.NutritionalInfo
		for _, nutritionalInfo := range nutritionalInfos {
			quantityUnit := getNutritionalInfoQuantityUnitExistenceConfKey(mapper, &nutritionalInfo, exconfErrMsg)
			wg2.Add(1)
			exReqTimes++
			go func() {
				if isZero(quantityUnit) {
					wg2.Done()
					return
				}
				res, err := c.quantityUnitExistenceConfRequest(quantityUnit, mapper, input, existenceMap, mtx, log)
				if err != nil {
					mtx.Lock()
					*errs = append(*errs, err)
					mtx.Unlock()
				}
				if res != "" {
					*exconfErrMsg = res
				}
				wg2.Done()
			}()
		}
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) workSchedulingQuantityUnitExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	bpPlants := input.General.BPPlant
	for _, bpPlant := range bpPlants {
		workSchedulings := bpPlant.WorkScheduling
		for _, workScheduling := range workSchedulings {
			quantityUnit := getWorkSchedulingQuantityUnitExistenceConfKey(mapper, &workScheduling, exconfErrMsg)
			wg2.Add(1)
			exReqTimes++
			go func() {
				if isZero(quantityUnit) {
					wg2.Done()
					return
				}
				res, err := c.quantityUnitExistenceConfRequest(quantityUnit, mapper, input, existenceMap, mtx, log)
				if err != nil {
					mtx.Lock()
					*errs = append(*errs, err)
					mtx.Unlock()
				}
				if res != "" {
					*exconfErrMsg = res
				}
				wg2.Done()
			}()
		}
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) quantityUnitExistenceConfRequest(quantityUnit string, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"QuantityUnit": quantityUnit,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.QuantityUnitReturn.QuantityUnit = quantityUnit

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getGeneralQuantityUnitExistenceConfKey(mapper ExConfMapper, general *dpfm_api_input_reader.General, exconfErrMsg *string) string {
	var quantityUnit string

	switch mapper.Field {
	case "BaseUnit":
		if general.BaseUnit == nil {
			quantityUnit = ""
		} else {
			quantityUnit = *general.BaseUnit
		}
	case "WeightUnit":
		if general.WeightUnit == nil {
			quantityUnit = ""
		} else {
			quantityUnit = *general.WeightUnit
		}
	case "InternalCapacityQuantityUnit":
		if general.InternalCapacityQuantityUnit == nil {
			quantityUnit = ""
		} else {
			quantityUnit = *general.InternalCapacityQuantityUnit
		}
	}

	return quantityUnit
}

func getBPPlantQuantityUnitExistenceConfKey(mapper ExConfMapper, general *dpfm_api_input_reader.BPPlant, exconfErrMsg *string) string {
	var quantityUnit string

	switch mapper.Field {
	case "InventoryUnit":
		if general.InventoryUnit == nil {
			quantityUnit = ""
		} else {
			quantityUnit = *general.InventoryUnit
		}
	}

	return quantityUnit
}

func getCaloriesQuantityUnitExistenceConfKey(mapper ExConfMapper, calories *dpfm_api_input_reader.Calories, exconfErrMsg *string) string {
	var quantityUnit string

	switch mapper.Field {
	case "CaloryUnit":
		if calories.CaloryUnit == nil {
			quantityUnit = ""
		} else {
			quantityUnit = *calories.CaloryUnit
		}
	}

	return quantityUnit
}

func getNutritionalInfoQuantityUnitExistenceConfKey(mapper ExConfMapper, nutritionalInfo *dpfm_api_input_reader.NutritionalInfo, exconfErrMsg *string) string {
	var quantityUnit string

	switch mapper.Field {
	case "NutrientContentUnit":
		if nutritionalInfo.NutrientContentUnit == nil {
			quantityUnit = ""
		} else {
			quantityUnit = *nutritionalInfo.NutrientContentUnit
		}
	}

	return quantityUnit
}

func getWorkSchedulingQuantityUnitExistenceConfKey(mapper ExConfMapper, workScheduling *dpfm_api_input_reader.WorkScheduling, exconfErrMsg *string) string {
	var quantityUnit string

	switch mapper.Field {
	case "ProductProductionQuantityUnit":
		if workScheduling.ProductProductionQuantityUnit == nil {
			quantityUnit = ""
		} else {
			quantityUnit = *workScheduling.ProductProductionQuantityUnit
		}
	}

	return quantityUnit
}
