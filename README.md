# n8n Automation Stack — Docker

A self-hosted automation stack built around [n8n](https://n8n.io/), containerised with Docker Compose. Demonstrates practical workflow automation with a vector database, in-memory cache, and a custom web frontend served by Caddy.

**Stack:** n8n · Caddy · Qdrant · Redis · Docker Compose · HTMX

---

## What's in the box

| Service | Purpose |
|---|---|
| **n8n** | Visual workflow automation — triggers, AI nodes, webhooks |
| **Caddy** | Reverse proxy + static file server for the custom frontend |
| **Qdrant** | Vector database for semantic search / RAG pipelines |
| **Redis** | In-memory cache and queue backend |

### Workflows included

- **`rag.json`** — RAG (Retrieval-Augmented Generation) pipeline: accepts file uploads via a form trigger, splits and embeds documents into Qdrant, then answers questions against the stored vectors.
- **`textbook-assistant.json` / `textbook-assistant-gpt40.json`** — Textbook Q&A assistant variants (standard and GPT-4o).
- **`maybe-parallel.json`** — Experimental parallel execution workflow.

### Frontend

A lightweight HTML/HTMX interface served by Caddy (`interface/site/`) with:
- File upload form (feeds into the n8n RAG webhook)
- Chat panel (streams responses from n8n back to the browser)

Caddy is configured with its template engine enabled, so HTML partials (`includes/`) are composed server-side.

---

## Running it

```bash
# 1. Clone the repo
git clone https://github.com/JAC-1/n8n-caddy-docker
cd n8n-caddy-docker

# 2. (Optional) create a .env file
cp .env.example .env   # set N8N_USER, N8N_PASSWORD, etc.

# 3. Create the external volume (keeps your workflows safe across restarts)
docker volume create n8n-docker_n8n_data

# 4. Start everything
docker compose up -d
```

| Service | URL |
|---|---|
| n8n editor | http://localhost:5678 |
| Frontend | http://localhost |
| Qdrant dashboard | http://localhost:6333/dashboard |

---

## Useful commands

```bash
docker compose pull          # pull latest images
docker compose up -d         # (re)start in background
docker compose down          # stop all containers
docker compose logs -f n8n   # tail n8n logs
```

---

## Project layout

```
.
├── docker-compose.yml        # all four services
├── interface/
│   ├── conf/Caddyfile        # Caddy config (templates + file server)
│   └── site/                 # frontend HTML + HTMX partials
├── rag.json                  # n8n RAG workflow export
├── textbook-assistant.json   # n8n textbook assistant workflow
└── maybe-parallel.json       # experimental parallel workflow
```

---

## Environment variables

| Variable | Default | Description |
|---|---|---|
| `N8N_USER` | `admin` | n8n basic auth username |
| `N8N_PASSWORD` | `change_me_strong` | n8n basic auth password |
| `N8N_PROTOCOL` | `http` | `http` or `https` |

---

---

# n8n 自動化スタック — Docker

[n8n](https://n8n.io/) を中心に構築した自己ホスト型自動化スタックです。Docker Compose でコンテナ化し、ベクターデータベース・インメモリキャッシュ・Caddy で配信するカスタム Web フロントエンドを組み合わせています。

**技術スタック：** n8n · Caddy · Qdrant · Redis · Docker Compose · HTMX

---

## 構成

| サービス | 役割 |
|---|---|
| **n8n** | ビジュアルワークフロー自動化 — トリガー・AI ノード・Webhook |
| **Caddy** | リバースプロキシ + カスタムフロントエンド用静的ファイルサーバー |
| **Qdrant** | セマンティック検索 / RAG パイプライン用ベクターデータベース |
| **Redis** | インメモリキャッシュおよびキューバックエンド |

### 収録ワークフロー

- **`rag.json`** — RAG（検索拡張生成）パイプライン：フォームトリガーでファイルを受け取り、ドキュメントを分割して Qdrant に埋め込み、格納済みベクターに対して質問に回答します。
- **`textbook-assistant.json` / `textbook-assistant-gpt40.json`** — 教科書 Q&A アシスタント（標準版と GPT-4o 版）。
- **`maybe-parallel.json`** — 並列実行の実験的ワークフロー。

### フロントエンド

Caddy が配信する軽量 HTML/HTMX インターフェース（`interface/site/`）：
- ファイルアップロードフォーム（n8n の RAG Webhook に送信）
- チャットパネル（n8n からブラウザへレスポンスをストリーミング）

Caddy のテンプレートエンジンを有効化しており、HTML パーシャル（`includes/`）はサーバーサイドで合成されます。

---

## 起動方法

```bash
# 1. リポジトリをクローン
git clone https://github.com/JAC-1/n8n-caddy-docker
cd n8n-caddy-docker

# 2. （任意）.env ファイルを作成
cp .env.example .env   # N8N_USER、N8N_PASSWORD などを設定

# 3. 外部ボリュームを作成（再起動後もワークフローを保持）
docker volume create n8n-docker_n8n_data

# 4. すべて起動
docker compose up -d
```

| サービス | URL |
|---|---|
| n8n エディター | http://localhost:5678 |
| フロントエンド | http://localhost |
| Qdrant ダッシュボード | http://localhost:6333/dashboard |

---

## よく使うコマンド

```bash
docker compose pull          # 最新イメージを取得
docker compose up -d         # バックグラウンドで（再）起動
docker compose down          # 全コンテナを停止
docker compose logs -f n8n   # n8n のログをリアルタイム表示
```

---

## 環境変数

| 変数 | デフォルト | 説明 |
|---|---|---|
| `N8N_USER` | `admin` | n8n ベーシック認証のユーザー名 |
| `N8N_PASSWORD` | `change_me_strong` | n8n ベーシック認証のパスワード |
| `N8N_PROTOCOL` | `http` | `http` または `https` |
