package entity

type Course struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Days          int    `json:"days"`
	Price         int    `json:"price"`
	Status        int    `json:"status"`
	Beginner      int    `json:"beginner"`
	Intermediate  int    `json:"intermediate"`
	Order_in_list int    `json:"order_in_list"`
}
