# godcr

[![Build Status](https://github.com/raedahgroup/godcr/workflows/Build/badge.svg)](https://github.com/raedahgroup/godcr/actions)
[![Tests Status](https://github.com/raedahgroup/godcr/workflows/Tests/badge.svg)](https://github.com/raedahgroup/godcr/actions)

A cross-platform desktop [SPV](https://docs.decred.org/wallets/spv/) wallet for [decred](https://decred.org/) built with [gio](https://gioui.org/).

## Building

Note: You need to have [Go 1.13](https://golang.org/dl/) or above to build.

Follow the [installation](https://gioui.org/doc/install) instructions for gio.

Pkger is required to bundle static resources before building. Install Pkger globally by running
`go get github.com/markbates/pkger/cmd/pkger`

In the root directory, run
`pkger`. A pkged.go file should be silently created in your root directory.

Then `go build`.

## Contributing

See [CONTRIBUTING.md](https://github.com/raedahgroup/godcr/blob/master/.github/CONTRIBUTING.md)
