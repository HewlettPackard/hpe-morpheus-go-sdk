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
	"time"
)

// checks if the Health type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Health{}

// Health struct for Health
type Health struct {
	Success              *bool                                     `json:"success,omitempty"`
	StatusMessage        *string                                   `json:"statusMessage,omitempty"`
	ApplianceUrl         *string                                   `json:"applianceUrl,omitempty"`
	BuildVersion         *string                                   `json:"buildVersion,omitempty"`
	Uuid                 *string                                   `json:"uuid,omitempty"`
	SetupNeeded          *bool                                     `json:"setupNeeded,omitempty"`
	Date                 *time.Time                                `json:"date,omitempty"`
	Cpu                  *ListHealth200ResponseAllOfHealthCpu      `json:"cpu,omitempty"`
	Memory               *ListHealth200ResponseAllOfHealthMemory   `json:"memory,omitempty"`
	Threads              *ListHealth200ResponseAllOfHealthThreads  `json:"threads,omitempty"`
	Database             *ListHealth200ResponseAllOfHealthDatabase `json:"database,omitempty"`
	Elastic              *ListHealth200ResponseAllOfHealthElastic  `json:"elastic,omitempty"`
	Rabbit               *ListHealth200ResponseAllOfHealthRabbit   `json:"rabbit,omitempty"`
	AdditionalProperties map[string]interface{}                    `json:",remain"`
}

type _Health Health

// NewHealth instantiates a new Health object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealth() *Health {
	this := Health{}
	return &this
}

// NewHealthWithDefaults instantiates a new Health object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthWithDefaults() *Health {
	this := Health{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *Health) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *Health) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *Health) SetSuccess(v bool) {
	o.Success = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *Health) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// IsSetStatusMessage returns a boolean if a field has been set.
func (o *Health) IsSetStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *Health) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetApplianceUrl returns the ApplianceUrl field value if set, zero value otherwise.
func (o *Health) GetApplianceUrl() string {
	if o == nil || IsNil(o.ApplianceUrl) {
		var ret string
		return ret
	}
	return *o.ApplianceUrl
}

// GetApplianceUrlOk returns a tuple with the ApplianceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetApplianceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ApplianceUrl) {
		return nil, false
	}
	return o.ApplianceUrl, true
}

// IsSetApplianceUrl returns a boolean if a field has been set.
func (o *Health) IsSetApplianceUrl() bool {
	if o != nil && !IsNil(o.ApplianceUrl) {
		return true
	}

	return false
}

// SetApplianceUrl gets a reference to the given string and assigns it to the ApplianceUrl field.
func (o *Health) SetApplianceUrl(v string) {
	o.ApplianceUrl = &v
}

// GetBuildVersion returns the BuildVersion field value if set, zero value otherwise.
func (o *Health) GetBuildVersion() string {
	if o == nil || IsNil(o.BuildVersion) {
		var ret string
		return ret
	}
	return *o.BuildVersion
}

// GetBuildVersionOk returns a tuple with the BuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetBuildVersionOk() (*string, bool) {
	if o == nil || IsNil(o.BuildVersion) {
		return nil, false
	}
	return o.BuildVersion, true
}

// IsSetBuildVersion returns a boolean if a field has been set.
func (o *Health) IsSetBuildVersion() bool {
	if o != nil && !IsNil(o.BuildVersion) {
		return true
	}

	return false
}

// SetBuildVersion gets a reference to the given string and assigns it to the BuildVersion field.
func (o *Health) SetBuildVersion(v string) {
	o.BuildVersion = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Health) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// IsSetUuid returns a boolean if a field has been set.
func (o *Health) IsSetUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Health) SetUuid(v string) {
	o.Uuid = &v
}

// GetSetupNeeded returns the SetupNeeded field value if set, zero value otherwise.
func (o *Health) GetSetupNeeded() bool {
	if o == nil || IsNil(o.SetupNeeded) {
		var ret bool
		return ret
	}
	return *o.SetupNeeded
}

// GetSetupNeededOk returns a tuple with the SetupNeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetSetupNeededOk() (*bool, bool) {
	if o == nil || IsNil(o.SetupNeeded) {
		return nil, false
	}
	return o.SetupNeeded, true
}

// IsSetSetupNeeded returns a boolean if a field has been set.
func (o *Health) IsSetSetupNeeded() bool {
	if o != nil && !IsNil(o.SetupNeeded) {
		return true
	}

	return false
}

// SetSetupNeeded gets a reference to the given bool and assigns it to the SetupNeeded field.
func (o *Health) SetSetupNeeded(v bool) {
	o.SetupNeeded = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Health) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// IsSetDate returns a boolean if a field has been set.
