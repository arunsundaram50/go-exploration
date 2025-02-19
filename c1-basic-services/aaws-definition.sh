# Do a 
#   source aaws-definition,sh
# and use it in place of aws
# example: 
#   aaws lambda invoke --function-name myLambda1 x.json
aaws() {
  if [ "$AWS_PROFILE" = "localstack" ]; then
    aws --endpoint-url=http://localhost:4566 "$@"
  else
    aws "$@"
  fi
}
