package sub_func_complementer

type SDC struct {
	ConnectionKey       string   `json:"connection_key"`
	Result              bool     `json:"result"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartnerID   *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	APIType             string   `json:"api_type"`
	Message             Message  `json:"message"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	Deleted             bool     `json:"deleted"`
	SQLUpdateResult     *bool    `json:"sql_update_result"`
	SQLUpdateError      string   `json:"sql_update_error"`
	SubfuncResult       *bool    `json:"subfunc_result"`
	SubfuncError        string   `json:"subfunc_error"`
	ExconfResult        *bool    `json:"exconf_result"`
	ExconfError         string   `json:"exconf_error"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type Message struct {
	Header *Header `json:"Header"`
}

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
