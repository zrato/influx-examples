
Influxd can not be running when **influxd generate** is called...

```
influxd generate simple --help
infdgen help-schema
infdgen simple --print
infdgen simple
infdgen simple --t=1,1,1
infdgen simple --t=1,1,1 --p=15
infdgen simple --t=2,1,1 --p=5
```

----------------------------------

infdgen simple --t=1,2,1 --p=5
1 measurement, 2 tags, 1 field

Tag cardinalities        [1,2,1]
Points per series        5
Total points             10
Total series             2

cardinality = 2

----------------------------------

infdgen simple --t=1,2,2 --p=5
1 measurement, 2 tags, 2 fields

Tag cardinalities        [1,2,2]
Points per series        5
Total points             20
Total series             4

cardinality = 4

----------------------------------

infdgen simple --t=2,2,2 --p=5
2 measurements, 2 tags, 2 fields

Tag cardinalities        [2,2,2]
Points per series        5
Total points             40
Total series             8

cardinality = 8

----------------------------------

```
infdgen schema02.toml --print
infdgen schema02.toml --print --clean all
infdgen schema02.toml --clean all
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
