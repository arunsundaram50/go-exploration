Schedule a Lambda function `myLambda` to run every hour using **Amazon EventBridge** (formerly CloudWatch Events) via the AWS CLI.

### Steps:

#### 1. **Create an EventBridge Rule**
Use the following AWS CLI command to create a rule that triggers every hour:

```sh
aaws events put-rule \
    --name myLambdaHourlyTrigger \
    --schedule-expression "rate(1 hour)" \
    --state ENABLED
```

Alternatively, if a cron expression (UTC timezone) can be used:

```sh
aaws events put-rule \
    --name myLambdaHourlyTrigger \
    --schedule-expression "cron(0 * * * ? *)" \
    --state ENABLED
```
This cron expression means **"Run at minute 0 of every hour."**

#### 2. **Grant Permissions for EventBridge to Invoke Lambda**
Attach the necessary permissions to the Lambda function:

```sh
aaws lambda add-permission \
    --function-name myLambda \
    --statement-id myLambdaHourlyPermission \
    --action "lambda:InvokeFunction" \
    --principal events.amazonaws.com \
    --source-arn arn:aws:events:REGION:ACCOUNT_ID:rule/myLambdaHourlyTrigger
```
Replace:
- `REGION` (e.g., `us-east-1`)
- `ACCOUNT_ID` (the AWS account ID)

#### 3. **Associate the Rule with the Lambda Function**
Now, create the target for the rule:

```sh
aaws events put-targets \
    --rule myLambdaHourlyTrigger \
    --targets "Id"="1","Arn"="$(aaws lambda get-function --function-name myLambda --query 'Configuration.FunctionArn' --output text)"
```

### Verification:
Check if the rule was created:
```sh
aaws events list-rules --name-prefix myLambdaHourlyTrigger
```

Check if the target is attached:
```sh
aaws events list-targets-by-rule --rule myLambdaHourlyTrigger
```

---

This setup ensures that the Lambda function `myLambda` runs **every hour** automatically. 🚀