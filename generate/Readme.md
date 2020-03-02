
Influxd can not be running when **influxd generate** is called...

```
infdgen simple --print
infdgen simple
infdgen schema01.toml
```

Random values get generated to build up a series file.    
Here are the core tenants to how it works.

[influxd generate](https://v2.docs.influxdata.com/v2.0/reference/cli/influxd/generate)

The <schema.toml> files are located in the schema directory.

```
infdgen schema01.toml
```

where infdgen is

```
alias infdgen='influxd generate --org ag --bucket rick'
```
