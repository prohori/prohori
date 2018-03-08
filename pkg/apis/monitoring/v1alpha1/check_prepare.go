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
	"github.com/prohori/prohori/pkg/apis/monitoring"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"
)

func (s CheckStrategy) PrepareForCreate(ctx request.Context, obj runtime.Object) {
	// Invoke the parent implementation to strip the Status
	s.DefaultStorageStrategy.PrepareForCreate(ctx, obj)

	// Cast the element
	check := obj.(*monitoring.Check)

	// Custom PrepareForCreate logic here
	check.SetFinalizers([]string{"prohori/controller"})
	check.Status = monitoring.CheckStatus{
		Phase: monitoring.CheckPending,
	}
}
