unit "service" {
  # source = "service"
  source = "${find_in_parent_folders("test_extras/catalog/units")}//versioned-module"
  path = "service"

  values = {
    version = "v2"
  }
}
