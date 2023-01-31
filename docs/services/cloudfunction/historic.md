---
layout: default
title: Historic Cloud Function
parent: Cloud Function
nav_order: 5
---

# Historic Cloud Function
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