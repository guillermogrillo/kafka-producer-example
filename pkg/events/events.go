package events

import "time"

type TollRecord struct {
	Id        string
	CreatedAt time.Time
	TollId    string
}
