```bash
$ k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
INFO[0000] portmapping '8081:80' targets the loadbalancer: defaulting to [servers:*:proxy agents:*:proxy]
INFO[0000] Prep: Network
INFO[0000] Created network 'k3d-k3s-default'
INFO[0000] Created image volume k3d-k3s-default-images
INFO[0000] Starting new tools node...
INFO[0000] Starting node 'k3d-k3s-default-tools'
INFO[0001] Creating node 'k3d-k3s-default-server-0'
INFO[0002] Creating node 'k3d-k3s-default-agent-0'
INFO[0002] Creating node 'k3d-k3s-default-agent-1'
INFO[0002] Creating LoadBalancer 'k3d-k3s-default-serverlb'
INFO[0002] Using the k3d-tools node to gather environment information
INFO[0002] Starting new tools node...
INFO[0002] Starting node 'k3d-k3s-default-tools'
INFO[0004] Starting cluster 'k3s-default'
INFO[0004] Starting servers...
INFO[0004] Starting node 'k3d-k3s-default-server-0'
INFO[0010] Starting agents...
INFO[0010] Starting node 'k3d-k3s-default-agent-1'
INFO[0010] Starting node 'k3d-k3s-default-agent-0'
INFO[0016] Starting helpers...
INFO[0016] Starting node 'k3d-k3s-default-serverlb'
INFO[0024] Injecting records for hostAliases (incl. host.k3d.internal) and for 5 network members into CoreDNS configmap...
INFO[0026] Cluster 'k3s-default' created successfully!
INFO[0026] You can now use it like this:
kubectl cluster-info
$ kubectl apply -f manifests/deployment.yaml
deployment.apps/crud-app created
$ kubectl get pods
NAME                       READY   STATUS    RESTARTS   AGE
crud-app-9c89f7454-5784k   1/1     Running   0          3m13s
$ kubectl apply -f manifests/service.yaml
service/crud-app-svc created
$ kubectl get services
NAME           TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
crud-app-svc   NodePort    10.43.49.179   <none>        1234:30080/TCP   8s
kubernetes     ClusterIP   10.43.0.1      <none>        443/TCP          7m2s
$ curl http://localhost:8082/
Hello, World!
