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

// DefaultingFunction sets default Alert field values
func (AlertSchemeFns) DefaultingFunction(o interface{}) {
	alert := o.(*Alert)
	// set default field values here
	log.Printf("Defaulting fields for Alert %s\n", alert.Name)

	alertSpec := alert.Spec

	// Defaulting selector
	if alertSpec.Type == TypePodAlert {
		selector := alertSpec.Selector
		if selector == nil {
			alert.Spec.Selector = new(ObjectSelector)
		}
		if alert.Spec.Selector.Namespace == "" {
			alert.Spec.Selector.Namespace = alert.Namespace
		}
	}

	// Defaulting PluginPullPolicy
	plugin := alertSpec.Plugin
	if plugin != nil {
		if plugin.PluginPullPolicy == "" {
			alert.Spec.Plugin.PluginPullPolicy = PullPluginAlways
		}
	}

	// Defaulting CheckInterval
	if alertSpec.CheckInterval == 0 {
		alert.Spec.CheckInterval = 60
	}

	// Defaulting AlertInterval
	if alertSpec.AlertInterval == 0 {
		alert.Spec.AlertInterval = 300
	}
}
