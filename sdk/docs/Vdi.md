# Vdi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AllocationStatus** | Pointer to **string** |  | [optional] 
**Allocation** | Pointer to [**ListVdi200ResponseAllOfDesktopsInnerAllocation**](ListVdi200ResponseAllOfDesktopsInnerAllocation.md) |  | [optional] 

## Methods

### NewVdi

`func NewVdi() *Vdi`

NewVdi instantiates a new Vdi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdiWithDefaults

`func NewVdiWithDefaults() *Vdi`

NewVdiWithDefaults instantiates a new Vdi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vdi) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vdi) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vdi) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Vdi) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogo

`func (o *Vdi) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Vdi) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Vdi) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Vdi) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *Vdi) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vdi) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vdi) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vdi) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Vdi) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Vdi) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Vdi) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Vdi) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *Vdi) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Vdi) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Vdi) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Vdi) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAllocationStatus

`func (o *Vdi) GetAllocationStatus() string`

GetAllocationStatus returns the AllocationStatus field if non-nil, zero value otherwise.

### GetAllocationStatusOk

`func (o *Vdi) GetAllocationStatusOk() (*string, bool)`

GetAllocationStatusOk returns a tuple with the AllocationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationStatus

`func (o *Vdi) SetAllocationStatus(v string)`

SetAllocationStatus sets AllocationStatus field to given value.

### HasAllocationStatus

`func (o *Vdi) HasAllocationStatus() bool`

HasAllocationStatus returns a boolean if a field has been set.

### GetAllocation

`func (o *Vdi) GetAllocation() ListVdi200ResponseAllOfDesktopsInnerAllocation`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *Vdi) GetAllocationOk() (*ListVdi200ResponseAllOfDesktopsInnerAllocation, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *Vdi) SetAllocation(v ListVdi200ResponseAllOfDesktopsInnerAllocation)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *Vdi) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


