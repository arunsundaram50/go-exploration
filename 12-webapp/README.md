1) Source code main.go
2) Executable (go build): funny-app.exe
3) Create a docker image (embed funny-app.exe)
  3a) Build and test it locally
  ```
    env GOOS=linux GOARCH=amd64 go build -o funny-app 
    docker build -t funny-app . 
    docker images|grep funny
    docker run -p8080:8080 IMAGE_ID 
  ```
  3b) Upload it to GCP and deploy it as a cloud run service
  ```
  docker build -t gcr.io/$PROJECT_ID/funny-app .
  gcloud auth login
  gcloud auth print-access-token | docker login -u oauth2accesstoken --password-stdin https://gcr.io/$PROJECT_ID\
  docker push gcr.io/$PROJECT_ID/funny-app
  gcloud run deploy funny-app --image gcr.io/$PROJECT_ID/funny-app \
    --platform managed --region us-east1 --allow-unauthenticated --port=8080
  ```
  The last command should give us the URL of the deployed RestAPI service
