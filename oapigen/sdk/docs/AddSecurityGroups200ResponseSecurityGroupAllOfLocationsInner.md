# AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**IacId** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to [**AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInnerZone**](AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInnerZone.md) |  | [optional] 
**ZonePool** | Pointer to [**AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInnerZonePool**](AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInnerZonePool.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **NullableString** |  | [optional] 
**GroupLayer** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner

`func NewAddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner() *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner`

NewAddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner instantiates a new AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIacId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetIacId() string`

GetIacId returns the IacId field if non-nil, zero value otherwise.

### GetIacIdOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetIacIdOk() (*string, bool)`

GetIacIdOk returns a tuple with the IacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetIacId(v string)`

SetIacId sets IacId field to given value.

### HasIacId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) HasIacId() bool`

HasIacId returns a boolean if a field has been set.

### SetIacIdNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetIacIdNil(b bool)`

 SetIacIdNil sets the value for IacId to be an explicit nil

### UnsetIacId
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) UnsetIacId()`

UnsetIacId ensures that no value is present for IacId, not even an explicit nil
### GetZone

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetZone() AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetZoneOk() (*AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetZone(v AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetZonePool() AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInnerZonePool`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetZonePoolOk() (*AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInnerZonePool, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetZonePool(v AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInnerZonePool)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetStatus

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetGroupLayer

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetGroupLayer() string`

GetGroupLayer returns the GroupLayer field if non-nil, zero value otherwise.

### GetGroupLayerOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) GetGroupLayerOk() (*string, bool)`

GetGroupLayerOk returns a tuple with the GroupLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLayer

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetGroupLayer(v string)`

SetGroupLayer sets GroupLayer field to given value.

### HasGroupLayer

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) HasGroupLayer() bool`

HasGroupLayer returns a boolean if a field has been set.

### SetGroupLayerNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) SetGroupLayerNil(b bool)`

 SetGroupLayerNil sets the value for GroupLayer to be an explicit nil

### UnsetGroupLayer
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfLocationsInner) UnsetGroupLayer()`

UnsetGroupLayer ensures that no value is present for GroupLayer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


