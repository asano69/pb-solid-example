## Init
プロジェクトは[テンプレートリポジトリ](https://gitea.internal/tmpl/solid-pocketbase-tmpl)を作っておくのがシンプル。

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
export tmpl=tmpl/solid-pocketbase-tmpl

mkdir -p $repo_dir && cd $repo_dir
tea repos create-from-template --login gitea.internal --template $tmpl --name $name --owner $org
git clone git@gitea.internal:$org/$name $repo_dir

cd $repo_dir
make init
```
## 構成
### frontend
src/utils
- UIに関係しない汎用ロジック、 自分で書いた小さなヘルパー関数の置き場。可能であれば、同じ入力には必ず同じ出力を返し副作用のない関数（＝純粋関数）をおくことが好ましい。
- SvelteKitではlibが使われるがreactではutilsが使われる。
- JSXをおかず拡張子は、JSまたはTSになるはず。

src/components/
- 再利用可能なUIパーツ。特定のページに依存しないもの：Button.jsx, Modal.jsx, Avatar.jsx

src/pages/（または views/）
- ルートに対応する画面単位のコンポーネント。
- Next.js/SvelteKitなどファイルベースルーティングのフレームワークではフレームワーク側が管理するので自分で作らないことも多い

src/auth
- 認証関連のコンポーネント、関数置き場。LoginForm.jsx、useAuth.js、authApi.js

src/data
- 静的データ置き場


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


---
参考：
- https://github.com/TomaszSkrzyp/good-game
- https://github.com/swerve731/go-solid-todo-app/tree/main/frontend
- https://github.com/OpenListTeam/OpenList-Frontend/tree/main/src

---
Wiki:
=> https://gitea.internal/tmpl/sveltekit-pocketbase-tmpl/wiki
=> https://gollum.app.internal/pocketbase/Home

Komodo:
=> https://komodo.app.internal/actions/6a14306c936ff851f0261730



