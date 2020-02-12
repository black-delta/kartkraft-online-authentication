package authentication

import (
	"context"
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestAuthentication(test *testing.T) {

	md := metadata.New(map[string]string{
		"ticket":   "CF5B949877FF52592DC8ACBB62AF3F4B",
		"secret":   "DKLFSHFLSHFFKJDHFJHUIGYRGEIWRYIUWEFBVIDBHDJSGF",
		"platform": "0",
		"id":       "FDHJKNERQIO3132",
	})

	ctx := context.Background()
	ctx = metadata.NewIncomingContext(ctx, md)

	err := AuthenticateContext(ctx)

	if err != nil {
		test.Errorf("Error: %s", err)
		test.Fail()
	}

}
