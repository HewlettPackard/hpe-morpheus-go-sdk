# ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner

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
**OptionTypes** | Pointer to [**[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner.md) |  | [optional] 
**RouteOptionTypes** | Pointer to [**[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner.md) |  | [optional] 

## Methods

### NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner

`func NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner() *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner instantiates a new ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults

`func NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults() *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerWithDefaults instantiates a new ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetExternalType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetCreatable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetOverlay

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOverlay() bool`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOverlayOk() (*bool, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetOverlay(v bool)`

SetOverlay sets Overlay field to given value.

### HasOverlay

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasOverlay() bool`

HasOverlay returns a boolean if a field has been set.

### GetNameEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameEditable() bool`

GetNameEditable returns the NameEditable field if non-nil, zero value otherwise.

### GetNameEditableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNameEditableOk() (*bool, bool)`

GetNameEditableOk returns a tuple with the NameEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetNameEditable(v bool)`

SetNameEditable sets NameEditable field to given value.

### HasNameEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasNameEditable() bool`

HasNameEditable returns a boolean if a field has been set.

### GetCidrRequired

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrRequired() bool`

GetCidrRequired returns the CidrRequired field if non-nil, zero value otherwise.

### GetCidrRequiredOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrRequiredOk() (*bool, bool)`

GetCidrRequiredOk returns a tuple with the CidrRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrRequired

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCidrRequired(v bool)`

SetCidrRequired sets CidrRequired field to given value.

### HasCidrRequired

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCidrRequired() bool`

HasCidrRequired returns a boolean if a field has been set.

### GetCidrEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrEditable() bool`

GetCidrEditable returns the CidrEditable field if non-nil, zero value otherwise.

### GetCidrEditableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCidrEditableOk() (*bool, bool)`

GetCidrEditableOk returns a tuple with the CidrEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCidrEditable(v bool)`

SetCidrEditable sets CidrEditable field to given value.

### HasCidrEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCidrEditable() bool`

HasCidrEditable returns a boolean if a field has been set.

### GetDhcpServerEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDhcpServerEditable() bool`

GetDhcpServerEditable returns the DhcpServerEditable field if non-nil, zero value otherwise.

### GetDhcpServerEditableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDhcpServerEditableOk() (*bool, bool)`

GetDhcpServerEditableOk returns a tuple with the DhcpServerEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDhcpServerEditable(v bool)`

SetDhcpServerEditable sets DhcpServerEditable field to given value.

### HasDhcpServerEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDhcpServerEditable() bool`

HasDhcpServerEditable returns a boolean if a field has been set.

### GetDnsEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDnsEditable() bool`

GetDnsEditable returns the DnsEditable field if non-nil, zero value otherwise.

### GetDnsEditableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDnsEditableOk() (*bool, bool)`

GetDnsEditableOk returns a tuple with the DnsEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDnsEditable(v bool)`

SetDnsEditable sets DnsEditable field to given value.

### HasDnsEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDnsEditable() bool`

HasDnsEditable returns a boolean if a field has been set.

### GetGatewayEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetGatewayEditable() bool`

GetGatewayEditable returns the GatewayEditable field if non-nil, zero value otherwise.

### GetGatewayEditableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetGatewayEditableOk() (*bool, bool)`

GetGatewayEditableOk returns a tuple with the GatewayEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetGatewayEditable(v bool)`

SetGatewayEditable sets GatewayEditable field to given value.

### HasGatewayEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasGatewayEditable() bool`

HasGatewayEditable returns a boolean if a field has been set.

### GetVlanIdEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetVlanIdEditable() bool`

GetVlanIdEditable returns the VlanIdEditable field if non-nil, zero value otherwise.

### GetVlanIdEditableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetVlanIdEditableOk() (*bool, bool)`

GetVlanIdEditableOk returns a tuple with the VlanIdEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIdEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetVlanIdEditable(v bool)`

SetVlanIdEditable sets VlanIdEditable field to given value.

### HasVlanIdEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasVlanIdEditable() bool`

HasVlanIdEditable returns a boolean if a field has been set.

### GetStaticOverrideEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetStaticOverrideEditable() bool`

GetStaticOverrideEditable returns the StaticOverrideEditable field if non-nil, zero value otherwise.

### GetStaticOverrideEditableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetStaticOverrideEditableOk() (*bool, bool)`

GetStaticOverrideEditableOk returns a tuple with the StaticOverrideEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticOverrideEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetStaticOverrideEditable(v bool)`

SetStaticOverrideEditable sets StaticOverrideEditable field to given value.

### HasStaticOverrideEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasStaticOverrideEditable() bool`

HasStaticOverrideEditable returns a boolean if a field has been set.

### GetNetworkDomainEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNetworkDomainEditable() bool`

GetNetworkDomainEditable returns the NetworkDomainEditable field if non-nil, zero value otherwise.

### GetNetworkDomainEditableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetNetworkDomainEditableOk() (*bool, bool)`

GetNetworkDomainEditableOk returns a tuple with the NetworkDomainEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomainEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetNetworkDomainEditable(v bool)`

SetNetworkDomainEditable sets NetworkDomainEditable field to given value.

### HasNetworkDomainEditable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasNetworkDomainEditable() bool`

HasNetworkDomainEditable returns a boolean if a field has been set.

### GetCanAssignPool

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCanAssignPool() bool`

GetCanAssignPool returns the CanAssignPool field if non-nil, zero value otherwise.

### GetCanAssignPoolOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetCanAssignPoolOk() (*bool, bool)`

GetCanAssignPoolOk returns a tuple with the CanAssignPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAssignPool

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetCanAssignPool(v bool)`

SetCanAssignPool sets CanAssignPool field to given value.

### HasCanAssignPool

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasCanAssignPool() bool`

HasCanAssignPool returns a boolean if a field has been set.

### GetDeletable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetHasNetworkServer

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasNetworkServer() bool`

GetHasNetworkServer returns the HasNetworkServer field if non-nil, zero value otherwise.

### GetHasNetworkServerOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasNetworkServerOk() (*bool, bool)`

GetHasNetworkServerOk returns a tuple with the HasNetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworkServer

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasNetworkServer(v bool)`

SetHasNetworkServer sets HasNetworkServer field to given value.

### HasHasNetworkServer

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasNetworkServer() bool`

HasHasNetworkServer returns a boolean if a field has been set.

### GetHasCidr

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasCidr() bool`

GetHasCidr returns the HasCidr field if non-nil, zero value otherwise.

### GetHasCidrOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasCidrOk() (*bool, bool)`

GetHasCidrOk returns a tuple with the HasCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCidr

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasCidr(v bool)`

SetHasCidr sets HasCidr field to given value.

### HasHasCidr

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasCidr() bool`

HasHasCidr returns a boolean if a field has been set.

### GetHasStaticRoutes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasStaticRoutes() bool`

GetHasStaticRoutes returns the HasStaticRoutes field if non-nil, zero value otherwise.

### GetHasStaticRoutesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasStaticRoutesOk() (*bool, bool)`

GetHasStaticRoutesOk returns a tuple with the HasStaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStaticRoutes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasStaticRoutes(v bool)`

SetHasStaticRoutes sets HasStaticRoutes field to given value.

### HasHasStaticRoutes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasStaticRoutes() bool`

HasHasStaticRoutes returns a boolean if a field has been set.

### GetHasFloatingIps

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasFloatingIps() bool`

GetHasFloatingIps returns the HasFloatingIps field if non-nil, zero value otherwise.

### GetHasFloatingIpsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetHasFloatingIpsOk() (*bool, bool)`

GetHasFloatingIpsOk returns a tuple with the HasFloatingIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFloatingIps

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetHasFloatingIps(v bool)`

SetHasFloatingIps sets HasFloatingIps field to given value.

### HasHasFloatingIps

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasHasFloatingIps() bool`

HasHasFloatingIps returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOptionTypes() []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetOptionTypesOk() (*[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetOptionTypes(v []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetRouteOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetRouteOptionTypes() []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner`

GetRouteOptionTypes returns the RouteOptionTypes field if non-nil, zero value otherwise.

### GetRouteOptionTypesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) GetRouteOptionTypesOk() (*[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner, bool)`

GetRouteOptionTypesOk returns a tuple with the RouteOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) SetRouteOptionTypes(v []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInnerRouteOptionTypesInner)`

SetRouteOptionTypes sets RouteOptionTypes field to given value.

### HasRouteOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner) HasRouteOptionTypes() bool`

HasRouteOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


