```bash

$ kubectl apply -f manifests/ingress.yaml
ingress.networking.k8s.io/crud-app-ingress created
$ kubectl apply -f manifests/service.yaml
service/crud-app-svc created
$ kubectl get pods
NAME                       READY   STATUS    RESTARTS   AGE
crud-app-9c89f7454-5784k   1/1     Running   0          122m
$ kubectl get svc,ing
NAME                   TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
service/crud-app-svc   NodePort    10.43.146.109   <none>        2345:32275/TCP   2m29s
service/kubernetes     ClusterIP   10.43.0.1       <none>        443/TCP          123m

NAME                                         CLASS     HOSTS   ADDRESS                            PORTS   AGE
ingress.networking.k8s.io/crud-app-ingress   traefik   *       172.19.0.3,172.19.0.4,172.19.0.5   80      2m58s
$ curl http://localhost:8081/
Hello, World!
