//go:build windows

// Copyright 2026 Arctel.net
// SPDX-License-Identifier: AGPL-3.0-only

package updater

import "errors"

func replaceAndRestart(_, _ string) error {
	return errors.New(errAutomaticUpgradeBlocked)
}
