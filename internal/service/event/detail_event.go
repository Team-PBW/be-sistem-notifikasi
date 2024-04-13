package event

import (
	"strconv"

	"golang.org/x/e-calender/internal/dto"
)

func (e *EventService) DetailEvent(id string) (*dto.EventDTO, error) {

	err := e.Validate.TryValidate(id)
	if  err != nil {
		return nil, err
	}

	evt, err := e.EventRepository.FindEventByID(id)
	if err != nil {
		return nil, err
	}

	guests, err := e.EventRepository.FindGuestsInEvent(id)
	if err != nil {
		return nil, err
	}

	var personConfirmed []*dto.UserDto
	for _, guest := range guests {
		strPhone := strconv.Itoa(guest.PhoneNumber)
		userDto := &dto.UserDto{
			Username: guest.Username,
			PhoneNumber: strPhone,
		}
		personConfirmed = append(personConfirmed, userDto)
	}

	eventObject := &dto.EventDTO{
		EventName:       evt.EventName,
		FromDate:        evt.FromDate,
		ToDate:          evt.ToDate,
		Location:        evt.EventLocation,
		Description:     evt.Descriptions,
		PersonConfirmed: personConfirmed,
	}

	return eventObject, nil
}