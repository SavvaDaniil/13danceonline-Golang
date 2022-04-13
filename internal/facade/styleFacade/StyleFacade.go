package styleFacade

import (
	"danceonline/internal/entity"
	"danceonline/internal/repository/styleRepository"
	"danceonline/internal/viewmodel"
	"danceonline/internal/viewmodel/styleViewModel"
)

func ListAllActiveLite() []styleViewModel.StyleLiteViewModel {
	var styles []entity.Style = styleRepository.ListAllActive()

	var styleLiteViewModels []styleViewModel.StyleLiteViewModel
	for index, style := range styles {
		styleLiteViewModel := styleViewModel.StyleLiteViewModel{
			Id:          style.Id,
			Index:       (index + 1),
			IsIndexDiv5: IsIndexDiv5(index + 1),
			Name:        style.Name,
			Description: style.Description,
		}
		styleLiteViewModels = append(styleLiteViewModels, styleLiteViewModel)
	}

	return styleLiteViewModels
}

func JsonListAllActiveMicro() viewmodel.JsonAnswerStatus {
	return viewmodel.JsonAnswerStatus{
		Status:               "success",
		StyleMicroViewModels: ListAllActiveMicro(),
	}
}

func ListAllActiveMicro() []styleViewModel.StyleMicroViewModel {
	var styles []entity.Style = styleRepository.ListAllActive()

	var styleMicroViewModels []styleViewModel.StyleMicroViewModel
	for _, style := range styles {
		styleMicroViewModel := styleViewModel.StyleMicroViewModel{
			Id:   style.Id,
			Name: style.Name,
		}
		styleMicroViewModels = append(styleMicroViewModels, styleMicroViewModel)
	}

	return styleMicroViewModels
}

func IsIndexInt(index int) bool {
	return index%2 == 0
}
func IsIndexDiv5(index int) bool {
	return index%5 == 0
}
