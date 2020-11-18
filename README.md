# エラーを返すサーバ

* [maizy/errserv](https://github.com/maizy/errserv) の改良版
* ServeHTTP処理のタイムアウトを設定可能に

# How to Use

* Main関数を呼び出してサーバ起動（example/errserv.go参照）

* リクエストのポート、エンドポイントはエラーのステータスコード
  * リクエスト例
  <pre>
  curl localhost:10405/10405
  curl localhost:10500/10500
  </pre>
  
* 対象ステータスコード
<pre>
400	http.StatusBadRequest
401	http.StatusUnauthorized
403	http.StatusForbidden
404	http.StatusNotFound
405	http.StatusMethodNotAllowed
406	http.StatusNotAcceptable
410	http.StatusGone
411	http.StatusLengthRequired
412	http.StatusPreconditionFailed
413	http.StatusRequestEntityTooLarge
414	http.StatusRequestURITooLong
415	http.StatusUnsupportedMediaType
416	http.StatusRequestedRangeNotSatisfiable
500	http.StatusInternalServerError
501	http.StatusNotImplemented
502	http.StatusBadGateway
503	http.StatusServiceUnavailable
504	http.StatusGatewayTimeout
</pre>

* ServeHTTPのcallのtimeoutは、Main関数の第1引数の数値
  * ナノ秒単位
  * 任意指定
