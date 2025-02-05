/*
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

package client

import (
	"encoding/json"
	"net"

	"github.com/apache/trafficcontrol/lib/go-tc"
)

const API_V13_SERVERCHECK = "/api/1.3/servercheck"
const API_V13_SERVERCHECK_GET = "/api/1.3/servers/checks"

// Update a Server Check Status
func (to *Session) UpdateCheckStatus(status tc.ServercheckPostNullable) (*tc.ServercheckPostResponse, ReqInf, error) {
	uri := API_V13_SERVERCHECK
	var remoteAddr net.Addr
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	jsonReq, err := json.Marshal(status)
	if err != nil {
		return nil, reqInf, err
	}
	resp := tc.ServercheckPostResponse{}
	reqInf, err = post(to, uri, jsonReq, &resp)
	if err != nil {
		return nil, reqInf, err
	}
	return &resp, reqInf, nil
}

// Get Server Check Data
func (to *Session) GetCheckData() (*tc.ServerchecksResponse, ReqInf, error) {
	uri := API_V13_SERVERCHECK_GET
	var remoteAddr net.Addr
	reqInf := ReqInf{CacheHitStatus: CacheHitStatusMiss, RemoteAddr: remoteAddr}
	resp := tc.ServerchecksResponse{}
	reqInf, err := get(to, uri, &resp)
	if err != nil {
		return nil, reqInf, err
	}
	return &resp, reqInf, nil
}
