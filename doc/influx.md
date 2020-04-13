
To clean out just the data but not the credentials

```
alias infclean='cd ~/.influxdbv2; rm -fr engine'
```

And then you can get in there and write out new data

```
alias infw='influx write -b rick -o ag -p s'
infw @filename.txt
```

Complete Sequence for Manual Testing

```
cd .influxdbv2
rm -fr *
influxd
infsetup
```

```
alias infsetup='influx setup --username storm --password jjjjjjjj --org ag --bucket rick'
```

You must then populate the filename **.influxenv** with this line of code...

```
export INFLUX_TOKEN=ieL6P0KeyZ_DKlxGQiv1VSZELOolY8ISSFVqZX8uw5e-eBolJdnlCUiByAVjoZety-EQPTkaAGveKo-aHjK03Q==
```

This file then gets sourced at the end of .profile

To trigger it do your standard command

```
sp
```

##### Talks about creating your new token from scratch

https://github.com/influxdata/influxdb#getting-started

##### References

https://v2.docs.influxdata.com/v2.0/write-data/use-telegraf/manual-config/

https://v2.docs.influxdata.com/v2.0/security/tokens/view-tokens/

In .influxdbv2 there is a **config** file which looks like this

```
[default]
  url = "http://localhost:9999"
  token = "ieL6P0KeyZ_DKlxGQiv1VSZELOolY8ISSFVqZX8uw5e-eBolJdnlCUiByAVjoZety-EQPTkaAGveKo-aHjK03Q=="
  org = "ag"
  active = true
```

```
env | grep INFLUX
INFLUX_TOKEN=ieL6P0KeyZ_DKlxGQiv1VSZELOolY8ISSFVqZX8uw5e-eBolJdnlCUiByAVjoZety-EQPTkaAGveKo-aHjK03Q==
```

```
cat .influxenv
produces the following result...

export INFLUX_TOKEN=ieL6P0KeyZ_DKlxGQiv1VSZELOolY8ISSFVqZX8uw5e-eBolJdnlCUiByAVjoZety-EQPTkaAGveKo-aHjK03Q==
```

which is sourced inside .profile at the end of the file...

Do not delete the influxd.bolt file because   
the organization and bucket information is stored inside the file   
~/.influxdbv2/influxd.bolt

so unless you want to go through this whole dance again...   
simply execute the command **infclean**

```
alias infclean='cd ~/.influxdbv2; rm -fr engine'
```

##### Start up scenarios

```
alias igr='gr cmd/influxd/main.go'
```

##### How to run e2e_test

**The test only runs once**, if you try to run it a second time
with out bringing down the server and then bringing it back up it fails

```
alias ie2e='gr cmd/influxd/main.go --e2e-testing --store=memory'
alias ie2ec='influxd --e2e-testing --store=memory'
```

##### On the client side run

Inside the repo
[influxdb-client-go](https://github.com/influxdata/influxdb-client-go)

```
gtvr TestE2E --e2e
```

[You must then set the INFLUX_TOKEN](https://github.com/stormasm/go-examples/blob/master/filenv/Readme.md) prior to writing out data to influxdb.

To directly write data to influxdb set the most recent time in the data file
below using the
[EpochConverter](https://www.epochconverter.com/) to get the times.

```
Go to the directory where the influxdb data file exists

alias inf1='influx write -b rick -o ag -p s @./temp1.txt'
alias infw='influx write -b rick -o ag -p s'
This is a key point about how bash aliases work...

infw @./t03.txt

or you can use this syntax

infw @t04.txt

```

[How to delete a specific organization](https://v2.docs.influxdata.com/v2.0/organizations/delete-org/)

```
influx org find
influx org delete -i 03db76dc3b09b000
```

[influxd - InfluxDB daemon](https://v2.docs.influxdata.com/v2.0/reference/cli/influxd/)

The initial config files by default are located in
```
cd
~/.influxdbv2
```

[Simple example of writing data](https://v2.docs.influxdata.com/v2.0/write-data/)
```
influx write -b rick -o ag -p s 'myMeasurement,host=myHost testField="testData" 1556896326'
influx write -b rick -o ag -p s 'weather,location=us-midwest temperature=82 1556896328'
```

```
influx write -b rick -o ag -p s 'weather,location=oregon temperature=42'
influx write -b rick -o ag -p s 'weather,location=oregon temperature=45'
influx write -b rick -o ag -p s 'weather,location=oregon temperature=41'
influx write -b rick -o ag -p s 'weather,location=oregon temperature=39'
influx write -b rick -o ag -p s 'weather,location=oregon temperature=49'
influx write -b rick -o ag -p s 'weather,location=oregon temperature=45'
influx write -b rick -o ag -p s 'weather,location=oregon temperature=47'
influx write 'weather,location=oregon temperature=100'
```

[Writing a data file](https://v2.docs.influxdata.com/v2.0/write-data/#example-influx-write-commands)

```
influx write -b rick -o ag -p s @./temp1.txt
```

[Here is a nice timestamp visualizer called EpochConverter](https://www.epochconverter.com/)

### Influxd Generate and Inspect Commands

[influxd generate and inspect](https://v2.docs.influxdata.com/v2.0/reference/cli/influxd/)

In order for the commands below to work you must go through the

```
infclean
infd
infsetup
and then bring down infd
```

print to stdout all of the parameters
```
influxd generate simple --bucket rick --org ag --print
```

write out a sample file
```
influxd generate simple --bucket rick --org ag
```

write out a sample schema
```
influxd generate help-schema > myschema.toml
```

Comment out the line
```
{ name = "rack", source = { type = "file", path = "files/racks.txt" } },
```

```
infd generate myschema.toml --bucket rick --org ag
```

You can then go ahead and inspect the file...

```
infd inspect report-tsm
```

##### alias references from .golang

aliases for testing across all filenames in a directory

```
alias gt='go test'
alias gtv='gt -v'
```

aliases for testing names of tests inside files

```
alias gtr='go test -run'
alias gtvr='gtv -run'
```

##### Cmd Help

```
influxd generate help-schema
```

[More details here](https://v2.docs.influxdata.com/v2.0/reference/cli/influxd/generate/)

In the code if you want to see where Token data gets read do a
```
rg Token
```
in the cmd directory
