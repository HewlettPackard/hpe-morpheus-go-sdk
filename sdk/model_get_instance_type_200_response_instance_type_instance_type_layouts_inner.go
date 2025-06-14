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

// checks if the GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner{}

// GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner struct for GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner
type GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner struct {
	Id           *int64                                                      `json:"id,omitempty"`
	InstanceType *ListBackupSettings200ResponseBackupSettingsDefaultSchedule `json:"instanceType,omitempty"`
	Account      *GetAlerts200ResponseAllOfCheckGroupsInnerInstance          `json:"account,omitempty"`
	Code         *string                                                     `json:"code,omitempty"`
	Name         *string                                                     `json:"name,omitempty"`
	// Array of label strings, can be used for filtering.
	Labels                   []string                                                                                             `json:"labels,omitempty"`
	InstanceVersion          *string                                                                                              `json:"instanceVersion,omitempty"`
	Description              *string                                                                                              `json:"description,omitempty"`
	Creatable                *bool                                                                                                `json:"creatable,omitempty"`
	MemoryRequirement        *int64                                                                                               `json:"memoryRequirement,omitempty"`
	SortOrder                *int64                                                                                               `json:"sortOrder,omitempty"`
	SupportsConvertToManaged *bool                                                                                                `json:"supportsConvertToManaged,omitempty"`
	ProvisionType            *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType        `json:"provisionType,omitempty"`
	TaskSets                 []map[string]interface{}                                                                             `json:"taskSets,omitempty"`
	ContainerTypes           []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner `json:"containerTypes,omitempty"`
	Mounts                   []map[string]interface{}                                                                             `json:"mounts,omitempty"`
	Ports                    []map[string]interface{}                                                                             `json:"ports,omitempty"`
	OptionTypes              []map[string]interface{}                                                                             `json:"optionTypes,omitempty"`
	EnvironmentVariables     []map[string]interface{}                                                                             `json:"environmentVariables,omitempty"`
	PriceSets                []map[string]interface{}                                                                             `json:"priceSets,omitempty"`
	SpecTemplates            []map[string]interface{}                                                                             `json:"specTemplates,omitempty"`
	TfvarSecret              *string                                                                                              `json:"tfvarSecret,omitempty"`
	Permissions              *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerPermissions          `json:"permissions,omitempty"`
	AdditionalProperties     map[string]interface{}                                                                               `json:",remain"`
}

type _GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner

// NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner instantiates a new GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner() *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner {
	this := GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner{}
	return &this
}

// NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerWithDefaults instantiates a new GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerWithDefaults() *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner {
	this := GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetId(v int64) {
	o.Id = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetInstanceType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule {
	if o == nil || IsNil(o.InstanceType) {
		var ret ListBackupSettings200ResponseBackupSettingsDefaultSchedule
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetInstanceTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool) {
	if o == nil || IsNil(o.InstanceType) {
		return nil, false
	}
	return o.InstanceType, true
}

// IsSetInstanceType returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetInstanceType() bool {
	if o != nil && !IsNil(o.InstanceType) {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given ListBackupSettings200ResponseBackupSettingsDefaultSchedule and assigns it to the InstanceType field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetInstanceType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule) {
	o.InstanceType = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance {
	if o == nil || IsNil(o.Account) {
		var ret GetAlerts200ResponseAllOfCheckGroupsInnerInstance
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// IsSetAccount returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given GetAlerts200ResponseAllOfCheckGroupsInnerInstance and assigns it to the Account field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance) {
	o.Account = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// IsSetLabels returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetLabels(v []string) {
	o.Labels = v
}

// GetInstanceVersion returns the InstanceVersion field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetInstanceVersion() string {
	if o == nil || IsNil(o.InstanceVersion) {
		var ret string
		return ret
	}
	return *o.InstanceVersion
}

// GetInstanceVersionOk returns a tuple with the InstanceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetInstanceVersionOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceVersion) {
		return nil, false
	}
	return o.InstanceVersion, true
}

// IsSetInstanceVersion returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetInstanceVersion() bool {
	if o != nil && !IsNil(o.InstanceVersion) {
		return true
	}

	return false
}

// SetInstanceVersion gets a reference to the given string and assigns it to the InstanceVersion field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetInstanceVersion(v string) {
	o.InstanceVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetDescription(v string) {
	o.Description = &v
}

// GetCreatable returns the Creatable field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetCreatable() bool {
	if o == nil || IsNil(o.Creatable) {
		var ret bool
		return ret
	}
	return *o.Creatable
}

// GetCreatableOk returns a tuple with the Creatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetCreatableOk() (*bool, bool) {
	if o == nil || IsNil(o.Creatable) {
		return nil, false
	}
	return o.Creatable, true
}

// IsSetCreatable returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetCreatable() bool {
	if o != nil && !IsNil(o.Creatable) {
		return true
	}

	return false
}

// SetCreatable gets a reference to the given bool and assigns it to the Creatable field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetCreatable(v bool) {
	o.Creatable = &v
}

