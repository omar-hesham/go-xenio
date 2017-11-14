// Copyright 2017 The go-xenio Authors
// Copyright 2016 The go-ethereum Authors
//
// This file is part of the go-xenio library.
//
// The go-xenio library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-xenio library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-xenio library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"fmt"
)

const (
	XenioVersion = "0.2.3"
	XenioMeta = "unstable"
	VersionMajor = 1        // Major version component of the current release
	VersionMinor = 7        // Minor version component of the current release
	VersionPatch = 2        // Patch version component of the current release
	VersionMeta  = "stable" // Version metadata to append to the version string
)

// Version holds the textual version string.
var Version = func() string {
	v := fmt.Sprintf("%s%s/Geth %d.%d.%d",XenioVersion,XenioMeta, VersionMajor, VersionMinor, VersionPatch)
	if VersionMeta != "" {
		v += VersionMeta
	}
	return v
}()

func VersionWithCommit(gitCommit string) string {
	vsn := Version
	if len(gitCommit) >= 8 {
		vsn += "-" + gitCommit[:8]
	}
	return vsn
}
