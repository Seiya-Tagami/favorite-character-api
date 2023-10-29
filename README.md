# favorite-character-api

この API は、いわゆるアニメの”推し”を簡易的に保存するための API です。キャラクターの名前と属するアニメを保存できます。

## ベース URL

開発環境では、`http://localhost:8080`です。

## エンドポイント

### 全推しキャラの取得

- **URI**: `/characters`
- **method**: GET
- **リクエストの例**:

```sh
curl -X GET https://localhost:8080/characters
```

- **レスポンス**:
  - `200 OK`: 正常に取得
  - `500 Internal Server Error`: サーバ側の異常

### 条件付き推しキャラの取得

- **URI**: `/characters/:id`
- **method**: GET
- **リクエストの例**:

```sh
curl -X GET https://localhost:8080/characters/1
```

- **レスポンス**:
  - `200 OK`: 正常に取得
  - `500 Internal Server Error`: サーバ側の異常

### 推しキャラの追加

- **URI**: `/characters`
- **method**: POST
- **body**:
  - `name` (文字列): キャラクターの名前。
  - `belonging` (文字列): キャラクターが所属するアニメ
- **リクエストの例**:

```sh
curl -X POST -H "Content-Type: application/json" -d '{"name":"後藤ひとり","belonging":"ぼっちざろっく"}' http://localhost:8080/characters

```

- **レスポンス**:
  - `200 OK`: 正常にリソースを作成
  - `400 Bad Request`: クライアント側の不正
  - `500 Internal Server Error`: サーバ側の異常

### 推しキャラの更新

- **URI**: `/characters/:id`
- **method**: PATCH
- **body**:
  - `name` (文字列): キャラクターの名前。
  - `belonging` (文字列): キャラクターが所属するアニメ
- **リクエストの例**:

```sh
curl -X PATCH -H "Content-Type: application/json" -d '{"name":"山田りょう","belonging":"ぼっちざろっく"}' http://localhost:8080/characters/1

```

- **レスポンス**:
  - `200 OK`: 正常にリソースを更新
  - `400 Bad Request`: クライアント側の不正
  - `500 Internal Server Error`: サーバ側の異常

### 推しキャラの削除

- **URI**: `/characters/:id`
- **method**: DELETE
- **リクエストの例**:

```sh
curl -X POST https://localhost:8080/characters/1
```

- **レスポンス**:
  - `204 No Content`: 正常にリソースを削除し、何も返さない
  - `500 Internal Server Error`: サーバ側の異常
