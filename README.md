# Access Token (JWT) Fetcher

Fetches and prints an access token (JWT) for a given environment.

## Build

From this folder execute

```bash
scripts/build.sh
```

## Run

### Prerequisites

Make sure to have an `.env` file in this folder containing the following content (the values must be replaced, of course):

```bash
UrlInt=https://your-url-to-int/protocol/openid-connect/token
UrlPre=https://your-url-to-pre/protocol/openid-connect/token
UrlProd=https://your-url-to-prod/protocol/openid-connect/token

BasicAuthInt="your-basic-auth-secret-for-int"
BasicAuthPre="your-basic-auth-secret-for-pre"
BasicAuthProd="your-basic-auth-secret-for-prod"
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
