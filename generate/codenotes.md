
Random values get generated to build up a series file.    
Here are the core tenants to how it works.

```
rg math\/rand
```

### Everything happens in specs.go visit

Inside exec of command.go NewSeriesGeneratorFromSpec gets
called and this is what drives the start of the **generator**

```
	sg := gen.NewSeriesGeneratorFromSpec(spec, tr)
```

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

##### Generate simple data sets using only CLI flags

cmd/influxd/generate/command_simple.go

##### How does this work ?

```
rg AddFlags
```

So this shows you that

People write schemas using toml and then specs get generated from
the toml files which represent the schema...

Schema --> Spec

**exec gets called** on both command.go and command_simple.go

##### See where the pkg/data/gen code is actually used ?

```
rg \/data\/gen
```

* cmd/influxd/generate
* mock
* storage/reads

```
rg SeriesGenerator
```

* cmd/influxd/generate
* mock
* storage/reads
