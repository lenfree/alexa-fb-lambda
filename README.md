A lambda function that process myfacebook skill
which queries facebook inbox count new/unread
messages and from whom. This is triggered by
Alexa.

Usage:
=====

Start a docker-machine:

```bash
$ docker-machine create -d virtualbox aws-lambda
$ eval $(docker-machine env aws-lambda)
```

Build a zip file for lambda function:

```bash
$ cd build
$ make
```

Use Terraform to manage AWS resources:

```bash
$ cd terraform
$ terraform apply
$ terraform destroy
```
