# GetHostType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerType** | Pointer to [**ListHostTypes200ResponseAllOfServerTypesInner**](ListHostTypes200ResponseAllOfServerTypesInner.md) |  | [optional] 

## Methods

### NewGetHostType200Response

`func NewGetHostType200Response() *GetHostType200Response`

NewGetHostType200Response instantiates a new GetHostType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHostType200ResponseWithDefaults

`func NewGetHostType200ResponseWithDefaults() *GetHostType200Response`

NewGetHostType200ResponseWithDefaults instantiates a new GetHostType200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerType

`func (o *GetHostType200Response) GetServerType() ListHostTypes200ResponseAllOfServerTypesInner`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *GetHostType200Response) GetServerTypeOk() (*ListHostTypes200ResponseAllOfServerTypesInner, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *GetHostType200Response) SetServerType(v ListHostTypes200ResponseAllOfServerTypesInner)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *GetHostType200Response) HasServerType() bool`

HasServerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


