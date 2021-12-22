package authentication

import (
	"context"
	"fmt"
	"strings"

	"github.com/motorsportgames/kartkraft-online-authentication/platform"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Authentication is the container for the credentials
type Authentication struct {
	Ticket   string
	Secret   string
	Platform platform.Platform
	ID       string
}

// AuthenticateContext takes a context, extracts the metadata containing the authentication credentials,
// and verifies it before returning an error or nil if the verification was successful
func AuthenticateContext(ctx context.Context) error {
	if md, ok := metadata.FromIncomingContext(ctx); ok {

		id := strings.Join(md["id"], "")

		if id == "" {
			return status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid ID: %s", id))
		}

		return nil
	}

	return fmt.Errorf("No credentials supplied")
}

// GetIDFromContext returns the ID from the context
func GetIDFromContext(ctx context.Context) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {

		// TODO: Add id to the authentication struct
		id := strings.Join(md["id"], "")

		if id == "" {
			return "", fmt.Errorf("Couldn't extract id from metadata after authentication. This should not occur")
		}

		return id, nil

	}

	return "", fmt.Errorf("No credentials supplied")
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

// GetRequestMetadata extracts embedded metadata from the request
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"ticket":   a.Ticket,
		"secret":   a.Secret,
		"platform": a.Platform.String(),
		"id":       a.ID,
	}, nil
}
