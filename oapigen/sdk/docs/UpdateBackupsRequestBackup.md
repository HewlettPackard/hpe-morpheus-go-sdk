# UpdateBackupsRequestBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the backup | [optional] 
**JobId** | Pointer to **int64** | The Backup Job ID to assign the backup to. This determines when the backup is run. | [optional] 
**Enabled** | Pointer to **bool** | Can be used to enable or disable the backup | [optional] 

## Usage

Instantiate with a Go composite literal:

```go
obj := &UpdateBackupsRequestBackup{
    // Set fields directly
}
```


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


