package entities

import (
	"time"

	"github.com/google/uuid"

	"gitlab.com/4nd3rs0n/username"
)

type User struct {
	Id        uuid.UUID
	Username  username.Username
	Name      string
	EMail     string
	CreatedAt time.Time
}

type PublicUserInfo struct {
	Id        uuid.UUID
	Username  username.Username
	Name      string
	CreatedAt time.Time
}
