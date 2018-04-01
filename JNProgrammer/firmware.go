//JNProgrammer;JN516x firmware programmer written in go
//Copyright (C) HanWool Lee <kudnya@gmail.com>
//
//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
//along with this program.  If not, see <http://www.gnu.org/licenses/>.

package JNProgrammer

import (
	"os"
	"encoding/binary"
)

const (
	E_FW_OK = iota
	E_FW_ERROR
	E_FW_INVALID_FILE
)

type ImageVersion uint8
const (
	E_IMAGE_VERSION_1 ImageVersion = 1
	E_IMAGE_VERSION_2 ImageVersion = 2

)

type FirmwareInfo struct {
	ImageVersion ImageVersion

	ROMVersion uint32
	TextSectionLoadAddress uint32
	TextSectionLength uint32
	BssSectionLoadAddress uint32
	BssSectionLength uint32
	WakeUpEntryPoint uint32
	ResetEntryPoint uint32

	ImageDataV1 [2]uint8
	ImageDataV2 [4]uint32
	ImageLength uint32

	MacAddressLocation uint32
	TextData []uint8
}

type binHeaderV1 struct {
	ConfigByte0 uint8
	ConfigByte1 uint8
	SpiScrambleIndex uint16
	TextStartAddress uint32
	TextLength uint32
	ROMVersion uint32
	Unused1 uint32
	BssStartAddress uint32
	BssLength uint32
	WakeUpEntryPoint uint32
	ResetEntryPoint uint32
	OadData [12]uint8
	TextDataStart uint8
}

type binHeaderV2 struct {
	ROMVersion uint32
	BootImageRecord [4]uint32
	MacAddress uint64
	AEncryptionInitialisationVector [4]uint32
	DataSectionInfo uint32
	BssSectionInfo uint32
	WakeUpEntryPoint uint32
	ResetEntryPoint uint32
	TextDataStart uint8
}

func FwOpen(fwInfo *FirmwareInfo, firmwareFileName string) int {
	file, err := os.Open(firmwareFileName)
	if err != nil {
		Logger.Printf("Can not open firmware file %v\n", err)
		return E_FW_INVALID_FILE
	}
	fstat, err := file.Stat()
	if err != nil {
		Logger.Printf("Can not get file stat %v\n", err)
	}
	fwInfo.ImageLength = uint32(fstat.Size())
	return fwReadInfo(fwInfo, file)
}

func fwReadInfo(fwInfo *FirmwareInfo, firmware *os.File) int {
	var headerV2 binHeaderV2
	if err := binary.Read(firmware, binary.BigEndian, &headerV2); err != nil {
		Logger.Printf("Can not read info %v\n", err)
	}
	if (headerV2.BootImageRecord[0] == 0x12345678) &&
		(headerV2.BootImageRecord[1] == 0x11223344) &&
		(headerV2.BootImageRecord[2] == 0x55667788) {
		if LogLevel > E_VERVOSITY_VERVOSE {
			Logger.Printf("binHeaderV2\n")
		}

		fwInfo.ImageVersion = E_IMAGE_VERSION_2
		fwInfo.ROMVersion = headerV2.ROMVersion
		fwInfo.TextSectionLoadAddress = 0x04000000 + ((headerV2.DataSectionInfo >> 16) * 4)
		fwInfo.TextSectionLength = (headerV2.DataSectionInfo & 0x0000FFFF) * 4
		fwInfo.BssSectionLoadAddress = 0x04000000 + ((headerV2.BssSectionInfo >> 16) * 4)
		fwInfo.BssSectionLength = (headerV2.BssSectionInfo & 0x0000FFFF) * 4

		fwInfo.ResetEntryPoint = headerV2.ResetEntryPoint
		fwInfo.WakeUpEntryPoint = headerV2.WakeUpEntryPoint

		fwInfo.ImageDataV2 = headerV2.BootImageRecord

		fwInfo.MacAddressLocation = 0x10

		//fwInfo.TextData = &binHeaderV2.TextDataStart
	} else {
		var headerV1 binHeaderV1
		if err := binary.Read(firmware, binary.BigEndian, &headerV1); err != nil {
			Logger.Printf("Can not read info %v\n", err)
		}
		if LogLevel > E_VERVOSITY_VERVOSE {
			Logger.Printf("binHeaderV1\n")
		}

		fwInfo.ImageVersion = E_IMAGE_VERSION_1

		fwInfo.ROMVersion = headerV1.ROMVersion

		fwInfo.TextSectionLoadAddress = headerV1.TextStartAddress
		fwInfo.TextSectionLength = headerV1.TextLength
		fwInfo.BssSectionLoadAddress = headerV1.BssStartAddress
		fwInfo.BssSectionLength = headerV1.BssLength

		fwInfo.ResetEntryPoint = headerV1.ResetEntryPoint
		fwInfo.WakeUpEntryPoint = headerV1.WakeUpEntryPoint

		fwInfo.ImageDataV1 = [2]uint8{headerV1.ConfigByte0, headerV1.ConfigByte1}

		fwInfo.MacAddressLocation = 0x30
	}

	if LogLevel > E_VERVOSITY_WARN {
		Logger.Printf("Firmware INFO\n" +
			"\tImageVersion: %d\n" +
			"\tRom Version: 0x%08x\n" +
			"\tTextSectionLength: 0x%08x\n" +
			"\tBssSectionLength: 0x%08x\n",
			fwInfo.ImageVersion,
			fwInfo.ROMVersion,
			fwInfo.TextSectionLength,
			fwInfo.BssSectionLength)
	}
	if DEBUG {
		Logger.Printf("Firmware INFO2\n\t VALUE:: %v", fwInfo)
	}

	return E_FW_OK
}