package repository

import (
	"errors"
	"log"
	"time"

	"golang.org/x/e-calender/entity"
	// "golang.org/x/e-calender/internal/dto"
	"golang.org/x/e-calender/model"
	"gorm.io/gorm"
)

var (
	errDatabase = errors.New("Something wrong")
)

type EventRepository struct {
	TX *gorm.DB
}

func NewEventRepository(tx *gorm.DB) *EventRepository {
	return &EventRepository{
		TX: tx,
	}
}

func (e *EventRepository) BeginTransaction() *gorm.DB {
	return e.TX.Begin()
}

func (e *EventRepository) FindEventByMonth(username string, start string, end string) ([]entity.EventEntity, error) {
    var events []entity.EventEntity

    // startDate, err := time.Parse("2006-01-02", start)
    // if err != nil {
    //     return nil, err
    // }
    // endDate, err := time.Parse("2006-01-02", end)
    // if err != nil {
    //     return nil, err
    // }

	const (
		dateLayout     = "2006-01-02"
		timeLayout     = "15:04:05"
		dateTimeLayout = "2006-01-02 15:04:05"
	)

	startFilterDate, err := time.Parse(dateLayout, start)
	if err != nil {
		log.Println("Error parsing date:", err)
		return nil, err
	}

	endFilterDate, err := time.Parse(dateLayout, end)
	if err != nil {
		log.Println("Error parsing date:", err)
		return nil, err
	}

	startDate := startFilterDate.Format("2006-01-02")
	endDate := endFilterDate.Format("2006-01-02")

	log.Println(start)
	log.Println(startDate)
	log.Println(endDate)

	var data []entity.EventEntity

    err = e.TX.Joins("JOIN followed_event_entities ON followed_event_entities.event_id = event_entities.id").
		// Select("id", "title", "location", "distance", "description", "date", "category_id", "start_time").
		Find(&data).
        Where("followed_event_entities.username = ?", username).
        Where("event_entities.date BETWEEN ? AND ?", startDate, endDate).
		Where("event_entities.bentrok = ?", 0).
        Find(&events).Error

    if err != nil {
        return nil, err
    }

    // for i, event := range events {
	// 	// avoid error: unsupported Scan, storing driver.Value type []uint8 into type *time.Time
    //     events[i].StartTime, err = time.Parse(time.RFC3339, event.StartTime.(string))
    //     if err != nil {
    //         return nil, err
    //     }
    // }

    return events, nil
}

func (e *EventRepository) CreateEvent(user interface{}, event *entity.EventEntity, person map[string][]string) error {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		e.TX.Rollback()
	// 	}
	// }()

	if err := e.TX.Create(&event).Error; err != nil {
		// e.TX.Rollback()
		return errDatabase
	}

	// var allPerson []*model.EventPersonConfirmed

	// for _, username := range person[event.Id] {
	// 	personConfirmed := &model.EventPersonConfirmed{
	// 		Id:          event.Id,
	// 		Username:    username,
	// 		IsConfirmed: false,
	// 	}
	// 	allPerson = append(allPerson, personConfirmed)
	// }

	// err := e.TX.CreateInBatches(allPerson, len(person[event.Id]))
	// if err != nil {
	// 	// e.TX.Rollback()
	// 	return errDatabase
	// }

	return nil
}

func (e *EventRepository) CheckPersonExist(id string, persons []string) ([]*entity.UserEntity, error) {
	var users []*entity.UserEntity

	err := e.TX.Where(persons).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (e *EventRepository) InvitePersonToEvent(users []*entity.FollowedEventEntity) error {
	// var users []entity.UserEntity

	err := e.TX.Create(&users).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *EventRepository) Update(id string, evtEntity *model.Event) (*model.Event, error) {
	defer func() {
		if r := recover(); r != nil {
			e.TX.Rollback()
		}
	}()

	if err := e.TX.Model(&model.Event{}).Updates(evtEntity).Error; err != nil {
		e.TX.Rollback()
		return nil, errDatabase
	}

	return evtEntity, e.TX.Commit().Error
}

func (e *EventRepository) FindGuestsInEvent(idEvt string) ([]*entity.FollowedEventEntity, error) {
	var guestRecorded []*entity.FollowedEventEntity
	// var guests []*model.User

	err := e.TX.Model(&entity.FollowedEventEntity{}).Where("id = ?", idEvt).Find(&guestRecorded).Error
	if err != nil {
		// e.TX.Rollback()
		return nil, errDatabase
	}

	// for _, guestEvt := range guestRecorded {
	// 	guest := &model.User{
	// 		Username: guestEvt.Username,
	// 		PhoneNumber: guestEvt.PhoneNumber,
	// 	}
	// 	guests = append(guests, guest)
	
	// }

	// if err := e.TX.Commit().Error; err != nil {
	// 	return nil, errDatabase
	// }

	return guestRecorded, nil
}

func (e *EventRepository) Delete(id string) error {
	defer func() {
		if r := recover(); r != nil {
			e.TX.Rollback()
		}
	}()

	err := e.TX.Where("id = ?", id).Delete(&model.Event{}).Error
	if err != nil {
		e.TX.Rollback()
		return err
	}

	err = e.TX.Where("id = ?", id).Delete(&model.EventPersonConfirmed{}).Error
	if err != nil {
		e.TX.Rollback()
		return err
	}
	return e.TX.Commit().Error
}

func (e *EventRepository) FindEventByID(id string) (*entity.EventEntity, error) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		e.TX.Rollback()
	// 	}
	// }()

	var event *entity.EventEntity

	if err := e.TX.Where("id = ?").Find(&event).Error; err != nil {
		// e.TX.Rollback()
		return nil, err
	}

	return event, nil
}

func (e *EventRepository) FindEventsByHost(username string) ([]*model.Event, error) {
	defer func() {
		if r := recover(); r != nil {
			e.TX.Rollback()
		}
	}()

	var event []*model.Event
	err := e.TX.First(&event, "username = ?", username).Order("from_date DESC")
	if err != nil {
		e.TX.Rollback()
		return nil, errDatabase
	}

	return event, e.TX.Commit().Error
}

// func (e *EventRepository) EventByMonth() {
// 	var events map[string]interface{}
// 	err := e.TX.Model()
// }

// func (e *EventRepository) UpdateGuestByEventID(id string, guests []*model.EveryPerson) (*model.EventPersonConfirmed, error) {

// }

func (e *EventRepository) CheckEventExist(start time.Time, end time.Time, date time.Time) (bool) {
	// convert uint8 to time.Time
	dateString := date.Format("2006-01-02")

	// startTimeString := start.Format("15:04:05")
	// endTimeString := end.Format("15:04:05")

	// startNew := date + start
	// endNew := date + end

	startConvert := start.Add(-8 * time.Hour)
	endConvert := end.Add(-8 * time.Hour)


	var count int64
	err := e.TX.Table("event_entities").Where("date = ?", dateString).Where("start_time BETWEEN ? AND ? OR end_time BETWEEN ? AND ?", startConvert, endConvert, startConvert, endConvert).Count(&count).Error
	if err != nil || count == 1 {
		return false
	}

	log.Println("wii: ", count)

	return count > 1
}

func (e *EventRepository) UpdateBentrok(idEvt string) error {
	return e.TX.Table("event_entities").Where("id = ?", idEvt).Update("bentrok", 1).Error
}