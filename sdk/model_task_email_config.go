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

// checks if the TaskEmailConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskEmailConfig{}

// TaskEmailConfig struct for TaskEmailConfig
type TaskEmailConfig struct {
	Id                   *int64                                                `json:"id,omitempty"`
	AccountId            *int64                                                `json:"accountId,omitempty"`
	Name                 *string                                               `json:"name,omitempty"`
	Code                 *string                                               `json:"code,omitempty"`
	TaskType             *ListTasks200ResponseAllOfTasksInnerAnyOf3TaskType    `json:"taskType,omitempty"`
	Labels               []string                                              `json:"labels,omitempty"`
	Visibility           *string                                               `json:"visibility,omitempty"`
	TaskOptions          *ListTasks200ResponseAllOfTasksInnerAnyOf3TaskOptions `json:"taskOptions,omitempty"`
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

type _TaskEmailConfig TaskEmailConfig

// NewTaskEmailConfig instantiates a new TaskEmailConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskEmailConfig() *TaskEmailConfig {
	this := TaskEmailConfig{}
	return &this
}

// NewTaskEmailConfigWithDefaults instantiates a new TaskEmailConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskEmailConfigWithDefaults() *TaskEmailConfig {
	this := TaskEmailConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TaskEmailConfig) SetId(v int64) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetAccountId() int64 {
	if o == nil || IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// IsSetAccountId returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *TaskEmailConfig) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskEmailConfig) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TaskEmailConfig) SetCode(v string) {
	o.Code = &v
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetTaskType() ListTasks200ResponseAllOfTasksInnerAnyOf3TaskType {
	if o == nil || IsNil(o.TaskType) {
		var ret ListTasks200ResponseAllOfTasksInnerAnyOf3TaskType
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetTaskTypeOk() (*ListTasks200ResponseAllOfTasksInnerAnyOf3TaskType, bool) {
	if o == nil || IsNil(o.TaskType) {
		return nil, false
	}
	return o.TaskType, true
}

// IsSetTaskType returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetTaskType() bool {
	if o != nil && !IsNil(o.TaskType) {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given ListTasks200ResponseAllOfTasksInnerAnyOf3TaskType and assigns it to the TaskType field.
func (o *TaskEmailConfig) SetTaskType(v ListTasks200ResponseAllOfTasksInnerAnyOf3TaskType) {
	o.TaskType = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *TaskEmailConfig) SetLabels(v []string) {
	o.Labels = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *TaskEmailConfig) SetVisibility(v string) {
	o.Visibility = &v
}

// GetTaskOptions returns the TaskOptions field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetTaskOptions() ListTasks200ResponseAllOfTasksInnerAnyOf3TaskOptions {
	if o == nil || IsNil(o.TaskOptions) {
		var ret ListTasks200ResponseAllOfTasksInnerAnyOf3TaskOptions
		return ret
	}
	return *o.TaskOptions
}

// GetTaskOptionsOk returns a tuple with the TaskOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetTaskOptionsOk() (*ListTasks200ResponseAllOfTasksInnerAnyOf3TaskOptions, bool) {
	if o == nil || IsNil(o.TaskOptions) {
		return nil, false
	}
	return o.TaskOptions, true
}

// IsSetTaskOptions returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetTaskOptions() bool {
	if o != nil && !IsNil(o.TaskOptions) {
		return true
	}

	return false
}

// SetTaskOptions gets a reference to the given ListTasks200ResponseAllOfTasksInnerAnyOf3TaskOptions and assigns it to the TaskOptions field.
func (o *TaskEmailConfig) SetTaskOptions(v ListTasks200ResponseAllOfTasksInnerAnyOf3TaskOptions) {
	o.TaskOptions = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetFile() ListTasks200ResponseAllOfTasksInnerAnyOfFile {
	if o == nil || IsNil(o.File) {
		var ret ListTasks200ResponseAllOfTasksInnerAnyOfFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetFileOk() (*ListTasks200ResponseAllOfTasksInnerAnyOfFile, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// IsSetFile returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given ListTasks200ResponseAllOfTasksInnerAnyOfFile and assigns it to the File field.
func (o *TaskEmailConfig) SetFile(v ListTasks200ResponseAllOfTasksInnerAnyOfFile) {
	o.File = &v
}

// GetResultType returns the ResultType field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetResultType() string {
	if o == nil || IsNil(o.ResultType) {
		var ret string
		return ret
	}
	return *o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetResultTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultType) {
		return nil, false
	}
	return o.ResultType, true
}

// IsSetResultType returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetResultType() bool {
	if o != nil && !IsNil(o.ResultType) {
		return true
	}

	return false
}

// SetResultType gets a reference to the given string and assigns it to the ResultType field.
func (o *TaskEmailConfig) SetResultType(v string) {
	o.ResultType = &v
}

// GetExecuteTarget returns the ExecuteTarget field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetExecuteTarget() string {
	if o == nil || IsNil(o.ExecuteTarget) {
		var ret string
		return ret
	}
	return *o.ExecuteTarget
}

// GetExecuteTargetOk returns a tuple with the ExecuteTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetExecuteTargetOk() (*string, bool) {
	if o == nil || IsNil(o.ExecuteTarget) {
		return nil, false
	}
	return o.ExecuteTarget, true
}

