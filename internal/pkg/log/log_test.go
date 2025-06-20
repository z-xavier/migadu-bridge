package log

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"migadu-bridge/internal/pkg/common"
)

func TestInfoConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 并发调用日志方法
			ctx := context.WithValue(context.Background(), common.XRequestIDKey, fmt.Sprintf("req-%d", id))
			l := C(ctx)
			l.WithField("goroutine", id).Info("concurrent log test")

			Infow("concurrent log test", "goroutine", id, common.XRequestIDKey, fmt.Sprintf("req-%d", id))
		}(i)
	}
	wg.Wait()
}
