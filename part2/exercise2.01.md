```bash
$ kubectl get pods
NAME                        READY   STATUS    RESTARTS   AGE
log-output-fc58ccb9-qlwhm   2/2     Running   0          54s
pong-6cf4b88ccf-mgzpt       1/1     Running   0          60s
$ kubectl exec -it log-output-fc58ccb9-qlwhm -- bash
Defaulted container "log-output-worker" out of: log-output-worker, log-output-web
root@log-output-fc58ccb9-qlwhm:/app# curl http://log-output-svc:2345
2024-12-26T14:54:00Z: 2d220be8-dd85-44ba-ba9e-2181d9d22f8b
Ping / Pongs: 5
```
