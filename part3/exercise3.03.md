## Files

- https://github.com/peltomaa/kubernetes-course/blob/159d26ac143e5e60d01b412097e664921b9bfc38/.github/workflows/deploy.yaml

## Service running

```bash
➜  ~ curl http://34.8.33.189
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>CRUD App</title>
  </head>
  <body>
    <img alt="Random image" width="500" src="/img" />
    <form id="todo-form">
      <input type="text" name="task" id="form-task" maxlength="140" required />
      <button type="submit" id="form-submit">Create TODO</button>
    </form>
    <ul id="todos"></ul>

    <script>
      const getTodos = async () => {
        const res = await fetch("/todos");
        const json = await res.json();
        return json;
      };

      const postTodo = async (todo) => {
        const res = await fetch("/todos", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(todo),
        });
        const json = await res.json();
        return json;
      };

      const injectTodos = async () => {
        const todos = await getTodos();
        const ul = document.getElementById("todos");
        ul.innerHTML = "";
        todos.forEach((todo) => {
          ul.innerHTML += `<li>${todo.task}</li>`;
        });
      };

      const handleForm = async (e) => {
        e.preventDefault();
        const input = document.getElementById("form-task");
        const todos = await postTodo({ task: input.value.trim() });
        input.value = "";

        const ul = document.getElementById("todos");
        ul.innerHTML = "";
        todos.forEach((todo) => {
          ul.innerHTML += `<li>${todo.task}</li>`;
        });
      };

      document
        .getElementById("todo-form")
        .addEventListener("submit", handleForm);

      injectTodos();
    </script>
  </body>
</html>
➜  ~
```

## Github action deployment logs

