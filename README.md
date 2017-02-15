## Production Release

Push to Google Container Registry:

```sh
gcloud container builds submit --tag us.gcr.io/invenio-ml/healthz:<version> .
```