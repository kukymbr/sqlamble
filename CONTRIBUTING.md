# Contributing

You are very welcome to contribute to the `sqlamble`.

## Committing generated code

If you are contributing to the project and make changes to the code generator, please commit the
generated code in example as well. This is to make sure that the generated code is always up-to-date.

To update generated code run:

```shell
make generate_example
```

Generated code should be committed in a one separate commit `chore: commit generated files`.

```shell
git add ./example
git commit -m "chore: commit generated files"
```