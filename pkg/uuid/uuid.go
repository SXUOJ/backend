package uuid

import (
	uuid "github.com/satori/go.uuid"
)

func Getuuid() (id string, err error) {
	u1 := uuid.Must(uuid.NewV4(), err)
	id = u1.String()
	return id, err
}

func ParseUuid(id string) (string, error) {
	u2, err := uuid.FromString(id)
	if err != nil {
		return "", err
	}
	return u2.String(), nil
}
