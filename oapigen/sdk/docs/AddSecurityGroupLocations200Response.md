# AddSecurityGroupLocations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupLocation** | Pointer to [**AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation**](AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddSecurityGroupLocations200Response

`func NewAddSecurityGroupLocations200Response() *AddSecurityGroupLocations200Response`

NewAddSecurityGroupLocations200Response instantiates a new AddSecurityGroupLocations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityGroupLocations200ResponseWithDefaults

`func NewAddSecurityGroupLocations200ResponseWithDefaults() *AddSecurityGroupLocations200Response`

NewAddSecurityGroupLocations200ResponseWithDefaults instantiates a new AddSecurityGroupLocations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupLocation

`func (o *AddSecurityGroupLocations200Response) GetSecurityGroupLocation() AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation`

GetSecurityGroupLocation returns the SecurityGroupLocation field if non-nil, zero value otherwise.

### GetSecurityGroupLocationOk

`func (o *AddSecurityGroupLocations200Response) GetSecurityGroupLocationOk() (*AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation, bool)`

GetSecurityGroupLocationOk returns a tuple with the SecurityGroupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupLocation

`func (o *AddSecurityGroupLocations200Response) SetSecurityGroupLocation(v AddSecurityGroupLocations200ResponseAllOfSecurityGroupLocation)`

SetSecurityGroupLocation sets SecurityGroupLocation field to given value.

### HasSecurityGroupLocation

`func (o *AddSecurityGroupLocations200Response) HasSecurityGroupLocation() bool`

HasSecurityGroupLocation returns a boolean if a field has been set.

### GetSuccess

`func (o *AddSecurityGroupLocations200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddSecurityGroupLocations200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddSecurityGroupLocations200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddSecurityGroupLocations200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


