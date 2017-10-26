
## Sample Operator
The sample operator creates a custom resources and watches for changes.

### Build
```bash
# pull all the libraries needed for operator-kit (this may take a while with all the Kubernetes dependencies)
glide install  --strip-vendor

# change directory to sample-operator
cd sample-operator

# build the sample operator binary
CGO_ENABLED=0 GOOS=linux go build

# build the docker container
docker build -t sample-operator:0.1 .
```

### Start the Operator

```bash
# Create the sample operator
$ kubectl create -f sample-operator.yaml

# Wait for the pod status to be Running
$ kubectl get pod -l app=sample-operator
NAME                              READY     STATUS    RESTARTS   AGE
sample-operator-821691060-m5vqp   1/1       Running   0          3m

# View the samples CRD
$ kubectl get crd samples.mycompany.io
NAME                   KIND
samples.mycompany.io   CustomResourceDefinition.v1beta1.apiextensions.k8s.io
```

### Create the Sample Resource
```bash
# Create the sample
$ kubectl create -f sample-resource.yaml

# See the sample resource
$ kubectl get samples.mycompany.io
NAME       KIND
mysample   Sample.v1alpha1.mycompany.io
```

### Modify the Sample Resource
Change the value of the `Hello` property in `sample-resource.yaml`, then apply the new yaml.
```bash
kubectl apply -f sample-resource.yaml
```

### Logs

Notice the added and modified Hello= text in the log below

```bash
$ kubectl logs -l app=sample-operator
2017-06-29 00:11:22.738629 I | sample: Creating the sample resource
2017-06-29 00:11:22.749325 I | op-kit: creating sample resource
2017-06-29 00:11:22.794635 I | op-kit: did not yet find resource sample at apis/mycompany.io/v1alpha1/samples. the server could not find the requested resource
2017-06-29 00:11:28.797620 I | op-kit: did not yet find resource sample at apis/mycompany.io/v1alpha1/samples. the server could not find the requested resource
2017-06-29 00:11:34.797912 I | sample: Managing the sample resource
2017-06-29 00:11:34.797932 I | sample: finding existing samples...
2017-06-29 00:11:34.799276 I | sample: found 0 samples.
2017-06-29 00:11:34.799296 I | op-kit: start watching sample resource in namespace default at 27064
2017-06-29 00:12:07.605948 I | sample: Added Sample 'mysample' with Hello=World!
2017-06-29 00:14:29.553035 I | sample: Modified Sample 'mysample' with Hello=ANOTHER World!!
```

### Cleanup
```bash
kubectl delete -f sample-resource.yaml
kubectl delete -f sample-operator.yaml
kubectl delete crd samples.mycompany.io
```
