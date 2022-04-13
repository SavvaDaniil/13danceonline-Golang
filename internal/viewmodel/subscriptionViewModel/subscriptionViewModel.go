package subscriptionViewModel

type SubscriptionLiteViewModel struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Days  int    `json:"days"`
}
