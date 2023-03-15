/*
Copyright
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"io"

	"github.com/markandersontrocme/helm-edit/pkg/common"
	"github.com/markandersontrocme/helm-edit/pkg/helm"
	"github.com/spf13/cobra"
)

type EditOptions struct {
	AllValues        bool
	DryRun           bool
	ReleaseName      string
	ReleaseNamespace string
}

var (
	settings *EnvSettings
)

func newEditCmd(out io.Writer, args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [RELEASE]",
		Short: "edit helm values",
		Args:  cobra.ExactArgs(1),
		RunE:  runEdit,
	}

	flags := cmd.PersistentFlags()
	flags.Parse(args)
	settings = new(EnvSettings)

	settings.AddFlags(flags)

	return cmd
}

func runEdit(cmd *cobra.Command, args []string) error {
	releaseName := args[0]

	editOptions := EditOptions{
		AllValues:        settings.AllValues,
		DryRun:           settings.DryRun,
		ReleaseName:      releaseName,
		ReleaseNamespace: settings.Namespace,
	}

	return Edit(editOptions)
}

func Edit(editOptions EditOptions) error {
	options := common.EditOptions{
		AllValues:        editOptions.AllValues,
		DryRun:           editOptions.DryRun,
		ReleaseName:      editOptions.ReleaseName,
		ReleaseNamespace: editOptions.ReleaseNamespace,
	}

	if err := helm.EditRelease(options); err != nil {
		return err
	}
	return nil
}
