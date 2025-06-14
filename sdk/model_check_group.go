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

// checks if the CheckGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckGroup{}

// CheckGroup struct for CheckGroup
type CheckGroup struct {
	Id                   *int64                                                      `json:"id,omitempty"`
	Account              *GetAlerts200ResponseAllOfChecksInnerAccount                `json:"account,omitempty"`
	Instance             *GetAlerts200ResponseAllOfCheckGroupsInnerInstance          `json:"instance,omitempty"`
	Name                 *string                                                     `json:"name,omitempty"`
	Description          *string                                                     `json:"description,omitempty"`
	InUptime             *bool                                                       `json:"inUptime,omitempty"`
	LastCheckStatus      *string                                                     `json:"lastCheckStatus,omitempty"`
	LastWarningDate      *time.Time                                                  `json:"lastWarningDate,omitempty"`
	LastErrorDate        *time.Time                                                  `json:"lastErrorDate,omitempty"`
	LastSuccessDate      *time.Time                                                  `json:"lastSuccessDate,omitempty"`
	LastRunDate          *time.Time                                                  `json:"lastRunDate,omitempty"`
	LastError            *string                                                     `json:"lastError,omitempty"`
	OutageTime           *int64                                                      `json:"outageTime,omitempty"`
	LastTimer            *int64                                                      `json:"lastTimer,omitempty"`
	Health               *int64                                                      `json:"health,omitempty"`
	History              *string                                                     `json:"history,omitempty"`
	MinHappy             *int64                                                      `json:"minHappy,omitempty"`
	LastMetric           *string                                                     `json:"lastMetric,omitempty"`
	Severity             *string                                                     `json:"severity,omitempty"`
	CreateIncident       *bool                                                       `json:"createIncident,omitempty"`
	Muted                *bool                                                       `json:"muted,omitempty"`
	CreatedBy            *ListActivity200ResponseAllOfActivityInnerActivityInnerUser `json:"createdBy,omitempty"`
	DateCreated          *time.Time                                                  `json:"dateCreated,omitempty"`
	LastUpdated          *time.Time                                                  `json:"lastUpdated,omitempty"`
	Availability         *float32                                                    `json:"availability,omitempty"`
	CheckType            *GetAlerts200ResponseAllOfChecksInnerCheckType              `json:"checkType,omitempty"`
	Checks               []int64                                                     `json:"checks,omitempty"`
	AdditionalProperties map[string]interface{}                                      `json:",remain"`
}

type _CheckGroup CheckGroup

// NewCheckGroup instantiates a new CheckGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckGroup() *CheckGroup {
	this := CheckGroup{}
	return &this
}

// NewCheckGroupWithDefaults instantiates a new CheckGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckGroupWithDefaults() *CheckGroup {
	this := CheckGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CheckGroup) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *CheckGroup) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CheckGroup) SetId(v int64) {
	o.Id = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *CheckGroup) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount {
	if o == nil || IsNil(o.Account) {
		var ret GetAlerts200ResponseAllOfChecksInnerAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// IsSetAccount returns a boolean if a field has been set.
func (o *CheckGroup) IsSetAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given GetAlerts200ResponseAllOfChecksInnerAccount and assigns it to the Account field.
func (o *CheckGroup) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount) {
	o.Account = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *CheckGroup) GetInstance() GetAlerts200ResponseAllOfCheckGroupsInnerInstance {
	if o == nil || IsNil(o.Instance) {
		var ret GetAlerts200ResponseAllOfCheckGroupsInnerInstance
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetInstanceOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// IsSetInstance returns a boolean if a field has been set.
func (o *CheckGroup) IsSetInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given GetAlerts200ResponseAllOfCheckGroupsInnerInstance and assigns it to the Instance field.
func (o *CheckGroup) SetInstance(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance) {
	o.Instance = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CheckGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *CheckGroup) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CheckGroup) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CheckGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *CheckGroup) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CheckGroup) SetDescription(v string) {
	o.Description = &v
}

// GetInUptime returns the InUptime field value if set, zero value otherwise.
func (o *CheckGroup) GetInUptime() bool {
	if o == nil || IsNil(o.InUptime) {
		var ret bool
		return ret
	}
	return *o.InUptime
}

