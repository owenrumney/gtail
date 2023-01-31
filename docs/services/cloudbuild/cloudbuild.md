---
layout: default
title: Cloud Build
nav_order: 20
has_children: true
permalink: /cloudbuild
parent: Supported Services
---

# Cloud Build
{: .no_toc }

gtail allows you to tail logs from Google Cloud Build both as streams or historic logs.

## Streaming Cloud Build

gtail can start streaming a build based on the ID or the trigger name.

```bash
gtail cloud-build -h
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

Passing the `--build-id` flag will start streaming the logs for that build or you can use `--trigger-name` to specify a trigger and start streaming that.

## Historic Cloud Build
{: .no_toc }

gtail can get historic logs for a Cloud Build that has already completed

```bash
gtail cloud-build historic -h
```
```text
Get the Cloud Build logs for a trigger that has already completed

Usage:
  gtail cloud-build historic [flags]

Flags:
  -h, --help            help for historic
      --hours-ago int   Roughly how many hours ago the build happened. Searches a window of time from then till now (default 24)
      --last-run        Get the logs for the last run of the trigger

Global Flags:
      --build-id string       The cloud build ID
  -d, --debug                 Enable debug logging
  -o, --output string         The output format either json or a template string
  -p, --project string        The GCP project ID (default "gs-app-iac")
  -r, --region string         The GCP region (default "us-central1")
      --severity strings      The severity of logs to include
      --trigger-name string   The name of the cloud build trigger to use
```

The `--hours-ago` flag will search for a build that started within the last `n` hours. If you don't specify this flag it will search for a build that started within the last 24 hours.

| Note: starting a historic build with the `--trigger-name` flag will tail all builds for that trigger in the given time period.

The `--last-run` flag will get the logs for the last run of the trigger, it needs to be used with the `--trigger-name` flag. The time it ran will be pulled from the build details so `--hours-ago` will be ignored.