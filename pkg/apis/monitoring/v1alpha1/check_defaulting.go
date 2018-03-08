/*
Copyright 2018 The Prohori Authors.

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

package v1alpha1

import (
	"log"
)

// DefaultingFunction sets default Check field values
func (CheckSchemeFns) DefaultingFunction(obj interface{}) {
	check := obj.(*Check)
	// set default field values here
	log.Printf("Defaulting fields for Check %s\n", check.Name)

	checkSpec := check.Spec

	// Defaulting selector
	if checkSpec.Type == CheckTypePod {
		selector := checkSpec.Selector
		if selector == nil {
			check.Spec.Selector = new(ObjectSelector)
		}
		if check.Spec.Selector.Namespace == "" {
			check.Spec.Selector.Namespace = check.Namespace
		}
	}

	// Defaulting PluginPullPolicy
	plugin := checkSpec.Plugin
	if plugin != nil {
		if plugin.PluginPullPolicy == "" {
			check.Spec.Plugin.PluginPullPolicy = PullPluginAlways
		}
	}

	// Defaulting CheckInterval
	if checkSpec.CheckInterval == 0 {
		check.Spec.CheckInterval = 60
	}

	// Defaulting AlertInterval
	if checkSpec.AlertInterval == 0 {
		check.Spec.AlertInterval = 300
	}
}