// GetInUptimeOk returns a tuple with the InUptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetInUptimeOk() (*bool, bool) {
	if o == nil || IsNil(o.InUptime) {
		return nil, false
	}
	return o.InUptime, true
}

// IsSetInUptime returns a boolean if a field has been set.
func (o *CheckGroup) IsSetInUptime() bool {
	if o != nil && !IsNil(o.InUptime) {
		return true
	}

	return false
}

// SetInUptime gets a reference to the given bool and assigns it to the InUptime field.
func (o *CheckGroup) SetInUptime(v bool) {
	o.InUptime = &v
}

// GetLastCheckStatus returns the LastCheckStatus field value if set, zero value otherwise.
func (o *CheckGroup) GetLastCheckStatus() string {
	if o == nil || IsNil(o.LastCheckStatus) {
		var ret string
		return ret
	}
	return *o.LastCheckStatus
}

// GetLastCheckStatusOk returns a tuple with the LastCheckStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetLastCheckStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LastCheckStatus) {
		return nil, false
	}
	return o.LastCheckStatus, true
}

// IsSetLastCheckStatus returns a boolean if a field has been set.
func (o *CheckGroup) IsSetLastCheckStatus() bool {
	if o != nil && !IsNil(o.LastCheckStatus) {
		return true
	}

	return false
}

// SetLastCheckStatus gets a reference to the given string and assigns it to the LastCheckStatus field.
func (o *CheckGroup) SetLastCheckStatus(v string) {
	o.LastCheckStatus = &v
}

// GetLastWarningDate returns the LastWarningDate field value if set, zero value otherwise.
func (o *CheckGroup) GetLastWarningDate() time.Time {
	if o == nil || IsNil(o.LastWarningDate) {
		var ret time.Time
		return ret
	}
	return *o.LastWarningDate
}

// GetLastWarningDateOk returns a tuple with the LastWarningDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetLastWarningDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastWarningDate) {
		return nil, false
	}
	return o.LastWarningDate, true
}

// IsSetLastWarningDate returns a boolean if a field has been set.
func (o *CheckGroup) IsSetLastWarningDate() bool {
	if o != nil && !IsNil(o.LastWarningDate) {
		return true
	}

	return false
}

// SetLastWarningDate gets a reference to the given time.Time and assigns it to the LastWarningDate field.
func (o *CheckGroup) SetLastWarningDate(v time.Time) {
	o.LastWarningDate = &v
}

// GetLastErrorDate returns the LastErrorDate field value if set, zero value otherwise.
func (o *CheckGroup) GetLastErrorDate() time.Time {
	if o == nil || IsNil(o.LastErrorDate) {
		var ret time.Time
		return ret
	}
	return *o.LastErrorDate
}

// GetLastErrorDateOk returns a tuple with the LastErrorDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetLastErrorDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastErrorDate) {
		return nil, false
	}
	return o.LastErrorDate, true
}

// IsSetLastErrorDate returns a boolean if a field has been set.
func (o *CheckGroup) IsSetLastErrorDate() bool {
	if o != nil && !IsNil(o.LastErrorDate) {
		return true
	}

	return false
}

// SetLastErrorDate gets a reference to the given time.Time and assigns it to the LastErrorDate field.
func (o *CheckGroup) SetLastErrorDate(v time.Time) {
	o.LastErrorDate = &v
}

// GetLastSuccessDate returns the LastSuccessDate field value if set, zero value otherwise.
func (o *CheckGroup) GetLastSuccessDate() time.Time {
	if o == nil || IsNil(o.LastSuccessDate) {
		var ret time.Time
		return ret
	}
	return *o.LastSuccessDate
}

// GetLastSuccessDateOk returns a tuple with the LastSuccessDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetLastSuccessDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSuccessDate) {
		return nil, false
	}
	return o.LastSuccessDate, true
}

// IsSetLastSuccessDate returns a boolean if a field has been set.
func (o *CheckGroup) IsSetLastSuccessDate() bool {
	if o != nil && !IsNil(o.LastSuccessDate) {
		return true
	}

	return false
}

// SetLastSuccessDate gets a reference to the given time.Time and assigns it to the LastSuccessDate field.
func (o *CheckGroup) SetLastSuccessDate(v time.Time) {
	o.LastSuccessDate = &v
}

