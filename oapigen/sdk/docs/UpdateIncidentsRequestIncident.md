# UpdateIncidentsRequestIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolution** | Pointer to **string** | Description of the resolution to this incident | [optional] 
**Comment** | Pointer to **string** | Comment on this incident, updates summary field | [optional] 
**Status** | Pointer to **string** | Set status | [optional] 
**Severity** | Pointer to **string** | Set severity | [optional] 
**Name** | Pointer to **string** | Set display name | [optional] 
**StartDate** | Pointer to **time.Time** | Set start time | [optional] 
**EndDate** | Pointer to **time.Time** | Set start time | [optional] 
**InUptime** | Pointer to **bool** | Set &#39;In Availability&#39; | [optional] 

## Methods

### NewUpdateIncidentsRequestIncident

`func NewUpdateIncidentsRequestIncident() *UpdateIncidentsRequestIncident`

NewUpdateIncidentsRequestIncident instantiates a new UpdateIncidentsRequestIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIncidentsRequestIncidentWithDefaults

`func NewUpdateIncidentsRequestIncidentWithDefaults() *UpdateIncidentsRequestIncident`

NewUpdateIncidentsRequestIncidentWithDefaults instantiates a new UpdateIncidentsRequestIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolution

`func (o *UpdateIncidentsRequestIncident) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *UpdateIncidentsRequestIncident) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *UpdateIncidentsRequestIncident) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *UpdateIncidentsRequestIncident) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetComment

`func (o *UpdateIncidentsRequestIncident) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateIncidentsRequestIncident) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateIncidentsRequestIncident) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateIncidentsRequestIncident) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateIncidentsRequestIncident) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateIncidentsRequestIncident) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateIncidentsRequestIncident) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateIncidentsRequestIncident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeverity

`func (o *UpdateIncidentsRequestIncident) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *UpdateIncidentsRequestIncident) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *UpdateIncidentsRequestIncident) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *UpdateIncidentsRequestIncident) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetName

`func (o *UpdateIncidentsRequestIncident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIncidentsRequestIncident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIncidentsRequestIncident) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateIncidentsRequestIncident) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateIncidentsRequestIncident) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateIncidentsRequestIncident) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateIncidentsRequestIncident) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateIncidentsRequestIncident) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdateIncidentsRequestIncident) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateIncidentsRequestIncident) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateIncidentsRequestIncident) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateIncidentsRequestIncident) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInUptime

`func (o *UpdateIncidentsRequestIncident) GetInUptime() bool`

GetInUptime returns the InUptime field if non-nil, zero value otherwise.

### GetInUptimeOk

`func (o *UpdateIncidentsRequestIncident) GetInUptimeOk() (*bool, bool)`

GetInUptimeOk returns a tuple with the InUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUptime

`func (o *UpdateIncidentsRequestIncident) SetInUptime(v bool)`

SetInUptime sets InUptime field to given value.

### HasInUptime

`func (o *UpdateIncidentsRequestIncident) HasInUptime() bool`

HasInUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


