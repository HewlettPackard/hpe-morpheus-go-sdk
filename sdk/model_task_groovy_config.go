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

// checks if the TaskGroovyConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskGroovyConfig{}

// TaskGroovyConfig struct for TaskGroovyConfig
type TaskGroovyConfig struct {
	Id                   *int64                                                `json:"id,omitempty"`
	AccountId            *int64                                                `json:"accountId,omitempty"`
	Name                 *string                                               `json:"name,omitempty"`
	Code                 *string                                               `json:"code,omitempty"`
	TaskType             *ListTasks200ResponseAllOfTasksInnerAnyOf4TaskType    `json:"taskType,omitempty"`
	Labels               []string                                              `json:"labels,omitempty"`
	Visibility           *string                                               `json:"visibility,omitempty"`
	TaskOptions          *ListTasks200ResponseAllOfTasksInnerAnyOf4TaskOptions `json:"taskOptions,omitempty"`
	File                 *ListTasks200ResponseAllOfTasksInnerAnyOfFile         `json:"file,omitempty"`
	ResultType           *string                                               `json:"resultType,omitempty"`
	ExecuteTarget        *string                                               `json:"executeTarget,omitempty"`
	Retryable            *bool                                                 `json:"retryable,omitempty"`
	RetryCount           *int64                                                `json:"retryCount,omitempty"`
	RetryDelaySeconds    *int64                                                `json:"retryDelaySeconds,omitempty"`
	AllowCustomConfig    *bool                                                 `json:"allowCustomConfig,omitempty"`
	Credential           *ListClouds200ResponseAllOfZonesInnerCredentialAnyOf  `json:"credential,omitempty"`
	DateCreated          *time.Time                                            `json:"dateCreated,omitempty"`
	LastUpdated          *time.Time                                            `json:"lastUpdated,omitempty"`
	AdditionalProperties map[string]interface{}                                `json:",remain"`
}

type _TaskGroovyConfig TaskGroovyConfig

// NewTaskGroovyConfig instantiates a new TaskGroovyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskGroovyConfig() *TaskGroovyConfig {
	this := TaskGroovyConfig{}
	return &this
}

// NewTaskGroovyConfigWithDefaults instantiates a new TaskGroovyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskGroovyConfigWithDefaults() *TaskGroovyConfig {
	this := TaskGroovyConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TaskGroovyConfig) SetId(v int64) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetAccountId() int64 {
	if o == nil || IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// IsSetAccountId returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *TaskGroovyConfig) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskGroovyConfig) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TaskGroovyConfig) SetCode(v string) {
	o.Code = &v
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetTaskType() ListTasks200ResponseAllOfTasksInnerAnyOf4TaskType {
	if o == nil || IsNil(o.TaskType) {
		var ret ListTasks200ResponseAllOfTasksInnerAnyOf4TaskType
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetTaskTypeOk() (*ListTasks200ResponseAllOfTasksInnerAnyOf4TaskType, bool) {
	if o == nil || IsNil(o.TaskType) {
		return nil, false
	}
	return o.TaskType, true
}

// IsSetTaskType returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetTaskType() bool {
	if o != nil && !IsNil(o.TaskType) {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given ListTasks200ResponseAllOfTasksInnerAnyOf4TaskType and assigns it to the TaskType field.
func (o *TaskGroovyConfig) SetTaskType(v ListTasks200ResponseAllOfTasksInnerAnyOf4TaskType) {
	o.TaskType = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *TaskGroovyConfig) SetLabels(v []string) {
	o.Labels = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *TaskGroovyConfig) SetVisibility(v string) {
	o.Visibility = &v
}

// GetTaskOptions returns the TaskOptions field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetTaskOptions() ListTasks200ResponseAllOfTasksInnerAnyOf4TaskOptions {
	if o == nil || IsNil(o.TaskOptions) {
		var ret ListTasks200ResponseAllOfTasksInnerAnyOf4TaskOptions
		return ret
	}
	return *o.TaskOptions
}

// GetTaskOptionsOk returns a tuple with the TaskOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetTaskOptionsOk() (*ListTasks200ResponseAllOfTasksInnerAnyOf4TaskOptions, bool) {
	if o == nil || IsNil(o.TaskOptions) {
		return nil, false
	}
	return o.TaskOptions, true
}

// IsSetTaskOptions returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetTaskOptions() bool {
	if o != nil && !IsNil(o.TaskOptions) {
		return true
	}

	return false
}

// SetTaskOptions gets a reference to the given ListTasks200ResponseAllOfTasksInnerAnyOf4TaskOptions and assigns it to the TaskOptions field.
func (o *TaskGroovyConfig) SetTaskOptions(v ListTasks200ResponseAllOfTasksInnerAnyOf4TaskOptions) {
	o.TaskOptions = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetFile() ListTasks200ResponseAllOfTasksInnerAnyOfFile {
	if o == nil || IsNil(o.File) {
		var ret ListTasks200ResponseAllOfTasksInnerAnyOfFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetFileOk() (*ListTasks200ResponseAllOfTasksInnerAnyOfFile, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// IsSetFile returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given ListTasks200ResponseAllOfTasksInnerAnyOfFile and assigns it to the File field.
func (o *TaskGroovyConfig) SetFile(v ListTasks200ResponseAllOfTasksInnerAnyOfFile) {
	o.File = &v
}

// GetResultType returns the ResultType field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetResultType() string {
	if o == nil || IsNil(o.ResultType) {
		var ret string
		return ret
	}
	return *o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetResultTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultType) {
		return nil, false
	}
	return o.ResultType, true
}

// IsSetResultType returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetResultType() bool {
	if o != nil && !IsNil(o.ResultType) {
		return true
	}

	return false
}

// SetResultType gets a reference to the given string and assigns it to the ResultType field.
func (o *TaskGroovyConfig) SetResultType(v string) {
	o.ResultType = &v
}

// GetExecuteTarget returns the ExecuteTarget field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetExecuteTarget() string {
	if o == nil || IsNil(o.ExecuteTarget) {
		var ret string
		return ret
	}
	return *o.ExecuteTarget
}

// GetExecuteTargetOk returns a tuple with the ExecuteTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetExecuteTargetOk() (*string, bool) {
	if o == nil || IsNil(o.ExecuteTarget) {
		return nil, false
	}
	return o.ExecuteTarget, true
}

