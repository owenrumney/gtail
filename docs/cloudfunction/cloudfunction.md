---
layout: default
title: Cloud Function
nav_order: 30
has_children: true
permalink: /cloudfunction
---

# Cloud Functions
{: .no_toc }

gtail allows you to tail logs from Google Cloud Functions both as streams or historic logs.

## Cloud Functions Command


```bash
gtail cloud-funtion -h
```
```text
Tail logs for a Cloud Function revision

Usage:
  gtail cloud-function [flags]
  gtail cloud-function [command]

Aliases:
  cloud-function, cf

Available Commands:
  historic    Get the cloud run logs for a revision that has already exited

Flags:
      --function string    Cloud Function name for the logs to get
  -h, --help               help for cloud-function
  -o, --output string      The output format either json or a template string
      --severity strings   The severity of logs to include

Global Flags:
  -d, --debug            Enable debug logging
  -p, --project string   The GCP project ID
  -r, --region string    The GCP region (default "us-central1")

Use "gtail cloud-function [command] --help" for more information about a command.
```