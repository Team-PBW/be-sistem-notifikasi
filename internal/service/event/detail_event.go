package event

import (
	// "time"
	// "log"

	"log"

	"golang.org/x/e-calender/internal/dto"
)

func (e *EventService) DetailEvent(id string) (*dto.EventDto, error) {

	log.Println("tesssssssssss service")


	// err := e.Validate.TryValidate(id)
	// if err != nil {
	// 	return nil, err
	// }

	log.Println("validasi detail event sukses")

	evt, err := e.EventRepository.FindEventByID(id)
	if err != nil {
		log.Println("error fetch detail dari service")
		return nil, err
	}

	// guests, err := e.EventRepository.FindGuestsInEvent(id)
	// if err != nil {
	// 	return nil, err
	// }

	// var data []string
	// for _, guest := range guests {
	// 	data = append(data, guest.Username)
	// }

	const (
		dateLayout = "2006-01-02"
		timeLayout = "15:04:05"
	)

	eventObject := &dto.EventDto{
		Id: id,
		Title:       evt.Title,
		IdCategory:  evt.CategoryId,
		Location:    evt.Location,
		Description: evt.Description,
		// Distance:    evt.Distance,
		// TimeDistance: evt.TimeDistance,
		Date:        evt.Date.Format(dateLayout),
		StartTime:   evt.StartTime.Format(timeLayout),
		EndTime:     evt.EndTime.Format(timeLayout),
		// InvitedUser: data,
	}

	return eventObject, nil
}

// func (e *EventService) CheckEventByDateDay(category map[string][]string) (*dto.EventDto, error) {
// 	for key, datas := range category {
// 		if data, ok := category[key] {
// 			if !ok {
// 				return nil, err
// 			}

			
	
// 		}
// 	}

// 	log.Println(category["end"][0])

// 	return nil
// }