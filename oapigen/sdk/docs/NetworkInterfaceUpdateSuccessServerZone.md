# NetworkInterfaceUpdateSuccessServerZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Groups** | Pointer to **[]int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**ZoneTypeId** | Pointer to **int64** |  | [optional] 
**NetworkServer** | Pointer to [**NetworkInterfaceUpdateSuccessServerZoneNetworkServer**](NetworkInterfaceUpdateSuccessServerZoneNetworkServer.md) |  | [optional] 
**SecurityServer** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNetworkInterfaceUpdateSuccessServerZone

`func NewNetworkInterfaceUpdateSuccessServerZone() *NetworkInterfaceUpdateSuccessServerZone`

NewNetworkInterfaceUpdateSuccessServerZone instantiates a new NetworkInterfaceUpdateSuccessServerZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceUpdateSuccessServerZoneWithDefaults

`func NewNetworkInterfaceUpdateSuccessServerZoneWithDefaults() *NetworkInterfaceUpdateSuccessServerZone`

NewNetworkInterfaceUpdateSuccessServerZoneWithDefaults instantiates a new NetworkInterfaceUpdateSuccessServerZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkInterfaceUpdateSuccessServerZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *NetworkInterfaceUpdateSuccessServerZone) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetGroups

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetGroups() []int64`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetGroupsOk() (*[]int64, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetGroups(v []int64)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *NetworkInterfaceUpdateSuccessServerZone) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetName

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkInterfaceUpdateSuccessServerZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkInterfaceUpdateSuccessServerZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkInterfaceUpdateSuccessServerZone) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocation

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *NetworkInterfaceUpdateSuccessServerZone) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetVisibility

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *NetworkInterfaceUpdateSuccessServerZone) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetZoneTypeId

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetZoneTypeId() int64`

GetZoneTypeId returns the ZoneTypeId field if non-nil, zero value otherwise.

### GetZoneTypeIdOk

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetZoneTypeIdOk() (*int64, bool)`

GetZoneTypeIdOk returns a tuple with the ZoneTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTypeId

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetZoneTypeId(v int64)`

SetZoneTypeId sets ZoneTypeId field to given value.

### HasZoneTypeId

`func (o *NetworkInterfaceUpdateSuccessServerZone) HasZoneTypeId() bool`

HasZoneTypeId returns a boolean if a field has been set.

### GetNetworkServer

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetNetworkServer() NetworkInterfaceUpdateSuccessServerZoneNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetNetworkServerOk() (*NetworkInterfaceUpdateSuccessServerZoneNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetNetworkServer(v NetworkInterfaceUpdateSuccessServerZoneNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *NetworkInterfaceUpdateSuccessServerZone) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetSecurityServer

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetSecurityServer() string`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *NetworkInterfaceUpdateSuccessServerZone) GetSecurityServerOk() (*string, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetSecurityServer(v string)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *NetworkInterfaceUpdateSuccessServerZone) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### SetSecurityServerNil

`func (o *NetworkInterfaceUpdateSuccessServerZone) SetSecurityServerNil(b bool)`

 SetSecurityServerNil sets the value for SecurityServer to be an explicit nil

### UnsetSecurityServer
`func (o *NetworkInterfaceUpdateSuccessServerZone) UnsetSecurityServer()`

UnsetSecurityServer ensures that no value is present for SecurityServer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


