locals {
  prefix = "${values.prefix}-${include.root.locals.ahhhhhh}"
}

include "root" {
  path = find_in_parent_folders("root.hcl")
  expose = true
}

terraform {
  source = "."
}

inputs = {
  prefix = local.prefix
}