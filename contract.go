package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type RedisService interface {
	Context() context.Context
	WithContext(ctx context.Context) *redis.Client
	Options() *redis.Options
	PoolStats() *redis.PoolStats
	Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	TxPipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	TxPipeline() redis.Pipeliner
	Subscribe(channels ...string) *redis.PubSub
	PSubscribe(channels ...string) *redis.PubSub
	Conn(ctx context.Context) *redis.Conn

	Command(ctx context.Context) *redis.CommandsInfoCmd
	ClientGetName(ctx context.Context) *redis.StringCmd
	Echo(ctx context.Context, message interface{}) *redis.StringCmd
	Ping(ctx context.Context) *redis.StatusCmd
	Quit(ctx context.Context) *redis.StatusCmd
	Del(ctx context.Context, XClaim ...string) *redis.IntCmd
	Unlink(ctx context.Context, XClaim ...string) *redis.IntCmd
	Dump(ctx context.Context, key string) *redis.StringCmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd
	Keys(ctx context.Context, pattern string) *redis.StringSliceCmd
	Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *redis.StatusCmd
	Move(ctx context.Context, key string, db int) *redis.BoolCmd
	ObjectRefCount(ctx context.Context, key string) *redis.IntCmd
	ObjectEncoding(ctx context.Context, key string) *redis.StringCmd
	ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd
	Persist(ctx context.Context, key string) *redis.BoolCmd
	PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd
	PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd
	PTTL(ctx context.Context, key string) *redis.DurationCmd
	RandomKey(ctx context.Context) *redis.StringCmd
	Rename(ctx context.Context, key, newkey string) *redis.StatusCmd
	RenameNX(ctx context.Context, key, newkey string) *redis.BoolCmd
	Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd
	RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd
	Sort(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd
	SortStore(ctx context.Context, key, store string, sort *redis.Sort) *redis.IntCmd
	SortInterfaces(ctx context.Context, key string, sort *redis.Sort) *redis.SliceCmd
	Touch(ctx context.Context, keys ...string) *redis.IntCmd
	TTL(ctx context.Context, key string) *redis.DurationCmd
	Type(ctx context.Context, key string) *redis.StatusCmd
	Append(ctx context.Context, key, value string) *redis.IntCmd
	Decr(ctx context.Context, key string) *redis.IntCmd
	DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd
	GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
	IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd
	IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd
	MGet(ctx context.Context, keys ...string) *redis.SliceCmd
	MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd
	MSetNX(ctx context.Context, values ...interface{}) *redis.BoolCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd
	StrLen(ctx context.Context, key string) *redis.IntCmd

	GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd
	SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd
	BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd
	BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd
	BitOpNot(ctx context.Context, destKey string, key string) *redis.IntCmd
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd

	Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd

	HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd
	HExists(ctx context.Context, key, field string) *redis.BoolCmd
	HGet(ctx context.Context, key, field string) *redis.StringCmd
	HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd
	HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd
	HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd
	HKeys(ctx context.Context, key string) *redis.StringSliceCmd
	HLen(ctx context.Context, key string) *redis.IntCmd
	HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd
	HSetNX(ctx context.Context, key, field string, value interface{}) *redis.BoolCmd
	HVals(ctx context.Context, key string) *redis.StringSliceCmd

	BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd
	LIndex(ctx context.Context, key string, index int64) *redis.StringCmd
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd
	LLen(ctx context.Context, key string) *redis.IntCmd
	LPop(ctx context.Context, key string) *redis.StringCmd
	LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd
	LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd
	LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd
	RPop(ctx context.Context, key string) *redis.StringCmd
	RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd
	RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd

	SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	SCard(ctx context.Context, key string) *redis.IntCmd
	SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd
	SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd
	SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd
	SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd
	SMembers(ctx context.Context, key string) *redis.StringSliceCmd
	SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd
	SMove(ctx context.Context, source, destination string, member interface{}) *redis.BoolCmd
	SPop(ctx context.Context, key string) *redis.StringCmd
	SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd
	SRandMember(ctx context.Context, key string) *redis.StringCmd
	SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd
	SRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd
	SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd

	XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd
	XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd
	XLen(ctx context.Context, stream string) *redis.IntCmd
	XRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd
	XRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd
	XRevRange(ctx context.Context, stream string, start, stop string) *redis.XMessageSliceCmd
	XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *redis.XMessageSliceCmd
	XRead(ctx context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd
	XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd
	XGroupCreate(ctx context.Context, stream, group, start string) *redis.StatusCmd
	XGroupCreateMkStream(ctx context.Context, stream, group, start string) *redis.StatusCmd
	XGroupSetID(ctx context.Context, stream, group, start string) *redis.StatusCmd
	XGroupDestroy(ctx context.Context, stream, group string) *redis.IntCmd
	XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd
	XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd
	XAck(ctx context.Context, stream, group string, ids ...string) *redis.IntCmd
	XPending(ctx context.Context, stream, group string) *redis.XPendingCmd
	XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd
	XClaim(ctx context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd
	XClaimJustID(ctx context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd
	XTrim(ctx context.Context, key string, maxLen int64) *redis.IntCmd
	XTrimApprox(ctx context.Context, key string, maxLen int64) *redis.IntCmd

	BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd
	BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd
	ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddNX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddXX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddNXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZAddXXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZIncr(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	ZIncrNX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	ZIncrXX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd
	ZCard(ctx context.Context, key string) *redis.IntCmd
	ZCount(ctx context.Context, key, min, max string) *redis.IntCmd
	ZLexCount(ctx context.Context, key, min, max string) *redis.IntCmd
	ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd
	ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd
	ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd
	ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd
	ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd
	ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	ZRank(ctx context.Context, key, member string) *redis.IntCmd
	ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd
	ZRemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd
	ZRemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd
	ZRevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd
	ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd
	ZRevRank(ctx context.Context, key, member string) *redis.IntCmd
	ZScore(ctx context.Context, key, member string) *redis.FloatCmd
	ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd

	PFAdd(ctx context.Context, key string, els ...interface{}) *redis.IntCmd
	PFCount(ctx context.Context, keys ...string) *redis.IntCmd
	PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd

	BgRewriteAOF(ctx context.Context) *redis.StatusCmd
	BgSave(ctx context.Context) *redis.StatusCmd
	ClientKill(ctx context.Context, ipPort string) *redis.StatusCmd
	ClientKillByFilter(ctx context.Context, keys ...string) *redis.IntCmd
	ClientList(ctx context.Context) *redis.StringCmd
	ClientPause(ctx context.Context, dur time.Duration) *redis.BoolCmd
	ClientID(ctx context.Context) *redis.IntCmd
	ConfigGet(ctx context.Context, parameter string) *redis.SliceCmd
	ConfigResetStat(ctx context.Context) *redis.StatusCmd
	ConfigSet(ctx context.Context, parameter, value string) *redis.StatusCmd
	ConfigRewrite(ctx context.Context) *redis.StatusCmd
	DBSize(ctx context.Context) *redis.IntCmd
	FlushAll(ctx context.Context) *redis.StatusCmd
	FlushAllAsync(ctx context.Context) *redis.StatusCmd
	FlushDB(ctx context.Context) *redis.StatusCmd
	FlushDBAsync(ctx context.Context) *redis.StatusCmd
	Info(ctx context.Context, section ...string) *redis.StringCmd
	LastSave(ctx context.Context) *redis.IntCmd
	Save(ctx context.Context) *redis.StatusCmd
	Shutdown(ctx context.Context) *redis.StatusCmd
	ShutdownSave(ctx context.Context) *redis.StatusCmd
	ShutdownNoSave(ctx context.Context) *redis.StatusCmd
	SlaveOf(ctx context.Context, host, port string) *redis.StatusCmd
	Time(ctx context.Context) *redis.TimeCmd
	DebugObject(ctx context.Context, key string) *redis.StringCmd
	ReadOnly(ctx context.Context) *redis.StatusCmd
	ReadWrite(ctx context.Context) *redis.StatusCmd
	MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd

	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd
	ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd
	ScriptFlush(ctx context.Context) *redis.StatusCmd
	ScriptKill(ctx context.Context) *redis.StatusCmd
	ScriptLoad(ctx context.Context, script string) *redis.StringCmd

	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
	PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd
	PubSubNumSub(ctx context.Context, channels ...string) *redis.StringIntMapCmd
	PubSubNumPat(ctx context.Context) *redis.IntCmd

	ClusterSlots(ctx context.Context) *redis.ClusterSlotsCmd
	ClusterNodes(ctx context.Context) *redis.StringCmd
	ClusterMeet(ctx context.Context, host, port string) *redis.StatusCmd
	ClusterForget(ctx context.Context, nodeID string) *redis.StatusCmd
	ClusterReplicate(ctx context.Context, nodeID string) *redis.StatusCmd
	ClusterResetSoft(ctx context.Context) *redis.StatusCmd
	ClusterResetHard(ctx context.Context) *redis.StatusCmd
	ClusterInfo(ctx context.Context) *redis.StringCmd
	ClusterKeySlot(ctx context.Context, key string) *redis.IntCmd
	ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *redis.StringSliceCmd
	ClusterCountFailureReports(ctx context.Context, nodeID string) *redis.IntCmd
	ClusterCountKeysInSlot(ctx context.Context, slot int) *redis.IntCmd
	ClusterDelSlots(ctx context.Context, slots ...int) *redis.StatusCmd
	ClusterDelSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd
	ClusterSaveConfig(ctx context.Context) *redis.StatusCmd
	ClusterSlaves(ctx context.Context, nodeID string) *redis.StringSliceCmd
	ClusterFailover(ctx context.Context) *redis.StatusCmd
	ClusterAddSlots(ctx context.Context, slots ...int) *redis.StatusCmd
	ClusterAddSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd

	GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd
	GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd
	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.IntCmd
	GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd
	GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.IntCmd
	GeoDist(ctx context.Context, key string, member1, member2, unit string) *redis.FloatCmd
	GeoHash(ctx context.Context, key string, members ...string) *redis.StringSliceCmd
}
