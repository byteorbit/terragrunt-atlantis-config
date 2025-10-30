include "root" {
  path = find_in_parent_folders("root.hcl")
}

dependency "dependency" {
  config_path = values.parentPath
}
