# ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner

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
**OptionTypes** | Pointer to [**[]ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner.md) |  | [optional] 
**RouteOptionTypes** | Pointer to [**[]ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner.md) |  | [optional] 

## Methods

### NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner

`func NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner() *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner instantiates a new ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults

`func NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults() *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults instantiates a new ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetExternalType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetCreatable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetOverlay

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOverlay() bool`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOverlayOk() (*bool, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetOverlay(v bool)`

SetOverlay sets Overlay field to given value.

### HasOverlay

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasOverlay() bool`

HasOverlay returns a boolean if a field has been set.

### GetNameEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameEditable() bool`

GetNameEditable returns the NameEditable field if non-nil, zero value otherwise.

### GetNameEditableOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameEditableOk() (*bool, bool)`

GetNameEditableOk returns a tuple with the NameEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetNameEditable(v bool)`

SetNameEditable sets NameEditable field to given value.

### HasNameEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasNameEditable() bool`

HasNameEditable returns a boolean if a field has been set.

### GetCidrRequired

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrRequired() bool`

GetCidrRequired returns the CidrRequired field if non-nil, zero value otherwise.

### GetCidrRequiredOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrRequiredOk() (*bool, bool)`

GetCidrRequiredOk returns a tuple with the CidrRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrRequired

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCidrRequired(v bool)`

SetCidrRequired sets CidrRequired field to given value.

### HasCidrRequired

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCidrRequired() bool`

HasCidrRequired returns a boolean if a field has been set.

### GetCidrEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrEditable() bool`

GetCidrEditable returns the CidrEditable field if non-nil, zero value otherwise.

### GetCidrEditableOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrEditableOk() (*bool, bool)`

GetCidrEditableOk returns a tuple with the CidrEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCidrEditable(v bool)`

SetCidrEditable sets CidrEditable field to given value.

### HasCidrEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCidrEditable() bool`

HasCidrEditable returns a boolean if a field has been set.

### GetDhcpServerEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDhcpServerEditable() bool`

GetDhcpServerEditable returns the DhcpServerEditable field if non-nil, zero value otherwise.

### GetDhcpServerEditableOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDhcpServerEditableOk() (*bool, bool)`

GetDhcpServerEditableOk returns a tuple with the DhcpServerEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDhcpServerEditable(v bool)`

SetDhcpServerEditable sets DhcpServerEditable field to given value.

### HasDhcpServerEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDhcpServerEditable() bool`

HasDhcpServerEditable returns a boolean if a field has been set.

### GetDnsEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDnsEditable() bool`

GetDnsEditable returns the DnsEditable field if non-nil, zero value otherwise.

### GetDnsEditableOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDnsEditableOk() (*bool, bool)`

GetDnsEditableOk returns a tuple with the DnsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDnsEditable(v bool)`

SetDnsEditable sets DnsEditable field to given value.

### HasDnsEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDnsEditable() bool`

HasDnsEditable returns a boolean if a field has been set.

### GetGatewayEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetGatewayEditable() bool`

GetGatewayEditable returns the GatewayEditable field if non-nil, zero value otherwise.

### GetGatewayEditableOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetGatewayEditableOk() (*bool, bool)`

GetGatewayEditableOk returns a tuple with the GatewayEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetGatewayEditable(v bool)`

SetGatewayEditable sets GatewayEditable field to given value.

### HasGatewayEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasGatewayEditable() bool`

HasGatewayEditable returns a boolean if a field has been set.

### GetVlanIdEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetVlanIdEditable() bool`

GetVlanIdEditable returns the VlanIdEditable field if non-nil, zero value otherwise.

### GetVlanIdEditableOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetVlanIdEditableOk() (*bool, bool)`

GetVlanIdEditableOk returns a tuple with the VlanIdEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIdEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetVlanIdEditable(v bool)`

SetVlanIdEditable sets VlanIdEditable field to given value.

### HasVlanIdEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasVlanIdEditable() bool`

HasVlanIdEditable returns a boolean if a field has been set.

### GetStaticOverrideEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetStaticOverrideEditable() bool`

GetStaticOverrideEditable returns the StaticOverrideEditable field if non-nil, zero value otherwise.

### GetStaticOverrideEditableOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetStaticOverrideEditableOk() (*bool, bool)`

