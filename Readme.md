# Upload Service Go Server


## Docker 

### Build

Build the docker image using the following command

``docker build . -t upload_service:latest --build-arg GOOGLE_APPLICATION_CREDENTIALS_PATH=[PATH_TO_GOOGLE_AUTH_JSON_KEY] 
``

### Run

``
&& docker run  -p 9000:9000 -e DATASTORE_PROJECT_ID=[GCP_PROJECT_NAME] upload_service
``