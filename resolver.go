package fipsresolver

import (
	"github.com/aws/aws-sdk-go/aws/endpoints"
)

// FIPSEndpointResolver is an endpoint resolver which
// resolves exclusively to FIPS regional endpoints
type FIPSEndpointResolver struct {
	// TODO
}

// NewFIPSEndpointResolver is the constructor for a FIPSEndpointResolver
func NewFIPSEndpointResolver() endpoints.Resolver {
	return &FIPSEndpointResolver{
		// TODO
	}
}

// EndpointFor resolves for a given service, region, and option inputs
func (r *FIPSEndpointResolver) EndpointFor(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
	// TODO
	return endpoints.ResolvedEndpoint{}, nil
}
