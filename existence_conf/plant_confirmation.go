package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) bpPlantPlantExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	bpPlants := input.General.BPPlant
	for _, bpPlant := range bpPlants {
		plant, bpID := getBPPlantPlantExistenceConfKey(mapper, &bpPlant, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(plant) || isZero(bpID) {
				wg2.Done()
				return
			}
			res, err := c.plantExistenceConfRequest(plant, bpID, mapper, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) plantExistenceConfRequest(plant string, bpID int, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"BusinessPartner": bpID,
		"Plant":           plant,
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
	req.PlantGeneralReturn.Plant = plant
	req.PlantGeneralReturn.BusinessPartner = bpID

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getBPPlantPlantExistenceConfKey(mapper ExConfMapper, bpPlant *dpfm_api_input_reader.BPPlant, exconfErrMsg *string) (string, int) {
	var plant string
	var bpID int

	switch mapper.Field {
	case "Plant":
		plant = bpPlant.Plant
		bpID = bpPlant.BusinessPartner
	}

	return plant, bpID
}
