---
subcategory: "Kubernetes Marketplace"
page_title: "Yandex: yandex_kubernetes_marketplace_helm_release"
description: |-
  Allows management of Kubernetes product installed from Yandex Cloud Marketplace.
  For more information, see the official documentation https://yandex.cloud/ru/marketplace?type=K8S.
---

# yandex_kubernetes_marketplace_helm_release

Allows management of Kubernetes product installed from Yandex Cloud Marketplace.
For more information, see [the official documentation](https://yandex.cloud/ru/marketplace?type=K8S).

## Example Usage

```terraform
resource "yandex_kubernetes_marketplace_helm_release" "gatekeeper_helm_release" {
  cluster_id = yandex_kubernetes_cluster.cluster_resource_name.id

  product_version = "f2ecif2vt62k2637tgus" // Gatekeeper 3.12.0

  name      = "gatekeeper"
  namespace = kubernetes_namespace.namespace_resource_name.name

  user_values = {
    auditInterval             = "90"
    constraintViolationsLimit = "30"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_id` (String) The ID of the Kubernetes cluster where the product will be installed.
- `name` (String) The name of the deployment.
- `namespace` (String) The Kubernetes namespace where the product will be installed.
- `product_version` (String) The ID of the product version to be installed.

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `user_values` (Map of String, Sensitive) Values to be passed for the installation of the product. The block consists of attributes that accept string values. The exact structure depends on the particular product and may differ for different versions of the same product. Depending on the product, some values may be required, and the installation may fail if they are not provided. **Note:** `applicationName` and `namespace`, if provided in this block, override `name` and `namespace` arguments, respectively.

### Read-Only

- `created_at` (String) The Helm Release creation (first installation) timestamp.
- `id` (String) The ID of this resource.
- `product_id` (String) The ID of the Marketplace product.
- `product_name` (String) The name of the Marketplace product.
- `status` (String) Status of the deployment.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).
- `delete` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
- `update` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours).


## Import

Import is supported using the following syntax:

```shell
terraform import yandex_kubernetes_marketpalce_helm_release.default helm_release_id
```