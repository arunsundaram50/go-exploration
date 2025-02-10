# go-exploration

# GCP

1. Create a project
2. Create a service account
3. Create a key for the service account
4. Create a bucket
5. Create a VM instance
6. Create a VPC network
7. Create a firewall rule
8. Create a static IP address
9. Create a Cloud Run service
10. Create a Cloud Run job
11. Create a Cloud Run revision
12. Create a Cloud Run service revision
13. Create a Cloud Run service revision

# Using GCP to setup a CI/CD pipeline, starting from Go code in GitHub, using Cloud Build and Deploy
1. Create a GitHub Actions workflow
2. Create a Cloud Run service
3. Create a Cloud Run job
4. Create a Cloud Run revision
5. Create a Cloud Run service revision
6. Create a Cloud Build trigger
7. Create a Cloud Deploy trigger
8. Create a Cloud Deploy target
9. Create a Cloud Deploy release
10. Create a Cloud Deploy rollout

## How to create a GitHub Actions workflow?
1. Create a GitHub Actions workflow file in the .github/workflows directory
2. Add a trigger to the workflow file
3. Add a job to the workflow file
4. Add a step to the job file

### Create a GitHub Actions workflow file in the .github/workflows directory
1. Create a new file in the .github/workflows directory
2. Add the following content to the file:

```yaml
name: CI/CD Pipeline

on:
  push:
    branches:
      - main
```
