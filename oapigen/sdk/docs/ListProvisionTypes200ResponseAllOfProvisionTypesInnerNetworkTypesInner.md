# ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner

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
**OptionTypes** | Pointer to [**[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerOptionTypesInner**](ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerOptionTypesInner.md) |  | [optional] 
**RouteOptionTypes** | Pointer to [**[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerRouteOptionTypesInner**](ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerRouteOptionTypesInner.md) |  | [optional] 

## Methods

### NewListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner

`func NewListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner() *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner`

NewListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner instantiates a new ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerWithDefaults

`func NewListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerWithDefaults() *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner`

NewListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerWithDefaults instantiates a new ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetExternalType

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetCreatable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetOverlay

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetOverlay() bool`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetOverlayOk() (*bool, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetOverlay(v bool)`

SetOverlay sets Overlay field to given value.

### HasOverlay

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasOverlay() bool`

HasOverlay returns a boolean if a field has been set.

### GetNameEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetNameEditable() bool`

GetNameEditable returns the NameEditable field if non-nil, zero value otherwise.

### GetNameEditableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetNameEditableOk() (*bool, bool)`

GetNameEditableOk returns a tuple with the NameEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetNameEditable(v bool)`

SetNameEditable sets NameEditable field to given value.

### HasNameEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasNameEditable() bool`

HasNameEditable returns a boolean if a field has been set.

### GetCidrRequired

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCidrRequired() bool`

GetCidrRequired returns the CidrRequired field if non-nil, zero value otherwise.

### GetCidrRequiredOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCidrRequiredOk() (*bool, bool)`

GetCidrRequiredOk returns a tuple with the CidrRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrRequired

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetCidrRequired(v bool)`

SetCidrRequired sets CidrRequired field to given value.

### HasCidrRequired

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasCidrRequired() bool`

HasCidrRequired returns a boolean if a field has been set.

### GetCidrEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCidrEditable() bool`

GetCidrEditable returns the CidrEditable field if non-nil, zero value otherwise.

### GetCidrEditableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCidrEditableOk() (*bool, bool)`

GetCidrEditableOk returns a tuple with the CidrEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetCidrEditable(v bool)`

SetCidrEditable sets CidrEditable field to given value.

### HasCidrEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasCidrEditable() bool`

HasCidrEditable returns a boolean if a field has been set.

### GetDhcpServerEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetDhcpServerEditable() bool`

GetDhcpServerEditable returns the DhcpServerEditable field if non-nil, zero value otherwise.

### GetDhcpServerEditableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetDhcpServerEditableOk() (*bool, bool)`

GetDhcpServerEditableOk returns a tuple with the DhcpServerEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetDhcpServerEditable(v bool)`

SetDhcpServerEditable sets DhcpServerEditable field to given value.

### HasDhcpServerEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasDhcpServerEditable() bool`

HasDhcpServerEditable returns a boolean if a field has been set.

### GetDnsEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetDnsEditable() bool`

GetDnsEditable returns the DnsEditable field if non-nil, zero value otherwise.

### GetDnsEditableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetDnsEditableOk() (*bool, bool)`

GetDnsEditableOk returns a tuple with the DnsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetDnsEditable(v bool)`

SetDnsEditable sets DnsEditable field to given value.

### HasDnsEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasDnsEditable() bool`

HasDnsEditable returns a boolean if a field has been set.

### GetGatewayEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetGatewayEditable() bool`

GetGatewayEditable returns the GatewayEditable field if non-nil, zero value otherwise.

### GetGatewayEditableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetGatewayEditableOk() (*bool, bool)`

GetGatewayEditableOk returns a tuple with the GatewayEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetGatewayEditable(v bool)`

SetGatewayEditable sets GatewayEditable field to given value.

### HasGatewayEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasGatewayEditable() bool`

HasGatewayEditable returns a boolean if a field has been set.

### GetVlanIdEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetVlanIdEditable() bool`

GetVlanIdEditable returns the VlanIdEditable field if non-nil, zero value otherwise.

### GetVlanIdEditableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetVlanIdEditableOk() (*bool, bool)`

GetVlanIdEditableOk returns a tuple with the VlanIdEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIdEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetVlanIdEditable(v bool)`

SetVlanIdEditable sets VlanIdEditable field to given value.

### HasVlanIdEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasVlanIdEditable() bool`

HasVlanIdEditable returns a boolean if a field has been set.

### GetStaticOverrideEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetStaticOverrideEditable() bool`

GetStaticOverrideEditable returns the StaticOverrideEditable field if non-nil, zero value otherwise.

### GetStaticOverrideEditableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetStaticOverrideEditableOk() (*bool, bool)`

GetStaticOverrideEditableOk returns a tuple with the StaticOverrideEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticOverrideEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetStaticOverrideEditable(v bool)`

SetStaticOverrideEditable sets StaticOverrideEditable field to given value.

### HasStaticOverrideEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasStaticOverrideEditable() bool`

HasStaticOverrideEditable returns a boolean if a field has been set.

### GetNetworkDomainEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetNetworkDomainEditable() bool`

GetNetworkDomainEditable returns the NetworkDomainEditable field if non-nil, zero value otherwise.

### GetNetworkDomainEditableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetNetworkDomainEditableOk() (*bool, bool)`

GetNetworkDomainEditableOk returns a tuple with the NetworkDomainEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomainEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetNetworkDomainEditable(v bool)`

SetNetworkDomainEditable sets NetworkDomainEditable field to given value.

### HasNetworkDomainEditable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasNetworkDomainEditable() bool`

HasNetworkDomainEditable returns a boolean if a field has been set.

### GetCanAssignPool

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCanAssignPool() bool`

GetCanAssignPool returns the CanAssignPool field if non-nil, zero value otherwise.

### GetCanAssignPoolOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetCanAssignPoolOk() (*bool, bool)`

GetCanAssignPoolOk returns a tuple with the CanAssignPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAssignPool

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetCanAssignPool(v bool)`

SetCanAssignPool sets CanAssignPool field to given value.

### HasCanAssignPool

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasCanAssignPool() bool`

HasCanAssignPool returns a boolean if a field has been set.

### GetDeletable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetHasNetworkServer

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetHasNetworkServer() bool`

GetHasNetworkServer returns the HasNetworkServer field if non-nil, zero value otherwise.

### GetHasNetworkServerOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetHasNetworkServerOk() (*bool, bool)`

GetHasNetworkServerOk returns a tuple with the HasNetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworkServer

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetHasNetworkServer(v bool)`

SetHasNetworkServer sets HasNetworkServer field to given value.

### HasHasNetworkServer

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasHasNetworkServer() bool`

HasHasNetworkServer returns a boolean if a field has been set.

### GetHasCidr

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetHasCidr() bool`

GetHasCidr returns the HasCidr field if non-nil, zero value otherwise.

### GetHasCidrOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetHasCidrOk() (*bool, bool)`

GetHasCidrOk returns a tuple with the HasCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCidr

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetHasCidr(v bool)`

SetHasCidr sets HasCidr field to given value.

### HasHasCidr

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasHasCidr() bool`

HasHasCidr returns a boolean if a field has been set.

### GetHasStaticRoutes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetHasStaticRoutes() bool`

GetHasStaticRoutes returns the HasStaticRoutes field if non-nil, zero value otherwise.

### GetHasStaticRoutesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetHasStaticRoutesOk() (*bool, bool)`

GetHasStaticRoutesOk returns a tuple with the HasStaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStaticRoutes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetHasStaticRoutes(v bool)`

SetHasStaticRoutes sets HasStaticRoutes field to given value.

### HasHasStaticRoutes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasHasStaticRoutes() bool`

HasHasStaticRoutes returns a boolean if a field has been set.

### GetHasFloatingIps

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetHasFloatingIps() bool`

GetHasFloatingIps returns the HasFloatingIps field if non-nil, zero value otherwise.

### GetHasFloatingIpsOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetHasFloatingIpsOk() (*bool, bool)`

GetHasFloatingIpsOk returns a tuple with the HasFloatingIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFloatingIps

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetHasFloatingIps(v bool)`

SetHasFloatingIps sets HasFloatingIps field to given value.

### HasHasFloatingIps

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasHasFloatingIps() bool`

HasHasFloatingIps returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetOptionTypes() []ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetOptionTypesOk() (*[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetOptionTypes(v []ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetRouteOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetRouteOptionTypes() []ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerRouteOptionTypesInner`

GetRouteOptionTypes returns the RouteOptionTypes field if non-nil, zero value otherwise.

### GetRouteOptionTypesOk

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) GetRouteOptionTypesOk() (*[]ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerRouteOptionTypesInner, bool)`

GetRouteOptionTypesOk returns a tuple with the RouteOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) SetRouteOptionTypes(v []ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInnerRouteOptionTypesInner)`

SetRouteOptionTypes sets RouteOptionTypes field to given value.

### HasRouteOptionTypes

`func (o *ListProvisionTypes200ResponseAllOfProvisionTypesInnerNetworkTypesInner) HasRouteOptionTypes() bool`

HasRouteOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


