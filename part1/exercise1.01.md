```bash
$ kubectl create deployment log-output --image toukop/log-output:latest
deployment.apps/log-output created
$ kubectl get pods
NAME                                 READY   STATUS    RESTARTS   AGE
log-output-6ff7b49f5b-5lwkt   1/1     Running   0          39s
$ kubectl logs -f log-output-6ff7b49f5b-5lwkt
2024-12-25T12:15:38Z: e6ba7043-1f04-4d78-b35a-fee9868da0b1
2024-12-25T12:15:43Z: e6ba7043-1f04-4d78-b35a-fee9868da0b1
2024-12-25T12:15:48Z: e6ba7043-1f04-4d78-b35a-fee9868da0b1
2024-12-25T12:15:53Z: e6ba7043-1f04-4d78-b35a-fee9868da0b1```
```
