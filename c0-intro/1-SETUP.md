**LocalStack** is a tool for running **AWS services locally** for testing and development. It provides a **Docker-based** local environment that emulates a wide range of AWS services.

### 🔹 **Services Provided by LocalStack**
Here’s a categorized list of AWS services supported by LocalStack:

| Category              | AWS Service            | LocalStack Support |
|----------------------|----------------------|-------------------|
| **Compute**          | Lambda               | ✅ Yes |
|                      | EC2                  | ✅ Yes (Limited) |
| **Storage**         | S3                   | ✅ Yes |
|                      | EBS                  | ✅ Yes (Experimental) |
|                      | EFS                  | ✅ Yes (Experimental) |
| **Databases**        | DynamoDB             | ✅ Yes |
|                      | RDS (PostgreSQL, MySQL) | ✅ Yes (Limited) |
|                      | ElastiCache (Redis)  | ✅ Yes |
| **Messaging**        | SQS                  | ✅ Yes |
|                      | SNS                  | ✅ Yes |
| **Networking**       | VPC                  | ✅ Yes (Limited) |
|                      | Route 53             | ✅ Yes |
|                      | API Gateway          | ✅ Yes |
| **Security**         | IAM                  | ✅ Yes |
|                      | Secrets Manager      | ✅ Yes |
|                      | Cognito              | ✅ Yes |
| **CI/CD**           | CodeCommit           | ✅ Yes |
|                      | CodeBuild            | ✅ Yes |
|                      | CodePipeline         | ✅ Yes |
| **Container Services** | ECS (Fargate)        | ✅ Yes (Limited) |
|                      | EKS                  | ❌ No |
| **Machine Learning** | SageMaker            | ❌ No |
| **Monitoring**      | CloudWatch           | ✅ Yes |
|                      | X-Ray                | ✅ Yes |

---

### 🔹 **How to Run LocalStack**
To spin up LocalStack using **Docker**, use:

```sh
docker run --rm -it -p 4566:4566 localstack/localstack
```

Then, interact with LocalStack using the **AWS CLI**:

```sh
aws --endpoint-url=http://localhost:4566 s3 ls
aws --endpoint-url=http://localhost:4566 sqs list-queues
aws --endpoint-url=http://localhost:4566 lambda list-functions
```

For **advanced setup**, you can use `docker-compose` or install **LocalStack Pro** for more features.

---

### 🔹 **Why Use LocalStack?**
- ✅ **No cloud costs** – Run AWS locally without using real AWS.
- ✅ **Fast iterations** – No need to deploy to AWS for testing.
- ✅ **Works offline** – No internet connection required.
- ✅ **Mimics real AWS APIs** – Ideal for integration testing.
