locals {
  prefix = "${values.prefix}-${include.root.locals.ahhhhhh}"
}

include "root" {
  path = find_in_parent_folders("root.hcl")
  expose = true
}

terraform {
  source = find_in_parent_folders("test_extras/catalog/modules/service")
}

inputs = {
  prefix = local.prefix
}