package mail_repo

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const timelife = time.Hour

type MailVerificationRepository interface {
	AddMail(ctx context.Context, mail string, approvementID string) error
	GetMailByID(ctx context.Context, id string) (mail string, err error)
	RemoveID(ctx context.Context, id string) error
	// TODO:
	// IsMailInRepo() (bool, error)
	// ResendMail(id string) (newId string, err error)
}

func NewMailApprovementsRepository(rdb *redis.Client) MailVerificationRepository {
	return &redisMailVerificationRepository{
		rdb: rdb,
	}
}
