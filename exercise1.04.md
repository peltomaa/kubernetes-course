```bash
$ kubectl create -f manifests/deployment.yaml
deployment.apps/crud-app created
peltomaa@DESKTOP-3PT38CP:~/School/kubernetes-course/crud-app$ kubectl get pods
NAME                       READY   STATUS    RESTARTS   AGE
crud-app-9c89f7454-hmckz   1/1     Running   0          5s
$ kubectl logs -f crud-app-9c89f7454-hmckz
Starting server
```
### Note:

```basj
Server started in port 3000
```

is not displayed as the port 3000 is not available for the application
