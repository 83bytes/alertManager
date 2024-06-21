- Use https://github.com/alecthomas/kong for config management (or use [flag](https://pkg.go.dev/flag) from the STL)
  - CLI params
    - --server-port
    - --pprof-port
    - --alerting-config
    - --auth-config
- Use [net/http]
  - needs 2 min endpoints
    - POST / (for getting webhook)
    - POST /~/reload to reload config
    - GET /~/config to get current loaded config
    - GET /health
- Use logrus for logging
- Tests

# Tol

## app cli style

alertmanager server --port abcd --pprof-port efgh --config /path-to-file --loglevel

alertmanager config generate-template --out-file
alertmanager config validate --config-file

## Flow

Alert -> Enrichment(s) -> Action(s)
