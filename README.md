# bosh CLI

* Documentation: [bosh.io/docs/cli-v2](https://bosh.io/docs/cli-v2.html)
* Slack: #bosh on <https://slack.cloudfoundry.org>
* Mailing list: [cf-bosh](https://lists.cloudfoundry.org/pipermail/cf-bosh)
* CI: <https://main.bosh-ci.cf-app.com/teams/main/pipelines/bosh:cli>
* Roadmap: [Pivotal Tracker](https://www.pivotaltracker.com/n/projects/956238)

## Special Note about This Fork

This fork allows bosh v2 CLI to skip TLS verification on communicating with a BOSH director.

### Why?

I can understand the CFF's decision to force certificate verification ([1](https://github.com/cloudfoundry/bosh-cli/issues/82), [2](https://github.com/cloudfoundry/bosh-cli/issues/168)), but feel too strict to conform it for some existing environments.

Thare are still so many BOSH environments built with bosh CLI v1 and/or bosh-init. To communicate with the BOSH director in those environments using bosh CLI v2, it may sometimes be difficult to provide the right certificates especially when the director is using a self-signed certificate.

So I've made this fork for tentative use until bosh CLI v2 will become the mainstream.

### How to Build

#### Prerequisites

* Go 1.6 or later

#### Build (only for Linux/bash environment)

```
git clone https://github.com/nota-ja/bosh-cli.git
cd bosh-cli
git remote add upstream https://github.com/cloudfoundry/bosh-cli.git  # For the sake of version of build
git fetch upstream -p                                                 # For the sake of version of build
bin/build-fork
```

And you'll get the `bosh` executable binary in the top of the repository.

```
./bosh -v
```

### How to Skip Certificate Verification

Just set the environment variable `BOSH_DIRECTOR_INSECURE_HTTPS` to `true`:

```
export BOSH_DIRECTOR_INSECURE_HTTPS=true
```

## Usage

- [Install](https://bosh.io/docs/cli-v2.html)

### Installing using a package manager

**Mac OS X** (using [Homebrew](http://brew.sh/) via the [cloudfoundry tap](https://github.com/cloudfoundry/homebrew-tap)):

```sh
$ brew install cloudfoundry/tap/bosh-cli
```

## Client Library

This project includes [`director`](director/interfaces.go) and [`uaa`](uaa/interfaces.go) packages meant to be used in your project for programmatic access to the Director API.

See [docs/example.go](docs/example.go) for a live short usage example.

## Developer Notes

- [Workstation setup docs](docs/build.md)
- [Test docs](docs/test.md)
- [CLI workflow](docs/cli_workflow.md)
  - [Architecture docs](docs/architecture.md)
