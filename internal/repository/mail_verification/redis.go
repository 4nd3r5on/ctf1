package mail_repo

import (
	"context"

	"github.com/4nd3r5on/ctf1/internal/repository"
	redis "github.com/redis/go-redis/v9"
)

type redisMailVerificationRepository struct {
	rdb *redis.Client
}

func (mvr *redisMailVerificationRepository) AddMail(ctx context.Context, mail string, id string) error {
	err := mvr.rdb.Set(ctx, "waitlist."+id, mail, timelife).Err()
	return err
}

func (mvr *redisMailVerificationRepository) GetMailByID(ctx context.Context, id string) (string, error) {
	val, err := mvr.rdb.Get(ctx, "waitlist."+id).Result()
	if err != nil {
		if err == redis.Nil {
			return "", repository.ErrEntityNotFound
		}
		return "", err
	}
	return val, nil
}

func (mvr *redisMailVerificationRepository) RemoveID(ctx context.Context, id string) error {
	err := mvr.rdb.Del(ctx, "waitlist"+id).Err()
	return err
}

func (mvr *redisMailVerificationRepository) IsInRepo(mail string) (bool, error) {
	return false, nil
}
