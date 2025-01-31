# Kill Runner - CKAD Scenario Simulator

A Golang tool to validate CKAD test scenarios!

# Answers To Questions

<details><summary>Answer to Question 1</summary>
<h3>Obtain Namespaces</h3>

The following BASH command explains how to complete this task

```bash
k get ns > Q1/namespaces.txt
```
</details>

<details><summary>Answer to Question 2</summary>
<h3>Create A Single Pods</h3>

The following YAML & BASH command explains how to complete this task

```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: pod1
  name: pod1
  namespace: q2-ns
spec:
  containers:
  - image: httpd:2.4.41-alpine
    name: pod1-container
  dnsPolicy: ClusterFirst
  restartPolicy: Always
```

```bash
k -n q2-ns describe po pod1 | grep -i status: | cut -d ":" -f2 | grep -o "[A-Z].*" > Q2/pod1-status-command.sh
```
</details>



<details><summary>Answer to Question 3</summary>
<h3>Create A Single Job</h3>

The following YAML file explains how to create this job

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: awesome-job
  namespace: q3-ns
spec:
  completions: 3
  parallelism: 2
  template:
    metadata:
      labels:
        id: awesome-job
    spec:
      containers:
      - command:
        - /bin/sh
        - -c
        - sleep 2 && echo done
        image: busybox:1.31.0
        name: awesome-job
      restartPolicy: Never
```
</details>
