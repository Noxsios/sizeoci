# sizeoci

Replicate `oras manifest fetch | jq '[.layers[].size] | add' | awk ...` in a simple Go CLI.

![demo](https://github.com/user-attachments/assets/20dd9a55-1b09-4c57-840a-0833307096d8)

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
