# Release History

## 0.2.0 (2022-04-15)
### Breaking Changes

- Function `*OperationsClient.List` has been removed
- Function `*GrafanaClient.ListByResourceGroup` has been removed
- Function `*GrafanaClient.List` has been removed

### Features Added

- New function `*OperationsClient.NewListPager(*OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse]`
- New function `*GrafanaClient.NewListPager(*GrafanaClientListOptions) *runtime.Pager[GrafanaClientListResponse]`
- New function `*GrafanaClient.NewListByResourceGroupPager(string, *GrafanaClientListByResourceGroupOptions) *runtime.Pager[GrafanaClientListByResourceGroupResponse]`


## 0.1.0 (2022-04-11)

- Init release.