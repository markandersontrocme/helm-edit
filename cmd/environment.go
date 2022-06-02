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
	"github.com/spf13/pflag"
)

type EnvSettings struct {
	AllValues bool
	DryRun    bool
	Namespace string
}

func New() *EnvSettings {
	envSettings := EnvSettings{}
	return &envSettings
}

func (s *EnvSettings) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&s.DryRun, "dry-run", false, "simulate upgrade command")
	fs.StringVarP(&s.Namespace, "namespace", "n", s.Namespace, "namespace scope of the release")
	fs.BoolVarP(&s.AllValues, "all", "a", false, "get all values")
}
