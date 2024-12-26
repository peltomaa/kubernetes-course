```bash
$ kubectl get pods,ing,svc,pv,pvc -n crud
NAME                                READY   STATUS              RESTARTS   AGE
pod/crud-app-6677ff8dd8-jflfq       0/1     Pending             0          80s
pod/crud-backend-557f789665-v6rxz   0/1     ContainerCreating   0          74s
                                                                                                                                                                                            NAME                                             CLASS     HOSTS   ADDRESS                            PORTS   AGE
ingress.networking.k8s.io/crud-app-ingress       traefik   *       172.18.0.3,172.18.0.4,172.18.0.5   80      80s
ingress.networking.k8s.io/crud-backend-ingress   traefik   *       172.18.0.3,172.18.0.4,172.18.0.5   80      74s

NAME                       TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
service/crud-app-svc       NodePort   10.43.112.213   <none>        2347:31828/TCP   80s
service/crud-backend-svc   NodePort   10.43.35.78     <none>        2348:30592/TCP   74s

NAME                              CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                     STORAGECLASS     VOLUMEATTRIBUTESCLASS   REASON   AGE
persistentvolume/kube-course-pv   1Gi        RWX            Retain           Bound    log-pong/log-pong-claim   kube-course-pv   <unset>                          5m14s

NAME                               STATUS    VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS     VOLUMEATTRIBUTESCLASS   AGE
persistentvolumeclaim/crud-claim   Pending                                      kube-course-pv   <unset>                 115s
```
