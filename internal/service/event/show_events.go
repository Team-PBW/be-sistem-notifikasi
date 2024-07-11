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
		// date := event.Date.Format("2006-05-01")
		// startTime := event.StartTime.Format("08:01:09")
		// endTime := event.EndTime.Format("08:01:09")
		newEvent := &dto.EventDto {
			Id: event.Id,
			Title: event.Title,
			IdCategory: event.CategoryId,
			Location: event.Location,
			Description: event.Description,
			Date:        event.Date.Format("2006-01-02"),      // Format YYYY-MM-DD
			StartTime:   event.StartTime.Format("15:04:05"),   // Format HH:MM:SS
			EndTime:     event.EndTime.Format("15:04:05"),     // Format HH:MM:SS
		}

		allEvents = append(allEvents, newEvent)
	}

	if err != nil {
		return nil, err
	}

	log.Println(allEvents)

	return allEvents, nil
}
