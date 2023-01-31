---
layout: default
title: Cloud Run
nav_order: 40
has_children: true
permalink: /cloudrun
parent: Supported Services
---

# Cloud Run
{: .no_toc }

gtail allows you to tail logs from Google Cloud Run both as streams or historic logs.

# Streaming Cloud Run
{: .no_toc }

{: .no_toc }

gtail can start streaming a run based on the `revision-id` or the `service` name.

If you're tailing the logs as a stream, it's best to use the `service` flag.

```bash
gtail cloud-run -h
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

Passing the `--revision-id` flag will start streaming the logs for that revision or you can use `--service` to specify a service and start streaming the latest from that.

# Historic Cloud Run
{: .no_toc }

gtail can get historic logs for a Cloud Run that has already completed

```bash
gtail cloud-run historic -h
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

You need to pass a `--revision-id` or `--service` flag to get the logs for a specific revision or service.

The `--hours-ago` flag will search for a build that started within the last `n` hours. If you don't specify this flag it will search for a build that started within the last 24 hours.

| Note: if you do a historic build with the `--service` flag it will tail everything in the period and you may hit the rate limit if it is a particularly busy service