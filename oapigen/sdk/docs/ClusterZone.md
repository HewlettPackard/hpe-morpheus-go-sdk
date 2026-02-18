# ClusterZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ZoneType** | Pointer to [**ClusterZoneZoneType**](ClusterZoneZoneType.md) |  | [optional] 

## Methods

### NewClusterZone

`func NewClusterZone() *ClusterZone`

NewClusterZone instantiates a new ClusterZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterZoneWithDefaults

`func NewClusterZoneWithDefaults() *ClusterZone`

NewClusterZoneWithDefaults instantiates a new ClusterZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClusterZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetZoneType

`func (o *ClusterZone) GetZoneType() ClusterZoneZoneType`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *ClusterZone) GetZoneTypeOk() (*ClusterZoneZoneType, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *ClusterZone) SetZoneType(v ClusterZoneZoneType)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *ClusterZone) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


