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

// checks if the ListJobs200ResponseAllOfJobsInnerAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListJobs200ResponseAllOfJobsInnerAnyOf1{}

// ListJobs200ResponseAllOfJobsInnerAnyOf1 struct for ListJobs200ResponseAllOfJobsInnerAnyOf1
type ListJobs200ResponseAllOfJobsInnerAnyOf1 struct {
	Id                   *int64                                                      `json:"id,omitempty"`
	Name                 *string                                                     `json:"name,omitempty"`
	Labels               []string                                                    `json:"labels,omitempty"`
	Type                 *ListBackupSettings200ResponseBackupSettingsDefaultSchedule `json:"type,omitempty"`
	Task                 *GetAlerts200ResponseAllOfChecksInnerAccount                `json:"task,omitempty"`
	JobSummary           *string                                                     `json:"jobSummary,omitempty"`
	ScheduleMode         *ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode         `json:"scheduleMode,omitempty"`
	DateTime             *string                                                     `json:"dateTime,omitempty"`
	Status               *string                                                     `json:"status,omitempty"`
	Namespace            *string                                                     `json:"namespace,omitempty"`
	Category             *string                                                     `json:"category,omitempty"`
	Description          *string                                                     `json:"description,omitempty"`
	Enabled              *bool                                                       `json:"enabled,omitempty"`
	DateCreated          *time.Time                                                  `json:"dateCreated,omitempty"`
	LastUpdated          *time.Time                                                  `json:"lastUpdated,omitempty"`
	LastRun              *time.Time                                                  `json:"lastRun,omitempty"`
	LastResult           *string                                                     `json:"lastResult,omitempty"`
	CreatedBy            *ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy            `json:"createdBy,omitempty"`
	TargetType           *string                                                     `json:"targetType,omitempty"`
	Targets              []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner        `json:"targets,omitempty"`
	CustomConfig         *string                                                     `json:"customConfig,omitempty"`
	CustomOptions        *ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions        `json:"customOptions,omitempty"`
	AdditionalProperties map[string]interface{}                                      `json:",remain"`
}

type _ListJobs200ResponseAllOfJobsInnerAnyOf1 ListJobs200ResponseAllOfJobsInnerAnyOf1

// NewListJobs200ResponseAllOfJobsInnerAnyOf1 instantiates a new ListJobs200ResponseAllOfJobsInnerAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListJobs200ResponseAllOfJobsInnerAnyOf1() *ListJobs200ResponseAllOfJobsInnerAnyOf1 {
	this := ListJobs200ResponseAllOfJobsInnerAnyOf1{}
	return &this
}

// NewListJobs200ResponseAllOfJobsInnerAnyOf1WithDefaults instantiates a new ListJobs200ResponseAllOfJobsInnerAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListJobs200ResponseAllOfJobsInnerAnyOf1WithDefaults() *ListJobs200ResponseAllOfJobsInnerAnyOf1 {
	this := ListJobs200ResponseAllOfJobsInnerAnyOf1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetLabels(v []string) {
	o.Labels = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule {
	if o == nil || IsNil(o.Type) {
		var ret ListBackupSettings200ResponseBackupSettingsDefaultSchedule
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ListBackupSettings200ResponseBackupSettingsDefaultSchedule and assigns it to the Type field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule) {
	o.Type = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetTask() GetAlerts200ResponseAllOfChecksInnerAccount {
	if o == nil || IsNil(o.Task) {
		var ret GetAlerts200ResponseAllOfChecksInnerAccount
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetTaskOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// IsSetTask returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given GetAlerts200ResponseAllOfChecksInnerAccount and assigns it to the Task field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetTask(v GetAlerts200ResponseAllOfChecksInnerAccount) {
	o.Task = &v
}

// GetJobSummary returns the JobSummary field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetJobSummary() string {
	if o == nil || IsNil(o.JobSummary) {
		var ret string
		return ret
	}
	return *o.JobSummary
}

// GetJobSummaryOk returns a tuple with the JobSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetJobSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.JobSummary) {
		return nil, false
	}
	return o.JobSummary, true
}

// IsSetJobSummary returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetJobSummary() bool {
	if o != nil && !IsNil(o.JobSummary) {
		return true
	}

	return false
}

// SetJobSummary gets a reference to the given string and assigns it to the JobSummary field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetJobSummary(v string) {
	o.JobSummary = &v
}

// GetScheduleMode returns the ScheduleMode field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetScheduleMode() ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode {
	if o == nil || IsNil(o.ScheduleMode) {
		var ret ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode
		return ret
	}
	return *o.ScheduleMode
}

// GetScheduleModeOk returns a tuple with the ScheduleMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetScheduleModeOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode, bool) {
	if o == nil || IsNil(o.ScheduleMode) {
		return nil, false
	}
	return o.ScheduleMode, true
}

// IsSetScheduleMode returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetScheduleMode() bool {
	if o != nil && !IsNil(o.ScheduleMode) {
		return true
	}

	return false
}

// SetScheduleMode gets a reference to the given ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode and assigns it to the ScheduleMode field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetScheduleMode(v ListJobs200ResponseAllOfJobsInnerAnyOfScheduleMode) {
	o.ScheduleMode = &v
}

