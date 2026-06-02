# ListBackups200ResponseAllOfBackupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Backup ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**LocationType** | Pointer to **string** | Source Type (instance, server, storage) | [optional] 
**Instance** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerInstance**](ListBackups200ResponseAllOfBackupsInnerInstance.md) |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**Job** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerJob**](ListBackups200ResponseAllOfBackupsInnerJob.md) |  | [optional] 
**Schedule** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerSchedule**](ListBackups200ResponseAllOfBackupsInnerSchedule.md) |  | [optional] 
**RetentionCount** | Pointer to **NullableInt64** |  | [optional] 
**BackupType** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerBackupType**](ListBackups200ResponseAllOfBackupsInnerBackupType.md) |  | [optional] 
**StorageProvider** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerStorageProvider**](ListBackups200ResponseAllOfBackupsInnerStorageProvider.md) |  | [optional] 
**BackupProvider** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerBackupProvider**](ListBackups200ResponseAllOfBackupsInnerBackupProvider.md) |  | [optional] 
**BackupRespository** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerBackupRespository**](ListBackups200ResponseAllOfBackupsInnerBackupRespository.md) |  | [optional] 
**CronExpression** | Pointer to **NullableString** | Cron Expression | [optional] 
**NextFire** | Pointer to **NullableTime** | Next Fire | [optional] 
**LastStatus** | Pointer to **NullableString** | Last Status | [optional] 
**LastResult** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerLastResult**](ListBackups200ResponseAllOfBackupsInnerLastResult.md) |  | [optional] 
**Stats** | Pointer to [**ListBackups200ResponseAllOfBackupsInnerStats**](ListBackups200ResponseAllOfBackupsInnerStats.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enabled | [optional] 
**DateCreated** | Pointer to **time.Time** | Date Created | [optional] 
**LastUpdated** | Pointer to **time.Time** | Last Updated | [optional] 

## Usage

Instantiate with a Go composite literal:

```go
obj := &ListBackups200ResponseAllOfBackupsInner{
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


