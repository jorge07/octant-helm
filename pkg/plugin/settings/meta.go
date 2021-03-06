/*
Copyright 2019 Blood Orange

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

package settings // import "github.com/bloodorangeio/octant-helm/pkg/plugin/settings"

const (
	name        = "helm" // This should stay lowercase for routing purposes
	description = "Helm support (v3)"
	rootNavIcon = "boat" // See https://clarity.design/icons for all options
)

func GetName() string {
	return name
}

func GetDescription() string {
	return description
}
