
```bash
$ kubectl create deployment crud-app --image toukop/crud-app:latest
deployment.apps/crud-app created
$ kubectl get pods
NAME                       READY   STATUS    RESTARTS   AGE
crud-app-9c89f7454-5q59w   1/1     Running   0          7s
$ kubectl logs -f crud-app-9c89f7454-5q59w
Starting server
```

### Note:

```basj
Server started in port 3000
```

is not displayed as the port 3000 is not available for the application
