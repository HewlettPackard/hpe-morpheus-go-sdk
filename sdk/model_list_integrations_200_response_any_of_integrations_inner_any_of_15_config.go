/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 8.0.7
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config{}

// ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config struct for ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config
type ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config struct {
	IncidentAccess               *bool                                                                                           `json:"incidentAccess,omitempty"`
	RequestAccess                *bool                                                                                           `json:"requestAccess,omitempty"`
	ServiceNowCMDBBusinessObject *string                                                                                         `json:"serviceNowCMDBBusinessObject,omitempty"`
	ServiceNowCustomCmdbMapping  *string                                                                                         `json:"serviceNowCustomCmdbMapping,omitempty"`
	ServiceNowCmdbClassMapping   []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigServiceNowCmdbClassMappingInner `json:"serviceNowCmdbClassMapping,omitempty"`
	WebServiceImportUrl          *string                                                                                         `json:"webServiceImportUrl,omitempty"`
	WebServiceImportSysId        *string                                                                                         `json:"webServiceImportSysId,omitempty"`
	WebServiceOperationUrl       *string                                                                                         `json:"webServiceOperationUrl,omitempty"`
	CmdbMode                     *string                                                                                         `json:"cmdbMode,omitempty"`
	PreparedForSync              *bool                                                                                           `json:"preparedForSync,omitempty"`
	AdditionalProperties         map[string]interface{}                                                                          `json:",remain"`
}

type _ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config

// NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config {
	this := ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config{}
	var cmdbMode string = "TABLE"
	this.CmdbMode = &cmdbMode
	return &this
}

// NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigWithDefaults instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigWithDefaults() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config {
	this := ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config{}
	var cmdbMode string = "TABLE"
	this.CmdbMode = &cmdbMode
	return &this
}

// GetIncidentAccess returns the IncidentAccess field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetIncidentAccess() bool {
	if o == nil || IsNil(o.IncidentAccess) {
		var ret bool
		return ret
	}
	return *o.IncidentAccess
}

// GetIncidentAccessOk returns a tuple with the IncidentAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetIncidentAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.IncidentAccess) {
		return nil, false
	}
	return o.IncidentAccess, true
}

// IsSetIncidentAccess returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) IsSetIncidentAccess() bool {
	if o != nil && !IsNil(o.IncidentAccess) {
		return true
	}

	return false
}

// SetIncidentAccess gets a reference to the given bool and assigns it to the IncidentAccess field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetIncidentAccess(v bool) {
	o.IncidentAccess = &v
}

// GetRequestAccess returns the RequestAccess field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetRequestAccess() bool {
	if o == nil || IsNil(o.RequestAccess) {
		var ret bool
		return ret
	}
	return *o.RequestAccess
}

// GetRequestAccessOk returns a tuple with the RequestAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetRequestAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestAccess) {
		return nil, false
	}
	return o.RequestAccess, true
}

// IsSetRequestAccess returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) IsSetRequestAccess() bool {
	if o != nil && !IsNil(o.RequestAccess) {
		return true
	}

	return false
}

// SetRequestAccess gets a reference to the given bool and assigns it to the RequestAccess field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetRequestAccess(v bool) {
	o.RequestAccess = &v
}

// GetServiceNowCMDBBusinessObject returns the ServiceNowCMDBBusinessObject field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCMDBBusinessObject() string {
	if o == nil || IsNil(o.ServiceNowCMDBBusinessObject) {
		var ret string
		return ret
	}
	return *o.ServiceNowCMDBBusinessObject
}

// GetServiceNowCMDBBusinessObjectOk returns a tuple with the ServiceNowCMDBBusinessObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCMDBBusinessObjectOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceNowCMDBBusinessObject) {
		return nil, false
	}
	return o.ServiceNowCMDBBusinessObject, true
}

// IsSetServiceNowCMDBBusinessObject returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) IsSetServiceNowCMDBBusinessObject() bool {
	if o != nil && !IsNil(o.ServiceNowCMDBBusinessObject) {
		return true
	}

	return false
}

