package main

import (
	"context"
	"log"
	"whatever"

	compute "cloud.google.com/go/compute/apiv1"
	"google.golang.org/api/option"
)

func main() {
	secret := "don't hardcode me"
	ctx, cancel := context.WithCancel(context.Background()) // for instance
	defer cancel()
	c, err := compute.NewAddressesRESTClient(ctx, option.WithCredentialsFile(secret))
	if err != nil {
		log.Fatal(err) // or deal with the error in some way
	}
	defer c.Close()
	cp := whatever.Compute{Lister: &whatever.RealLister{Client: c}}
	if err := cp.Res(ctx, "my-project", "us-east-1", "my-vpc"); err != nil {
		log.Fatal(err) // or deal with the error in some way
	}
}
