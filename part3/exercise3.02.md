```bash
➜  pong git:(main) ✗ kubectl get pods,ing,services
NAME                             READY   STATUS    RESTARTS   AGE
pod/log-output-668977f44-w7fxk   2/2     Running   0          5m48s
pod/pong-b889858d4-4hp99         1/1     Running   0          100s
pod/pong-postgres-stset-0        1/1     Running   0          62m

NAME                                           CLASS    HOSTS   ADDRESS        PORTS   AGE
ingress.networking.k8s.io/log-output-ingress   <none>   *       34.8.132.181   80      3m45s
ingress.networking.k8s.io/pong-ingress         <none>   *       34.8.81.65     80      24m

NAME                        TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
service/log-output-svc      NodePort    34.118.231.125   <none>        2345:30845/TCP   40m
service/pong-postgres-svc   ClusterIP   None             <none>        5432/TCP         72m
service/pong-svc            NodePort    34.118.232.157   <none>        2346:32334/TCP   32m
➜  pong git:(main) ✗ curl http://34.8.132.181/
file content: this text is from file
MESSAGE=hello world
2025-01-05T16:29:36Z: 1f1f196b-a4b4-4e2d-ace5-bed5ae800a4e
Ping / Pongs: CANNOT_READ_PONG_FILE
➜  pong git:(main) ✗ curl http://34.8.81.65/pingpong
pong 323%
```
