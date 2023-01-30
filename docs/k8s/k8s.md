---
layout: default
title: GKE Cluster
nav_order: 50
has_children: true
permalink: /k8s
---

# GKE Cluster
{: .no_toc }

gtail allows you to tail logs from GKE Clusters.

## K8s Command
```bash
gtail k8s -h
```
```text
Tail logs for GKE Cluster

Usage:
  gtail k8s [flags]

Flags:
      --cluster string     GKE Cluster you want to tail logs for
  -h, --help               help for k8s
  -o, --output string      The output format either json or a template string
      --severity strings   The severity of logs to include

Global Flags:
  -d, --debug            Enable debug logging
  -p, --project string   The GCP project ID
  -r, --region string    The GCP region (default "us-central1")
  ```