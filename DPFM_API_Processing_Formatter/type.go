package dpfm_api_processing_formatter

type HeaderUpdates struct {
	PlannedFreight                  int      `json:"PlannedFreight"`
	FreightType                     string   `json:"FreightType"`
	FreightSpec                     string   `json:"FreightSpec"`
	FreightCalendar                 string   `json:"FreightCalendar"`
	PlannedFreightDepartureDate     string   `json:"PlannedFreightDepartureDate"`
	PlannedFreightDepartureTime     string   `json:"PlannedFreightDepartureTime"`
	PlannedFreightArrivalDate       string   `json:"PlannedFreightArrivalDate"`
	PlannedFreightArrivalTime       string   `json:"PlannedFreightArrivalTime"`
	FreightPartner                  *int     `json:"FreightPartner"`
	Buyer                           *int     `json:"Buyer"`
	Seller                          *int     `json:"Seller"`
	PlannedFreightNumberInCharacter *string  `json:"PlannedFreightNumberInCharacter"`
	PlannedFreightNumberDescription *string  `json:"PlannedFreightNumberDescription"`
	FreightCapacityWeight           *float32 `json:"FreightCapacityWeight"`
	FreightCapacityWeightUnit       *string  `json:"FreightCapacityWeightUnit"`
	PlannedFreightLongText          *string  `json:"PlannedFreightLongText"`
}
