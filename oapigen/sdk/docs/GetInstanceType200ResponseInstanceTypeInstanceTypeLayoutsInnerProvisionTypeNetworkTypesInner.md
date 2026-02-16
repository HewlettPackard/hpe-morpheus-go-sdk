# GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner

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
**OptionTypes** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner.md) |  | [optional] 
**RouteOptionTypes** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner.md) |  | [optional] 

## Methods

### NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner

`func NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner() *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner instantiates a new GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults

`func NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults() *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults instantiates a new GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetExternalType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetCreatable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetOverlay

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOverlay() bool`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOverlayOk() (*bool, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetOverlay(v bool)`

SetOverlay sets Overlay field to given value.

### HasOverlay

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasOverlay() bool`

HasOverlay returns a boolean if a field has been set.

### GetNameEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameEditable() bool`

GetNameEditable returns the NameEditable field if non-nil, zero value otherwise.

### GetNameEditableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameEditableOk() (*bool, bool)`

GetNameEditableOk returns a tuple with the NameEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetNameEditable(v bool)`

SetNameEditable sets NameEditable field to given value.

### HasNameEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasNameEditable() bool`

HasNameEditable returns a boolean if a field has been set.

### GetCidrRequired

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrRequired() bool`

GetCidrRequired returns the CidrRequired field if non-nil, zero value otherwise.

### GetCidrRequiredOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrRequiredOk() (*bool, bool)`

GetCidrRequiredOk returns a tuple with the CidrRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrRequired

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCidrRequired(v bool)`

SetCidrRequired sets CidrRequired field to given value.

### HasCidrRequired

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCidrRequired() bool`

HasCidrRequired returns a boolean if a field has been set.

### GetCidrEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrEditable() bool`

GetCidrEditable returns the CidrEditable field if non-nil, zero value otherwise.

### GetCidrEditableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrEditableOk() (*bool, bool)`

GetCidrEditableOk returns a tuple with the CidrEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCidrEditable(v bool)`

SetCidrEditable sets CidrEditable field to given value.

### HasCidrEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCidrEditable() bool`

HasCidrEditable returns a boolean if a field has been set.

### GetDhcpServerEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDhcpServerEditable() bool`

GetDhcpServerEditable returns the DhcpServerEditable field if non-nil, zero value otherwise.

### GetDhcpServerEditableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDhcpServerEditableOk() (*bool, bool)`

GetDhcpServerEditableOk returns a tuple with the DhcpServerEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDhcpServerEditable(v bool)`

SetDhcpServerEditable sets DhcpServerEditable field to given value.

### HasDhcpServerEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDhcpServerEditable() bool`

HasDhcpServerEditable returns a boolean if a field has been set.

### GetDnsEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDnsEditable() bool`

GetDnsEditable returns the DnsEditable field if non-nil, zero value otherwise.

### GetDnsEditableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDnsEditableOk() (*bool, bool)`

GetDnsEditableOk returns a tuple with the DnsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDnsEditable(v bool)`

SetDnsEditable sets DnsEditable field to given value.

### HasDnsEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDnsEditable() bool`

HasDnsEditable returns a boolean if a field has been set.

### GetGatewayEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetGatewayEditable() bool`

GetGatewayEditable returns the GatewayEditable field if non-nil, zero value otherwise.

### GetGatewayEditableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetGatewayEditableOk() (*bool, bool)`

GetGatewayEditableOk returns a tuple with the GatewayEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetGatewayEditable(v bool)`

SetGatewayEditable sets GatewayEditable field to given value.

### HasGatewayEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasGatewayEditable() bool`

HasGatewayEditable returns a boolean if a field has been set.

### GetVlanIdEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetVlanIdEditable() bool`

GetVlanIdEditable returns the VlanIdEditable field if non-nil, zero value otherwise.

### GetVlanIdEditableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetVlanIdEditableOk() (*bool, bool)`

GetVlanIdEditableOk returns a tuple with the VlanIdEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIdEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetVlanIdEditable(v bool)`

SetVlanIdEditable sets VlanIdEditable field to given value.

### HasVlanIdEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasVlanIdEditable() bool`

HasVlanIdEditable returns a boolean if a field has been set.

### GetStaticOverrideEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetStaticOverrideEditable() bool`

