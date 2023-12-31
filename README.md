# Mackerunn

"Mackerunn" executes scenario monitoring using [runn](https://github.com/k1LoW/runn) and sends the results to Mackerel.

## Usage

### mackerunn

```console
$ export MACKERUNN_MACKEREL_APIKEY="your_api_key"
$ go run ./cmd/mackerunn -runbook testdata/test.yml -service test -hostID 12345ABCD
2023/11/23 16:13:23 succeeded
```

### mackerunnd

Execute mackerunn for every minute.

```console
$ export MACKERUNN_MACKEREL_APIKEY="your_api_key"
$ go run ./cmd/mackerunnd -runbook testdata/test.yml -service test -hostID 12345ABCD
2023/11/23 16:13:23 succeeded
2023/11/23 16:14:23 succeeded
2023/11/23 16:15:23 succeeded
2023/11/23 16:16:23 succeeded
2023/11/23 16:17:23 succeeded
```
