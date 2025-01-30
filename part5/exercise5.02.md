## Injecting linkerd to crud-app

```bash
➜  kubernetes-course git:(main) kubectl kustomize crud-app | linkerd inject - | kubectl apply -f -


service "crud-app-svc" skipped
deployment "crud-app" injected
ingress "crud-app-ingress" skipped

service/crud-app-svc unchanged
deployment.apps/crud-app configured
ingress.networking.k8s.io/crud-app-ingress unchanged

➜  kubernetes-course git:(main) kubectl -n crud get po -o jsonpath='{.items[0].spec.containers[*].name}'

linkerd-proxy crud-app%
```

## Injecting all apps (rest) in crud namespace

```bash
➜  kubernetes-course git:(main) kubectl get -n crud deploy -o yaml | linkerd inject - | kubectl apply -f -


deployment "crud-app" injected
deployment "crud-backend" injected

deployment.apps/crud-app configured
deployment.apps/crud-backend configured
```
