# GetCloudTypes200ResponseZoneTypeServerTypesInner

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
**ProvisionType** | Pointer to [**GetCloudTypes200ResponseZoneTypeServerTypesInnerProvisionType**](GetCloudTypes200ResponseZoneTypeServerTypesInnerProvisionType.md) |  | [optional] 
**OptionTypes** | Pointer to [**[]GetCloudTypes200ResponseZoneTypeServerTypesInnerOptionTypesInner**](GetCloudTypes200ResponseZoneTypeServerTypesInnerOptionTypesInner.md) |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetCloudTypes200ResponseZoneTypeServerTypesInner

`func NewGetCloudTypes200ResponseZoneTypeServerTypesInner() *GetCloudTypes200ResponseZoneTypeServerTypesInner`

NewGetCloudTypes200ResponseZoneTypeServerTypesInner instantiates a new GetCloudTypes200ResponseZoneTypeServerTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNodeType

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetPlatform

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetEnabled

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSelectable

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetExternalDelete

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetExternalDelete() bool`

GetExternalDelete returns the ExternalDelete field if non-nil, zero value otherwise.

### GetExternalDeleteOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetExternalDeleteOk() (*bool, bool)`

GetExternalDeleteOk returns a tuple with the ExternalDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDelete

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetExternalDelete(v bool)`

SetExternalDelete sets ExternalDelete field to given value.

### HasExternalDelete

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasExternalDelete() bool`

HasExternalDelete returns a boolean if a field has been set.

### GetManaged

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetControlPower

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetControlPower() bool`

GetControlPower returns the ControlPower field if non-nil, zero value otherwise.

### GetControlPowerOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetControlPowerOk() (*bool, bool)`

GetControlPowerOk returns a tuple with the ControlPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPower

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetControlPower(v bool)`

SetControlPower sets ControlPower field to given value.

### HasControlPower

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasControlPower() bool`

HasControlPower returns a boolean if a field has been set.

### GetControlSuspend

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetControlSuspend() bool`

GetControlSuspend returns the ControlSuspend field if non-nil, zero value otherwise.

### GetControlSuspendOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetControlSuspendOk() (*bool, bool)`

GetControlSuspendOk returns a tuple with the ControlSuspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlSuspend

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetControlSuspend(v bool)`

SetControlSuspend sets ControlSuspend field to given value.

### HasControlSuspend

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasControlSuspend() bool`

HasControlSuspend returns a boolean if a field has been set.

### GetCreatable

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasAgent

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetHasAgent() bool`

GetHasAgent returns the HasAgent field if non-nil, zero value otherwise.

### GetHasAgentOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetHasAgentOk() (*bool, bool)`

GetHasAgentOk returns a tuple with the HasAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAgent

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetHasAgent(v bool)`

SetHasAgent sets HasAgent field to given value.

### HasHasAgent

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasHasAgent() bool`

HasHasAgent returns a boolean if a field has been set.

### GetVmHypervisor

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetVmHypervisor() bool`

GetVmHypervisor returns the VmHypervisor field if non-nil, zero value otherwise.

### GetVmHypervisorOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetVmHypervisorOk() (*bool, bool)`

GetVmHypervisorOk returns a tuple with the VmHypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHypervisor

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetVmHypervisor(v bool)`

SetVmHypervisor sets VmHypervisor field to given value.

### HasVmHypervisor

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasVmHypervisor() bool`

HasVmHypervisor returns a boolean if a field has been set.

### GetContainerHypervisor

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetContainerHypervisor() bool`

GetContainerHypervisor returns the ContainerHypervisor field if non-nil, zero value otherwise.

### GetContainerHypervisorOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetContainerHypervisorOk() (*bool, bool)`

GetContainerHypervisorOk returns a tuple with the ContainerHypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerHypervisor

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetContainerHypervisor(v bool)`

SetContainerHypervisor sets ContainerHypervisor field to given value.

### HasContainerHypervisor

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasContainerHypervisor() bool`

HasContainerHypervisor returns a boolean if a field has been set.

### GetBareMetalHost

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetBareMetalHost() bool`

GetBareMetalHost returns the BareMetalHost field if non-nil, zero value otherwise.

### GetBareMetalHostOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetBareMetalHostOk() (*bool, bool)`

GetBareMetalHostOk returns a tuple with the BareMetalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBareMetalHost

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetBareMetalHost(v bool)`

SetBareMetalHost sets BareMetalHost field to given value.

### HasBareMetalHost

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasBareMetalHost() bool`

HasBareMetalHost returns a boolean if a field has been set.

### GetGuestVm

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetGuestVm() bool`

GetGuestVm returns the GuestVm field if non-nil, zero value otherwise.

### GetGuestVmOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetGuestVmOk() (*bool, bool)`

GetGuestVmOk returns a tuple with the GuestVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVm

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetGuestVm(v bool)`

SetGuestVm sets GuestVm field to given value.

### HasGuestVm

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasGuestVm() bool`

HasGuestVm returns a boolean if a field has been set.

### GetHasAutomation

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetHasAutomation() bool`

GetHasAutomation returns the HasAutomation field if non-nil, zero value otherwise.

### GetHasAutomationOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetHasAutomationOk() (*bool, bool)`

GetHasAutomationOk returns a tuple with the HasAutomation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomation

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetHasAutomation(v bool)`

SetHasAutomation sets HasAutomation field to given value.

### HasHasAutomation

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasHasAutomation() bool`

HasHasAutomation returns a boolean if a field has been set.

### GetProvisionType

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetProvisionType() GetCloudTypes200ResponseZoneTypeServerTypesInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetProvisionTypeOk() (*GetCloudTypes200ResponseZoneTypeServerTypesInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetProvisionType(v GetCloudTypes200ResponseZoneTypeServerTypesInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetOptionTypes() []GetCloudTypes200ResponseZoneTypeServerTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetOptionTypesOk() (*[]GetCloudTypes200ResponseZoneTypeServerTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetOptionTypes(v []GetCloudTypes200ResponseZoneTypeServerTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *GetCloudTypes200ResponseZoneTypeServerTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