GetStaticOverrideEditableOk returns a tuple with the StaticOverrideEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticOverrideEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetStaticOverrideEditable(v bool)`

SetStaticOverrideEditable sets StaticOverrideEditable field to given value.

### HasStaticOverrideEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasStaticOverrideEditable() bool`

HasStaticOverrideEditable returns a boolean if a field has been set.

### GetNetworkDomainEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNetworkDomainEditable() bool`

GetNetworkDomainEditable returns the NetworkDomainEditable field if non-nil, zero value otherwise.

### GetNetworkDomainEditableOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNetworkDomainEditableOk() (*bool, bool)`

GetNetworkDomainEditableOk returns a tuple with the NetworkDomainEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomainEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetNetworkDomainEditable(v bool)`

SetNetworkDomainEditable sets NetworkDomainEditable field to given value.

### HasNetworkDomainEditable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasNetworkDomainEditable() bool`

HasNetworkDomainEditable returns a boolean if a field has been set.

### GetCanAssignPool

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCanAssignPool() bool`

GetCanAssignPool returns the CanAssignPool field if non-nil, zero value otherwise.

### GetCanAssignPoolOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCanAssignPoolOk() (*bool, bool)`

GetCanAssignPoolOk returns a tuple with the CanAssignPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAssignPool

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCanAssignPool(v bool)`

SetCanAssignPool sets CanAssignPool field to given value.

### HasCanAssignPool

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCanAssignPool() bool`

HasCanAssignPool returns a boolean if a field has been set.

### GetDeletable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetHasNetworkServer

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasNetworkServer() bool`

GetHasNetworkServer returns the HasNetworkServer field if non-nil, zero value otherwise.

### GetHasNetworkServerOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasNetworkServerOk() (*bool, bool)`

GetHasNetworkServerOk returns a tuple with the HasNetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworkServer

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasNetworkServer(v bool)`

SetHasNetworkServer sets HasNetworkServer field to given value.

### HasHasNetworkServer

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasNetworkServer() bool`

HasHasNetworkServer returns a boolean if a field has been set.

### GetHasCidr

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasCidr() bool`

GetHasCidr returns the HasCidr field if non-nil, zero value otherwise.

### GetHasCidrOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasCidrOk() (*bool, bool)`

GetHasCidrOk returns a tuple with the HasCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCidr

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasCidr(v bool)`

SetHasCidr sets HasCidr field to given value.

### HasHasCidr

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasCidr() bool`

HasHasCidr returns a boolean if a field has been set.

### GetHasStaticRoutes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasStaticRoutes() bool`

GetHasStaticRoutes returns the HasStaticRoutes field if non-nil, zero value otherwise.

### GetHasStaticRoutesOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasStaticRoutesOk() (*bool, bool)`

GetHasStaticRoutesOk returns a tuple with the HasStaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStaticRoutes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasStaticRoutes(v bool)`

SetHasStaticRoutes sets HasStaticRoutes field to given value.

### HasHasStaticRoutes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasStaticRoutes() bool`

HasHasStaticRoutes returns a boolean if a field has been set.

### GetHasFloatingIps

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasFloatingIps() bool`

GetHasFloatingIps returns the HasFloatingIps field if non-nil, zero value otherwise.

### GetHasFloatingIpsOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasFloatingIpsOk() (*bool, bool)`

GetHasFloatingIpsOk returns a tuple with the HasFloatingIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFloatingIps

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasFloatingIps(v bool)`

SetHasFloatingIps sets HasFloatingIps field to given value.

### HasHasFloatingIps

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasFloatingIps() bool`

HasHasFloatingIps returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOptionTypes() []ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOptionTypesOk() (*[]ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetOptionTypes(v []ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetRouteOptionTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetRouteOptionTypes() []ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner`

GetRouteOptionTypes returns the RouteOptionTypes field if non-nil, zero value otherwise.

### GetRouteOptionTypesOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetRouteOptionTypesOk() (*[]ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner, bool)`

GetRouteOptionTypesOk returns a tuple with the RouteOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteOptionTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetRouteOptionTypes(v []ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner)`

SetRouteOptionTypes sets RouteOptionTypes field to given value.

### HasRouteOptionTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasRouteOptionTypes() bool`

HasRouteOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


