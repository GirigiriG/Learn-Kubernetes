## Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mongo-deployment
  labels:
    app: mongo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mongo
  template:
    metadata:
      labels:
        app: mongo
    spec:
      containers:
      - name: mongodb
        image: mongo:5.0
        ports:
        - containerPort: 27017
```

### Template:

- Configuration for a kubernetes Pod.

### Labels:

- You can give any K8s resource a label.
- The label is a key/value pair that are attached to K8s resources.
- Labels do not provide uniqueness. All Pod replicas will have the same label.
- For Pods, label is a required field

```yaml
labels:
    app: mongo
```

- Identify sets of resources for a `Deployment` using `matchlabels`

```yaml
spec:
  replicas: 3
  selector:
    matchLabels:
      app: mongo
```

## Service

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mongo-service
spec:
  selector:
  #this is how the service and pod will find each other (ln: 15)
    app: mongo
  ports:
    - protocol: TCP
      port: 27017
      # target port should always be the same and the container port (ln: 21)
      targetPort: 27017
```

![Screen Shot 2022-10-17 at 9.51.37 PM.png](./images/service-flow.png)