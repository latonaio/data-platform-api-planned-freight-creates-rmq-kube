package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-freight-agreement-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-freight-agreement-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-freight-agreement-creates-rmq-kube/sub_func_complementer"
	dpfm_api_input_reader "data-platform-api-project-creates-rmq-kube/DPFM_API_Input_Reader"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Item, error) {
	item := make([]Item, 0)

	for _, data := range *subfuncSDC.Message.Item {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		item = append(item, *item)
	}

	return &item, nil
}

func ConvertToItemAvailableFreightCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemAvailableFreight, error) {
	itemAvailableFreight := make([]ItemAvailableFreight, 0)

	for _, data := range *subfuncSDC.Message.ItemAvailableFreight {
		itemAvailableFreight, err := TypeConverter[*ItemAvailableFreight](data)
		if err != nil {
			return nil, err
		}

		itemAvailableFreight = append(itemAvailableFreight, *itemAvailableFreight)
	}

	return &itemAvailableFreight, nil
}

func ConvertToAddressCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Address, error) {
	address := make([]Address, 0)

	for _, data := range *subfuncSDC.Message.Address {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		address = append(address, *address)
	}

	return &address, nil
}

func ConvertToPartnerCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Partner, error) {
	partner := make([]Partner, 0)

	for _, data := range *subfuncSDC.Message.Partner {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partner = append(partner, *partner)
	}

	return &partner, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) (*[]Item, error) {
	item := make([]Item, 0)

	for _, data := range *&itemUpdates {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		item = append(item, *item)
	}

	return &item, nil
}

func ConvertToItemAvailableFreightUpdates(itemAvailableFreightUpdates *[]dpfm_api_processing_formatter.ItemAvailableFreightUpdates) (*[]ItemAvailableFreight, error) {
	itemAvailableFreight := make([]ItemAvailableFreight, 0)

	for _, data := range *&itemAvailableFreightUpdates {
		itemAvailableFreight, err := TypeConverter[*ItemAvailableFreight](data)
		if err != nil {
			return nil, err
		}

		itemAvailableFreight = append(itemAvailableFreight, *itemAvailableFreight)
	}

	return &itemAvailableFreight, nil
}
func ConvertToAddressUpdates(addressUpdates *[]dpfm_api_processing_formatter.AddressUpdates) (*[]Address, error) {
	address := make([]Address, 0)

	for _, data := range *&addressUpdates {
		address, err := TypeConverter[*Address](data)
		if err != nil {
			return nil, err
		}

		address = append(address, *address)
	}

	return &address, nil
}
func ConvertToPartnerUpdates(partnerUpdates *[]dpfm_api_processing_formatter.PartnerUpdates) (*[]Partner, error) {
	partner := make([]Partner, 0)

	for _, data := range *&partnerUpdates {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partner = append(partner, *partner)
	}

	return &partner, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
