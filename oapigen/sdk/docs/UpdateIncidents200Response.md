# UpdateIncidents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incident** | Pointer to [**UpdateIncidents200ResponseAllOfIncident**](UpdateIncidents200ResponseAllOfIncident.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateIncidents200Response

`func NewUpdateIncidents200Response() *UpdateIncidents200Response`

NewUpdateIncidents200Response instantiates a new UpdateIncidents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIncidents200ResponseWithDefaults

`func NewUpdateIncidents200ResponseWithDefaults() *UpdateIncidents200Response`

NewUpdateIncidents200ResponseWithDefaults instantiates a new UpdateIncidents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncident

`func (o *UpdateIncidents200Response) GetIncident() UpdateIncidents200ResponseAllOfIncident`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *UpdateIncidents200Response) GetIncidentOk() (*UpdateIncidents200ResponseAllOfIncident, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *UpdateIncidents200Response) SetIncident(v UpdateIncidents200ResponseAllOfIncident)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *UpdateIncidents200Response) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateIncidents200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateIncidents200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateIncidents200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateIncidents200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