// IsSetExecuteTarget returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetExecuteTarget() bool {
	if o != nil && !IsNil(o.ExecuteTarget) {
		return true
	}

	return false
}

// SetExecuteTarget gets a reference to the given string and assigns it to the ExecuteTarget field.
func (o *TaskGroovyConfig) SetExecuteTarget(v string) {
	o.ExecuteTarget = &v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetRetryable() bool {
	if o == nil || IsNil(o.Retryable) {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetRetryableOk() (*bool, bool) {
	if o == nil || IsNil(o.Retryable) {
		return nil, false
	}
	return o.Retryable, true
}

// IsSetRetryable returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetRetryable() bool {
	if o != nil && !IsNil(o.Retryable) {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *TaskGroovyConfig) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetRetryCount() int64 {
	if o == nil || IsNil(o.RetryCount) {
		var ret int64
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetRetryCountOk() (*int64, bool) {
	if o == nil || IsNil(o.RetryCount) {
		return nil, false
	}
	return o.RetryCount, true
}

// IsSetRetryCount returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetRetryCount() bool {
	if o != nil && !IsNil(o.RetryCount) {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given int64 and assigns it to the RetryCount field.
func (o *TaskGroovyConfig) SetRetryCount(v int64) {
	o.RetryCount = &v
}

// GetRetryDelaySeconds returns the RetryDelaySeconds field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetRetryDelaySeconds() int64 {
	if o == nil || IsNil(o.RetryDelaySeconds) {
		var ret int64
		return ret
	}
	return *o.RetryDelaySeconds
}

// GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetRetryDelaySecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.RetryDelaySeconds) {
		return nil, false
	}
	return o.RetryDelaySeconds, true
}

// IsSetRetryDelaySeconds returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetRetryDelaySeconds() bool {
	if o != nil && !IsNil(o.RetryDelaySeconds) {
		return true
	}

	return false
}

// SetRetryDelaySeconds gets a reference to the given int64 and assigns it to the RetryDelaySeconds field.
func (o *TaskGroovyConfig) SetRetryDelaySeconds(v int64) {
	o.RetryDelaySeconds = &v
}

// GetAllowCustomConfig returns the AllowCustomConfig field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetAllowCustomConfig() bool {
	if o == nil || IsNil(o.AllowCustomConfig) {
		var ret bool
		return ret
	}
	return *o.AllowCustomConfig
}

// GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetAllowCustomConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCustomConfig) {
		return nil, false
	}
	return o.AllowCustomConfig, true
}

// IsSetAllowCustomConfig returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetAllowCustomConfig() bool {
	if o != nil && !IsNil(o.AllowCustomConfig) {
		return true
	}

	return false
}

// SetAllowCustomConfig gets a reference to the given bool and assigns it to the AllowCustomConfig field.
func (o *TaskGroovyConfig) SetAllowCustomConfig(v bool) {
	o.AllowCustomConfig = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetCredential() ListClouds200ResponseAllOfZonesInnerCredentialAnyOf {
	if o == nil || IsNil(o.Credential) {
		var ret ListClouds200ResponseAllOfZonesInnerCredentialAnyOf
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetCredentialOk() (*ListClouds200ResponseAllOfZonesInnerCredentialAnyOf, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// IsSetCredential returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given ListClouds200ResponseAllOfZonesInnerCredentialAnyOf and assigns it to the Credential field.
func (o *TaskGroovyConfig) SetCredential(v ListClouds200ResponseAllOfZonesInnerCredentialAnyOf) {
	o.Credential = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *TaskGroovyConfig) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *TaskGroovyConfig) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskGroovyConfig) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *TaskGroovyConfig) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *TaskGroovyConfig) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o TaskGroovyConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskGroovyConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.TaskType) {
		toSerialize["taskType"] = o.TaskType
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.TaskOptions) {
		toSerialize["taskOptions"] = o.TaskOptions
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.ResultType) {
		toSerialize["resultType"] = o.ResultType
	}
	if !IsNil(o.ExecuteTarget) {
		toSerialize["executeTarget"] = o.ExecuteTarget
	}
	if !IsNil(o.Retryable) {
		toSerialize["retryable"] = o.Retryable
	}
	if !IsNil(o.RetryCount) {
		toSerialize["retryCount"] = o.RetryCount
	}
	if !IsNil(o.RetryDelaySeconds) {
		toSerialize["retryDelaySeconds"] = o.RetryDelaySeconds
	}
	if !IsNil(o.AllowCustomConfig) {
		toSerialize["allowCustomConfig"] = o.AllowCustomConfig
	}
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *TaskGroovyConfig) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
