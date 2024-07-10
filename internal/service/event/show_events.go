package event

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/e-calender/internal/dto"
	// "golang.org/x/e-calender/entity"
)

func (s *EventService) CheckEventByDateDay(username string, queries map[string][]string) ([]*dto.EventDto, error) {
	start := queries["start_time"][0]
	end := queries["end_time"][0]

	events, err := s.EventRepository.FindEventByMonth(username, start, end)
	// log.Println(events)

	var allEvents []*dto.EventDto

	for _, event := range events {
		newEvent := &dto.EventDto {
			Id: event.Id,
			Title: event.Title,
			IdCategory: event.CategoryId,
			// Location: event.Location,
			Date: event.Date.String(),
		}

		allEvents = append(allEvents, newEvent)
	}

	if err != nil {
		return nil, err
	}

	log.Println(allEvents)

	return allEvents, nil
}
