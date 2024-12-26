```bash
$ kubectl get pods
NAME                                READY   STATUS    RESTARTS   AGE
crud-app-9c89f7454-hmckz            1/1     Running   0          5m10s
hashresponse-dep-755b5b5dd7-v9lvk   1/1     Running   0          30s
$ kubectl port-forward crud-app-9c89f7454-hmckz 3003:3000
Forwarding from 127.0.0.1:3003 -> 3000
Forwarding from [::1]:3003 -> 3000
Handling connection for 3003
```
### Another terminal window:

```bash
$ curl http://localhost:3003/
Hello, World!
```
