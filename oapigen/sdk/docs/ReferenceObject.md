# ReferenceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Resource ID | 
**Name** | Pointer to **string** | Resource name | [optional] 
**Code** | Pointer to **string** | Resource code | [optional] 

## Methods

### NewReferenceObject

`func NewReferenceObject(id int64, ) *ReferenceObject`

NewReferenceObject instantiates a new ReferenceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceObjectWithDefaults

`func NewReferenceObjectWithDefaults() *ReferenceObject`

NewReferenceObjectWithDefaults instantiates a new ReferenceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReferenceObject) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReferenceObject) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReferenceObject) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ReferenceObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReferenceObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReferenceObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReferenceObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ReferenceObject) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReferenceObject) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReferenceObject) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ReferenceObject) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


