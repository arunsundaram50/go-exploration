```
   gcloud functions deploy hello-world-fiber-function \
       --runtime go121 \
       --trigger-http \
       --allow-unauthenticated \
       --entry-point HelloWorld \
       --region us-east1
```
