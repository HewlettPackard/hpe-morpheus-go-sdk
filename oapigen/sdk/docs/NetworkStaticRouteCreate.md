# NetworkStaticRouteCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Source CIDR | 
**Destination** | **string** | Destination address | 

## Methods

### NewNetworkStaticRouteCreate

`func NewNetworkStaticRouteCreate(source string, destination string, ) *NetworkStaticRouteCreate`

NewNetworkStaticRouteCreate instantiates a new NetworkStaticRouteCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkStaticRouteCreateWithDefaults

`func NewNetworkStaticRouteCreateWithDefaults() *NetworkStaticRouteCreate`

NewNetworkStaticRouteCreateWithDefaults instantiates a new NetworkStaticRouteCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *NetworkStaticRouteCreate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NetworkStaticRouteCreate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NetworkStaticRouteCreate) SetSource(v string)`

SetSource sets Source field to given value.


### GetDestination

`func (o *NetworkStaticRouteCreate) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *NetworkStaticRouteCreate) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *NetworkStaticRouteCreate) SetDestination(v string)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


