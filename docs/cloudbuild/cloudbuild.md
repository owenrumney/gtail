---
layout: default
title: Cloud Build
nav_order: 20
has_children: true
permalink: /cloudbuild
---

# Cloud Build
{: .no_toc }

gtail allows you to tail logs from Google Cloud Build both as streams or historic logs.

## Cloud Build Command

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
