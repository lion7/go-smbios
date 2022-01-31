// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// GroupAssociations represents the SMBIOS group associations.
type GroupAssociations struct {
	// GroupName returns the group name.
	GroupName string
}

func NewGroupAssociations(s *smbios.Structure) *GroupAssociations {
	return &GroupAssociations{
		GroupName: *GetString(s, 0x04),
	}
}
