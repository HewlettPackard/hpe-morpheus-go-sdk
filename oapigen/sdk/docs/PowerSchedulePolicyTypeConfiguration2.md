# PowerSchedulePolicyTypeConfiguration2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PowerScheduleType** | **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict schedule) | 
**PowerSchedule** | Pointer to **string** | ID of the power schedule | [optional] 
**PowerScheduleHideFixed** | Pointer to **bool** | Hide fixed schedule from users | [optional] 

## Methods

### NewPowerSchedulePolicyTypeConfiguration2

`func NewPowerSchedulePolicyTypeConfiguration2(powerScheduleType string, ) *PowerSchedulePolicyTypeConfiguration2`

NewPowerSchedulePolicyTypeConfiguration2 instantiates a new PowerSchedulePolicyTypeConfiguration2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerSchedulePolicyTypeConfiguration2WithDefaults

`func NewPowerSchedulePolicyTypeConfiguration2WithDefaults() *PowerSchedulePolicyTypeConfiguration2`

NewPowerSchedulePolicyTypeConfiguration2WithDefaults instantiates a new PowerSchedulePolicyTypeConfiguration2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPowerScheduleType

`func (o *PowerSchedulePolicyTypeConfiguration2) GetPowerScheduleType() string`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *PowerSchedulePolicyTypeConfiguration2) GetPowerScheduleTypeOk() (*string, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *PowerSchedulePolicyTypeConfiguration2) SetPowerScheduleType(v string)`

SetPowerScheduleType sets PowerScheduleType field to given value.


### GetPowerSchedule

`func (o *PowerSchedulePolicyTypeConfiguration2) GetPowerSchedule() string`

GetPowerSchedule returns the PowerSchedule field if non-nil, zero value otherwise.

### GetPowerScheduleOk

`func (o *PowerSchedulePolicyTypeConfiguration2) GetPowerScheduleOk() (*string, bool)`

GetPowerScheduleOk returns a tuple with the PowerSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSchedule

`func (o *PowerSchedulePolicyTypeConfiguration2) SetPowerSchedule(v string)`

SetPowerSchedule sets PowerSchedule field to given value.

### HasPowerSchedule

`func (o *PowerSchedulePolicyTypeConfiguration2) HasPowerSchedule() bool`

HasPowerSchedule returns a boolean if a field has been set.

### GetPowerScheduleHideFixed

`func (o *PowerSchedulePolicyTypeConfiguration2) GetPowerScheduleHideFixed() bool`

GetPowerScheduleHideFixed returns the PowerScheduleHideFixed field if non-nil, zero value otherwise.

### GetPowerScheduleHideFixedOk

`func (o *PowerSchedulePolicyTypeConfiguration2) GetPowerScheduleHideFixedOk() (*bool, bool)`

GetPowerScheduleHideFixedOk returns a tuple with the PowerScheduleHideFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleHideFixed

`func (o *PowerSchedulePolicyTypeConfiguration2) SetPowerScheduleHideFixed(v bool)`

SetPowerScheduleHideFixed sets PowerScheduleHideFixed field to given value.

### HasPowerScheduleHideFixed

`func (o *PowerSchedulePolicyTypeConfiguration2) HasPowerScheduleHideFixed() bool`

HasPowerScheduleHideFixed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


