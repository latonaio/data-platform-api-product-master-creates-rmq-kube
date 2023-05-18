package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FileReader struct{}

func NewFileReader() *FileReader {
	return &FileReader{}
}

func (*FileReader) ReadECMC(path string) EC_MC {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("input data read error :%#v", err.Error())
		os.Exit(1)
	}
	ec := EC_MC{}
	err = json.Unmarshal(raw, &ec)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return ec
}

func (*FileReader) ReadSDC(path string) SDC {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("input data read error :%#v", err.Error())
		os.Exit(1)
	}
	sdc := SDC{}
	err = json.Unmarshal(raw, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}

func ConvertToSDC(raw []byte) SDC {
	sdc := SDC{}
	err := json.Unmarshal(raw, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	product := sdc.General.Product
	for i, businessPartner := range sdc.General.BusinessPartner {
		bp := businessPartner.BusinessPartner
		sdc.General.BusinessPartner[i].Product = product
		for j := range businessPartner.Allergen {
			sdc.General.BusinessPartner[i].Allergen[j].Product = product
			sdc.General.BusinessPartner[i].Allergen[j].BusinessPartner = bp
		}
		for j := range businessPartner.Calories {
			sdc.General.BusinessPartner[i].Calories[j].Product = product
			sdc.General.BusinessPartner[i].Calories[j].BusinessPartner = bp
		}
		for j := range businessPartner.NutritionalInfo {
			sdc.General.BusinessPartner[i].NutritionalInfo[j].Product = product
			sdc.General.BusinessPartner[i].NutritionalInfo[j].BusinessPartner = bp
		}
		for j, productDescription := range businessPartner.ProductDescription {
			sdc.General.BusinessPartner[i].ProductDescription[j].Product = product
			for k := range productDescription.ProductDescByBP {
				sdc.General.BusinessPartner[i].ProductDescription[j].ProductDescByBP[k].Product = product
				sdc.General.BusinessPartner[i].ProductDescription[j].ProductDescByBP[k].BusinessPartner = bp
			}
		}
	}
	for i, bpPlant := range sdc.General.BPPlant {
		bp := bpPlant.BusinessPartner
		plant := bpPlant.Plant
		sdc.General.BPPlant[i].Product = product
		for j := range bpPlant.Accounting {
			sdc.General.BPPlant[i].Accounting[j].Product = product
			sdc.General.BPPlant[i].Accounting[j].BusinessPartner = bp
			sdc.General.BPPlant[i].Accounting[j].Plant = plant
		}
		for j := range bpPlant.BPPlantDoc {
			sdc.General.BPPlant[i].BPPlantDoc[j].Product = product
			sdc.General.BPPlant[i].BPPlantDoc[j].BusinessPartner = bp
			sdc.General.BPPlant[i].BPPlantDoc[j].Plant = plant
		}
		for j := range bpPlant.MRPArea {
			sdc.General.BPPlant[i].MRPArea[j].Product = product
			sdc.General.BPPlant[i].MRPArea[j].BusinessPartner = bp
			sdc.General.BPPlant[i].MRPArea[j].Plant = plant
		}
		for j := range bpPlant.Quality {
			sdc.General.BPPlant[i].Quality[j].Product = product
			sdc.General.BPPlant[i].Quality[j].BusinessPartner = bp
			sdc.General.BPPlant[i].Quality[j].Plant = plant
		}
		for j, storageLocation := range bpPlant.StorageLocation {
			sl := storageLocation.StorageLocation
			sdc.General.BPPlant[i].StorageLocation[j].Product = product
			sdc.General.BPPlant[i].StorageLocation[j].BusinessPartner = bp
			sdc.General.BPPlant[i].StorageLocation[j].Plant = plant
			for k := range storageLocation.StorageBin {
				sdc.General.BPPlant[i].StorageLocation[j].StorageBin[k].Product = product
				sdc.General.BPPlant[i].StorageLocation[j].StorageBin[k].BusinessPartner = bp
				sdc.General.BPPlant[i].StorageLocation[j].StorageBin[k].Plant = plant
				sdc.General.BPPlant[i].StorageLocation[j].StorageBin[k].StorageLocation = sl
			}
		}
		for j := range bpPlant.WorkScheduling {
			sdc.General.BPPlant[i].WorkScheduling[j].Product = product
			sdc.General.BPPlant[i].WorkScheduling[j].BusinessPartner = bp
			sdc.General.BPPlant[i].WorkScheduling[j].Plant = plant
		}
	}
	for i := range sdc.General.GeneralDoc {
		sdc.General.GeneralDoc[i].Product = product
	}
	for i := range sdc.General.Tax {
		sdc.General.Tax[i].Product = product
	}

	return sdc
}
