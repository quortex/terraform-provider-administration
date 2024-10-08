---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "administration Provider"
subcategory: ""
description: |-
  Interact with Administration.
---

# administration Provider

Interact with Administration.

## Example Usage

```terraform
# Configuration-based authentication
provider "administration" {
  client_id     = "my_client_id"
  client_secret = "my_client_secret"
  host          = "my_host"
  auth_server   = "my_auth_server"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `client_id` (String) ClientId for Administration API. May also be provided via ADMINISTRATION_CLIENT_ID environment variable.
- `client_secret` (String, Sensitive) ClientSecret for Administration API. May also be provided via ADMINISTRATION_CLIENT_SECRET environment variable.

### Optional

- `auth_server` (String) Auth server for Administration API. May also be provided via ADMINISTRATION_AUTH_SERVER environment variable.
- `host` (String) Host for Administration API. May also be provided via ADMINISTRATION_HOST environment variable.
