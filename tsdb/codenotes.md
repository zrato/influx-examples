
### Tsdb Notes

All writes to **tsdb** happens through generator.go which talks to internal/shard/writer.go

```
rg \/tsdb
```

And inside pkg/data/gen we have these references...

* arrays.gen.go
* values_sequence.gen.go
* specs.go
* merged_series_generator_test.go

##### Besides pkg/data/gen tsdb is globally referenced here...

* cmd/influxd
* gather
* mock
* predicate
* query
* storage

### Tsm1 Notes

tsm1 is not found or referenced anywhere in the tsdb package

```
rg tsm1 *.go
```

These are the core locations of who referenes tsm1

storage/config.go
storage/engine.go
storage/retention.go

### Tsi1 Notes

tsi1 is only found in one place in the tsdb package and that is index_test.go

```
rg tsi1 *.go
```

```
rg \/tsdb\/tsi1
```

tsi1 is called mainly in storage and cmd and...

tsdb/tsm1/engine.go

No reference to the SeriesFile inside the tsm1.

### storage package

storage.NewEngine

```
rg 'storage\.NewEngine'
```

tsdb/series_file.go
Opening Series File

See here specifically to storage...

rg \/tsm1 *.go
config.go
engine.go
retention.go

rg \/tsi1 *.go
config.go
engine.go
series_cursor.go
