// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2026 Kernloom Contributors

package adapter

import (
	"context"
	"testing"

	contractv1 "github.com/kernloom/kernloom-protocol/contract/adapter/v1"
)

func TestAdapterPassesServiceContract(t *testing.T) {
	contractv1.RunServiceContract(t, New())
}

func TestDescriptorReportsManifestDigest(t *testing.T) {
	desc, err := NewWithManifestDigest(" sha256:ziti-manifest ").Descriptor(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if desc.GetManifestDigest() != "sha256:ziti-manifest" {
		t.Fatalf("expected manifest digest to be reported, got %q", desc.GetManifestDigest())
	}
}
