
### Storage and Tsdb

The main entry point for tsm1 and tsi1 when influxd is running is via
storage/Engine.go.

The cmd/influxd/generate facility writes directly to tsdb but independently
of the above code path.

[See the codenotes for specific references](./codenotes.md)
