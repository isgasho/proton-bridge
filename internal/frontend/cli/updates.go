// Copyright (c) 2021 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package cli

import (
	"strings"

	"github.com/ProtonMail/proton-bridge/internal/bridge"
	"github.com/abiosoft/ishell"
)

func (f *frontendCLI) checkUpdates(c *ishell.Context) {
	f.Println("Your version is up to date.")
}

func (f *frontendCLI) printCredits(c *ishell.Context) {
	for _, pkg := range strings.Split(bridge.Credits, ";") {
		f.Println(pkg)
	}
}
