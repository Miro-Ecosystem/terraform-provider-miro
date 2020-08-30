# miro_board Resource

Provides a Miro board resource.

## Example Usage

```hcl
# Create a new Miro board
resource "miro_board" "myboard" {
    name        = "My board"
    description = "Testing with Miro provider"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Pipeline name.
* `description` - (Optional) Pipeline JSON content.

## Import

Applications can be imported using their Spinnaker application and pipeline name, e.g.

```
$ terraform import miro_.myboard o9J_km_OSX8=
```
