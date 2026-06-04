
テンプレートの作成方法

```
mkdir -p my-app && cd my-app

# SvelteKit のインストール
npx sv create frontend --template minimal --types ts --no-add-ons --install npm

# PocketBase の JS SDK を追加
cd frontend
npm install pocketbase
```
