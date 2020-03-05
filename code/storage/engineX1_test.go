package storage

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/influxdata/influxdb/logger"
	"github.com/influxdata/influxdb/toml"
	"go.uber.org/zap"
)

func TestEngineX(t *testing.T) {
	t.Parallel()
	c := NewConfig()
	c.RetentionInterval = toml.Duration(time.Second)
	log := zap.NewNop()
	if testing.Verbose() {
		log = logger.New(os.Stdout)
	}

	t.Run("limiter", func(t *testing.T) {
		t.Parallel()

		path := MustTempDir()
		defer os.RemoveAll(path)

		var mu sync.Mutex
		limiter := func() func() {
			mu.Lock()
			return func() { mu.Unlock() }
		}

		engine1 := NewEngine(path, c, WithNodeID(2), WithEngineID(1), WithRetentionEnforcerLimiter(limiter))
		engine1.WithLogger(log)
		engine2 := NewEngine(path, c, WithNodeID(3), WithEngineID(2), WithRetentionEnforcerLimiter(limiter))
		engine2.WithLogger(log)

		var runner1, runner2 MockRunner
		engine1.retentionEnforcer = &runner1
		engine2.retentionEnforcer = &runner2

		var running int64
		errCh := make(chan error, 2)

		runner1.runf = func() {
			x := atomic.AddInt64(&running, 1)
			if x > 1 {
				errCh <- errors.New("runner 1 ran concurrently with runner 2")
				return
			}

			time.Sleep(time.Second) //Running retention

			atomic.AddInt64(&running, -1)
			runner1.runf = func() {} // Don't run again.
			errCh <- nil
		}

		runner2.runf = func() {
			x := atomic.AddInt64(&running, 1)
			if x > 1 {
				errCh <- errors.New("runner 2 ran concurrently with runner 1")
				return
			}

			time.Sleep(time.Second) //Running retention

			atomic.AddInt64(&running, -1)
			runner2.runf = func() {} // Don't run again.
			errCh <- nil
		}

		if err := engine1.Open(context.Background()); err != nil {
			t.Fatal(err)
		} else if err := engine2.Open(context.Background()); err != nil {
			t.Fatal(err)
		}
		defer engine1.Close()
		defer engine2.Close()

		for i := 0; i < 2; i++ {
			if err := <-errCh; err != nil {
				t.Fatal(err)
			}
		}
	})
}

type MockRunner struct {
	runf func()
}

func (r *MockRunner) run() {
	if r.runf == nil {
		return
	}
	r.runf()
}

func MustTempDir() string {
	dir, err := ioutil.TempDir("", "storage-engine-test")
	if err != nil {
		panic(fmt.Sprintf("failed to create temp dir: %v", err))
	}
	return dir
}
