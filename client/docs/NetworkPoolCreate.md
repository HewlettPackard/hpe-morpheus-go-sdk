# NetworkPoolCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Type** | Pointer to [**CreateNetworkPoolRequestNetworkPoolType**](CreateNetworkPoolRequestNetworkPoolType.md) |  | [optional] 
**IpRanges** | Pointer to [**[]CreateNetworkPoolRequestNetworkPoolIpRangesInner**](CreateNetworkPoolRequestNetworkPoolIpRangesInner.md) | Array of IP range objects. Type &#39;morpheus&#39; expects startAddress and endAddress. Type &#39;morpheusipv6&#39; expects a cidrIPv6. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by pool type. | [optional] 

## Methods

### NewNetworkPoolCreate

`func NewNetworkPoolCreate() *NetworkPoolCreate`

NewNetworkPoolCreate instantiates a new NetworkPoolCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPoolCreateWithDefaults

`func NewNetworkPoolCreateWithDefaults() *NetworkPoolCreate`

NewNetworkPoolCreateWithDefaults instantiates a new NetworkPoolCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkPoolCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPoolCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPoolCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkPoolCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *NetworkPoolCreate) GetType() CreateNetworkPoolRequestNetworkPoolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkPoolCreate) GetTypeOk() (*CreateNetworkPoolRequestNetworkPoolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkPoolCreate) SetType(v CreateNetworkPoolRequestNetworkPoolType)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkPoolCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIpRanges

`func (o *NetworkPoolCreate) GetIpRanges() []CreateNetworkPoolRequestNetworkPoolIpRangesInner`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *NetworkPoolCreate) GetIpRangesOk() (*[]CreateNetworkPoolRequestNetworkPoolIpRangesInner, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *NetworkPoolCreate) SetIpRanges(v []CreateNetworkPoolRequestNetworkPoolIpRangesInner)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *NetworkPoolCreate) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkPoolCreate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkPoolCreate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkPoolCreate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkPoolCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


