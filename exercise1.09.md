```bash
$ kubectl get pods,svc,ing
NAME                             READY   STATUS    RESTARTS   AGE
pod/log-output-f79f9b568-ctl9j   1/1     Running   0          49s
pod/pong-8c7d6b9d-nvckq          1/1     Running   0          93s
                                                                                                                                                             NAME                     TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
service/kubernetes       ClusterIP   10.43.0.1      <none>        443/TCP          41m
service/log-output-svc   NodePort    10.43.227.19   <none>        2345:31270/TCP   49s
service/pong-svc         NodePort    10.43.98.32    <none>        2346:31692/TCP   93s

NAME                                           CLASS     HOSTS   ADDRESS                            PORTS   AGE
ingress.networking.k8s.io/log-output-ingress   traefik   *       172.20.0.3,172.20.0.4,172.20.0.5   80      49s
ingress.networking.k8s.io/pong-ingress         traefik   *       172.20.0.3,172.20.0.4,172.20.0.5   80      93s
$ curl http://localhost:8081/
2024-12-26T08:11:00Z: 6ca8482e-faf9-4967-8454-158120ef323b
$ curl http://localhost:8081/pingpong
pong 4
