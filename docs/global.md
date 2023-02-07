---
layout: default
title: Global Arguments
nav_order: 4
---

# Global Arguments
{: .no_toc }

There are a number of global arguments that can be used with any command.

## `--project` or `-p`
This is the GCP project ID to use. If not specified, the project ID will be inferred from the current GCP project.

Alternatively, it can be set using the `GCP_CLOUD_PROJECT` environment variable.


## `--region` or `-r`
This is the GCP region to use. If not specified, the region will be inferred from the current GCP project. 
This is defaulted to `us-central1`.

## `--debug` or `-d`

This enables debug logging. This is useful for debugging issues with gtail as it will give you information about what log queries are being used.
