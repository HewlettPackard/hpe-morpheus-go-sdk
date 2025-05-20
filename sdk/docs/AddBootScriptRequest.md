# AddBootScriptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootScript** | Pointer to [**AddBootScriptRequestBootScript**](AddBootScriptRequestBootScript.md) |  | [optional] 

## Methods

### NewAddBootScriptRequest

`func NewAddBootScriptRequest() *AddBootScriptRequest`

NewAddBootScriptRequest instantiates a new AddBootScriptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBootScriptRequestWithDefaults

`func NewAddBootScriptRequestWithDefaults() *AddBootScriptRequest`

NewAddBootScriptRequestWithDefaults instantiates a new AddBootScriptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootScript

`func (o *AddBootScriptRequest) GetBootScript() AddBootScriptRequestBootScript`

GetBootScript returns the BootScript field if non-nil, zero value otherwise.

### GetBootScriptOk

`func (o *AddBootScriptRequest) GetBootScriptOk() (*AddBootScriptRequestBootScript, bool)`

GetBootScriptOk returns a tuple with the BootScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScript

`func (o *AddBootScriptRequest) SetBootScript(v AddBootScriptRequestBootScript)`

SetBootScript sets BootScript field to given value.

### HasBootScript

`func (o *AddBootScriptRequest) HasBootScript() bool`

HasBootScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


