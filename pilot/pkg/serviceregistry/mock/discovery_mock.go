// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mock

import (
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/config/host"
)

var (
	// HelloService is a mock service with `hello.default.svc.cluster.local` as
	// a hostname and `10.1.0.0` for ip
	HelloService = MakeService("hello.default.svc.cluster.local", "10.1.0.0", []string{}, "cluster-1")

	// ReplicatedFooServiceName is a service replicated in all clusters.
	ReplicatedFooServiceName = host.Name("foo.default.svc.cluster.local")
	ReplicatedFooServiceV1   = MakeService(ReplicatedFooServiceName, "10.3.0.0", []string{
		"spiffe://cluster.local/ns/default/sa/foo1",
		"spiffe://cluster.local/ns/default/sa/foo-share",
	}, "")
	ReplicatedFooServiceV2 = MakeService(ReplicatedFooServiceName, "10.3.0.1", []string{
		"spiffe://cluster.local/ns/default/sa/foo2",
		"spiffe://cluster.local/ns/default/sa/foo-share",
	}, "")

	// WorldService is a mock service with `world.default.svc.cluster.local` as
	// a hostname and `10.2.0.0` for ip
	WorldService = MakeService("world.default.svc.cluster.local", "10.2.0.0", []string{
		"spiffe://cluster.local/ns/default/sa/world1",
		"spiffe://cluster.local/ns/default/sa/world2",
	}, "cluster-2")

	// ExtHTTPService is a mock external HTTP service
	ExtHTTPService = MakeExternalHTTPService("httpbin.default.svc.cluster.local",
		true, "")

	// ExtHTTPSService is a mock external HTTPS service
	ExtHTTPSService = MakeExternalHTTPSService("httpsbin.default.svc.cluster.local",
		true, "")

	// HelloInstanceV0 is a mock IP address for v0 of HelloService
	HelloInstanceV0 = MakeIP(HelloService, 0)

	// HelloProxyV0 is a mock proxy v0 of HelloService
	HelloProxyV0 = model.Proxy{
		Type:         model.SidecarProxy,
		IPAddresses:  []string{HelloInstanceV0},
		ID:           "v0.default",
		DNSDomain:    "default.svc.cluster.local",
		IstioVersion: model.MaxIstioVersion,
		Metadata:     &model.NodeMetadata{},
	}

	// MockDiscovery is an in-memory ServiceDiscover with mock services
	MockDiscovery = &ServiceDiscovery{
		services: map[host.Name]*model.Service{
			HelloService.ClusterLocal.Hostname:   HelloService,
			WorldService.ClusterLocal.Hostname:   WorldService,
			ExtHTTPService.ClusterLocal.Hostname: ExtHTTPService,
			// TODO external https is not currently supported - this service
			// should NOT be in any of the .golden json files
			ExtHTTPSService.ClusterLocal.Hostname: ExtHTTPSService,
		},
		versions: 2,
	}
)