// GetMemoryRequirement returns the MemoryRequirement field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetMemoryRequirement() int64 {
	if o == nil || IsNil(o.MemoryRequirement) {
		var ret int64
		return ret
	}
	return *o.MemoryRequirement
}

// GetMemoryRequirementOk returns a tuple with the MemoryRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetMemoryRequirementOk() (*int64, bool) {
	if o == nil || IsNil(o.MemoryRequirement) {
		return nil, false
	}
	return o.MemoryRequirement, true
}

// IsSetMemoryRequirement returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetMemoryRequirement() bool {
	if o != nil && !IsNil(o.MemoryRequirement) {
		return true
	}

	return false
}

// SetMemoryRequirement gets a reference to the given int64 and assigns it to the MemoryRequirement field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetMemoryRequirement(v int64) {
	o.MemoryRequirement = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSortOrder() int64 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int64
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSortOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// IsSetSortOrder returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int64 and assigns it to the SortOrder field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetSortOrder(v int64) {
	o.SortOrder = &v
}

// GetSupportsConvertToManaged returns the SupportsConvertToManaged field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSupportsConvertToManaged() bool {
	if o == nil || IsNil(o.SupportsConvertToManaged) {
		var ret bool
		return ret
	}
	return *o.SupportsConvertToManaged
}

// GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSupportsConvertToManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsConvertToManaged) {
		return nil, false
	}
	return o.SupportsConvertToManaged, true
}

// IsSetSupportsConvertToManaged returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetSupportsConvertToManaged() bool {
	if o != nil && !IsNil(o.SupportsConvertToManaged) {
		return true
	}

	return false
}

// SetSupportsConvertToManaged gets a reference to the given bool and assigns it to the SupportsConvertToManaged field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetSupportsConvertToManaged(v bool) {
	o.SupportsConvertToManaged = &v
}

// GetProvisionType returns the ProvisionType field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetProvisionType() GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType {
	if o == nil || IsNil(o.ProvisionType) {
		var ret GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType
		return ret
	}
	return *o.ProvisionType
}

// GetProvisionTypeOk returns a tuple with the ProvisionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetProvisionTypeOk() (*GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType, bool) {
	if o == nil || IsNil(o.ProvisionType) {
		return nil, false
	}
	return o.ProvisionType, true
}

// IsSetProvisionType returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetProvisionType() bool {
	if o != nil && !IsNil(o.ProvisionType) {
		return true
	}

	return false
}

// SetProvisionType gets a reference to the given GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType and assigns it to the ProvisionType field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetProvisionType(v GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType) {
	o.ProvisionType = &v
}

// GetTaskSets returns the TaskSets field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetTaskSets() []map[string]interface{} {
	if o == nil || IsNil(o.TaskSets) {
		var ret []map[string]interface{}
		return ret
	}
	return o.TaskSets
}

// GetTaskSetsOk returns a tuple with the TaskSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetTaskSetsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.TaskSets) {
		return nil, false
	}
	return o.TaskSets, true
}

// IsSetTaskSets returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetTaskSets() bool {
	if o != nil && !IsNil(o.TaskSets) {
		return true
	}

	return false
}

// SetTaskSets gets a reference to the given []map[string]interface{} and assigns it to the TaskSets field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetTaskSets(v []map[string]interface{}) {
	o.TaskSets = v
}

// GetContainerTypes returns the ContainerTypes field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetContainerTypes() []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner {
	if o == nil || IsNil(o.ContainerTypes) {
		var ret []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner
		return ret
	}
	return o.ContainerTypes
}

// GetContainerTypesOk returns a tuple with the ContainerTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetContainerTypesOk() ([]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner, bool) {
	if o == nil || IsNil(o.ContainerTypes) {
		return nil, false
	}
	return o.ContainerTypes, true
}

// IsSetContainerTypes returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetContainerTypes() bool {
	if o != nil && !IsNil(o.ContainerTypes) {
		return true
	}

	return false
}

// SetContainerTypes gets a reference to the given []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner and assigns it to the ContainerTypes field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetContainerTypes(v []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) {
	o.ContainerTypes = v
}

// GetMounts returns the Mounts field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetMounts() []map[string]interface{} {
	if o == nil || IsNil(o.Mounts) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Mounts
}

// GetMountsOk returns a tuple with the Mounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetMountsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Mounts) {
		return nil, false
	}
	return o.Mounts, true
}

// IsSetMounts returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetMounts() bool {
	if o != nil && !IsNil(o.Mounts) {
		return true
	}

	return false
}

// SetMounts gets a reference to the given []map[string]interface{} and assigns it to the Mounts field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetMounts(v []map[string]interface{}) {
	o.Mounts = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPorts() []map[string]interface{} {
	if o == nil || IsNil(o.Ports) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPortsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// IsSetPorts returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []map[string]interface{} and assigns it to the Ports field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetPorts(v []map[string]interface{}) {
	o.Ports = v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetOptionTypes() []map[string]interface{} {
	if o == nil || IsNil(o.OptionTypes) {
		var ret []map[string]interface{}
		return ret
	}
	return o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetOptionTypesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.OptionTypes) {
		return nil, false
	}
	return o.OptionTypes, true
}

