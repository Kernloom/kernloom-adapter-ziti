// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2026 Kernloom Contributors

package adapter

import (
	"context"
	"strings"

	adapterv1 "github.com/kernloom/kernloom-protocol/sdk/go/adapter/v1"
)

type Adapter struct {
	adapterv1.UnimplementedAdapterServiceServer
	manifestDigest string
}

func New() *Adapter {
	return NewWithManifestDigest("")
}

func NewWithManifestDigest(manifestDigest string) *Adapter {
	return &Adapter{manifestDigest: strings.TrimSpace(manifestDigest)}
}

func (a *Adapter) Descriptor(context.Context) (*adapterv1.AdapterDescriptor, error) {
	return &adapterv1.AdapterDescriptor{
		AdapterId:       "kernloom.adapter.ziti",
		Name:            "Kernloom OpenZiti Adapter",
		ProtocolVersion: adapterv1.ProtocolVersion,
		ManifestDigest:  a.manifestDigest,
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
		FacetDescriptors: []*adapterv1.FacetDescriptor{
			{Name: adapterv1.FacetDescribe, Status: adapterv1.FacetStatusImplemented},
			{Name: adapterv1.FacetHealth, Status: adapterv1.FacetStatusImplemented},
			{Name: adapterv1.FacetReadObservedState, Status: adapterv1.FacetStatusPlanned, Message: "OpenZiti observed-state reads are planned after Slice 2."},
			{Name: adapterv1.FacetPlanConfig, Status: adapterv1.FacetStatusPlanned, Message: "OpenZiti config planning is planned after Slice 2."},
			{Name: adapterv1.FacetValidateConfig, Status: adapterv1.FacetStatusPlanned, Message: "OpenZiti config validation is planned after Slice 2."},
			{Name: adapterv1.FacetProvideRelationships, Status: adapterv1.FacetStatusPlanned, Message: "OpenZiti relationship reads are planned after Slice 2."},
			{Name: adapterv1.FacetProvideConformanceEvidence, Status: adapterv1.FacetStatusPlanned, Message: "OpenZiti conformance evidence is planned after Slice 2."},
		},
	}, nil
}

func (a *Adapter) Describe(ctx context.Context, _ *adapterv1.DescribeRequest) (*adapterv1.DescribeResponse, error) {
	desc, err := a.Descriptor(ctx)
	if err != nil {
		return nil, err
	}
	return &adapterv1.DescribeResponse{Adapter: desc}, nil
}

func (a *Adapter) Health(context.Context, *adapterv1.HealthRequest) (*adapterv1.HealthResponse, error) {
	return &adapterv1.HealthResponse{
		Status:  adapterv1.HealthServing,
		Message: "ziti adapter bootstrap is serving",
	}, nil
}
