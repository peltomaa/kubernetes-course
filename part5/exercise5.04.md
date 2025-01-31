## My wikipedia app running

```bash
➜  ~ kubectl get pod,services
NAME                READY   STATUS    RESTARTS   AGE
pod/wikipedia-app   2/2     Running   0          61s

NAME                            TYPE           CLUSTER-IP       EXTERNAL-IP      PORT(S)        AGE
service/wikipedia-app-service   LoadBalancer   34.118.234.171   35.228.124.218   80:31700/TCP   61s
➜  ~ curl http://35.228.124.218/
<!DOCTYPE html>
<html class="client-nojs vector-feature-language-in-header-enabled vector-feature-language-in-main-page-header-disabled vector-feature-page-tools-pinned-disabled vector-feature-toc-pinned-clientpref-1 vector-feature-main-menu-pinned-disabled vector-feature-limited-width-clientpref-1 vector-feature-limited-width-content-enabled vector-feature-custom-font-size-clientpref-1 vector-feature-appearance-pinned-clientpref-1 vector-feature-night-mode-enabled skin-theme-clientpref-day vector-sticky-header-enabled vector-toc-available" lang="en" dir="ltr">
<head>
...rest of HTML would be here...
```
