This directory exists purely to exclude test assets from the `test_examples` directory. This is required to prevent false test failures when the `test_examples`
directory is scraped for `terragrunt.hcl` files in `TestEnvHCLProjectsExternalChilds` and
`TestEnvHCLProjectsAllChilds`.