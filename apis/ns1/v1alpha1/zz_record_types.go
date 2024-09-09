// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AnswersInitParameters struct {

	// Space delimited string of RDATA fields dependent on the record type.
	Answer *string `json:"answer,omitempty" tf:"answer,omitempty"`

	// meta is supported at the regions level. Meta
	// is documented below.
	// +mapType=granular
	Meta map[string]*string `json:"meta,omitempty" tf:"meta,omitempty"`

	// The region (Answer Group really) that this answer
	// belongs to. This should be one of the names specified in regions. Only a
	// single region per answer is currently supported. If you want an answer in
	// multiple regions, duplicating the answer (including metadata) is the correct
	// approach.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type AnswersObservation struct {

	// Space delimited string of RDATA fields dependent on the record type.
	Answer *string `json:"answer,omitempty" tf:"answer,omitempty"`

	// meta is supported at the regions level. Meta
	// is documented below.
	// +mapType=granular
	Meta map[string]*string `json:"meta,omitempty" tf:"meta,omitempty"`

	// The region (Answer Group really) that this answer
	// belongs to. This should be one of the names specified in regions. Only a
	// single region per answer is currently supported. If you want an answer in
	// multiple regions, duplicating the answer (including metadata) is the correct
	// approach.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type AnswersParameters struct {

	// Space delimited string of RDATA fields dependent on the record type.
	// +kubebuilder:validation:Optional
	Answer *string `json:"answer,omitempty" tf:"answer,omitempty"`

	// meta is supported at the regions level. Meta
	// is documented below.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Meta map[string]*string `json:"meta,omitempty" tf:"meta,omitempty"`

	// The region (Answer Group really) that this answer
	// belongs to. This should be one of the names specified in regions. Only a
	// single region per answer is currently supported. If you want an answer in
	// multiple regions, duplicating the answer (including metadata) is the correct
	// approach.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type FiltersInitParameters struct {

	// The filters' configuration. Simple key/value pairs
	// determined by the filter type.
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// Determines whether the filter is applied in the
	// filter chain.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The type of filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`
}

type FiltersObservation struct {

	// The filters' configuration. Simple key/value pairs
	// determined by the filter type.
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// Determines whether the filter is applied in the
	// filter chain.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The type of filter.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`
}

type FiltersParameters struct {

	// The filters' configuration. Simple key/value pairs
	// determined by the filter type.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// Determines whether the filter is applied in the
	// filter chain.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The type of filter.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter" tf:"filter,omitempty"`
}