// SetServiceNowCMDBBusinessObject gets a reference to the given string and assigns it to the ServiceNowCMDBBusinessObject field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetServiceNowCMDBBusinessObject(v string) {
	o.ServiceNowCMDBBusinessObject = &v
}

// GetServiceNowCustomCmdbMapping returns the ServiceNowCustomCmdbMapping field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCustomCmdbMapping() string {
	if o == nil || IsNil(o.ServiceNowCustomCmdbMapping) {
		var ret string
		return ret
	}
	return *o.ServiceNowCustomCmdbMapping
}

// GetServiceNowCustomCmdbMappingOk returns a tuple with the ServiceNowCustomCmdbMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCustomCmdbMappingOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceNowCustomCmdbMapping) {
		return nil, false
	}
	return o.ServiceNowCustomCmdbMapping, true
}

// IsSetServiceNowCustomCmdbMapping returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) IsSetServiceNowCustomCmdbMapping() bool {
	if o != nil && !IsNil(o.ServiceNowCustomCmdbMapping) {
		return true
	}

	return false
}

// SetServiceNowCustomCmdbMapping gets a reference to the given string and assigns it to the ServiceNowCustomCmdbMapping field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetServiceNowCustomCmdbMapping(v string) {
	o.ServiceNowCustomCmdbMapping = &v
}

// GetServiceNowCmdbClassMapping returns the ServiceNowCmdbClassMapping field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCmdbClassMapping() []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigServiceNowCmdbClassMappingInner {
	if o == nil || IsNil(o.ServiceNowCmdbClassMapping) {
		var ret []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigServiceNowCmdbClassMappingInner
		return ret
	}
	return o.ServiceNowCmdbClassMapping
}

// GetServiceNowCmdbClassMappingOk returns a tuple with the ServiceNowCmdbClassMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetServiceNowCmdbClassMappingOk() ([]ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigServiceNowCmdbClassMappingInner, bool) {
	if o == nil || IsNil(o.ServiceNowCmdbClassMapping) {
		return nil, false
	}
	return o.ServiceNowCmdbClassMapping, true
}

// IsSetServiceNowCmdbClassMapping returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) IsSetServiceNowCmdbClassMapping() bool {
	if o != nil && !IsNil(o.ServiceNowCmdbClassMapping) {
		return true
	}

	return false
}

// SetServiceNowCmdbClassMapping gets a reference to the given []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigServiceNowCmdbClassMappingInner and assigns it to the ServiceNowCmdbClassMapping field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetServiceNowCmdbClassMapping(v []ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15ConfigServiceNowCmdbClassMappingInner) {
	o.ServiceNowCmdbClassMapping = v
}

// GetWebServiceImportUrl returns the WebServiceImportUrl field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceImportUrl() string {
	if o == nil || IsNil(o.WebServiceImportUrl) {
		var ret string
		return ret
	}
	return *o.WebServiceImportUrl
}

// GetWebServiceImportUrlOk returns a tuple with the WebServiceImportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceImportUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebServiceImportUrl) {
		return nil, false
	}
	return o.WebServiceImportUrl, true
}

// IsSetWebServiceImportUrl returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) IsSetWebServiceImportUrl() bool {
	if o != nil && !IsNil(o.WebServiceImportUrl) {
		return true
	}

	return false
}

// SetWebServiceImportUrl gets a reference to the given string and assigns it to the WebServiceImportUrl field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetWebServiceImportUrl(v string) {
	o.WebServiceImportUrl = &v
}

// GetWebServiceImportSysId returns the WebServiceImportSysId field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceImportSysId() string {
	if o == nil || IsNil(o.WebServiceImportSysId) {
		var ret string
		return ret
	}
	return *o.WebServiceImportSysId
}

// GetWebServiceImportSysIdOk returns a tuple with the WebServiceImportSysId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceImportSysIdOk() (*string, bool) {
	if o == nil || IsNil(o.WebServiceImportSysId) {
		return nil, false
	}
	return o.WebServiceImportSysId, true
}

// IsSetWebServiceImportSysId returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) IsSetWebServiceImportSysId() bool {
	if o != nil && !IsNil(o.WebServiceImportSysId) {
		return true
	}

	return false
}

