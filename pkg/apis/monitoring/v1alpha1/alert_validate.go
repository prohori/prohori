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

	"fmt"
	"github.com/prohori/prohori/pkg/apis/monitoring"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/endpoints/request"
)

// Validate checks that an instance of Alert is well formed
func (AlertStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	alert := obj.(*monitoring.Alert)
	log.Printf("Validating fields for Alert %s\n", alert.Name)
	errors := field.ErrorList{}

	alertSpec := alert.Spec

	// Validate spec.type
	alertType := alertSpec.Type
	if !(alertType == monitoring.TypePodAlert || alertType == monitoring.TypeNodeAlert || alertType == monitoring.TypeClusterAlert) {
		errors = append(errors, field.NotSupported(
			field.NewPath("spec", "type"),
			alertType,
			[]string{"PodAlert", "NodeAlert", "ClusterAlert"},
		))
	}

	// Validate selector
	selector := alertSpec.Selector
	if selector != nil {
		switch alertType {
		case monitoring.TypeClusterAlert:
			errors = append(errors, field.Forbidden(
				field.NewPath("spec", "selector"),
				"You can't use selector for ClusterAlert type",
			))
		case monitoring.TypeNodeAlert:
			if selector.Namespace != "" {
				errors = append(errors, field.Forbidden(
					field.NewPath("spec", "selector"),
					"You can't use namespace for NodeAlert type",
				))
			}
		}
	}

	// Validate plugin
	plugin := alertSpec.Plugin
	if plugin != nil {
		pp := plugin.PluginPullPolicy
		if pp != "" {
			if !(pp == monitoring.PullPluginAlways || pp == monitoring.PullPluginIfNotPresent) {
				errors = append(errors, field.NotSupported(
					field.NewPath("spec", "plugin", "pluginPullPolicy"),
					pp,
					[]string{"Always", "IfNotPresent"},
				))
			}
		}

		if plugin.Binary == "" {
			errors = append(errors, field.Required(
				field.NewPath("spec", "plugin", "binary"),
				"Must provided plugin download link",
			))
		}
	}

	// Validate CheckInterval
	ci := alertSpec.CheckInterval
	if ci != 0 && ci < 30 {
		errors = append(errors, field.Invalid(
			field.NewPath("spec", "checkInterval"),
			ci,
			"Must be at least 30 second. Default to 60",
		))
	}

	// Validate AlertInterval
	ai := alertSpec.AlertInterval
	if ai != 0 && ai < 60 {
		errors = append(errors, field.Invalid(
			field.NewPath("spec", "checkInterval"),
			ai,
			"Must be at least 60 second. Default to 300",
		))
	}

	// Validate ProblemState
	receivers := alertSpec.Receivers
	for i, r := range receivers {
		if !(r.State == monitoring.StateOK || r.State == monitoring.StateWarning || r.State == monitoring.StateCritical) {
			errors = append(errors, field.NotSupported(
				field.NewPath("spec", fmt.Sprintf("receivers[%d]", i), "state"),
				r.State,
				[]string{"OK", "WARNING", "CRITICAL"},
			))
		}
	}

	// perform validation here and add to errors using field.Invalid
	return errors
}
