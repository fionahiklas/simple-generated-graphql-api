// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AlarmSystem struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Sensors     []*Sensor `json:"sensors"`
}

type Home struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	AlarmSystem *AlarmSystem `json:"alarmSystem"`
}

type NewHome struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Sensor struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
