```bash
$ kubectl get svc pong-postgres-svc -n log-pong

NAME                TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE
pong-postgres-svc   ClusterIP   None         <none>        5432/TCP   12m
➜  kubernetes-cource git:(main) ✗ kubectl run -it --rm --restart=Never --image postgres psql-for-debugging sh
If you don't see a command prompt, try pressing enter.

# psql -h pong-postgres-svc -p 5432 -U postgres
Password for user postgres:
psql (17.2 (Debian 17.2-1.pgdg120+1), server 14.15 (Debian 14.15-1.pgdg120+1))
Type "help" for help.

postgres=#
```

## Pongs from the database

```bash
postgres=# SELECT COUNT(*) FROM pong_clicks;
 count
-------
    17
(1 row)

postgres=# SELECT COUNT(*) FROM pong_clicks;
 count
-------
    19
(1 row)

postgres=# exit;
root@psql-for-debugging:/# exit
exit
pod "psql-for-debugging" deleted
➜  ~ curl http://localhost:8081/pingpong
pong 19%
```
