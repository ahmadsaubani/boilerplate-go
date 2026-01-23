package Config

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

type DbSqlConfigName string
type dbRedisConfigName string

const (
	databaseMain DbSqlConfigName   = "mainDB"
	redisMainDb  dbRedisConfigName = "mainDBRedis"
)

type DbConInterface interface {
	PostgreMainCon() *sql.DB
	setSqlConnection(conName DbSqlConfigName, con *sql.DB)
	setRedisConnection(conName dbRedisConfigName, con *redis.Client)
}

type DbCon struct {
	sql   map[DbSqlConfigName]*sql.DB
	redis map[dbRedisConfigName]*redis.Client
}

func (d DbCon) PostgreMainCon() *sql.DB {
	return d.sql[databaseMain]
}

func (d DbCon) RedisMainCon() redisFunctionInterface {
	return &redisFunc{connection: d.redis[redisMainDb]}
}

func (d *DbCon) setSqlConnection(conName DbSqlConfigName, con *sql.DB) {
	if d.sql == nil {
		d.sql = make(map[DbSqlConfigName]*sql.DB)
	}
	d.sql[conName] = con
}

func (d *DbCon) setRedisConnection(conName dbRedisConfigName, con *redis.Client) {
	if d.redis == nil {
		d.redis = make(map[dbRedisConfigName]*redis.Client)
	}
	d.redis[conName] = con
}

type redisFunctionInterface interface {
	Set(key string, data interface{}, expired time.Duration) (err error)
	Get(key string, data interface{}) (err error)
	Del(keys ...string) (err error)
	HMSet(key string, data map[string]interface{}, expired time.Duration) (err error)
	HGetAll(key string) (result map[string]string, err error)
	HDel(key, field string) (err error)
	HGet(key, field string) (result string, err error)
	FindKeysInRedis(formatKeys string) (keys []string, err error)
	SetExpire(key string, expTime time.Duration) (err error)
	RedisPipeline(txFunc func(redis.Pipeliner) error) (cmd []redis.Cmder, err error)
	Exists(key string) (exists bool, err error)
	HIncrBy(key string, field string, value int64) (err error)
	SMembers(key string) (members []string, err error)
	RedisBeginTx() redis.Pipeliner
	SisMember(cacheKey string, key string) (exists bool, err error)
	Copy(sourceKey, destKey string) (res bool, err error)
	ZRevRank(key, memberKey string) (rank int64, err error)
	SRem(formatKeys, member string) (value int64, err error)
	HMSetExpiredAt(key string, data map[string]interface{}, expiredAt time.Time) (err error)
}

type redisFunc struct {
	connection *redis.Client
}

func (r redisFunc) Copy(sourceKey, destKey string) (res bool, err error) {
	res, err = r.connection.Do("COPY", sourceKey, destKey).Bool()
	return
}

func (r redisFunc) SMembers(key string) (members []string, err error) {
	members, err = r.connection.SMembers(key).Result()
	return
}

func (r redisFunc) RedisBeginTx() redis.Pipeliner {
	return r.connection.TxPipeline()
}

func (r redisFunc) Set(key string, data interface{}, expired time.Duration) (err error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = r.connection.Set(key, string(jsonBytes), expired).Err()
	return
}

func (r redisFunc) Get(key string, data interface{}) (err error) {
	res, err := r.connection.Get(key).Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(res), &data)
	return
}

func (r redisFunc) Del(keys ...string) (err error) {
	err = r.connection.Del(keys...).Err()
	if err != nil {
		return
	}
	return
}

func (r redisFunc) HMSet(key string, data map[string]interface{}, expired time.Duration) (err error) {
	err = r.connection.HMSet(key, data).Err()
	if err != nil {
		return
	}
	err = r.connection.Expire(key, expired).Err()
	return
}

func (r redisFunc) HGetAll(key string) (result map[string]string, err error) {
	result, err = r.connection.HGetAll(key).Result()
	return
}

func (r redisFunc) HDel(key, field string) (err error) {
	_, err = r.connection.HDel(key, field).Result()
	return
}

func (r redisFunc) HGet(key, field string) (result string, err error) {
	result, err = r.connection.HGet(key, field).Result()
	if err != nil {
		return
	}
	return
}

func (r redisFunc) FindKeysInRedis(formatKeys string) (keys []string, err error) {
	keys, err = r.connection.Keys(formatKeys).Result()
	if err != nil {
		return
	}
	return
}

func (r redisFunc) SetExpire(key string, expTime time.Duration) (err error) {
	err = r.connection.Expire(key, expTime).Err()
	if err != nil {
		return
	}
	return
}

func (r redisFunc) RedisPipeline(txFunc func(redis.Pipeliner) error) (cmd []redis.Cmder, err error) {
	cmd, err = r.connection.TxPipelined(txFunc)
	return
}

func (r redisFunc) Exists(key string) (exists bool, err error) {
	res, err := r.connection.Exists(key).Result()
	if err != nil {
		return
	}

	if res > 0 {
		exists = true
		return
	}
	return
}

func (r redisFunc) HIncrBy(key string, field string, value int64) (err error) {
	err = r.connection.HIncrBy(key, field, value).Err()
	return
}

func (r redisFunc) SisMember(cacheKey string, key string) (exists bool, err error) {
	exists, err = r.connection.SIsMember(cacheKey, key).Result()
	return
}

func (r redisFunc) ZRevRank(key, memberKey string) (rank int64, err error) {
	rank, err = r.connection.ZRevRank(key, memberKey).Result()
	return
}

func (r redisFunc) SRem(formatKeys, member string) (value int64, err error) {
	_, err = r.connection.SRem(formatKeys, member).Result()
	if err != nil {
		return
	}
	return
}

func (r redisFunc) HMSetExpiredAt(key string, data map[string]interface{}, expiredAt time.Time) (err error) {
	err = r.connection.HMSet(key, data).Err()
	if err != nil {
		return
	}
	err = r.connection.ExpireAt(key, expiredAt).Err()
	return
}
