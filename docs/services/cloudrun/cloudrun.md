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

## Cloud Run Command


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