---
layout: default
title: Streaming Cloud Function
parent: Cloud Function
nav_order: 5
---

# Streaming Cloud Function
{: .no_toc }


gtail can start streaming a function run based on the `execution-id` or the `function` name.

If you're tailing the logs as a stream, it's best to use the `function` flag.

```bash
gtail cloud-run -h
```
```text
Tail logs for a Cloud Function revision

Usage:
  gtail cloud-function [flags]
  gtail cloud-function [command]

Aliases:
  cloud-function, cf

Available Commands:
  historic    Get the cloud function logs for a run that has already exited

Flags:
      --execution-id string   The Cloud Function execution ID
      --function string       Cloud Function name for the logs to get
  -h, --help                  help for cloud-function
  -o, --output string         The output format either json or a template string
      --severity strings      The severity of logs to include

Global Flags:
  -d, --debug            Enable debug logging
  -p, --project string   The GCP project ID
  -r, --region string    The GCP region (default "us-central1")

Use "gtail cloud-function [command] --help" for more information about a command.
```

Passing the `--execution-id` flag will start streaming the logs for that execution or you can use `--function` to specify a service and start streaming the latest from that.

