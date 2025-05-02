# sizeoci

Replicate `oras manifest fetch | jq '[.layers[].size] | add' | awk ...` in a simple Go CLI.

```bash
go install github.com/noxsios/sizeoci@latest
```

```bash
$ sizeoci -h
  -h    Print this message and exit.
  -p string
        request platform in the form of os[/arch][/variant][:os_version] (default "linux")
  -v    Print the version number of sizeoci and exit.
```
