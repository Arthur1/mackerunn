# Mackerunn

"Mackerunn" executes scenario monitoring using [runn](https://github.com/k1LoW/runn) and sends the results to Mackerel.

## Usage

### mackerunn

```console
$ export MACKERUNN_MACKEREL_APIKEY="your_api_key"
$ go run ./cmd/mackerunn -runbook testdata/test.yml -service test -hostID 12345ABCD
2023/11/23 16:13:23 succeeded
```
