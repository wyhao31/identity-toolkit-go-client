// Copyright 2014 Google Inc. All Rights Reserved.
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

// +build appengine

package gitkit

import (
	"net/http"

	"google.golang.org/appengine/urlfetch"

	"golang.org/x/net/context"
)

// defaultTransport returns a urlfetcher HTTP transport in AppEngine.
func defaultTransport(ctx context.Context) http.RoundTripper {
	return urlfetch.Client(ctx).Transport
}

// apiClient creates a new APIClient based on the current context.
func (c *Client) apiClient(ctx context.Context) *APIClient {
	// newAPIClient should never return error on App Engine.
	api, _ := newAPIClient(ctx, c.jc)
	return api
}
