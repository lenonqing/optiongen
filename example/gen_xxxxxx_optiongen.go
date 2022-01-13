// Code generated by optiongen. DO NOT EDIT.
// optiongen: github.com/timestee/optiongen

package example

import (
	"sync/atomic"
	"time"
	"unsafe"
)

// XXXXXX should use NewXXXXXX to initialize it
type XXXXXX struct {
	OptionUsage      string            `xconf:"option_usage"`
	Endpoints        []string          `xconf:"endpoints"`
	ReadTimeout      time.Duration     `xconf:"read_timeout"`
	TypeMapIntString map[int]string    `xconf:"type_map_int_string"`
	TypeSliceInt64   []int64           `xconf:"type_slice_int64"`
	TypeBool         bool              `xconf:"type_bool"`
	MapRedis         map[string]*Redis `xconf:"map_redis"`
	// annotation@Redis(getter="RedisVisitor",deprecated="use MapRedis intead")
	Redis              *Redis                                                   `xconf:"redis,deprecated"` // 辅助指定类型为*Redis
	OnWatchError       WatchError                                               `xconf:"on_watch_error"`   // 辅助指定类型为WatchError
	OnWatchErrorNotNil func(loaderName string, confPath string, watchErr error) `xconf:"on_watch_error_not_nil"`
	TypeSliceDuratuon  []time.Duration                                          `xconf:"type_slice_duratuon"` // 辅助指定类型为WatchError
}

// ApplyOption apply mutiple new option and return the old ones
// sample:
// old := cc.ApplyOption(WithTimeout(time.Second))
// defer cc.ApplyOption(old...)
func (cc *XXXXXX) ApplyOption(opts ...XXXXXXOption) []XXXXXXOption {
	var previous []XXXXXXOption
	for _, opt := range opts {
		previous = append(previous, opt(cc))
	}
	return previous
}

// XXXXXXOption option func
type XXXXXXOption func(cc *XXXXXX) XXXXXXOption

// WithXXXXXXOptionUsage option func for filed OptionUsage
func WithXXXXXXOptionUsage(v string) XXXXXXOption {
	return func(cc *XXXXXX) XXXXXXOption {
		previous := cc.OptionUsage
		cc.OptionUsage = v
		return WithXXXXXXOptionUsage(previous)
	}
}

// WithXXXXXXEndpoints option func for filed Endpoints
func WithXXXXXXEndpoints(v ...string) XXXXXXOption {
	return func(cc *XXXXXX) XXXXXXOption {
		previous := cc.Endpoints
		cc.Endpoints = v
		return WithXXXXXXEndpoints(previous...)
	}
}

// WithXXXXXXReadTimeout option func for filed ReadTimeout
func WithXXXXXXReadTimeout(v time.Duration) XXXXXXOption {
	return func(cc *XXXXXX) XXXXXXOption {
		previous := cc.ReadTimeout
		cc.ReadTimeout = v
		return WithXXXXXXReadTimeout(previous)
	}
}

// WithXXXXXXTypeMapIntString option func for filed TypeMapIntString
func WithXXXXXXTypeMapIntString(v map[int]string) XXXXXXOption {
	return func(cc *XXXXXX) XXXXXXOption {
		previous := cc.TypeMapIntString
		cc.TypeMapIntString = v
		return WithXXXXXXTypeMapIntString(previous)
	}
}

// WithXXXXXXTypeSliceInt64 option func for filed TypeSliceInt64
func WithXXXXXXTypeSliceInt64(v ...int64) XXXXXXOption {
	return func(cc *XXXXXX) XXXXXXOption {
		previous := cc.TypeSliceInt64
		cc.TypeSliceInt64 = v
		return WithXXXXXXTypeSliceInt64(previous...)
	}
}

// WithXXXXXXTypeBool option func for filed TypeBool
func WithXXXXXXTypeBool(v bool) XXXXXXOption {
	return func(cc *XXXXXX) XXXXXXOption {
		previous := cc.TypeBool
		cc.TypeBool = v
		return WithXXXXXXTypeBool(previous)
	}
}

