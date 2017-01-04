# Download latest binary
wget http://jenkins/view/Arkflex%20QA/job/Anita_Build/lastSuccessfulBuild/artifact/ReportServer_Linux.zip

# Kill process
if [ -z $(pidof ./reportserver) ];
then
	echo Report server is not running...
else
    echo Report server is running, kill process...
    proecessid=$(pidof ./reportserver)
    echo $proecessid
	kill -9 $proecessid
fi

# Clean deploy directory
cd $HOME/deploy
rm -rf ReportServer/

# unzip file to depoly directory
unzip $WORKSPACE/ReportServer_Linux.zip -d .

# Execute binary
cd ReportServer
BUILD_ID=dontKillMe nohup ./reportserver >> output 2>&1 &

# Check process
if [ -z $(pidof ./reportserver) ];
then
    echo Report server is not running...
    exit 1
else
    echo Report server is running succesfully!
fi
