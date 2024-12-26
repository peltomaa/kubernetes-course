```bash
$ kubectl get svc,ing,pods,pv,pvc
NAME                   TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
service/crud-app-svc   NodePort    10.43.63.49   <none>        2347:31887/TCP   2m58s
service/kubernetes     ClusterIP   10.43.0.1     <none>        443/TCP          102m

NAME                                         CLASS     HOSTS   ADDRESS                            PORTS   AGE
ingress.networking.k8s.io/crud-app-ingress   traefik   *       172.24.0.3,172.24.0.4,172.24.0.5   80      2m58s

NAME                           READY   STATUS    RESTARTS   AGE
pod/crud-app-c457457c5-fkdds   1/1     Running   0          2m58s

NAME                              CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                       STORAGECLASS     VOLUMEATTRIBUTESCLASS   REASON   AGE
persistentvolume/kube-course-pv   1Gi        RWX            Retain           Bound    default/kube-course-claim   kube-course-pv   <unset>                          102m

NAME                                      STATUS   VOLUME           CAPACITY   ACCESS MODES   STORAGECLASS     VOLUMEATTRIBUTESCLASS   AGE
persistentvolumeclaim/kube-course-claim   Bound    kube-course-pv   1Gi        RWX            kube-course-pv   <unset>                 102m
$ curl -I http://localhost:8081/img
HTTP/1.1 200 OK
Accept-Ranges: bytes
Content-Length: 115560
Content-Type: image/png
Date: Thu, 26 Dec 2024 12:40:34 GMT
Last-Modified: Thu, 26 Dec 2024 12:40:34 GMT
