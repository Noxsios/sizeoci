Output demo.gif

Set Shell "bash"
Set FontSize 24
Set Width 1600
Set Height 800

Type "go install github.com/noxsios/sizeoci@latest" Sleep 750ms Enter Enter
Sleep 500ms Type "sizeoci docker.io/library/nginx:latest" Sleep 750ms Enter
Sleep 750ms Type "sizeoci -p multi/amd64 ghcr.io/zarf-dev/packages/dos-games:1.2.0" Sleep 1000ms Enter

Sleep 5s

Enter
