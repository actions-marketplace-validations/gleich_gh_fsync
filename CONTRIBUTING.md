# Contributing

๐ Welcome to gh_fsync! Thank you for showing interest in contributing to gh_fsync, we would love to have your contribution. Below are some details on how to contribute to gh_fsync. Please read carefully!

- [Contributing](#contributing)
  - [๐ Requesting Features/Reporting Bugs](#-requesting-featuresreporting-bugs)
  - [โ Adding/Changing code](#-addingchanging-code)
    - [โ ๏ธ Notice](#๏ธ-notice)
    - [๐งพ Process](#-process)
    - [โ Checking Code](#-checking-code)
  - [โน๏ธ General](#โน๏ธ-general)

## ๐ Requesting Features/Reporting Bugs

1. Click on the "Issues" tab in the repo.
2. Make sure that the issue doesn't exist already by searching for it.
3. Pick the issue template.
4. Fill in the issue template.

## โ Adding/Changing code

### โ ๏ธ Notice

This project uses [golangci-lint](https://github.com/golangci/golangci-lint) for code linting, please install it and format your code with `make lint-golangci`.

### ๐งพ Process

1. Make an issue (see above) and check to make sure what you are trying to add/change doesn't already exist.
2. Create a branch with the name being the issue number. If you don't have contributor access just fork the repo.
3. Make and commit the changes.
4. [Validate code](#-checking-code)
5. Make the pull request!
6. Now someone on the team will review your PR. Congrats!
7. **Warning** once your PR gets merged the branch for it will automatically get deleted (only for contributors with contributor access).

### โ Checking Code

If both the options stated below don't work for you can just make the changes if the CI jobs fail.

#### ๐ณ Docker Container

You can check all the code inside of a docker container with all the dependencies installed by running `make docker-lint` and `make test-in-docker`. Both of these commands will build the image for you and run it. No need to install anything! Check the output to make sure you don't have any issues to resolve.

#### โ๏ธ Manually

| **Program**                                                | **Description**                    |
| ---------------------------------------------------------- | ---------------------------------- |
| [golangci-lint](https://github.com/golangci/golangci-lint) | Linter for all golang files        |
| [goreleaser](https://github.com/goreleaser/goreleaser)     | Release automation for the program |
| [hadolint](https://github.com/hadolint/hadolint)           | Linter for all Dockerfiles         |

Once you have the programs above installed please run `make local-test` and `make local-lint`. If you don't get any errors your code checks out!

## โน๏ธ General

- When you take on an issue please set yourself as the assignee or leave a comment. This will let the maintainers and other contributors that know that you are going to work on it.
- **This project syncs files from other repos**. To check where certain files come from check the `fsync.yml` file (sometimes located in `/github/fsync.yml`). This means if you wanna change that file you need to change the file from the source repo.