// WithXXXXXXMapRedis option func for filed MapRedis
func WithXXXXXXMapRedis(v map[string]*Redis) XXXXXXOption {
	return func(cc *XXXXXX) XXXXXXOption {
		previous := cc.MapRedis
		cc.MapRedis = v
		return WithXXXXXXMapRedis(previous)
	}
}

// WithXXXXXXRedis option func for filed Redis
//
// Deprecated: use MapRedis intead
func WithXXXXXXRedis(v *Redis) XXXXXXOption {
	return func(cc *XXXXXX) XXXXXXOption {
		previous := cc.Redis
		cc.Redis = v
		return WithXXXXXXRedis(previous)
	}
}

// WithXXXXXXOnWatchError option func for filed OnWatchError
func WithXXXXXXOnWatchError(v WatchError) XXXXXXOption {
	return func(cc *XXXXXX) XXXXXXOption {
		previous := cc.OnWatchError
		cc.OnWatchError = v
		return WithXXXXXXOnWatchError(previous)
	}
}

// WithXXXXXXOnWatchErrorNotNil option func for filed OnWatchErrorNotNil
func WithXXXXXXOnWatchErrorNotNil(v func(loaderName string, confPath string, watchErr error)) XXXXXXOption {
	return func(cc *XXXXXX) XXXXXXOption {
		previous := cc.OnWatchErrorNotNil
		cc.OnWatchErrorNotNil = v
		return WithXXXXXXOnWatchErrorNotNil(previous)
	}
}

// WithXXXXXXTypeSliceDuratuon option func for filed TypeSliceDuratuon
func WithXXXXXXTypeSliceDuratuon(v ...time.Duration) XXXXXXOption {
	return func(cc *XXXXXX) XXXXXXOption {
		previous := cc.TypeSliceDuratuon
		cc.TypeSliceDuratuon = v
		return WithXXXXXXTypeSliceDuratuon(previous...)
	}
}

// NewXXXXXX new XXXXXX
func NewXXXXXX(opts ...XXXXXXOption) *XXXXXX {
	cc := newDefaultXXXXXX()
	for _, opt := range opts {
		opt(cc)
	}
	if watchDogXXXXXX != nil {
		watchDogXXXXXX(cc)
	}
	return cc
}

// InstallXXXXXXWatchDog the installed func will called when NewXXXXXX  called
func InstallXXXXXXWatchDog(dog func(cc *XXXXXX)) { watchDogXXXXXX = dog }

// watchDogXXXXXX global watch dog
var watchDogXXXXXX func(cc *XXXXXX)

// newDefaultXXXXXX new default XXXXXX
func newDefaultXXXXXX() *XXXXXX {
	cc := &XXXXXX{}

	for _, opt := range [...]XXXXXXOption{
		WithXXXXXXOptionUsage(optionUsage),
		WithXXXXXXEndpoints([]string{"10.0.0.1", "10.0.0.2"}...),
		WithXXXXXXReadTimeout(time.Second),
		WithXXXXXXTypeMapIntString(map[int]string{1: "a", 2: "b"}),
		WithXXXXXXTypeSliceInt64([]int64{1, 2, 3, 4}...),
		WithXXXXXXTypeBool(false),
		WithXXXXXXMapRedis(map[string]*Redis{"test": NewRedis()}),
		WithXXXXXXRedis(NewRedis()),
		WithXXXXXXOnWatchError(nil),
		WithXXXXXXOnWatchErrorNotNil(func(loaderName string, confPath string, watchErr error) {
		}),
		WithXXXXXXTypeSliceDuratuon([]time.Duration{time.Second, time.Minute, time.Hour}...),
	} {
		opt(cc)
	}

	return cc
}

