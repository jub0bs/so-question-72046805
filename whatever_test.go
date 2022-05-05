package whatever_test

import (
	"context"
	"errors"
	"testing"
	"whatever"
	"whatever/stub"

	"google.golang.org/api/iterator"
)

var dummyError = errors.New("oh no")

func TestCompute_Res(t *testing.T) {
	dummyContext := context.Background()
	dummyProject := "test"
	dummyRegion := "test"
	dummyVpc := "test"
	type TestCase struct {
	}
	testCases := []struct {
		desc   string
		name   string
		status string
		err    error
		want   error
	}{
		{
			desc:   "iterator done",
			name:   "dummyName",
			status: "dummyStatus",
			err:    iterator.Done,
			want:   nil,
		},
		{
			desc:   "error",
			name:   "dummyName",
			status: "dummyStatus",
			err:    dummyError,
			want:   dummyError,
		},
		{
			desc:   "not in use",
			name:   "dummyName",
			status: "NOT_IN_USE",
			err:    nil,
			want:   nil,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			cp := &whatever.Compute{
				Lister: stub.NewLister(tt.name, tt.status, tt.err),
			}
			if err := cp.Res(dummyContext, dummyProject, dummyRegion, dummyVpc); err != tt.want {
				t.Errorf("Res() error = %v, want %v", err, tt.want)
			}
		})
	}
}
