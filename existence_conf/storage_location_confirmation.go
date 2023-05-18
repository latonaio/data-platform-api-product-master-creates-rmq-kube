package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-product-master-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) mrpAreaStorageLocationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	bpPlants := input.General.BPPlant
	for _, bpPlant := range bpPlants {
		mrpAreas := bpPlant.MRPArea
		for _, mrpArea := range mrpAreas {
			bpID, plant, storageLocation := getMRPAreaStorageLocationExistenceConfKey(mapper, &mrpArea, exconfErrMsg)
			wg2.Add(1)
			exReqTimes++
			go func() {
				if isZero(bpID) || isZero(plant) || isZero(storageLocation) {
					wg2.Done()
					return
				}
				res, err := c.storageLocationExistenceConfRequest(bpID, plant, storageLocation, mapper, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) storageLocationStorageLocationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	bpPlants := input.General.BPPlant
	for _, bpPlant := range bpPlants {
		storageLocations := bpPlant.StorageLocation
		for _, storageLocation := range storageLocations {
			bpID, plant, storageLocation := getStorageLocationStorageLocationExistenceConfKey(mapper, &storageLocation, exconfErrMsg)
			wg2.Add(1)
			exReqTimes++
			go func() {
				if isZero(bpID) || isZero(plant) || isZero(storageLocation) {
					wg2.Done()
					return
				}
				res, err := c.storageLocationExistenceConfRequest(bpID, plant, storageLocation, mapper, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) workSchedulingStorageLocationExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	bpPlants := input.General.BPPlant
	for _, bpPlant := range bpPlants {
		workSchedulings := bpPlant.WorkScheduling
		for _, workScheduling := range workSchedulings {
			bpID, plant, storageLocation := getWorkSchedulingStorageLocationExistenceConfKey(mapper, &workScheduling, exconfErrMsg)
			wg2.Add(1)
			exReqTimes++
			go func() {
				if isZero(bpID) || isZero(plant) || isZero(storageLocation) {
					wg2.Done()
					return
				}
				res, err := c.storageLocationExistenceConfRequest(bpID, plant, storageLocation, mapper, input, existenceMap, mtx, log)
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

func (c *ExistenceConf) storageLocationExistenceConfRequest(bpID int, plant string, storageLocation string, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"BusinessPartner": bpID,
		"Plant":           plant,
		"StorageLocation": storageLocation,
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
	req.StorageLocationReturn.BusinessPartner = bpID
	req.StorageLocationReturn.Plant = plant
	req.StorageLocationReturn.StorageLocation = storageLocation

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getMRPAreaStorageLocationExistenceConfKey(mapper ExConfMapper, mrpArea *dpfm_api_input_reader.MRPArea, exconfErrMsg *string) (int, string, string) {
	var bpID int
	var plant string
	var sl string

	switch mapper.Field {
	case "StorageLocation":
		bpID = mrpArea.BusinessPartner
		plant = mrpArea.Plant
		if mrpArea.StorageLocationForMRP == nil {
			sl = ""
		} else {
			sl = *mrpArea.StorageLocationForMRP
		}
	}

	return bpID, plant, sl
}

func getStorageLocationStorageLocationExistenceConfKey(mapper ExConfMapper, storageLocation *dpfm_api_input_reader.StorageLocation, exconfErrMsg *string) (int, string, string) {
	var bpID int
	var plant string
	var sl string

	switch mapper.Field {
	case "StorageLocation":
		bpID = storageLocation.BusinessPartner
		plant = storageLocation.Plant
		sl = storageLocation.StorageLocation

	}

	return bpID, plant, sl
}

func getWorkSchedulingStorageLocationExistenceConfKey(mapper ExConfMapper, workScheduling *dpfm_api_input_reader.WorkScheduling, exconfErrMsg *string) (int, string, string) {
	var bpID int
	var plant string
	var sl string

	switch mapper.Field {
	case "StorageLocation":
		bpID = workScheduling.BusinessPartner
		plant = workScheduling.Plant
		if workScheduling.ProductionInvtryManagedLoc == nil {
			sl = ""
		} else {
			sl = *workScheduling.ProductionInvtryManagedLoc
		}
	}

	return bpID, plant, sl
}
