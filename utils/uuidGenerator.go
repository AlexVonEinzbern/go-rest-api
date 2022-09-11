package utils

import "github.com/google/uuid"

func IdGenerator() string {
	id := uuid.NewString()
	return id
}
