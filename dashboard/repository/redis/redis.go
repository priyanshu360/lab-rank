package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type RedisSessionRepository struct {
	client *redis.Client
}

func NewRedisSessionRepository(client *redis.Client) *RedisSessionRepository {
	return &RedisSessionRepository{client: client}
}

const sessionKeyPrefix = "session:"

func (r *RedisSessionRepository) GetSession(ctx context.Context, sessionID uuid.UUID) (*models.AuthSession, models.AppError) {
	key := sessionKeyPrefix + sessionID.String()
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, models.InternalError.Add(err)
	}

	var session models.AuthSession
	err = json.Unmarshal([]byte(val), &session)
	if err != nil {
		return nil, models.InternalError.Add(err)
	}

	return &session, models.NoError
}

func (r *RedisSessionRepository) SetSession(ctx context.Context, session *models.AuthSession) (uuid.UUID, models.AppError) {
	sessionID := uuid.New()
	key := sessionKeyPrefix + sessionID.String()

	// Serialize session to JSON
	sessionData, err := json.Marshal(session)
	if err != nil {
		// Handle error (failed to marshal JSON, etc.)
		return uuid.Nil, models.InternalError.Add(err)
	}

	// Set session in Redis with an expiration time (adjust as needed)
	err = r.client.Set(ctx, key, sessionData, 24*time.Hour).Err()
	if err != nil {
		// Handle error (failed to set session in Redis, etc.)
		return uuid.Nil, models.InternalError.Add(err)
	}

	return sessionID, models.NoError
}
