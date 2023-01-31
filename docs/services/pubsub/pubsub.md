---
layout: default
title: PubSub
nav_order: 60
has_children: true
permalink: /pubsub
parent: Supported Services
---

# PubSub Topic

gtail will create a temporary subscription for a topic. 

> :warning: This is a temporary subscription and will be deleted when gtail exits and you need permissions to create subscriptions

```bash
$ gtail pubsub
```
```text
You must specify a project ID using the -p flag or GCP_PROJECT_ID envvar.

Usage:
  gtail pubsub [flags]

Aliases:
  pubsub, ps

Flags:
  -h, --help                   help for pubsub
  -o, --output string          The output format json or a template string
      --tail-duration string   The duration to tail for (default "10m")
  -t, --topic string           The pubsub topic to tail

Global Flags:
  -d, --debug            Enable debug logging
  -p, --project string   The GCP project ID
  -r, --region string    The GCP region (default "us-central1")
```

The subscription that will be created by gtail will be named `gtail-<topic>-<number>` and will be deleted when gtail exits. The `number` is just the Unix ticks.

If you `Ctrl-C` gtail it will delete the subscription. If you `kill -9` gtail it will not delete the subscription and you will need to clean up.

## Tail PubSub Topic for a Duration

```bash
$ gtail pubsub -p my-project -t my-topic --tail-duration 1h
```

When passing a `--tail-duration` gtail will create a subscription and then delete it when the duration has passed.
You can use sensible durations like `1h` or `10m` or `1h30m` or `1h30m10s`.