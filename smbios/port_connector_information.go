// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// PortConnectorInformation represents the port connector information.
type PortConnectorInformation struct {
	// InternalReferenceDesignator returns the port connector internal reference designator.
	InternalReferenceDesignator *string
	// ExternalReferenceDesignator returns the port connector external reference designator.
	ExternalReferenceDesignator *string
}

func NewPortConnectorInformation(s *smbios.Structure) *PortConnectorInformation {
	return &PortConnectorInformation{
		InternalReferenceDesignator: GetString(s, 0x04),
		ExternalReferenceDesignator: GetString(s, 0x06),
	}
}
