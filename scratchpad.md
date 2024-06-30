- Use https://github.com/alecthomas/kong for config management (or use [flag](https://pkg.go.dev/flag) from the STL)
  - CLI params
    - --server-port
    - --alerting-config
    - --auth-config
- Use [net/http]
  - needs 2 min endpoints
    - POST /webhook (for getting webhook)
    - POST /~/reload to reload config
    - GET /~/config to get current loaded config
    - GET /ping
- Use logrus for logging
- Tests

# Tol

## app cli style

alertmanager server --port abcd --pprof-port efgh --config /path-to-file --loglevel

alertmanager config generate-template > outfile
alertmanager config validate --config-file

## Flow

Alert -> Enrichment(s) -> Action(s)

## OverAll Plan

- I load the config to see

# promQL enrichment

process this query for now

```
curl 'http://localhost:9090/api/v1/query?query=sum%28rate%28node_cpu_seconds_total%5B1m%5D%29%29&time=1719778255.812'
```
