package entities

import (
	"time"

	"github.com/google/uuid"

	"gitlab.com/4nd3rs0n/username"
)

type User struct {
	Id        uuid.UUID         `json:"id"`
	Username  username.Username `json:"username"`
	Name      string            `json:"name"`
	EMail     string            `json:"email"`
	CreatedAt time.Time         `json:"createdAt"`
}

type PublicUserInfo struct {
	Id        uuid.UUID         `json:"id"`
	Username  username.Username `json:"username"`
	Name      string            `json:"name"`
	CreatedAt time.Time         `json:"createdAt"`
}
