// Created by interfacer; DO NOT EDIT

package iredis

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v7"
)

// Client is an interface generated for "github.com/go-redis/redis/v7.Client".
type Client interface {
	AddHook(redis.Hook)
	Append(string, string) *redis.IntCmd
	BLPop(time.Duration, ...string) *redis.StringSliceCmd
	BRPop(time.Duration, ...string) *redis.StringSliceCmd
	BRPopLPush(string, string, time.Duration) *redis.StringCmd
	BZPopMax(time.Duration, ...string) *redis.ZWithKeyCmd
	BZPopMin(time.Duration, ...string) *redis.ZWithKeyCmd
	BgRewriteAOF() *redis.StatusCmd
	BgSave() *redis.StatusCmd
	BitCount(string, *redis.BitCount) *redis.IntCmd
	BitField(string, ...interface{}) *redis.IntSliceCmd
	BitOpAnd(string, ...string) *redis.IntCmd
	BitOpNot(string, string) *redis.IntCmd
	BitOpOr(string, ...string) *redis.IntCmd
	BitOpXor(string, ...string) *redis.IntCmd
	BitPos(string, int64, ...int64) *redis.IntCmd
	ClientGetName() *redis.StringCmd
	ClientID() *redis.IntCmd
	ClientKill(string) *redis.StatusCmd
	ClientKillByFilter(...string) *redis.IntCmd
	ClientList() *redis.StringCmd
	ClientPause(time.Duration) *redis.BoolCmd
	ClientUnblock(int64) *redis.IntCmd
	ClientUnblockWithError(int64) *redis.IntCmd
	Close() error
	ClusterAddSlots(...int) *redis.StatusCmd
	ClusterAddSlotsRange(int, int) *redis.StatusCmd
	ClusterCountFailureReports(string) *redis.IntCmd
	ClusterCountKeysInSlot(int) *redis.IntCmd
	ClusterDelSlots(...int) *redis.StatusCmd
	ClusterDelSlotsRange(int, int) *redis.StatusCmd
	ClusterFailover() *redis.StatusCmd
	ClusterForget(string) *redis.StatusCmd
	ClusterGetKeysInSlot(int, int) *redis.StringSliceCmd
	ClusterInfo() *redis.StringCmd
	ClusterKeySlot(string) *redis.IntCmd
	ClusterMeet(string, string) *redis.StatusCmd
	ClusterNodes() *redis.StringCmd
	ClusterReplicate(string) *redis.StatusCmd
	ClusterResetHard() *redis.StatusCmd
	ClusterResetSoft() *redis.StatusCmd
	ClusterSaveConfig() *redis.StatusCmd
	ClusterSlaves(string) *redis.StringSliceCmd
	ClusterSlots() *redis.ClusterSlotsCmd
	Command() *redis.CommandsInfoCmd
	ConfigGet(string) *redis.SliceCmd
	ConfigResetStat() *redis.StatusCmd
	ConfigRewrite() *redis.StatusCmd
	ConfigSet(string, string) *redis.StatusCmd
	Conn() *redis.Conn
	Context() context.Context
	DBSize() *redis.IntCmd
	DbSize() *redis.IntCmd
	DebugObject(string) *redis.StringCmd
	Decr(string) *redis.IntCmd
	DecrBy(string, int64) *redis.IntCmd
	Del(...string) *redis.IntCmd
	Do(...interface{}) *redis.Cmd
	DoContext(context.Context, ...interface{}) *redis.Cmd
	Dump(string) *redis.StringCmd
	Echo(interface{}) *redis.StringCmd
	Eval(string, []string, ...interface{}) *redis.Cmd
	EvalSha(string, []string, ...interface{}) *redis.Cmd
	Exists(...string) *redis.IntCmd
	Expire(string, time.Duration) *redis.BoolCmd
	ExpireAt(string, time.Time) *redis.BoolCmd
	FlushAll() *redis.StatusCmd
	FlushAllAsync() *redis.StatusCmd
	FlushDB() *redis.StatusCmd
	FlushDBAsync() *redis.StatusCmd
	GeoAdd(string, ...*redis.GeoLocation) *redis.IntCmd
	GeoDist(string, string, string, string) *redis.FloatCmd
	GeoHash(string, ...string) *redis.StringSliceCmd
	GeoPos(string, ...string) *redis.GeoPosCmd
	GeoRadius(string, float64, float64, *redis.GeoRadiusQuery) *redis.GeoLocationCmd
	GeoRadiusByMember(string, string, *redis.GeoRadiusQuery) *redis.GeoLocationCmd
	GeoRadiusByMemberStore(string, string, *redis.GeoRadiusQuery) *redis.IntCmd
	GeoRadiusStore(string, float64, float64, *redis.GeoRadiusQuery) *redis.IntCmd
	Get(string) *redis.StringCmd
	GetBit(string, int64) *redis.IntCmd
	GetRange(string, int64, int64) *redis.StringCmd
	GetSet(string, interface{}) *redis.StringCmd
	HDel(string, ...string) *redis.IntCmd
	HExists(string, string) *redis.BoolCmd
	HGet(string, string) *redis.StringCmd
	HGetAll(string) *redis.StringStringMapCmd
	HIncrBy(string, string, int64) *redis.IntCmd
	HIncrByFloat(string, string, float64) *redis.FloatCmd
	HKeys(string) *redis.StringSliceCmd
	HLen(string) *redis.IntCmd
	HMGet(string, ...string) *redis.SliceCmd
	HMSet(string, map[string]interface{}) *redis.StatusCmd
	HScan(string, uint64, string, int64) *redis.ScanCmd
	HSet(string, string, interface{}) *redis.BoolCmd
	HSetNX(string, string, interface{}) *redis.BoolCmd
	HVals(string) *redis.StringSliceCmd
	Incr(string) *redis.IntCmd
	IncrBy(string, int64) *redis.IntCmd
	IncrByFloat(string, float64) *redis.FloatCmd
	Info(...string) *redis.StringCmd
	Keys(string) *redis.StringSliceCmd
	LIndex(string, int64) *redis.StringCmd
	LInsert(string, string, interface{}, interface{}) *redis.IntCmd
	LInsertAfter(string, interface{}, interface{}) *redis.IntCmd
	LInsertBefore(string, interface{}, interface{}) *redis.IntCmd
	LLen(string) *redis.IntCmd
	LPop(string) *redis.StringCmd
	LPush(string, ...interface{}) *redis.IntCmd
	LPushX(string, ...interface{}) *redis.IntCmd
	LRange(string, int64, int64) *redis.StringSliceCmd
	LRem(string, int64, interface{}) *redis.IntCmd
	LSet(string, int64, interface{}) *redis.StatusCmd
	LTrim(string, int64, int64) *redis.StatusCmd
	LastSave() *redis.IntCmd
	Lock()
	MGet(...string) *redis.SliceCmd
	MSet(...interface{}) *redis.StatusCmd
	MSetNX(...interface{}) *redis.BoolCmd
	MemoryUsage(string, ...int) *redis.IntCmd
	Migrate(string, string, string, int, time.Duration) *redis.StatusCmd
	Move(string, int) *redis.BoolCmd
	ObjectEncoding(string) *redis.StringCmd
	ObjectIdleTime(string) *redis.DurationCmd
	ObjectRefCount(string) *redis.IntCmd
	Options() *redis.Options
	PExpire(string, time.Duration) *redis.BoolCmd
	PExpireAt(string, time.Time) *redis.BoolCmd
	PFAdd(string, ...interface{}) *redis.IntCmd
	PFCount(...string) *redis.IntCmd
	PFMerge(string, ...string) *redis.StatusCmd
	PSubscribe(...string) *redis.PubSub
	PTTL(string) *redis.DurationCmd
	Persist(string) *redis.BoolCmd
	Ping() *redis.StatusCmd
	Pipeline() redis.Pipeliner
	Pipelined(func(redis.Pipeliner) error) ([]redis.Cmder, error)
	PoolStats() *redis.PoolStats
	Process(redis.Cmder) error
	ProcessContext(context.Context, redis.Cmder) error
	PubSubChannels(string) *redis.StringSliceCmd
	PubSubNumPat() *redis.IntCmd
	PubSubNumSub(...string) *redis.StringIntMapCmd
	Publish(string, interface{}) *redis.IntCmd
	Quit() *redis.StatusCmd
	RPop(string) *redis.StringCmd
	RPopLPush(string, string) *redis.StringCmd
	RPush(string, ...interface{}) *redis.IntCmd
	RPushX(string, ...interface{}) *redis.IntCmd
	RandomKey() *redis.StringCmd
	ReadOnly() *redis.StatusCmd
	ReadWrite() *redis.StatusCmd
	Rename(string, string) *redis.StatusCmd
	RenameNX(string, string) *redis.BoolCmd
	Restore(string, time.Duration, string) *redis.StatusCmd
	RestoreReplace(string, time.Duration, string) *redis.StatusCmd
	SAdd(string, ...interface{}) *redis.IntCmd
	SCard(string) *redis.IntCmd
	SDiff(...string) *redis.StringSliceCmd
	SDiffStore(string, ...string) *redis.IntCmd
	SInter(...string) *redis.StringSliceCmd
	SInterStore(string, ...string) *redis.IntCmd
	SIsMember(string, interface{}) *redis.BoolCmd
	SMembers(string) *redis.StringSliceCmd
	SMembersMap(string) *redis.StringStructMapCmd
	SMove(string, string, interface{}) *redis.BoolCmd
	SPop(string) *redis.StringCmd
	SPopN(string, int64) *redis.StringSliceCmd
	SRandMember(string) *redis.StringCmd
	SRandMemberN(string, int64) *redis.StringSliceCmd
	SRem(string, ...interface{}) *redis.IntCmd
	SScan(string, uint64, string, int64) *redis.ScanCmd
	SUnion(...string) *redis.StringSliceCmd
	SUnionStore(string, ...string) *redis.IntCmd
	Save() *redis.StatusCmd
	Scan(uint64, string, int64) *redis.ScanCmd
	ScriptExists(...string) *redis.BoolSliceCmd
	ScriptFlush() *redis.StatusCmd
	ScriptKill() *redis.StatusCmd
	ScriptLoad(string) *redis.StringCmd
	Set(string, interface{}, time.Duration) *redis.StatusCmd
	SetBit(string, int64, int) *redis.IntCmd
	SetLimiter(redis.Limiter) *redis.Client
	SetNX(string, interface{}, time.Duration) *redis.BoolCmd
	SetRange(string, int64, string) *redis.IntCmd
	SetXX(string, interface{}, time.Duration) *redis.BoolCmd
	Shutdown() *redis.StatusCmd
	ShutdownNoSave() *redis.StatusCmd
	ShutdownSave() *redis.StatusCmd
	SlaveOf(string, string) *redis.StatusCmd
	SlowLog()
	Sort(string, *redis.Sort) *redis.StringSliceCmd
	SortInterfaces(string, *redis.Sort) *redis.SliceCmd
	SortStore(string, string, *redis.Sort) *redis.IntCmd
	StrLen(string) *redis.IntCmd
	String() string
	Subscribe(...string) *redis.PubSub
	Sync()
	TTL(string) *redis.DurationCmd
	Time() *redis.TimeCmd
	Touch(...string) *redis.IntCmd
	TxPipeline() redis.Pipeliner
	TxPipelined(func(redis.Pipeliner) error) ([]redis.Cmder, error)
	Type(string) *redis.StatusCmd
	Unlink(...string) *redis.IntCmd
	Wait(int, time.Duration) *redis.IntCmd
	Watch(func(*redis.Tx) error, ...string) error
	WatchContext(context.Context, func(*redis.Tx) error, ...string) error
	WithContext(context.Context) *redis.Client
	XAck(string, string, ...string) *redis.IntCmd
	XAdd(*redis.XAddArgs) *redis.StringCmd
	XClaim(*redis.XClaimArgs) *redis.XMessageSliceCmd
	XClaimJustID(*redis.XClaimArgs) *redis.StringSliceCmd
	XDel(string, ...string) *redis.IntCmd
	XGroupCreate(string, string, string) *redis.StatusCmd
	XGroupCreateMkStream(string, string, string) *redis.StatusCmd
	XGroupDelConsumer(string, string, string) *redis.IntCmd
	XGroupDestroy(string, string) *redis.IntCmd
	XGroupSetID(string, string, string) *redis.StatusCmd
	XLen(string) *redis.IntCmd
	XPending(string, string) *redis.XPendingCmd
	XPendingExt(*redis.XPendingExtArgs) *redis.XPendingExtCmd
	XRange(string, string, string) *redis.XMessageSliceCmd
	XRangeN(string, string, string, int64) *redis.XMessageSliceCmd
	XRead(*redis.XReadArgs) *redis.XStreamSliceCmd
	XReadGroup(*redis.XReadGroupArgs) *redis.XStreamSliceCmd
	XReadStreams(...string) *redis.XStreamSliceCmd
	XRevRange(string, string, string) *redis.XMessageSliceCmd
	XRevRangeN(string, string, string, int64) *redis.XMessageSliceCmd
	XTrim(string, int64) *redis.IntCmd
	XTrimApprox(string, int64) *redis.IntCmd
	ZAdd(string, ...*redis.Z) *redis.IntCmd
	ZAddCh(string, ...*redis.Z) *redis.IntCmd
	ZAddNX(string, ...*redis.Z) *redis.IntCmd
	ZAddNXCh(string, ...*redis.Z) *redis.IntCmd
	ZAddXX(string, ...*redis.Z) *redis.IntCmd
	ZAddXXCh(string, ...*redis.Z) *redis.IntCmd
	ZCard(string) *redis.IntCmd
	ZCount(string, string, string) *redis.IntCmd
	ZIncr(string, *redis.Z) *redis.FloatCmd
	ZIncrBy(string, float64, string) *redis.FloatCmd
	ZIncrNX(string, *redis.Z) *redis.FloatCmd
	ZIncrXX(string, *redis.Z) *redis.FloatCmd
	ZInterStore(string, *redis.ZStore) *redis.IntCmd
	ZLexCount(string, string, string) *redis.IntCmd
	ZPopMax(string, ...int64) *redis.ZSliceCmd
	ZPopMin(string, ...int64) *redis.ZSliceCmd
	ZRange(string, int64, int64) *redis.StringSliceCmd
	ZRangeByLex(string, *redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByScore(string, *redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByScoreWithScores(string, *redis.ZRangeBy) *redis.ZSliceCmd
	ZRangeWithScores(string, int64, int64) *redis.ZSliceCmd
	ZRank(string, string) *redis.IntCmd
	ZRem(string, ...interface{}) *redis.IntCmd
	ZRemRangeByLex(string, string, string) *redis.IntCmd
	ZRemRangeByRank(string, int64, int64) *redis.IntCmd
	ZRemRangeByScore(string, string, string) *redis.IntCmd
	ZRevRange(string, int64, int64) *redis.StringSliceCmd
	ZRevRangeByLex(string, *redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByScore(string, *redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByScoreWithScores(string, *redis.ZRangeBy) *redis.ZSliceCmd
	ZRevRangeWithScores(string, int64, int64) *redis.ZSliceCmd
	ZRevRank(string, string) *redis.IntCmd
	ZScan(string, uint64, string, int64) *redis.ScanCmd
	ZScore(string, string) *redis.FloatCmd
	ZUnionStore(string, *redis.ZStore) *redis.IntCmd
}
