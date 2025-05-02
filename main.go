// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2025-Present Harry Randazzo

// Package main is the entrypoint for the sizeoci CLI
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"slices"
	"strings"
	"syscall"

	"github.com/defenseunicorns/pkg/oci"
	"github.com/dustin/go-humanize"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

func main() {
	var platform string
	flag.StringVar(&platform, "p", runtime.GOOS, "request platform in the form of os[/arch][/variant][:os_version]")

	var help bool
	flag.BoolVar(&help, "h", false, "Print this message and exit.")

	var ver bool
	flag.BoolVar(&ver, "v", false, "Print the version number of sizeoci and exit.")

	flag.Parse()

	if help || slices.Contains(os.Args[1:], "--help") {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if ver {
		bi, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("version information not available")
			os.Exit(1)
		}
		fmt.Println(bi.Main.Version)
		os.Exit(0)
	}

	if len(flag.Args()) > 1 || len(flag.Args()) == 0 {
		fatal(fmt.Errorf("invalid number of args: want 1, got %d", len(flag.Args())))
	}

	targetPlatform, err := parsePlatform(platform)
	if err != nil {
		fatal(err)
	}

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	distributionReference := flag.Args()[0]
	client, err := oci.NewOrasRemote(distributionReference, targetPlatform)
	if err != nil {
		fatal(err)
	}

	root, err := client.FetchRoot(ctx)
	if err != nil {
		fatal(err)
	}

	total := oci.SumDescsSize(root.Layers)

	fmt.Println(humanize.Bytes(uint64(total)))
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

// https://github.com/oras-project/oras/blob/main/cmd/oras/internal/option/platform.go#L43
func parsePlatform(in string) (ocispec.Platform, error) {
	if in == "" {
		return ocispec.Platform{}, nil
	}

	// OS[/Arch[/Variant]][:OSVersion]
	// If Arch is not provided, will use GOARCH instead
	var platformStr string
	var p ocispec.Platform
	platformStr, p.OSVersion, _ = strings.Cut(in, ":")
	parts := strings.Split(platformStr, "/")
	switch len(parts) {
	case 3:
		p.Variant = parts[2]
		fallthrough
	case 2:
		p.Architecture = parts[1]
	case 1:
		p.Architecture = runtime.GOARCH
	default:
		return ocispec.Platform{}, fmt.Errorf("failed to parse platform %q: expected format os[/arch[/variant]]", in)
	}
	p.OS = parts[0]
	if p.OS == "" {
		return ocispec.Platform{}, fmt.Errorf("invalid platform: OS cannot be empty")
	}
	if p.Architecture == "" {
		return ocispec.Platform{}, fmt.Errorf("invalid platform: Architecture cannot be empty")
	}

	return p, nil
}
