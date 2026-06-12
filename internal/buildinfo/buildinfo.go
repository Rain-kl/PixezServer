// Copyright 2026 Arctel.net
// SPDX-License-Identifier: AGPL-3.0-only

// Package buildinfo exposes metadata injected by the release workflow.
package buildinfo

var (
	// Version is the application version.
	Version = "dev"
	// BuildTime is the UTC release build timestamp.
	BuildTime = ""
)
