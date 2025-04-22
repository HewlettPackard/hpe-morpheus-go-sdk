# ListIncidents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incidents** | Pointer to [**[]GetCheckApps200ResponseOpenIncidentsInner**](GetCheckApps200ResponseOpenIncidentsInner.md) |  | [optional] 
**Issues** | Pointer to [**[]ListIncidents200ResponseAllOfIssuesInner**](ListIncidents200ResponseAllOfIssuesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListIncidents200Response

`func NewListIncidents200Response() *ListIncidents200Response`

NewListIncidents200Response instantiates a new ListIncidents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIncidents200ResponseWithDefaults

`func NewListIncidents200ResponseWithDefaults() *ListIncidents200Response`

NewListIncidents200ResponseWithDefaults instantiates a new ListIncidents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidents

`func (o *ListIncidents200Response) GetIncidents() []GetCheckApps200ResponseOpenIncidentsInner`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *ListIncidents200Response) GetIncidentsOk() (*[]GetCheckApps200ResponseOpenIncidentsInner, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *ListIncidents200Response) SetIncidents(v []GetCheckApps200ResponseOpenIncidentsInner)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *ListIncidents200Response) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetIssues

`func (o *ListIncidents200Response) GetIssues() []ListIncidents200ResponseAllOfIssuesInner`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *ListIncidents200Response) GetIssuesOk() (*[]ListIncidents200ResponseAllOfIssuesInner, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *ListIncidents200Response) SetIssues(v []ListIncidents200ResponseAllOfIssuesInner)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *ListIncidents200Response) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetMeta

`func (o *ListIncidents200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListIncidents200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListIncidents200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListIncidents200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


