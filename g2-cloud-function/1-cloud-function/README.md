### **Step 1: Create a New Go Function**
1. Create a directory for your project:
   ```bash
   mkdir hello-world-function
   cd hello-world-function
   ```
2. Create the `go.mod` file:
   ```bash
   go mod init hello-world-function
   ```
3. Create the main Go file (`main.go`):
   ```go
   package hello

   import (
       "fmt"
       "net/http"
   )

   // HelloWorld is an HTTP Cloud Function.
   func HelloWorld(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintln(w, "Hello, World!")
   }
   ```

---

### **Step 2: Prepare Deployment Files**
1. Create a `function.yaml` file (optional but recommended for organized deployment):
   ```yaml
   name: hello-world-function
   runtime: go121
   entryPoint: HelloWorld
   region: us-east1
   ```

---

### **Step 3: Deploy the Function**
1. Deploy using the Google Cloud CLI:
   ```bash
   gcloud functions deploy hello-world-function \
       --runtime go121 \
       --trigger-http \
       --allow-unauthenticated \
       --entry-point HelloWorld \
       --region us-east1
   ```

2. Wait for the deployment to complete. Once deployed, you’ll receive a URL for your Cloud Function.

---

### **Step 4: Test the Function**
1. Use `curl` or a browser to access the Cloud Function:
   ```bash
   curl https://[REGION]-[PROJECT_ID].cloudfunctions.net/hello-world-function
   ```

   The response should be:
   ```
   Hello, World!
   ```

---

### **Step 5: Clean Up**
1. If you no longer need the function, delete it to avoid incurring costs:
   ```bash
   gcloud functions delete hello-world-function --region us-central1
   ```

---

### **Tips**
- **Testing Locally**: Use tools like `go run` and `curl` for quick tests.
- **Authentication**: By default, the function is public (`--allow-unauthenticated`). If you want it private, skip `--allow-unauthenticated` and use `gcloud auth print-identity-token` for testing.