// AtomicSetFunc used for XConf
func (cc *XXXXXX) AtomicSetFunc() func(interface{}) { return AtomicXXXXXXSet }

// atomicXXXXXX global *XXXXXX holder
var atomicXXXXXX unsafe.Pointer

// onAtomicXXXXXXSet global call back when  AtomicXXXXXXSet called by XConf.
// use XXXXXXInterface.ApplyOption to modify the updated cc
// if passed in cc not valid, then return false, cc will not set to atomicXXXXXX
var onAtomicXXXXXXSet func(cc XXXXXXInterface) bool

// InstallCallbackOnAtomicXXXXXXSet install callback
func InstallCallbackOnAtomicXXXXXXSet(callback func(cc XXXXXXInterface) bool) {
	onAtomicXXXXXXSet = callback
}

// AtomicXXXXXXSet atomic setter for *XXXXXX
func AtomicXXXXXXSet(update interface{}) {
	cc := update.(*XXXXXX)
	if onAtomicXXXXXXSet != nil && !onAtomicXXXXXXSet(cc) {
		return
	}
	atomic.StorePointer(&atomicXXXXXX, (unsafe.Pointer)(cc))
}

// AtomicXXXXXX return atomic *XXXXXXVisitor
func AtomicXXXXXX() XXXXXXVisitor {
	current := (*XXXXXX)(atomic.LoadPointer(&atomicXXXXXX))
	if current == nil {
		defaultOne := newDefaultXXXXXX()
		if watchDogXXXXXX != nil {
			watchDogXXXXXX(defaultOne)
		}
		atomic.CompareAndSwapPointer(&atomicXXXXXX, nil, (unsafe.Pointer)(defaultOne))
		return (*XXXXXX)(atomic.LoadPointer(&atomicXXXXXX))
	}
	return current
}

// all getter func
func (cc *XXXXXX) GetOptionUsage() string              { return cc.OptionUsage }
func (cc *XXXXXX) GetEndpoints() []string              { return cc.Endpoints }
func (cc *XXXXXX) GetReadTimeout() time.Duration       { return cc.ReadTimeout }
func (cc *XXXXXX) GetTypeMapIntString() map[int]string { return cc.TypeMapIntString }
func (cc *XXXXXX) GetTypeSliceInt64() []int64          { return cc.TypeSliceInt64 }
func (cc *XXXXXX) GetTypeBool() bool                   { return cc.TypeBool }
func (cc *XXXXXX) GetMapRedis() map[string]*Redis      { return cc.MapRedis }

// GetRedis visitor func for filed Redis
//
// Deprecated: use MapRedis intead
func (cc *XXXXXX) GetRedis() RedisVisitor      { return cc.Redis }
func (cc *XXXXXX) GetOnWatchError() WatchError { return cc.OnWatchError }
func (cc *XXXXXX) GetOnWatchErrorNotNil() func(loaderName string, confPath string, watchErr error) {
	return cc.OnWatchErrorNotNil
}
func (cc *XXXXXX) GetTypeSliceDuratuon() []time.Duration { return cc.TypeSliceDuratuon }

// XXXXXXVisitor visitor interface for XXXXXX
type XXXXXXVisitor interface {
	GetOptionUsage() string
	GetEndpoints() []string
	GetReadTimeout() time.Duration
	GetTypeMapIntString() map[int]string
	GetTypeSliceInt64() []int64
	GetTypeBool() bool
	GetMapRedis() map[string]*Redis
	// GetRedis visitor func for filed Redis
	//
	// Deprecated: use MapRedis intead
	GetRedis() RedisVisitor
	GetOnWatchError() WatchError
	GetOnWatchErrorNotNil() func(loaderName string, confPath string, watchErr error)
	GetTypeSliceDuratuon() []time.Duration
}

// XXXXXXInterface visitor + ApplyOption interface for XXXXXX
type XXXXXXInterface interface {
	XXXXXXVisitor
	ApplyOption(...XXXXXXOption) []XXXXXXOption
}