// GetLastRunDate returns the LastRunDate field value if set, zero value otherwise.
func (o *CheckGroup) GetLastRunDate() time.Time {
	if o == nil || IsNil(o.LastRunDate) {
		var ret time.Time
		return ret
	}
	return *o.LastRunDate
}

// GetLastRunDateOk returns a tuple with the LastRunDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetLastRunDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastRunDate) {
		return nil, false
	}
	return o.LastRunDate, true
}

// IsSetLastRunDate returns a boolean if a field has been set.
func (o *CheckGroup) IsSetLastRunDate() bool {
	if o != nil && !IsNil(o.LastRunDate) {
		return true
	}

	return false
}

// SetLastRunDate gets a reference to the given time.Time and assigns it to the LastRunDate field.
func (o *CheckGroup) SetLastRunDate(v time.Time) {
	o.LastRunDate = &v
}

// GetLastError returns the LastError field value if set, zero value otherwise.
func (o *CheckGroup) GetLastError() string {
	if o == nil || IsNil(o.LastError) {
		var ret string
		return ret
	}
	return *o.LastError
}

// GetLastErrorOk returns a tuple with the LastError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetLastErrorOk() (*string, bool) {
	if o == nil || IsNil(o.LastError) {
		return nil, false
	}
	return o.LastError, true
}

// IsSetLastError returns a boolean if a field has been set.
func (o *CheckGroup) IsSetLastError() bool {
	if o != nil && !IsNil(o.LastError) {
		return true
	}

	return false
}

// SetLastError gets a reference to the given string and assigns it to the LastError field.
func (o *CheckGroup) SetLastError(v string) {
	o.LastError = &v
}

// GetOutageTime returns the OutageTime field value if set, zero value otherwise.
func (o *CheckGroup) GetOutageTime() int64 {
	if o == nil || IsNil(o.OutageTime) {
		var ret int64
		return ret
	}
	return *o.OutageTime
}

// GetOutageTimeOk returns a tuple with the OutageTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetOutageTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.OutageTime) {
		return nil, false
	}
	return o.OutageTime, true
}

// IsSetOutageTime returns a boolean if a field has been set.
func (o *CheckGroup) IsSetOutageTime() bool {
	if o != nil && !IsNil(o.OutageTime) {
		return true
	}

	return false
}

// SetOutageTime gets a reference to the given int64 and assigns it to the OutageTime field.
func (o *CheckGroup) SetOutageTime(v int64) {
	o.OutageTime = &v
}

// GetLastTimer returns the LastTimer field value if set, zero value otherwise.
func (o *CheckGroup) GetLastTimer() int64 {
	if o == nil || IsNil(o.LastTimer) {
		var ret int64
		return ret
	}
	return *o.LastTimer
}

// GetLastTimerOk returns a tuple with the LastTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetLastTimerOk() (*int64, bool) {
	if o == nil || IsNil(o.LastTimer) {
		return nil, false
	}
	return o.LastTimer, true
}

// IsSetLastTimer returns a boolean if a field has been set.
func (o *CheckGroup) IsSetLastTimer() bool {
	if o != nil && !IsNil(o.LastTimer) {
		return true
	}

	return false
}

