// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "fmt"

// SMBIOSSummary summarizes the System Management BIOS.
type SMBIOSSummary struct { //nolint:govet
	BIOS       BIOS
	System     System
	Enclosure  Enclosure
	Baseboard  Baseboard
	Processors ProcessorSummary
	Memory     MemorySummary
}

type BIOS struct {
	Vendor      *string
	Version     string
	ReleaseDate string
}

type System struct {
	Manufacturer string
	ProductName  string
	SerialNumber *string
	UUID         string
	SKUNumber    *string
	Family       *string
}

type Enclosure struct {
	Manufacturer string
	SerialNumber *string
}

type Baseboard struct {
	Manufacturer string
	ProductName  string
	SerialNumber string
	BoardType    string
}

type ProcessorSummary struct {
	ProcessorCount   int
	TotalCoreCount   int
	TotalThreadCount int
	Processors       []Processor
}

type Processor struct {
	Manufacturer *string
	ProductName  *string
	SerialNumber *string
	Speed        string
	CoreCount    int
	ThreadCount  int
}

type MemorySummary struct {
	ModuleCount int
	TotalSize   string
	Modules     []MemoryModule
}

type MemoryModule struct {
	Manufacturer *string
	ProductName  *string
	SerialNumber *string
	Type         string
	Size         string
	Speed        string
}

func (s *SMBIOS) Summarize() SMBIOSSummary {
	return SMBIOSSummary{
		BIOS{
			Vendor:      s.BIOSInformation.Vendor,
			Version:     s.BIOSInformation.Version,
			ReleaseDate: s.BIOSInformation.ReleaseDate,
		},
		System{
			Manufacturer: s.SystemInformation.Manufacturer,
			ProductName:  s.SystemInformation.ProductName,
			SerialNumber: s.SystemInformation.SerialNumber,
			UUID:         s.SystemInformation.UUID,
			SKUNumber:    s.SystemInformation.SKUNumber,
			Family:       s.SystemInformation.Family,
		},
		Enclosure{
			Manufacturer: s.SystemEnclosure.Manufacturer,
			SerialNumber: s.SystemEnclosure.SerialNumber,
		},
		Baseboard{
			Manufacturer: s.BaseboardInformation.Manufacturer,
			ProductName:  s.BaseboardInformation.Product,
			SerialNumber: s.BaseboardInformation.SerialNumber,
			BoardType:    s.BaseboardInformation.BoardType.String(),
		},
		MapProcessorSummary(s.ProcessorInformation),
		MapMemorySummary(s.MemoryDevices),
	}
}

func MapProcessorSummary(s []ProcessorInformation) ProcessorSummary {
	var processorCount = 0
	var totalCoreCount = 0
	var totalThreadCount = 0
	var processors []Processor
	for _, v := range s {
		if v.Status.SocketPopulated() {
			processorCount += 1
			totalCoreCount += int(v.CoreCount)
			totalThreadCount += int(v.ThreadCount)
			var processor = Processor{
				Manufacturer: v.ProcessorManufacturer,
				ProductName:  v.ProcessorVersion,
				SerialNumber: v.SerialNumber,
				Speed:        fmt.Sprintf("%d Mhz", v.CurrentSpeed),
				CoreCount:    int(v.CoreCount),
				ThreadCount:  int(v.ThreadCount),
			}
			processors = append(processors, processor)
		}
	}
	return ProcessorSummary{
		ProcessorCount:   processorCount,
		TotalCoreCount:   totalCoreCount,
		TotalThreadCount: totalThreadCount,
		Processors:       processors,
	}
}

func MapMemorySummary(s []MemoryDevice) MemorySummary {
	var moduleCount = 0
	var totalSize = 0
	var modules []MemoryModule
	for _, v := range s {
		if v.Size != 0 && v.Size != 0xFFFF {
			moduleCount += 1
			var size string
			if v.Size == 0x7FFF {
				totalSize += int(v.ExtendedSize)
				size = v.ExtendedSize.String()
			} else {
				totalSize += v.Size.Megabytes()
				size = v.Size.String()
			}
			var memoryModule = MemoryModule{
				Manufacturer: v.Manufacturer,
				ProductName:  v.PartNumber,
				SerialNumber: v.SerialNumber,
				Type:         v.MemoryType.String(),
				Size:         size,
				Speed:        v.Speed.String(),
			}
			modules = append(modules, memoryModule)
		}
	}
	return MemorySummary{
		ModuleCount: moduleCount,
		TotalSize:   fmt.Sprintf("%d GB", totalSize/1024),
		Modules:     modules,
	}
}
