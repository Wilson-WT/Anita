# Set variable Project name
project_name=Anita
# Set variable project path
project_path=gitlab.hopebaytech.com/autox/$project_name
# Delete project
rm -rf $GOPATH/src/$project_path
# Get project again
go get -v $project_path
cd $GOPATH/src/$project_path
git checkout develop
go get -v ./...

# Run tests
go get -u github.com/jstemmer/go-junit-report
go test -v -cover ./... | go-junit-report > report.xml
cp $GOPATH/src/$project_path/report.xml $WORKSPACE/report.xml