// IsSetOptionTypes returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetOptionTypes() bool {
	if o != nil && !IsNil(o.OptionTypes) {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []map[string]interface{} and assigns it to the OptionTypes field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetOptionTypes(v []map[string]interface{}) {
	o.OptionTypes = v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetEnvironmentVariables() []map[string]interface{} {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret []map[string]interface{}
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetEnvironmentVariablesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// IsSetEnvironmentVariables returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetEnvironmentVariables() bool {
	if o != nil && !IsNil(o.EnvironmentVariables) {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given []map[string]interface{} and assigns it to the EnvironmentVariables field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetEnvironmentVariables(v []map[string]interface{}) {
	o.EnvironmentVariables = v
}

// GetPriceSets returns the PriceSets field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPriceSets() []map[string]interface{} {
	if o == nil || IsNil(o.PriceSets) {
		var ret []map[string]interface{}
		return ret
	}
	return o.PriceSets
}

// GetPriceSetsOk returns a tuple with the PriceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPriceSetsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.PriceSets) {
		return nil, false
	}
	return o.PriceSets, true
}

// IsSetPriceSets returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetPriceSets() bool {
	if o != nil && !IsNil(o.PriceSets) {
		return true
	}

	return false
}

// SetPriceSets gets a reference to the given []map[string]interface{} and assigns it to the PriceSets field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetPriceSets(v []map[string]interface{}) {
	o.PriceSets = v
}

// GetSpecTemplates returns the SpecTemplates field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSpecTemplates() []map[string]interface{} {
	if o == nil || IsNil(o.SpecTemplates) {
		var ret []map[string]interface{}
		return ret
	}
	return o.SpecTemplates
}

// GetSpecTemplatesOk returns a tuple with the SpecTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSpecTemplatesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.SpecTemplates) {
		return nil, false
	}
	return o.SpecTemplates, true
}

// IsSetSpecTemplates returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetSpecTemplates() bool {
	if o != nil && !IsNil(o.SpecTemplates) {
		return true
	}

	return false
}

// SetSpecTemplates gets a reference to the given []map[string]interface{} and assigns it to the SpecTemplates field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetSpecTemplates(v []map[string]interface{}) {
	o.SpecTemplates = v
}

// GetTfvarSecret returns the TfvarSecret field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetTfvarSecret() string {
	if o == nil || IsNil(o.TfvarSecret) {
		var ret string
		return ret
	}
	return *o.TfvarSecret
}

// GetTfvarSecretOk returns a tuple with the TfvarSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetTfvarSecretOk() (*string, bool) {
	if o == nil || IsNil(o.TfvarSecret) {
		return nil, false
	}
	return o.TfvarSecret, true
}

// IsSetTfvarSecret returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetTfvarSecret() bool {
	if o != nil && !IsNil(o.TfvarSecret) {
		return true
	}

	return false
}

// SetTfvarSecret gets a reference to the given string and assigns it to the TfvarSecret field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetTfvarSecret(v string) {
	o.TfvarSecret = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPermissions() GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerPermissions {
	if o == nil || IsNil(o.Permissions) {
		var ret GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPermissionsOk() (*GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerPermissions, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// IsSetPermissions returns a boolean if a field has been set.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) IsSetPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerPermissions and assigns it to the Permissions field.
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetPermissions(v GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerPermissions) {
	o.Permissions = &v
}

func (o GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InstanceType) {
		toSerialize["instanceType"] = o.InstanceType
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.InstanceVersion) {
		toSerialize["instanceVersion"] = o.InstanceVersion
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Creatable) {
		toSerialize["creatable"] = o.Creatable
	}
	if !IsNil(o.MemoryRequirement) {
		toSerialize["memoryRequirement"] = o.MemoryRequirement
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if !IsNil(o.SupportsConvertToManaged) {
		toSerialize["supportsConvertToManaged"] = o.SupportsConvertToManaged
	}
	if !IsNil(o.ProvisionType) {
		toSerialize["provisionType"] = o.ProvisionType
	}
	if !IsNil(o.TaskSets) {
		toSerialize["taskSets"] = o.TaskSets
	}
	if !IsNil(o.ContainerTypes) {
		toSerialize["containerTypes"] = o.ContainerTypes
	}
	if !IsNil(o.Mounts) {
		toSerialize["mounts"] = o.Mounts
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.OptionTypes) {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	if !IsNil(o.PriceSets) {
		toSerialize["priceSets"] = o.PriceSets
	}
	if !IsNil(o.SpecTemplates) {
		toSerialize["specTemplates"] = o.SpecTemplates
	}
	if !IsNil(o.TfvarSecret) {
		toSerialize["tfvarSecret"] = o.TfvarSecret
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
