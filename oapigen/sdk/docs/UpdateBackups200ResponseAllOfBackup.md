# UpdateBackups200ResponseAllOfBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**LocationType** | Pointer to **string** | Source Type (instance, server, storage) | [optional] 
**Instance** | Pointer to [**UpdateBackups200ResponseAllOfBackupInstance**](UpdateBackups200ResponseAllOfBackupInstance.md) |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**Job** | Pointer to [**UpdateBackups200ResponseAllOfBackupJob**](UpdateBackups200ResponseAllOfBackupJob.md) |  | [optional] 
**Schedule** | Pointer to [**UpdateBackups200ResponseAllOfBackupSchedule**](UpdateBackups200ResponseAllOfBackupSchedule.md) |  | [optional] 
**RetentionCount** | Pointer to **NullableInt64** |  | [optional] 
**BackupType** | Pointer to [**UpdateBackups200ResponseAllOfBackupBackupType**](UpdateBackups200ResponseAllOfBackupBackupType.md) |  | [optional] 
**StorageProvider** | Pointer to [**UpdateBackups200ResponseAllOfBackupStorageProvider**](UpdateBackups200ResponseAllOfBackupStorageProvider.md) |  | [optional] 
**BackupProvider** | Pointer to [**UpdateBackups200ResponseAllOfBackupBackupProvider**](UpdateBackups200ResponseAllOfBackupBackupProvider.md) |  | [optional] 
**BackupRespository** | Pointer to [**UpdateBackups200ResponseAllOfBackupBackupRespository**](UpdateBackups200ResponseAllOfBackupBackupRespository.md) |  | [optional] 
**CronExpression** | Pointer to **NullableString** | Cron Expression | [optional] 
**NextFire** | Pointer to **NullableTime** | Next Fire | [optional] 
**LastStatus** | Pointer to **NullableString** | Last Status | [optional] 
**LastResult** | Pointer to [**UpdateBackups200ResponseAllOfBackupLastResult**](UpdateBackups200ResponseAllOfBackupLastResult.md) |  | [optional] 
**Stats** | Pointer to [**UpdateBackups200ResponseAllOfBackupStats**](UpdateBackups200ResponseAllOfBackupStats.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Usage

Instantiate with a Go composite literal:

```go
obj := &UpdateBackups200ResponseAllOfBackup{
    // Set fields directly
}
```

### ContainerId (Nullable)

Use the Nullable wrapper methods:
- `obj.ContainerId.IsSet()` — check if set
- `obj.ContainerId.Get()` — get the inner value (returns pointer)
- `obj.ContainerId.Set(&val)` — set the value
- `obj.ContainerId.Unset()` — clear the value
### RetentionCount (Nullable)

Use the Nullable wrapper methods:
- `obj.RetentionCount.IsSet()` — check if set
- `obj.RetentionCount.Get()` — get the inner value (returns pointer)
- `obj.RetentionCount.Set(&val)` — set the value
- `obj.RetentionCount.Unset()` — clear the value
### CronExpression (Nullable)

Use the Nullable wrapper methods:
- `obj.CronExpression.IsSet()` — check if set
- `obj.CronExpression.Get()` — get the inner value (returns pointer)
- `obj.CronExpression.Set(&val)` — set the value
- `obj.CronExpression.Unset()` — clear the value
### NextFire (Nullable)

Use the Nullable wrapper methods:
- `obj.NextFire.IsSet()` — check if set
- `obj.NextFire.Get()` — get the inner value (returns pointer)
- `obj.NextFire.Set(&val)` — set the value
- `obj.NextFire.Unset()` — clear the value
### LastStatus (Nullable)

Use the Nullable wrapper methods:
- `obj.LastStatus.IsSet()` — check if set
- `obj.LastStatus.Get()` — get the inner value (returns pointer)
- `obj.LastStatus.Set(&val)` — set the value
- `obj.LastStatus.Unset()` — clear the value

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


