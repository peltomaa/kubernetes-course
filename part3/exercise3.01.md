```bash
➜  kubernetes-course git:(main) ✗ kubectl get pods,
➜  kubernetes-course git:(main) ✗ kubectl cluster-info
Kubernetes control plane is running at https://34.88.8.169
GLBCDefaultBackend is running at https://34.88.8.169/api/v1/namespaces/kube-system/services/default-http-backend:http/proxy
KubeDNS is running at https://34.88.8.169/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
Metrics-server is running at https://34.88.8.169/api/v1/namespaces/kube-system/services/https:metrics-server:/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
➜  kubernetes-course git:(main) ✗ kubectl get pods,pvc,statefulset,services
NAME                        READY   STATUS    RESTARTS   AGE
pod/pong-b889858d4-bnt6p    1/1     Running   0          28m
pod/pong-postgres-stset-0   1/1     Running   0          13m

NAME                                                                     STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   VOLUMEATTRIBUTESCLASS   AGE
persistentvolumeclaim/log-pong-claim                                     Bound    pvc-c6fee384-73a4-41ee-a86e-4cc300e85a79   30Gi       RWO            standard-rwo   <unset>                 24m
persistentvolumeclaim/pong-postgres-data-storage-pong-postgres-stset-0   Bound    pvc-c590df40-2a55-4811-9467-843428f7cc81   1Gi        RWO            standard-rwo   <unset>                 16m

NAME                                   READY   AGE
statefulset.apps/pong-postgres-stset   1/1     13m

NAME                        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
service/pong-postgres-svc   ClusterIP   None            <none>        5432/TCP         23m
service/pong-svc            NodePort    34.118.231.85   <none>        2346:30728/TCP   28m
```
