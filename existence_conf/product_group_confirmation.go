package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) generalProductGroupExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	generals := make([]dpfm_api_input_reader.General, 0, 1)
	generals = append(generals, input.General)
	for _, general := range generals {
		productGroup := getGeneralProductGroupExistenceConfKey(mapper, &general, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(productGroup) {
				wg2.Done()
				return
			}
			res, err := c.productGroupExistenceConfRequest(productGroup, mapper, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) productGroupExistenceConfRequest(productGroup string, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"ProductGroup": productGroup,
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
	req.ProductGroupReturn.ProductGroup = productGroup

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getGeneralProductGroupExistenceConfKey(mapper ExConfMapper, general *dpfm_api_input_reader.General, exconfErrMsg *string) string {
	var productGroup string

	switch mapper.Field {
	case "ProductGroup":
		if general.ProductGroup == nil {
			productGroup = ""
		} else {
			productGroup = *general.ProductGroup
		}
	}

	return productGroup
}
