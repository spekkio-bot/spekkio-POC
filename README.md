# Spekkio
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/spekkio-bot/spekkio?label=latest)](https://github.com/spekkio-bot/spekkio/tags)
[![Server status](https://img.shields.io/website?down_color=red&down_message=offline&label=server&up_message=online&url=https%3A%2F%2Fjunha.netlify.com)](https://5ila6fw37k.execute-api.us-west-1.amazonaws.com/api)
[![Go Report Card](https://goreportcard.com/badge/github.com/spekkio-bot/spekkio)](https://goreportcard.com/report/github.com/spekkio-bot/spekkio)

*"I'm Spekkio, the Master of ~War~ GitHub!"*

Spekkio is a set of tools and services that delivers a better developer experience to teams on GitHub by automating tedious things.

## Related Repositories

Listed below are related repositories for Spekkio:
- [Command Line Interface](https://github.com/spekkio-bot/spekkio-cli)
- [Database Schema](https://github.com/spekkio-bot/spekkio-dbschema)

## Available Scripts

### `./run.sh`

To run the app, simply run `./run.sh`. You will need to do some [setup](https://github.com/spekkio-bot/spekkio/blob/master/src/README.md#first-time-setup) before you can run the app.

You can also navigate to `src/` and directly build / run the app there.

### `./run.sh t`, `./run.sh test`

This will run all unit tests, print your test results on the console, and open a test coverage report on your browser.

### `./run.sh d`, `./run.sh deploy`

You can deploy the app to AWS Lambda with `./deploy.sh` - in a nutshell, this script will first build your app to `bin/`, then deploy your executable app to AWS if it built successfully.

Deploy settings are determined by `serverless.yml`.

### `./run.sh b`, `./run.sh build`

If you wish, you may choose to build a Linux/AMD64 binary without deploying to test if the app builds properly.

## License

Licensed by the GNU General Public License, version 3.