// GetDateTime returns the DateTime field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetDateTime() string {
	if o == nil || IsNil(o.DateTime) {
		var ret string
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DateTime) {
		return nil, false
	}
	return o.DateTime, true
}

// IsSetDateTime returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetDateTime() bool {
	if o != nil && !IsNil(o.DateTime) {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given string and assigns it to the DateTime field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetDateTime(v string) {
	o.DateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// IsSetStatus returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetStatus(v string) {
	o.Status = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// IsSetNamespace returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetNamespace(v string) {
	o.Namespace = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// IsSetCategory returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// IsSetEnabled returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastRun returns the LastRun field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetLastRun() time.Time {
	if o == nil || IsNil(o.LastRun) {
		var ret time.Time
		return ret
	}
	return *o.LastRun
}

// GetLastRunOk returns a tuple with the LastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetLastRunOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastRun) {
		return nil, false
	}
	return o.LastRun, true
}

// IsSetLastRun returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetLastRun() bool {
	if o != nil && !IsNil(o.LastRun) {
		return true
	}

	return false
}

// SetLastRun gets a reference to the given time.Time and assigns it to the LastRun field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetLastRun(v time.Time) {
	o.LastRun = &v
}

// GetLastResult returns the LastResult field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetLastResult() string {
	if o == nil || IsNil(o.LastResult) {
		var ret string
		return ret
	}
	return *o.LastResult
}

// GetLastResultOk returns a tuple with the LastResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetLastResultOk() (*string, bool) {
	if o == nil || IsNil(o.LastResult) {
		return nil, false
	}
	return o.LastResult, true
}

// IsSetLastResult returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetLastResult() bool {
	if o != nil && !IsNil(o.LastResult) {
		return true
	}

	return false
}

// SetLastResult gets a reference to the given string and assigns it to the LastResult field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetLastResult(v string) {
	o.LastResult = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetCreatedBy() ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy {
	if o == nil || IsNil(o.CreatedBy) {
		var ret ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetCreatedByOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// IsSetCreatedBy returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy and assigns it to the CreatedBy field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetCreatedBy(v ListJobs200ResponseAllOfJobsInnerAnyOfCreatedBy) {
	o.CreatedBy = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// IsSetTargetType returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetTargetType(v string) {
	o.TargetType = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetTargets() []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner {
	if o == nil || IsNil(o.Targets) {
		var ret []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetTargetsOk() ([]ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// IsSetTargets returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner and assigns it to the Targets field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetTargets(v []ListJobs200ResponseAllOfJobsInnerAnyOfTargetsInner) {
	o.Targets = v
}

// GetCustomConfig returns the CustomConfig field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetCustomConfig() string {
	if o == nil || IsNil(o.CustomConfig) {
		var ret string
		return ret
	}
	return *o.CustomConfig
}

// GetCustomConfigOk returns a tuple with the CustomConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetCustomConfigOk() (*string, bool) {
	if o == nil || IsNil(o.CustomConfig) {
		return nil, false
	}
	return o.CustomConfig, true
}

// IsSetCustomConfig returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetCustomConfig() bool {
	if o != nil && !IsNil(o.CustomConfig) {
		return true
	}

	return false
}

// SetCustomConfig gets a reference to the given string and assigns it to the CustomConfig field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetCustomConfig(v string) {
	o.CustomConfig = &v
}

// GetCustomOptions returns the CustomOptions field value if set, zero value otherwise.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetCustomOptions() ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions {
	if o == nil || IsNil(o.CustomOptions) {
		var ret ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions
		return ret
	}
	return *o.CustomOptions
}

// GetCustomOptionsOk returns a tuple with the CustomOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) GetCustomOptionsOk() (*ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions, bool) {
	if o == nil || IsNil(o.CustomOptions) {
		return nil, false
	}
	return o.CustomOptions, true
}

// IsSetCustomOptions returns a boolean if a field has been set.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) IsSetCustomOptions() bool {
	if o != nil && !IsNil(o.CustomOptions) {
		return true
	}

	return false
}

// SetCustomOptions gets a reference to the given ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions and assigns it to the CustomOptions field.
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) SetCustomOptions(v ListJobs200ResponseAllOfJobsInnerAnyOfCustomOptions) {
	o.CustomOptions = &v
}

func (o ListJobs200ResponseAllOfJobsInnerAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListJobs200ResponseAllOfJobsInnerAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	if !IsNil(o.JobSummary) {
		toSerialize["jobSummary"] = o.JobSummary
	}
	if !IsNil(o.ScheduleMode) {
		toSerialize["scheduleMode"] = o.ScheduleMode
	}
	if !IsNil(o.DateTime) {
		toSerialize["dateTime"] = o.DateTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.LastRun) {
		toSerialize["lastRun"] = o.LastRun
	}
	if !IsNil(o.LastResult) {
		toSerialize["lastResult"] = o.LastResult
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.TargetType) {
		toSerialize["targetType"] = o.TargetType
	}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	if !IsNil(o.CustomConfig) {
		toSerialize["customConfig"] = o.CustomConfig
	}
	if !IsNil(o.CustomOptions) {
		toSerialize["customOptions"] = o.CustomOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListJobs200ResponseAllOfJobsInnerAnyOf1) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
