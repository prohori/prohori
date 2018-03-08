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

// Validate checks that an instance of Check is well formed
func (CheckStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	check := obj.(*monitoring.Check)
	log.Printf("Validating fields for Check %s\n", check.Name)
	errors := field.ErrorList{}

	checkSpec := check.Spec

	// Validate spec.type
	checkType := checkSpec.Type
	if !(checkType == monitoring.CheckTypePod || checkType == monitoring.CheckTypeNode || checkType == monitoring.CheckTypeCluster) {
		errors = append(errors, field.NotSupported(
			field.NewPath("spec", "type"),
			checkType,
			[]string{"PodCheck", "NodeCheck", "ClusterCheck"},
		))
	}

	// Validate selector
	selector := checkSpec.Selector
	if selector != nil {
		switch checkType {
		case monitoring.CheckTypeCluster:
			errors = append(errors, field.Forbidden(
				field.NewPath("spec", "selector"),
				"You can't use selector for ClusterCheck type",
			))
		case monitoring.CheckTypeNode:
			if selector.Namespace != "" {
				errors = append(errors, field.Forbidden(
					field.NewPath("spec", "selector"),
					"You can't use namespace for NodeCheck type",
				))
			}
		}
	}

	// Validate plugin
	plugin := checkSpec.Plugin
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
	ci := checkSpec.CheckInterval
	if ci != 0 && ci < 30 {
		errors = append(errors, field.Invalid(
			field.NewPath("spec", "checkInterval"),
			ci,
			"Must be at least 30 second. Default to 60",
		))
	}

	// Validate AlertInterval
	ai := checkSpec.AlertInterval
	if ai != 0 && ai < 60 {
		errors = append(errors, field.Invalid(
			field.NewPath("spec", "checkInterval"),
			ai,
			"Must be at least 60 second. Default to 300",
		))
	}

	// Validate CheckState
	receivers := checkSpec.Receivers
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
