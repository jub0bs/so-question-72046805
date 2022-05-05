package stub

import (
	"context"
	"whatever"

	gax "github.com/googleapis/gax-go/v2"
	computev1 "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

type Lister func(ctx context.Context, req *computev1.ListAddressesRequest, opts ...gax.CallOption) whatever.Iterator

func (l Lister) List(ctx context.Context, req *computev1.ListAddressesRequest, opts ...gax.CallOption) whatever.Iterator {
	return l(ctx, req, opts...)
}

var _ whatever.Lister = (*Lister)(nil)

type Iterator struct {
	name   string
	status string
	err    error
}

func (i *Iterator) Next() (*computev1.Address, error) {
	addr := computev1.Address{
		Name:   &i.name,
		Status: &i.status,
	}
	return &addr, i.err
}

var _ whatever.Iterator = (*Iterator)(nil)

func NewLister(name, status string, err error) whatever.Lister {
	return Lister(func(_ context.Context, _ *computev1.ListAddressesRequest, _ ...gax.CallOption) whatever.Iterator {
		return &Iterator{
			status: status,
			name:   name,
			err:    err,
		}
	})
}
