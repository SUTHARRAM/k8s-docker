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

## minikube addons enable metrics-server

```bash
minikube addons enable metrics-server
```

## kubectl top pods

```bash
kubectl top pods
```

## kubectl top pods --watch

```bash
kubectl top pods --watch
```

## kubectl get pods -n kube-system


```bash
kubectl get pods -n kube-system
```

## kubectl top pods -n kube-system

```bash
kubectl top pods -n kube-system
```

## kubectl top pods --all-namespaces

```bash
kubectl top pods --all-namespaces
```

## kubectl top nodes

```bash
kubectl top nodes
```

## kubectl top nodes --watch

```bash
kubectl top nodes --watch
```

## kubectl top nodes --all-namespaces

```bash
kubectl top nodes --all-namespaces
```

## kubectl top nodes --all-namespaces --watch


```bash
kubectl top nodes --all-namespaces --watch
```

## kubectl top nodes --all-namespaces --watch

```bash
kubectl top nodes --all-namespaces --watch
```
<!-- For auto scaling testing  -->

## Pod Watching 

```bash
kubectl get pods -w
```

## HPA Watching 

```bash
kubectl get hpa -w
```

## For stress on pod 

```bash 
kubectl exec -it user-service-5d8f9966cd-rbtxn -- sh
```

```bash 
while true; do curl http://localhost:8080/user-service/health; done
```

<!-- OR  run following command on pod >


while :; do :; done
