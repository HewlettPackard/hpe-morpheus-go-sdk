# UpdateHostResizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**UpdateHostResizeRequestServer**](UpdateHostResizeRequestServer.md) |  | [optional] 
**ServicePlanOptions** | Pointer to [**UpdateHostResizeRequestServicePlanOptions**](UpdateHostResizeRequestServicePlanOptions.md) |  | [optional] 
**Volumes** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner.md) | List of volumes with their new sizes. | [optional] 
**DeleteOriginalVolumes** | Pointer to **bool** | Delete the original volumes after resizing. (Amazon only) | [optional] [default to false]
**NetworkInterfaces** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner.md) | Key for network configurations. Include id to update an existing interface. | [optional] 

## Methods

### NewUpdateHostResizeRequest

`func NewUpdateHostResizeRequest() *UpdateHostResizeRequest`

NewUpdateHostResizeRequest instantiates a new UpdateHostResizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostResizeRequestWithDefaults

`func NewUpdateHostResizeRequestWithDefaults() *UpdateHostResizeRequest`

NewUpdateHostResizeRequestWithDefaults instantiates a new UpdateHostResizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *UpdateHostResizeRequest) GetServer() UpdateHostResizeRequestServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *UpdateHostResizeRequest) GetServerOk() (*UpdateHostResizeRequestServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *UpdateHostResizeRequest) SetServer(v UpdateHostResizeRequestServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *UpdateHostResizeRequest) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *UpdateHostResizeRequest) GetServicePlanOptions() UpdateHostResizeRequestServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *UpdateHostResizeRequest) GetServicePlanOptionsOk() (*UpdateHostResizeRequestServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *UpdateHostResizeRequest) SetServicePlanOptions(v UpdateHostResizeRequestServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *UpdateHostResizeRequest) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetVolumes

`func (o *UpdateHostResizeRequest) GetVolumes() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *UpdateHostResizeRequest) GetVolumesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *UpdateHostResizeRequest) SetVolumes(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *UpdateHostResizeRequest) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetDeleteOriginalVolumes

`func (o *UpdateHostResizeRequest) GetDeleteOriginalVolumes() bool`

GetDeleteOriginalVolumes returns the DeleteOriginalVolumes field if non-nil, zero value otherwise.

### GetDeleteOriginalVolumesOk

`func (o *UpdateHostResizeRequest) GetDeleteOriginalVolumesOk() (*bool, bool)`

GetDeleteOriginalVolumesOk returns a tuple with the DeleteOriginalVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOriginalVolumes

`func (o *UpdateHostResizeRequest) SetDeleteOriginalVolumes(v bool)`

SetDeleteOriginalVolumes sets DeleteOriginalVolumes field to given value.

### HasDeleteOriginalVolumes

`func (o *UpdateHostResizeRequest) HasDeleteOriginalVolumes() bool`

HasDeleteOriginalVolumes returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *UpdateHostResizeRequest) GetNetworkInterfaces() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *UpdateHostResizeRequest) GetNetworkInterfacesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *UpdateHostResizeRequest) SetNetworkInterfaces(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *UpdateHostResizeRequest) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


