# Make aaws work for localstack and for cloud-AWS
```
export AWS_PROFILE=localstack
source aaws-definition.sh
```

## Play with buckets
```
aaws s3 mb s3://my-bucket
aaws s3 ls

echo "Hello LocalStack" > test.txt
aaws s3 cp test.txt s3://my-bucket/
rm test.txt
aaws s3 ls s3://my-bucket/
# just the object contents
aaws s3 cp s3://my-bucket/test.txt -
# When needing metadata along with the content.
aaws s3api get-object --bucket my-bucket --key test.txt /dev/stdout 
```

## Play with Lambda
```
cat <<EOF > lambda_function.py
def lambda_handler(event, context):
    return {"message": "Hello from LocalStack Lambda!"}
EOF

zip lambda.zip lambda_function.py

aaws lambda create-function \
    --function-name myLambda \
    --runtime python3.8 \
    --role arn:aws:iam::000000000000:role/lambda-role \
    --handler lambda_function.lambda_handler \
    --zip-file fileb://lambda.zip

rm lambda.zip

aaws lambda invoke \
    --function-name myLambda output.json
cat output.json
rm output.json
```
