### AutoRest Configuration

> see https://aka.ms/autorest

``` yaml
azure-arm: true
require:
- https://github.com/Azure/azure-rest-api-specs/blob/330702f6561bceb832d51f0a65090c0652f8ca9b/specification/mysql/resource-manager/readme.md
- https://github.com/Azure/azure-rest-api-specs/blob/330702f6561bceb832d51f0a65090c0652f8ca9b/specification/mysql/resource-manager/readme.go.md
license-header: MICROSOFT_MIT_NO_VERSION
module-version: 0.2.0
package-singleservers: true
directive:
- from: Servers.json
  where: $.definitions.CloudError.properties.error
  transform: >
    $["description"] = undefined
```