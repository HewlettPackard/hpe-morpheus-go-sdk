# CreateNetworkPoolRequestNetworkPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Type** | Pointer to [**CreateNetworkPoolRequestNetworkPoolType**](CreateNetworkPoolRequestNetworkPoolType.md) |  | [optional] 
**IpRanges** | Pointer to [**[]CreateNetworkPoolRequestNetworkPoolIpRangesInner**](CreateNetworkPoolRequestNetworkPoolIpRangesInner.md) | Array of IP range objects. Type &#39;morpheus&#39; expects startAddress and endAddress. Type &#39;morpheusipv6&#39; expects a cidrIPv6. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by pool type. | [optional] 

## Methods

### NewCreateNetworkPoolRequestNetworkPool

`func NewCreateNetworkPoolRequestNetworkPool() *CreateNetworkPoolRequestNetworkPool`

NewCreateNetworkPoolRequestNetworkPool instantiates a new CreateNetworkPoolRequestNetworkPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkPoolRequestNetworkPoolWithDefaults

`func NewCreateNetworkPoolRequestNetworkPoolWithDefaults() *CreateNetworkPoolRequestNetworkPool`

NewCreateNetworkPoolRequestNetworkPoolWithDefaults instantiates a new CreateNetworkPoolRequestNetworkPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkPoolRequestNetworkPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkPoolRequestNetworkPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkPoolRequestNetworkPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkPoolRequestNetworkPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateNetworkPoolRequestNetworkPool) GetType() CreateNetworkPoolRequestNetworkPoolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworkPoolRequestNetworkPool) GetTypeOk() (*CreateNetworkPoolRequestNetworkPoolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworkPoolRequestNetworkPool) SetType(v CreateNetworkPoolRequestNetworkPoolType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateNetworkPoolRequestNetworkPool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIpRanges

`func (o *CreateNetworkPoolRequestNetworkPool) GetIpRanges() []CreateNetworkPoolRequestNetworkPoolIpRangesInner`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *CreateNetworkPoolRequestNetworkPool) GetIpRangesOk() (*[]CreateNetworkPoolRequestNetworkPoolIpRangesInner, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *CreateNetworkPoolRequestNetworkPool) SetIpRanges(v []CreateNetworkPoolRequestNetworkPoolIpRangesInner)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *CreateNetworkPoolRequestNetworkPool) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetConfig

`func (o *CreateNetworkPoolRequestNetworkPool) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateNetworkPoolRequestNetworkPool) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateNetworkPoolRequestNetworkPool) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateNetworkPoolRequestNetworkPool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


