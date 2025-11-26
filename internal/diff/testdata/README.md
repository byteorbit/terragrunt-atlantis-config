atlantis.yaml generated for projects using:

```shell
terragrunt-atlantis-config generate --output atlantis.yaml --autoplan --parallel  --execution-order-groups --create-project-name --depends-on
```
> NOTE ON THE TESTDATA
> To ensure the diff commands unit tests are not coupled to terragrunt and the generate command, testdata is pre rendered.
> Generating stacks and creating atlantis files from scratch is out of scope of this command.
> This ensures very fast unit tests. Normally these assets would be inlined in the go test files, this reduces legibility of the go files.

The `cmd/diff/scripts/reset_test_data.sh` script should be executed if any project changes are made, to ensure the tests
accurately reflect changes.

## Test scenarios:

1. `emptyBase`: Base project is empty and Target has modules generated from a Stack.
2. `unchangedFiles`: Assert no changes results in autoplan false.
3. `changedRefsSelf`: Reference for this unit changed (eg. version change).
4. `changedRefsDependent`: Reference of a dependent changed (eg. version of dependent module changed, both modules
   should change then) without the unit itself changing versions.
5. `changedRefsDeep`: Reference of an indirect dependent changed (eg. version of dependents dependent module changed,
   all the modules in the chain should change then) without the final unit itself changing versions.
6. `changedFiles`: Same list of files in the `when_modified` though with changes files matched by the set. For example
   an additional tf file in the unit directory.
7. `changedFilesSet`: List of files in the `when_modified` changed. For example another include was added.