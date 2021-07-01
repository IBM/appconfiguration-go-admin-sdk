# Contributing

Review the following guidelines for submitting questions, issues, or changes to this repository.

## Questions

If you have questions about using the IBM Cloud App Configuration SDKs, you can ask questions in
[Stack Overflow](https://stackoverflow.com/questions/tagged/ibm-appconfiguration). Be sure to include
the `ibm-appconfiguration` tags.

You can also check out the [App Configuration documentation](https://cloud.ibm.com/docs/app-configuration) and [API reference](https://cloud.ibm.com/apidocs/app-configuration?code=go) for more information about the service.

## Coding Style

The SDK follows the Go coding conventions documented [here](https://golang.org/doc/effective_go.html). 

### Running Linter
Linter is a tool that analyzes source code to flag programming errors, bugs, stylistic errors, and suspicious constructs.

Perform the following steps to run lint check on the SDK :-
- Ensure that local GOPATH and GOBIN are fully setup by using the go env command.
- Install the package by using the command 
    ```
    go get github.com/golangci/golangci-lint/cmd/golangci-lint
    ```
- To run lint on the SDK, run the following from the SDK root directory.
    ```
    make lint
    ```

## Issues

If you encounter an issue with Go Admin SDK, you're welcome to submit
a [bug report](https://github.com/IBM/appconfiguration-go-admin-sdk/issues). Before that, please search for similar issues. It's possible somebody has encountered this issue already.

## Pull Requests

If you want to contribute to the repository, follow these steps:

1. Fork the repository
2. Develop and test your code changes:
    - To build/test: `make runUnitTests`
3. Please add one or more tests to validate your changes.
4. Make sure everything builds/tests cleanly
5. Check your code for lint issues:
    ```
    make lint
    ```
6. Commit your changes
7. Push to your fork and submit a pull request to the `master` branch

## Running the tests

The tests within the SDK consists of unit tests.

To run the test, run the following from the root folder:

```
make runUnitTests
```

# Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
   have the right to submit it under the open source license
   indicated in the file; or

(b) The contribution is based upon previous work that, to the best
   of my knowledge, is covered under an appropriate open source
   license and I have the right under that license to submit that
   work with modifications, whether created in whole or in part
   by me, under the same open source license (unless I am
   permitted to submit under a different license), as indicated
   in the file; or

(c) The contribution was provided directly to me by some other
   person who certified (a), (b) or (c) and I have not modified
   it.

(d) I understand and agree that this project and the contribution
   are public and that a record of the contribution (including all
   personal information I submit with it, including my sign-off) is
   maintained indefinitely and may be redistributed consistent with
   this project or the open source license(s) involved.

