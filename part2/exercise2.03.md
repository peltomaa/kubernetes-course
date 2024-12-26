```bash
$ kubectl get pods,ing,svc,pv,pvc -n log-ping
NAME                            READY   STATUS    RESTARTS   AGE
pod/log-output-fc58ccb9-fq495   0/2     Pending   0          2m10s
pod/pong-6cf4b88ccf-mqxkl       0/1     Pending   0          2m2s

NAME                                           CLASS     HOSTS   ADDRESS                            PORTS   AGE
ingress.networking.k8s.io/log-output-ingress   traefik   *       172.24.0.3,172.24.0.4,172.24.0.5   80      2m10s
ingress.networking.k8s.io/pong-ingress         traefik   *       172.24.0.3,172.24.0.4,172.24.0.5   80      2m2s
                                                                                                                                                                                            NAME                     TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
service/log-output-svc   NodePort   10.43.242.55    <none>        2345:32252/TCP   2m10s
service/pong-svc         NodePort   10.43.133.142   <none>        2346:31801/TCP   2m2s

NAME                              CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                       STORAGECLASS     VOLUMEATTRIBUTESCLASS   REASON   AGE
persistentvolume/kube-course-pv   1Gi        RWX            Retain           Bound    default/kube-course-claim   kube-course-pv   <unset>                          5h29m

NAME                                      STATUS    VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS     VOLUMEATTRIBUTESCLASS   AGE
persistentvolumeclaim/kube-course-claim   Pending                                      kube-course-pv   <unset>                 15s
```
