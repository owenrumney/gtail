---
layout: default
title: Streaming Cloud Build
parent: Cloud Build
nav_order: 5
---

# Streaming Cloud Build
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

Passing the `--build-id` flag will start streaming the logs for that build or you can use `--trigger-name` to specify a trigger and start streaming that.
