
## Init
プロジェクトは[テンプレートリポジトリ](https://gitea.internal/tmpl/sveltekit-pocketbase-tmpl)を作っておくのがシンプル。

teaでgiteaにlogin
```sh
export GITEA_SERVER_URL=https://gitea.internal
export GITEA_SERVER_TOKEN=$(sudo -A cat /run/secrets/tea_api_token)
tea login add  --name gitea.internal
tea login default gitea.internal
```

repositoryの作成
```sh
export name=$(date +%y%m%d)-my-app
export org=test
export root_dir=$HOME/test
export repo_dir=$root_dir/$name
export tmpl=tmpl/sveltekit-pocketbase-tmpl

mkdir -p $repo_dir && cd $repo_dir
tea repos create-from-template --login gitea.internal --template $tmpl --name $name --owner $org
git clone $org/$name $repo_dir

cd $repo_dir
make init
```

## Setup

```sh
# 1. Create project with official Solid template
npm create vite@latest frontend -- --template solid


# 2. Add PocketBase + prettier + eslint
cd frontend
npm install @solidjs/router pocketbase
npm install --save-dev prettier eslint eslint-plugin-solid

# 3. Static build (no adapter needed — Vite outputs static by default)


# 以下のファイルを編集
# vite.config.js
# .env
# src/lib/pocketbase.js
# src/routes/+layout.js
```

pure-js+prettier, eslint


---
Wiki:
=> https://gitea.internal/tmpl/sveltekit-pocketbase-tmpl/wiki
=> https://gollum.app.internal/pocketbase/Home

Komodo:
=> https://komodo.app.internal/actions/6a14306c936ff851f0261730
