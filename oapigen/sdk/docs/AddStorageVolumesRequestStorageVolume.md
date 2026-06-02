# AddStorageVolumesRequestStorageVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name scoped to your account for the storage volume | 
**Type** | **string** | Storage Type Code or ID | 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by &#x60;type&#x60;. | [optional] 
**StorageServer** | [**AddStorageVolumesRequestStorageVolumeStorageServer**](AddStorageVolumesRequestStorageVolumeStorageServer.md) |  | 
**StorageGroup** | Pointer to [**AddStorageVolumesRequestStorageVolumeStorageGroup**](AddStorageVolumesRequestStorageVolumeStorageGroup.md) |  | [optional] 

## Usage

Instantiate with a Go composite literal:

```go
obj := &AddStorageVolumesRequestStorageVolume{
    // Set fields directly
}
```


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


