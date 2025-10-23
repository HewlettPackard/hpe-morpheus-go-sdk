# ListVDIApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdiApps** | Pointer to [**[]ListVDIApps200ResponseAllOfVdiAppsInner**](ListVDIApps200ResponseAllOfVdiAppsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListVDIApps200Response

`func NewListVDIApps200Response() *ListVDIApps200Response`

NewListVDIApps200Response instantiates a new ListVDIApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVDIApps200ResponseWithDefaults

`func NewListVDIApps200ResponseWithDefaults() *ListVDIApps200Response`

NewListVDIApps200ResponseWithDefaults instantiates a new ListVDIApps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdiApps

`func (o *ListVDIApps200Response) GetVdiApps() []ListVDIApps200ResponseAllOfVdiAppsInner`

GetVdiApps returns the VdiApps field if non-nil, zero value otherwise.

### GetVdiAppsOk

`func (o *ListVDIApps200Response) GetVdiAppsOk() (*[]ListVDIApps200ResponseAllOfVdiAppsInner, bool)`

GetVdiAppsOk returns a tuple with the VdiApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiApps

`func (o *ListVDIApps200Response) SetVdiApps(v []ListVDIApps200ResponseAllOfVdiAppsInner)`

SetVdiApps sets VdiApps field to given value.

### HasVdiApps

`func (o *ListVDIApps200Response) HasVdiApps() bool`

HasVdiApps returns a boolean if a field has been set.

### GetMeta

`func (o *ListVDIApps200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVDIApps200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVDIApps200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVDIApps200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


