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

// environmentCreateCmd represents the environment list command
var environmentCreateCmd = &cobra.Command{
	Use:   "create",
	Aliases: []string{"new", "c", "n"},
	Short: "create a new environment",
	Long: `The environment create command requests from O² Control the
creation of a new O² environment.`,
	Run:   control.WrapCall(control.CreateEnvironment),
}

func init() {
	environmentCmd.AddCommand(environmentCreateCmd)

	environmentCreateCmd.Flags().StringArrayP("roles", "r", []string{}, "list of roles to be included in the environment")
	environmentCreateCmd.MarkFlagRequired("roles")
}
