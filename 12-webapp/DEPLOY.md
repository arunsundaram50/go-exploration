### **1. Prepare Environment**
1. **Enable APIs**:
   Ensure the required APIs are enabled:
   ```bash
   gcloud services enable cloudbuild.googleapis.com deploy.googleapis.com
   ```

2. **Set Up IAM Roles**:
   Grant necessary roles to service account:
   ```bash
   gcloud projects add-iam-policy-binding $PROJECT_ID \
       --member="serviceAccount:$PROJECT_NUMBER@cloudbuild.gserviceaccount.com" \
       --role="roles/clouddeploy.admin"
   ```

---

### **2. Define  Cloud Deploy Pipeline**
1. Create a **Cloud Deploy delivery pipeline configuration file** (e.g., `delivery-pipeline.yaml`):
   ```yaml
   apiVersion: clouddeploy.googleapis.com/v1
   kind: DeliveryPipeline
   metadata:
     name: hello-world-pipeline
   spec:
     serialPipeline:
       stages:
       - targetId: dev
       - targetId: prod
   ```

2. Create **targets** for each stage (e.g., `dev`, `prod`):
   ```yaml
   apiVersion: clouddeploy.googleapis.com/v1
   kind: Target
   metadata:
     name: dev
   description: "Development environment"
   ```
   Repeat for `prod`.

3. Apply the pipeline and targets:
   ```bash
   gcloud deploy apply --file=delivery-pipeline.yaml --region=us-east1
   ```

---

### **3. Prepare Cloud Build Config**
Create a `cloudbuild.yaml` file to build container image and deploy it.

Example:
```yaml
steps:
- name: 'gcr.io/cloud-builders/docker'
  args: ['build', '-t', 'gcr.io/$PROJECT_ID/hello-world:$COMMIT_SHA', '.']
- name: 'gcr.io/cloud-builders/docker'
  args: ['push', 'gcr.io/$PROJECT_ID/hello-world:$COMMIT_SHA']
- name: 'gcr.io/google.com/cloudsdktool/cloud-sdk'
  args:
  - gcloud
  - deploy
  - releases
  - create
  - hello-world-release-$COMMIT_SHA
  - --delivery-pipeline=hello-world-pipeline
  - --region=us-east1
  - --images=gcr.io/$PROJECT_ID/hello-world=$COMMIT_SHA
substitutions:
  _COMMIT_SHA: "$(git rev-parse --short HEAD)"
```

---

### **4. Trigger Cloud Build**
Set up a trigger in Cloud Build to automate builds and deployments when changes are pushed to a branch.

1. Go to the **Triggers** page in Google Cloud Console.
2. Create a new trigger and specify:
   - Source: repository and branch.
   - Build Configuration: Use the `cloudbuild.yaml` file.
3. Save the trigger.

---

### **5. Deploy Using Cloud Deploy**
When the Cloud Build pipeline completes, Cloud Deploy automatically creates a release and deploys it to the specified targets (`dev` and `prod`).

---

### **6. Verify Deployment**
1. Check the status of delivery pipeline:
   ```bash
   gcloud deploy releases list --delivery-pipeline=hello-world-pipeline --region=us-east1
   ```
2. View deployment logs:
   ```bash
   gcloud deploy releases describe hello-world-release-$COMMIT_SHA --region=us-east1
   ```

---

### **Summary**
- **Cloud Build** handles containerization and triggers.
- **Cloud Deploy** orchestrates multi-environment deployments using delivery pipelines.


## Next steps
- Automate the process by configuring triggers in Cloud Build to kick off the build and deploy when code is committed.