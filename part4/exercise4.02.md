## Commits:

- Adding readiness probe to crud-backend: https://github.com/peltomaa/kubernetes-course/commit/410c4566f898e3cab2a50fc69c3a1b6353beef48
- Adding `GET /healthz` endpoint to crud-backend: https://github.com/peltomaa/kubernetes-course/commit/55eae40c5526480ea9731d1b050138d6788a129f
  - Improving the check: https://github.com/peltomaa/kubernetes-course/commit/6f084e745fdfeeae92d54de9de8b59526f7981c1
- Adding preadiness proble to crud-app index: https://github.com/peltomaa/kubernetes-course/commit/4e4a4a24d546165a5e8d0241be2ed8b8b242c64f

```bash
➜  kubernetes-course git:(main) ✗ kubectl describe pod crud-backend-7bc947689f-tkdnn
...
    Readiness:  http-get http://:3000/healthz delay=60s timeout=1s period=60s #success=1 #failure=10
...

➜  kubernetes-course git:(main) ✗ kubectl describe pod crud-app-6cc5bb7484-fbdhq
...
    Readiness:  http-get http://:3000/ delay=60s timeout=1s period=60s #success=1 #failure=10
...
```

## Testing failing

```bash
➜  crud-backend git:(main) ✗ kubectl get po
NAME                            READY   STATUS             RESTARTS       AGE
crud-app-6cc5bb7484-fbdhq       1/1     Running            0              5h51m
crud-backend-7c99879d4f-64pr4   0/1     CrashLoopBackOff   7 (2m9s ago)   12m
crud-postgres-stset-0           1/1     Running            0              6h11m
```

CRUD backend is not starting due to failing `GET /healthz`
