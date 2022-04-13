package paymentDTO

type PaymentInitDTO struct {
	Id_of_tutorial     int `json:"id_of_tutorial"`
	Id_of_course       int `json:"id_of_course"`
	Id_of_subscription int `json:"id_of_subscription"`
}

type PaymentResultRobokassaDTO struct {
	Id_of_payment int `json:"id_of_payment"`
}
