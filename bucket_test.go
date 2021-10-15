package tokenbucket

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestNewBucket(t *testing.T) {
	type args struct {
		designation string
		size        int16
	}
	tests := []struct {
		name string
		args args
		want *Bucket
	}{
		{
			name: "Test bucket creation",
			args: args{
				designation: "Test Bucket",
				size:        15,
			},
			want: &Bucket{
				size:                15,
				designation:         "Test Bucket",
				availableTokens:     15,
				lastAvailableTokens: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBucket(tt.args.designation, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucket_hit(t *testing.T) {
	tests := []struct {
		name   string
		bucket *Bucket
		want   bool
	}{
		{
			name:   "Successful token assignment",
			bucket: NewBucket("Test Bucket", 1),
			want:   true,
		},
		{
			name:   "UnSuccessful token assignment",
			bucket: NewBucket("Test Bucket", 0),
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bucket.hit(); got != tt.want {
				t.Errorf("hit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucket_fill(t *testing.T) {
	t.Run("Test refilled", func(t *testing.T) {
		b := &Bucket{
			size:                15,
			designation:         "Test Bucket",
			availableTokens:     0,
			lastAvailableTokens: 0,
		}
		b.fill()
		assert.Assert(t, b.availableTokens == b.size)
	})
}