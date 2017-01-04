# Set variable Project name
project_name=Anita
# Set variable project path
project_path=gitlab.hopebaytech.com/autox/$project_name
# Delete project
rm -rf $GOPATH/src/$project_path
# Get project again
go get $project_path
cd $GOPATH/src/$project_path
git checkout develop
go get -v ./...

# Build binary
## for Linux
echo Building for Linux
go build
## for windows
echo Building for Windows
GOOS=windows GOARCH=386 go build

# Package binary and config
## for Linux
echo Packaging for Linux
mkdir Scheduler
cp $project_name Scheduler/
cp config.ini Scheduler/

# Zip Scheduler
zip -r Scheduler_Linux.zip Scheduler/*
rm -rf Scheduler/

## for Windows
echo Packaging for Windows
mkdir Scheduler
cp $project_name.exe Scheduler/
cp config.ini Scheduler/

# Zip Scheduler
zip -r Scheduler_Windows.zip Scheduler/*
rm -rf Scheduler/

# Copy to workspace
cp $GOPATH/src/$project_path/Scheduler_Linux.zip $WORKSPACE/Scheduler_Linux.zip
cp $GOPATH/src/$project_path/Scheduler_Windows.zip $WORKSPACE/Scheduler_Windows.zip


## Report Server Build and Zip
cd reportserver

# Build
go build
GOOS=windows GOARCH=386 go build

# Packaging for Linux
mkdir ReportServer
cp reportserver ReportServer/
cp config.ini ReportServer/
cp favicon.ico ReportServer/

# Zip ReportServer
zip -r ../ReportServer_Linux.zip ReportServer/*
rm -rf ReportServer/

# Packaging for Windows
mkdir ReportServer
cp reportserver.exe ReportServer/
cp config.ini ReportServer/
cp favicon.ico ReportServer/

# Zip ReportServer
zip -r ../ReportServer_Windows.zip ReportServer/*
rm -rf ReportServer/

# Copy to workspace
cp $GOPATH/src/$project_path/ReportServer_Linux.zip $WORKSPACE/ReportServer_Linux.zip
cp $GOPATH/src/$project_path/ReportServer_Windows.zip $WORKSPACE/ReportServer_Windows.zip


## Sheet Server Build and Zip
cd ..
cd sheetserver

# Build
go build
GOOS=windows GOARCH=386 go build

# Packaging for Linux
mkdir SheetServer
cp sheetserver SheetServer/
cp config.ini SheetServer/
cp client_secret.json SheetServer/

# Zip SheetServer
zip -r ../SheetServer_Linux.zip SheetServer/*
rm -rf SheetServer/

# Packaging for Windows
mkdir SheetServer
cp sheetserver.exe SheetServer/
cp config.ini SheetServer/
cp client_secret.json SheetServer/

# Zip SheetServer
zip -r ../SheetServer_Windows.zip SheetServer/*
rm -rf SheetServer/

# Copy to workspace
cp $GOPATH/src/$project_path/SheetServer_Linux.zip $WORKSPACE/SheetServer_Linux.zip
cp $GOPATH/src/$project_path/SheetServer_Windows.zip $WORKSPACE/SheetServer_Windows.zip