```bash
2025-01-05T18:28:03.3534723Z Current runner version: '2.321.0'
2025-01-05T18:28:03.3562063Z ##[group]Operating System
2025-01-05T18:28:03.3563099Z Ubuntu
2025-01-05T18:28:03.3563895Z 24.04.1
2025-01-05T18:28:03.3564694Z LTS
2025-01-05T18:28:03.3565440Z ##[endgroup]
2025-01-05T18:28:03.3566204Z ##[group]Runner Image
2025-01-05T18:28:03.3567076Z Image: ubuntu-24.04
2025-01-05T18:28:03.3567861Z Version: 20241215.1.0
2025-01-05T18:28:03.3569477Z Included Software: https://github.com/actions/runner-images/blob/ubuntu24/20241215.1/images/ubuntu/Ubuntu2404-Readme.md
2025-01-05T18:28:03.3571196Z Image Release: https://github.com/actions/runner-images/releases/tag/ubuntu24%2F20241215.1
2025-01-05T18:28:03.3572624Z ##[endgroup]
2025-01-05T18:28:03.3573424Z ##[group]Runner Image Provisioner
2025-01-05T18:28:03.3574378Z 2.0.404.1
2025-01-05T18:28:03.3575084Z ##[endgroup]
2025-01-05T18:28:03.3576802Z ##[group]GITHUB_TOKEN Permissions
2025-01-05T18:28:03.3579469Z Contents: read
2025-01-05T18:28:03.3580255Z Metadata: read
2025-01-05T18:28:03.3581291Z Packages: read
2025-01-05T18:28:03.3582169Z ##[endgroup]
2025-01-05T18:28:03.3585564Z Secret source: Actions
2025-01-05T18:28:03.3586927Z Prepare workflow directory
2025-01-05T18:28:03.4077765Z Prepare all required actions
2025-01-05T18:28:03.4113950Z Getting action download info
2025-01-05T18:28:03.6931310Z Download action repository 'actions/checkout@v4' (SHA:11bd71901bbe5b1630ceea73d27597364c9af683)
2025-01-05T18:28:03.7887144Z Download action repository 'google-github-actions/auth@v2' (SHA:6fc4af4b145ae7821d527454aa9bd537d1f2dc5f)
2025-01-05T18:28:04.2586246Z Download action repository 'google-github-actions/setup-gcloud@v2' (SHA:6189d56e4096ee891640bb02ac264be376592d6a)
2025-01-05T18:28:04.7297295Z Download action repository 'google-github-actions/get-gke-credentials@v2' (SHA:9025e8f90f2d8e0c3dafc3128cc705a26d992a6a)
2025-01-05T18:28:05.2945486Z Complete job name: build
2025-01-05T18:28:05.3873309Z ##[group]Run actions/checkout@v4
2025-01-05T18:28:05.3874793Z with:
2025-01-05T18:28:05.3875768Z   repository: peltomaa/kubernetes-course
2025-01-05T18:28:05.3877320Z   token: ***
2025-01-05T18:28:05.3878404Z   ssh-strict: true
2025-01-05T18:28:05.3879372Z   ssh-user: git
2025-01-05T18:28:05.3880351Z   persist-credentials: true
2025-01-05T18:28:05.3881416Z   clean: true
2025-01-05T18:28:05.3882407Z   sparse-checkout-cone-mode: true
2025-01-05T18:28:05.3883606Z   fetch-depth: 1
2025-01-05T18:28:05.3884558Z   fetch-tags: false
2025-01-05T18:28:05.3885545Z   show-progress: true
2025-01-05T18:28:05.3886553Z   lfs: false
2025-01-05T18:28:05.3887449Z   submodules: false
2025-01-05T18:28:05.3888552Z   set-safe-directory: true
2025-01-05T18:28:05.3889875Z env:
2025-01-05T18:28:05.3890868Z   PROJECT_ID: ***
2025-01-05T18:28:05.3891832Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:28:05.3892906Z   GKE_ZONE: europe-north1-b
2025-01-05T18:28:05.3893952Z   BRANCH: main
2025-01-05T18:28:05.3894900Z ##[endgroup]
2025-01-05T18:28:05.5747966Z Syncing repository: peltomaa/kubernetes-course
2025-01-05T18:28:05.5752318Z ##[group]Getting Git version info
2025-01-05T18:28:05.5755853Z Working directory is '/home/runner/work/kubernetes-course/kubernetes-course'
2025-01-05T18:28:05.5759492Z [command]/usr/bin/git version
2025-01-05T18:28:05.5774826Z git version 2.47.1
2025-01-05T18:28:05.5816311Z ##[endgroup]
2025-01-05T18:28:05.5827846Z Temporarily overriding HOME='/home/runner/work/_temp/a700b2b4-81db-4623-be7d-373813e2b2da' before making global git config changes
2025-01-05T18:28:05.5833838Z Adding repository directory to the temporary git global config as a safe directory
2025-01-05T18:28:05.5853929Z [command]/usr/bin/git config --global --add safe.directory /home/runner/work/kubernetes-course/kubernetes-course
2025-01-05T18:28:05.5895069Z Deleting the contents of '/home/runner/work/kubernetes-course/kubernetes-course'
2025-01-05T18:28:05.5900406Z ##[group]Initializing the repository
2025-01-05T18:28:05.5905522Z [command]/usr/bin/git init /home/runner/work/kubernetes-course/kubernetes-course
2025-01-05T18:28:05.5982184Z hint: Using 'master' as the name for the initial branch. This default branch name
2025-01-05T18:28:05.5985244Z hint: is subject to change. To configure the initial branch name to use in all
2025-01-05T18:28:05.5987268Z hint: of your new repositories, which will suppress this warning, call:
2025-01-05T18:28:05.5989657Z hint:
2025-01-05T18:28:05.5991599Z hint: 	git config --global init.defaultBranch <name>
2025-01-05T18:28:05.5993979Z hint:
2025-01-05T18:28:05.5996231Z hint: Names commonly chosen instead of 'master' are 'main', 'trunk' and
2025-01-05T18:28:05.6000246Z hint: 'development'. The just-created branch can be renamed via this command:
2025-01-05T18:28:05.6003366Z hint:
2025-01-05T18:28:05.6004901Z hint: 	git branch -m <name>
2025-01-05T18:28:05.6007901Z Initialized empty Git repository in /home/runner/work/kubernetes-course/kubernetes-course/.git/
2025-01-05T18:28:05.6013901Z [command]/usr/bin/git remote add origin https://github.com/peltomaa/kubernetes-course
2025-01-05T18:28:05.6054777Z ##[endgroup]
2025-01-05T18:28:05.6057785Z ##[group]Disabling automatic garbage collection
2025-01-05T18:28:05.6060728Z [command]/usr/bin/git config --local gc.auto 0
2025-01-05T18:28:05.6097214Z ##[endgroup]
2025-01-05T18:28:05.6100316Z ##[group]Setting up auth
2025-01-05T18:28:05.6106715Z [command]/usr/bin/git config --local --name-only --get-regexp core\.sshCommand
2025-01-05T18:28:05.6137407Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'core\.sshCommand' && git config --local --unset-all 'core.sshCommand' || :"
2025-01-05T18:28:05.6455280Z [command]/usr/bin/git config --local --name-only --get-regexp http\.https\:\/\/github\.com\/\.extraheader
2025-01-05T18:28:05.6491568Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'http\.https\:\/\/github\.com\/\.extraheader' && git config --local --unset-all 'http.https://github.com/.extraheader' || :"
2025-01-05T18:28:05.6733956Z [command]/usr/bin/git config --local http.https://github.com/.extraheader AUTHORIZATION: basic ***
2025-01-05T18:28:05.6783982Z ##[endgroup]
2025-01-05T18:28:05.6789346Z ##[group]Fetching the repository
2025-01-05T18:28:05.6807514Z [command]/usr/bin/git -c protocol.version=2 fetch --no-tags --prune --no-recurse-submodules --depth=1 origin +ee29c5b4f7135cc0ed7ad009ad5929306216d337:refs/remotes/origin/main
2025-01-05T18:28:06.2880085Z From https://github.com/peltomaa/kubernetes-course
2025-01-05T18:28:06.2881409Z  * [new ref]         ee29c5b4f7135cc0ed7ad009ad5929306216d337 -> origin/main
2025-01-05T18:28:06.2883808Z ##[endgroup]
2025-01-05T18:28:06.2884814Z ##[group]Determining the checkout info
2025-01-05T18:28:06.2885802Z ##[endgroup]
2025-01-05T18:28:06.2886404Z [command]/usr/bin/git sparse-checkout disable
2025-01-05T18:28:06.2898518Z [command]/usr/bin/git config --local --unset-all extensions.worktreeConfig
2025-01-05T18:28:06.2930616Z ##[group]Checking out the ref
2025-01-05T18:28:06.2934079Z [command]/usr/bin/git checkout --progress --force -B main refs/remotes/origin/main
2025-01-05T18:28:06.3112605Z Switched to a new branch 'main'
2025-01-05T18:28:06.3119921Z branch 'main' set up to track 'origin/main'.
2025-01-05T18:28:06.3126548Z ##[endgroup]
2025-01-05T18:28:06.3166250Z [command]/usr/bin/git log -1 --format=%H
2025-01-05T18:28:06.3189774Z ee29c5b4f7135cc0ed7ad009ad5929306216d337
2025-01-05T18:28:06.3511148Z ##[group]Run google-github-actions/auth@v2
2025-01-05T18:28:06.3511586Z with:
2025-01-05T18:28:06.3521923Z   credentials_json: ***
2025-01-05T18:28:06.3522510Z   create_credentials_file: true
2025-01-05T18:28:06.3523175Z   export_environment_variables: true
2025-01-05T18:28:06.3523809Z   universe: googleapis.com
2025-01-05T18:28:06.3524387Z   cleanup_credentials: true
2025-01-05T18:28:06.3525044Z   access_token_lifetime: 3600s
2025-01-05T18:28:06.3525756Z   access_token_scopes: https://www.googleapis.com/auth/cloud-platform
2025-01-05T18:28:06.3526243Z   id_token_include_email: false
2025-01-05T18:28:06.3526579Z env:
2025-01-05T18:28:06.3526958Z   PROJECT_ID: ***
2025-01-05T18:28:06.3527517Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:28:06.3527864Z   GKE_ZONE: europe-north1-b
2025-01-05T18:28:06.3528460Z   BRANCH: main
2025-01-05T18:28:06.3528767Z ##[endgroup]
2025-01-05T18:28:06.4458561Z Created credentials file at "/home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json"
2025-01-05T18:28:06.4591246Z ##[group]Run google-github-actions/setup-gcloud@v2
2025-01-05T18:28:06.4591708Z with:
2025-01-05T18:28:06.4592148Z   install_components: kubectl
2025-01-05T18:28:06.4592506Z   version: latest
2025-01-05T18:28:06.4592827Z   skip_install: false
2025-01-05T18:28:06.4593141Z env:
2025-01-05T18:28:06.4593512Z   PROJECT_ID: ***
2025-01-05T18:28:06.4593833Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:28:06.4594180Z   GKE_ZONE: europe-north1-b
2025-01-05T18:28:06.4594514Z   BRANCH: main
2025-01-05T18:28:06.4595143Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:28:06.4596142Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:28:06.4597058Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:28:06.4597723Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:28:06.4598357Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:28:06.4598737Z   GCLOUD_PROJECT: ***
2025-01-05T18:28:06.4599098Z   GCP_PROJECT: ***
2025-01-05T18:28:06.4599460Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:28:06.4599806Z ##[endgroup]
2025-01-05T18:28:07.8153653Z [command]/usr/bin/tar xz --warning=no-unknown-keyword --overwrite -C /home/runner/work/_temp/d95434c0-6f66-464d-971a-b8d801ae9248 -f /home/runner/work/_temp/ef1358d9-b51c-4f89-a747-1d62c746f29f
2025-01-05T18:29:05.0457357Z Successfully authenticated
2025-01-05T18:29:05.0621313Z ##[group]Run google-github-actions/get-gke-credentials@v2
2025-01-05T18:29:05.0621745Z with:
2025-01-05T18:29:05.0621997Z   cluster_name: dwk-cluster
2025-01-05T18:29:05.0622358Z   project_id: ***
2025-01-05T18:29:05.0622632Z   location: europe-north1-b
2025-01-05T18:29:05.0622932Z   use_auth_provider: false
2025-01-05T18:29:05.0623217Z   use_internal_ip: false
2025-01-05T18:29:05.0623542Z   use_connect_gateway: false
2025-01-05T18:29:05.0623827Z   use_dns_based_endpoint: false
2025-01-05T18:29:05.0624110Z env:
2025-01-05T18:29:05.0624375Z   PROJECT_ID: ***
2025-01-05T18:29:05.0624662Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:29:05.0624948Z   GKE_ZONE: europe-north1-b
2025-01-05T18:29:05.0625228Z   BRANCH: main
2025-01-05T18:29:05.0625791Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:05.0626691Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:05.0627487Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:05.0628371Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:29:05.0628745Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:29:05.0629047Z   GCLOUD_PROJECT: ***
2025-01-05T18:29:05.0629329Z   GCP_PROJECT: ***
2025-01-05T18:29:05.0629626Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:29:05.0629975Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-05T18:29:05.0630377Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-05T18:29:05.0630706Z ##[endgroup]
2025-01-05T18:29:06.3812072Z Successfully created and exported "KUBECONFIG" at: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:06.3907234Z ##[group]Run gcloud --quiet auth configure-docker europe-north1-docker.pkg.dev
2025-01-05T18:29:06.3908554Z [36;1mgcloud --quiet auth configure-docker europe-north1-docker.pkg.dev[0m
2025-01-05T18:29:06.3939472Z shell: /usr/bin/bash -e ***0***
2025-01-05T18:29:06.3940028Z env:
2025-01-05T18:29:06.3940431Z   PROJECT_ID: ***
2025-01-05T18:29:06.3941085Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:29:06.3941539Z   GKE_ZONE: europe-north1-b
2025-01-05T18:29:06.3942001Z   BRANCH: main
2025-01-05T18:29:06.3942671Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:06.3943749Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:06.3944641Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:06.3945432Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:29:06.3946016Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:29:06.3946463Z   GCLOUD_PROJECT: ***
2025-01-05T18:29:06.3946930Z   GCP_PROJECT: ***
2025-01-05T18:29:06.3947341Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:29:06.3947904Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-05T18:29:06.3948649Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-05T18:29:06.3949371Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:06.3950295Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:06.3950938Z ##[endgroup]
2025-01-05T18:29:09.7013692Z Adding credentials for: europe-north1-docker.pkg.dev
2025-01-05T18:29:09.7187538Z Docker configuration file updated.
2025-01-05T18:29:09.7852392Z ##[group]Run docker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-app:$BRANCH-$GITHUB_SHA crud-app/
2025-01-05T18:29:09.7853514Z [36;1mdocker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-app:$BRANCH-$GITHUB_SHA crud-app/[0m
2025-01-05T18:29:09.7881317Z shell: /usr/bin/bash -e ***0***
2025-01-05T18:29:09.7881769Z env:
2025-01-05T18:29:09.7882261Z   PROJECT_ID: ***
2025-01-05T18:29:09.7882628Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:29:09.7882906Z   GKE_ZONE: europe-north1-b
2025-01-05T18:29:09.7883186Z   BRANCH: main
2025-01-05T18:29:09.7883736Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:09.7884630Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:09.7885455Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:09.7886027Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:29:09.7886337Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:29:09.7886634Z   GCLOUD_PROJECT: ***
2025-01-05T18:29:09.7886908Z   GCP_PROJECT: ***
2025-01-05T18:29:09.7887220Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:29:09.7887560Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-05T18:29:09.7887947Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-05T18:29:09.7888739Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:09.7889467Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:09.7889961Z ##[endgroup]
2025-01-05T18:29:13.0319276Z #0 building with "default" instance using docker driver
2025-01-05T18:29:13.0319731Z
2025-01-05T18:29:13.0320151Z #1 [internal] load build definition from Dockerfile
2025-01-05T18:29:13.0324153Z #1 transferring dockerfile: 130B done
2025-01-05T18:29:13.0325259Z #1 DONE 0.0s
2025-01-05T18:29:13.0326398Z
2025-01-05T18:29:13.0326912Z #2 [internal] load metadata for docker.io/library/golang:1.23
2025-01-05T18:29:13.2759217Z #2 ...
2025-01-05T18:29:13.2759502Z
2025-01-05T18:29:13.2759824Z #3 [auth] library/golang:pull token for registry-1.docker.io
2025-01-05T18:29:13.2760444Z #3 DONE 0.0s
2025-01-05T18:29:13.4250681Z
2025-01-05T18:29:13.4251183Z #2 [internal] load metadata for docker.io/library/golang:1.23
2025-01-05T18:29:13.8700265Z #2 DONE 1.0s
2025-01-05T18:29:13.9982709Z
2025-01-05T18:29:13.9984094Z #4 [internal] load .dockerignore
2025-01-05T18:29:13.9985366Z #4 transferring context: 2B done
2025-01-05T18:29:13.9986023Z #4 DONE 0.0s
2025-01-05T18:29:13.9986255Z
2025-01-05T18:29:13.9986442Z #5 [internal] load build context
2025-01-05T18:29:13.9986942Z #5 transferring context: 6.01kB done
2025-01-05T18:29:13.9987425Z #5 DONE 0.0s
2025-01-05T18:29:13.9987636Z
2025-01-05T18:29:13.9988447Z #6 [1/3] FROM docker.io/library/golang:1.23@sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041
2025-01-05T18:29:13.9989713Z #6 resolve docker.io/library/golang:1.23@sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041 done
2025-01-05T18:29:13.9991065Z #6 sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041 9.74kB / 9.74kB done
2025-01-05T18:29:13.9992210Z #6 sha256:aff7bd25e6a162e9db0a284663d6aff04d456416cb3cc94d692a89be72b0e605 2.32kB / 2.32kB done
2025-01-05T18:29:13.9993373Z #6 sha256:7012e31470cb237fd56d72b6a7d16892ea12e7f8fd361be9010444423f28c821 2.80kB / 2.80kB done
2025-01-05T18:29:13.9994571Z #6 sha256:0a96bdb8280554b560ffee0f2e5f9843dc7b625f28192021ee103ecbcc2d629b 0B / 48.50MB 0.1s
2025-01-05T18:29:13.9995752Z #6 sha256:54c7be425079efba0003054ee884bf72f1ffebca733bedd6f077d2809ee9aa6f 0B / 23.87MB 0.1s
2025-01-05T18:29:13.9996907Z #6 sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 0B / 64.39MB 0.1s
2025-01-05T18:29:14.0986460Z #6 sha256:0a96bdb8280554b560ffee0f2e5f9843dc7b625f28192021ee103ecbcc2d629b 29.36MB / 48.50MB 0.2s
2025-01-05T18:29:14.3009501Z #6 sha256:0a96bdb8280554b560ffee0f2e5f9843dc7b625f28192021ee103ecbcc2d629b 48.50MB / 48.50MB 0.4s done
2025-01-05T18:29:14.3011428Z #6 sha256:54c7be425079efba0003054ee884bf72f1ffebca733bedd6f077d2809ee9aa6f 23.87MB / 23.87MB 0.4s done
2025-01-05T18:29:14.3012869Z #6 sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 24.12MB / 64.39MB 0.4s
2025-01-05T18:29:14.3014677Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 0B / 74.05MB 0.4s
2025-01-05T18:29:14.3016781Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 0B / 92.31MB 0.4s
2025-01-05T18:29:14.4099508Z #6 sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 50.11MB / 64.39MB 0.5s
2025-01-05T18:29:14.4112861Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 5.24MB / 92.31MB 0.5s
2025-01-05T18:29:14.4120197Z #6 extracting sha256:0a96bdb8280554b560ffee0f2e5f9843dc7b625f28192021ee103ecbcc2d629b
2025-01-05T18:29:14.5890301Z #6 sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 62.82MB / 64.39MB 0.6s
2025-01-05T18:29:14.5891497Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 16.55MB / 74.05MB 0.6s
2025-01-05T18:29:14.5892532Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 18.34MB / 92.31MB 0.6s
2025-01-05T18:29:14.6943509Z #6 sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 64.39MB / 64.39MB 0.7s done
2025-01-05T18:29:14.6953471Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 34.60MB / 74.05MB 0.7s
2025-01-05T18:29:14.6954620Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 35.80MB / 92.31MB 0.7s
2025-01-05T18:29:14.6955726Z #6 sha256:3fd114183f3282d111ed7eaa48e1f94ff3018db89a43f47239fed2180f2d1084 0B / 125B 0.7s
2025-01-05T18:29:14.8113411Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 48.23MB / 74.05MB 0.8s
2025-01-05T18:29:14.8117755Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 52.43MB / 92.31MB 0.8s
2025-01-05T18:29:14.8119137Z #6 sha256:3fd114183f3282d111ed7eaa48e1f94ff3018db89a43f47239fed2180f2d1084 125B / 125B 0.8s
2025-01-05T18:29:14.9202205Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 74.05MB / 74.05MB 1.0s
2025-01-05T18:29:14.9211734Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 90.50MB / 92.31MB 1.0s
2025-01-05T18:29:14.9212937Z #6 sha256:3fd114183f3282d111ed7eaa48e1f94ff3018db89a43f47239fed2180f2d1084 125B / 125B 0.8s done
2025-01-05T18:29:15.0202818Z #6 sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1 0B / 32B 1.1s
2025-01-05T18:29:15.1257218Z #6 sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1 32B / 32B 1.2s
2025-01-05T18:29:15.3982236Z #6 sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 74.05MB / 74.05MB 1.4s done
2025-01-05T18:29:15.8459102Z #6 sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 92.31MB / 92.31MB 1.9s done
2025-01-05T18:29:15.8460357Z #6 sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1 32B / 32B 1.9s done
2025-01-05T18:29:16.6083914Z #6 extracting sha256:0a96bdb8280554b560ffee0f2e5f9843dc7b625f28192021ee103ecbcc2d629b 2.3s done
2025-01-05T18:29:16.7741939Z #6 extracting sha256:54c7be425079efba0003054ee884bf72f1ffebca733bedd6f077d2809ee9aa6f 0.1s
2025-01-05T18:29:17.2209028Z #6 extracting sha256:54c7be425079efba0003054ee884bf72f1ffebca733bedd6f077d2809ee9aa6f 0.5s done
2025-01-05T18:29:18.9872084Z #6 extracting sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1
2025-01-05T18:29:21.4999916Z #6 extracting sha256:7aa8176e6d893aff4b57b2c22574ec2afadff4673b8e0954e275244bfd3d7bc1 2.4s done
2025-01-05T18:29:22.1020115Z #6 extracting sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4
2025-01-05T18:29:24.3049029Z #6 extracting sha256:4930ffbfb2152fc9d9ccd8712b7162244c1b95a0998025070dbb4229bc0564d4 2.2s done
2025-01-05T18:29:25.3786133Z #6 extracting sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c
2025-01-05T18:29:29.6725854Z #6 extracting sha256:06f05ace1117d62b655e04f6f73c83617e3e0febc38681dbedf58f477dd0658c 4.3s done
2025-01-05T18:29:31.0622902Z #6 extracting sha256:3fd114183f3282d111ed7eaa48e1f94ff3018db89a43f47239fed2180f2d1084
2025-01-05T18:29:31.2935159Z #6 extracting sha256:3fd114183f3282d111ed7eaa48e1f94ff3018db89a43f47239fed2180f2d1084 done
2025-01-05T18:29:31.2937752Z #6 extracting sha256:4f4fb700ef54461cfa02571ae0db9a0dc1e0cdb5577484a6d75e68dc38e8acc1 done
2025-01-05T18:29:31.2938743Z #6 DONE 17.2s
2025-01-05T18:29:31.2938980Z
2025-01-05T18:29:31.2939142Z #7 [2/3] WORKDIR /app
2025-01-05T18:29:31.2939525Z #7 DONE 0.0s
2025-01-05T18:29:31.2939728Z
2025-01-05T18:29:31.2939881Z #8 [3/3] COPY . .
2025-01-05T18:29:31.2940263Z #8 DONE 0.0s
2025-01-05T18:29:31.2940527Z
2025-01-05T18:29:31.2940694Z #9 exporting to image
2025-01-05T18:29:31.2941090Z #9 exporting layers
2025-01-05T18:29:33.0446501Z #9 exporting layers 1.9s done
2025-01-05T18:29:33.0672315Z #9 writing image sha256:849e9a7323e56ef327b5b9c7d82dea46239d4b00fb90a0a1f2ea0631f4693b47 done
2025-01-05T18:29:33.0674348Z #9 naming to europe-north1-docker.pkg.dev/***/kubernetes-course/crud-app:main-ee29c5b4f7135cc0ed7ad009ad5929306216d337 done
2025-01-05T18:29:33.0675498Z #9 DONE 1.9s
2025-01-05T18:29:33.0769053Z ##[group]Run docker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend:$BRANCH-$GITHUB_SHA crud-backend/
2025-01-05T18:29:33.0770736Z [36;1mdocker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend:$BRANCH-$GITHUB_SHA crud-backend/[0m
2025-01-05T18:29:33.0809178Z shell: /usr/bin/bash -e ***0***
2025-01-05T18:29:33.0809717Z env:
2025-01-05T18:29:33.0810171Z   PROJECT_ID: ***
2025-01-05T18:29:33.0810604Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:29:33.0811109Z   GKE_ZONE: europe-north1-b
2025-01-05T18:29:33.0811601Z   BRANCH: main
2025-01-05T18:29:33.0812562Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:33.0814113Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:33.0815201Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:33.0815809Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:29:33.0816376Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:29:33.0816672Z   GCLOUD_PROJECT: ***
2025-01-05T18:29:33.0816950Z   GCP_PROJECT: ***
2025-01-05T18:29:33.0817243Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:29:33.0817583Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-05T18:29:33.0818174Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-05T18:29:33.0818751Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:33.0819486Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:33.0820001Z ##[endgroup]
2025-01-05T18:29:33.3693417Z #0 building with "default" instance using docker driver
2025-01-05T18:29:33.3693939Z
2025-01-05T18:29:33.3694213Z #1 [internal] load build definition from Dockerfile
2025-01-05T18:29:33.3695694Z #1 transferring dockerfile: 189B done
2025-01-05T18:29:33.3696096Z #1 DONE 0.0s
2025-01-05T18:29:33.3696244Z
2025-01-05T18:29:33.3696560Z #2 [internal] load metadata for docker.io/library/golang:1.23
2025-01-05T18:29:33.3697019Z #2 DONE 0.1s
2025-01-05T18:29:33.5550282Z
2025-01-05T18:29:33.5550916Z #3 [internal] load .dockerignore
2025-01-05T18:29:33.5551442Z #3 transferring context: 2B done
2025-01-05T18:29:33.5551877Z #3 DONE 0.0s
2025-01-05T18:29:33.5552091Z
2025-01-05T18:29:33.5552630Z #4 [1/5] FROM docker.io/library/golang:1.23@sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041
2025-01-05T18:29:33.5553429Z #4 DONE 0.0s
2025-01-05T18:29:33.5553630Z
2025-01-05T18:29:33.5553782Z #5 [2/5] WORKDIR /app
2025-01-05T18:29:33.5554208Z #5 CACHED
2025-01-05T18:29:33.5554414Z
2025-01-05T18:29:33.5554601Z #6 [internal] load build context
2025-01-05T18:29:33.5555082Z #6 transferring context: 9.51kB done
2025-01-05T18:29:33.5555552Z #6 DONE 0.0s
2025-01-05T18:29:33.5555763Z
2025-01-05T18:29:33.5555975Z #7 [3/5] COPY go.mod go.sum ./
2025-01-05T18:29:33.5556419Z #7 DONE 0.0s
2025-01-05T18:29:33.5556627Z
2025-01-05T18:29:33.5556813Z #8 [4/5] RUN go mod download
2025-01-05T18:29:33.6650560Z #8 DONE 0.3s
2025-01-05T18:29:33.8417697Z
2025-01-05T18:29:33.8418365Z #9 [5/5] COPY . .
2025-01-05T18:29:33.8418993Z #9 DONE 0.0s
2025-01-05T18:29:33.8419338Z
2025-01-05T18:29:33.8419642Z #10 exporting to image
2025-01-05T18:29:33.8420094Z #10 exporting layers
2025-01-05T18:29:36.2901787Z #10 exporting layers 2.6s done
2025-01-05T18:29:36.3157892Z #10 writing image sha256:4342c7c9076170ab78a1bbe893c8233cb4fa9843ac098192058e58276c5d01d5 done
2025-01-05T18:29:36.3164262Z #10 naming to europe-north1-docker.pkg.dev/***/kubernetes-course/crud-backend:main-ee29c5b4f7135cc0ed7ad009ad5929306216d337 done
2025-01-05T18:29:36.3165481Z #10 DONE 2.6s
2025-01-05T18:29:36.3270671Z ##[group]Run docker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend-worker:$BRANCH-$GITHUB_SHA -f crud-backend/Dockerfile-worker crud-backend/
2025-01-05T18:29:36.3272783Z [36;1mdocker build -t europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend-worker:$BRANCH-$GITHUB_SHA -f crud-backend/Dockerfile-worker crud-backend/[0m
2025-01-05T18:29:36.3311748Z shell: /usr/bin/bash -e ***0***
2025-01-05T18:29:36.3312270Z env:
2025-01-05T18:29:36.3312747Z   PROJECT_ID: ***
2025-01-05T18:29:36.3313208Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:29:36.3313712Z   GKE_ZONE: europe-north1-b
2025-01-05T18:29:36.3314199Z   BRANCH: main
2025-01-05T18:29:36.3315176Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:36.3316691Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:36.3318384Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:36.3319425Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:29:36.3320002Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:29:36.3320514Z   GCLOUD_PROJECT: ***
2025-01-05T18:29:36.3321309Z   GCP_PROJECT: ***
2025-01-05T18:29:36.3321819Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:29:36.3322411Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-05T18:29:36.3323107Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-05T18:29:36.3324055Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:36.3325398Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:36.3326344Z ##[endgroup]
2025-01-05T18:29:36.6056161Z #0 building with "default" instance using docker driver
2025-01-05T18:29:36.6056732Z
2025-01-05T18:29:36.6057037Z #1 [internal] load build definition from Dockerfile-worker
2025-01-05T18:29:36.6057665Z #1 transferring dockerfile: 196B done
2025-01-05T18:29:36.6058303Z #1 DONE 0.0s
2025-01-05T18:29:36.6058534Z
2025-01-05T18:29:36.6058830Z #2 [internal] load metadata for docker.io/library/golang:1.23
2025-01-05T18:29:36.6059435Z #2 DONE 0.1s
2025-01-05T18:29:36.6358350Z
2025-01-05T18:29:36.6359194Z #3 [internal] load .dockerignore
2025-01-05T18:29:36.6360797Z #3 transferring context: 2B done
2025-01-05T18:29:36.6361297Z #3 DONE 0.0s
2025-01-05T18:29:36.6361528Z
2025-01-05T18:29:36.6362163Z #4 [1/5] FROM docker.io/library/golang:1.23@sha256:7ea4c9dcb2b97ff8ee80a67db3d44f98c8ffa0d191399197007d8459c1453041
2025-01-05T18:29:36.6363085Z #4 DONE 0.0s
2025-01-05T18:29:36.6363309Z
2025-01-05T18:29:36.6363493Z #5 [internal] load build context
2025-01-05T18:29:36.6363993Z #5 transferring context: 1.00kB done
2025-01-05T18:29:36.6364486Z #5 DONE 0.0s
2025-01-05T18:29:36.6364707Z
2025-01-05T18:29:36.6364901Z #6 [3/5] COPY go.mod go.sum ./
2025-01-05T18:29:36.6365375Z #6 CACHED
2025-01-05T18:29:36.6365580Z
2025-01-05T18:29:36.6365779Z #7 [4/5] RUN go mod download
2025-01-05T18:29:36.6366232Z #7 CACHED
2025-01-05T18:29:36.6366443Z
2025-01-05T18:29:36.6366626Z #8 [2/5] WORKDIR /app
2025-01-05T18:29:36.6367065Z #8 CACHED
2025-01-05T18:29:36.6367258Z
2025-01-05T18:29:36.6367396Z #9 [5/5] COPY . .
2025-01-05T18:29:36.6367780Z #9 CACHED
2025-01-05T18:29:36.6368217Z
2025-01-05T18:29:36.6368405Z #10 exporting to image
2025-01-05T18:29:36.6368838Z #10 exporting layers done
2025-01-05T18:29:36.6369601Z #10 writing image sha256:e1eb318bd72c96e1e971308efc7499212ebc1352ca94a668c5d890c21f11767e done
2025-01-05T18:29:36.6371193Z #10 naming to europe-north1-docker.pkg.dev/***/kubernetes-course/crud-backend-worker:main-ee29c5b4f7135cc0ed7ad009ad5929306216d337 done
2025-01-05T18:29:36.6372267Z #10 DONE 0.0s
2025-01-05T18:29:36.6472738Z ##[group]Run docker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-app:$BRANCH-$GITHUB_SHA
2025-01-05T18:29:36.6474188Z [36;1mdocker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-app:$BRANCH-$GITHUB_SHA[0m
2025-01-05T18:29:36.6513176Z shell: /usr/bin/bash -e ***0***
2025-01-05T18:29:36.6513687Z env:
2025-01-05T18:29:36.6514152Z   PROJECT_ID: ***
2025-01-05T18:29:36.6514599Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:29:36.6515113Z   GKE_ZONE: europe-north1-b
2025-01-05T18:29:36.6515585Z   BRANCH: main
2025-01-05T18:29:36.6516559Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:36.6518272Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:36.6519772Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:36.6520810Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:29:36.6521356Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:29:36.6521864Z   GCLOUD_PROJECT: ***
2025-01-05T18:29:36.6522352Z   GCP_PROJECT: ***
2025-01-05T18:29:36.6522863Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:29:36.6523480Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-05T18:29:36.6524190Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-05T18:29:36.6525178Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:36.6526845Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:36.6527736Z ##[endgroup]
2025-01-05T18:29:37.4927345Z The push refers to repository [europe-north1-docker.pkg.dev/***/kubernetes-course/crud-app]
2025-01-05T18:29:38.5338281Z a5a11bfa35ef: Preparing
2025-01-05T18:29:38.5341439Z 4b4be28e5339: Preparing
2025-01-05T18:29:38.5343738Z 5f70bf18a086: Preparing
2025-01-05T18:29:38.5344189Z 4e061853116c: Preparing
2025-01-05T18:29:38.5344585Z f3fe8ef9af87: Preparing
2025-01-05T18:29:38.5344994Z c3c4291cf0a2: Preparing
2025-01-05T18:29:38.5345416Z 85c6f0cfb532: Preparing
2025-01-05T18:29:38.5345855Z a4fd1e7df47e: Preparing
2025-01-05T18:29:38.5346291Z 2f7b6d216a37: Preparing
2025-01-05T18:29:38.5346686Z c3c4291cf0a2: Waiting
2025-01-05T18:29:38.5347094Z 85c6f0cfb532: Waiting
2025-01-05T18:29:38.5347487Z a4fd1e7df47e: Waiting
2025-01-05T18:29:38.5347898Z 2f7b6d216a37: Waiting
2025-01-05T18:29:40.6667726Z 4e061853116c: Layer already exists
2025-01-05T18:29:40.6780361Z f3fe8ef9af87: Layer already exists
2025-01-05T18:29:40.6880272Z 5f70bf18a086: Layer already exists
2025-01-05T18:29:41.7327881Z c3c4291cf0a2: Layer already exists
2025-01-05T18:29:41.7409966Z 85c6f0cfb532: Layer already exists
2025-01-05T18:29:41.7443794Z a4fd1e7df47e: Layer already exists
2025-01-05T18:29:42.7994327Z 2f7b6d216a37: Layer already exists
2025-01-05T18:29:45.0087880Z a5a11bfa35ef: Pushed
2025-01-05T18:29:45.8531131Z 4b4be28e5339: Pushed
2025-01-05T18:29:50.9531462Z main-ee29c5b4f7135cc0ed7ad009ad5929306216d337: digest: sha256:daf742ab47d4de1ad15c04aae586c11ec19914fccd225f084fb85dcca9669ec8 size: 2204
2025-01-05T18:29:50.9585829Z ##[group]Run docker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend:$BRANCH-$GITHUB_SHA
2025-01-05T18:29:50.9586658Z [36;1mdocker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend:$BRANCH-$GITHUB_SHA[0m
2025-01-05T18:29:50.9613712Z shell: /usr/bin/bash -e ***0***
2025-01-05T18:29:50.9614004Z env:
2025-01-05T18:29:50.9614263Z   PROJECT_ID: ***
2025-01-05T18:29:50.9614513Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:29:50.9614784Z   GKE_ZONE: europe-north1-b
2025-01-05T18:29:50.9615047Z   BRANCH: main
2025-01-05T18:29:50.9615591Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:50.9616447Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:50.9617250Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:29:50.9617816Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:29:50.9618345Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:29:50.9618643Z   GCLOUD_PROJECT: ***
2025-01-05T18:29:50.9618914Z   GCP_PROJECT: ***
2025-01-05T18:29:50.9619200Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:29:50.9619546Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-05T18:29:50.9619940Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-05T18:29:50.9620467Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:50.9621200Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:29:50.9621694Z ##[endgroup]
2025-01-05T18:29:51.4997820Z The push refers to repository [europe-north1-docker.pkg.dev/***/kubernetes-course/crud-backend]
2025-01-05T18:29:51.7233857Z ebedfaddd57a: Preparing
2025-01-05T18:29:51.7236775Z 197d6bcdc63d: Preparing
2025-01-05T18:29:51.7237112Z 6cdd0047cf36: Preparing
2025-01-05T18:29:51.7237392Z 4b4be28e5339: Preparing
2025-01-05T18:29:51.7237642Z 5f70bf18a086: Preparing
2025-01-05T18:29:51.7237880Z 4e061853116c: Preparing
2025-01-05T18:29:51.7238324Z f3fe8ef9af87: Preparing
2025-01-05T18:29:51.7238577Z c3c4291cf0a2: Preparing
2025-01-05T18:29:51.7239151Z 85c6f0cfb532: Preparing
2025-01-05T18:29:51.7239395Z a4fd1e7df47e: Preparing
2025-01-05T18:29:51.7254110Z 2f7b6d216a37: Preparing
2025-01-05T18:29:51.7254538Z f3fe8ef9af87: Waiting
2025-01-05T18:29:51.7254937Z c3c4291cf0a2: Waiting
2025-01-05T18:29:51.7255334Z 85c6f0cfb532: Waiting
2025-01-05T18:29:51.7255695Z a4fd1e7df47e: Waiting
2025-01-05T18:29:51.7256058Z 2f7b6d216a37: Waiting
2025-01-05T18:29:51.7256442Z 4e061853116c: Waiting
2025-01-05T18:29:52.2263867Z 5f70bf18a086: Layer already exists
2025-01-05T18:29:52.4745681Z 4e061853116c: Layer already exists
2025-01-05T18:29:52.7228830Z f3fe8ef9af87: Layer already exists
2025-01-05T18:29:52.9639040Z c3c4291cf0a2: Layer already exists
2025-01-05T18:29:53.0387682Z 4b4be28e5339: Layer already exists
2025-01-05T18:29:54.0313805Z 85c6f0cfb532: Layer already exists
2025-01-05T18:29:54.0948695Z a4fd1e7df47e: Layer already exists
2025-01-05T18:29:55.0920963Z 2f7b6d216a37: Layer already exists
2025-01-05T18:29:57.2792114Z ebedfaddd57a: Pushed
2025-01-05T18:29:57.3390775Z 6cdd0047cf36: Pushed
2025-01-05T18:29:57.7309804Z 197d6bcdc63d: Pushed
2025-01-05T18:30:04.4235467Z main-ee29c5b4f7135cc0ed7ad009ad5929306216d337: digest: sha256:0a00b67093ed37455298272f0deb4f30484797feb40e2924f041ec64b7f02de7 size: 2621
2025-01-05T18:30:04.4306309Z ##[group]Run docker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend-worker:$BRANCH-$GITHUB_SHA
2025-01-05T18:30:04.4307347Z [36;1mdocker push europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend-worker:$BRANCH-$GITHUB_SHA[0m
2025-01-05T18:30:04.4335380Z shell: /usr/bin/bash -e ***0***
2025-01-05T18:30:04.4335884Z env:
2025-01-05T18:30:04.4336330Z   PROJECT_ID: ***
2025-01-05T18:30:04.4336771Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:30:04.4337256Z   GKE_ZONE: europe-north1-b
2025-01-05T18:30:04.4337732Z   BRANCH: main
2025-01-05T18:30:04.4338665Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:04.4339577Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:04.4340408Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:04.4340988Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:30:04.4341293Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:30:04.4341576Z   GCLOUD_PROJECT: ***
2025-01-05T18:30:04.4341850Z   GCP_PROJECT: ***
2025-01-05T18:30:04.4342136Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:30:04.4342476Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-05T18:30:04.4342868Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-05T18:30:04.4343388Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:30:04.4344113Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:30:04.4344634Z ##[endgroup]
2025-01-05T18:30:04.9587936Z The push refers to repository [europe-north1-docker.pkg.dev/***/kubernetes-course/crud-backend-worker]
2025-01-05T18:30:05.1768472Z ebedfaddd57a: Preparing
2025-01-05T18:30:05.1769056Z 197d6bcdc63d: Preparing
2025-01-05T18:30:05.1771601Z 6cdd0047cf36: Preparing
2025-01-05T18:30:05.1773292Z 4b4be28e5339: Preparing
2025-01-05T18:30:05.1776412Z 5f70bf18a086: Preparing
2025-01-05T18:30:05.1776904Z 4e061853116c: Preparing
2025-01-05T18:30:05.1777407Z f3fe8ef9af87: Preparing
2025-01-05T18:30:05.1777840Z c3c4291cf0a2: Preparing
2025-01-05T18:30:05.1778469Z 85c6f0cfb532: Preparing
2025-01-05T18:30:05.1778749Z a4fd1e7df47e: Preparing
2025-01-05T18:30:05.1779026Z 2f7b6d216a37: Preparing
2025-01-05T18:30:05.1779285Z 4e061853116c: Waiting
2025-01-05T18:30:05.1779529Z f3fe8ef9af87: Waiting
2025-01-05T18:30:05.1779766Z c3c4291cf0a2: Waiting
2025-01-05T18:30:05.1780013Z 85c6f0cfb532: Waiting
2025-01-05T18:30:05.1780250Z a4fd1e7df47e: Waiting
2025-01-05T18:30:05.1780793Z 2f7b6d216a37: Waiting
2025-01-05T18:30:05.6863715Z 5f70bf18a086: Layer already exists
2025-01-05T18:30:06.4827597Z 6cdd0047cf36: Layer already exists
2025-01-05T18:30:06.4877918Z ebedfaddd57a: Layer already exists
2025-01-05T18:30:06.4931626Z 197d6bcdc63d: Layer already exists
2025-01-05T18:30:06.5138555Z 4b4be28e5339: Layer already exists
2025-01-05T18:30:06.7509646Z 4e061853116c: Layer already exists
2025-01-05T18:30:07.5342326Z c3c4291cf0a2: Layer already exists
2025-01-05T18:30:07.5351541Z f3fe8ef9af87: Layer already exists
2025-01-05T18:30:07.5490544Z 85c6f0cfb532: Layer already exists
2025-01-05T18:30:07.5686446Z a4fd1e7df47e: Layer already exists
2025-01-05T18:30:07.8430112Z 2f7b6d216a37: Layer already exists
2025-01-05T18:30:13.6555575Z main-ee29c5b4f7135cc0ed7ad009ad5929306216d337: digest: sha256:e29fd16bc06cb2b0f59c82fd4493523e7711f8fab95b5db0fcb4589f950e7189 size: 2621
2025-01-05T18:30:13.6615973Z ##[group]Run kubectl config set-context --current --namespace=crud
2025-01-05T18:30:13.6616491Z [36;1mkubectl config set-context --current --namespace=crud[0m
2025-01-05T18:30:13.6643584Z shell: /usr/bin/bash -e ***0***
2025-01-05T18:30:13.6643873Z env:
2025-01-05T18:30:13.6644124Z   PROJECT_ID: ***
2025-01-05T18:30:13.6644374Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:30:13.6644647Z   GKE_ZONE: europe-north1-b
2025-01-05T18:30:13.6644918Z   BRANCH: main
2025-01-05T18:30:13.6645465Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:13.6646330Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:13.6647129Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:13.6647689Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:30:13.6648221Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:30:13.6648514Z   GCLOUD_PROJECT: ***
2025-01-05T18:30:13.6648782Z   GCP_PROJECT: ***
2025-01-05T18:30:13.6649074Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:30:13.6649407Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-05T18:30:13.6649795Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-05T18:30:13.6650317Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:30:13.6651021Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:30:13.6651515Z ##[endgroup]
2025-01-05T18:30:14.2906651Z Context "gke_***_europe-north1-b_dwk-cluster" modified.
2025-01-05T18:30:14.2947675Z ##[group]Run cd crud-app
2025-01-05T18:30:14.2948258Z [36;1mcd crud-app[0m
2025-01-05T18:30:14.2948875Z [36;1mkustomize edit set image crud/app=europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-app:$BRANCH-$GITHUB_SHA[0m
2025-01-05T18:30:14.2949512Z [36;1mkustomize build . | kubectl apply -f -[0m
2025-01-05T18:30:14.2949857Z [36;1mkubectl rollout status deployment[0m
2025-01-05T18:30:14.2950194Z [36;1mkubectl get services -o wide[0m
2025-01-05T18:30:14.2976312Z shell: /usr/bin/bash -e ***0***
2025-01-05T18:30:14.2976608Z env:
2025-01-05T18:30:14.2976868Z   PROJECT_ID: ***
2025-01-05T18:30:14.2977118Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:30:14.2977386Z   GKE_ZONE: europe-north1-b
2025-01-05T18:30:14.2977644Z   BRANCH: main
2025-01-05T18:30:14.2978467Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:14.2979386Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:14.2980177Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:14.2980737Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:30:14.2981046Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:30:14.2981337Z   GCLOUD_PROJECT: ***
2025-01-05T18:30:14.2981606Z   GCP_PROJECT: ***
2025-01-05T18:30:14.2982083Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:30:14.2982419Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-05T18:30:14.2982811Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-05T18:30:14.2983334Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:30:14.2984042Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:30:14.2984534Z ##[endgroup]
2025-01-05T18:30:20.4092408Z service/crud-app-svc unchanged
2025-01-05T18:30:20.8645957Z deployment.apps/crud-app configured
2025-01-05T18:30:21.0837571Z ingress.networking.k8s.io/crud-app-ingress unchanged
2025-01-05T18:30:22.4272571Z Waiting for deployment "crud-app" rollout to finish: 1 old replicas are pending termination...
2025-01-05T18:30:56.9526354Z Waiting for deployment "crud-app" rollout to finish: 1 old replicas are pending termination...
2025-01-05T18:30:57.0153034Z deployment "crud-app" successfully rolled out
2025-01-05T18:30:58.3011243Z NAME                TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE     SELECTOR
2025-01-05T18:30:58.3012155Z crud-app-svc        NodePort    34.118.237.4   <none>        2347:32035/TCP   24m     app=crud-app
2025-01-05T18:30:58.3013135Z crud-postgres-svc   ClusterIP   None           <none>        5432/TCP         3h13m   app=crud-postgres
2025-01-05T18:30:58.3070764Z ##[group]Run cd crud-backend
2025-01-05T18:30:58.3071346Z [36;1mcd crud-backend[0m
2025-01-05T18:30:58.3072398Z [36;1mkustomize edit set image crud/backend=europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend:$BRANCH-$GITHUB_SHA[0m
2025-01-05T18:30:58.3074176Z [36;1mkustomize edit set image crud/backend-worker=europe-north1-docker.pkg.dev/$PROJECT_ID/kubernetes-course/crud-backend-worker:$BRANCH-$GITHUB_SHA[0m
2025-01-05T18:30:58.3075413Z [36;1mkustomize build . | kubectl apply -f -[0m
2025-01-05T18:30:58.3076040Z [36;1mkubectl rollout status deployment[0m
2025-01-05T18:30:58.3076627Z [36;1mkubectl get services -o wide[0m
2025-01-05T18:30:58.3114726Z shell: /usr/bin/bash -e ***0***
2025-01-05T18:30:58.3115243Z env:
2025-01-05T18:30:58.3115700Z   PROJECT_ID: ***
2025-01-05T18:30:58.3116152Z   GKE_CLUSTER: dwk-cluster
2025-01-05T18:30:58.3116653Z   GKE_ZONE: europe-north1-b
2025-01-05T18:30:58.3117128Z   BRANCH: main
2025-01-05T18:30:58.3118286Z   CLOUDSDK_AUTH_CREDENTIAL_FILE_OVERRIDE: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:58.3119847Z   GOOGLE_APPLICATION_CREDENTIALS: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:58.3121281Z   GOOGLE_GHA_CREDS_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json
2025-01-05T18:30:58.3122326Z   CLOUDSDK_CORE_PROJECT: ***
2025-01-05T18:30:58.3122868Z   CLOUDSDK_PROJECT: ***
2025-01-05T18:30:58.3123371Z   GCLOUD_PROJECT: ***
2025-01-05T18:30:58.3123854Z   GCP_PROJECT: ***
2025-01-05T18:30:58.3124369Z   GOOGLE_CLOUD_PROJECT: ***
2025-01-05T18:30:58.3124963Z   CLOUDSDK_METRICS_ENVIRONMENT: github-actions-setup-gcloud
2025-01-05T18:30:58.3125676Z   CLOUDSDK_METRICS_ENVIRONMENT_VERSION: 2.1.2
2025-01-05T18:30:58.3126579Z   KUBECONFIG: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:30:58.3127876Z   KUBE_CONFIG_PATH: /home/runner/work/kubernetes-course/kubernetes-course/gha-kubeconfig-8f2a6a0d132181af
2025-01-05T18:30:58.3129093Z ##[endgroup]
2025-01-05T18:31:01.0335250Z service/crud-backend-svc created
2025-01-05T18:31:01.6025307Z deployment.apps/crud-backend created
2025-01-05T18:31:02.2750780Z cronjob.batch/crud-backend-worker created
2025-01-05T18:31:02.8341456Z ingress.networking.k8s.io/crud-backend-ingress created
2025-01-05T18:31:04.1361100Z deployment "crud-app" successfully rolled out
2025-01-05T18:31:04.4951402Z Waiting for deployment "crud-backend" rollout to finish: 0 of 1 updated replicas are available...
2025-01-05T18:31:27.9673549Z deployment "crud-backend" successfully rolled out
2025-01-05T18:31:29.1247942Z NAME                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE     SELECTOR
2025-01-05T18:31:29.1253847Z crud-app-svc        NodePort    34.118.237.4     <none>        2347:32035/TCP   25m     app=crud-app
2025-01-05T18:31:29.1254846Z crud-backend-svc    NodePort    34.118.225.158   <none>        2348:30398/TCP   29s     app=crud-backend
2025-01-05T18:31:29.1255849Z crud-postgres-svc   ClusterIP   None             <none>        5432/TCP         3h13m   app=crud-postgres
2025-01-05T18:31:29.1336307Z Post job cleanup.
2025-01-05T18:31:29.2188850Z Removed exported credentials at "/home/runner/work/kubernetes-course/kubernetes-course/gha-creds-e42b5d7cd011977b.json".
2025-01-05T18:31:29.2285118Z Post job cleanup.
2025-01-05T18:31:29.3256047Z [command]/usr/bin/git version
2025-01-05T18:31:29.3318659Z git version 2.47.1
2025-01-05T18:31:29.3360147Z Temporarily overriding HOME='/home/runner/work/_temp/dc4303b8-cd47-45f6-88e0-86015bef962c' before making global git config changes
2025-01-05T18:31:29.3365812Z Adding repository directory to the temporary git global config as a safe directory
2025-01-05T18:31:29.3379861Z [command]/usr/bin/git config --global --add safe.directory /home/runner/work/kubernetes-course/kubernetes-course
2025-01-05T18:31:29.3457229Z [command]/usr/bin/git config --local --name-only --get-regexp core\.sshCommand
2025-01-05T18:31:29.3518643Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'core\.sshCommand' && git config --local --unset-all 'core.sshCommand' || :"
2025-01-05T18:31:29.3889298Z [command]/usr/bin/git config --local --name-only --get-regexp http\.https\:\/\/github\.com\/\.extraheader
2025-01-05T18:31:29.3938864Z http.https://github.com/.extraheader
2025-01-05T18:31:29.3948678Z [command]/usr/bin/git config --local --unset-all http.https://github.com/.extraheader
2025-01-05T18:31:29.4008602Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'http\.https\:\/\/github\.com\/\.extraheader' && git config --local --unset-all 'http.https://github.com/.extraheader' || :"
2025-01-05T18:31:29.4402147Z Cleaning up orphan processes
```