// SetLastTimer gets a reference to the given int64 and assigns it to the LastTimer field.
func (o *CheckGroup) SetLastTimer(v int64) {
	o.LastTimer = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *CheckGroup) GetHealth() int64 {
	if o == nil || IsNil(o.Health) {
		var ret int64
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetHealthOk() (*int64, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// IsSetHealth returns a boolean if a field has been set.
func (o *CheckGroup) IsSetHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given int64 and assigns it to the Health field.
func (o *CheckGroup) SetHealth(v int64) {
	o.Health = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *CheckGroup) GetHistory() string {
	if o == nil || IsNil(o.History) {
		var ret string
		return ret
	}
	return *o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetHistoryOk() (*string, bool) {
	if o == nil || IsNil(o.History) {
		return nil, false
	}
	return o.History, true
}

// IsSetHistory returns a boolean if a field has been set.
func (o *CheckGroup) IsSetHistory() bool {
	if o != nil && !IsNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given string and assigns it to the History field.
func (o *CheckGroup) SetHistory(v string) {
	o.History = &v
}

// GetMinHappy returns the MinHappy field value if set, zero value otherwise.
func (o *CheckGroup) GetMinHappy() int64 {
	if o == nil || IsNil(o.MinHappy) {
		var ret int64
		return ret
	}
	return *o.MinHappy
}

// GetMinHappyOk returns a tuple with the MinHappy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetMinHappyOk() (*int64, bool) {
	if o == nil || IsNil(o.MinHappy) {
		return nil, false
	}
	return o.MinHappy, true
}

// IsSetMinHappy returns a boolean if a field has been set.
func (o *CheckGroup) IsSetMinHappy() bool {
	if o != nil && !IsNil(o.MinHappy) {
		return true
	}

	return false
}

// SetMinHappy gets a reference to the given int64 and assigns it to the MinHappy field.
func (o *CheckGroup) SetMinHappy(v int64) {
	o.MinHappy = &v
}

// GetLastMetric returns the LastMetric field value if set, zero value otherwise.
func (o *CheckGroup) GetLastMetric() string {
	if o == nil || IsNil(o.LastMetric) {
		var ret string
		return ret
	}
	return *o.LastMetric
}

// GetLastMetricOk returns a tuple with the LastMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetLastMetricOk() (*string, bool) {
	if o == nil || IsNil(o.LastMetric) {
		return nil, false
	}
	return o.LastMetric, true
}

// IsSetLastMetric returns a boolean if a field has been set.
func (o *CheckGroup) IsSetLastMetric() bool {
	if o != nil && !IsNil(o.LastMetric) {
		return true
	}

	return false
}

// SetLastMetric gets a reference to the given string and assigns it to the LastMetric field.
func (o *CheckGroup) SetLastMetric(v string) {
	o.LastMetric = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *CheckGroup) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// IsSetSeverity returns a boolean if a field has been set.
func (o *CheckGroup) IsSetSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *CheckGroup) SetSeverity(v string) {
	o.Severity = &v
}

// GetCreateIncident returns the CreateIncident field value if set, zero value otherwise.
func (o *CheckGroup) GetCreateIncident() bool {
	if o == nil || IsNil(o.CreateIncident) {
		var ret bool
		return ret
	}
	return *o.CreateIncident
}

// GetCreateIncidentOk returns a tuple with the CreateIncident field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetCreateIncidentOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateIncident) {
		return nil, false
	}
	return o.CreateIncident, true
}

// IsSetCreateIncident returns a boolean if a field has been set.
func (o *CheckGroup) IsSetCreateIncident() bool {
	if o != nil && !IsNil(o.CreateIncident) {
		return true
	}

	return false
}

// SetCreateIncident gets a reference to the given bool and assigns it to the CreateIncident field.
func (o *CheckGroup) SetCreateIncident(v bool) {
	o.CreateIncident = &v
}

// GetMuted returns the Muted field value if set, zero value otherwise.
func (o *CheckGroup) GetMuted() bool {
	if o == nil || IsNil(o.Muted) {
		var ret bool
		return ret
	}
	return *o.Muted
}

// GetMutedOk returns a tuple with the Muted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetMutedOk() (*bool, bool) {
	if o == nil || IsNil(o.Muted) {
		return nil, false
	}
	return o.Muted, true
}

// IsSetMuted returns a boolean if a field has been set.
func (o *CheckGroup) IsSetMuted() bool {
	if o != nil && !IsNil(o.Muted) {
		return true
	}

	return false
}

