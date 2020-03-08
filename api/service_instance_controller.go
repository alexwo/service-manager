/*
 * Copyright 2018 The Service Manager Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/pkg/web"
)

// ServiceInstanceController implements api.Controller by providing service Instances API logic
type ServiceInstanceController struct {
	*BaseController
}

func NewServiceInstanceController(ctx context.Context, options *Options) *ServiceInstanceController {
	return &ServiceInstanceController{
		BaseController: NewAsyncController(ctx, options, web.ServiceInstancesURL, types.ServiceInstanceType, true, func() types.Object {
			return &types.ServiceInstance{}
		}),
	}
}

func (c *ServiceInstanceController) Routes() []web.Route {
	return []web.Route{
		{
			Endpoint: web.Endpoint{
				Method: http.MethodOptions,
				Path:   c.resourceBaseURL,
			},
		},

		{
			Endpoint: web.Endpoint{
				Method: http.MethodOptions,
				Path:   fmt.Sprintf("%s/{%s}", c.resourceBaseURL, web.PathParamResourceID),
			},
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPost,
				Path:   c.resourceBaseURL,
			},
			Handler: c.CreateObject,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   fmt.Sprintf("%s/{%s}", c.resourceBaseURL, web.PathParamResourceID),
			},
			Handler: c.GetSingleObject,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   fmt.Sprintf("%s/{%s}%s/{%s}", c.resourceBaseURL, web.PathParamResourceID, web.OperationsURL, web.PathParamID),
			},
			Handler: c.GetOperation,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   c.resourceBaseURL,
			},
			Handler: c.ListObjects,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodDelete,
				Path:   fmt.Sprintf("%s/{%s}", c.resourceBaseURL, web.PathParamResourceID),
			},
			Handler: c.DeleteSingleObject,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPatch,
				Path:   fmt.Sprintf("%s/{%s}", c.resourceBaseURL, web.PathParamResourceID),
			},
			Handler: c.PatchObject,
		},
	}
}
