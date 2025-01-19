## Commits:

- Adding readiness probe to crud-backend: https://github.com/peltomaa/kubernetes-course/commit/410c4566f898e3cab2a50fc69c3a1b6353beef48
- Adding `GET /healthz` endpoint to crud-backend: https://github.com/peltomaa/kubernetes-course/commit/55eae40c5526480ea9731d1b050138d6788a129f
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
