---
page_title: "Resource Billing Tag"
---

# Billing Tag Resource

Resource Billing Tag will be available after running BAAS service for client
## Example Usage

```
resource "ochk_billing_tag" "res-bt-cc2" {
    display_name = "billing-tag-cc2"
}
```


## Argument Reference
The following arguments are supported:
* `display_name` - (Required) The name of the billing tag
