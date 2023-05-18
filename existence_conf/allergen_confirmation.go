package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) allergenAllergenExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	businessPartners := input.General.BusinessPartner
	for _, businessPartner := range businessPartners {
		allergens := businessPartner.Allergen
		for _, allergen := range allergens {
			allergen := getAllergenAllergenExistenceConfKey(mapper, &allergen, exconfErrMsg)
			wg2.Add(1)
			exReqTimes++
			go func() {
				if isZero(allergen) {
					wg2.Done()
					return
				}
				res, err := c.allergenExistenceConfRequest(allergen, mapper, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) allergenExistenceConfRequest(allergen string, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"Allergen": allergen,
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
	req.AllergenReturn.Allergen = allergen

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getAllergenAllergenExistenceConfKey(mapper ExConfMapper, allergen *dpfm_api_input_reader.Allergen, exconfErrMsg *string) string {
	var ag string

	switch mapper.Field {
	case "Allergen":
		ag = allergen.Allergen

	}

	return ag
}
