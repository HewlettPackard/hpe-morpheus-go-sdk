# GetNetworkRouter200ResponseNetworkRouterInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**InterfaceType** | Pointer to **NullableString** |  | [optional] 
**NetworkPosition** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**ExternalLink** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Network** | Pointer to [**GetNetworkRouter200ResponseNetworkRouterInterfacesInnerNetwork**](GetNetworkRouter200ResponseNetworkRouterInterfacesInnerNetwork.md) |  | [optional] 

## Methods

### NewGetNetworkRouter200ResponseNetworkRouterInterfacesInner

`func NewGetNetworkRouter200ResponseNetworkRouterInterfacesInner() *GetNetworkRouter200ResponseNetworkRouterInterfacesInner`

NewGetNetworkRouter200ResponseNetworkRouterInterfacesInner instantiates a new GetNetworkRouter200ResponseNetworkRouterInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRouter200ResponseNetworkRouterInterfacesInnerWithDefaults

`func NewGetNetworkRouter200ResponseNetworkRouterInterfacesInnerWithDefaults() *GetNetworkRouter200ResponseNetworkRouterInterfacesInner`

NewGetNetworkRouter200ResponseNetworkRouterInterfacesInnerWithDefaults instantiates a new GetNetworkRouter200ResponseNetworkRouterInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetInterfaceType

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### SetInterfaceTypeNil

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetInterfaceTypeNil(b bool)`

 SetInterfaceTypeNil sets the value for InterfaceType to be an explicit nil

### UnsetInterfaceType
`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) UnsetInterfaceType()`

UnsetInterfaceType ensures that no value is present for InterfaceType, not even an explicit nil
### GetNetworkPosition

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetNetworkPosition() string`

GetNetworkPosition returns the NetworkPosition field if non-nil, zero value otherwise.

### GetNetworkPositionOk

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetNetworkPositionOk() (*string, bool)`

GetNetworkPositionOk returns a tuple with the NetworkPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPosition

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetNetworkPosition(v string)`

SetNetworkPosition sets NetworkPosition field to given value.

### HasNetworkPosition

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) HasNetworkPosition() bool`

HasNetworkPosition returns a boolean if a field has been set.

### SetNetworkPositionNil

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetNetworkPositionNil(b bool)`

 SetNetworkPositionNil sets the value for NetworkPosition to be an explicit nil

### UnsetNetworkPosition
`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) UnsetNetworkPosition()`

UnsetNetworkPosition ensures that no value is present for NetworkPosition, not even an explicit nil
### GetIpAddress

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetCidr

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetExternalLink

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### SetExternalLinkNil

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetExternalLinkNil(b bool)`

 SetExternalLinkNil sets the value for ExternalLink to be an explicit nil

### UnsetExternalLink
`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) UnsetExternalLink()`

UnsetExternalLink ensures that no value is present for ExternalLink, not even an explicit nil
### GetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNetwork

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetNetwork() GetNetworkRouter200ResponseNetworkRouterInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) GetNetworkOk() (*GetNetworkRouter200ResponseNetworkRouterInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) SetNetwork(v GetNetworkRouter200ResponseNetworkRouterInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetNetworkRouter200ResponseNetworkRouterInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


