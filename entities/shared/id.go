package shared

import (
	"log"
	"github.com/google/uuid"
)

func GenID() uuid.UUID {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Error on generating UUID: ", err)
	}
	return uuid
}

func GetUUIDFromString(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func GetUUIDEmpty() (uuid.UUID) {
	return uuid.Nil
}