/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/apache/camel-k/pkg/apis/camel/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IntegrationPlatformBuildSpecApplyConfiguration represents an declarative configuration of the IntegrationPlatformBuildSpec type for use
// with apply.
type IntegrationPlatformBuildSpecApplyConfiguration struct {
	BuildStrategy          *v1.BuildStrategy                           `json:"buildStrategy,omitempty"`
	PublishStrategy        *v1.IntegrationPlatformBuildPublishStrategy `json:"publishStrategy,omitempty"`
	RuntimeVersion         *string                                     `json:"runtimeVersion,omitempty"`
	RuntimeProvider        *v1.RuntimeProvider                         `json:"runtimeProvider,omitempty"`
	BaseImage              *string                                     `json:"baseImage,omitempty"`
	Registry               *RegistrySpecApplyConfiguration             `json:"registry,omitempty"`
	Timeout                *metav1.Duration                            `json:"timeout,omitempty"`
	Maven                  *MavenSpecApplyConfiguration                `json:"maven,omitempty"`
	PublishStrategyOptions map[string]string                           `json:"PublishStrategyOptions,omitempty"`
}

// IntegrationPlatformBuildSpecApplyConfiguration constructs an declarative configuration of the IntegrationPlatformBuildSpec type for use with
// apply.
func IntegrationPlatformBuildSpec() *IntegrationPlatformBuildSpecApplyConfiguration {
	return &IntegrationPlatformBuildSpecApplyConfiguration{}
}

// WithBuildStrategy sets the BuildStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BuildStrategy field is set to the value of the last call.
func (b *IntegrationPlatformBuildSpecApplyConfiguration) WithBuildStrategy(value v1.BuildStrategy) *IntegrationPlatformBuildSpecApplyConfiguration {
	b.BuildStrategy = &value
	return b
}

// WithPublishStrategy sets the PublishStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PublishStrategy field is set to the value of the last call.
func (b *IntegrationPlatformBuildSpecApplyConfiguration) WithPublishStrategy(value v1.IntegrationPlatformBuildPublishStrategy) *IntegrationPlatformBuildSpecApplyConfiguration {
	b.PublishStrategy = &value
	return b
}

// WithRuntimeVersion sets the RuntimeVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RuntimeVersion field is set to the value of the last call.
func (b *IntegrationPlatformBuildSpecApplyConfiguration) WithRuntimeVersion(value string) *IntegrationPlatformBuildSpecApplyConfiguration {
	b.RuntimeVersion = &value
	return b
}

// WithRuntimeProvider sets the RuntimeProvider field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RuntimeProvider field is set to the value of the last call.
func (b *IntegrationPlatformBuildSpecApplyConfiguration) WithRuntimeProvider(value v1.RuntimeProvider) *IntegrationPlatformBuildSpecApplyConfiguration {
	b.RuntimeProvider = &value
	return b
}

// WithBaseImage sets the BaseImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BaseImage field is set to the value of the last call.
func (b *IntegrationPlatformBuildSpecApplyConfiguration) WithBaseImage(value string) *IntegrationPlatformBuildSpecApplyConfiguration {
	b.BaseImage = &value
	return b
}

// WithRegistry sets the Registry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Registry field is set to the value of the last call.
func (b *IntegrationPlatformBuildSpecApplyConfiguration) WithRegistry(value *RegistrySpecApplyConfiguration) *IntegrationPlatformBuildSpecApplyConfiguration {
	b.Registry = value
	return b
}

// WithTimeout sets the Timeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Timeout field is set to the value of the last call.
func (b *IntegrationPlatformBuildSpecApplyConfiguration) WithTimeout(value metav1.Duration) *IntegrationPlatformBuildSpecApplyConfiguration {
	b.Timeout = &value
	return b
}

// WithMaven sets the Maven field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Maven field is set to the value of the last call.
func (b *IntegrationPlatformBuildSpecApplyConfiguration) WithMaven(value *MavenSpecApplyConfiguration) *IntegrationPlatformBuildSpecApplyConfiguration {
	b.Maven = value
	return b
}

// WithPublishStrategyOptions puts the entries into the PublishStrategyOptions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the PublishStrategyOptions field,
// overwriting an existing map entries in PublishStrategyOptions field with the same key.
func (b *IntegrationPlatformBuildSpecApplyConfiguration) WithPublishStrategyOptions(entries map[string]string) *IntegrationPlatformBuildSpecApplyConfiguration {
	if b.PublishStrategyOptions == nil && len(entries) > 0 {
		b.PublishStrategyOptions = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.PublishStrategyOptions[k] = v
	}
	return b
}
