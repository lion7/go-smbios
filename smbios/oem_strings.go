// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// OEMStrings represents the SMBIOS OEM strings.
type OEMStrings struct {
	// Count return the number of strings
	Count uint8
	// Strings returns the actual strings.
	Strings []string
}

func NewOEMStrings(s *smbios.Structure) *OEMStrings {
	return &OEMStrings{
		Count:   GetByte(s, 0x04),
		Strings: s.Strings,
	}
}
