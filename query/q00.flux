from(bucket: "rick")
  |> range(start: -30d)
  |> filter(fn: (r) => r._measurement == "mem")
