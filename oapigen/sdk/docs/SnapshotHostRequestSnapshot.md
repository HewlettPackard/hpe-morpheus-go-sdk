# SnapshotHostRequestSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Optional name for the snapshot being created. | [optional] [default to "{serverName}.{timestamp}"]
**Description** | Pointer to **string** | Optional description for the snapshot | [optional] 
**MemorySnapshot** | Pointer to **bool** | Whether to include the memory state in the snapshot. Only supported by certain provision types such as VMware | [optional] [default to false]
**ForExport** | Pointer to **bool** | For Export? Indicates the snapshot is intended for export to storage. | [optional] [default to false]

## Usage

Instantiate with a Go composite literal:

```go
obj := &SnapshotHostRequestSnapshot{
    // Set fields directly
}
```


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