type RecordInitParameters struct {

	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	Answers []AnswersInitParameters `json:"answers,omitempty" tf:"answers,omitempty"`

	// map of tags in the form of "key" = "value" where both key and value are strings
	BlockedTags []*string `json:"blockedTags,omitempty" tf:"blocked_tags,omitempty"`

	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and FQDN formatting below.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters []FiltersInitParameters `json:"filters,omitempty" tf:"filters,omitempty"`

	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link *string `json:"link,omitempty" tf:"link,omitempty"`

	// meta is supported at the regions level. Meta
	// is documented below.
	// +mapType=granular
	Meta map[string]*string `json:"meta,omitempty" tf:"meta,omitempty"`

	// The records' time to live (in seconds).
	OverrideTTL *bool `json:"overrideTtl,omitempty" tf:"override_ttl,omitempty"`

	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain regions here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions []RegionsInitParameters `json:"regions,omitempty" tf:"regions,omitempty"`

	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	ShortAnswers []*string `json:"shortAnswers,omitempty" tf:"short_answers,omitempty"`

	// The records' time to live (in seconds).
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// map of tags in the form of "key" = "value" where both key and value are strings
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The records' RR type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Whether to use EDNS client subnet data when
	// available(in filter chain).
	UseClientSubnet *bool `json:"useClientSubnet,omitempty" tf:"use_client_subnet,omitempty"`

	// The zone the record belongs to. Cannot have leading or
	// trailing dots (".") - see the example above and FQDN formatting below.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type RecordObservation struct {

	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	Answers []AnswersObservation `json:"answers,omitempty" tf:"answers,omitempty"`

	// map of tags in the form of "key" = "value" where both key and value are strings
	BlockedTags []*string `json:"blockedTags,omitempty" tf:"blocked_tags,omitempty"`

	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and FQDN formatting below.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	Filters []FiltersObservation `json:"filters,omitempty" tf:"filters,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	Link *string `json:"link,omitempty" tf:"link,omitempty"`

	// meta is supported at the regions level. Meta
	// is documented below.
	// +mapType=granular
	Meta map[string]*string `json:"meta,omitempty" tf:"meta,omitempty"`

	// The records' time to live (in seconds).
	OverrideTTL *bool `json:"overrideTtl,omitempty" tf:"override_ttl,omitempty"`

	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain regions here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	Regions []RegionsObservation `json:"regions,omitempty" tf:"regions,omitempty"`

	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	ShortAnswers []*string `json:"shortAnswers,omitempty" tf:"short_answers,omitempty"`

	// The records' time to live (in seconds).
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// map of tags in the form of "key" = "value" where both key and value are strings
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The records' RR type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Whether to use EDNS client subnet data when
	// available(in filter chain).
	UseClientSubnet *bool `json:"useClientSubnet,omitempty" tf:"use_client_subnet,omitempty"`

	// The zone the record belongs to. Cannot have leading or
	// trailing dots (".") - see the example above and FQDN formatting below.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type RecordParameters struct {

	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	// +kubebuilder:validation:Optional
	Answers []AnswersParameters `json:"answers,omitempty" tf:"answers,omitempty"`

	// map of tags in the form of "key" = "value" where both key and value are strings
	// +kubebuilder:validation:Optional
	BlockedTags []*string `json:"blockedTags,omitempty" tf:"blocked_tags,omitempty"`

	// The records' domain. Cannot have leading or trailing
	// dots - see the example above and FQDN formatting below.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// One or more NS1 filters for the record(order matters).
	// Filters are documented below.
	// +kubebuilder:validation:Optional
	Filters []FiltersParameters `json:"filters,omitempty" tf:"filters,omitempty"`

	// The target record to link to. This means this record is a
	// 'linked' record, and it inherits all properties from its target.
	// +kubebuilder:validation:Optional
	Link *string `json:"link,omitempty" tf:"link,omitempty"`

	// meta is supported at the regions level. Meta
	// is documented below.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Meta map[string]*string `json:"meta,omitempty" tf:"meta,omitempty"`

	// The records' time to live (in seconds).
	// +kubebuilder:validation:Optional
	OverrideTTL *bool `json:"overrideTtl,omitempty" tf:"override_ttl,omitempty"`

	// One or more "regions" for the record. These are really
	// just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
	// but remain regions here for legacy reasons. Regions are
	// documented below. Please note the ordering requirement!
	// +kubebuilder:validation:Optional
	Regions []RegionsParameters `json:"regions,omitempty" tf:"regions,omitempty"`

	// One or more NS1 answers for the records' specified type.
	// Answers are documented below.
	// +kubebuilder:validation:Optional
	ShortAnswers []*string `json:"shortAnswers,omitempty" tf:"short_answers,omitempty"`

	// The records' time to live (in seconds).
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// map of tags in the form of "key" = "value" where both key and value are strings
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The records' RR type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Whether to use EDNS client subnet data when
	// available(in filter chain).
	// +kubebuilder:validation:Optional
	UseClientSubnet *bool `json:"useClientSubnet,omitempty" tf:"use_client_subnet,omitempty"`

	// The zone the record belongs to. Cannot have leading or
	// trailing dots (".") - see the example above and FQDN formatting below.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type RegionsInitParameters struct {

	// meta is supported at the regions level. Meta
	// is documented below.
	// +mapType=granular
	Meta map[string]*string `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of the region (or Answer Group).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RegionsObservation struct {

	// meta is supported at the regions level. Meta
	// is documented below.
	// +mapType=granular
	Meta map[string]*string `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of the region (or Answer Group).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RegionsParameters struct {

	// meta is supported at the regions level. Meta
	// is documented below.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Meta map[string]*string `json:"meta,omitempty" tf:"meta,omitempty"`

	// Name of the region (or Answer Group).
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

// RecordSpec defines the desired state of Record
type RecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RecordInitParameters `json:"initProvider,omitempty"`
}

// RecordStatus defines the observed state of Record.
type RecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Record is the Schema for the Records API. Provides a NS1 Record resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ns1}
type Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || (has(self.initProvider) && has(self.initProvider.domain))",message="spec.forProvider.domain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zone) || (has(self.initProvider) && has(self.initProvider.zone))",message="spec.forProvider.zone is a required parameter"
	Spec   RecordSpec   `json:"spec"`
	Status RecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordList contains a list of Records
type RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Record `json:"items"`
}

// Repository type metadata.
var (
	Record_Kind             = "Record"
	Record_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Record_Kind}.String()
	Record_KindAPIVersion   = Record_Kind + "." + CRDGroupVersion.String()
	Record_GroupVersionKind = CRDGroupVersion.WithKind(Record_Kind)
)

func init() {
	SchemeBuilder.Register(&Record{}, &RecordList{})
}
