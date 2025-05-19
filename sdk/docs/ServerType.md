# ServerType

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
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 

## Methods

### NewServerType

`func NewServerType() *ServerType`

NewServerType instantiates a new ServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTypeWithDefaults

`func NewServerTypeWithDefaults() *ServerType`

NewServerTypeWithDefaults instantiates a new ServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServerType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ServerType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServerType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServerType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServerType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ServerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ServerType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNodeType

`func (o *ServerType) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ServerType) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ServerType) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *ServerType) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetPlatform

`func (o *ServerType) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ServerType) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ServerType) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ServerType) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetEnabled

`func (o *ServerType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServerType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServerType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServerType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSelectable

`func (o *ServerType) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *ServerType) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *ServerType) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *ServerType) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetExternalDelete

`func (o *ServerType) GetExternalDelete() bool`

GetExternalDelete returns the ExternalDelete field if non-nil, zero value otherwise.

### GetExternalDeleteOk

`func (o *ServerType) GetExternalDeleteOk() (*bool, bool)`

GetExternalDeleteOk returns a tuple with the ExternalDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDelete

`func (o *ServerType) SetExternalDelete(v bool)`

SetExternalDelete sets ExternalDelete field to given value.

### HasExternalDelete

`func (o *ServerType) HasExternalDelete() bool`

HasExternalDelete returns a boolean if a field has been set.

### GetManaged

`func (o *ServerType) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ServerType) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ServerType) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ServerType) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetControlPower

`func (o *ServerType) GetControlPower() bool`

GetControlPower returns the ControlPower field if non-nil, zero value otherwise.

### GetControlPowerOk

`func (o *ServerType) GetControlPowerOk() (*bool, bool)`

GetControlPowerOk returns a tuple with the ControlPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPower

`func (o *ServerType) SetControlPower(v bool)`

SetControlPower sets ControlPower field to given value.

### HasControlPower

`func (o *ServerType) HasControlPower() bool`

HasControlPower returns a boolean if a field has been set.

### GetControlSuspend

`func (o *ServerType) GetControlSuspend() bool`

GetControlSuspend returns the ControlSuspend field if non-nil, zero value otherwise.

### GetControlSuspendOk

`func (o *ServerType) GetControlSuspendOk() (*bool, bool)`

GetControlSuspendOk returns a tuple with the ControlSuspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlSuspend

`func (o *ServerType) SetControlSuspend(v bool)`

SetControlSuspend sets ControlSuspend field to given value.

### HasControlSuspend

`func (o *ServerType) HasControlSuspend() bool`

HasControlSuspend returns a boolean if a field has been set.

### GetCreatable

`func (o *ServerType) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ServerType) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ServerType) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ServerType) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasAgent

`func (o *ServerType) GetHasAgent() bool`

GetHasAgent returns the HasAgent field if non-nil, zero value otherwise.

### GetHasAgentOk

`func (o *ServerType) GetHasAgentOk() (*bool, bool)`

GetHasAgentOk returns a tuple with the HasAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAgent

`func (o *ServerType) SetHasAgent(v bool)`

SetHasAgent sets HasAgent field to given value.

### HasHasAgent

`func (o *ServerType) HasHasAgent() bool`

HasHasAgent returns a boolean if a field has been set.

### GetVmHypervisor

`func (o *ServerType) GetVmHypervisor() bool`

GetVmHypervisor returns the VmHypervisor field if non-nil, zero value otherwise.

### GetVmHypervisorOk

`func (o *ServerType) GetVmHypervisorOk() (*bool, bool)`

GetVmHypervisorOk returns a tuple with the VmHypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHypervisor

`func (o *ServerType) SetVmHypervisor(v bool)`

SetVmHypervisor sets VmHypervisor field to given value.

### HasVmHypervisor

`func (o *ServerType) HasVmHypervisor() bool`

HasVmHypervisor returns a boolean if a field has been set.

### GetContainerHypervisor

`func (o *ServerType) GetContainerHypervisor() bool`

GetContainerHypervisor returns the ContainerHypervisor field if non-nil, zero value otherwise.

### GetContainerHypervisorOk

`func (o *ServerType) GetContainerHypervisorOk() (*bool, bool)`

GetContainerHypervisorOk returns a tuple with the ContainerHypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerHypervisor

`func (o *ServerType) SetContainerHypervisor(v bool)`

SetContainerHypervisor sets ContainerHypervisor field to given value.

### HasContainerHypervisor

`func (o *ServerType) HasContainerHypervisor() bool`

HasContainerHypervisor returns a boolean if a field has been set.

### GetBareMetalHost

`func (o *ServerType) GetBareMetalHost() bool`

GetBareMetalHost returns the BareMetalHost field if non-nil, zero value otherwise.

### GetBareMetalHostOk

`func (o *ServerType) GetBareMetalHostOk() (*bool, bool)`

GetBareMetalHostOk returns a tuple with the BareMetalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBareMetalHost

`func (o *ServerType) SetBareMetalHost(v bool)`

SetBareMetalHost sets BareMetalHost field to given value.

### HasBareMetalHost

`func (o *ServerType) HasBareMetalHost() bool`

HasBareMetalHost returns a boolean if a field has been set.

### GetGuestVm

`func (o *ServerType) GetGuestVm() bool`

GetGuestVm returns the GuestVm field if non-nil, zero value otherwise.

### GetGuestVmOk

`func (o *ServerType) GetGuestVmOk() (*bool, bool)`

GetGuestVmOk returns a tuple with the GuestVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVm

`func (o *ServerType) SetGuestVm(v bool)`

SetGuestVm sets GuestVm field to given value.

### HasGuestVm

`func (o *ServerType) HasGuestVm() bool`

HasGuestVm returns a boolean if a field has been set.

### GetHasAutomation

`func (o *ServerType) GetHasAutomation() bool`

GetHasAutomation returns the HasAutomation field if non-nil, zero value otherwise.

### GetHasAutomationOk

`func (o *ServerType) GetHasAutomationOk() (*bool, bool)`

GetHasAutomationOk returns a tuple with the HasAutomation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutomation

`func (o *ServerType) SetHasAutomation(v bool)`

SetHasAutomation sets HasAutomation field to given value.

### HasHasAutomation

`func (o *ServerType) HasHasAutomation() bool`

HasHasAutomation returns a boolean if a field has been set.

### GetProvisionType

`func (o *ServerType) GetProvisionType() ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ServerType) GetProvisionTypeOk() (*ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ServerType) SetProvisionType(v ListCloudTypes200ResponseAllOfZoneTypesInnerServerTypesInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ServerType) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ServerType) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ServerType) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ServerType) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ServerType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ServerType) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ServerType) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ServerType) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ServerType) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


