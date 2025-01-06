## Files

- https://github.com/peltomaa/kubernetes-course/blob/3ae4eea759e53cba27221bd73c1e0ffee14b2c39/.github/workflows/deploy.yaml

## Feature namespace running

```bash
➜  kubernetes-course git:(main) ✗ kubens crud-feat-add-crud-app-title
Context "gke_nomadic-charge-446815-j2_europe-north1-b_dwk-cluster" modified.
Active namespace is "crud-feat-add-crud-app-title".
➜  kubernetes-course git:(main) ✗ kubectl get pods,ing,pvc,services
NAME                              READY   STATUS    RESTARTS   AGE
pod/crud-app-68fc466f45-wnvmk     1/1     Running   0          81s
pod/crud-backend-8c9d6bb6-z9x24   1/1     Running   0          65s
pod/crud-postgres-stset-0         1/1     Running   0          3m28s
pod/pong-postgres-stset-0         1/1     Running   0          3m28s

NAME                                             CLASS    HOSTS   ADDRESS   PORTS   AGE
ingress.networking.k8s.io/crud-app-ingress       <none>   *                 80      81s
ingress.networking.k8s.io/crud-backend-ingress   <none>   *                 80      64s

NAME                                                                     STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   VOLUMEATTRIBUTESCLASS   AGE
persistentvolumeclaim/crud-claim                                         Bound     pvc-0278f556-03be-41b0-b4bf-4589ed1d6781   30Gi       RWO            standard-rwo   <unset>                 22m
persistentvolumeclaim/crud-postgres-data-storage-crud-postgres-stset-0   Bound     pvc-46bd78b0-cfae-49ad-b055-8eac7c903b07   1Gi        RWO            standard-rwo   <unset>                 22m
persistentvolumeclaim/log-pong-claim                                     Pending                                                                        standard-rwo   <unset>                 22m
persistentvolumeclaim/pong-postgres-data-storage-pong-postgres-stset-0   Bound     pvc-2bd239b3-0536-4661-bed9-bacafcd0ca93   1Gi        RWO            standard-rwo   <unset>                 22m

NAME                        TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
service/crud-app-svc        NodePort    34.118.230.151   <none>        2347:32468/TCP   81s
service/crud-backend-svc    NodePort    34.118.225.207   <none>        2348:30697/TCP   65s
service/crud-postgres-svc   ClusterIP   None             <none>        5432/TCP         85s
```

## Github action deployment logs

