package event

import("time")

type Event struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type DTOEvent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        time.Time `json:"date"`
}
