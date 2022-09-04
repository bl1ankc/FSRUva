package utils

import uuid "github.com/satori/go.uuid"

func GetUid() string {
	uid := uuid.NewV4()

	NewUid := uid.String()

	return NewUid
}
