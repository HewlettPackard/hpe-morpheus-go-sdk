# AddVirtualImage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualImage** | Pointer to [**ListVirtualImages200ResponseAllOfVirtualImagesInner**](ListVirtualImages200ResponseAllOfVirtualImagesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddVirtualImage200Response

`func NewAddVirtualImage200Response() *AddVirtualImage200Response`

NewAddVirtualImage200Response instantiates a new AddVirtualImage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVirtualImage200ResponseWithDefaults

`func NewAddVirtualImage200ResponseWithDefaults() *AddVirtualImage200Response`

NewAddVirtualImage200ResponseWithDefaults instantiates a new AddVirtualImage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualImage

`func (o *AddVirtualImage200Response) GetVirtualImage() ListVirtualImages200ResponseAllOfVirtualImagesInner`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *AddVirtualImage200Response) GetVirtualImageOk() (*ListVirtualImages200ResponseAllOfVirtualImagesInner, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *AddVirtualImage200Response) SetVirtualImage(v ListVirtualImages200ResponseAllOfVirtualImagesInner)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *AddVirtualImage200Response) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.

### GetSuccess

`func (o *AddVirtualImage200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddVirtualImage200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddVirtualImage200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddVirtualImage200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


