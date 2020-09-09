# CHANGELOG

Version log for Spekkio.

## 0.2.9
### Modified
- Fixed typos in controller test error logs

## 0.2.8
### Modified
- Replaced personal fork of `github.com/davyzhang-agw` with updated master copy (main repository now supports `httptest.ResponseRecorder` within `agw.WriteResponse`)

## 0.2.7
### Modified
- Created separate common functions for initializing GitHub GraphQL API request

## 0.2.6
### Modified
- Created separate common functions for Spekkio error messages

## 0.2.5
### New
- Added `controller_test.go` to `src/` README

## 0.2.4
### New
- Added build shell script that does not publish to AWS Lambda

## 0.2.3
### New
- Added unit testing for controller package
### Modified
- Use personal fork of `github.com/davyzhang-agw` to support unit testing of `agw.WriteResponse`
### Removed
- Removed unused `getSqlFrom()` common function

## 0.2.2
### Modified
- Fixed errors / warnings thrown by `gofmt -s` and `golint` on `v0.2.1`

## 0.2.1
### New
- Added SQL query builder
### Modified
- Deploy shell script no longer copies SQL queries directory to `bin/`

## 0.2.0
### New
- Added Scrumify route / controller
- Added ScrumifyRequest struct to models
- Added queries for Scrumify
- Added POST handler wrapper
- Added GraphQL query builder
- Added links to related repositories to README
- Added latest version badge to README
- Deploy shell script now copies queries to build directory
### Modified
- Enhanced Error response struct in models

## 0.1.2
### New
- Added middleware to log request info

## 0.1.1
### Modified
- Added comments to exported functions to suppress warnings from golint
- Unexported functions formerly in server package (now in main)
### Removed
- Removed server package and combined it into main

## 0.1.0
### New
- App can run on traditional server or AWS Lambda (configurable)
- App now has a 404 handler
- Shell scripts added to deploy and run the app
### Modified
- Additional levels of abstraction added to app directory structure
- HTTP handlers slightly modified to support both default and AWS Lambda platforms
### Removed
- Alice logging / middleware temporarily disabled

## 0.0.1
### New
- App now has an index route
- App logs request and response details
- App components further divided into controller and model

## 0.0.0
### New
- Initial app!
- App server / database interfaces and CORS origins can be configured
- App supports Postgres databases only
