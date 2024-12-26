```bash
$ kubectl apply -f manifests/deployment.yaml
deployment.apps/log-output created
$ kubectl get pods
NAME                                 READY   STATUS    RESTARTS   AGE
crud-app-9c89f7454-5q59w             1/1     Running   0          7m40s
log-output-6ff7b49f5b-hjmxr   1/1     Running   0          12s
$ kubectl logs -f log-output-6ff7b49f5b-hjmxr
2024-12-25T12:59:22Z: 45fd4f42-04e4-404a-98be-e4d71f469168
2024-12-25T12:59:27Z: 45fd4f42-04e4-404a-98be-e4d71f469168
2024-12-25T12:59:32Z: 45fd4f42-04e4-404a-98be-e4d71f469168
2024-12-25T12:59:37Z: 45fd4f42-04e4-404a-98be-e4d71f469168
```
