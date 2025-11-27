unit "ancestor" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//versioned-module"
  path   = "ancestor"

  values = {
    version    = "v1"
    prefix = "testservice"
  }
}

unit "parent1" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//with-dependency"
  path   = "parent1"

  values = {
    version    = "v1"
    parentPath = "../ancestor"
  }
}

unit "parent2" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//with-dependency"
  path   = "parent2"

  values = {
    parentPath = "../ancestor"
  }
}

unit "child" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//with-dependency"
  path   = "child"

  values = {
    parentPath = "../parent1"
  }
}
