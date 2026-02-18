# ClustersZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ZoneType** | Pointer to [**ClustersZoneZoneType**](ClustersZoneZoneType.md) |  | [optional] 

## Methods

### NewClustersZone

`func NewClustersZone() *ClustersZone`

NewClustersZone instantiates a new ClustersZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClustersZoneWithDefaults

`func NewClustersZoneWithDefaults() *ClustersZone`

NewClustersZoneWithDefaults instantiates a new ClustersZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClustersZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClustersZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClustersZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClustersZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClustersZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClustersZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClustersZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClustersZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZoneType

`func (o *ClustersZone) GetZoneType() ClustersZoneZoneType`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *ClustersZone) GetZoneTypeOk() (*ClustersZoneZoneType, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *ClustersZone) SetZoneType(v ClustersZoneZoneType)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *ClustersZone) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


