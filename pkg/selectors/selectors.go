/*
Copyright The Kubernetes NMState Authors.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package selectors

import (
	"github.com/go-logr/logr"

	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	nmstatev1 "github.com/nmstate/kubernetes-nmstate/api/v1"
)

type Selectors struct {
	client client.Client
	policy nmstatev1.NodeNetworkConfigurationPolicy
	logger logr.Logger
}

func NewFromPolicy(client client.Client, policy nmstatev1.NodeNetworkConfigurationPolicy) Selectors {
	selectors := Selectors{
		client: client,
		policy: policy,
	}
	selectors.logger = logf.Log.WithName("policy/selectors").WithValues("policy", policy.Name)
	return selectors
}
