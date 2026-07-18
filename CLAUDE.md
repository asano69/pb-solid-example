# Overview

- このプロジェクトの目的は、汎用的な目的に使えるテンプレートとなるリポジトリをつくることです。
- このアプリは、SPA＋CSRの個人向けのアプリです。

## Rules

- 後方互換性は維持しなくてよい。
- データベースのマイグレーションはPocketBaseのWEB UIから行うのでマイグレーションコードを作成する必要はない。
- When fixing bugs, add a failing regression test first.
- All errors are user-facing, so messages should be clear.
- Keep functions small and focused.
- Module files should re-export what's needed, hide implementation details.


## Tech Stack

- backend: Go + PocketBase **v0.39+**
- frontend: solid.js + **tailwind v4**


## Work in progress
