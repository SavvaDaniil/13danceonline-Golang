package entity

type PurchaseSubscription struct {
	Id                      int
	Id_of_user              int
	Id_of_payment           int
	Id_of_subscription      int
	Days                    int
	Active                  int
	Date_of_add             *string
	Date_of_activation      *string
	Date_of_must_be_used_to *string
}