```bash
2025-01-06T09:49:01.9235485Z Current runner version: '2.321.0'
2025-01-06T09:49:01.9264968Z ##[group]Operating System
2025-01-06T09:49:01.9266130Z Ubuntu
2025-01-06T09:49:01.9266864Z 24.04.1
2025-01-06T09:49:01.9267706Z LTS
2025-01-06T09:49:01.9268453Z ##[endgroup]
2025-01-06T09:49:01.9269240Z ##[group]Runner Image
2025-01-06T09:49:01.9270396Z Image: ubuntu-24.04
2025-01-06T09:49:01.9271257Z Version: 20241215.1.0
2025-01-06T09:49:01.9272756Z Included Software: https://github.com/actions/runner-images/blob/ubuntu24/20241215.1/images/ubuntu/Ubuntu2404-Readme.md
2025-01-06T09:49:01.9274650Z Image Release: https://github.com/actions/runner-images/releases/tag/ubuntu24%2F20241215.1
2025-01-06T09:49:01.9276025Z ##[endgroup]
2025-01-06T09:49:01.9276819Z ##[group]Runner Image Provisioner
2025-01-06T09:49:01.9277770Z 2.0.404.1
2025-01-06T09:49:01.9278534Z ##[endgroup]
2025-01-06T09:49:01.9280449Z ##[group]GITHUB_TOKEN Permissions
2025-01-06T09:49:01.9282934Z Contents: read
2025-01-06T09:49:01.9283800Z Metadata: read
2025-01-06T09:49:01.9284802Z Packages: read
2025-01-06T09:49:01.9285732Z ##[endgroup]
2025-01-06T09:49:01.9289138Z Secret source: Actions
2025-01-06T09:49:01.9290725Z Prepare workflow directory
2025-01-06T09:49:01.9785317Z Prepare all required actions
2025-01-06T09:49:01.9822023Z Getting action download info
2025-01-06T09:49:02.1960023Z Download action repository 'actions/checkout@v4' (SHA:11bd71901bbe5b1630ceea73d27597364c9af683)
2025-01-06T09:49:02.3064411Z Download action repository 'google-github-actions/auth@v2' (SHA:6fc4af4b145ae7821d527454aa9bd537d1f2dc5f)
2025-01-06T09:49:02.5087251Z Download action repository 'google-github-actions/setup-gcloud@v2' (SHA:6189d56e4096ee891640bb02ac264be376592d6a)
2025-01-06T09:49:02.6803906Z Download action repository 'google-github-actions/get-gke-credentials@v2' (SHA:9025e8f90f2d8e0c3dafc3128cc705a26d992a6a)
2025-01-06T09:49:03.0003553Z Complete job name: build
2025-01-06T09:49:03.1059238Z ##[group]Run actions/checkout@v4
2025-01-06T09:49:03.1061362Z with:
2025-01-06T09:49:03.1062550Z   repository: peltomaa/kubernetes-course
2025-01-06T09:49:03.1064460Z   token: ***
2025-01-06T09:49:03.1065576Z   ssh-strict: true
2025-01-06T09:49:03.1066760Z   ssh-user: git
2025-01-06T09:49:03.1067964Z   persist-credentials: true
2025-01-06T09:49:03.1069587Z   clean: true
2025-01-06T09:49:03.1070802Z   sparse-checkout-cone-mode: true
2025-01-06T09:49:03.1072257Z   fetch-depth: 1
2025-01-06T09:49:03.1073414Z   fetch-tags: false
2025-01-06T09:49:03.1074605Z   show-progress: true
2025-01-06T09:49:03.1075816Z   lfs: false
2025-01-06T09:49:03.1076910Z   submodules: false
2025-01-06T09:49:03.1078116Z   set-safe-directory: true
2025-01-06T09:49:03.1079866Z env:
2025-01-06T09:49:03.1081050Z   PROJECT_ID: ***
2025-01-06T09:49:03.1082239Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:49:03.1083583Z   GKE_ZONE: europe-north1-b
2025-01-06T09:49:03.1084955Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:49:03.1087359Z   SOPS_AGE_KEY: ***
2025-01-06T09:49:03.1088534Z ##[endgroup]
2025-01-06T09:49:03.3363975Z Syncing repository: peltomaa/kubernetes-course
2025-01-06T09:49:03.3367387Z ##[group]Getting Git version info
2025-01-06T09:49:03.3369243Z Working directory is '/home/runner/work/kubernetes-course/kubernetes-course'
2025-01-06T09:49:03.3387709Z [command]/usr/bin/git version
2025-01-06T09:49:03.3390142Z git version 2.47.1
2025-01-06T09:49:03.3430066Z ##[endgroup]
2025-01-06T09:49:03.3441452Z Temporarily overriding HOME='/home/runner/work/_temp/5d79ec59-ee16-42bc-9de3-76d2eb5ae875' before making global git config changes
2025-01-06T09:49:03.3446763Z Adding repository directory to the temporary git global config as a safe directory
2025-01-06T09:49:03.3451339Z [command]/usr/bin/git config --global --add safe.directory /home/runner/work/kubernetes-course/kubernetes-course
2025-01-06T09:49:03.3519892Z Deleting the contents of '/home/runner/work/kubernetes-course/kubernetes-course'
2025-01-06T09:49:03.3524063Z ##[group]Initializing the repository
2025-01-06T09:49:03.3527023Z [command]/usr/bin/git init /home/runner/work/kubernetes-course/kubernetes-course
2025-01-06T09:49:03.3588088Z hint: Using 'master' as the name for the initial branch. This default branch name
2025-01-06T09:49:03.3593571Z hint: is subject to change. To configure the initial branch name to use in all
2025-01-06T09:49:03.3599035Z hint: of your new repositories, which will suppress this warning, call:
2025-01-06T09:49:03.3602076Z hint:
2025-01-06T09:49:03.3604148Z hint: 	git config --global init.defaultBranch <name>
2025-01-06T09:49:03.3606841Z hint:
2025-01-06T09:49:03.3609247Z hint: Names commonly chosen instead of 'master' are 'main', 'trunk' and
2025-01-06T09:49:03.3612888Z hint: 'development'. The just-created branch can be renamed via this command:
2025-01-06T09:49:03.3615570Z hint:
2025-01-06T09:49:03.3617263Z hint: 	git branch -m <name>
2025-01-06T09:49:03.3620742Z Initialized empty Git repository in /home/runner/work/kubernetes-course/kubernetes-course/.git/
2025-01-06T09:49:03.3627864Z [command]/usr/bin/git remote add origin https://github.com/peltomaa/kubernetes-course
2025-01-06T09:49:03.3656399Z ##[endgroup]
2025-01-06T09:49:03.3659756Z ##[group]Disabling automatic garbage collection
2025-01-06T09:49:03.3662527Z [command]/usr/bin/git config --local gc.auto 0
2025-01-06T09:49:03.3694263Z ##[endgroup]
2025-01-06T09:49:03.3697245Z ##[group]Setting up auth
2025-01-06T09:49:03.3728414Z [command]/usr/bin/git config --local --name-only --get-regexp core\.sshCommand
2025-01-06T09:49:03.3745200Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'core\.sshCommand' && git config --local --unset-all 'core.sshCommand' || :"
2025-01-06T09:49:03.4050333Z [command]/usr/bin/git config --local --name-only --get-regexp http\.https\:\/\/github\.com\/\.extraheader
2025-01-06T09:49:03.4084199Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'http\.https\:\/\/github\.com\/\.extraheader' && git config --local --unset-all 'http.https://github.com/.extraheader' || :"
2025-01-06T09:49:03.4318873Z [command]/usr/bin/git config --local http.https://github.com/.extraheader AUTHORIZATION: basic ***
2025-01-06T09:49:03.4362758Z ##[endgroup]
2025-01-06T09:49:03.4367666Z ##[group]Fetching the repository
2025-01-06T09:49:03.4386450Z [command]/usr/bin/git -c protocol.version=2 fetch --no-tags --prune --no-recurse-submodules --depth=1 origin +3ae4eea759e53cba27221bd73c1e0ffee14b2c39:refs/remotes/origin/feat-add-crud-app-title
2025-01-06T09:49:03.7736388Z From https://github.com/peltomaa/kubernetes-course
2025-01-06T09:49:03.7738992Z  * [new ref]         3ae4eea759e53cba27221bd73c1e0ffee14b2c39 -> origin/feat-add-crud-app-title
2025-01-06T09:49:03.7785113Z ##[endgroup]
2025-01-06T09:49:03.7788046Z ##[group]Determining the checkout info
2025-01-06T09:49:03.7793013Z ##[endgroup]
2025-01-06T09:49:03.7794893Z [command]/usr/bin/git sparse-checkout disable
2025-01-06T09:49:03.7835395Z [command]/usr/bin/git config --local --unset-all extensions.worktreeConfig
2025-01-06T09:49:03.7869777Z ##[group]Checking out the ref
2025-01-06T09:49:03.7874974Z [command]/usr/bin/git checkout --progress --force -B feat-add-crud-app-title refs/remotes/origin/feat-add-crud-app-title
2025-01-06T09:49:03.8079665Z Switched to a new branch 'feat-add-crud-app-title'
2025-01-06T09:49:03.8084372Z branch 'feat-add-crud-app-title' set up to track 'origin/feat-add-crud-app-title'.
2025-01-06T09:49:03.8091076Z ##[endgroup]
2025-01-06T09:49:03.8115154Z [command]/usr/bin/git log -1 --format=%H
2025-01-06T09:49:03.8138011Z 3ae4eea759e53cba27221bd73c1e0ffee14b2c39
2025-01-06T09:49:03.8467686Z ##[group]Run google-github-actions/auth@v2
2025-01-06T09:49:03.8469066Z with:
2025-01-06T09:49:03.8503503Z   credentials_json: ***
2025-01-06T09:49:03.8504674Z   create_credentials_file: true
2025-01-06T09:49:03.8505910Z   export_environment_variables: true
2025-01-06T09:49:03.8507196Z   universe: googleapis.com
2025-01-06T09:49:03.8508362Z   cleanup_credentials: true
2025-01-06T09:49:03.8509658Z   access_token_lifetime: 3600s
2025-01-06T09:49:03.8511204Z   access_token_scopes: https://www.googleapis.com/auth/cloud-platform
2025-01-06T09:49:03.8513064Z   id_token_include_email: false
2025-01-06T09:49:03.8514232Z env:
2025-01-06T09:49:03.8515279Z   PROJECT_ID: ***
2025-01-06T09:49:03.8516342Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:49:03.8517502Z   GKE_ZONE: europe-north1-b
2025-01-06T09:49:03.8518690Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:49:03.8520863Z   SOPS_AGE_KEY: ***
2025-01-06T09:49:03.8521945Z ##[endgroup]
2025-01-06T09:49:03.9588827Z Created credentials file at "/home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json"
2025-01-06T09:49:03.9765926Z ##[group]Run google-github-actions/setup-gcloud@v2
2025-01-06T09:49:03.9768323Z with:
2025-01-06T09:49:03.9770343Z   install_components: kubectl
2025-01-06T09:49:03.9772421Z   version: latest
2025-01-06T09:49:03.9774281Z   skip_install: false
2025-01-06T09:49:03.9776160Z env:
2025-01-06T09:49:03.9778021Z   PROJECT_ID: ***
2025-01-06T09:49:03.9780165Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:49:03.9782251Z   GKE_ZONE: europe-north1-b
2025-01-06T09:49:03.9784356Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:49:03.9787774Z   SOPS_AGE_KEY: ***
2025-01-06T09:49:03.9791721Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:49:03.9826430Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:49:03.9832175Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:49:03.9836299Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:49:03.9838662Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:49:03.9841053Z   GCLOUD_PROJECT: ***
2025-01-06T09:49:03.9843179Z   GCP_PROJECT: ***
2025-01-06T09:49:03.9845345Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:49:03.9847365Z ##[endgroup]
2025-01-06T09:49:05.6342247Z [command]/usr/bin/tar xz --warning=no-unknown-keyword --overwrite -C /home/runner/work/_temp/5b22fc04-a4bb-49ab-a2e5-3ef883a68a83 -f /home/runner/work/_temp/ca9f0657-cb7d-4f14-959f-5d5a8d7781f7
2025-01-06T09:50:01.7556495Z Successfully authenticated
2025-01-06T09:50:01.7768754Z ##[group]Run google-github-actions/get-gke-credentials@v2
2025-01-06T09:50:01.7769194Z with:
2025-01-06T09:50:01.7769819Z   cluster_name: dwk-cluster
2025-01-06T09:50:01.7770192Z   project_id: ***
2025-01-06T09:50:01.7770485Z   location: europe-north1-b
2025-01-06T09:50:01.7770785Z   use_auth_provider: false
2025-01-06T09:50:01.7771072Z   use_internal_ip: false
2025-01-06T09:50:01.7771362Z   use_connect_gateway: false
2025-01-06T09:50:01.7771666Z   use_dns_based_endpoint: false
2025-01-06T09:50:01.7771956Z env:
2025-01-06T09:50:01.7772227Z   PROJECT_ID: ***
2025-01-06T09:50:01.7772506Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:01.7772822Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:01.7773128Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:01.7773665Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:01.7774269Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:01.7775197Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:01.7776026Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:01.7776645Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:01.7776988Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:01.7777305Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:01.7777605Z   GCP_PROJECT: ***
2025-01-06T09:50:01.7777920Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:01.7778296Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:01.7778717Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:01.7779054Z ##[endgroup]
2025-01-06T09:50:02.7391811Z Successfully created and exported "KUBECONFIG" at: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:02.7487962Z ##[group]Run gcloud --quiet auth configure-docker europe-north1-docker.pkg.dev
2025-01-06T09:50:02.7488698Z [36;1mgcloud --quiet auth configure-docker europe-north1-docker.pkg.dev[0m
2025-01-06T09:50:02.7519196Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:50:02.7519852Z env:
2025-01-06T09:50:02.7520172Z   PROJECT_ID: ***
2025-01-06T09:50:02.7520463Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:02.7520771Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:02.7521081Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:02.7521614Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:02.7522196Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:02.7523094Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:02.7523946Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:02.7524548Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:02.7524886Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:02.7525195Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:02.7525514Z   GCP_PROJECT: ***
2025-01-06T09:50:02.7525837Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:02.7526210Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:02.7526633Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:02.7527198Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:02.7527951Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:02.7528479Z ##[endgroup]
2025-01-06T09:50:03.1963659Z Adding credentials for: europe-north1-docker.pkg.dev
2025-01-06T09:50:03.2233073Z Docker configuration file updated.
2025-01-06T09:50:03.3188907Z ##[group]Run curl -O -L -C - https://github.com/mozilla/sops/releases/download/v3.7.3/sops-v3.7.3.linux
2025-01-06T09:50:03.3190426Z [36;1mcurl -O -L -C - https://github.com/mozilla/sops/releases/download/v3.7.3/sops-v3.7.3.linux[0m
2025-01-06T09:50:03.3191032Z [36;1msudo mv sops-v3.7.3.linux /usr/bin/sops[0m
2025-01-06T09:50:03.3191820Z [36;1msudo chmod +x /usr/bin/sops[0m
2025-01-06T09:50:03.3192611Z [36;1msops --decrypt crud-db/manifests/secrets.enc.yaml > crud-db/manifests/secrets.yaml[0m
2025-01-06T09:50:03.3219853Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:50:03.3220190Z env:
2025-01-06T09:50:03.3220469Z   PROJECT_ID: ***
2025-01-06T09:50:03.3220749Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:03.3221057Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:03.3221370Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:03.3221903Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:03.3222536Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:03.3223462Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:03.3224293Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:03.3224886Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:03.3225263Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:03.3225575Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:03.3225868Z   GCP_PROJECT: ***
2025-01-06T09:50:03.3226182Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:03.3226545Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:03.3226965Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:03.3227523Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:03.3228270Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:03.3228802Z ##[endgroup]
2025-01-06T09:50:03.3345335Z   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
2025-01-06T09:50:03.3346589Z                                  Dload  Upload   Total   Spent    Left  Speed
2025-01-06T09:50:03.3347031Z
2025-01-06T09:50:03.4189243Z   0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
2025-01-06T09:50:03.4190575Z   0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
2025-01-06T09:50:03.4982553Z
2025-01-06T09:50:03.4984247Z   0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
2025-01-06T09:50:03.6534155Z
2025-01-06T09:50:03.6535142Z 100 27.7M  100 27.7M    0     0  87.1M      0 --:--:-- --:--:-- --:--:-- 87.1M
2025-01-06T09:50:03.6895828Z ##[group]Run docker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-app:$BRANCH-$GITHUB_SHA crud-app/
2025-01-06T09:50:03.6896997Z [36;1mdocker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-app:$BRANCH-$GITHUB_SHA crud-app/[0m
2025-01-06T09:50:03.6925409Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:50:03.6925998Z env:
2025-01-06T09:50:03.6926432Z   PROJECT_ID: ***
2025-01-06T09:50:03.6926918Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:03.6927323Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:03.6945486Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:03.6946468Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:03.6947440Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:03.6948869Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:03.6950504Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:03.6951574Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:03.6952137Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:03.6952662Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:03.6953180Z   GCP_PROJECT: ***
2025-01-06T09:50:03.6953721Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:03.6954358Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:03.6955076Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:03.6956068Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:03.6957409Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:03.6958351Z ##[endgroup]
2025-01-06T09:50:04.2322193Z #0 building with "default" instance using docker driver
2025-01-06T09:50:04.2334798Z
2025-01-06T09:50:04.2335218Z #1 [internal] load build definition from Dockerfile
2025-01-06T09:50:04.2336228Z #1 transferring dockerfile: 130B done
2025-01-06T09:50:04.2338517Z #1 DONE 0.0s
2025-01-06T09:50:04.2339857Z
2025-01-06T09:50:04.2340957Z #2 [auth] library/golang:pull token for registry-1.docker.io
2025-01-06T09:50:04.2342328Z #2 DONE 0.0s
2025-01-06T09:50:04.2343431Z
2025-01-06T09:50:04.2344003Z #3 [internal] load metadata for docker.io/library/golang:1.23
2025-01-06T09:50:04.3611227Z #3 DONE 0.3s
2025-01-06T09:50:04.4934792Z
2025-01-06T09:50:04.4944703Z #4 [internal] load .dockerignore
2025-01-06T09:50:04.4955446Z #4 transferring context: 2B done
2025-01-06T09:50:04.4972010Z #4 DONE 0.0s
2025-01-06T09:50:04.4972738Z
2025-01-06T09:50:04.4973095Z #5 [internal] load build context
2025-01-06T09:50:04.4974054Z #5 transferring context: 6.04kB 0.0s done
2025-01-06T09:50:04.4974912Z #5 DONE 0.0s
2025-01-06T09:50:04.4993782Z
2025-01-06T09:50:04.4995597Z #6 [1/3] FROM docker.io/library/golang:1.23@sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041
2025-01-06T09:50:04.4998499Z #6 resolve docker.io/library/golang:1.23@sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041 0.0s done
2025-01-06T09:50:04.5001665Z #6 sha256:0a96bdb8280554b560ffee0f2e5f9843dc7b625f28192021ee103ecbcc2d629b 3.15MB / 48.50MB 0.1s
2025-01-06T09:50:04.5004455Z #6 sha256:54c7be425079efba0003054ee884bf72f1ffebca733bedd6f077d2809ee9aa6f 0B / 23.87MB 0.1s
2025-01-06T09:50:04.5006903Z #6 sha256:aff7bd25e6a162e9db0a284663d6aff04d456416cb3cc94d692a89be72b0e605 2.32kB / 2.32kB done
2025-01-06T09:50:04.5025358Z #6 sha256:7012e31470cb237fd56d72b6a7d16892ea12e7f8fd361be9010444423f28c821 2.80kB / 2.80kB done
2025-01-06T09:50:04.5051373Z #6 sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041 9.74kB / 9.74kB done
2025-01-06T09:50:04.5054660Z #6 sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 0B / 64.39MB 0.1s
2025-01-06T09:50:04.6925687Z #6 sha256:0a96bdb8280554b560ffee0f2e5f9843dc7b625f28192021ee103ecbcc2d629b 44.04MB / 48.50MB 0.3s
2025-01-06T09:50:04.6927395Z #6 sha256:54c7be425079efba0003054ee884bf72f1ffebca733bedd6f077d2809ee9aa6f 23.87MB / 23.87MB 0.3s
2025-01-06T09:50:04.6928672Z #6 sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 35.65MB / 64.39MB 0.3s
2025-01-06T09:50:04.7954990Z #6 sha256:0a96bdb8280554b560ffee0f2e5f9843dc7b625f28192021ee103ecbcc2d629b 48.50MB / 48.50MB 0.4s done
2025-01-06T09:50:04.7963310Z #6 sha256:54c7be425079efba0003054ee884bf72f1ffebca733bedd6f077d2809ee9aa6f 23.87MB / 23.87MB 0.3s done
2025-01-06T09:50:04.7968717Z #6 sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 45.09MB / 64.39MB 0.4s
2025-01-06T09:50:04.7971594Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 0B / 92.31MB 0.4s
2025-01-06T09:50:04.7973223Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 0B / 74.05MB 0.4s
2025-01-06T09:50:04.9185422Z #6 sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 53.48MB / 64.39MB 0.5s
2025-01-06T09:50:04.9201252Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 8.39MB / 92.31MB 0.5s
2025-01-06T09:50:04.9202447Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 7.34MB / 74.05MB 0.5s
2025-01-06T09:50:04.9203615Z #6 extracting sha256:0a96bdb8280554b560ffee0f2e5f9843dc7b625f28192021ee103ecbcc2d629b
2025-01-06T09:50:05.0311447Z #6 sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 64.39MB / 64.39MB 0.6s
2025-01-06T09:50:05.0329693Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 25.17MB / 92.31MB 0.6s
2025-01-06T09:50:05.0331041Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 20.97MB / 74.05MB 0.6s
2025-01-06T09:50:05.1461865Z #6 sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 64.39MB / 64.39MB 0.7s done
2025-01-06T09:50:05.1464625Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 40.89MB / 92.31MB 0.7s
2025-01-06T09:50:05.1465933Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 38.80MB / 74.05MB 0.7s
2025-01-06T09:50:05.2607769Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 58.63MB / 92.31MB 0.8s
2025-01-06T09:50:05.2621449Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 58.72MB / 74.05MB 0.8s
2025-01-06T09:50:05.2625857Z #6 sha256:3fd114183f3282d111ed7eaa48e1f94ff3018db89a43f47239fed2180f2d1084 125B / 125B 0.8s done
2025-01-06T09:50:05.2627365Z #6 sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1 0B / 32B 0.8s
2025-01-06T09:50:05.3803744Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 73.40MB / 92.31MB 0.9s
2025-01-06T09:50:05.3808014Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 72.48MB / 74.05MB 0.9s
2025-01-06T09:50:05.3810347Z #6 sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1 32B / 32B 0.9s
2025-01-06T09:50:05.4896560Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 92.31MB / 92.31MB 1.1s
2025-01-06T09:50:05.5931716Z #6 sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1 32B / 32B 1.2s done
2025-01-06T09:50:05.7075891Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 74.05MB / 74.05MB 1.3s done
2025-01-06T09:50:06.0532236Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 92.31MB / 92.31MB 1.5s done
2025-01-06T09:50:06.9363869Z #6 extracting sha256:0a96bdb8280554b560ffee0f2e5f9843dc7b625f28192021ee103ecbcc2d629b 2.0s done
2025-01-06T09:50:07.8169192Z #6 extracting sha256:54c7be425079efba0003054ee884bf72f1ffebca733bedd6f077d2809ee9aa6f
2025-01-06T09:50:08.4687447Z #6 extracting sha256:54c7be425079efba0003054ee884bf72f1ffebca733bedd6f077d2809ee9aa6f 0.5s done
2025-01-06T09:50:10.4910220Z #6 extracting sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1
2025-01-06T09:50:12.8061355Z #6 extracting sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 2.2s done
2025-01-06T09:50:13.4020844Z #6 extracting sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4
2025-01-06T09:50:15.4272185Z #6 extracting sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 2.0s done
2025-01-06T09:50:15.5448908Z #6 extracting sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c
2025-01-06T09:50:19.8304021Z #6 extracting sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 4.3s done
2025-01-06T09:50:20.5609324Z #6 extracting sha256:3fd114183f3282d111ed7eaa48e1f94ff3018db89a43f47239fed2180f2d1084
2025-01-06T09:50:20.7944218Z #6 extracting sha256:3fd114183f3282d111ed7eaa48e1f94ff3018db89a43f47239fed2180f2d1084 done
2025-01-06T09:50:20.7945517Z #6 extracting sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1 done
2025-01-06T09:50:20.7946385Z #6 DONE 16.2s
2025-01-06T09:50:20.7947778Z
2025-01-06T09:50:20.7948147Z #7 [2/3] WORKDIR /app
2025-01-06T09:50:20.7949876Z #7 DONE 0.0s
2025-01-06T09:50:20.7953123Z
2025-01-06T09:50:20.7955672Z #8 [3/3] COPY . .
2025-01-06T09:50:20.7957287Z #8 DONE 0.0s
2025-01-06T09:50:20.7957633Z
2025-01-06T09:50:20.7957834Z #9 exporting to image
2025-01-06T09:50:20.7958308Z #9 exporting layers
2025-01-06T09:50:22.4790885Z #9 exporting layers 1.8s done
2025-01-06T09:50:22.4793044Z #9 writing image sha256:66d5493133e984574f2dc249917bcfd09b699772ef04f17a8ad2893a38bddf1e
2025-01-06T09:50:22.5024684Z #9 writing image sha256:66d5493133e984574f2dc249917bcfd09b699772ef04f17a8ad2893a38bddf1e done
2025-01-06T09:50:22.5033750Z #9 naming to europe-north1-docker.pkg.dev/***/kubernetes-course/crud-app:feat-add-crud-app-title-3ae4eea759e53cba27221bd73c1e0ffee14b2c39 done
2025-01-06T09:50:22.5035431Z #9 DONE 1.8s
2025-01-06T09:50:22.5110894Z ##[group]Run docker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend:$BRANCH-$GITHUB_SHA crud-backend/
2025-01-06T09:50:22.5111956Z [36;1mdocker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend:$BRANCH-$GITHUB_SHA crud-backend/[0m
2025-01-06T09:50:22.5138914Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:50:22.5139248Z env:
2025-01-06T09:50:22.5139797Z   PROJECT_ID: ***
2025-01-06T09:50:22.5140097Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:22.5140407Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:22.5140738Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:22.5141398Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:22.5141996Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:22.5142885Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:22.5143756Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:22.5144352Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:22.5144690Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:22.5144993Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:22.5145282Z   GCP_PROJECT: ***
2025-01-06T09:50:22.5145587Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:22.5145961Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:22.5146398Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:22.5146941Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:22.5147892Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:22.5148424Z ##[endgroup]
2025-01-06T09:50:22.7588442Z #0 building with "default" instance using docker driver
2025-01-06T09:50:22.7590949Z
2025-01-06T09:50:22.7591323Z #1 [internal] load build definition from Dockerfile
2025-01-06T09:50:22.7592022Z #1 transferring dockerfile: 189B done
2025-01-06T09:50:22.7592578Z #1 DONE 0.0s
2025-01-06T09:50:22.7592819Z
2025-01-06T09:50:22.7593157Z #2 [internal] load metadata for docker.io/library/golang:1.23
2025-01-06T09:50:22.7593849Z #2 DONE 0.1s
2025-01-06T09:50:22.7594102Z
2025-01-06T09:50:22.7594327Z #3 [internal] load .dockerignore
2025-01-06T09:50:22.7594905Z #3 transferring context: 2B done
2025-01-06T09:50:22.7595443Z #3 DONE 0.0s
2025-01-06T09:50:22.7595698Z
2025-01-06T09:50:22.7596331Z #4 [1/5] FROM docker.io/library/golang:1.23@sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041
2025-01-06T09:50:22.7597285Z #4 DONE 0.0s
2025-01-06T09:50:22.7597543Z
2025-01-06T09:50:22.7597748Z #5 [2/5] WORKDIR /app
2025-01-06T09:50:22.7598234Z #5 CACHED
2025-01-06T09:50:22.7598400Z
2025-01-06T09:50:22.7598527Z #6 [internal] load build context
2025-01-06T09:50:22.7598854Z #6 transferring context: 9.51kB done
2025-01-06T09:50:22.7599280Z #6 DONE 0.0s
2025-01-06T09:50:22.7599694Z
2025-01-06T09:50:22.7599907Z #7 [3/5] COPY go.mod go.sum ./
2025-01-06T09:50:22.7600419Z #7 DONE 0.0s
2025-01-06T09:50:22.9117429Z
2025-01-06T09:50:22.9117731Z #8 [4/5] RUN go mod download
2025-01-06T09:50:23.0256564Z #8 DONE 0.3s
2025-01-06T09:50:23.2045629Z
2025-01-06T09:50:23.2061142Z #9 [5/5] COPY . .
2025-01-06T09:50:23.2061606Z #9 DONE 0.0s
2025-01-06T09:50:23.2061844Z
2025-01-06T09:50:23.2062037Z #10 exporting to image
2025-01-06T09:50:23.2062494Z #10 exporting layers
2025-01-06T09:50:25.6130876Z #10 exporting layers 2.6s done
2025-01-06T09:50:25.6355291Z #10 writing image sha256:270834e83eaabc4adbbe36571c893ccf661d71d351b2a14a5efad2248d384b9e done
2025-01-06T09:50:25.6357363Z #10 naming to europe-north1-docker.pkg.dev/***/kubernetes-course/crud-backend:feat-add-crud-app-title-3ae4eea759e53cba27221bd73c1e0ffee14b2c39 done
2025-01-06T09:50:25.6358807Z #10 DONE 2.6s
2025-01-06T09:50:25.6444777Z ##[group]Run docker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend-worker:$BRANCH-$GITHUB_SHA -f crud-backend/Dockerfile-worker crud-backend/
2025-01-06T09:50:25.6447052Z [36;1mdocker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend-worker:$BRANCH-$GITHUB_SHA -f crud-backend/Dockerfile-worker crud-backend/[0m
2025-01-06T09:50:25.6482187Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:50:25.6482525Z env:
2025-01-06T09:50:25.6482810Z   PROJECT_ID: ***
2025-01-06T09:50:25.6483097Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:25.6483403Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:25.6483711Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:25.6484260Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:25.6484890Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:25.6485829Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:25.6486656Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:25.6487267Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:25.6487609Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:25.6487927Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:25.6488232Z   GCP_PROJECT: ***
2025-01-06T09:50:25.6488539Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:25.6488910Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:25.6489481Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:25.6490099Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:25.6491084Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:25.6491617Z ##[endgroup]
2025-01-06T09:50:25.9093119Z #0 building with "default" instance using docker driver
2025-01-06T09:50:25.9094457Z
2025-01-06T09:50:25.9095039Z #1 [internal] load build definition from Dockerfile-worker
2025-01-06T09:50:25.9096268Z #1 transferring dockerfile: 196B done
2025-01-06T09:50:25.9097277Z #1 DONE 0.0s
2025-01-06T09:50:25.9101449Z
2025-01-06T09:50:25.9101835Z #2 [internal] load metadata for docker.io/library/golang:1.23
2025-01-06T09:50:25.9102528Z #2 DONE 0.1s
2025-01-06T09:50:25.9102785Z
2025-01-06T09:50:25.9103023Z #3 [internal] load .dockerignore
2025-01-06T09:50:25.9103588Z #3 transferring context: 2B done
2025-01-06T09:50:25.9104118Z #3 DONE 0.0s
2025-01-06T09:50:25.9104365Z
2025-01-06T09:50:25.9105018Z #4 [1/5] FROM docker.io/library/golang:1.23@sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041
2025-01-06T09:50:25.9105999Z #4 DONE 0.0s
2025-01-06T09:50:25.9106233Z
2025-01-06T09:50:25.9106435Z #5 [internal] load build context
2025-01-06T09:50:25.9106977Z #5 transferring context: 1.00kB done
2025-01-06T09:50:25.9107500Z #5 DONE 0.0s
2025-01-06T09:50:25.9107761Z
2025-01-06T09:50:25.9107959Z #6 [4/5] RUN go mod download
2025-01-06T09:50:25.9108457Z #6 CACHED
2025-01-06T09:50:25.9108678Z
2025-01-06T09:50:25.9110517Z #7 [2/5] WORKDIR /app
2025-01-06T09:50:25.9111105Z #7 CACHED
2025-01-06T09:50:25.9111333Z
2025-01-06T09:50:25.9111547Z #8 [3/5] COPY go.mod go.sum ./
2025-01-06T09:50:25.9112057Z #8 CACHED
2025-01-06T09:50:25.9112297Z
2025-01-06T09:50:25.9112483Z #9 [5/5] COPY . .
2025-01-06T09:50:25.9112925Z #9 CACHED
2025-01-06T09:50:25.9113169Z
2025-01-06T09:50:25.9113365Z #10 exporting to image
2025-01-06T09:50:25.9113819Z #10 exporting layers done
2025-01-06T09:50:25.9114618Z #10 writing image sha256:a977d0aff095504d2f49ce6afbe46f29a837399041d7998e68bbd2deb056bef3 done
2025-01-06T09:50:25.9116486Z #10 naming to europe-north1-docker.pkg.dev/***/kubernetes-course/crud-backend-worker:feat-add-crud-app-title-3ae4eea759e53cba27221bd73c1e0ffee14b2c39 done
2025-01-06T09:50:25.9117709Z #10 DONE 0.0s
2025-01-06T09:50:25.9198776Z ##[group]Run docker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-app:$BRANCH-$GITHUB_SHA
2025-01-06T09:50:25.9200681Z [36;1mdocker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-app:$BRANCH-$GITHUB_SHA[0m
2025-01-06T09:50:25.9250034Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:50:25.9250875Z env:
2025-01-06T09:50:25.9251391Z   PROJECT_ID: ***
2025-01-06T09:50:25.9251900Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:25.9252442Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:25.9253009Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:25.9253993Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:25.9255035Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:25.9256719Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:25.9258300Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:25.9259707Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:25.9260339Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:25.9260893Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:25.9270299Z   GCP_PROJECT: ***
2025-01-06T09:50:25.9270922Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:25.9271607Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:25.9272409Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:25.9273469Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:25.9274854Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:25.9276152Z ##[endgroup]
2025-01-06T09:50:26.6700372Z The push refers to repository [europe-north1-docker.pkg.dev/***/kubernetes-course/crud-app]
2025-01-06T09:50:27.3362488Z 72bed4b6594b: Preparing
2025-01-06T09:50:27.3363008Z 820e3f47fc89: Preparing
2025-01-06T09:50:27.3363440Z 5f70bf18a086: Preparing
2025-01-06T09:50:27.3363847Z 4e061853116c: Preparing
2025-01-06T09:50:27.3364314Z f3fe8ef9af87: Preparing
2025-01-06T09:50:27.3364926Z c3c4291cf0a2: Preparing
2025-01-06T09:50:27.3365414Z 85c6f0cfb532: Preparing
2025-01-06T09:50:27.3365895Z a4fd1e7df47e: Preparing
2025-01-06T09:50:27.3366391Z 2f7b6d216a37: Preparing
2025-01-06T09:50:27.3366868Z c3c4291cf0a2: Waiting
2025-01-06T09:50:27.3367337Z 85c6f0cfb532: Waiting
2025-01-06T09:50:27.3367794Z a4fd1e7df47e: Waiting
2025-01-06T09:50:27.3368254Z 2f7b6d216a37: Waiting
2025-01-06T09:50:28.2271496Z f3fe8ef9af87: Layer already exists
2025-01-06T09:50:28.7396695Z 5f70bf18a086: Layer already exists
2025-01-06T09:50:28.7532072Z 4e061853116c: Layer already exists
2025-01-06T09:50:28.9185880Z c3c4291cf0a2: Layer already exists
2025-01-06T09:50:29.4082349Z 85c6f0cfb532: Layer already exists
2025-01-06T09:50:29.4394049Z a4fd1e7df47e: Layer already exists
2025-01-06T09:50:29.6017321Z 2f7b6d216a37: Layer already exists
2025-01-06T09:50:31.5474546Z 820e3f47fc89: Pushed
2025-01-06T09:50:32.0439822Z 72bed4b6594b: Pushed
2025-01-06T09:50:35.5345778Z feat-add-crud-app-title-3ae4eea759e53cba27221bd73c1e0ffee14b2c39: digest: sha256:dfd5201604551585168c6ec52bf2ae8f9ed0f4e06da4502b4e7e6932efe7e9b2 size: 2204
2025-01-06T09:50:35.5416345Z ##[group]Run docker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend:$BRANCH-$GITHUB_SHA
2025-01-06T09:50:35.5417946Z [36;1mdocker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend:$BRANCH-$GITHUB_SHA[0m
2025-01-06T09:50:35.5453471Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:50:35.5453803Z env:
2025-01-06T09:50:35.5454083Z   PROJECT_ID: ***
2025-01-06T09:50:35.5454359Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:35.5454680Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:35.5454993Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:35.5455531Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:35.5456134Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:35.5457031Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:35.5457875Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:35.5458467Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:35.5458803Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:35.5459113Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:35.5459722Z   GCP_PROJECT: ***
2025-01-06T09:50:35.5460232Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:35.5460848Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:35.5461416Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:35.5462193Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:35.5463066Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:35.5463814Z ##[endgroup]
2025-01-06T09:50:36.0694108Z The push refers to repository [europe-north1-docker.pkg.dev/***/kubernetes-course/crud-backend]
2025-01-06T09:50:36.2058152Z d60bdb388cad: Preparing
2025-01-06T09:50:36.2059131Z ca3bc334f67d: Preparing
2025-01-06T09:50:36.2061481Z fdaf7008ac9e: Preparing
2025-01-06T09:50:36.2063085Z 820e3f47fc89: Preparing
2025-01-06T09:50:36.2063651Z 5f70bf18a086: Preparing
2025-01-06T09:50:36.2064044Z 4e061853116c: Preparing
2025-01-06T09:50:36.2064572Z f3fe8ef9af87: Preparing
2025-01-06T09:50:36.2064994Z c3c4291cf0a2: Preparing
2025-01-06T09:50:36.2065418Z 85c6f0cfb532: Preparing
2025-01-06T09:50:36.2065831Z a4fd1e7df47e: Preparing
2025-01-06T09:50:36.2066642Z 2f7b6d216a37: Preparing
2025-01-06T09:50:36.2067116Z f3fe8ef9af87: Waiting
2025-01-06T09:50:36.2067574Z c3c4291cf0a2: Waiting
2025-01-06T09:50:36.2068072Z 85c6f0cfb532: Waiting
2025-01-06T09:50:36.2068489Z a4fd1e7df47e: Waiting
2025-01-06T09:50:36.2068921Z 2f7b6d216a37: Waiting
2025-01-06T09:50:36.2069528Z 4e061853116c: Waiting
2025-01-06T09:50:36.5367960Z 820e3f47fc89: Layer already exists
2025-01-06T09:50:36.6939556Z 4e061853116c: Layer already exists
2025-01-06T09:50:36.8545225Z f3fe8ef9af87: Layer already exists
2025-01-06T09:50:37.0688720Z 5f70bf18a086: Layer already exists
2025-01-06T09:50:37.5146853Z c3c4291cf0a2: Layer already exists
2025-01-06T09:50:37.6673712Z a4fd1e7df47e: Layer already exists
2025-01-06T09:50:37.7542263Z 85c6f0cfb532: Layer already exists
2025-01-06T09:50:37.8246808Z 2f7b6d216a37: Layer already exists
2025-01-06T09:50:39.2634884Z d60bdb388cad: Pushed
2025-01-06T09:50:39.3537943Z fdaf7008ac9e: Pushed
2025-01-06T09:50:40.6425756Z ca3bc334f67d: Pushed
2025-01-06T09:50:43.8996421Z feat-add-crud-app-title-3ae4eea759e53cba27221bd73c1e0ffee14b2c39: digest: sha256:21ae9a1687103b712b5213572f859495d7c499337c3ac1f189ff12d2d7789691 size: 2621
2025-01-06T09:50:43.9067518Z ##[group]Run docker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend-worker:$BRANCH-$GITHUB_SHA
2025-01-06T09:50:43.9068489Z [36;1mdocker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend-worker:$BRANCH-$GITHUB_SHA[0m
2025-01-06T09:50:43.9099221Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:50:43.9100044Z env:
2025-01-06T09:50:43.9100546Z   PROJECT_ID: ***
2025-01-06T09:50:43.9101042Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:43.9101393Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:43.9101712Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:43.9102251Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:43.9102846Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:43.9103731Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:43.9104582Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:43.9105172Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:43.9105503Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:43.9105805Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:43.9106097Z   GCP_PROJECT: ***
2025-01-06T09:50:43.9106396Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:43.9106751Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:43.9107159Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:43.9107705Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:43.9108446Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:43.9108966Z ##[endgroup]
2025-01-06T09:50:44.4427004Z The push refers to repository [europe-north1-docker.pkg.dev/***/kubernetes-course/crud-backend-worker]
2025-01-06T09:50:44.5785575Z d60bdb388cad: Preparing
2025-01-06T09:50:44.5786538Z ca3bc334f67d: Preparing
2025-01-06T09:50:44.5787068Z fdaf7008ac9e: Preparing
2025-01-06T09:50:44.5787524Z 820e3f47fc89: Preparing
2025-01-06T09:50:44.5787958Z 5f70bf18a086: Preparing
2025-01-06T09:50:44.5788395Z 4e061853116c: Preparing
2025-01-06T09:50:44.5788843Z f3fe8ef9af87: Preparing
2025-01-06T09:50:44.5789187Z c3c4291cf0a2: Preparing
2025-01-06T09:50:44.5789751Z 85c6f0cfb532: Preparing
2025-01-06T09:50:44.5790049Z a4fd1e7df47e: Preparing
2025-01-06T09:50:44.5790320Z 2f7b6d216a37: Preparing
2025-01-06T09:50:44.5790589Z 4e061853116c: Waiting
2025-01-06T09:50:44.5790858Z f3fe8ef9af87: Waiting
2025-01-06T09:50:44.5791113Z c3c4291cf0a2: Waiting
2025-01-06T09:50:44.5791364Z 85c6f0cfb532: Waiting
2025-01-06T09:50:44.5791613Z a4fd1e7df47e: Waiting
2025-01-06T09:50:44.5791866Z 2f7b6d216a37: Waiting
2025-01-06T09:50:44.8816211Z 5f70bf18a086: Layer already exists
2025-01-06T09:50:44.8947565Z ca3bc334f67d: Layer already exists
2025-01-06T09:50:44.8976610Z fdaf7008ac9e: Layer already exists
2025-01-06T09:50:44.8990147Z d60bdb388cad: Layer already exists
2025-01-06T09:50:44.9041592Z 820e3f47fc89: Layer already exists
2025-01-06T09:50:45.0307190Z 4e061853116c: Layer already exists
2025-01-06T09:50:45.0486581Z f3fe8ef9af87: Layer already exists
2025-01-06T09:50:45.0501196Z c3c4291cf0a2: Layer already exists
2025-01-06T09:50:45.0548999Z a4fd1e7df47e: Layer already exists
2025-01-06T09:50:45.0560523Z 85c6f0cfb532: Layer already exists
2025-01-06T09:50:45.7337487Z 2f7b6d216a37: Layer already exists
2025-01-06T09:50:48.4931985Z feat-add-crud-app-title-3ae4eea759e53cba27221bd73c1e0ffee14b2c39: digest: sha256:8540735de085ac750db19d24b1a567eac36fa81190f6dc2d50b10eabad9ea81b size: 2621
2025-01-06T09:50:48.5013514Z ##[group]Run kubectl config set-context --current --namespace=crud-$***GITHUB_REF#refs/heads/***
2025-01-06T09:50:48.5014265Z [36;1mkubectl config set-context --current --namespace=crud-$***GITHUB_REF#refs/heads/***[0m
2025-01-06T09:50:48.5014899Z [36;1mkubectl create namespace crud-$***GITHUB_REF#refs/heads/*** || true[0m
2025-01-06T09:50:48.5015329Z [36;1mcd crud-db[0m
2025-01-06T09:50:48.5015700Z [36;1mkustomize edit set namespace crud-$***GITHUB_REF#refs/heads/***[0m
2025-01-06T09:50:48.5016115Z [36;1mcd ..[0m
2025-01-06T09:50:48.5016368Z [36;1mcd crud-app[0m
2025-01-06T09:50:48.5016739Z [36;1mkustomize edit set namespace crud-$***GITHUB_REF#refs/heads/***[0m
2025-01-06T09:50:48.5017145Z [36;1mcd ..[0m
2025-01-06T09:50:48.5017403Z [36;1mcd crud-backend[0m
2025-01-06T09:50:48.5017784Z [36;1mkustomize edit set namespace crud-$***GITHUB_REF#refs/heads/***[0m
2025-01-06T09:50:48.5044233Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:50:48.5044545Z env:
2025-01-06T09:50:48.5044856Z   PROJECT_ID: ***
2025-01-06T09:50:48.5045133Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:48.5045426Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:48.5045738Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:48.5046260Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:48.5046842Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:48.5047728Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:48.5048535Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:48.5049128Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:48.5049672Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:48.5049984Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:48.5050275Z   GCP_PROJECT: ***
2025-01-06T09:50:48.5050572Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:48.5050930Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:48.5051354Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:48.5051905Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:48.5052656Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:48.5053182Z ##[endgroup]
2025-01-06T09:50:48.9476977Z Context "gke_***_europe-north1-b_dwk-cluster" modified.
2025-01-06T09:50:49.7353203Z Error from server (AlreadyExists): namespaces "crud-feat-add-crud-app-title" already exists
2025-01-06T09:50:49.7967067Z ##[group]Run cd crud-db
2025-01-06T09:50:49.7967656Z [36;1mcd crud-db[0m
2025-01-06T09:50:49.7968207Z [36;1mkustomize build . | kubectl apply -f -[0m
2025-01-06T09:50:49.7968878Z [36;1mkubectl rollout status deployment[0m
2025-01-06T09:50:49.7969783Z [36;1mkubectl get services -o wide[0m
2025-01-06T09:50:49.8006915Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:50:49.8007438Z env:
2025-01-06T09:50:49.8007896Z   PROJECT_ID: ***
2025-01-06T09:50:49.8008410Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:49.8008935Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:49.8010062Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:49.8011048Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:49.8012115Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:49.8013796Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:49.8015346Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:49.8016466Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:49.8017132Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:49.8017717Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:49.8018274Z   GCP_PROJECT: ***
2025-01-06T09:50:49.8018852Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:49.8019805Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:49.8020517Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:49.8021438Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:49.8022757Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:49.8023687Z ##[endgroup]
2025-01-06T09:50:52.0234238Z secret/crud-postgres-secrets unchanged
2025-01-06T09:50:52.4424482Z service/crud-postgres-svc created
2025-01-06T09:50:52.6356051Z persistentvolumeclaim/crud-claim unchanged
2025-01-06T09:50:52.9510004Z statefulset.apps/crud-postgres-stset configured
2025-01-06T09:50:53.7360726Z No resources found in crud-feat-add-crud-app-title namespace.
2025-01-06T09:50:54.5263869Z NAME                TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE   SELECTOR
2025-01-06T09:50:54.5265468Z crud-postgres-svc   ClusterIP   None         <none>        5432/TCP   2s    app=crud-postgres
2025-01-06T09:50:54.5319138Z ##[group]Run cd crud-app
2025-01-06T09:50:54.5320021Z [36;1mcd crud-app[0m
2025-01-06T09:50:54.5320638Z [36;1mkustomize edit set image crud/app=europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-app:$BRANCH-$GITHUB_SHA[0m
2025-01-06T09:50:54.5321320Z [36;1mkustomize build . | kubectl apply -f -[0m
2025-01-06T09:50:54.5321703Z [36;1mkubectl rollout status deployment[0m
2025-01-06T09:50:54.5322058Z [36;1mkubectl get services -o wide[0m
2025-01-06T09:50:54.5348460Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:50:54.5348773Z env:
2025-01-06T09:50:54.5349051Z   PROJECT_ID: ***
2025-01-06T09:50:54.5349513Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:50:54.5349878Z   GKE_ZONE: europe-north1-b
2025-01-06T09:50:54.5350189Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:50:54.5350730Z   SOPS_AGE_KEY: ***
2025-01-06T09:50:54.5351314Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:54.5352235Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:54.5353059Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:50:54.5353656Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:50:54.5354022Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:50:54.5354328Z   GCLOUD_PROJECT: ***
2025-01-06T09:50:54.5354626Z   GCP_PROJECT: ***
2025-01-06T09:50:54.5354930Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:50:54.5355303Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:50:54.5355733Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:50:54.5356277Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:54.5357025Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:50:54.5357555Z ##[endgroup]
2025-01-06T09:50:56.2850325Z service/crud-app-svc created
2025-01-06T09:50:56.6672304Z deployment.apps/crud-app created
2025-01-06T09:50:57.0574270Z ingress.networking.k8s.io/crud-app-ingress created
2025-01-06T09:50:57.9533919Z Waiting for deployment "crud-app" rollout to finish: 0 of 1 updated replicas are available...
2025-01-06T09:51:09.8136729Z deployment "crud-app" successfully rolled out
2025-01-06T09:51:10.5753448Z NAME                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE   SELECTOR
2025-01-06T09:51:10.5754435Z crud-app-svc        NodePort    34.118.230.151   <none>        2347:32468/TCP   14s   app=crud-app
2025-01-06T09:51:10.5755516Z crud-postgres-svc   ClusterIP   None             <none>        5432/TCP         18s   app=crud-postgres
2025-01-06T09:51:10.5815512Z ##[group]Run cd crud-backend
2025-01-06T09:51:10.5816119Z [36;1mcd crud-backend[0m
2025-01-06T09:51:10.5817217Z [36;1mkustomize edit set image crud/backend=europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend:$BRANCH-$GITHUB_SHA[0m
2025-01-06T09:51:10.5819043Z [36;1mkustomize edit set image crud/backend-worker=europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend-worker:$BRANCH-$GITHUB_SHA[0m
2025-01-06T09:51:10.5820829Z [36;1mkustomize build . | kubectl apply -f -[0m
2025-01-06T09:51:10.5821500Z [36;1mkubectl rollout status deployment[0m
2025-01-06T09:51:10.5822110Z [36;1mkubectl get services -o wide[0m
2025-01-06T09:51:10.5860373Z shell: /usr/bin/bash -e ***0***
2025-01-06T09:51:10.5860910Z env:
2025-01-06T09:51:10.5861967Z   PROJECT_ID: ***
2025-01-06T09:51:10.5862453Z   GKE_CLUSTER: dwk-cluster
2025-01-06T09:51:10.5862958Z   GKE_ZONE: europe-north1-b
2025-01-06T09:51:10.5863481Z   BRANCH: feat-add-crud-app-title
2025-01-06T09:51:10.5864478Z   SOPS_AGE_KEY: ***
2025-01-06T09:51:10.5865496Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:51:10.5867132Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:51:10.5868607Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json
2025-01-06T09:51:10.5869871Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-06T09:51:10.5870446Z   CLOUDSDK_PROJECT: ***
2025-01-06T09:51:10.5870985Z   GCLOUD_PROJECT: ***
2025-01-06T09:51:10.5871505Z   GCP_PROJECT: ***
2025-01-06T09:51:10.5872038Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-06T09:51:10.5872664Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-06T09:51:10.5873391Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-06T09:51:10.5874378Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:51:10.5875726Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-93df279e7933b0ae
2025-01-06T09:51:10.5876706Z ##[endgroup]
2025-01-06T09:51:12.3877433Z service/crud-backend-svc created
2025-01-06T09:51:12.7528600Z deployment.apps/crud-backend created
2025-01-06T09:51:13.0644091Z cronjob.batch/crud-backend-worker configured
2025-01-06T09:51:13.5031292Z ingress.networking.k8s.io/crud-backend-ingress created
2025-01-06T09:51:14.4981611Z deployment "crud-app" successfully rolled out
2025-01-06T09:51:14.7347751Z deployment "crud-backend" successfully rolled out
2025-01-06T09:51:15.7233169Z NAME                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE   SELECTOR
2025-01-06T09:51:15.7234240Z crud-app-svc        NodePort    34.118.230.151   <none>        2347:32468/TCP   19s   app=crud-app
2025-01-06T09:51:15.7235259Z crud-backend-svc    NodePort    34.118.225.207   <none>        2348:30697/TCP   3s    app=crud-backend
2025-01-06T09:51:15.7236351Z crud-postgres-svc   ClusterIP   None             <none>        5432/TCP         23s   app=crud-postgres
2025-01-06T09:51:15.7321277Z Post job cleanup.
2025-01-06T09:51:15.8169634Z Removed exported credentials at "/home/runner/work/kubernetes-course/kubernetes-course/gha-creds-05f0e4707c2cd11b.json".
2025-01-06T09:51:15.8286487Z Post job cleanup.
2025-01-06T09:51:15.9258375Z [command]/usr/bin/git version
2025-01-06T09:51:15.9319282Z git version 2.47.1
2025-01-06T09:51:15.9398049Z Temporarily overriding HOME='/home/runner/work/_temp/3c0201a3-116e-4dc1-a121-84042bbb3162' before making global git config changes
2025-01-06T09:51:15.9399578Z Adding repository directory to the temporary git global config as a safe directory
2025-01-06T09:51:15.9403937Z [command]/usr/bin/git config --global --add safe.directory /home/runner/work/kubernetes-course/kubernetes-course
2025-01-06T09:51:15.9480076Z [command]/usr/bin/git config --local --name-only --get-regexp core\.sshCommand
2025-01-06T09:51:15.9558786Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'core\.sshCommand' && git config --local --unset-all 'core.sshCommand' || :"
2025-01-06T09:51:15.9955099Z [command]/usr/bin/git config --local --name-only --get-regexp http\.https\:\/\/github\.com\/\.extraheader
2025-01-06T09:51:16.0002680Z http.https://github.com/.extraheader
2025-01-06T09:51:16.0033621Z [command]/usr/bin/git config --local --unset-all http.https://github.com/.extraheader
2025-01-06T09:51:16.0094111Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'http\.https\:\/\/github\.com\/\.extraheader' && git config --local --unset-all 'http.https://github.com/.extraheader' || :"
2025-01-06T09:51:16.0486965Z Cleaning up orphan processes
```
