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

package health

import (
	"fmt"
	"github.com/InVisionApp/go-health"
)

// Settings type to be loaded from the environment
type Settings struct {
	FailuresTreshold int64 `mapstructure:"failures_treshold" description:"maximum failures in a row until component is considered down"`
	Interval         int64 `description:"seconds between health checks of components"`
}

// DefaultSettings returns default values for health settings
func DefaultSettings() *Settings {
	return &Settings{
		FailuresTreshold: 3,
		Interval:         60,
	}
}

// Validate validates the health settings
func (s *Settings) Validate() error {
	if s.FailuresTreshold < 0 {
		return fmt.Errorf("validate Settings: FailuresTreshold must be >= 0")
	}
	if s.Interval < 0 {
		return fmt.Errorf("valudate Settings: Interval must be >= 0")
	}
	return nil
}

// Status represents the overall health status of a component
type Status string

const (
	// StatusUp indicates that the checked component is up and running
	StatusUp Status = "UP"
	// StatusDown indicates the the checked component has an issue and is unavailable
	StatusDown Status = "DOWN"
	// StatusUnknown indicates that the health of the checked component cannot be determined
	StatusUnknown Status = "UNKNOWN"
)

// Health contains information about the health of a component.
type Health struct {
	Status  Status                 `json:"status"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// New returns a new Health with an unknown status an empty details.
func New() *Health {
	return &Health{
		Status:  StatusUnknown,
		Details: make(map[string]interface{}),
	}
}

// WithStatus sets the status of the health
func (h *Health) WithStatus(status Status) *Health {
	h.Status = status
	return h
}

// WithError sets the status of the health to DOWN and adds an error detail
func (h *Health) WithError(err error) *Health {
	h.Status = StatusDown
	return h.WithDetail("error", err)
}

// WithDetail adds a detail to the health
func (h *Health) WithDetail(key string, val interface{}) *Health {
	h.Details[key] = val
	return h
}

// Up sets the health status to up
func (h *Health) Up() *Health {
	h.Status = StatusUp
	return h
}

// Down sets the health status to down
func (h *Health) Down() *Health {
	h.Status = StatusDown
	return h
}

// Unknown sets the health status to unknown
func (h *Health) Unknown() *Health {
	h.Status = StatusUnknown
	return h
}

// WithDetails adds the given details to the health
func (h *Health) WithDetails(details map[string]interface{}) *Health {
	for k, v := range details {
		h.Details[k] = v
	}
	return h
}

// Indicator is an interface to provide the health of a component
type Indicator interface {
	// Name returns the name of the component
	Name() string

	// Status returns the health information of the component
	Status() (interface{}, error)
}

// AggregationPolicy is an interface to provide aggregated health information
//go:generate counterfeiter . AggregationPolicy
type AggregationPolicy interface {
	// Apply processes the given healths to build a single health
	Apply(healths map[string]health.State, failureTreshold int64) *Health
}

// NewDefaultRegistry returns a default health registry with a single ping indicator and a default aggregation policy
func NewDefaultRegistry() *Registry {
	return &Registry{
		HealthIndicators:        []Indicator{&pingIndicator{}},
		HealthAggregationPolicy: &DefaultAggregationPolicy{},
	}
}

// Registry is an interface to store and fetch health indicators
type Registry struct {
	// HealthIndicators are the currently registered health indicators
	HealthIndicators []Indicator

	// HealthAggregationPolicy is the registered health aggregationPolicy
	HealthAggregationPolicy AggregationPolicy
}
