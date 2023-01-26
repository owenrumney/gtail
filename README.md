# :page_with_curl: gtail

Tailing for GCP services to make reading logs easier.

> :warning: gtail is still in early development and you kind of use it at your own risk :grimacing:

- [gtail](#gtail)
  - [Important Notes](#important-notes)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Cloud Build Logs](#cloud-build-logs)
      - [Historical Cloud Build Logs](#historical-cloud-build-logs)
  - [Cloud Run Logs](#cloud-run-logs)
    - [Historical Cloud Run Logs](#historical-cloud-run-logs)
  - [Tail PubSub Topic](#tail-pubsub-topic)
  - [Template Strings](#template-strings)


## Important Notes

- :warning: gtail requires you to have active ADC creds to authenticate against GCP. This can be done by running `gcloud auth application-default login` if you have the gcloud CLI installed.
- :warning: gtail assumes you have the permissions to the project or service you are trying to tail logs for.
- :warning: gtail assumes you have the permissions to create subscriptions for the pubsub topic you are trying to tail.

## Installation

From source

```bash
go install github.com/owenrumney/gtail/cmd/gtail@latest
```

Alternatively, download the latest release from the [releases page](https://github.com/owenrumney/gtail/releases)

## Usage

You need to have authenticated against GCP before using gtail. This can be done by running `gcloud auth application-default login` if you have the gcloud CLI installed.

```bash
gcloud auth application-default login
```

`gtail` uses the Application Default Credentials to authenticate against GCP, so once you have done this, `gtail` can run.

### Cloud Build Logs

Tail the logs for a Cloud Build job or pull back previous logs from a give time ago

```bash
$ gtail cloud-build -h
```
```text
Tail logs for a CloudBuild Job

Usage:
  gtail cloud-build [flags]
  gtail cloud-build [command]

Aliases:
  cloud-build, cb

Available Commands:
  historic    Get the cloud build logs for a trigger that has already completed

Flags:
      --build-id string       The cloud build ID
  -h, --help                  help for cloud-build
  -o, --output string         The output format either json or a template string
      --severity strings      The severity of logs to include
      --trigger-name string   The name of the cloud build trigger to use

Global Flags:
  -d, --debug            Enable debug logging
  -p, --project string   The GCP project ID
  -r, --region string    The GCP region (default "us-central1")

Use "gtail cloud-build [command] --help" for more information about a command.
```

An example usage to stream the logs for a trigger called `my-trigger` in project `my-project`:

```bash
$ gtail cloud-build --trigger-name my-trigger --project my-project
```

This will stream any logs that come through for that build trigger from this point onwards.


#### Historical Cloud Build Logs

If you want to pull back logs for a build that has already completed, you can use the `historic` subcommand.

```bash
$ gtail cloud-build historic -h
```
```text
Get the cloud build logs for a trigger that has already completed

Usage:
  gtail cloud-build historic [flags]

Flags:
  -h, --help            help for historic
      --hours-ago int   Roughly how many hours ago the build happened. Searches a window of time from then till now (default 24)

Global Flags:
      --build-id string       The cloud build ID
  -d, --debug                 Enable debug logging
  -o, --output string         The output format either json or a template string
  -p, --project string        The GCP project ID
  -r, --region string         The GCP region (default "us-central1")
      --severity strings      The severity of logs to include
      --trigger-name string   The name of the cloud build trigger to use
```


## Cloud Run Logs

Tail the logs for a Cloud Run service

```bash
$ gtail cloud-run -h
```
```text
Tail logs for a Cloud Run revision

Usage:
  gtail cloud-run [flags]
  gtail cloud-run [command]

Aliases:
  cloud-run, cr

Available Commands:
  historic    Get the cloud run logs for a revision that has already exited

Flags:
  -h, --help                 help for cloud-run
  -o, --output string        The output format either json or a template string
      --revision-id string   The cloud run revision ID
      --service string       Cloud Run service for the logs to get
      --severity strings     The severity of logs to include

Global Flags:
  -d, --debug            Enable debug logging
  -p, --project string   The GCP project ID
  -r, --region string    The GCP region (default "us-central1")

Use "gtail cloud-run [command] --help" for more information about a command.
```

### Historical Cloud Run Logs

You can get the Cloud Run logs for a revision that has already exited by using the `historic` subcommand.

```bash
$ gtail cloud-run historic -h
```
```text
Get the cloud run logs for a revision that has already exited

Usage:
  gtail cloud-run historic [flags]

Flags:
  -h, --help            help for historic
      --hours-ago int   Roughly how many hours ago the log happened. Searches a window of time from then till now (default 24)

Global Flags:
  -d, --debug                Enable debug logging
  -o, --output string        The output format either json or a template string
  -p, --project string       The GCP project ID
  -r, --region string        The GCP region (default "us-central1")
      --revision-id string   The cloud run revision ID
      --service string       Cloud Run service for the logs to get
      --severity strings     The severity of logs to include
```


## Tail PubSub Topic

gtail will create a temporary subscription for a topic. 

> :warning: This is a temporary subscription and will be deleted when gtail exits and you need permissions to create subscriptions

```bash
$ gtail pubsub
```
```text
You must specify a project ID using the -p flag or GCP_PROJECT_ID envvar.

Usage:
  gtail pubsub [flags]

Aliases:
  pubsub, ps

Flags:
  -h, --help                   help for pubsub
  -o, --output string          The output format json or a template string
      --tail-duration string   The duration to tail for (default "10m")
  -t, --topic string           The pubsub topic to tail

Global Flags:
  -d, --debug            Enable debug logging
  -p, --project string   The GCP project ID
  -r, --region string    The GCP region (default "us-central1")
```

## Template Strings

You can use a template string to format the output of gtail. Doing so assumes an understanding of the structure of the messages that are being returned and a knowledge of the [Go text/template package](https://golang.org/pkg/text/template/).

An example might be 

```bash
gtail cloud-build historic -p my-project --trigger-name my-trigger-build --output '{{ .Severity}} - {{ index .Resource.Labels "build_id"}} {{ .Payload }}'
```

Which would give you something fairly unhelpful like below, where every line is the same;

```log
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a Something that happened
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a Something else that happened
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a Something you might need to know about
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a Something else
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a A few empty lines then another thing that's interesting
```

