![Dynamic XML Badge](https://img.shields.io/badge/dynamic/xml?url=https%3A%2F%2Fspetix.github.io%2Fbar-out-adapters%2Fcoverage.xml&query=round(100*%2F%2Fcoverage%2F%40line-rate)&suffix=%25&style=plastic&label=Code%20Coverage)


Library to output your blocket to different bars (waybar, i3bar...)
## Installation

TBD

# Build details

* [Coverage](https://bar-out-adapters.cassano.fe.it/coverage/index.html)

## Make targets

- `make clean` — remove build artifacts, coverage, and generated files
- `make build` — build release binaries
- `make test` — run tests and generate `test-results.json`
- `make coverage` — run tests and create `coverage.txt` and `coverage.xml`
- `make project-stats` — generate `docs-site/static/project-stats.json`
- `make local-site` — build the docs site locally with coverage included
- `make run-local-site` — build the local site and serve it locally
- `make run` — start the docs dev server locally

## License

MIT

## Author

[spetix](https://github.com/spetix)

