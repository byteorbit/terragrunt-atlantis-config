unit "parent" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//service"
  path   = "parent"

  values = {
    prefix = "testparent"
  }
}

unit "child" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//with-dependency"
  path   = "child"

  values = {
    prefix     = "testchild"
    parentPath = "../parent"
  }
}
