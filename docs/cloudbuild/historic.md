---
layout: default
title: Historic Cloud Build
parent: Cloud Build
nav_order: 5
---

# Historic Cloud Build
{: .no_toc }

gtail can get historic logs for a Cloud Build that has already completed

```bash
gtail cloud-build historic -h
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

The `--hours-ago` flag will search for a build that started within the last `n` hours. If you don't specify this flag it will search for a build that started within the last 24 hours.

| Note: starting a historic build with the `--trigger-name` flag will tail all builds for that trigger in the given time period.