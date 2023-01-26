---
layout: default
title: Historic Cloud Run
parent: Cloud Run
nav_order: 5
---

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