---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "administration_billing_plan Resource - administration"
subcategory: ""
description: |-
  Manages a plan.
---

# administration_billing_plan (Resource)

Manages a plan.

## Example Usage

```terraform
# Manage example order.
resource "administration_billing_plan" "premium" {
  name = "premium"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `limits` (Attributes List) List of limits of the plan. (see [below for nested schema](#nestedatt--limits))
- `name` (String) Name of the plan.
- `pricing` (Attributes List) List of pricing of the plan. (see [below for nested schema](#nestedatt--pricing))

### Optional

- `features` (List of String) List of features of the plan.

### Read-Only

- `id` (String) Numeric identifier of the plan.
- `last_updated` (String)

<a id="nestedatt--limits"></a>
### Nested Schema for `limits`

Required:

- `name` (String) Name of limit.
- `value` (Number) Value of limit.


<a id="nestedatt--pricing"></a>
### Nested Schema for `pricing`

Required:

- `monthly_price` (Number) Monthly pricing.
- `monthly_price_currency` (String) Monthly currency.
- `subscribe_for_year` (Number) Number of year of subscription.

## Import

Import is supported using the following syntax:

```shell
# Order can be imported by specifying the numeric identifier.
terraform import administration_billing_plan.premium 123
```
