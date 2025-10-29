include "root" {
  path   = find_in_parent_folders("root.hcl")
  expose = true
}

include "common_configs" {
  path   = format("%s/common/terragrunt.hcl", dirname(find_in_parent_folders("root.hcl")))
  expose = true
}

terraform {
  source = "git::git@github.com:transcend-io/terraform-aws-fargate-container?ref=v0.0.4"
}

inputs = {
  foo = "bar"
}
