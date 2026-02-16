# InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**ExternalType** | Pointer to **NullableString** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**Overlay** | Pointer to **bool** |  | [optional] 
**NameEditable** | Pointer to **bool** |  | [optional] 
**CidrRequired** | Pointer to **bool** |  | [optional] 
**CidrEditable** | Pointer to **bool** |  | [optional] 
**DhcpServerEditable** | Pointer to **bool** |  | [optional] 
**DnsEditable** | Pointer to **bool** |  | [optional] 
**GatewayEditable** | Pointer to **bool** |  | [optional] 
**VlanIdEditable** | Pointer to **bool** |  | [optional] 
**StaticOverrideEditable** | Pointer to **bool** |  | [optional] 
**NetworkDomainEditable** | Pointer to **bool** |  | [optional] 
**CanAssignPool** | Pointer to **bool** |  | [optional] 
**Deletable** | Pointer to **bool** |  | [optional] 
**HasNetworkServer** | Pointer to **bool** |  | [optional] 
**HasCidr** | Pointer to **bool** |  | [optional] 
**HasStaticRoutes** | Pointer to **bool** |  | [optional] 
**HasFloatingIps** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to [**[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner**](InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner.md) |  | [optional] 
**RouteOptionTypes** | Pointer to [**[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner**](InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner.md) |  | [optional] 

## Methods

### NewInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner

`func NewInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner() *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

NewInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner instantiates a new InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults

`func NewInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults() *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

NewInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults instantiates a new InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetExternalType

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetCreatable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetOverlay

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOverlay() bool`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOverlayOk() (*bool, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetOverlay(v bool)`

SetOverlay sets Overlay field to given value.

### HasOverlay

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasOverlay() bool`

HasOverlay returns a boolean if a field has been set.

### GetNameEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameEditable() bool`

GetNameEditable returns the NameEditable field if non-nil, zero value otherwise.

### GetNameEditableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameEditableOk() (*bool, bool)`

GetNameEditableOk returns a tuple with the NameEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetNameEditable(v bool)`

SetNameEditable sets NameEditable field to given value.

### HasNameEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasNameEditable() bool`

HasNameEditable returns a boolean if a field has been set.

### GetCidrRequired

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrRequired() bool`

GetCidrRequired returns the CidrRequired field if non-nil, zero value otherwise.

### GetCidrRequiredOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrRequiredOk() (*bool, bool)`

GetCidrRequiredOk returns a tuple with the CidrRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrRequired

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCidrRequired(v bool)`

SetCidrRequired sets CidrRequired field to given value.

### HasCidrRequired

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCidrRequired() bool`

HasCidrRequired returns a boolean if a field has been set.

### GetCidrEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrEditable() bool`

GetCidrEditable returns the CidrEditable field if non-nil, zero value otherwise.

### GetCidrEditableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrEditableOk() (*bool, bool)`

GetCidrEditableOk returns a tuple with the CidrEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCidrEditable(v bool)`

SetCidrEditable sets CidrEditable field to given value.

### HasCidrEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCidrEditable() bool`

HasCidrEditable returns a boolean if a field has been set.

### GetDhcpServerEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDhcpServerEditable() bool`

GetDhcpServerEditable returns the DhcpServerEditable field if non-nil, zero value otherwise.

### GetDhcpServerEditableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDhcpServerEditableOk() (*bool, bool)`

GetDhcpServerEditableOk returns a tuple with the DhcpServerEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDhcpServerEditable(v bool)`

SetDhcpServerEditable sets DhcpServerEditable field to given value.

### HasDhcpServerEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDhcpServerEditable() bool`

HasDhcpServerEditable returns a boolean if a field has been set.

### GetDnsEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDnsEditable() bool`

GetDnsEditable returns the DnsEditable field if non-nil, zero value otherwise.

### GetDnsEditableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDnsEditableOk() (*bool, bool)`

GetDnsEditableOk returns a tuple with the DnsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDnsEditable(v bool)`

SetDnsEditable sets DnsEditable field to given value.

### HasDnsEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDnsEditable() bool`

HasDnsEditable returns a boolean if a field has been set.

### GetGatewayEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetGatewayEditable() bool`

GetGatewayEditable returns the GatewayEditable field if non-nil, zero value otherwise.

### GetGatewayEditableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetGatewayEditableOk() (*bool, bool)`

GetGatewayEditableOk returns a tuple with the GatewayEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetGatewayEditable(v bool)`

SetGatewayEditable sets GatewayEditable field to given value.

### HasGatewayEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasGatewayEditable() bool`

HasGatewayEditable returns a boolean if a field has been set.

### GetVlanIdEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetVlanIdEditable() bool`

GetVlanIdEditable returns the VlanIdEditable field if non-nil, zero value otherwise.

### GetVlanIdEditableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetVlanIdEditableOk() (*bool, bool)`

GetVlanIdEditableOk returns a tuple with the VlanIdEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIdEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetVlanIdEditable(v bool)`

SetVlanIdEditable sets VlanIdEditable field to given value.

### HasVlanIdEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasVlanIdEditable() bool`

HasVlanIdEditable returns a boolean if a field has been set.

### GetStaticOverrideEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetStaticOverrideEditable() bool`

GetStaticOverrideEditable returns the StaticOverrideEditable field if non-nil, zero value otherwise.

### GetStaticOverrideEditableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetStaticOverrideEditableOk() (*bool, bool)`

GetStaticOverrideEditableOk returns a tuple with the StaticOverrideEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticOverrideEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetStaticOverrideEditable(v bool)`

SetStaticOverrideEditable sets StaticOverrideEditable field to given value.

### HasStaticOverrideEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasStaticOverrideEditable() bool`

HasStaticOverrideEditable returns a boolean if a field has been set.

### GetNetworkDomainEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNetworkDomainEditable() bool`

GetNetworkDomainEditable returns the NetworkDomainEditable field if non-nil, zero value otherwise.

### GetNetworkDomainEditableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNetworkDomainEditableOk() (*bool, bool)`

GetNetworkDomainEditableOk returns a tuple with the NetworkDomainEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomainEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetNetworkDomainEditable(v bool)`

SetNetworkDomainEditable sets NetworkDomainEditable field to given value.

### HasNetworkDomainEditable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasNetworkDomainEditable() bool`

HasNetworkDomainEditable returns a boolean if a field has been set.

### GetCanAssignPool

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCanAssignPool() bool`

GetCanAssignPool returns the CanAssignPool field if non-nil, zero value otherwise.

### GetCanAssignPoolOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCanAssignPoolOk() (*bool, bool)`

GetCanAssignPoolOk returns a tuple with the CanAssignPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAssignPool

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCanAssignPool(v bool)`

SetCanAssignPool sets CanAssignPool field to given value.

### HasCanAssignPool

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCanAssignPool() bool`

HasCanAssignPool returns a boolean if a field has been set.

### GetDeletable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetHasNetworkServer

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasNetworkServer() bool`

GetHasNetworkServer returns the HasNetworkServer field if non-nil, zero value otherwise.

### GetHasNetworkServerOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasNetworkServerOk() (*bool, bool)`

GetHasNetworkServerOk returns a tuple with the HasNetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworkServer

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasNetworkServer(v bool)`

SetHasNetworkServer sets HasNetworkServer field to given value.

### HasHasNetworkServer

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasNetworkServer() bool`

HasHasNetworkServer returns a boolean if a field has been set.

### GetHasCidr

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasCidr() bool`

GetHasCidr returns the HasCidr field if non-nil, zero value otherwise.

### GetHasCidrOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasCidrOk() (*bool, bool)`

GetHasCidrOk returns a tuple with the HasCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCidr

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasCidr(v bool)`

SetHasCidr sets HasCidr field to given value.

### HasHasCidr

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasCidr() bool`

HasHasCidr returns a boolean if a field has been set.

### GetHasStaticRoutes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasStaticRoutes() bool`

GetHasStaticRoutes returns the HasStaticRoutes field if non-nil, zero value otherwise.

### GetHasStaticRoutesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasStaticRoutesOk() (*bool, bool)`

GetHasStaticRoutesOk returns a tuple with the HasStaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStaticRoutes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasStaticRoutes(v bool)`

SetHasStaticRoutes sets HasStaticRoutes field to given value.

### HasHasStaticRoutes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasStaticRoutes() bool`

HasHasStaticRoutes returns a boolean if a field has been set.

### GetHasFloatingIps

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasFloatingIps() bool`

GetHasFloatingIps returns the HasFloatingIps field if non-nil, zero value otherwise.

### GetHasFloatingIpsOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasFloatingIpsOk() (*bool, bool)`

GetHasFloatingIpsOk returns a tuple with the HasFloatingIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFloatingIps

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasFloatingIps(v bool)`

SetHasFloatingIps sets HasFloatingIps field to given value.

### HasHasFloatingIps

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasFloatingIps() bool`

HasHasFloatingIps returns a boolean if a field has been set.

### GetOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOptionTypes() []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOptionTypesOk() (*[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetOptionTypes(v []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetRouteOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetRouteOptionTypes() []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner`

GetRouteOptionTypes returns the RouteOptionTypes field if non-nil, zero value otherwise.

### GetRouteOptionTypesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetRouteOptionTypesOk() (*[]InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner, bool)`

GetRouteOptionTypesOk returns a tuple with the RouteOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetRouteOptionTypes(v []InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner)`

SetRouteOptionTypes sets RouteOptionTypes field to given value.

### HasRouteOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasRouteOptionTypes() bool`

HasRouteOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