// IsSetExecuteTarget returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetExecuteTarget() bool {
	if o != nil && !IsNil(o.ExecuteTarget) {
		return true
	}

	return false
}

// SetExecuteTarget gets a reference to the given string and assigns it to the ExecuteTarget field.
func (o *TaskEmailConfig) SetExecuteTarget(v string) {
	o.ExecuteTarget = &v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetRetryable() bool {
	if o == nil || IsNil(o.Retryable) {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetRetryableOk() (*bool, bool) {
	if o == nil || IsNil(o.Retryable) {
		return nil, false
	}
	return o.Retryable, true
}

// IsSetRetryable returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetRetryable() bool {
	if o != nil && !IsNil(o.Retryable) {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *TaskEmailConfig) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetRetryCount() int64 {
	if o == nil || IsNil(o.RetryCount) {
		var ret int64
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetRetryCountOk() (*int64, bool) {
	if o == nil || IsNil(o.RetryCount) {
		return nil, false
	}
	return o.RetryCount, true
}

// IsSetRetryCount returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetRetryCount() bool {
	if o != nil && !IsNil(o.RetryCount) {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given int64 and assigns it to the RetryCount field.
func (o *TaskEmailConfig) SetRetryCount(v int64) {
	o.RetryCount = &v
}

// GetRetryDelaySeconds returns the RetryDelaySeconds field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetRetryDelaySeconds() int64 {
	if o == nil || IsNil(o.RetryDelaySeconds) {
		var ret int64
		return ret
	}
	return *o.RetryDelaySeconds
}

// GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetRetryDelaySecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.RetryDelaySeconds) {
		return nil, false
	}
	return o.RetryDelaySeconds, true
}

// IsSetRetryDelaySeconds returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetRetryDelaySeconds() bool {
	if o != nil && !IsNil(o.RetryDelaySeconds) {
		return true
	}

	return false
}

// SetRetryDelaySeconds gets a reference to the given int64 and assigns it to the RetryDelaySeconds field.
func (o *TaskEmailConfig) SetRetryDelaySeconds(v int64) {
	o.RetryDelaySeconds = &v
}

// GetAllowCustomConfig returns the AllowCustomConfig field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetAllowCustomConfig() bool {
	if o == nil || IsNil(o.AllowCustomConfig) {
		var ret bool
		return ret
	}
	return *o.AllowCustomConfig
}

// GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetAllowCustomConfigOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCustomConfig) {
		return nil, false
	}
	return o.AllowCustomConfig, true
}

// IsSetAllowCustomConfig returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetAllowCustomConfig() bool {
	if o != nil && !IsNil(o.AllowCustomConfig) {
		return true
	}

	return false
}

// SetAllowCustomConfig gets a reference to the given bool and assigns it to the AllowCustomConfig field.
func (o *TaskEmailConfig) SetAllowCustomConfig(v bool) {
	o.AllowCustomConfig = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetCredential() ListClouds200ResponseAllOfZonesInnerCredentialAnyOf {
	if o == nil || IsNil(o.Credential) {
		var ret ListClouds200ResponseAllOfZonesInnerCredentialAnyOf
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetCredentialOk() (*ListClouds200ResponseAllOfZonesInnerCredentialAnyOf, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// IsSetCredential returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given ListClouds200ResponseAllOfZonesInnerCredentialAnyOf and assigns it to the Credential field.
func (o *TaskEmailConfig) SetCredential(v ListClouds200ResponseAllOfZonesInnerCredentialAnyOf) {
	o.Credential = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *TaskEmailConfig) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *TaskEmailConfig) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEmailConfig) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *TaskEmailConfig) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *TaskEmailConfig) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o TaskEmailConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskEmailConfig) ToMap() (map[string]interface{}, error) {
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
func (o *TaskEmailConfig) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
