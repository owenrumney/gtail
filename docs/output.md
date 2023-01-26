---
layout: default
title: Output
nav_order: 5
---

# Output

Most commands in gtail support outputting the logs in a number of different formats. This is done using the `--output` flag.

## json 

If you use `--output json` you will get your results in json format. This is useful for seeing what is actually in the message.

## Template Strings

You can use a template string to format the output of gtail. Doing so assumes an understanding of the structure of the messages that are being returned and a knowledge of the [Go text/template package](https://golang.org/pkg/text/template/).

An example might be 

```bash
gtail cloud-build historic -p my-project --trigger-name my-trigger-build --output '{{ .Severity}} - {{ index .Resource.Labels "build_id"}} {{ .Payload }}'
```

Which would give you something fairly unhelpful like below, where every line is the same;

```log
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a Something that happened
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a Something else that happened
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a Something you might need to know about
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a Something else
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a
INFO - 7f3b0b1a-1b8a-4b0f-8b1a-1b8a4b0f8b1a A few empty lines then another thing that's interesting
```

