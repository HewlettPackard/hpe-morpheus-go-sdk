# ZoneResourcePoolConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | **string** | Provide the base CIDR Block to use for this VPC (must be between a /16 and /28 Block) | 
**Tenancy** | **string** |  | [default to "default"]

## Methods

### NewZoneResourcePoolConfigAnyOf

`func NewZoneResourcePoolConfigAnyOf(cidrBlock string, tenancy string, ) *ZoneResourcePoolConfigAnyOf`

NewZoneResourcePoolConfigAnyOf instantiates a new ZoneResourcePoolConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneResourcePoolConfigAnyOfWithDefaults

`func NewZoneResourcePoolConfigAnyOfWithDefaults() *ZoneResourcePoolConfigAnyOf`

NewZoneResourcePoolConfigAnyOfWithDefaults instantiates a new ZoneResourcePoolConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *ZoneResourcePoolConfigAnyOf) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *ZoneResourcePoolConfigAnyOf) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *ZoneResourcePoolConfigAnyOf) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetTenancy

`func (o *ZoneResourcePoolConfigAnyOf) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *ZoneResourcePoolConfigAnyOf) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *ZoneResourcePoolConfigAnyOf) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


