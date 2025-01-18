FROM google/cloud-sdk

COPY service_account_key.json /app/service_account_key.json

RUN gcloud auth activate-service-account --key-file=$GOOGLE_APPLICATION_CREDENTIALS
RUN gcloud config set project $GOOGLE_CLOUD_PROJECT
RUN gcloud config set compute/region $CLOUDSDK_REGION
