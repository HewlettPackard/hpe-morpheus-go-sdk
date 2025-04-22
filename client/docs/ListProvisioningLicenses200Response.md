# ListProvisioningLicenses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Licenses** | Pointer to [**[]ListProvisioningLicenses200ResponseAllOfLicensesInner**](ListProvisioningLicenses200ResponseAllOfLicensesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListProvisioningLicenses200Response

`func NewListProvisioningLicenses200Response() *ListProvisioningLicenses200Response`

NewListProvisioningLicenses200Response instantiates a new ListProvisioningLicenses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProvisioningLicenses200ResponseWithDefaults

`func NewListProvisioningLicenses200ResponseWithDefaults() *ListProvisioningLicenses200Response`

NewListProvisioningLicenses200ResponseWithDefaults instantiates a new ListProvisioningLicenses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenses

`func (o *ListProvisioningLicenses200Response) GetLicenses() []ListProvisioningLicenses200ResponseAllOfLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *ListProvisioningLicenses200Response) GetLicensesOk() (*[]ListProvisioningLicenses200ResponseAllOfLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *ListProvisioningLicenses200Response) SetLicenses(v []ListProvisioningLicenses200ResponseAllOfLicensesInner)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *ListProvisioningLicenses200Response) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMeta

`func (o *ListProvisioningLicenses200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListProvisioningLicenses200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListProvisioningLicenses200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListProvisioningLicenses200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


