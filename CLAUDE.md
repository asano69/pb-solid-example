## Overview


# Rules

- データベースのマイグレーションはフロントエンドから行うのでマイグレーションコードを作成する必要はまったくない。
- When fixing bugs, add a failing regression test first.
- All errors are user-facing, so messages should be clear.
- Keep functions small and focused.
- Module files should re-export what's needed, hide implementation details.
- Don't persist changes to the database during drilling. Use the cache.
- Don't use timezones: dates are naive for a reason. Due dates etc. are more like the dates in a journal entry than precise points in time.


# Work in progress

## Backend 
- PocketBase **v0.39+** 

## Frontend
- frontendをsolid.js + **tailwind v4**
