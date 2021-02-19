# Network policies

Start minikube with `minikube.sh`.

[Install](https://docs.projectcalico.org/getting-started/kubernetes/minikube) Calico

If needed, restart minikube.

Deploy the service 
`kubectl apply -f deployment.yaml`.

Service is of type `LoadBalancer`, if using with Minikube use `minikube tunnel` in separate terminal.
Check `kubectl get svc nginx -n web-demo -o jsonpath='{.status.loadBalancer.ingress[0].ip}'` for External IP.

Run another pod
```
kubectl run -i --tty --rm access --image=busybox --namespace=web-demo -- /bin/sh
```

Try

```
wget -q --timeout=5 nginx -O -
```

Run another pod
```
kubectl run -i --tty --rm no-access --image=busybox --namespace=web-demo -- /bin/sh
```

Try

```
wget -q --timeout=5 nginx -O -
```

And see, that it won't work