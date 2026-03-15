/*
Copyright 2026 The Kubernetes Authors.

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

package hub

// PolicySpec describes a reference to a policy object in the same namespace.
type PolicySpec LocalObjectRef

// PolicyStatus describes an observed policy applied to this VM.
type PolicyStatus struct {
	PolicySpec `json:",inline"`

	// Generation describes the observed generation of the policy applied to
	// this VM.
	Generation int64 `json:"generation"`
}
