package requests

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
