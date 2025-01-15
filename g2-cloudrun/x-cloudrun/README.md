Here are the steps to create and run a Golang "Hello World" application in Google Cloud Run:

### 1. **Set Up Your Environment**
   - Install the [Google Cloud SDK](https://cloud.google.com/sdk/docs/install).
   - Install [Docker](https://www.docker.com/products/docker-desktop).
   - Install Golang if not already installed: [Go installation](https://go.dev/dl/).
   - Authenticate with Google Cloud:
     ```bash
     gcloud auth login
     gcloud config set project $PROJECT_ID
     gcloud auth configure-docker
     gcloud components update
     ```

---

### 2. **Write the Go Application**
   Create a file named `main.go`:
   ```go
   package main

   import (
       "fmt"
       "net/http"
   )

   func helloWorld(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintln(w, "Hello, World!")
   }

   func main() {
       http.HandleFunc("/", helloWorld)
       http.ListenAndServe(":8080", nil)
   }
   ```

---

### 3. **Create a Dockerfile**
   In the same directory, create a `Dockerfile`:
   ```dockerfile
   # Use the official Golang image to build the app
   FROM golang:1.20 as builder

   # Set the working directory
   WORKDIR /app

   # Copy go files
   COPY . .

   # Build the Go app
   RUN go build -o main .

   # Use a minimal base image
   FROM gcr.io/distroless/base-debian10

   # Copy the app from the builder stage
   COPY --from=builder /app/main /

   # Command to run the executable
   CMD ["/main"]
   ```

---

### 4. **Build the Docker Image**
   Build the container image:
   ```bash
   docker build -t gcr.io/$PROJECT_ID/hello-world .
   ```

---

### Test it locally
```
docker run -p 8080:8080 gcr.io/$PROJECT_ID/hello-world
```

### 5. **Push the Image to Container Registry**
   Push the image to Google Container Registry:
   ```bash
   docker push gcr.io/$PROJECT_ID/hello-world
   ```

---

### 6. **Deploy to Cloud Run**
   Deploy the container to Cloud Run:
   ```bash
   gcloud run deploy hello-world \
       --image gcr.io/$PROJECT_ID/hello-world \
       --platform managed \
       --region us-east1 \
       --allow-unauthenticated
   ```

---

### 7. **Access the Application**
   Once deployed, you’ll get a URL. Open it in your browser to see "Hello, World!" displayed.


### Response
```
https://hello-world-363357136974.us-east1.run.app
```


### 8. Delete it
```
gcloud run services list
gcloud run services delete hello-world --region us-east1

```