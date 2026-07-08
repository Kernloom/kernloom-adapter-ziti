// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2026 Kernloom Contributors

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/kernloom/kernloom-adapter-ziti/internal/adapter"
	adapterv1 "github.com/kernloom/kernloom-protocol/sdk/go/adapter/v1"
	"google.golang.org/grpc"
)

var logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{}))

func main() {
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		serve(os.Args[2:])
		return
	}
	describe()
}

func describe() {
	desc, err := adapter.New().Descriptor(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := json.NewEncoder(os.Stdout).Encode(desc); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func serve(args []string) {
	fs := flag.NewFlagSet("kernloom-adapter-ziti serve", flag.ExitOnError)
	addr := fs.String("addr", "127.0.0.1:18081", "gRPC listen address")
	manifestDigest := fs.String("manifest-digest", os.Getenv("KERNLOOM_ADAPTER_MANIFEST_DIGEST"), "sha256 digest of the adapter manifest reported by Describe")
	if err := fs.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		logger.Error("ziti_adapter_listen_failed", "addr", *addr, "error", err.Error())
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	server := grpc.NewServer()
	adapterv1.RegisterAdapterServiceServer(server, adapter.NewWithManifestDigest(*manifestDigest))
	logger.Info("adapter_server_starting", "adapter_id", "kernloom.adapter.ziti", "addr", *addr)
	if err := server.Serve(listener); err != nil {
		logger.Error("adapter_server_failed", "adapter_id", "kernloom.adapter.ziti", "addr", *addr, "error", err.Error())
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
