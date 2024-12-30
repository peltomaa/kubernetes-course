```bash
➜  crud-backend git:(main) ✗ kubectl get pods
NAME                                 READY   STATUS      RESTARTS   AGE
crud-app-6677ff8dd8-gx9mf            1/1     Running     0          19h
crud-backend-748fbb4567-85dt2        1/1     Running     0          87s
crud-backend-worker-28926319-9cxfx   0/1     Completed   0          86s
crud-backend-worker-28926320-w5ch6   0/1     Completed   0          26s
crud-postgres-stset-0                1/1     Running     0          2m52s
psql-for-debugging                   1/1     Running     0          19h
➜  crud-backend git:(main) ✗ curl http://localhost:8081/todos
[{"id":1,"task":"New task","created_at":"2024-12-29T22:14:07.462535Z"},{"id":2,"task":"New task 2","created_at":"2024-12-29T22:17:29.752494Z"},{"id":3,"task":"Newer task","created_at":"2024-12-30T16:38:22.162671Z"},{"id":4,"task":"Task again","created_at":"2024-12-30T16:41:05.883521Z"},{"id":5,"task":"Read https://en.wikipedia.org/wiki/NKVD_Order_No._00439","created_at":"2024-12-30T17:19:09.154048Z"},{"id":6,"task":"Read https://en.wikipedia.org/wiki/Jacinta_Allan","created_at":"2024-12-30T17:20:05.451412Z"}]
➜  crud-backend git:(main) ✗ kubectl get cronjobs
NAME                  SCHEDULE    TIMEZONE   SUSPEND   ACTIVE   LAST SCHEDULE   AGE
crud-backend-worker   0 * * * *   <none>     False     0        3m6s            5m7s
```