// SetMuted gets a reference to the given bool and assigns it to the Muted field.
func (o *CheckGroup) SetMuted(v bool) {
	o.Muted = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *CheckGroup) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser {
	if o == nil || IsNil(o.CreatedBy) {
		var ret ListActivity200ResponseAllOfActivityInnerActivityInnerUser
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// IsSetCreatedBy returns a boolean if a field has been set.
func (o *CheckGroup) IsSetCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given ListActivity200ResponseAllOfActivityInnerActivityInnerUser and assigns it to the CreatedBy field.
func (o *CheckGroup) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser) {
	o.CreatedBy = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *CheckGroup) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *CheckGroup) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *CheckGroup) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *CheckGroup) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *CheckGroup) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *CheckGroup) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *CheckGroup) GetAvailability() float32 {
	if o == nil || IsNil(o.Availability) {
		var ret float32
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetAvailabilityOk() (*float32, bool) {
	if o == nil || IsNil(o.Availability) {
		return nil, false
	}
	return o.Availability, true
}

// IsSetAvailability returns a boolean if a field has been set.
func (o *CheckGroup) IsSetAvailability() bool {
	if o != nil && !IsNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given float32 and assigns it to the Availability field.
func (o *CheckGroup) SetAvailability(v float32) {
	o.Availability = &v
}

// GetCheckType returns the CheckType field value if set, zero value otherwise.
func (o *CheckGroup) GetCheckType() GetAlerts200ResponseAllOfChecksInnerCheckType {
	if o == nil || IsNil(o.CheckType) {
		var ret GetAlerts200ResponseAllOfChecksInnerCheckType
		return ret
	}
	return *o.CheckType
}

// GetCheckTypeOk returns a tuple with the CheckType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetCheckTypeOk() (*GetAlerts200ResponseAllOfChecksInnerCheckType, bool) {
	if o == nil || IsNil(o.CheckType) {
		return nil, false
	}
	return o.CheckType, true
}

// IsSetCheckType returns a boolean if a field has been set.
func (o *CheckGroup) IsSetCheckType() bool {
	if o != nil && !IsNil(o.CheckType) {
		return true
	}

	return false
}

// SetCheckType gets a reference to the given GetAlerts200ResponseAllOfChecksInnerCheckType and assigns it to the CheckType field.
func (o *CheckGroup) SetCheckType(v GetAlerts200ResponseAllOfChecksInnerCheckType) {
	o.CheckType = &v
}

// GetChecks returns the Checks field value if set, zero value otherwise.
func (o *CheckGroup) GetChecks() []int64 {
	if o == nil || IsNil(o.Checks) {
		var ret []int64
		return ret
	}
	return o.Checks
}

// GetChecksOk returns a tuple with the Checks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGroup) GetChecksOk() ([]int64, bool) {
	if o == nil || IsNil(o.Checks) {
		return nil, false
	}
	return o.Checks, true
}

// IsSetChecks returns a boolean if a field has been set.
func (o *CheckGroup) IsSetChecks() bool {
	if o != nil && !IsNil(o.Checks) {
		return true
	}

	return false
}

// SetChecks gets a reference to the given []int64 and assigns it to the Checks field.
func (o *CheckGroup) SetChecks(v []int64) {
	o.Checks = v
}

func (o CheckGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InUptime) {
		toSerialize["inUptime"] = o.InUptime
	}
	if !IsNil(o.LastCheckStatus) {
		toSerialize["lastCheckStatus"] = o.LastCheckStatus
	}
	if !IsNil(o.LastWarningDate) {
		toSerialize["lastWarningDate"] = o.LastWarningDate
	}
	if !IsNil(o.LastErrorDate) {
		toSerialize["lastErrorDate"] = o.LastErrorDate
	}
	if !IsNil(o.LastSuccessDate) {
		toSerialize["lastSuccessDate"] = o.LastSuccessDate
	}
	if !IsNil(o.LastRunDate) {
		toSerialize["lastRunDate"] = o.LastRunDate
	}
	if !IsNil(o.LastError) {
		toSerialize["lastError"] = o.LastError
	}
	if !IsNil(o.OutageTime) {
		toSerialize["outageTime"] = o.OutageTime
	}
	if !IsNil(o.LastTimer) {
		toSerialize["lastTimer"] = o.LastTimer
	}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !IsNil(o.History) {
		toSerialize["history"] = o.History
	}
	if !IsNil(o.MinHappy) {
		toSerialize["minHappy"] = o.MinHappy
	}
	if !IsNil(o.LastMetric) {
		toSerialize["lastMetric"] = o.LastMetric
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.CreateIncident) {
		toSerialize["createIncident"] = o.CreateIncident
	}
	if !IsNil(o.Muted) {
		toSerialize["muted"] = o.Muted
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !IsNil(o.CheckType) {
		toSerialize["checkType"] = o.CheckType
	}
	if !IsNil(o.Checks) {
		toSerialize["checks"] = o.Checks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CheckGroup) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
