# NetworkPoolIpCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | IP Address | [optional] 
**Hostname** | Pointer to **string** | Hostname | [optional] 
**RefType** | Pointer to **string** | Type of associated resource such as a host/server | [optional] 
**RefId** | Pointer to **int64** | ID of associated resource such as a host/server | [optional] 

## Methods

### NewNetworkPoolIpCreate

`func NewNetworkPoolIpCreate() *NetworkPoolIpCreate`

NewNetworkPoolIpCreate instantiates a new NetworkPoolIpCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPoolIpCreateWithDefaults

`func NewNetworkPoolIpCreateWithDefaults() *NetworkPoolIpCreate`

NewNetworkPoolIpCreateWithDefaults instantiates a new NetworkPoolIpCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *NetworkPoolIpCreate) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkPoolIpCreate) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkPoolIpCreate) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkPoolIpCreate) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetHostname

`func (o *NetworkPoolIpCreate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NetworkPoolIpCreate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NetworkPoolIpCreate) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NetworkPoolIpCreate) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetRefType

`func (o *NetworkPoolIpCreate) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *NetworkPoolIpCreate) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *NetworkPoolIpCreate) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *NetworkPoolIpCreate) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *NetworkPoolIpCreate) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *NetworkPoolIpCreate) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *NetworkPoolIpCreate) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *NetworkPoolIpCreate) HasRefId() bool`

HasRefId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


