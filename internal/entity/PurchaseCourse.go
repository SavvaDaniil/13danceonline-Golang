package entity

type PurchaseCourse struct {
	Id                      int
	Id_of_user              int
	Id_of_payment           *int
	Id_of_course            int
	Days                    int
	Active                  int
	Date_of_add             *string
	Date_of_activation      *string
	Date_of_must_be_used_to *string
}
