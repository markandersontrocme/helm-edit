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

package helm

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/MarkAndersonTrocme/helm-edit/pkg/common"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"helm.sh/helm/pkg/chartutil"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func EditRelease(editOptions common.EditOptions) error {
	editor := getEnv("EDITOR", "vim")
	cfg, err := GetActionConfig(editOptions.ReleaseNamespace)
	if err != nil {
		return errors.Wrap(err, "failed to get Helm action configuration")
	}

	var releaseName = editOptions.ReleaseName
	releaseToEdit, err := getLatestRelease(releaseName, cfg)
	if err != nil {
		return errors.Wrapf(err, "failed to get release '%s' latest version", releaseName)
	}

	originalValues, err := getReleaseValues(releaseToEdit.Name, editOptions.AllValues, cfg)
	if err != nil {
		return errors.Wrapf(err, "failed to get values '%s' latest version", releaseName)
	}

	file, err := ioutil.TempFile("", "values.*.yaml")
	if err != nil {
		return errors.Wrap(err, "failed to create temp file")
	}
	defer os.Remove(file.Name())

	yOriginalValues, err := yaml.Marshal(&originalValues)
	if err != nil {
		return errors.Wrap(err, "failed to convert to YAML")
	}

	os.WriteFile(file.Name(), yOriginalValues, 0644)

	cmd := exec.Command(editor, file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	newValues, err := chartutil.ReadValuesFile(file.Name())
	if err != nil {
		return err
	}

	updatedValues := newValues.AsMap()
	if err != nil {
		return err
	}

	if !cmp.Equal(originalValues, updatedValues) {
		upgrade := action.NewUpgrade(cfg)

		upgrade.DryRun = editOptions.DryRun
		newRelease, err := upgrade.Run(
			releaseToEdit.Name,
			releaseToEdit.Chart,
			updatedValues,
		)
		if err != nil {
			return err
		}

		if editOptions.DryRun {
			log.Println("NOTE: This is in dry-run mode, the following actions will not be executed.")
			log.Println("Run without --dry-run to take the actions described below:")
			log.Println()
		}

		log.Printf("Release %q has been edited. Happy Helming!\n%s", releaseToEdit.Name, newRelease.Info.Notes)
	} else {
		log.Printf("Edit cancelled, no changes made!")
	}

	return nil
}

func getEnv(key, fallback string) string {
	env := os.Getenv(key)
	if env == "" {
		return fallback
	}
	return env
}

func getLatestRelease(releaseName string, cfg *action.Configuration) (*release.Release, error) {
	return cfg.Releases.Last(releaseName)
}

func getReleaseValues(releaseName string, allValues bool, cfg *action.Configuration) (map[string]interface{}, error) {
	getValues := action.NewGetValues(cfg)
	getValues.AllValues = allValues
	values, err := getValues.Run(releaseName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get values '%s' latest version", releaseName)
	}
	return values, nil
}
