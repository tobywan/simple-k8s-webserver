# Task

## Purpose

This exercise aims for you to create and deploy a simple hello world docker container to a
Kubernetes cluster and have a friendly technical discussion about your implementation.
During the exercise you will create a Dockerfile, push the image and deploy it to a Kubernetes
cluster.
There are no right or wrong answers to this exercise, and you are not expected to finish it. If you
want to timebox this to an hour, that is fine. If there are parts you are unsure of document, why you
made those decisions and let us discuss it. If there are steps you think you could shortcut in some
way to get to tasks further on, feel free to do that and talk about the caveats.

Exercise
- Create 'hello world' http server that returns 200 OK on port 8080
  - You can use a something simple like a golang example, or python flask, sinatra or
any other alternative web server in any language
- Create and package the app in a production ready docker container
  - You can build it inside the container or copy it in
- Create an ECR repo with terraform or use your personal Dockerhub repo and push your
container to the repo
- Create a k8s cluster (you could use terraform to create an EKS cluster, or just docker for mac
or kind)
- Create a Kubernetes deployment and service to deploy the docker image
  - You can use `helm create` if you wish, or copy and paste an example, or write the YAML
by hand

Considerations
- Bonus if any part of the solution employs what you deem as production “best practices” e.g.
  - The service emits metrics, logs or traces
  - The terraform is well structured and tested
  - The Kubernetes cluster is maintained via GitOps
- Well documented and commented code
- Notes about your approach and any shortcuts, caveats or thoughts whilst undertaking the
task

## Todo

- [x] Create web server
- [x] Add prometheus metrics
- [x] Add w3scv logger
- [x] Add docker file
- [x] Build docker image on push to GH
- [ ] Set up AWS account
- [ ] Terraform:
  - [ ] Create ECR repo with Lifecycle policy
  - [ ] IAM roles
  - [ ] EKS Cluster
- [ ] GitOps


