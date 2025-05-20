# UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Groups** | Pointer to **[]int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**ZoneTypeId** | Pointer to **int64** |  | [optional] 
**NetworkServer** | Pointer to [**UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZoneNetworkServer**](UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZoneNetworkServer.md) |  | [optional] 
**SecurityServer** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone

`func NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone() *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone`

NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone instantiates a new UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZoneWithDefaults

`func NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZoneWithDefaults() *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone`

NewUpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZoneWithDefaults instantiates a new UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetGroups

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetGroups() []int64`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetGroupsOk() (*[]int64, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) SetGroups(v []int64)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetZoneTypeId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetZoneTypeId() int64`

GetZoneTypeId returns the ZoneTypeId field if non-nil, zero value otherwise.

### GetZoneTypeIdOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetZoneTypeIdOk() (*int64, bool)`

GetZoneTypeIdOk returns a tuple with the ZoneTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTypeId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) SetZoneTypeId(v int64)`

SetZoneTypeId sets ZoneTypeId field to given value.

### HasZoneTypeId

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) HasZoneTypeId() bool`

HasZoneTypeId returns a boolean if a field has been set.

### GetNetworkServer

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetNetworkServer() UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZoneNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetNetworkServerOk() (*UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZoneNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) SetNetworkServer(v UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZoneNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetSecurityServer

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *UpdateInstanceNetworkInterface200ResponseAllOfOneOfServerZone) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


