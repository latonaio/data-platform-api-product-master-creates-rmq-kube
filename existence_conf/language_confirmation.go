package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) generalLanguageExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	generals := make([]dpfm_api_input_reader.General, 0, 1)
	generals = append(generals, input.General)
	for _, general := range generals {
		language := getGeneralLanguageExistenceConfKey(mapper, &general, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(language) {
				wg2.Done()
				return
			}
			res, err := c.languageExistenceConfRequest(language, mapper, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) productDescriptionLanguageExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	businessPartners := input.General.BusinessPartner
	for _, businessPartner := range businessPartners {
		productDescriptions := businessPartner.ProductDescription
		for _, productDescription := range productDescriptions {
			language := getProductDescriptionLanguageExistenceConfKey(mapper, &productDescription, exconfErrMsg)
			wg2.Add(1)
			exReqTimes++
			go func() {
				if isZero(language) {
					wg2.Done()
					return
				}
				res, err := c.languageExistenceConfRequest(language, mapper, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) languageExistenceConfRequest(language string, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"Language": language,
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
	req.LanguageReturn.Language = language

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getGeneralLanguageExistenceConfKey(mapper ExConfMapper, general *dpfm_api_input_reader.General, exconfErrMsg *string) string {
	var language string

	switch mapper.Field {
	case "CountryOfOriginLanguage":
		if general.CountryOfOriginLanguage == nil {
			language = ""
		} else {
			language = *general.CountryOfOriginLanguage
		}
	}

	return language
}

func getProductDescriptionLanguageExistenceConfKey(mapper ExConfMapper, productDescription *dpfm_api_input_reader.ProductDescription, exconfErrMsg *string) string {
	var language string

	switch mapper.Field {
	case "Language":
		language = productDescription.Language
	}

	return language
}
