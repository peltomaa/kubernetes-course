## Stateful sets in prometheus

```bash
➜  ~ kubectl get po
NAME                                                              READY   STATUS    RESTARTS   AGE
alertmanager-kube-prometheus-stack-1737-alertmanager-0            2/2     Running   0          9m16s
kube-prometheus-stack-1737-operator-76b59c655f-pdl6m              1/1     Running   0          9m27s
kube-prometheus-stack-1737306715-grafana-7d74d75585-twgkv         3/3     Running   0          9m26s
kube-prometheus-stack-1737306715-kube-state-metrics-9c66c4ckl4z   1/1     Running   0          9m27s
kube-prometheus-stack-1737306715-prometheus-node-exporter-26jdp   1/1     Running   0          9m26s
kube-prometheus-stack-1737306715-prometheus-node-exporter-rtx4r   1/1     Running   0          9m26s
kube-prometheus-stack-1737306715-prometheus-node-exporter-tlqc7   1/1     Running   0          9m26s
prometheus-kube-prometheus-stack-1737-prometheus-0                2/2     Running   0          9m16s
➜  ~ kubectl get statefulsets
NAME                                                   READY   AGE
alertmanager-kube-prometheus-stack-1737-alertmanager   1/1     9m23s
prometheus-kube-prometheus-stack-1737-prometheus       1/1     9m23s
```

## Query

```promql
sum(kube_pod_info{namespace="prometheus", created_by_kind="StatefulSet"})
```

## Getting data

![Getting data in Prometheus query page](/part4/exercise4.03.png)
