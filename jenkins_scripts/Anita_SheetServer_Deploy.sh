# Download latest binary
wget http://jenkins/view/Arkflex%20QA/job/Anita_Build/lastSuccessfulBuild/artifact/SheetServer_Linux.zip

# Kill process
if [ -z $(pidof ./sheetserver) ];
then
	echo Sheet Server is not running...
else
    echo Sheet Server is running, kill process...
    proecessid=$(pidof ./sheetserver)
    echo $proecessid
	kill -9 $proecessid
fi

# Clean deploy directory
cd $HOME/deploy
rm -rf SheetServer/

# unzip file to depoly directory
unzip $WORKSPACE/SheetServer_Linux.zip -d .

# Execute binary
cd SheetServer
BUILD_ID=dontKillMe nohup ./sheetserver >> output 2>&1 &

# Check process
if [ -z $(pidof ./sheetserver) ];
then
    echo Sheet Server is not running...
    exit 1
else
    echo Sheet Server is running succesfully!
fi
