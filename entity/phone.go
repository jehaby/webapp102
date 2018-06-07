package entity

import uuid "github.com/satori/go.uuid"

type Phone struct {
	UUID        uuid.UUID `json:"uuid"`
	CountryCode int64     `json:"country_code"`
	Number      string    `json:"number"`
	UserUUID    uuid.UUID `json:"user_uuid"`
}
