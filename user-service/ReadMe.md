## Docker commands

```bash
docker build -t user-service-img .
docker run -d -p 8080:8080 --name user-service user-service-img
```

## Kubernetes commands

```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```
## Remove Docker Images

```bash
docker rmi user-service-img
```

## Remove Kubernetes Resources

```bash
kubectl delete -f k8s/deployment.yaml
kubectl delete -f k8s/service.yaml
```

## Remove Docker Containers

```bash
docker rm user-service
```
