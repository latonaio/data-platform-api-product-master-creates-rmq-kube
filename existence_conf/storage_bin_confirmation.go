package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) storageBinStorageBinExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	bpPlants := input.General.BPPlant
	for _, bpPlant := range bpPlants {
		storageLocations := bpPlant.StorageLocation
		for _, storageLocation := range storageLocations {
			storageBins := storageLocation.StorageBin
			for _, storageBin := range storageBins {
				bpID, plant, storageLocation, storageBin := getStorageBinStorageBinExistenceConfKey(mapper, &storageBin, exconfErrMsg)
				wg2.Add(1)
				exReqTimes++
				go func() {
					if isZero(bpID) || isZero(plant) || isZero(storageLocation) || isZero(storageBin) {
						wg2.Done()
						return
					}
					res, err := c.storageBinExistenceConfRequest(bpID, plant, storageLocation, storageBin, mapper, input, existenceMap, mtx, log)
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
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) storageBinExistenceConfRequest(bpID int, plant string, storageLocation string, storageBin string, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"BusinessPartner": bpID,
		"Plant":           plant,
		"StorageLocation": storageLocation,
		"StorageBin":      storageBin,
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
	req.StorageBinReturn.BusinessPartner = bpID
	req.StorageBinReturn.Plant = plant
	req.StorageBinReturn.StorageLocation = storageLocation
	req.StorageBinReturn.StorageBin = storageBin

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getStorageBinStorageBinExistenceConfKey(mapper ExConfMapper, storageBin *dpfm_api_input_reader.StorageBin, exconfErrMsg *string) (int, string, string, string) {
	var bpID int
	var plant string
	var sl string
	var sb string

	switch mapper.Field {
	case "StorageBin":
		bpID = storageBin.BusinessPartner
		plant = storageBin.Plant
		sl = storageBin.StorageLocation
		sb = storageBin.StorageBin

	}

	return bpID, plant, sl, sb
}
