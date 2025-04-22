# ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**NodeType** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Selectable** | Pointer to **bool** |  | [optional] 
**ExternalDelete** | Pointer to **bool** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**ControlPower** | Pointer to **bool** |  | [optional] 
**ControlSuspend** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**HasAgent** | Pointer to **bool** |  | [optional] 
**VmHypervisor** | Pointer to **bool** |  | [optional] 
**ContainerHypervisor** | Pointer to **bool** |  | [optional] 
**BareMetalHost** | Pointer to **bool** |  | [optional] 
**GuestVm** | Pointer to **bool** |  | [optional] 
**HasAutomation** | Pointer to **bool** |  | [optional] 
**ProvisionType** | Pointer to [**ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerProvisionType**](ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerProvisionType.md) |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerOptionTypesInner**](ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerOptionTypesInner.md) |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 

## Methods

### NewListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner

`func NewListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner() *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner`

NewListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner instantiates a new ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerWithDefaults

`func NewListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerWithDefaults() *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner`

NewListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerWithDefaults instantiates a new ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNodeType

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetPlatform

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetEnabled

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSelectable

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetExternalDelete

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetExternalDelete() bool`

GetExternalDelete returns the ExternalDelete field if non-nil, zero value otherwise.

### GetExternalDeleteOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetExternalDeleteOk() (*bool, bool)`

GetExternalDeleteOk returns a tuple with the ExternalDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDelete

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetExternalDelete(v bool)`

SetExternalDelete sets ExternalDelete field to given value.

### HasExternalDelete

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasExternalDelete() bool`

HasExternalDelete returns a boolean if a field has been set.

### GetManaged

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetControlPower

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetControlPower() bool`

GetControlPower returns the ControlPower field if non-nil, zero value otherwise.

### GetControlPowerOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetControlPowerOk() (*bool, bool)`

GetControlPowerOk returns a tuple with the ControlPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPower

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetControlPower(v bool)`

SetControlPower sets ControlPower field to given value.

### HasControlPower

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasControlPower() bool`

HasControlPower returns a boolean if a field has been set.

### GetControlSuspend

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetControlSuspend() bool`

GetControlSuspend returns the ControlSuspend field if non-nil, zero value otherwise.

### GetControlSuspendOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetControlSuspendOk() (*bool, bool)`

GetControlSuspendOk returns a tuple with the ControlSuspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlSuspend

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetControlSuspend(v bool)`

SetControlSuspend sets ControlSuspend field to given value.

### HasControlSuspend

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasControlSuspend() bool`

HasControlSuspend returns a boolean if a field has been set.

### GetCreatable

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasAgent

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetHasAgent() bool`

GetHasAgent returns the HasAgent field if non-nil, zero value otherwise.

### GetHasAgentOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetHasAgentOk() (*bool, bool)`

GetHasAgentOk returns a tuple with the HasAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAgent

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetHasAgent(v bool)`

SetHasAgent sets HasAgent field to given value.

### HasHasAgent

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasHasAgent() bool`

HasHasAgent returns a boolean if a field has been set.

### GetVmHypervisor

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetVmHypervisor() bool`

GetVmHypervisor returns the VmHypervisor field if non-nil, zero value otherwise.

### GetVmHypervisorOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetVmHypervisorOk() (*bool, bool)`

GetVmHypervisorOk returns a tuple with the VmHypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHypervisor

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetVmHypervisor(v bool)`

SetVmHypervisor sets VmHypervisor field to given value.

### HasVmHypervisor

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasVmHypervisor() bool`

HasVmHypervisor returns a boolean if a field has been set.

### GetContainerHypervisor

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetContainerHypervisor() bool`

GetContainerHypervisor returns the ContainerHypervisor field if non-nil, zero value otherwise.

### GetContainerHypervisorOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetContainerHypervisorOk() (*bool, bool)`

GetContainerHypervisorOk returns a tuple with the ContainerHypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerHypervisor

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetContainerHypervisor(v bool)`

SetContainerHypervisor sets ContainerHypervisor field to given value.

### HasContainerHypervisor

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasContainerHypervisor() bool`

HasContainerHypervisor returns a boolean if a field has been set.

### GetBareMetalHost

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetBareMetalHost() bool`

GetBareMetalHost returns the BareMetalHost field if non-nil, zero value otherwise.

### GetBareMetalHostOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetBareMetalHostOk() (*bool, bool)`

GetBareMetalHostOk returns a tuple with the BareMetalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBareMetalHost

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetBareMetalHost(v bool)`

SetBareMetalHost sets BareMetalHost field to given value.

### HasBareMetalHost

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasBareMetalHost() bool`

HasBareMetalHost returns a boolean if a field has been set.

### GetGuestVm

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetGuestVm() bool`

GetGuestVm returns the GuestVm field if non-nil, zero value otherwise.

### GetGuestVmOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetGuestVmOk() (*bool, bool)`

GetGuestVmOk returns a tuple with the GuestVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVm

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetGuestVm(v bool)`

SetGuestVm sets GuestVm field to given value.

### HasGuestVm

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasGuestVm() bool`

HasGuestVm returns a boolean if a field has been set.

### GetHasAutomation

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetHasAutomation() bool`

GetHasAutomation returns the HasAutomation field if non-nil, zero value otherwise.

### GetHasAutomationOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetHasAutomationOk() (*bool, bool)`

GetHasAutomationOk returns a tuple with the HasAutomation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomation

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetHasAutomation(v bool)`

SetHasAutomation sets HasAutomation field to given value.

### HasHasAutomation

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasHasAutomation() bool`

HasHasAutomation returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetProvisionType() ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetProvisionTypeOk() (*ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetProvisionType(v ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetOptionTypes() []ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetOptionTypesOk() (*[]ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetOptionTypes(v []ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


