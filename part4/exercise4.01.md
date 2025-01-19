Final commit to adjust readiness probe: https://github.com/peltomaa/kubernetes-course/commit/59e0e151aed5705c85bc73b0033dbaef32adf03c

Commit to add `GET /healthz` endpoint: https://github.com/peltomaa/kubernetes-course/commit/7f13f2560cdefc4da17c0a5bdf5d26f92ba34559

```bash
➜  pong git:(main) ✗ kubectl describe pod pong-79d9df979b-jcl2s
...
    Readiness:  http-get http://:3000/healthz delay=60s timeout=1s period=60s #success=1 #failure=10
...
```
