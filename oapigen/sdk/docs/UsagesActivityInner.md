# UsagesActivityInner

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
**User** | Pointer to [**UsagesActivityInnerUser**](UsagesActivityInnerUser.md) |  | [optional] 
**Ts** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUsagesActivityInner

`func NewUsagesActivityInner() *UsagesActivityInner`

NewUsagesActivityInner instantiates a new UsagesActivityInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsagesActivityInnerWithDefaults

`func NewUsagesActivityInnerWithDefaults() *UsagesActivityInner`

NewUsagesActivityInnerWithDefaults instantiates a new UsagesActivityInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsagesActivityInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsagesActivityInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsagesActivityInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsagesActivityInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSuccess

`func (o *UsagesActivityInner) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UsagesActivityInner) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UsagesActivityInner) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UsagesActivityInner) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetActivityType

`func (o *UsagesActivityInner) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *UsagesActivityInner) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *UsagesActivityInner) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *UsagesActivityInner) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetName

`func (o *UsagesActivityInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsagesActivityInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsagesActivityInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsagesActivityInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessage

`func (o *UsagesActivityInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UsagesActivityInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UsagesActivityInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UsagesActivityInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetObjectType

`func (o *UsagesActivityInner) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UsagesActivityInner) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UsagesActivityInner) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *UsagesActivityInner) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *UsagesActivityInner) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *UsagesActivityInner) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *UsagesActivityInner) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *UsagesActivityInner) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetUser

`func (o *UsagesActivityInner) GetUser() UsagesActivityInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UsagesActivityInner) GetUserOk() (*UsagesActivityInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UsagesActivityInner) SetUser(v UsagesActivityInnerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UsagesActivityInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTs

`func (o *UsagesActivityInner) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *UsagesActivityInner) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *UsagesActivityInner) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *UsagesActivityInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


