package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-planned-freight-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		PlannedFreight: data.PlannedFreight,
	}
}
