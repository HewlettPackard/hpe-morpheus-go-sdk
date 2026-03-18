# ListMigrations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Migrations** | Pointer to **interface{}** |  | [optional] 
**Meta** | Pointer to [**ListAlerts200ResponseAllOfMeta**](ListAlerts200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListMigrations200Response

`func NewListMigrations200Response() *ListMigrations200Response`

NewListMigrations200Response instantiates a new ListMigrations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMigrations200ResponseWithDefaults

`func NewListMigrations200ResponseWithDefaults() *ListMigrations200Response`

NewListMigrations200ResponseWithDefaults instantiates a new ListMigrations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrations

`func (o *ListMigrations200Response) GetMigrations() interface{}`

GetMigrations returns the Migrations field if non-nil, zero value otherwise.

### GetMigrationsOk

`func (o *ListMigrations200Response) GetMigrationsOk() (*interface{}, bool)`

GetMigrationsOk returns a tuple with the Migrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrations

`func (o *ListMigrations200Response) SetMigrations(v interface{})`

SetMigrations sets Migrations field to given value.

### HasMigrations

`func (o *ListMigrations200Response) HasMigrations() bool`

HasMigrations returns a boolean if a field has been set.

### SetMigrationsNil

`func (o *ListMigrations200Response) SetMigrationsNil(b bool)`

 SetMigrationsNil sets the value for Migrations to be an explicit nil

### UnsetMigrations
`func (o *ListMigrations200Response) UnsetMigrations()`

UnsetMigrations ensures that no value is present for Migrations, not even an explicit nil
### GetMeta

`func (o *ListMigrations200Response) GetMeta() ListAlerts200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListMigrations200Response) GetMetaOk() (*ListAlerts200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListMigrations200Response) SetMeta(v ListAlerts200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListMigrations200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


