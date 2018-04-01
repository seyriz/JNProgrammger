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
	CHIP_ID_MANUFACTURER_ID_MASK    = 0x00000fff
	CHIP_ID_PART_MASK               = 0x003ff000
	CHIP_ID_MASK_VERSION_MASK       = 0x0fc00000
	CHIP_ID_REV_MASK                = 0xf0000000

	CHIP_ID_JN5121_REV1A            = 0x00000000  /* ZED001               JN5121 Development      */
	CHIP_ID_JN5121_REV2A            = 0x10000000  /* ZED002               JN5121 Development      */
	CHIP_ID_JN5121_REV3A            = 0x20000000  /* ZED003               JN5121 Production       */

	CHIP_ID_JN5131_REV1A            = 0x00001000  /* Alberich             Never Produced          */

	CHIP_ID_JN5139_REV1A            = 0x00002000  /* BAL01                JN5139R                 */
	CHIP_ID_JN5139_REV2A            = 0x10002000  /* BAL02A               JN5139R1                */
	CHIP_ID_JN5139_REV2B            = 0x10002000  /* BAL02B               Test Chip Only          */
	CHIP_ID_JN5139_REV2C            = 0x10802000  /* BAL02C   (Trent ROM) JN5139T01 & JN5139J01   */

	CHIP_ID_JN5147_REV1A            = 0x00004686  /* */

	CHIP_ID_JN5148_REV2A            = 0x10004686  /* JAG02A               JN5148                  */
	CHIP_ID_JN5148_REV2B            = 0x10404686  /* JAG02B               JN5148                  */
	CHIP_ID_JN5148_REV2C            = 0x10804686  /* JAG02C               JN5148                  */
	CHIP_ID_JN5148_REV2D            = 0x10C04686  /* JAG02D   (Trent2 ROM)JN5148T01 & JN5148J01   */
	CHIP_ID_JN5148_REV2E            = 0x11004686  /* JAG02E   (JAG03A?)   JN5148Z01               */

	CHIP_ID_JN5142_REV1A            = 0x00005686  /* PUM01A               JN5142                  */
	CHIP_ID_JN5142_REV1B            = 0x00425686  /* PUM01B               JN5142                  */
	CHIP_ID_JN5142_REV1C            = 0x00845686  /* PUM01C               JN5142J01               */

	CHIP_ID_COG03                   = 0x00006686  /* Cougar COG03                                 */
	CHIP_ID_COG04                   = 0x00007686  /* Cougar COG04                                 */
	CHIP_ID_JN5168                  = 0x00008686  /* Cougar COG05                                 */
	CHIP_ID_JN5168_COG07            = 0x10008686  /* Cougar COG07                                 */
	CHIP_ID_JN5169                  = 0x6000B686  /* Cougar COG05                                 */

	/* SPI Flash device ID's  */
	FLASH_ST_M25P10			= 0x1010	/* ST microelectronics M25P10		*/
	FLASH_ST_M25P40			= 0x1212	/* ST microelectronics M25P40		*/
	FLASH_SST_SST25VF010A	= 0xbf49	/* SST SST25VF010A					*/
	FLASH_INTERNAL			= 0xccee	/* Internal on-chip flash 			*/
)

func CHIP_ID_PART(a int) int {
	return a & CHIP_ID_PART_MASK
}