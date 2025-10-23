# ClusterTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DeployTargetService** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**ProviderType** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**HostService** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**HasMasters** | Pointer to **bool** |  | [optional] 
**HasWorkers** | Pointer to **bool** |  | [optional] 
**ViewSet** | Pointer to **string** |  | [optional] 
**ImageCode** | Pointer to **string** |  | [optional] 
**KubeCtlLocal** | Pointer to **bool** |  | [optional] 
**HasDatastore** | Pointer to **bool** |  | [optional] 
**SupportsCloudScaling** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**HasDefaultDataDisk** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**HasCluster** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**ControllerTypes** | Pointer to [**[]ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner**](ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner.md) |  | [optional] 
**WorkerTypes** | Pointer to [**[]ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner**](ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner.md) |  | [optional] 

## Methods

### NewClusterTypes

`func NewClusterTypes() *ClusterTypes`

NewClusterTypes instantiates a new ClusterTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterTypesWithDefaults

`func NewClusterTypesWithDefaults() *ClusterTypes`

NewClusterTypesWithDefaults instantiates a new ClusterTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterTypes) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterTypes) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterTypes) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterTypes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeployTargetService

`func (o *ClusterTypes) GetDeployTargetService() string`

GetDeployTargetService returns the DeployTargetService field if non-nil, zero value otherwise.

### GetDeployTargetServiceOk

`func (o *ClusterTypes) GetDeployTargetServiceOk() (*string, bool)`

GetDeployTargetServiceOk returns a tuple with the DeployTargetService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployTargetService

`func (o *ClusterTypes) SetDeployTargetService(v string)`

SetDeployTargetService sets DeployTargetService field to given value.

### HasDeployTargetService

`func (o *ClusterTypes) HasDeployTargetService() bool`

HasDeployTargetService returns a boolean if a field has been set.

### GetShortName

`func (o *ClusterTypes) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ClusterTypes) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ClusterTypes) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ClusterTypes) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetProviderType

`func (o *ClusterTypes) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *ClusterTypes) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *ClusterTypes) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *ClusterTypes) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetCode

`func (o *ClusterTypes) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClusterTypes) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClusterTypes) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClusterTypes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetHostService

`func (o *ClusterTypes) GetHostService() string`

GetHostService returns the HostService field if non-nil, zero value otherwise.

### GetHostServiceOk

`func (o *ClusterTypes) GetHostServiceOk() (*string, bool)`

GetHostServiceOk returns a tuple with the HostService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostService

`func (o *ClusterTypes) SetHostService(v string)`

SetHostService sets HostService field to given value.

### HasHostService

`func (o *ClusterTypes) HasHostService() bool`

HasHostService returns a boolean if a field has been set.

### GetManaged

`func (o *ClusterTypes) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ClusterTypes) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ClusterTypes) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ClusterTypes) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetHasMasters

`func (o *ClusterTypes) GetHasMasters() bool`

GetHasMasters returns the HasMasters field if non-nil, zero value otherwise.

### GetHasMastersOk

`func (o *ClusterTypes) GetHasMastersOk() (*bool, bool)`

GetHasMastersOk returns a tuple with the HasMasters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMasters

`func (o *ClusterTypes) SetHasMasters(v bool)`

SetHasMasters sets HasMasters field to given value.

### HasHasMasters

`func (o *ClusterTypes) HasHasMasters() bool`

HasHasMasters returns a boolean if a field has been set.

### GetHasWorkers

`func (o *ClusterTypes) GetHasWorkers() bool`

GetHasWorkers returns the HasWorkers field if non-nil, zero value otherwise.

### GetHasWorkersOk

`func (o *ClusterTypes) GetHasWorkersOk() (*bool, bool)`

GetHasWorkersOk returns a tuple with the HasWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWorkers

`func (o *ClusterTypes) SetHasWorkers(v bool)`

SetHasWorkers sets HasWorkers field to given value.

### HasHasWorkers

`func (o *ClusterTypes) HasHasWorkers() bool`

HasHasWorkers returns a boolean if a field has been set.

### GetViewSet

`func (o *ClusterTypes) GetViewSet() string`

GetViewSet returns the ViewSet field if non-nil, zero value otherwise.

### GetViewSetOk

`func (o *ClusterTypes) GetViewSetOk() (*string, bool)`

GetViewSetOk returns a tuple with the ViewSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewSet

`func (o *ClusterTypes) SetViewSet(v string)`

SetViewSet sets ViewSet field to given value.

### HasViewSet

`func (o *ClusterTypes) HasViewSet() bool`

HasViewSet returns a boolean if a field has been set.

### GetImageCode

`func (o *ClusterTypes) GetImageCode() string`

GetImageCode returns the ImageCode field if non-nil, zero value otherwise.

### GetImageCodeOk

`func (o *ClusterTypes) GetImageCodeOk() (*string, bool)`

GetImageCodeOk returns a tuple with the ImageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCode

