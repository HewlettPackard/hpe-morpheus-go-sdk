# ListClusterTypes200ResponseAllOfClusterTypesInner

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

### NewListClusterTypes200ResponseAllOfClusterTypesInner

`func NewListClusterTypes200ResponseAllOfClusterTypesInner() *ListClusterTypes200ResponseAllOfClusterTypesInner`

NewListClusterTypes200ResponseAllOfClusterTypesInner instantiates a new ListClusterTypes200ResponseAllOfClusterTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterTypes200ResponseAllOfClusterTypesInnerWithDefaults

`func NewListClusterTypes200ResponseAllOfClusterTypesInnerWithDefaults() *ListClusterTypes200ResponseAllOfClusterTypesInner`

NewListClusterTypes200ResponseAllOfClusterTypesInnerWithDefaults instantiates a new ListClusterTypes200ResponseAllOfClusterTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeployTargetService

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetDeployTargetService() string`

GetDeployTargetService returns the DeployTargetService field if non-nil, zero value otherwise.

### GetDeployTargetServiceOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetDeployTargetServiceOk() (*string, bool)`

GetDeployTargetServiceOk returns a tuple with the DeployTargetService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployTargetService

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetDeployTargetService(v string)`

SetDeployTargetService sets DeployTargetService field to given value.

### HasDeployTargetService

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasDeployTargetService() bool`

HasDeployTargetService returns a boolean if a field has been set.

### GetShortName

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetProviderType

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetCode

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetHostService

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHostService() string`

GetHostService returns the HostService field if non-nil, zero value otherwise.

### GetHostServiceOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHostServiceOk() (*string, bool)`

GetHostServiceOk returns a tuple with the HostService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostService

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetHostService(v string)`

SetHostService sets HostService field to given value.

### HasHostService

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasHostService() bool`

HasHostService returns a boolean if a field has been set.

### GetManaged

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetHasMasters

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHasMasters() bool`

GetHasMasters returns the HasMasters field if non-nil, zero value otherwise.

### GetHasMastersOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHasMastersOk() (*bool, bool)`

GetHasMastersOk returns a tuple with the HasMasters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMasters

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetHasMasters(v bool)`

SetHasMasters sets HasMasters field to given value.

### HasHasMasters

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasHasMasters() bool`

HasHasMasters returns a boolean if a field has been set.

### GetHasWorkers

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHasWorkers() bool`

GetHasWorkers returns the HasWorkers field if non-nil, zero value otherwise.

### GetHasWorkersOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHasWorkersOk() (*bool, bool)`

GetHasWorkersOk returns a tuple with the HasWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWorkers

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetHasWorkers(v bool)`

SetHasWorkers sets HasWorkers field to given value.

### HasHasWorkers

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasHasWorkers() bool`

HasHasWorkers returns a boolean if a field has been set.

### GetViewSet

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetViewSet() string`

GetViewSet returns the ViewSet field if non-nil, zero value otherwise.

### GetViewSetOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetViewSetOk() (*string, bool)`

GetViewSetOk returns a tuple with the ViewSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewSet

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetViewSet(v string)`

SetViewSet sets ViewSet field to given value.

### HasViewSet

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasViewSet() bool`

HasViewSet returns a boolean if a field has been set.

### GetImageCode

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetImageCode() string`

GetImageCode returns the ImageCode field if non-nil, zero value otherwise.

### GetImageCodeOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetImageCodeOk() (*string, bool)`

GetImageCodeOk returns a tuple with the ImageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCode

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetImageCode(v string)`

SetImageCode sets ImageCode field to given value.

### HasImageCode

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasImageCode() bool`

HasImageCode returns a boolean if a field has been set.

### GetKubeCtlLocal

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetKubeCtlLocal() bool`

GetKubeCtlLocal returns the KubeCtlLocal field if non-nil, zero value otherwise.

### GetKubeCtlLocalOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetKubeCtlLocalOk() (*bool, bool)`

GetKubeCtlLocalOk returns a tuple with the KubeCtlLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeCtlLocal

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetKubeCtlLocal(v bool)`

SetKubeCtlLocal sets KubeCtlLocal field to given value.

### HasKubeCtlLocal

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasKubeCtlLocal() bool`

HasKubeCtlLocal returns a boolean if a field has been set.

### GetHasDatastore

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetSupportsCloudScaling

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetSupportsCloudScaling() bool`

GetSupportsCloudScaling returns the SupportsCloudScaling field if non-nil, zero value otherwise.

### GetSupportsCloudScalingOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetSupportsCloudScalingOk() (*bool, bool)`

GetSupportsCloudScalingOk returns a tuple with the SupportsCloudScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCloudScaling

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetSupportsCloudScaling(v bool)`

SetSupportsCloudScaling sets SupportsCloudScaling field to given value.

### HasSupportsCloudScaling

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasSupportsCloudScaling() bool`

HasSupportsCloudScaling returns a boolean if a field has been set.

### GetName

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHasDefaultDataDisk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHasDefaultDataDisk() bool`

GetHasDefaultDataDisk returns the HasDefaultDataDisk field if non-nil, zero value otherwise.

### GetHasDefaultDataDiskOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHasDefaultDataDiskOk() (*bool, bool)`

GetHasDefaultDataDiskOk returns a tuple with the HasDefaultDataDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefaultDataDisk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetHasDefaultDataDisk(v bool)`

SetHasDefaultDataDisk sets HasDefaultDataDisk field to given value.

### HasHasDefaultDataDisk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasHasDefaultDataDisk() bool`

HasHasDefaultDataDisk returns a boolean if a field has been set.

### GetCanManage

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetHasCluster

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHasCluster() bool`

GetHasCluster returns the HasCluster field if non-nil, zero value otherwise.

### GetHasClusterOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetHasClusterOk() (*bool, bool)`

GetHasClusterOk returns a tuple with the HasCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCluster

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetHasCluster(v bool)`

SetHasCluster sets HasCluster field to given value.

### HasHasCluster

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasHasCluster() bool`

HasHasCluster returns a boolean if a field has been set.

### GetDescription

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetControllerTypes

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetControllerTypes() []ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner`

GetControllerTypes returns the ControllerTypes field if non-nil, zero value otherwise.

### GetControllerTypesOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetControllerTypesOk() (*[]ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner, bool)`

GetControllerTypesOk returns a tuple with the ControllerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerTypes

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetControllerTypes(v []ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner)`

SetControllerTypes sets ControllerTypes field to given value.

### HasControllerTypes

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasControllerTypes() bool`

HasControllerTypes returns a boolean if a field has been set.

### GetWorkerTypes

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetWorkerTypes() []ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner`

GetWorkerTypes returns the WorkerTypes field if non-nil, zero value otherwise.

### GetWorkerTypesOk

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) GetWorkerTypesOk() (*[]ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner, bool)`

GetWorkerTypesOk returns a tuple with the WorkerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerTypes

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) SetWorkerTypes(v []ListClusterTypes200ResponseAllOfClusterTypesInnerControllerTypesInner)`

SetWorkerTypes sets WorkerTypes field to given value.

### HasWorkerTypes

`func (o *ListClusterTypes200ResponseAllOfClusterTypesInner) HasWorkerTypes() bool`

HasWorkerTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


