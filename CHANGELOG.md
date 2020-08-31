# CHANGELOG

Version log for Spekkio.

## 0.2.0
### New
- Added Scrumify route / controller
- Added ScrumifyRequest struct to models
- Added queries for Scrumify
- Added POST handler wrapper
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
