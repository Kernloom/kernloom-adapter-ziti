// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2026 Kernloom Contributors

package adapter

import (
	"context"

	adapterv1 "github.com/kernloom/kernloom-protocol/sdk/go/adapter/v1"
)

type Adapter struct{}

func New() *Adapter {
	return &Adapter{}
}

func (a *Adapter) Describe(context.Context) (*adapterv1.AdapterDescriptor, error) {
	return &adapterv1.AdapterDescriptor{
		AdapterId:       "kernloom.adapter.ziti",
		Name:            "Kernloom OpenZiti Adapter",
		ProtocolVersion: adapterv1.ProtocolVersion,
		Capabilities: []*adapterv1.CapabilityDescriptor{
			{
				Id:          "ziti.observed_state.read",
				DisplayName: "Read OpenZiti observed state",
				Kind:        "observed_state",
				Actions:     []string{"read_services", "read_identities", "read_service_policies"},
			},
			{
				Id:          "ziti.config.plan",
				DisplayName: "Plan OpenZiti config proposals",
				Kind:        "config_planner",
				Actions:     []string{"plan_service_policy"},
			},
		},
		ContextRequirements: []*adapterv1.ContextRequirementDescriptor{
			{
				Fact:        "ziti.controller.reachable",
				Freshness:   "1m",
				Confidence:  "high",
				Sensitivity: "operational",
			},
		},
		Privileges: []*adapterv1.PrivilegeDescriptor{
			{
				Id:     "ziti.readonly.observation",
				Reason: "Read services, identities and service policies for planning and conformance.",
				Scope:  "ziti_controller",
				Access: "read",
			},
		},
		Facets: []string{
			adapterv1.FacetDescribe,
			adapterv1.FacetHealth,
			adapterv1.FacetReadObservedState,
			adapterv1.FacetPlanConfig,
			adapterv1.FacetValidateConfig,
			adapterv1.FacetProvideRelationships,
			adapterv1.FacetProvideConformanceEvidence,
		},
	}, nil
}

func (a *Adapter) Health(context.Context) (*adapterv1.HealthResponse, error) {
	return &adapterv1.HealthResponse{
		Status:  adapterv1.HealthServing,
		Message: "ziti adapter bootstrap is serving",
	}, nil
}
