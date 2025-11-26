unit "ancestor" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//versioned-module"
  path   = "ancestor"

  values = {
    version = "v1"
  }
}

unit "parent" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//with-dependency"
  path   = "parent"

  values = {
    parentPath = "../ancestor"
  }
}

unit "child" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//with-dependency"
  path   = "child"

  values = {
    parentPath = "../parent"
  }
}
