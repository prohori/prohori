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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Alert
// +k8s:openapi-gen=true
// +resource:path=alerts,strategy=AlertStrategy
type Alert struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata"`

	// Specification of the desired behavior of the alert.
	Spec AlertSpec `json:"spec"`

	// Most recently observed status of the alert.
	// This data may not be up to date.
	// Populated by the system.
	// Read-only.
	// +optional
	Status AlertStatus `json:"status,omitempty"`
}

// AlertSpec defines the desired state of Alert
type AlertSpec struct {
	// Alert type.
	// One of PodAlert, NodeAlert, ClusterAlert
	// Each alert must have a valid type.
	// Cannot be updated.
	Type AlertType `json:"type"`

	// Check command.
	// Can either be supported built-in methods or external plugin
	// Each alert must have a command.
	// Cannot be updated.
	Command string `json:"command"`

	// Flags to the check command.
	// Can be updated.
	// +optional
	Flags map[string]string `json:"flags,omitempty"`

	// To select Kubernetes objects to apply check command for.
	// Can be updated.
	// +optional
	Selector *ObjectSelector `json:"selector,omitempty"`

	// Check command plugin.
	// To download external plugin for check command.
	// Can be updated.
	// +optional
	Plugin *CommandPlugin `json:"plugin,omitempty"`

	// Check command interval
	// How frequently check command will be executed
	// Must be at least 30s.
	// Default 60s (1m) if not provided.
	// Can be updated.
	// +optional
	CheckInterval int64 `json:"checkInterval,omitempty"`

	// Notification interval
	// How frequently notifications will be send
	// Must be at least 60s (1m).
	// Default 300s (5m) if not provided.
	// Can be updated.
	// +optional
	AlertInterval int64 `json:"alertInterval,omitempty"`

	// Secret containing notifier credentials
	// If not specified, no notification will send
	// Can be updated.
	// +optional
	NotifierSecretName string `json:"notifierSecretName,omitempty"`

	// Receivers is an optional list of NotificationReceiver.
	// If provided, notifications will be send to specified receivers.
	// Can be updated.
	// +optional
	Receivers []NotificationReceiver `json:"receivers,omitempty"`
}

// AlertPhase is a label for the condition of a alert at the current time.
type AlertPhase string

// These are the valid statuses of alerts.
const (
	// AlertPending means the alert has been accepted by the system, but has not been started.
	// pulling plugins onto the host if necessary.
	AlertPending AlertPhase = "Pending"
	// AlertRunning means the alert has been set.
	AlertRunning AlertPhase = "Running"
	// AlertFailed means that alert has terminated in a failure.
	AlertFailed AlertPhase = "Failed"
	// AlertUnknown means that for some reason the state of the alert could not be obtained.
	AlertUnknown AlertPhase = "Unknown"
)

// AlertStatus defines the observed state of Alert
type AlertStatus struct {
	// Current condition of the alert.
	// +optional
	Phase AlertPhase `json:"phase,omitempty"`
}

type AlertType string

const (
	TypePodAlert     AlertType = "PodAlert"
	TypeNodeAlert    AlertType = "NodeAlert"
	TypeClusterAlert AlertType = "ClusterAlert"
)

type ObjectSelector struct {
	// Kubernetes object name
	// Can be updated.
	// +optional
	Name string `json:"name,omitempty"`

	// Kubernetes objects namespace
	// Default to alert object namespaces
	// Can be updated
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Selector is a label query over pods/nodes.
	// Can be updated.
	// +optional
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

type PullPolicy string

const (
	PullPluginAlways       PullPolicy = "Always"
	PullPluginIfNotPresent PullPolicy = "IfNotPresent"
)

type CommandPlugin struct {
	// Plugin pull policy.
	// One of Always, IfNotPresent.
	// Defaults to IfNotPresent
	// Can be updated.
	// +optional
	PluginPullPolicy PullPolicy `json:"pluginPullPolicy,omitempty"`

	// Plugin binary download link
	// Can be updated.
	// Update will work only if pluginPullPolicy is set to Always
	Binary string `json:"Binary"`
}

type ProblemState string

const (
	StateOK       ProblemState = "OK"
	StateWarning  ProblemState = "WARNING"
	StateCritical ProblemState = "CRITICAL"
)

type NotificationReceiver struct {
	// For which state notification will be sent
	State ProblemState `json:"state,omitempty"`

	// To whom notification will be sent
	To []string `json:"to,omitempty"`

	// How this notification will be sent
	Notifier string `json:"notifier,omitempty"`
}
