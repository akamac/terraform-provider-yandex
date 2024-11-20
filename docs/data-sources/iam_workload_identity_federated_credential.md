---
subcategory: "Identity and Access Management (IAM)"
page_title: "Yandex: yandex_iam_workload_identity_federated_credential"
description: |-
  Get information about a Yandex IAM federated credential.
---

# yandex_iam_workload_identity_federated_credential

Get information about a [Yandex.Cloud IAM federated credential](https://yandex.cloud/docs/iam/concepts/workload-identity#federated-credentials).

## Example Usage

```terraform
data "yandex_iam_workload_identity_federated_credential" "fc" {
  federated_credential_id = "some_fed_cred_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `federated_credential_id` (String) Id of the federated credential.

### Read-Only

- `created_at` (String) Creation timestamp.
- `external_subject_id` (String) Id of the external subject.
- `federation_id` (String) Id of the workload identity federation which is used for authentication.
- `id` (String) The ID of this resource.
- `service_account_id` (String) Id of the service account that the federated credential belongs to.

