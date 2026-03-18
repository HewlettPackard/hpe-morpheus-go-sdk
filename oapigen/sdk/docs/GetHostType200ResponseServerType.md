# GetHostType200ResponseServerType

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
**ControlStart** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**HasAgent** | Pointer to **bool** |  | [optional] 
**VmHypervisor** | Pointer to **bool** |  | [optional] 
**ContainerHypervisor** | Pointer to **bool** |  | [optional] 
**BareMetalHost** | Pointer to **bool** |  | [optional] 
**GuestVm** | Pointer to **bool** |  | [optional] 
**HasAutomation** | Pointer to **bool** |  | [optional] 
**ProvisionType** | Pointer to [**GetHostType200ResponseServerTypeProvisionType**](GetHostType200ResponseServerTypeProvisionType.md) |  | [optional] 
**OptionTypes** | Pointer to [**[]GetHostType200ResponseServerTypeOptionTypesInner**](GetHostType200ResponseServerTypeOptionTypesInner.md) |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetHostType200ResponseServerType

`func NewGetHostType200ResponseServerType() *GetHostType200ResponseServerType`

NewGetHostType200ResponseServerType instantiates a new GetHostType200ResponseServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHostType200ResponseServerTypeWithDefaults

`func NewGetHostType200ResponseServerTypeWithDefaults() *GetHostType200ResponseServerType`

NewGetHostType200ResponseServerTypeWithDefaults instantiates a new GetHostType200ResponseServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetHostType200ResponseServerType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetHostType200ResponseServerType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetHostType200ResponseServerType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetHostType200ResponseServerType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetHostType200ResponseServerType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetHostType200ResponseServerType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetHostType200ResponseServerType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetHostType200ResponseServerType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetHostType200ResponseServerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetHostType200ResponseServerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetHostType200ResponseServerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetHostType200ResponseServerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetHostType200ResponseServerType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetHostType200ResponseServerType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetHostType200ResponseServerType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetHostType200ResponseServerType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNodeType

`func (o *GetHostType200ResponseServerType) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *GetHostType200ResponseServerType) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *GetHostType200ResponseServerType) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *GetHostType200ResponseServerType) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetPlatform

`func (o *GetHostType200ResponseServerType) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetHostType200ResponseServerType) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetHostType200ResponseServerType) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetHostType200ResponseServerType) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetEnabled

`func (o *GetHostType200ResponseServerType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetHostType200ResponseServerType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetHostType200ResponseServerType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetHostType200ResponseServerType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSelectable

`func (o *GetHostType200ResponseServerType) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *GetHostType200ResponseServerType) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *GetHostType200ResponseServerType) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *GetHostType200ResponseServerType) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetExternalDelete

`func (o *GetHostType200ResponseServerType) GetExternalDelete() bool`

GetExternalDelete returns the ExternalDelete field if non-nil, zero value otherwise.

### GetExternalDeleteOk

`func (o *GetHostType200ResponseServerType) GetExternalDeleteOk() (*bool, bool)`

GetExternalDeleteOk returns a tuple with the ExternalDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDelete

`func (o *GetHostType200ResponseServerType) SetExternalDelete(v bool)`

SetExternalDelete sets ExternalDelete field to given value.

### HasExternalDelete

`func (o *GetHostType200ResponseServerType) HasExternalDelete() bool`

HasExternalDelete returns a boolean if a field has been set.

### GetManaged

`func (o *GetHostType200ResponseServerType) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *GetHostType200ResponseServerType) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *GetHostType200ResponseServerType) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *GetHostType200ResponseServerType) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetControlPower

`func (o *GetHostType200ResponseServerType) GetControlPower() bool`

GetControlPower returns the ControlPower field if non-nil, zero value otherwise.

### GetControlPowerOk

`func (o *GetHostType200ResponseServerType) GetControlPowerOk() (*bool, bool)`

GetControlPowerOk returns a tuple with the ControlPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPower

`func (o *GetHostType200ResponseServerType) SetControlPower(v bool)`

SetControlPower sets ControlPower field to given value.

### HasControlPower

`func (o *GetHostType200ResponseServerType) HasControlPower() bool`

HasControlPower returns a boolean if a field has been set.

### GetControlSuspend

`func (o *GetHostType200ResponseServerType) GetControlSuspend() bool`

GetControlSuspend returns the ControlSuspend field if non-nil, zero value otherwise.

### GetControlSuspendOk

`func (o *GetHostType200ResponseServerType) GetControlSuspendOk() (*bool, bool)`

GetControlSuspendOk returns a tuple with the ControlSuspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlSuspend

`func (o *GetHostType200ResponseServerType) SetControlSuspend(v bool)`

SetControlSuspend sets ControlSuspend field to given value.

### HasControlSuspend

`func (o *GetHostType200ResponseServerType) HasControlSuspend() bool`

HasControlSuspend returns a boolean if a field has been set.

### GetControlStart

`func (o *GetHostType200ResponseServerType) GetControlStart() bool`

GetControlStart returns the ControlStart field if non-nil, zero value otherwise.

### GetControlStartOk

`func (o *GetHostType200ResponseServerType) GetControlStartOk() (*bool, bool)`

GetControlStartOk returns a tuple with the ControlStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlStart

`func (o *GetHostType200ResponseServerType) SetControlStart(v bool)`

SetControlStart sets ControlStart field to given value.

### HasControlStart

`func (o *GetHostType200ResponseServerType) HasControlStart() bool`

HasControlStart returns a boolean if a field has been set.

### GetCreatable

`func (o *GetHostType200ResponseServerType) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetHostType200ResponseServerType) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetHostType200ResponseServerType) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetHostType200ResponseServerType) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasAgent

