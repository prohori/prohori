
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


package alert

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/prohori/prohori/pkg/apis/monitoring/v1alpha1"
	"github.com/prohori/prohori/pkg/controller/sharedinformers"
	listers "github.com/prohori/prohori/pkg/client/listers_generated/monitoring/v1alpha1"
)

// +controller:group=monitoring,version=v1alpha1,kind=Alert,resource=alerts
type AlertControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Alert
	lister listers.AlertLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *AlertControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing alerts labels
	c.lister = arguments.GetSharedInformers().Factory.Monitoring().V1alpha1().Alerts().Lister()
}

// Reconcile handles enqueued messages
func (c *AlertControllerImpl) Reconcile(u *v1alpha1.Alert) error {
	// Implement controller logic here
	log.Printf("Running reconcile Alert for %s\n", u.Name)
	return nil
}

func (c *AlertControllerImpl) Get(namespace, name string) (*v1alpha1.Alert, error) {
	return c.lister.Alerts(namespace).Get(name)
}
