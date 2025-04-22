# AddIncident200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incident** | Pointer to [**GetCheckApps200ResponseOpenIncidentsInner**](GetCheckApps200ResponseOpenIncidentsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddIncident200Response

`func NewAddIncident200Response() *AddIncident200Response`

NewAddIncident200Response instantiates a new AddIncident200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIncident200ResponseWithDefaults

`func NewAddIncident200ResponseWithDefaults() *AddIncident200Response`

NewAddIncident200ResponseWithDefaults instantiates a new AddIncident200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncident

`func (o *AddIncident200Response) GetIncident() GetCheckApps200ResponseOpenIncidentsInner`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *AddIncident200Response) GetIncidentOk() (*GetCheckApps200ResponseOpenIncidentsInner, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *AddIncident200Response) SetIncident(v GetCheckApps200ResponseOpenIncidentsInner)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *AddIncident200Response) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetSuccess

`func (o *AddIncident200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddIncident200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddIncident200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddIncident200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


