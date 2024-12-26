```bash
$ kubectl apply -f manifests/
deployment.apps/log-output created
ingress.networking.k8s.io/log-output-ingress created
service/log-output-svc created
$ kubectl get pods,svc,ing
NAME                                     READY   STATUS    RESTARTS   AGE
pod/log-output-6ff7b49f5b-n97t7   1/1     Running   0          25s

NAME                            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
service/kubernetes              ClusterIP   10.43.0.1       <none>        443/TCP          155m
service/log-output-svc   NodePort    10.43.134.251   <none>        2345:31030/TCP   25s

NAME                                                  CLASS     HOSTS   ADDRESS                            PORTS   AGE
ingress.networking.k8s.io/log-output-ingress   traefik   *       172.19.0.3,172.19.0.4,172.19.0.5   80      25s
$ kubectl logs -f log-output-6ff7b49f5b-n97t7
2024-12-25T20:28:34Z: 403f9975-9291-4c31-b89b-7bc60caa7a47
Server started in port 3000
2024-12-25T20:28:39Z: 403f9975-9291-4c31-b89b-7bc60caa7a47
2024-12-25T20:28:44Z: 403f9975-9291-4c31-b89b-7bc60caa7a47
2024-12-25T20:28:49Z: 403f9975-9291-4c31-b89b-7bc60caa7a47
2024-12-25T20:28:54Z: 403f9975-9291-4c31-b89b-7bc60caa7a47
$ curl http://localhost:8081/
2024-12-25T20:29:15Z: 403f9975-9291-4c31-b89b-7bc60caa7a47
