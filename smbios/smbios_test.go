// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/talos-systems/go-smbios/smbios"
)

func TestASRockSingleRyzen(t *testing.T) {
	DoTestDesktopManagementInterface(t, "ASRock-Single-Ryzen")
}

func TestDellPowerEdgeR630DualXeon(t *testing.T) {
	DoTestDesktopManagementInterface(t, "Dell-PowerEdge-R630-Dual-Xeon")
}

func TestDellSuperMicroDualXeon(t *testing.T) {
	DoTestDesktopManagementInterface(t, "SuperMicro-Dual-Xeon")
}

func TestDellSuperMicroQuadOpteron(t *testing.T) {
	DoTestDesktopManagementInterface(t, "SuperMicro-Quad-Opteron")
}

func DoTestDesktopManagementInterface(t *testing.T, name string) {
	stream, err := os.Open("../test/" + name + ".dmi")
	require.NoError(t, err)

	//nolint: errcheck
	defer stream.Close()

	version := smbios.Version{Major: 3, Minor: 3, Revision: 0} // dummy version
	actual, err := smbios.Decode(stream, version)
	require.NoError(t, err)

	actualJson, err := json.MarshalIndent(actual, "", "\t")
	require.NoError(t, err)

	summary := actual.Summarize()
	actualSummaryJson, err := json.MarshalIndent(summary, "", "\t")
	require.NoError(t, err)

	//os.WriteFile("../test/" + name + ".json", actualJson, 0777)
	//os.WriteFile("../test/" + name + "-summary.json", actualSummaryJson, 0777)

	expectedJson, err := os.ReadFile("../test/" + name + ".json")
	require.NoError(t, err)
	require.Exactly(t, string(expectedJson), string(actualJson))

	expectedSummaryJson, err := os.ReadFile("../test/" + name + "-summary.json")
	require.NoError(t, err)
	require.Exactly(t, string(expectedSummaryJson), string(actualSummaryJson))
}
