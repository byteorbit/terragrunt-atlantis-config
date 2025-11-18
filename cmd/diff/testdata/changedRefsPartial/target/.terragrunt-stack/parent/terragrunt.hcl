include "root" {
  path = find_in_parent_folders("root.hcl")
}

terraform {
  source = "${find_in_parent_folders("test_extras/catalog/modules/versioned")}/${values.version}"
}
