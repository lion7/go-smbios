// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// BIOSInformation represents the BIOS information.
type BIOSInformation struct {
	// Vendor returns the BIOS vendor.
	Vendor *string
	// Version returns the BIOS version.
	Version string
	// ReleaseDate returns the BIOS release date.
	ReleaseDate string
}

func NewBIOSInformation(s *smbios.Structure) *BIOSInformation {
	return &BIOSInformation{
		GetString(s, 0x04),
		*GetString(s, 0x05),
		*GetString(s, 0x08),
	}
}
