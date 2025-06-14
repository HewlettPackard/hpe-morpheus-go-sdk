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

// checks if the GuidanceAzureReservations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuidanceAzureReservations{}

// GuidanceAzureReservations struct for GuidanceAzureReservations
type GuidanceAzureReservations struct {
	Id                   *int64                                                     `json:"id,omitempty"`
	DateCreated          *time.Time                                                 `json:"dateCreated,omitempty"`
	LastUpdated          *time.Time                                                 `json:"lastUpdated,omitempty"`
	ActionCategory       *string                                                    `json:"actionCategory,omitempty"`
	ActionMessage        *string                                                    `json:"actionMessage,omitempty"`
	ActionTitle          *string                                                    `json:"actionTitle,omitempty"`
	ActionType           *string                                                    `json:"actionType,omitempty"`
	ActionValue          *string                                                    `json:"actionValue,omitempty"`
	ActionValueType      *string                                                    `json:"actionValueType,omitempty"`
	ActionPlanId         *string                                                    `json:"actionPlanId,omitempty"`
	StatusMessage        *string                                                    `json:"statusMessage,omitempty"`
	AccountId            *int64                                                     `json:"accountId,omitempty"`
	UserId               *string                                                    `json:"userId,omitempty"`
	SiteId               *int64                                                     `json:"siteId,omitempty"`
	Zone                 *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone    `json:"zone,omitempty"`
	State                *string                                                    `json:"state,omitempty"`
	StateMessage         *string                                                    `json:"stateMessage,omitempty"`
	Severity             *string                                                    `json:"severity,omitempty"`
	Resolved             *bool                                                      `json:"resolved,omitempty"`
	ResolvedMessage      *string                                                    `json:"resolvedMessage,omitempty"`
	RefType              *string                                                    `json:"refType,omitempty"`
	RefId                *int64                                                     `json:"refId,omitempty"`
	RefName              *string                                                    `json:"refName,omitempty"`
	Type                 *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType    `json:"type,omitempty"`
	Savings              *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings `json:"savings,omitempty"`
	Config               *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}                                     `json:",remain"`
}

type _GuidanceAzureReservations GuidanceAzureReservations

// NewGuidanceAzureReservations instantiates a new GuidanceAzureReservations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuidanceAzureReservations() *GuidanceAzureReservations {
	this := GuidanceAzureReservations{}
	return &this
}

// NewGuidanceAzureReservationsWithDefaults instantiates a new GuidanceAzureReservations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuidanceAzureReservationsWithDefaults() *GuidanceAzureReservations {
	this := GuidanceAzureReservations{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GuidanceAzureReservations) SetId(v int64) {
	o.Id = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *GuidanceAzureReservations) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *GuidanceAzureReservations) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetActionCategory returns the ActionCategory field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetActionCategory() string {
	if o == nil || IsNil(o.ActionCategory) {
		var ret string
		return ret
	}
	return *o.ActionCategory
}

// GetActionCategoryOk returns a tuple with the ActionCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetActionCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ActionCategory) {
		return nil, false
	}
	return o.ActionCategory, true
}

// IsSetActionCategory returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetActionCategory() bool {
	if o != nil && !IsNil(o.ActionCategory) {
		return true
	}

	return false
}

// SetActionCategory gets a reference to the given string and assigns it to the ActionCategory field.
func (o *GuidanceAzureReservations) SetActionCategory(v string) {
	o.ActionCategory = &v
}

// GetActionMessage returns the ActionMessage field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetActionMessage() string {
	if o == nil || IsNil(o.ActionMessage) {
		var ret string
		return ret
	}
	return *o.ActionMessage
}

// GetActionMessageOk returns a tuple with the ActionMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetActionMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ActionMessage) {
		return nil, false
	}
	return o.ActionMessage, true
}

// IsSetActionMessage returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetActionMessage() bool {
	if o != nil && !IsNil(o.ActionMessage) {
		return true
	}

	return false
}

// SetActionMessage gets a reference to the given string and assigns it to the ActionMessage field.
func (o *GuidanceAzureReservations) SetActionMessage(v string) {
	o.ActionMessage = &v
}

// GetActionTitle returns the ActionTitle field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetActionTitle() string {
	if o == nil || IsNil(o.ActionTitle) {
		var ret string
		return ret
	}
	return *o.ActionTitle
}

// GetActionTitleOk returns a tuple with the ActionTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetActionTitleOk() (*string, bool) {
	if o == nil || IsNil(o.ActionTitle) {
		return nil, false
	}
	return o.ActionTitle, true
}

// IsSetActionTitle returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetActionTitle() bool {
	if o != nil && !IsNil(o.ActionTitle) {
		return true
	}

	return false
}

