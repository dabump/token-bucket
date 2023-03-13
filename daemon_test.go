package tokenbucket

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func Test_flag_has(t *testing.T) {
	type args struct {
		flag flag
	}
	tests := []struct {
		name string
		a    flag
		args args
		want bool
	}{
		{
			name: "Test single flag",
			a:    Retryable,
			args: args{
				flag: Retryable,
			},
			want: true,
		},
		{
			name: "Test multiple flags",
			a:    Retryable | Forgiving,
			args: args{
				flag: Retryable,
			},
			want: true,
		},
		{
			name: "Test flag not found",
			a:    Retryable,
			args: args{
				flag: Forgiving,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.has(tt.args.flag); got != tt.want {
				t.Errorf("has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaemon_Hit(t *testing.T) {
	t.Run("Test happy day successful", func(t *testing.T) {
		bucket := NewBucket("Test Bucket", 2)
		dm := NewDaemon(bucket, NA)
		dm.Start()
		// drain the bucket
		assert.Assert(t, dm.Hit())
		assert.Assert(t, dm.Hit())
		// wait for it to fill
		time.Sleep(time.Duration(1) * time.Second)
		// drain the bucket again
		assert.Assert(t, dm.Hit())
		assert.Assert(t, dm.Hit())
		dm.Stop()
	})
	t.Run("Test drain bucket and fail to get token", func(t *testing.T) {
		bucket := NewBucket("Test Bucket", 2)
		dm := NewDaemon(bucket, NA)
		dm.Start()
		// drain the bucket
		assert.Assert(t, dm.Hit())
		assert.Assert(t, dm.Hit())
		assert.Assert(t, !dm.Hit())
		dm.Stop()
	})
	t.Run("Test forgiving", func(t *testing.T) {
		bucket := NewBucket("Test Bucket", 2)
		dm := NewDaemon(bucket, Forgiving)
		dm.Start()
		// drain the bucket
		assert.Assert(t, dm.Hit())
		assert.Assert(t, dm.Hit())
		assert.Assert(t, dm.Hit())
		dm.Stop()
	})
}
