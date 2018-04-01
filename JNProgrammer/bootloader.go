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

const (
	BL_MAX_CHUNK_SIZE = 248

	BL_TIMEOUT_1S = 1000000
	BL_TIMEOUT_10S = 10000000


	/* JN513x / JN514x definitions */
	JN514X_ROM_ID_ADDR = 0x00000004
	JN514X_MAC_ADDRESS_LOCATION = 0x00000010

	/* JN516x definitions */

	/* Location of MAC address in Index sector */
	JN516X_MAC_INDEX_SECTOR_PAGE = 5
	JN516X_MAC_INDEX_SECTOR_WORD = 7

	/* Location of MAC address in memory map */
	JN516X_CUSTOMER_MAC_ADDRESS_LOCATION = 0x01001570
	JN516X_MAC_ADDRESS_LOCATION = 0x01001580

	/* Location of bootloader information in memory map */
	JN516X_BOOTLOADER_VERSION_ADDRESS = 0x00000062
	JN516X_BOOTLOADER_ENTRY = 0x00000066

	/* Location of device configuration in memory map */
	JN516X_INDEX_SECTOR_DEVICE_CONFIG_ADDR = 0x01001500

	/* Location of Customer Settings in memory map */
	JN516X_INDEX_CUSTOMER_SETTINGS_ADDR = 0x01001510
)

const (
	E_BL_MSG_TYPE_FLASH_ERASE_REQUEST 					= 0x07
	E_BL_MSG_TYPE_FLASH_ERASE_RESPONSE					= 0x08
	E_BL_MSG_TYPE_FLASH_PROGRAM_REQUEST					= 0x09
	E_BL_MSG_TYPE_FLASH_PROGRAM_RESPONSE				= 0x0a
	E_BL_MSG_TYPE_FLASH_READ_REQUEST					= 0x0b
	E_BL_MSG_TYPE_FLASH_READ_RESPONSE					= 0x0c
	E_BL_MSG_TYPE_FLASH_SECTOR_ERASE_REQUEST			= 0x0d
	E_BL_MSG_TYPE_FLASH_SECTOR_ERASE_RESPONSE			= 0x0e
	E_BL_MSG_TYPE_FLASH_WRITE_STATUS_REGISTER_REQUEST	= 0x0f
	E_BL_MSG_TYPE_FLASH_WRITE_STATUS_REGISTER_RESPONSE	= 0x10
	E_BL_MSG_TYPE_RAM_WRITE_REQUEST						= 0x1d
	E_BL_MSG_TYPE_RAM_WRITE_RESPONSE					= 0x1e
	E_BL_MSG_TYPE_RAM_READ_REQUEST						= 0x1f
	E_BL_MSG_TYPE_RAM_READ_RESPONSE						= 0x20
	E_BL_MSG_TYPE_RAM_RUN_REQUEST						= 0x21
	E_BL_MSG_TYPE_RAM_RUN_RESPONSE						= 0x22
	E_BL_MSG_TYPE_FLASH_READ_ID_REQUEST					= 0x25
	E_BL_MSG_TYPE_FLASH_READ_ID_RESPONSE				= 0x26
	E_BL_MSG_TYPE_SET_BAUD_REQUEST						= 0x27
	E_BL_MSG_TYPE_SET_BAUD_RESPONSE						= 0x28
	E_BL_MSG_TYPE_FLASH_SELECT_TYPE_REQUEST				= 0x2c
	E_BL_MSG_TYPE_FLASH_SELECT_TYPE_RESPONSE			= 0x2d

	E_BL_MSG_TYPE_GET_CHIPID_REQUEST                    = 0x32
	E_BL_MSG_TYPE_GET_CHIPID_RESPONSE                   = 0x33
)

const (
	E_BL_RESPONSE_OK									= 0x00
	E_BL_RESPONSE_NOT_SUPPORTED							= 0xff
	E_BL_RESPONSE_WRITE_FAIL							= 0xfe
	E_BL_RESPONSE_INVALID_RESPONSE						= 0xfd
	E_BL_RESPONSE_CRC_ERROR								= 0xfc
	E_BL_RESPONSE_ASSERT_FAIL							= 0xfb
	E_BL_RESPONSE_USER_INTERRUPT						= 0xfa
	E_BL_RESPONSE_READ_FAIL								= 0xf9
	E_BL_RESPONSE_TST_ERROR								= 0xf8
	E_BL_RESPONSE_AUTH_ERROR							= 0xf7
	E_BL_RESPONSE_NO_RESPONSE							= 0xf6
	E_BL_RESPONSE_ERROR									= 0xf0
)

type ChipDetails struct {
	ChipId uint32
	SupportedFirmware uint32

	RamSize uint32
	FlashSize uint32

	BootloaderVersion uint32

	CustomerSettings uint32

	MaxAddress [8]uint8
}

type FlashDevice struct {
	FlashId uint16
	FlashType uint8
	FlashName string
}

var FlashDevices = []FlashDevice{
	FlashDevice{
		0x0505,
		4,
		"ST M25P05-A",
	},
	FlashDevice{
		0x1010,
		0,
		"ST M25P10-A",
	},
	FlashDevice{
		0x1111,
		5,
		"ST M25P20-A",
	},
	FlashDevice{
		0x1212,
		3,
		"ST M25P40",
	},
	FlashDevice{
		0xbf49,
		1,
		"SST 25VF010A",
	},
	FlashDevice{
		0x1f60,
		2,
		"Atmel 25F512",
	},
	FlashDevice{
		0xccee,
		8,
		"JN516x",
	},
}


