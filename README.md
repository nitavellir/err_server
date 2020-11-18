# エラーを返すサーバ

* [maizy/errserv](https://github.com/maizy/errserv) の改良版
* ServeHTTP処理のタイムアウトを設定可能に

# How to Use

* Main関数を呼び出してサーバ起動（example/errserv.go参照）

* リクエストのポート、エンドポイントはエラーステータス
  * リクエスト例
  <pre>
  curl localhost:10405/10405
  curl localhost:10500/10500
  </pre>
  
* 対象エラーステータス
<pre>
http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden,
		http.StatusNotFound, http.StatusMethodNotAllowed, http.StatusNotAcceptable,
		http.StatusGone, http.StatusLengthRequired, http.StatusPreconditionFailed,
		http.StatusRequestEntityTooLarge, http.StatusRequestURITooLong, http.StatusUnsupportedMediaType,
		http.StatusRequestedRangeNotSatisfiable,

		http.StatusInternalServerError, http.StatusNotImplemented, http.StatusBadGateway,
		http.StatusServiceUnavailable, http.StatusGatewayTimeout
</pre>
