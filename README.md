# get_access_token - an Access Token (JWT) Fetcher

![Build status](https://github.com/tztz/get_access_token/actions/workflows/build.yml/badge.svg)

The CLI tool `get_access_token` fetches and prints an access token (JWT) for a given environment/stage like e.g. `int`.
The URLs and credentials for the different environments/stages are configured via environment variables or an `.env` file.

## Build

### Build prerequisites

Install all needed dependencies via

```bash
scripts/install_deps.sh
```

Now, tools like `staticcheck` and `gosec` are installed and ready to be used during the build.

### Build the application

From this project's root directory execute

```bash
scripts/build.sh
```

This builds the application and puts the executable in the `bin` folder of this project's root directory.

The `bin` folder is ignored and not committed to Git.

## Run

### Run prerequisites

Make sure you have set the following environment variables (the example values must be replaced, of course):

```bash
RD_OIDC_TOKEN_UrlInt=https://your-url-to-int/protocol/openid-connect/token
RD_OIDC_TOKEN_UrlPre=https://your-url-to-pre/protocol/openid-connect/token
RD_OIDC_TOKEN_UrlProd=https://your-url-to-prod/protocol/openid-connect/token

RD_OIDC_TOKEN_BasicAuthInt=your-basic-auth-secret-for-int
RD_OIDC_TOKEN_BasicAuthPre=your-basic-auth-secret-for-pre
RD_OIDC_TOKEN_BasicAuthProd=your-basic-auth-secret-for-prod
```

Alternatively, you can have an `.env` file in this project's root directory containing the above variables.

The `.env` file is ignored and not committed to Git.

### Run the application

From this project's root directory execute the application and pass the desired environment/stage as first argument, e.g.:

```bash
./bin/get_access_token int
```

This prints an access token (JWT) for the `int` environment/stage.

In order to get a verbose output pass `-v` as second argument:

```bash
./bin/get_access_token int -v
```
