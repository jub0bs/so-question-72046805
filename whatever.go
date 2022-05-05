package whatever

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	computev1 "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

type Iterator interface {
	Next() (*computev1.Address, error)
}

type Lister interface {
	List(ctx context.Context, req *computev1.ListAddressesRequest, opts ...gax.CallOption) Iterator
}

type RealLister struct {
	Client *compute.AddressesClient
}

func (rl *RealLister) List(ctx context.Context, req *computev1.ListAddressesRequest, opts ...gax.CallOption) Iterator {
	return rl.Client.List(ctx, req, opts...)
}

type Compute struct {
	Lister Lister // some appropriate interface type
}

func (cp *Compute) Res(ctx context.Context, project, region, vpc string) error {
	addrReq := &computev1.ListAddressesRequest{
		Project: project,
		Region:  region,
	}
	it := cp.Lister.List(ctx, addrReq)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		if *(resp.Status) != "IN_USE" {
			return ipConverter(*resp.Name, vpc)
		}
	}
	return nil
}

func ipConverter(name, vpc string) error {
	return nil
}
