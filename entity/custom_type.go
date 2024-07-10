package entity

import (
    "database/sql/driver"
    "fmt"
    "time"
)

// CustomTime is a custom type for handling TIME fields
type CustomTime struct {
    time.Time
}

// Scan implements the sql.Scanner interface for CustomTime
func (ct *CustomTime) Scan(value interface{}) error {
    if value == nil {
        *ct = CustomTime{Time: time.Time{}}
        return nil
    }

    switch v := value.(type) {
    case time.Time:
        *ct = CustomTime{Time: v}
    case []byte:
        t, err := time.Parse("15:04:05", string(v))
        if err != nil {
            return err
        }
        *ct = CustomTime{Time: t}
    case string:
        t, err := time.Parse("15:04:05", v)
        if err != nil {
            return err
        }
        *ct = CustomTime{Time: t}
    default:
        return fmt.Errorf("cannot scan type %T into CustomTime", value)
    }
    return nil
}

// Value implements the driver.Valuer interface for CustomTime
func (ct CustomTime) Value() (driver.Value, error) {
    return ct.Time.Format("15:04:05"), nil
}

// CustomDatetime is a custom type for handling DATETIME fields
type CustomDatetime struct {
    time.Time
}

// Scan implements the sql.Scanner interface for CustomDatetime
func (cd *CustomDatetime) Scan(value interface{}) error {
    if value == nil {
        *cd = CustomDatetime{Time: time.Time{}}
        return nil
    }

    switch v := value.(type) {
    case time.Time:
        *cd = CustomDatetime{Time: v}
    case []byte:
        t, err := time.Parse("2006-01-02 15:04:05", string(v))
        if err != nil {
            return err
        }
        *cd = CustomDatetime{Time: t}
    case string:
        t, err := time.Parse("2006-01-02 15:04:05", v)
        if err != nil {
            return err
        }
        *cd = CustomDatetime{Time: t}
    default:
        return fmt.Errorf("cannot scan type %T into CustomDatetime", value)
    }
    return nil
}

// Value implements the driver.Valuer interface for CustomDatetime
func (cd CustomDatetime) Value() (driver.Value, error) {
    return cd.Time.Format("2006-01-02 15:04:05"), nil
}