// SetActionTitle gets a reference to the given string and assigns it to the ActionTitle field.
func (o *GuidanceAzureReservations) SetActionTitle(v string) {
	o.ActionTitle = &v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetActionType() string {
	if o == nil || IsNil(o.ActionType) {
		var ret string
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetActionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// IsSetActionType returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given string and assigns it to the ActionType field.
func (o *GuidanceAzureReservations) SetActionType(v string) {
	o.ActionType = &v
}

// GetActionValue returns the ActionValue field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetActionValue() string {
	if o == nil || IsNil(o.ActionValue) {
		var ret string
		return ret
	}
	return *o.ActionValue
}

// GetActionValueOk returns a tuple with the ActionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetActionValueOk() (*string, bool) {
	if o == nil || IsNil(o.ActionValue) {
		return nil, false
	}
	return o.ActionValue, true
}

// IsSetActionValue returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetActionValue() bool {
	if o != nil && !IsNil(o.ActionValue) {
		return true
	}

	return false
}

// SetActionValue gets a reference to the given string and assigns it to the ActionValue field.
func (o *GuidanceAzureReservations) SetActionValue(v string) {
	o.ActionValue = &v
}

// GetActionValueType returns the ActionValueType field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetActionValueType() string {
	if o == nil || IsNil(o.ActionValueType) {
		var ret string
		return ret
	}
	return *o.ActionValueType
}

// GetActionValueTypeOk returns a tuple with the ActionValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetActionValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActionValueType) {
		return nil, false
	}
	return o.ActionValueType, true
}

// IsSetActionValueType returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetActionValueType() bool {
	if o != nil && !IsNil(o.ActionValueType) {
		return true
	}

	return false
}

// SetActionValueType gets a reference to the given string and assigns it to the ActionValueType field.
func (o *GuidanceAzureReservations) SetActionValueType(v string) {
	o.ActionValueType = &v
}

// GetActionPlanId returns the ActionPlanId field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetActionPlanId() string {
	if o == nil || IsNil(o.ActionPlanId) {
		var ret string
		return ret
	}
	return *o.ActionPlanId
}

// GetActionPlanIdOk returns a tuple with the ActionPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetActionPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActionPlanId) {
		return nil, false
	}
	return o.ActionPlanId, true
}

// IsSetActionPlanId returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetActionPlanId() bool {
	if o != nil && !IsNil(o.ActionPlanId) {
		return true
	}

	return false
}

// SetActionPlanId gets a reference to the given string and assigns it to the ActionPlanId field.
func (o *GuidanceAzureReservations) SetActionPlanId(v string) {
	o.ActionPlanId = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// IsSetStatusMessage returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *GuidanceAzureReservations) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetAccountId() int64 {
	if o == nil || IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// IsSetAccountId returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *GuidanceAzureReservations) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// IsSetUserId returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *GuidanceAzureReservations) SetUserId(v string) {
	o.UserId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetSiteId() int64 {
	if o == nil || IsNil(o.SiteId) {
		var ret int64
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetSiteIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// IsSetSiteId returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given int64 and assigns it to the SiteId field.
func (o *GuidanceAzureReservations) SetSiteId(v int64) {
	o.SiteId = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetZone() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone {
	if o == nil || IsNil(o.Zone) {
		var ret ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetZoneOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// IsSetZone returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone and assigns it to the Zone field.
func (o *GuidanceAzureReservations) SetZone(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone) {
	o.Zone = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// IsSetState returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GuidanceAzureReservations) SetState(v string) {
	o.State = &v
}

// GetStateMessage returns the StateMessage field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetStateMessage() string {
	if o == nil || IsNil(o.StateMessage) {
		var ret string
		return ret
	}
	return *o.StateMessage
}

// GetStateMessageOk returns a tuple with the StateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetStateMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StateMessage) {
		return nil, false
	}
	return o.StateMessage, true
}

// IsSetStateMessage returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetStateMessage() bool {
	if o != nil && !IsNil(o.StateMessage) {
		return true
	}

	return false
}

// SetStateMessage gets a reference to the given string and assigns it to the StateMessage field.
func (o *GuidanceAzureReservations) SetStateMessage(v string) {
	o.StateMessage = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// IsSetSeverity returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *GuidanceAzureReservations) SetSeverity(v string) {
	o.Severity = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetResolved() bool {
	if o == nil || IsNil(o.Resolved) {
		var ret bool
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetResolvedOk() (*bool, bool) {
	if o == nil || IsNil(o.Resolved) {
		return nil, false
	}
	return o.Resolved, true
}

// IsSetResolved returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetResolved() bool {
	if o != nil && !IsNil(o.Resolved) {
		return true
	}

	return false
}

// SetResolved gets a reference to the given bool and assigns it to the Resolved field.
func (o *GuidanceAzureReservations) SetResolved(v bool) {
	o.Resolved = &v
}

// GetResolvedMessage returns the ResolvedMessage field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetResolvedMessage() string {
	if o == nil || IsNil(o.ResolvedMessage) {
		var ret string
		return ret
	}
	return *o.ResolvedMessage
}

// GetResolvedMessageOk returns a tuple with the ResolvedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetResolvedMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResolvedMessage) {
		return nil, false
	}
	return o.ResolvedMessage, true
}

// IsSetResolvedMessage returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetResolvedMessage() bool {
	if o != nil && !IsNil(o.ResolvedMessage) {
		return true
	}

	return false
}

