# ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**IacId** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**ZonePool** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**GroupLayer** | Pointer to **string** |  | [optional] 

## Methods

### NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner

`func NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner() *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner`

NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner instantiates a new ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInnerWithDefaults

`func NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInnerWithDefaults() *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner`

NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInnerWithDefaults instantiates a new ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIacId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetIacId() string`

GetIacId returns the IacId field if non-nil, zero value otherwise.

### GetIacIdOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetIacIdOk() (*string, bool)`

GetIacIdOk returns a tuple with the IacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) SetIacId(v string)`

SetIacId sets IacId field to given value.

### HasIacId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) HasIacId() bool`

HasIacId returns a boolean if a field has been set.

### GetZone

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetZone() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetZoneOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) SetZone(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetZonePool

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetZonePool() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetZonePoolOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) SetZonePool(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### GetStatus

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetGroupLayer

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetGroupLayer() string`

GetGroupLayer returns the GroupLayer field if non-nil, zero value otherwise.

### GetGroupLayerOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) GetGroupLayerOk() (*string, bool)`

GetGroupLayerOk returns a tuple with the GroupLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLayer

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) SetGroupLayer(v string)`

SetGroupLayer sets GroupLayer field to given value.

### HasGroupLayer

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerLocationsInner) HasGroupLayer() bool`

HasGroupLayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


