# CatalogItemTypeInstanceScribePortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int64** | Port number. | [optional] 
**Name** | Pointer to **string** | A name for the port. | [optional] 
**Lb** | Pointer to **NullableString** | The load balancer protocol. HTTP, HTTPS, or TCP. | [optional] 

## Methods

### NewCatalogItemTypeInstanceScribePortsInner

`func NewCatalogItemTypeInstanceScribePortsInner() *CatalogItemTypeInstanceScribePortsInner`

NewCatalogItemTypeInstanceScribePortsInner instantiates a new CatalogItemTypeInstanceScribePortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemTypeInstanceScribePortsInnerWithDefaults

`func NewCatalogItemTypeInstanceScribePortsInnerWithDefaults() *CatalogItemTypeInstanceScribePortsInner`

NewCatalogItemTypeInstanceScribePortsInnerWithDefaults instantiates a new CatalogItemTypeInstanceScribePortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *CatalogItemTypeInstanceScribePortsInner) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CatalogItemTypeInstanceScribePortsInner) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CatalogItemTypeInstanceScribePortsInner) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *CatalogItemTypeInstanceScribePortsInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetName

`func (o *CatalogItemTypeInstanceScribePortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogItemTypeInstanceScribePortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogItemTypeInstanceScribePortsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogItemTypeInstanceScribePortsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLb

`func (o *CatalogItemTypeInstanceScribePortsInner) GetLb() string`

GetLb returns the Lb field if non-nil, zero value otherwise.

### GetLbOk

`func (o *CatalogItemTypeInstanceScribePortsInner) GetLbOk() (*string, bool)`

GetLbOk returns a tuple with the Lb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLb

`func (o *CatalogItemTypeInstanceScribePortsInner) SetLb(v string)`

SetLb sets Lb field to given value.

### HasLb

`func (o *CatalogItemTypeInstanceScribePortsInner) HasLb() bool`

HasLb returns a boolean if a field has been set.

### SetLbNil

`func (o *CatalogItemTypeInstanceScribePortsInner) SetLbNil(b bool)`

 SetLbNil sets the value for Lb to be an explicit nil

### UnsetLb
`func (o *CatalogItemTypeInstanceScribePortsInner) UnsetLb()`

UnsetLb ensures that no value is present for Lb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


