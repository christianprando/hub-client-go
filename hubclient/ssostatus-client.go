// Copyright 2020 Synopsys, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hubclient

import "github.com/blackducksoftware/hub-client-go/hubapi"

func (c *Client) SsoStatus() (*hubapi.SsoStatus, error) {

	ssoStausURL := c.baseURL + "/api/sso/status"

	var ssoStatus hubapi.SsoStatus
	err := c.HttpGetJSON(ssoStausURL, &ssoStatus, 200)

	if err != nil {
		return nil, AnnotateHubClientError(err, "Error trying to get SSO status")
	}
	return &ssoStatus, nil
}
