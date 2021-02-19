# Network policies

Start minikube with `minikube.sh`.

Install linkerd-cni and linkerd.

```bash
linkerd install-cni | kubectl apply -f -
linkerd install --linkerd-cni-enabled | kubectl apply -f -
```

Deploy the service 
`kubectl apply -f deployment.yaml`.

Service is of type `LoadBalancer`, if using with Minikube use `minikube tunnel` in separate terminal.
Check `kubectl get svc nginx -n web-demo -o jsonpath='{.status.loadBalancer.ingress[0].ip}'` for External IP.