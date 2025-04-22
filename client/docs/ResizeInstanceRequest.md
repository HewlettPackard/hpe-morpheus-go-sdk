# ResizeInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**ResizeInstanceRequestInstance**](ResizeInstanceRequestInstance.md) |  | [optional] 
**ServicePlanOptions** | Pointer to [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions.md) |  | [optional] 
**Volumes** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner.md) | Can be used to grow just the logical volume of the instance instead of choosing a plan | [optional] 
**DeleteOriginalVolumes** | Pointer to **bool** | Delete the original volumes after resizing. (Amazon only) | [optional] [default to false]
**NetworkInterfaces** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner.md) | Key for network configuration. Include id to update an existing interface. The existing interfaces and their id can be retrieved with the hosts API. | [optional] 

## Methods

### NewResizeInstanceRequest

`func NewResizeInstanceRequest() *ResizeInstanceRequest`

NewResizeInstanceRequest instantiates a new ResizeInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeInstanceRequestWithDefaults

`func NewResizeInstanceRequestWithDefaults() *ResizeInstanceRequest`

NewResizeInstanceRequestWithDefaults instantiates a new ResizeInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *ResizeInstanceRequest) GetInstance() ResizeInstanceRequestInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ResizeInstanceRequest) GetInstanceOk() (*ResizeInstanceRequestInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ResizeInstanceRequest) SetInstance(v ResizeInstanceRequestInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ResizeInstanceRequest) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *ResizeInstanceRequest) GetServicePlanOptions() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *ResizeInstanceRequest) GetServicePlanOptionsOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *ResizeInstanceRequest) SetServicePlanOptions(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *ResizeInstanceRequest) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetVolumes

`func (o *ResizeInstanceRequest) GetVolumes() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ResizeInstanceRequest) GetVolumesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ResizeInstanceRequest) SetVolumes(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ResizeInstanceRequest) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetDeleteOriginalVolumes

`func (o *ResizeInstanceRequest) GetDeleteOriginalVolumes() bool`

GetDeleteOriginalVolumes returns the DeleteOriginalVolumes field if non-nil, zero value otherwise.

### GetDeleteOriginalVolumesOk

`func (o *ResizeInstanceRequest) GetDeleteOriginalVolumesOk() (*bool, bool)`

GetDeleteOriginalVolumesOk returns a tuple with the DeleteOriginalVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOriginalVolumes

`func (o *ResizeInstanceRequest) SetDeleteOriginalVolumes(v bool)`

SetDeleteOriginalVolumes sets DeleteOriginalVolumes field to given value.

### HasDeleteOriginalVolumes

`func (o *ResizeInstanceRequest) HasDeleteOriginalVolumes() bool`

HasDeleteOriginalVolumes returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ResizeInstanceRequest) GetNetworkInterfaces() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ResizeInstanceRequest) GetNetworkInterfacesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ResizeInstanceRequest) SetNetworkInterfaces(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ResizeInstanceRequest) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