// SetWebServiceImportSysId gets a reference to the given string and assigns it to the WebServiceImportSysId field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetWebServiceImportSysId(v string) {
	o.WebServiceImportSysId = &v
}

// GetWebServiceOperationUrl returns the WebServiceOperationUrl field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceOperationUrl() string {
	if o == nil || IsNil(o.WebServiceOperationUrl) {
		var ret string
		return ret
	}
	return *o.WebServiceOperationUrl
}

// GetWebServiceOperationUrlOk returns a tuple with the WebServiceOperationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetWebServiceOperationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebServiceOperationUrl) {
		return nil, false
	}
	return o.WebServiceOperationUrl, true
}

// IsSetWebServiceOperationUrl returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) IsSetWebServiceOperationUrl() bool {
	if o != nil && !IsNil(o.WebServiceOperationUrl) {
		return true
	}

	return false
}

// SetWebServiceOperationUrl gets a reference to the given string and assigns it to the WebServiceOperationUrl field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetWebServiceOperationUrl(v string) {
	o.WebServiceOperationUrl = &v
}

// GetCmdbMode returns the CmdbMode field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetCmdbMode() string {
	if o == nil || IsNil(o.CmdbMode) {
		var ret string
		return ret
	}
	return *o.CmdbMode
}

// GetCmdbModeOk returns a tuple with the CmdbMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetCmdbModeOk() (*string, bool) {
	if o == nil || IsNil(o.CmdbMode) {
		return nil, false
	}
	return o.CmdbMode, true
}

// IsSetCmdbMode returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) IsSetCmdbMode() bool {
	if o != nil && !IsNil(o.CmdbMode) {
		return true
	}

	return false
}

// SetCmdbMode gets a reference to the given string and assigns it to the CmdbMode field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetCmdbMode(v string) {
	o.CmdbMode = &v
}

// GetPreparedForSync returns the PreparedForSync field value if set, zero value otherwise.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetPreparedForSync() bool {
	if o == nil || IsNil(o.PreparedForSync) {
		var ret bool
		return ret
	}
	return *o.PreparedForSync
}

// GetPreparedForSyncOk returns a tuple with the PreparedForSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) GetPreparedForSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.PreparedForSync) {
		return nil, false
	}
	return o.PreparedForSync, true
}

// IsSetPreparedForSync returns a boolean if a field has been set.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) IsSetPreparedForSync() bool {
	if o != nil && !IsNil(o.PreparedForSync) {
		return true
	}

	return false
}

// SetPreparedForSync gets a reference to the given bool and assigns it to the PreparedForSync field.
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) SetPreparedForSync(v bool) {
	o.PreparedForSync = &v
}

func (o ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncidentAccess) {
		toSerialize["incidentAccess"] = o.IncidentAccess
	}
	if !IsNil(o.RequestAccess) {
		toSerialize["requestAccess"] = o.RequestAccess
	}
	if !IsNil(o.ServiceNowCMDBBusinessObject) {
		toSerialize["serviceNowCMDBBusinessObject"] = o.ServiceNowCMDBBusinessObject
	}
	if !IsNil(o.ServiceNowCustomCmdbMapping) {
		toSerialize["serviceNowCustomCmdbMapping"] = o.ServiceNowCustomCmdbMapping
	}
	if !IsNil(o.ServiceNowCmdbClassMapping) {
		toSerialize["serviceNowCmdbClassMapping"] = o.ServiceNowCmdbClassMapping
	}
	if !IsNil(o.WebServiceImportUrl) {
		toSerialize["webServiceImportUrl"] = o.WebServiceImportUrl
	}
	if !IsNil(o.WebServiceImportSysId) {
		toSerialize["webServiceImportSysId"] = o.WebServiceImportSysId
	}
	if !IsNil(o.WebServiceOperationUrl) {
		toSerialize["webServiceOperationUrl"] = o.WebServiceOperationUrl
	}
	if !IsNil(o.CmdbMode) {
		toSerialize["cmdbMode"] = o.CmdbMode
	}
	if !IsNil(o.PreparedForSync) {
		toSerialize["preparedForSync"] = o.PreparedForSync
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf15Config) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
