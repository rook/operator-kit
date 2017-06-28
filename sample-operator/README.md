
## Sample Operator
The sample operator creates a custom resources and watches for changes.

### Build
```bash
# pull all the libraries needed (this may take a while with all the Kubernetes dependencies)
glide install

# build the go binary
CGO_ENABLED=0 go build

# build the docker container
docker build -t sample-operator:0.1 .
```

### Start the Operator

```bash
kubectl create -f sample-operator.yaml
```

### Create the Sample Resource
```bash
kubectl create -f sample-resource.yaml
```

### Artifacts
Now look at the artifacts of the operator and resources created.
```bash
$ kubectl get pod -l app=sample-operator

$ kubectl get thirdpartyresource sample.mycompany.io

$ kubectl get sample.mycompany.io

$ kubectl logs -l app=sample-operator
```

### Cleanup
```bash
kubectl delete -f sample-resource.yaml
kubectl delete -f sample-operator.yaml
kubectl delete thirdpartyresource sample.mycompany.io
```