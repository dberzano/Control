/*
 * === This file is part of ALICE O² ===
 *
 * Copyright 2018 CERN and copyright holders of ALICE O².
 * Author: Teo Mrnjavac <teo.mrnjavac@cern.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * In applying this license CERN does not waive the privileges and
 * immunities granted to it by virtue of its status as an
 * Intergovernmental Organization or submit itself to any jurisdiction.
 */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/AliceO2Group/Control/coconut/control"
)

// environmentDestroyCmd represents the environment list command
var environmentDestroyCmd = &cobra.Command{
	Use:   "destroy [environment id]",
	Aliases: []string{"des", "d"},
	Short: "destroy an environment",
	Long: `The environment destroy command instructs O² Control to
teardown an existing O² environment. The environment must be in the 
CONFIGURED state.`,
	Run:   control.WrapCall(control.DestroyEnvironment),
	Args:  cobra.ExactArgs(1),
}

func init() {
	environmentCmd.AddCommand(environmentDestroyCmd)
}
