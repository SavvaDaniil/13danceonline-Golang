package subscriptionFacade

import (
	//"danceonline/internal/entity"
	"danceonline/internal/repository/subscriptionRepository"
	"danceonline/internal/viewmodel/subscriptionViewModel"
)

func ListAllLiteActive() ([]subscriptionViewModel.SubscriptionLiteViewModel, error) {

	var subscriptionLiteViewModels []subscriptionViewModel.SubscriptionLiteViewModel
	subscriptions, errSearch := subscriptionRepository.ListAllActive()
	if errSearch != nil {
		return nil, errSearch
	}

	for _, subscription := range subscriptions {
		var subscriptionLiteViewModel subscriptionViewModel.SubscriptionLiteViewModel = subscriptionViewModel.SubscriptionLiteViewModel{
			Id:    subscription.Id,
			Name:  subscription.Name,
			Price: subscription.Price,
			Days:  subscription.Days,
		}
		subscriptionLiteViewModels = append(subscriptionLiteViewModels, subscriptionLiteViewModel)
	}

	return subscriptionLiteViewModels, nil
}
