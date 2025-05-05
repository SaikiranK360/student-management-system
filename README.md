# grpc-go-practice-2

After cloning the project, check whether buf is installed in the system:
buf --version

If the version is not showing, install the buf:
brew install buf

To just generate the files corresponding to the proto files written:
buf generate

To see the structure of methods, check proto-gen/folder/filename_grpc.pb.go

The build any project:
make build

To run any service:
make run

To debug the code:
If the code is opened in VS Code, (Command + Shift + P) Go:Install/Update Tools, select all and click OK.
Then click Launch Package.
