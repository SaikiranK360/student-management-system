# grpc-go-practice-2

After cloning the project, check whether buf is installed in the system:
buf --version

If the version is not showing, install the buf:
brew install buf

The to generate files from proto files:
buf generate

Then to download the listed dependencies in go.mod file:
go mod tidy