GetStaticOverrideEditable returns the StaticOverrideEditable field if non-nil, zero value otherwise.

### GetStaticOverrideEditableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetStaticOverrideEditableOk() (*bool, bool)`

GetStaticOverrideEditableOk returns a tuple with the StaticOverrideEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticOverrideEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetStaticOverrideEditable(v bool)`

SetStaticOverrideEditable sets StaticOverrideEditable field to given value.

### HasStaticOverrideEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasStaticOverrideEditable() bool`

HasStaticOverrideEditable returns a boolean if a field has been set.

### GetNetworkDomainEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNetworkDomainEditable() bool`

GetNetworkDomainEditable returns the NetworkDomainEditable field if non-nil, zero value otherwise.

### GetNetworkDomainEditableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNetworkDomainEditableOk() (*bool, bool)`

GetNetworkDomainEditableOk returns a tuple with the NetworkDomainEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomainEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetNetworkDomainEditable(v bool)`

SetNetworkDomainEditable sets NetworkDomainEditable field to given value.

### HasNetworkDomainEditable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasNetworkDomainEditable() bool`

HasNetworkDomainEditable returns a boolean if a field has been set.

### GetCanAssignPool

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCanAssignPool() bool`

GetCanAssignPool returns the CanAssignPool field if non-nil, zero value otherwise.

### GetCanAssignPoolOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCanAssignPoolOk() (*bool, bool)`

GetCanAssignPoolOk returns a tuple with the CanAssignPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAssignPool

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCanAssignPool(v bool)`

SetCanAssignPool sets CanAssignPool field to given value.

### HasCanAssignPool

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCanAssignPool() bool`

HasCanAssignPool returns a boolean if a field has been set.

### GetDeletable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetHasNetworkServer

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasNetworkServer() bool`

GetHasNetworkServer returns the HasNetworkServer field if non-nil, zero value otherwise.

### GetHasNetworkServerOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasNetworkServerOk() (*bool, bool)`

GetHasNetworkServerOk returns a tuple with the HasNetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworkServer

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasNetworkServer(v bool)`

SetHasNetworkServer sets HasNetworkServer field to given value.

### HasHasNetworkServer

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasNetworkServer() bool`

HasHasNetworkServer returns a boolean if a field has been set.

### GetHasCidr

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasCidr() bool`

GetHasCidr returns the HasCidr field if non-nil, zero value otherwise.

### GetHasCidrOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasCidrOk() (*bool, bool)`

GetHasCidrOk returns a tuple with the HasCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCidr

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasCidr(v bool)`

SetHasCidr sets HasCidr field to given value.

### HasHasCidr

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasCidr() bool`

HasHasCidr returns a boolean if a field has been set.

### GetHasStaticRoutes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasStaticRoutes() bool`

GetHasStaticRoutes returns the HasStaticRoutes field if non-nil, zero value otherwise.

### GetHasStaticRoutesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasStaticRoutesOk() (*bool, bool)`

GetHasStaticRoutesOk returns a tuple with the HasStaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStaticRoutes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasStaticRoutes(v bool)`

SetHasStaticRoutes sets HasStaticRoutes field to given value.

### HasHasStaticRoutes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasStaticRoutes() bool`

HasHasStaticRoutes returns a boolean if a field has been set.

### GetHasFloatingIps

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasFloatingIps() bool`

GetHasFloatingIps returns the HasFloatingIps field if non-nil, zero value otherwise.

### GetHasFloatingIpsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasFloatingIpsOk() (*bool, bool)`

GetHasFloatingIpsOk returns a tuple with the HasFloatingIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFloatingIps

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasFloatingIps(v bool)`

SetHasFloatingIps sets HasFloatingIps field to given value.

### HasHasFloatingIps

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasFloatingIps() bool`

HasHasFloatingIps returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOptionTypes() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOptionTypesOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetOptionTypes(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetRouteOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetRouteOptionTypes() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner`

GetRouteOptionTypes returns the RouteOptionTypes field if non-nil, zero value otherwise.

### GetRouteOptionTypesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetRouteOptionTypesOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner, bool)`

GetRouteOptionTypesOk returns a tuple with the RouteOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetRouteOptionTypes(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner)`

SetRouteOptionTypes sets RouteOptionTypes field to given value.

### HasRouteOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasRouteOptionTypes() bool`

HasRouteOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


