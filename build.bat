echo Building go...
go build -o ./bin/windows
echo Building .msi
go-msi make --msi NetWatcher-Agent.msi --path ./wix/wix.json --src ./wix --out ./bin/windows --version 1.0.1