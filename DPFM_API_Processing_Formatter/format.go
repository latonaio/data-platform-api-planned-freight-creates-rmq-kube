package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-planned-freight-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		PlannedFreight:                  data.PlannedFreight,
		FreightType:                     data.FreightType,
		FreightSpec:                     data.FreightSpec,
		FreightCalendar:                 data.FreightCalendar,
		PlannedFreightDepartureDate:     data.PlannedFreightDepartureDate,
		PlannedFreightDepartureTime:     data.PlannedFreightDepartureTime,
		PlannedFreightArrivalDate:       data.PlannedFreightArrivalDate,
		PlannedFreightArrivalTime:       data.PlannedFreightArrivalTime,
		FreightPartner:                  data.FreightPartner,
		Buyer:                           data.Buyer,
		Seller:                          data.Seller,
		PlannedFreightNumberInCharacter: data.PlannedFreightNumberInCharacter,
		PlannedFreightNumberDescription: data.PlannedFreightNumberDescription,
		FreightCapacityWeight:           data.FreightCapacityWeight,
		FreightCapacityWeightUnit:       data.FreightCapacityWeightUnit,
		PlannedFreightLongText:          data.PlannedFreightLongText,
	}
}
