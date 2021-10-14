module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql

go 1.16

require (
	github.com/Azure/azure-sdk-for-go v58.1.0+incompatible
	github.com/Azure/azure-sdk-for-go/sdk/azcore v0.19.0
)

// fix module path github.com/Azure/azure-sdk-for-go/sdk/sql/armsql to github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql
retract v0.1.0
