Final commit to adjust readiness probe: https://github.com/peltomaa/kubernetes-course/commit/59e0e151aed5705c85bc73b0033dbaef32adf03c

```bash
➜  pong git:(main) ✗ kubectl describe pod pong-79d9df979b-jcl2s
...
    Readiness:  http-get http://:3000/healthz delay=60s timeout=1s period=60s #success=1 #failure=10
...
```
