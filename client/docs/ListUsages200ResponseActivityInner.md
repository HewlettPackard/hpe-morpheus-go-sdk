# ListUsages200ResponseActivityInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**ActivityType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**User** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerCreatedBy**](GetAlerts200ResponseAllOfChecksInnerCreatedBy.md) |  | [optional] 
**Ts** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListUsages200ResponseActivityInner

`func NewListUsages200ResponseActivityInner() *ListUsages200ResponseActivityInner`

NewListUsages200ResponseActivityInner instantiates a new ListUsages200ResponseActivityInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsages200ResponseActivityInnerWithDefaults

`func NewListUsages200ResponseActivityInnerWithDefaults() *ListUsages200ResponseActivityInner`

NewListUsages200ResponseActivityInnerWithDefaults instantiates a new ListUsages200ResponseActivityInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListUsages200ResponseActivityInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListUsages200ResponseActivityInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListUsages200ResponseActivityInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListUsages200ResponseActivityInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSuccess

`func (o *ListUsages200ResponseActivityInner) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListUsages200ResponseActivityInner) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListUsages200ResponseActivityInner) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListUsages200ResponseActivityInner) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetActivityType

`func (o *ListUsages200ResponseActivityInner) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *ListUsages200ResponseActivityInner) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *ListUsages200ResponseActivityInner) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *ListUsages200ResponseActivityInner) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetName

`func (o *ListUsages200ResponseActivityInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListUsages200ResponseActivityInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListUsages200ResponseActivityInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListUsages200ResponseActivityInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessage

`func (o *ListUsages200ResponseActivityInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListUsages200ResponseActivityInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListUsages200ResponseActivityInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListUsages200ResponseActivityInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetObjectType

`func (o *ListUsages200ResponseActivityInner) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ListUsages200ResponseActivityInner) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ListUsages200ResponseActivityInner) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ListUsages200ResponseActivityInner) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *ListUsages200ResponseActivityInner) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ListUsages200ResponseActivityInner) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ListUsages200ResponseActivityInner) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *ListUsages200ResponseActivityInner) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetUser

`func (o *ListUsages200ResponseActivityInner) GetUser() GetAlerts200ResponseAllOfChecksInnerCreatedBy`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListUsages200ResponseActivityInner) GetUserOk() (*GetAlerts200ResponseAllOfChecksInnerCreatedBy, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListUsages200ResponseActivityInner) SetUser(v GetAlerts200ResponseAllOfChecksInnerCreatedBy)`

SetUser sets User field to given value.

### HasUser

`func (o *ListUsages200ResponseActivityInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTs

`func (o *ListUsages200ResponseActivityInner) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ListUsages200ResponseActivityInner) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ListUsages200ResponseActivityInner) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ListUsages200ResponseActivityInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