`func (o *ClusterTypes) SetImageCode(v string)`

SetImageCode sets ImageCode field to given value.

### HasImageCode

`func (o *ClusterTypes) HasImageCode() bool`

HasImageCode returns a boolean if a field has been set.

### GetKubeCtlLocal

`func (o *ClusterTypes) GetKubeCtlLocal() bool`

GetKubeCtlLocal returns the KubeCtlLocal field if non-nil, zero value otherwise.

### GetKubeCtlLocalOk

`func (o *ClusterTypes) GetKubeCtlLocalOk() (*bool, bool)`

GetKubeCtlLocalOk returns a tuple with the KubeCtlLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeCtlLocal

`func (o *ClusterTypes) SetKubeCtlLocal(v bool)`

SetKubeCtlLocal sets KubeCtlLocal field to given value.

### HasKubeCtlLocal

`func (o *ClusterTypes) HasKubeCtlLocal() bool`

HasKubeCtlLocal returns a boolean if a field has been set.

### GetHasDatastore

`func (o *ClusterTypes) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *ClusterTypes) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *ClusterTypes) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *ClusterTypes) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetSupportsCloudScaling

`func (o *ClusterTypes) GetSupportsCloudScaling() bool`

GetSupportsCloudScaling returns the SupportsCloudScaling field if non-nil, zero value otherwise.

### GetSupportsCloudScalingOk

`func (o *ClusterTypes) GetSupportsCloudScalingOk() (*bool, bool)`

GetSupportsCloudScalingOk returns a tuple with the SupportsCloudScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCloudScaling

`func (o *ClusterTypes) SetSupportsCloudScaling(v bool)`

SetSupportsCloudScaling sets SupportsCloudScaling field to given value.

### HasSupportsCloudScaling

`func (o *ClusterTypes) HasSupportsCloudScaling() bool`

HasSupportsCloudScaling returns a boolean if a field has been set.

### GetName

`func (o *ClusterTypes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterTypes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterTypes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterTypes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHasDefaultDataDisk

`func (o *ClusterTypes) GetHasDefaultDataDisk() bool`

GetHasDefaultDataDisk returns the HasDefaultDataDisk field if non-nil, zero value otherwise.

### GetHasDefaultDataDiskOk

`func (o *ClusterTypes) GetHasDefaultDataDiskOk() (*bool, bool)`

GetHasDefaultDataDiskOk returns a tuple with the HasDefaultDataDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefaultDataDisk

`func (o *ClusterTypes) SetHasDefaultDataDisk(v bool)`

SetHasDefaultDataDisk sets HasDefaultDataDisk field to given value.

### HasHasDefaultDataDisk

`func (o *ClusterTypes) HasHasDefaultDataDisk() bool`

HasHasDefaultDataDisk returns a boolean if a field has been set.

### GetCanManage

`func (o *ClusterTypes) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ClusterTypes) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ClusterTypes) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ClusterTypes) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetHasCluster

`func (o *ClusterTypes) GetHasCluster() bool`

GetHasCluster returns the HasCluster field if non-nil, zero value otherwise.

### GetHasClusterOk

`func (o *ClusterTypes) GetHasClusterOk() (*bool, bool)`

GetHasClusterOk returns a tuple with the HasCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCluster

`func (o *ClusterTypes) SetHasCluster(v bool)`

SetHasCluster sets HasCluster field to given value.

### HasHasCluster

`func (o *ClusterTypes) HasHasCluster() bool`

HasHasCluster returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterTypes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterTypes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterTypes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterTypes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ClusterTypes) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ClusterTypes) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ClusterTypes) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ClusterTypes) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetControllerTypes

`func (o *ClusterTypes) GetControllerTypes() []ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner`

GetControllerTypes returns the ControllerTypes field if non-nil, zero value otherwise.

### GetControllerTypesOk

`func (o *ClusterTypes) GetControllerTypesOk() (*[]ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner, bool)`

GetControllerTypesOk returns a tuple with the ControllerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerTypes

`func (o *ClusterTypes) SetControllerTypes(v []ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner)`

SetControllerTypes sets ControllerTypes field to given value.

### HasControllerTypes

`func (o *ClusterTypes) HasControllerTypes() bool`

HasControllerTypes returns a boolean if a field has been set.

### GetWorkerTypes

`func (o *ClusterTypes) GetWorkerTypes() []ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner`

GetWorkerTypes returns the WorkerTypes field if non-nil, zero value otherwise.

### GetWorkerTypesOk

`func (o *ClusterTypes) GetWorkerTypesOk() (*[]ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner, bool)`

GetWorkerTypesOk returns a tuple with the WorkerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerTypes

`func (o *ClusterTypes) SetWorkerTypes(v []ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner)`

SetWorkerTypes sets WorkerTypes field to given value.

### HasWorkerTypes

`func (o *ClusterTypes) HasWorkerTypes() bool`

HasWorkerTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


