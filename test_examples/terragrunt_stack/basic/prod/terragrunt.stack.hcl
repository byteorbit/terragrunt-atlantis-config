unit "service" {
  source = "${find_in_parent_folders("test_extras/catalog/units")}//service"
  path = "service"

  values = {
    prefix = "testservice"
  }
}
