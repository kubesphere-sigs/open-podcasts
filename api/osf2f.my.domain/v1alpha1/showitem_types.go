/*
Copyright 2022 The open-podcasts Authors. All rights reserved.

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

// ShowItemSpec defines the desired state of ShowItem
type ShowItemSpec struct {
	ShowRef      string `json:"showRef"`
	Index        int    `json:"index,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Image        string `json:"image,omitempty"`
	Filename     string `json:"filename,omitempty"`
	LocalStorage string `json:"localStorage,omitempty"`
}

// ShowItemStatus defines the observed state of ShowItem
type ShowItemStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ShowItem is the Schema for the showitems API
type ShowItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ShowItemSpec   `json:"spec,omitempty"`
	Status ShowItemStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ShowItemList contains a list of ShowItem
type ShowItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShowItem `json:"items"`
}

const AnnotationKeyDownloadURL = "show.item.download.url"