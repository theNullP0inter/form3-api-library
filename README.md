# Form3 API Library

golang library for using Form3 APIs ( only account service )

## Libraries used

* [google/uuid](https://github.com/google/uuid)

* [stretchr/testify](https://github.com/stretchr/testify)

* [jarcoal/httpmock](https://github.com/jarcoal/httpmock)


## Project setup

```
\_ src
    \_ account           # client for account service
    \_ common            # common clients and utils
    \_ e2e               # end to end tests
    \_ form3             # client to combine all services
    \_ models            # all models
\_ scripts
    \_ db                # init scripts for postgres
\_ .github
    \_ workflows         # CI workflows using github actions
    \_ ISSUE_TEMPLATE    # Templates to raise issues on github

```


## Testing the library 

Clone the repository: `git clone https://github.com/theNullP0inter/form3-api-library.git`


Unit tests can be run without docker,

1. Install the dependencies with `go get -v -t -d ./...`

2. Run the unit tests with `go test -v ./...`

To run e2e tests, use docker-compose

* This command will raise all the required containers and close them after the tests are done
`docker-compose up --build --exit-code-from form3-lib-tester`

* To disable the e2e tests, remove `RUN_E2E=True` from the environment in `docker-compose.yml`


## CI/CD setup

CI/CD is managed using github-actions

* `unit-test-on-pr.yml` runs all the unit tests when a new PR is raised to the master branch 

* `e2e-test-on-merge` runs e2e and unit tests when a PR is merged to the master branch 

* `build-on-release` builds a docker image and pushes it to `ghcr` when a new release is made on github with tags in semantic versioning. => [ This is not relevant for an API library]


## Getting Started with the library

```

# Initialize your client
form3_cli := form3.NewForm3Client(&common.Config{
    Live: os.Getenv("ENV") == "production", # API endpoint will depend on this
    HttpTimeout: 10 * time.Second, # Timeout for Http client
})

# To Create an account
acc, err := form3_cli.Account.Create(Mockaccount)

# To fetch the list of all the accounts
accs, err := form3_cli.Account.List(common.NewPagination(0, 10), map[string]string{})

# To fetch an account by ID
acc, err = form3_cli.Account.Fetch(accID)

# To delete the account
err = form3_cli.Account.Delete(accID, accVersion)


```