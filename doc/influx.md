
[Getting Started](https://github.com/influxdata/influxdb#getting-started)

```
alias infsetup='influx setup --username storm --password 12345678 --org ag --bucket rick'
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