`func (o *GetHostType200ResponseServerType) GetHasAgent() bool`

GetHasAgent returns the HasAgent field if non-nil, zero value otherwise.

### GetHasAgentOk

`func (o *GetHostType200ResponseServerType) GetHasAgentOk() (*bool, bool)`

GetHasAgentOk returns a tuple with the HasAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAgent

`func (o *GetHostType200ResponseServerType) SetHasAgent(v bool)`

SetHasAgent sets HasAgent field to given value.

### HasHasAgent

`func (o *GetHostType200ResponseServerType) HasHasAgent() bool`

HasHasAgent returns a boolean if a field has been set.

### GetVmHypervisor

`func (o *GetHostType200ResponseServerType) GetVmHypervisor() bool`

GetVmHypervisor returns the VmHypervisor field if non-nil, zero value otherwise.

### GetVmHypervisorOk

`func (o *GetHostType200ResponseServerType) GetVmHypervisorOk() (*bool, bool)`

GetVmHypervisorOk returns a tuple with the VmHypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHypervisor

`func (o *GetHostType200ResponseServerType) SetVmHypervisor(v bool)`

SetVmHypervisor sets VmHypervisor field to given value.

### HasVmHypervisor

`func (o *GetHostType200ResponseServerType) HasVmHypervisor() bool`

HasVmHypervisor returns a boolean if a field has been set.

### GetContainerHypervisor

`func (o *GetHostType200ResponseServerType) GetContainerHypervisor() bool`

GetContainerHypervisor returns the ContainerHypervisor field if non-nil, zero value otherwise.

### GetContainerHypervisorOk

`func (o *GetHostType200ResponseServerType) GetContainerHypervisorOk() (*bool, bool)`

GetContainerHypervisorOk returns a tuple with the ContainerHypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerHypervisor

`func (o *GetHostType200ResponseServerType) SetContainerHypervisor(v bool)`

SetContainerHypervisor sets ContainerHypervisor field to given value.

### HasContainerHypervisor

`func (o *GetHostType200ResponseServerType) HasContainerHypervisor() bool`

HasContainerHypervisor returns a boolean if a field has been set.

### GetBareMetalHost

`func (o *GetHostType200ResponseServerType) GetBareMetalHost() bool`

GetBareMetalHost returns the BareMetalHost field if non-nil, zero value otherwise.

### GetBareMetalHostOk

`func (o *GetHostType200ResponseServerType) GetBareMetalHostOk() (*bool, bool)`

GetBareMetalHostOk returns a tuple with the BareMetalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBareMetalHost

`func (o *GetHostType200ResponseServerType) SetBareMetalHost(v bool)`

SetBareMetalHost sets BareMetalHost field to given value.

### HasBareMetalHost

`func (o *GetHostType200ResponseServerType) HasBareMetalHost() bool`

HasBareMetalHost returns a boolean if a field has been set.

### GetGuestVm

`func (o *GetHostType200ResponseServerType) GetGuestVm() bool`

GetGuestVm returns the GuestVm field if non-nil, zero value otherwise.

### GetGuestVmOk

`func (o *GetHostType200ResponseServerType) GetGuestVmOk() (*bool, bool)`

GetGuestVmOk returns a tuple with the GuestVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVm

`func (o *GetHostType200ResponseServerType) SetGuestVm(v bool)`

SetGuestVm sets GuestVm field to given value.

### HasGuestVm

`func (o *GetHostType200ResponseServerType) HasGuestVm() bool`

HasGuestVm returns a boolean if a field has been set.

### GetHasAutomation

`func (o *GetHostType200ResponseServerType) GetHasAutomation() bool`

GetHasAutomation returns the HasAutomation field if non-nil, zero value otherwise.

### GetHasAutomationOk

`func (o *GetHostType200ResponseServerType) GetHasAutomationOk() (*bool, bool)`

GetHasAutomationOk returns a tuple with the HasAutomation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomation

`func (o *GetHostType200ResponseServerType) SetHasAutomation(v bool)`

SetHasAutomation sets HasAutomation field to given value.

### HasHasAutomation

`func (o *GetHostType200ResponseServerType) HasHasAutomation() bool`

HasHasAutomation returns a boolean if a field has been set.

### GetProvisionType

`func (o *GetHostType200ResponseServerType) GetProvisionType() GetHostType200ResponseServerTypeProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GetHostType200ResponseServerType) GetProvisionTypeOk() (*GetHostType200ResponseServerTypeProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GetHostType200ResponseServerType) SetProvisionType(v GetHostType200ResponseServerTypeProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GetHostType200ResponseServerType) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetHostType200ResponseServerType) GetOptionTypes() []GetHostType200ResponseServerTypeOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetHostType200ResponseServerType) GetOptionTypesOk() (*[]GetHostType200ResponseServerTypeOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetHostType200ResponseServerType) SetOptionTypes(v []GetHostType200ResponseServerTypeOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetHostType200ResponseServerType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *GetHostType200ResponseServerType) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *GetHostType200ResponseServerType) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *GetHostType200ResponseServerType) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *GetHostType200ResponseServerType) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


