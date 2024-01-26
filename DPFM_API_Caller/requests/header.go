package requests

type Header struct {
	PlannedFreight                         int      `json:"PlannedFreight"`
	PlannedFreightType                     *string  `json:"PlannedFreightType"`
	FreightAgreement                       int      `json:"FreightAgreement"`
	FreightAgreementItem                   int      `json:"FreightAgreementItem"`
	FreightAgreementItemAvailableFreight   int      `json:"FreightAgreementItemAvailableFreight"`
	FreightType                            string   `json:"FreightType"`
	FreightSpec                            string   `json:"FreightSpec"`
	FreightCalendar                        string   `json:"FreightCalendar"`
	PlannedFreightDepartureDate            string   `json:"PlannedFreightDepartureDate"`
	PlannedFreightDepartureTime            string   `json:"PlannedFreightDepartureTime"`
	PlannedFreightArrivalDate              string   `json:"PlannedFreightArrivalDate"`
	PlannedFreightArrivalTime              string   `json:"PlannedFreightArrivalTime"`
	SupplyChainRelationshipID              int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipFreightID       int      `json:"SupplyChainRelationshipFreightID"`
	FreightPartner                         *int     `json:"FreightPartner"`
	Buyer                                  *int     `json:"Buyer"`
	Seller                                 *int     `json:"Seller"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	PlannedFreightNumberInCharacter        *string  `json:"PlannedFreightNumberInCharacter"`
	PlannedFreightNumberDescription        *string  `json:"PlannedFreightNumberDescription"`
	FRPArea                                *string  `json:"FRPArea"`
	FRPController                          *string  `json:"FRPController"`
	FreightCapacityWeight                  *float32 `json:"FreightCapacityWeight"`
	FreightCapacityWeightUnit              *string  `json:"FreightCapacityWeightUnit"`
	PlannedFreightLongText                 *string  `json:"PlannedFreightLongText"`
	CreationDate                           string   `json:"CreationDate"`
	CreationTime                           string   `json:"CreationTime"`
	LastChangeDate                         string   `json:"LastChangeDate"`
	LastChangeTime                         string   `json:"LastChangeTime"`
	IsReleased                             *bool    `json:"IsReleased"`
	IsMarkedForDeletion                    *bool    `json:"IsMarkedForDeletion"`
}
