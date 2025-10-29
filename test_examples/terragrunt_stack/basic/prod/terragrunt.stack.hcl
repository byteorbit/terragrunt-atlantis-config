unit "service" {
  source = "${find_in_parent_folders("catalog/units")}//service"
  path = "service"

  values = {
    prefix = "testservice"
  }
}
