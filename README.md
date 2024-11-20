# Access Token (JWT) Fetcher

![Build status](https://github.com/tztz/get_access_token/actions/workflows/build.yml/badge.svg)

Fetches and prints an access token (JWT) for a given environment.

## Build

### Build prerequisites

Install all needed dependencies via

```bash
scripts/install_deps.sh
```

Now, tools like `staticcheck` and `gosec` are installed and ready to be used during the build.

### Build the application

From this folder execute

```bash
scripts/build.sh
```

This builds the application and puts the executable in the `bin` folder.
The `bin` folder is ignored and not committed to Git.

## Run

### Run prerequisites

Make sure to have an `.env` file in this folder containing the following content (the example values must be replaced, of course):

```bash
RD_OIDC_TOKEN_UrlInt=https://your-url-to-int/protocol/openid-connect/token
RD_OIDC_TOKEN_UrlPre=https://your-url-to-pre/protocol/openid-connect/token
RD_OIDC_TOKEN_UrlProd=https://your-url-to-prod/protocol/openid-connect/token

RD_OIDC_TOKEN_BasicAuthInt=your-basic-auth-secret-for-int
RD_OIDC_TOKEN_BasicAuthPre=your-basic-auth-secret-for-pre
RD_OIDC_TOKEN_BasicAuthProd=your-basic-auth-secret-for-prod
```

The `.env` file is ignored and not committed to Git.

### Run the application

From this folder execute the application and pass the desired environment as first argument, e.g.:

```bash
./bin/get_access_token int
```

In order to get a verbose output pass `-v` as second argument:

```bash
./bin/get_access_token int -v
```
