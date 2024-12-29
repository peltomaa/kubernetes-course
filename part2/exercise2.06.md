```bash
$ kubectl get pods,configmaps -n log-pong
NAME                              READY   STATUS    RESTARTS   AGE
pod/log-output-7bfc8b5cf7-mmgfb   2/2     Running   0          3m53s
pod/pong-66557478cc-bjnt2         1/1     Running   0          2m56s

NAME                             DATA   AGE
configmap/kube-root-ca.crt       1      35m
configmap/log-output-configmap   1      3m53s
$ curl http://localhost:8081/
file content: this text is from file
MESSAGE=hello world
2024-12-26T17:39:59Z: 904965bd-5ec3-44c7-9eb8-b1f93ac98df0
Ping / Pongs: 2
```
