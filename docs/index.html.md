# Miro Provider

The [Miro](https://miro.com/) provider is used to interact with the
Miro resources. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Datadog provider
provider "miro" {
  access_token = "${var.access_token}"
}

# Create a new Board
resource "miro_board" "myboard" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `access_token` - (Required) Access token for Miro API.
