# GetNetworkPool200ResponseNetworkPoolIpRangesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**StartAddress** | Pointer to **NullableString** |  | [optional] 
**EndAddress** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AddressCount** | Pointer to **int64** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Cidr** | Pointer to **NullableString** |  | [optional] 
**CidrIPv6** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetNetworkPool200ResponseNetworkPoolIpRangesInner

`func NewGetNetworkPool200ResponseNetworkPoolIpRangesInner() *GetNetworkPool200ResponseNetworkPoolIpRangesInner`

NewGetNetworkPool200ResponseNetworkPoolIpRangesInner instantiates a new GetNetworkPool200ResponseNetworkPoolIpRangesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkPool200ResponseNetworkPoolIpRangesInnerWithDefaults

`func NewGetNetworkPool200ResponseNetworkPoolIpRangesInnerWithDefaults() *GetNetworkPool200ResponseNetworkPoolIpRangesInner`

NewGetNetworkPool200ResponseNetworkPoolIpRangesInnerWithDefaults instantiates a new GetNetworkPool200ResponseNetworkPoolIpRangesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartAddress

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### SetStartAddressNil

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetStartAddressNil(b bool)`

 SetStartAddressNil sets the value for StartAddress to be an explicit nil

### UnsetStartAddress
`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) UnsetStartAddress()`

UnsetStartAddress ensures that no value is present for StartAddress, not even an explicit nil
### GetEndAddress

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### SetEndAddressNil

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetEndAddressNil(b bool)`

 SetEndAddressNil sets the value for EndAddress to be an explicit nil

### UnsetEndAddress
`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) UnsetEndAddress()`

UnsetEndAddress ensures that no value is present for EndAddress, not even an explicit nil
### GetInternalId

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetDescription

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAddressCount

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetAddressCount() int64`

GetAddressCount returns the AddressCount field if non-nil, zero value otherwise.

### GetAddressCountOk

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetAddressCountOk() (*int64, bool)`

GetAddressCountOk returns a tuple with the AddressCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressCount

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetAddressCount(v int64)`

SetAddressCount sets AddressCount field to given value.

### HasAddressCount

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasAddressCount() bool`

HasAddressCount returns a boolean if a field has been set.

### GetActive

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCidr

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### SetCidrNil

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetCidrNil(b bool)`

 SetCidrNil sets the value for Cidr to be an explicit nil

### UnsetCidr
`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) UnsetCidr()`

UnsetCidr ensures that no value is present for Cidr, not even an explicit nil
### GetCidrIPv6

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetCidrIPv6() string`

GetCidrIPv6 returns the CidrIPv6 field if non-nil, zero value otherwise.

### GetCidrIPv6Ok

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) GetCidrIPv6Ok() (*string, bool)`

GetCidrIPv6Ok returns a tuple with the CidrIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrIPv6

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetCidrIPv6(v string)`

SetCidrIPv6 sets CidrIPv6 field to given value.

### HasCidrIPv6

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) HasCidrIPv6() bool`

HasCidrIPv6 returns a boolean if a field has been set.

### SetCidrIPv6Nil

`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) SetCidrIPv6Nil(b bool)`

 SetCidrIPv6Nil sets the value for CidrIPv6 to be an explicit nil

### UnsetCidrIPv6
`func (o *GetNetworkPool200ResponseNetworkPoolIpRangesInner) UnsetCidrIPv6()`

UnsetCidrIPv6 ensures that no value is present for CidrIPv6, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


