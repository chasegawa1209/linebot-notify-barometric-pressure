package model

import (
    "github.com/kyokomi/emoji"
)

const (
    PressureLevelIntOK0             = iota
    PressureLevelIntOK1
    PressureLevelIntSomewhatWarning
    PressureLevelIntWarning
    PressureLevelIntBomb
)

const (
    PressureLevelStrOK              = "通常"
    PressureLevelStrSomewhatWarning = "やや注意"
    PressureLevelStrWarning         = "注意"
    PressureLevelStrBomb            = "警戒"
)

var PressureLevelMap = map[int]string {
    PressureLevelIntOK0:             PressureLevelStrOK,
    PressureLevelIntOK1:             PressureLevelStrOK,
    PressureLevelIntSomewhatWarning: emoji.Sprint(":arrow_heading_down:") + PressureLevelStrSomewhatWarning,
    PressureLevelIntWarning:         emoji.Sprint(":warning:") + PressureLevelStrWarning,
    PressureLevelIntBomb:            emoji.Sprint(":bomb:") + PressureLevelStrBomb,
}

type BarometricPressure struct {
    NowLevel        int
    After1HourLevel int
    After2HourLevel int
}

type BarometricPressuresByZutool struct {
    PlaceName        string                        `json:"place_name"`
    PlaceID          int                           `json:"place_id"`
    PrefecturesID    int                           `json:"prefecture_id"`
    DateTime         string                        `json:"dateTime"`
    Yesterday        []*BarometricPressureByZutool `json:"yesterday"`
    Today            []*BarometricPressureByZutool `json:"today"`
    Tomorrow         []*BarometricPressureByZutool `json:"tomorrow"`
    DayAfterTomorrow []*BarometricPressureByZutool `json:"dayaftertomorrow"`
}

type BarometricPressureByZutool struct {
    Time          int     `json:"time"`
    Weather       int     `json:"weather"`
    Temp          float64 `json:"temp"`
    Pressure      float64 `json:"pressure"`
    PressureLevel int     `json:"pressure_level"`
}