// SetResolvedMessage gets a reference to the given string and assigns it to the ResolvedMessage field.
func (o *GuidanceAzureReservations) SetResolvedMessage(v string) {
	o.ResolvedMessage = &v
}

// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetRefType() string {
	if o == nil || IsNil(o.RefType) {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetRefTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RefType) {
		return nil, false
	}
	return o.RefType, true
}

// IsSetRefType returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetRefType() bool {
	if o != nil && !IsNil(o.RefType) {
		return true
	}

	return false
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *GuidanceAzureReservations) SetRefType(v string) {
	o.RefType = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetRefId() int64 {
	if o == nil || IsNil(o.RefId) {
		var ret int64
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetRefIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// IsSetRefId returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given int64 and assigns it to the RefId field.
func (o *GuidanceAzureReservations) SetRefId(v int64) {
	o.RefId = &v
}

// GetRefName returns the RefName field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetRefName() string {
	if o == nil || IsNil(o.RefName) {
		var ret string
		return ret
	}
	return *o.RefName
}

// GetRefNameOk returns a tuple with the RefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetRefNameOk() (*string, bool) {
	if o == nil || IsNil(o.RefName) {
		return nil, false
	}
	return o.RefName, true
}

// IsSetRefName returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetRefName() bool {
	if o != nil && !IsNil(o.RefName) {
		return true
	}

	return false
}

// SetRefName gets a reference to the given string and assigns it to the RefName field.
func (o *GuidanceAzureReservations) SetRefName(v string) {
	o.RefName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType {
	if o == nil || IsNil(o.Type) {
		var ret ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType and assigns it to the Type field.
func (o *GuidanceAzureReservations) SetType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType) {
	o.Type = &v
}

// GetSavings returns the Savings field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetSavings() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings {
	if o == nil || IsNil(o.Savings) {
		var ret ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings
		return ret
	}
	return *o.Savings
}

// GetSavingsOk returns a tuple with the Savings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetSavingsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings, bool) {
	if o == nil || IsNil(o.Savings) {
		return nil, false
	}
	return o.Savings, true
}

// IsSetSavings returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetSavings() bool {
	if o != nil && !IsNil(o.Savings) {
		return true
	}

	return false
}

// SetSavings gets a reference to the given ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings and assigns it to the Savings field.
func (o *GuidanceAzureReservations) SetSavings(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings) {
	o.Savings = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *GuidanceAzureReservations) GetConfig() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config {
	if o == nil || IsNil(o.Config) {
		var ret ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservations) GetConfigOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *GuidanceAzureReservations) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config and assigns it to the Config field.
func (o *GuidanceAzureReservations) SetConfig(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config) {
	o.Config = &v
}

func (o GuidanceAzureReservations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuidanceAzureReservations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.ActionCategory) {
		toSerialize["actionCategory"] = o.ActionCategory
	}
	if !IsNil(o.ActionMessage) {
		toSerialize["actionMessage"] = o.ActionMessage
	}
	if !IsNil(o.ActionTitle) {
		toSerialize["actionTitle"] = o.ActionTitle
	}
	if !IsNil(o.ActionType) {
		toSerialize["actionType"] = o.ActionType
	}
	if !IsNil(o.ActionValue) {
		toSerialize["actionValue"] = o.ActionValue
	}
	if !IsNil(o.ActionValueType) {
		toSerialize["actionValueType"] = o.ActionValueType
	}
	if !IsNil(o.ActionPlanId) {
		toSerialize["actionPlanId"] = o.ActionPlanId
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StateMessage) {
		toSerialize["stateMessage"] = o.StateMessage
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Resolved) {
		toSerialize["resolved"] = o.Resolved
	}
	if !IsNil(o.ResolvedMessage) {
		toSerialize["resolvedMessage"] = o.ResolvedMessage
	}
	if !IsNil(o.RefType) {
		toSerialize["refType"] = o.RefType
	}
	if !IsNil(o.RefId) {
		toSerialize["refId"] = o.RefId
	}
	if !IsNil(o.RefName) {
		toSerialize["refName"] = o.RefName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Savings) {
		toSerialize["savings"] = o.Savings
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GuidanceAzureReservations) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
