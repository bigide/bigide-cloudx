// Copyright 2019 Yunion
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

package cloudproxy

import (
	"yunion.io/x/onecloud/pkg/apis"
)

type ForwardCreateInput struct {
	apis.VirtualResourceCreateInput

	ProxyEndpointId string
	ProxyAgentId    string
	Type            string
	BindPortReq     int `json:",omitzero"`
	RemoteAddr      string
	RemotePort      string

	LastSeenTimeout int `json:",omitzero"`

	Opaque string
}

type ForwardCreateFromServerInput struct {
	ServerId string

	Type        string
	BindPortReq int `json:",omitzero"`
	RemotePort  string

	LastSeenTimeout int `json:",omitzero"`
}

type ForwardHeartbeatInput struct{}

type ForwardListInput struct {
	ProxyAgentId    string
	ProxyEndpointId string

	Type string

	Opaque string
}
