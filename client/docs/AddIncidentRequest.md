# AddIncidentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incident** | [**AddIncidentRequestIncident**](AddIncidentRequestIncident.md) |  | 

## Methods

### NewAddIncidentRequest

`func NewAddIncidentRequest(incident AddIncidentRequestIncident, ) *AddIncidentRequest`

NewAddIncidentRequest instantiates a new AddIncidentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIncidentRequestWithDefaults

`func NewAddIncidentRequestWithDefaults() *AddIncidentRequest`

NewAddIncidentRequestWithDefaults instantiates a new AddIncidentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncident

`func (o *AddIncidentRequest) GetIncident() AddIncidentRequestIncident`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *AddIncidentRequest) GetIncidentOk() (*AddIncidentRequestIncident, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *AddIncidentRequest) SetIncident(v AddIncidentRequestIncident)`

SetIncident sets Incident field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


