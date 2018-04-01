//JNProgrammer;JN516x firmware programmer written in go
// Copyright (C) HanWool Lee <kudnya@gmail.com>
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

import "log"

const (
	E_STATUS_OK = iota
	E_STATUS_ERROR
	E_STATUS_ERROR_WRITING
	E_STATUS_ERROR_READING
	E_STATUS_FAILED_TO_OPEN_FILE
	E_STATUS_BAD_PARAMETER
	E_STATUS_NULL_PARAMETER
	E_STATUS_INCOMPATIBLE
)

var Logger *log.Logger

type VERVOSITY uint64
var LogLevel VERVOSITY

const (
	E_VERVOSITY_NONE VERVOSITY = iota
	E_VERVOSITY_ERROR
	E_VERVOSITY_WARN
	E_VERVOSITY_INFO
	E_VERVOSITY_VERVOSE
)

var DEBUG bool