package dpfm_api_processing_formatter

type HeaderUpdates struct {
	InspectionLot           int     `json:"InspectionLot"`
	InspectionLotHeaderText *string `json:"InspectionLotHeaderText"`
}

type ItemUpdates struct {
	InspectionLot               int     `json:"InspectionLot"`
	Inspection                  int     `json:"Inspection"`
	InspectionLotInspectionText *string `json:"InspectionLotInspectionText"`
}

type ItemAvailableFreightUpdates struct {
	InspectionLot               int     `json:"InspectionLot"`
	Inspection                  int     `json:"Inspection"`
	InspectionLotInspectionText *string `json:"InspectionLotInspectionText"`
}
type AddressUpdates struct {
	InspectionLot               int     `json:"InspectionLot"`
	Inspection                  int     `json:"Inspection"`
	InspectionLotInspectionText *string `json:"InspectionLotInspectionText"`
}
type PartnerUpdates struct {
	InspectionLot               int     `json:"InspectionLot"`
	Inspection                  int     `json:"Inspection"`
	InspectionLotInspectionText *string `json:"InspectionLotInspectionText"`
}
