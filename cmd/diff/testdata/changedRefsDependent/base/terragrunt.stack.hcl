unit "parent" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//versioned-module"
  path   = "parent"

  values = {
    version = "v1"
  }
}

unit "child" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//with-dependency"
  path   = "child"

  values = {
    parentPath = "../parent"
  }
}