func (o *Health) IsSetDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *Health) SetDate(v time.Time) {
	o.Date = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Health) GetCpu() ListHealth200ResponseAllOfHealthCpu {
	if o == nil || IsNil(o.Cpu) {
		var ret ListHealth200ResponseAllOfHealthCpu
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetCpuOk() (*ListHealth200ResponseAllOfHealthCpu, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// IsSetCpu returns a boolean if a field has been set.
func (o *Health) IsSetCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given ListHealth200ResponseAllOfHealthCpu and assigns it to the Cpu field.
func (o *Health) SetCpu(v ListHealth200ResponseAllOfHealthCpu) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Health) GetMemory() ListHealth200ResponseAllOfHealthMemory {
	if o == nil || IsNil(o.Memory) {
		var ret ListHealth200ResponseAllOfHealthMemory
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetMemoryOk() (*ListHealth200ResponseAllOfHealthMemory, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// IsSetMemory returns a boolean if a field has been set.
func (o *Health) IsSetMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given ListHealth200ResponseAllOfHealthMemory and assigns it to the Memory field.
func (o *Health) SetMemory(v ListHealth200ResponseAllOfHealthMemory) {
	o.Memory = &v
}

// GetThreads returns the Threads field value if set, zero value otherwise.
func (o *Health) GetThreads() ListHealth200ResponseAllOfHealthThreads {
	if o == nil || IsNil(o.Threads) {
		var ret ListHealth200ResponseAllOfHealthThreads
		return ret
	}
	return *o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetThreadsOk() (*ListHealth200ResponseAllOfHealthThreads, bool) {
	if o == nil || IsNil(o.Threads) {
		return nil, false
	}
	return o.Threads, true
}

// IsSetThreads returns a boolean if a field has been set.
func (o *Health) IsSetThreads() bool {
	if o != nil && !IsNil(o.Threads) {
		return true
	}

	return false
}

// SetThreads gets a reference to the given ListHealth200ResponseAllOfHealthThreads and assigns it to the Threads field.
func (o *Health) SetThreads(v ListHealth200ResponseAllOfHealthThreads) {
	o.Threads = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *Health) GetDatabase() ListHealth200ResponseAllOfHealthDatabase {
	if o == nil || IsNil(o.Database) {
		var ret ListHealth200ResponseAllOfHealthDatabase
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetDatabaseOk() (*ListHealth200ResponseAllOfHealthDatabase, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// IsSetDatabase returns a boolean if a field has been set.
func (o *Health) IsSetDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given ListHealth200ResponseAllOfHealthDatabase and assigns it to the Database field.
func (o *Health) SetDatabase(v ListHealth200ResponseAllOfHealthDatabase) {
	o.Database = &v
}

// GetElastic returns the Elastic field value if set, zero value otherwise.
func (o *Health) GetElastic() ListHealth200ResponseAllOfHealthElastic {
	if o == nil || IsNil(o.Elastic) {
		var ret ListHealth200ResponseAllOfHealthElastic
		return ret
	}
	return *o.Elastic
}

// GetElasticOk returns a tuple with the Elastic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetElasticOk() (*ListHealth200ResponseAllOfHealthElastic, bool) {
	if o == nil || IsNil(o.Elastic) {
		return nil, false
	}
	return o.Elastic, true
}

// IsSetElastic returns a boolean if a field has been set.
func (o *Health) IsSetElastic() bool {
	if o != nil && !IsNil(o.Elastic) {
		return true
	}

	return false
}

// SetElastic gets a reference to the given ListHealth200ResponseAllOfHealthElastic and assigns it to the Elastic field.
func (o *Health) SetElastic(v ListHealth200ResponseAllOfHealthElastic) {
	o.Elastic = &v
}

// GetRabbit returns the Rabbit field value if set, zero value otherwise.
func (o *Health) GetRabbit() ListHealth200ResponseAllOfHealthRabbit {
	if o == nil || IsNil(o.Rabbit) {
		var ret ListHealth200ResponseAllOfHealthRabbit
		return ret
	}
	return *o.Rabbit
}

// GetRabbitOk returns a tuple with the Rabbit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetRabbitOk() (*ListHealth200ResponseAllOfHealthRabbit, bool) {
	if o == nil || IsNil(o.Rabbit) {
		return nil, false
	}
	return o.Rabbit, true
}

// IsSetRabbit returns a boolean if a field has been set.
func (o *Health) IsSetRabbit() bool {
	if o != nil && !IsNil(o.Rabbit) {
		return true
	}

	return false
}

// SetRabbit gets a reference to the given ListHealth200ResponseAllOfHealthRabbit and assigns it to the Rabbit field.
func (o *Health) SetRabbit(v ListHealth200ResponseAllOfHealthRabbit) {
	o.Rabbit = &v
}

func (o Health) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Health) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	if !IsNil(o.ApplianceUrl) {
		toSerialize["applianceUrl"] = o.ApplianceUrl
	}
	if !IsNil(o.BuildVersion) {
		toSerialize["buildVersion"] = o.BuildVersion
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.SetupNeeded) {
		toSerialize["setupNeeded"] = o.SetupNeeded
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.Threads) {
		toSerialize["threads"] = o.Threads
	}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.Elastic) {
		toSerialize["elastic"] = o.Elastic
	}
	if !IsNil(o.Rabbit) {
		toSerialize["rabbit"] = o.Rabbit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *Health) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
