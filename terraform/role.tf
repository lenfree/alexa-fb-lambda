resource "aws_iam_policy" "lambda_to_cloudwatch" {
  name        = "lambda-function-to-cloudwatch"
  description = "lambda to cloudwatch logs"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": [
        "logs:*"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:logs:us-east-1:*:*"
    }
  ]
}
EOF
}

resource "aws_iam_policy" "invoke_lambda_func" {
  name        = "invoke_lambda_func"
  description = "invoke_lambda_func"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "lambda:InvokeFunction",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "alexa_fb_lambda" {
  role       = "${aws_iam_role.alexa_fb_lambda.name}"
  policy_arn = "${aws_iam_policy.lambda_to_cloudwatch.arn}"
}

resource "aws_iam_role" "alexa_fb_lambda" {
  name = "alexa_fb_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_permission" "alexa" {
  statement_id  = "AllowExecutionFromAlexa"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.alexa_fb_lambda.function_name}"
  principal     = "alexa-appkit.amazon.com"
}

resource "aws_lambda_function" "alexa_fb_lambda" {
  filename      = "./../build/handler.zip"
  function_name = "myfacebook_alexa"
  role          = "${aws_iam_role.alexa_fb_lambda.arn}"
  handler       = "handler.Handle"
  runtime       = "python2.7"

  environment {
    variables = {
      ALEXA_APP_ID = "<change_me>"
      APP_ID       = "<change_me>"
      APP_SECRET   = "<change_me>"
    }
  }
}

output "lambda_role_arn" {
  value = "${aws_iam_role.alexa_fb_lambda.arn}"
}

output "lambda_arn" {
  value = "${aws_lambda_function.alexa_fb_lambda.arn}"
}