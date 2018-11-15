// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/client-go/config/clientset/versioned/typed/config/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeConfigV1 struct {
	*testing.Fake
}

func (c *FakeConfigV1) Authentications() v1.AuthenticationInterface {
	return &FakeAuthentications{c}
}

func (c *FakeConfigV1) Builds() v1.BuildInterface {
	return &FakeBuilds{c}
}

func (c *FakeConfigV1) ClusterOperators() v1.ClusterOperatorInterface {
	return &FakeClusterOperators{c}
}

func (c *FakeConfigV1) ClusterVersions() v1.ClusterVersionInterface {
	return &FakeClusterVersions{c}
}

func (c *FakeConfigV1) Consoles() v1.ConsoleInterface {
	return &FakeConsoles{c}
}

func (c *FakeConfigV1) DNSs() v1.DNSInterface {
	return &FakeDNSs{c}
}

func (c *FakeConfigV1) IdentityProviders() v1.IdentityProviderInterface {
	return &FakeIdentityProviders{c}
}

func (c *FakeConfigV1) Images() v1.ImageInterface {
	return &FakeImages{c}
}

func (c *FakeConfigV1) Infrastructures() v1.InfrastructureInterface {
	return &FakeInfrastructures{c}
}

func (c *FakeConfigV1) Ingresses() v1.IngressInterface {
	return &FakeIngresses{c}
}

func (c *FakeConfigV1) Networks() v1.NetworkInterface {
	return &FakeNetworks{c}
}

func (c *FakeConfigV1) OAuths() v1.OAuthInterface {
	return &FakeOAuths{c}
}

func (c *FakeConfigV1) Projects() v1.ProjectInterface {
	return &FakeProjects{c}
}

func (c *FakeConfigV1) Schedulings() v1.SchedulingInterface {
	return &FakeSchedulings{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeConfigV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
