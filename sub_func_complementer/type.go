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
	Header               *Header                 `json:"Header"`
	Item                 *[]Item                 `json:"Item"`
	ItemAvailableFreight *[]ItemAvailableFreight `json:"ItemAvailableFreight"`
	Address              *[]Address              `json:"Address"`
	Partner              *[]Partner              `json:"Partner"`
}

type Header struct {
	FreightAgreement                        int      `json:"FreightAgreement"`
	FreightAgreementDate                    string   `json:"FreightAgreementDate"`
	FreightAgreementType                    string   `json:"FreightAgreementType"`
	SupplyChainRelationshipID               int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipFreightID        int      `json:"SupplyChainRelationshipFreightID"`
	SupplyChainRelationshipFreightBillingID *int     `json:"SupplyChainRelationshipFreightBillingID"`
	SupplyChainRelationshipFreightPaymentID *int     `json:"SupplyChainRelationshipFreightPaymentID"`
	Buyer                                   int      `json:"Buyer"`
	Seller                                  int      `json:"Seller"`
	FreightPartner                          int      `json:"FreightPartner"`
	FreightBillToParty                      *int     `json:"FreightBillToParty"`
	FreightBillFromParty                    *int     `json:"FreightBillFromParty"`
	FreightBillToCountry                    *string  `json:"FreightBillToCountry"`
	FreightBillFromCountry                  *string  `json:"FreightBillFromCountry"`
	FreightPayer                            *int     `json:"FreightPayer"`
	FreightPayee                            *int     `json:"FreightPayee"`
	CreationDate                            string   `json:"CreationDate"`
	LastChangeDate                          string   `json:"LastChangeDate"`
	ContractType                            *string  `json:"ContractType"`
	FreightAgreementValidityStartDate       *string  `json:"FreightAgreementValidityStartDate"`
	FreightAgreementValidityEndDate         *string  `json:"FreightAgreementValidityEndDate"`
	InvoicePeriodStartDate                  *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                    *string  `json:"InvoicePeriodEndDate"`
	HeaderBillingStatus                     string   `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus                string   `json:"HeaderDocReferenceStatus"`
	TransactionCurrency                     string   `json:"TransactionCurrency"`
	PricingDate                             string   `json:"PricingDate"`
	PriceDetnExchangeRate                   *float32 `json:"PriceDetnExchangeRate"`
	Incoterms                               *string  `json:"Incoterms"`
	PaymentTerms                            string   `json:"PaymentTerms"`
	PaymentMethod                           string   `json:"PaymentMethod"`
	ReferenceDocument                       *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                   *int     `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup                  string   `json:"AccountAssignmentGroup"`
	AccountingExchangeRate                  *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate                     string   `json:"InvoiceDocumentDate"`
	IsExportImport                          *bool    `json:"IsExportImport"`
	HeaderText                              *string  `json:"HeaderText"`
	HeaderBlockStatus                       *bool    `json:"HeaderBlockStatus"`
	HeaderBillingBlockStatus                *bool    `json:"HeaderBillingBlockStatus"`
	IsCancelled                             *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                     *bool    `json:"IsMarkedForDeletion"`
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

type ItemComponentStockConfirmation struct {
	ProductionOrder           int     `json:"ProductionOrder"`
	ProductionOrderItem       int     `json:"ProductionOrderItem"`
	Operations                int     `json:"Operations"`
	OperationsItem            int     `json:"OperationsItem"`
	BillOfMaterial            int     `json:"BillOfMaterial"`
	BillOfMaterialItem        int     `json:"BillOfMaterialItem"`
	InventoryStockType        *string `json:"InventoryStockType"`
	InventorySpecialStockType *string `json:"InventorySpecialStockType"`
	AvailableProductStock     float32 `json:"AvailableProductStock"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
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
