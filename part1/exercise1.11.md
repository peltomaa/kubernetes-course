```bash
$ kubectl get svc,ing,pods
NAME                     TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
service/kubernetes       ClusterIP   10.43.0.1      <none>        443/TCP          48m
service/log-output-svc   NodePort    10.43.177.10   <none>        2345:31736/TCP   7m11s
service/pong-svc         NodePort    10.43.60.120   <none>        2346:31163/TCP   8m
                                                                                                                                                                                            NAME                                           CLASS     HOSTS   ADDRESS                            PORTS   AGE
ingress.networking.k8s.io/log-output-ingress   traefik   *       172.24.0.3,172.24.0.4,172.24.0.5   80      7m11s
ingress.networking.k8s.io/pong-ingress         traefik   *       172.24.0.3,172.24.0.4,172.24.0.5   80      8m
                                                                                                                                                                                            NAME                            READY   STATUS    RESTARTS   AGE
pod/log-output-fc58ccb9-m799z   2/2     Running   0          7m11s
pod/pong-6cf4b88ccf-nscvx       1/1     Running   0          8m
$ curl http://localhost:8081/pingpong
pong 4
$ curl http://localhost:8081/
2024-12-26T11:46:42Z: 8bf3db37-217a-446f-bfe0-8e9571d57884
Ping / Pongs: 5
