package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header               *[]Header               `json:"Header"`
	Item                 *[]Item                 `json:"Item"`
	ItemAvailableFreight *[]ItemAvailableFreight `json:"ItemAvailableFreight"`
	Address              *[]Address              `json:"Address"`
	Partner              *[]Partner              `json:"Partner"`
}

type Header struct {
	Project               int     `json:"Project"`
	ProjectCode           string  `json:"ProjectCode"`
	ProjectDescription    string  `json:"ProjectDescription"`
	OwnerBusinessPartner  int     `json:"OwnerBusinessPartner"`
	OwnerPlant            string  `json:"OwnerPlant"`
	ProjectProfile        *string `json:"ProjectProfile"`
	ResponsiblePerson     *int    `json:"ResponsiblePerson"`
	ResponsiblePersonName *string `json:"ResponsiblePersonName"`
	PlannedStartDate      *string `json:"PlannedStartDate"`
	PlannedEndDate        *string `json:"PlannedEndDate"`
	ActualStartDate       *string `json:"ActualStartDate"`
	ActualEndDate         *string `json:"ActualEndDate"`
	CreationDate          string  `json:"CreationDate"`
	LastChangeDate        string  `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}

type Item struct {
	FreightAgreement                        int      `json:"FreightAgreement"`
	FreightAgreementItem                    int      `json:"FreightAgreementItem"`
	FreightAgreementItemCategory            string   `json:"FreightAgreementItemCategory"`
	SupplyChainRelationshipID               int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID       *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID  *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	FreightAgreementItemText                string   `json:"FreightAgreementItemText"`
	Product                                 string   `json:"Product"`
	ProductStandardID                       string   `json:"ProductStandardID"`
	ProductGroup                            *string  `json:"ProductGroup"`
	BaseUnit                                string   `json:"BaseUnit"`
	DeliverToParty                          *int     `json:"DeliverToParty"`
	DeliverFromParty                        *int     `json:"DeliverFromParty"`
	CreationDate                            string   `json:"CreationDate"`
	LastChangeDate                          string   `json:"LastChangeDate"`
	DeliverToPlant                          *string  `json:"DeliverToPlant"`
	DeliverToPlantTimeZone                  *string  `json:"DeliverToPlantTimeZone"`
	DeliverFromPlant                        *string  `json:"DeliverFromPlant"`
	DeliverFromPlantTimeZone                *string  `json:"DeliverFromPlantTimeZone"`
	Incoterms                               *string  `json:"Incoterms"`
	TransactionTaxClassification            string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                string   `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                  string   `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup           string   `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                            string   `json:"PaymentTerms"`
	DueCalculationBaseDate                  *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                          *string  `json:"PaymentDueDate"`
	NetPaymentDays                          *int     `json:"NetPaymentDays"`
	PaymentMethod                           string   `json:"PaymentMethod"`
	Project                                 *int     `json:"Project"`
	WBSElement                              *int     `json:"WBSElement"`
	ItemBillingStatus                       *string  `json:"ItemBillingStatus"`
	TaxCode                                 *string  `json:"TaxCode"`
	TaxRate                                 *float32 `json:"TaxRate"`
	ItemBlockStatus                         bool     `json:"ItemBlockStatus"`
	ItemBillingBlockStatus                  bool     `json:"ItemBillingBlockStatus"`
	IsCancelled                             bool     `json:"IsCancelled"`
	IsMarkedForDeletion                     bool     `json:"IsMarkedForDeletion"`
}

type ItemAvailableFreight struct {
	FreightAgreement                     int    `json:"FreightAgreement"`
	FreightAgreementItem                 int    `json:"FreightAgreementItem"`
	FreightAgreementItemAvailableFreight int    `json:"FreightAgreementItemAvailableFreight"`
	AvailableFreightType                 string `json:"AvailableFreightType"`
	AvailableFreightSpec                 string `json:"AvailableFreightSpec"`
	AvailableFreightCalendar             string `json:"AvailableFreightCalendar"`
	CreationDate                         string `json:"CreationDate"`
	LastChangeDate                       string `json:"LastChangeDate"`
	IsCancelled                          *bool  `json:"IsCancelled"`
	IsMarkedForDeletion                  *bool  `json:"IsMarkedForDeletion"`
}

type Address struct {
	FreightAgreement int     `json:"FreightAgreement"`
	AddressID        int     `json:"AddressID"`
	PostalCode       *string `json:"PostalCode"`
	LocalRegion      *string `json:"LocalRegion"`
	Country          *string `json:"Country"`
	District         *string `json:"District"`
	StreetName       *string `json:"StreetName"`
	CityName         *string `json:"CityName"`
	Building         *string `json:"Building"`
	Floor            *int    `json:"Floor"`
	Room             *int    `json:"Room"`
}

type Partner struct {
	FreightAgreement        int     `json:"FreightAgreement"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}
