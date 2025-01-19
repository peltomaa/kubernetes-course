## Note: I didn't get Rollout working

I cannot get agro rollouts working unfortunaltely.
I've still added `analysistemplate.yaml` and `rollout.yaml` for crud-backend. I'll revisit this topic later and get it working.

## Code:

- `analysistemplate.yaml`: https://github.com/peltomaa/kubernetes-course/blob/332562c90ab5b17adfdd0440682d072f45c78476/crud-infra/manifests/analysistemplate.yaml
- `rollout.yaml`: https://github.com/peltomaa/kubernetes-course/blob/332562c90ab5b17adfdd0440682d072f45c78476/crud-backend/manifests/rollout.yaml

## Rollout not working :/

Seems like rollout is stuck at message: "more replicas need to be updated".

```bash
Name:            crud-backend
Namespace:       crud
Status:          ◌ Progressing
Message:         more replicas need to be updated
Strategy:        Canary
  Step:
  SetWeight:     20
  ActualWeight:  0
Replicas:
  Desired:       1
  Current:       0
  Updated:       0
  Ready:         0
  Available:     0

NAME            KIND     STATUS         AGE    INFO
⟳ crud-backend  Rollout  ◌ Progressing  9m44s
```
