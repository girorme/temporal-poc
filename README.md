Temporal POC
---
Project to test [temporal.io](https://temporal.io/) event sourcing

## Pre reqs
- [Setup temporal and golang](https://learn.temporal.io/getting_started/go/dev_environment/)

---

### Start temporal cluster
```bash
$ temporal server start-dev --db-filename database/temporal.db
```

### Run worker
```bash
$ go run main.go
```

### Start workflow
```bash
$ temporal workflow start --type Greet --task-queue greeting-tasks --workflow-id my-first-workflow --input '"Girorme"'
```