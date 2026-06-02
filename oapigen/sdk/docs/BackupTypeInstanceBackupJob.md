# BackupTypeInstanceBackupJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyntheticFullEnabled** | Pointer to **bool** | Enable synthetic full backups on this backup jobAction. Only applies to &#x60;kvm&#x60; backup type. | [optional] 
**SyntheticFullSchedule** | Pointer to **int64** | the ID of the execute schedule for the synthetic full backup to be created. Only applies to &#x60;kvm&#x60; backup type. | [optional] 

## Usage

Instantiate with a Go composite literal:

```go
obj := &BackupTypeInstanceBackupJob{
    // Set fields directly
}
```


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


