package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-freight-agreement-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		InspectionLot:           data.InspectionLot,
		InspectionLotHeaderText: data.InspectionLotHeaderText,
	}
}

func ConvertToItemUpdates(itemUpdates dpfm_api_input_reader.Item) *ItemUpdates {
	data := itemUpdates

	return &ItemUpdates{
		InspectionLot:               data.InspectionLot,
		InspectionLotInspectionText: data.InspectionLotInspectionText,
	}
}

func ConvertToItemAvailableFreightUpdates(itemAvailableFreightUpdates dpfm_api_input_reader.ItemAvailableFreight) *ItemAvailableFreightUpdates {
	data := itemAvailableFreightUpdates

	return &ItemAvailableFreightUpdates{
		InspectionLot:               data.InspectionLot,
		InspectionLotInspectionText: data.InspectionLotInspectionText,
	}
}

func ConvertToAddressUpdates(addressUpdates dpfm_api_input_reader.Address) *AddressUpdates {
	data := addressUpdates

	return &AddressUpdates{
		InspectionLot:               data.InspectionLot,
		InspectionLotInspectionText: data.InspectionLotInspectionText,
	}
}

func ConvertToPartnerUpdates(partnerUpdates dpfm_api_input_reader.Partner) *PartnerUpdates {
	data := partnerUpdates

	return &PartnerUpdates{
		InspectionLot:               data.InspectionLot,
		InspectionLotInspectionText: data.InspectionLotInspectionText,
	}
}
