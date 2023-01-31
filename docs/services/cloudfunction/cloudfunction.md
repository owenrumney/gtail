---
layout: default
title: Cloud Function
nav_order: 30
has_children: true
permalink: /cloudfunction
parent: Supported Services
---

# Cloud Functions
{: .no_toc }

gtail allows you to tail logs from Google Cloud Functions both as streams or historic logs.

# Streaming Cloud Function
{: .no_toc }


gtail can start streaming a function run based on the `execution-id` or the `function` name.

If you're tailing the logs as a stream, it's best to use the `function` flag.

```bash
gtail cloud-run -h
```
```text
Tail logs for a Cloud Function revision

Usage:
  gtail cloud-function [flags]
  gtail cloud-function [command]

Aliases:
  cloud-function, cf

Available Commands:
  historic    Get the cloud function logs for a run that has already exited

Flags:
      --execution-id string   The Cloud Function execution ID
      --function string       Cloud Function name for the logs to get
  -h, --help                  help for cloud-function
  -o, --output string         The output format either json or a template string
      --severity strings      The severity of logs to include

Global Flags:
  -d, --debug            Enable debug logging
  -p, --project string   The GCP project ID
  -r, --region string    The GCP region (default "us-central1")

Use "gtail cloud-function [command] --help" for more information about a command.
```

Passing the `--execution-id` flag will start streaming the logs for that execution or you can use `--function` to specify a service and start streaming the latest from that.

## Historic Cloud Function
{: .no_toc }

gtail can get historic logs for a Cloud Function that has already completed

```bash
Get the cloud function logs for a run that has already exited

Usage:
  gtail cloud-function historic [flags]

Flags:
  -h, --help            help for historic
      --hours-ago int   Roughly how many hours ago the log happened. Searches a window of time from then till now (default 24)

Global Flags:
  -d, --debug              Enable debug logging
      --function string    Cloud Function name for the logs to get
  -o, --output string      The output format either json or a template string
  -p, --project string     The GCP project ID
  -r, --region string      The GCP region (default "us-central1")
      --severity strings   The severity of logs to include
```

You need to pass a `--function` flag to get the logs for a specific function or --

The `--hours-ago` flag will search for a build that started within the last `n` hours. If you don't specify this flag it will search for a build that started within the last 24 hours.

| Note: if you do a historic build with the `--service` flag it will tail everything in the period and you may hit the rate limit if it is a particularly busy service