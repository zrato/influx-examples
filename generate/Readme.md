
[influxd generate](https://v2.docs.influxdata.com/v2.0/reference/cli/influxd/generate)

The <schema.toml> files are located in the schema directory.

```
infdgen schema01.toml
```

where infdgen is

```
alias infdgen='influxd generate --org ag --bucket rick'
```

```
influxd generate simple --bucket rick --org ag --print
influxd generate simple --bucket rick --org ag
```

Finds all places in the code where schemas are tested

```
rg '\[\[measurements\]\]'
```

##### Resultset

storage/reads/store.go

type ResultSet interface {
