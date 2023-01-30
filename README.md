# :page_with_curl: gtail

Tailing for GCP services to make reading logs easier.

## Important Notes

- :warning: gtail requires you to have active ADC creds to authenticate against GCP. This can be done by running `gcloud auth application-default login` if you have the gcloud CLI installed.
- :warning: gtail assumes you have the permissions to the project or service you are trying to tail logs for.
- :warning: gtail assumes you have the permissions to create subscriptions for the pubsub topic you are trying to tail.

## Installation

From source

```bash
go install github.com/owenrumney/gtail/cmd/gtail@latest
```

Alternatively, download the latest release from the [releases page](https://github.com/owenrumney/gtail/releases)

## Usage

You need to have authenticated against GCP before using gtail. This can be done by running `gcloud auth application-default login` if you have the gcloud CLI installed.

```bash
gcloud auth application-default login
```

`gtail` uses the Application Default Credentials to authenticate against GCP, so once you have done this, `gtail` can run.


## Supported Services

- [Cloud Build](https://gtail.app/cloudbuild)
- [Cloud Functions](https://gtail.app/cloudfunction)
- [Cloud Run](https://gtail.app/cloudrun)
- [K8s](https://gtail.app/k8s)
- [Pub/Sub](https://gtail.app/pubsub)
