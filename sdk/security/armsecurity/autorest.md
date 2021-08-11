### AutoRest Configuration

> see https://aka.ms/autorest

``` yaml
require:
- https://github.com/Azure/azure-rest-api-specs/blob/1dc2a754513cdd59c3f5da4125bcb9f0df4e2bf3/specification/security/resource-manager/readme.md
- https://github.com/Azure/azure-rest-api-specs/blob/1dc2a754513cdd59c3f5da4125bcb9f0df4e2bf3/specification/security/resource-manager/readme.go.md
module-version: 0.1.0
diretive:
  - from: externalSecuritySolutions.json
    where: $.definitions.ExternalSecuritySolutionKind
    transform: >
      $['x-ms-client-name'] = 'ExternalSecuritySolution'
```