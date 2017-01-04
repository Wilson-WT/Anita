# Download latest binary
wget http://jenkins/view/Arkflex%20QA/job/Anita_Build/lastSuccessfulBuild/artifact/Scheduler_Linux.zip

# Kill process
if [ -z $(pidof ./Anita) ];
then
	echo Scheduler is not running...
else
    echo Scheduler is running, kill process...
    proecessid=$(pidof ./Anita)
    echo $proecessid
	kill -9 $proecessid
fi

# Clean deploy directory
cd $HOME/deploy
rm -rf Scheduler/

# unzip file to depoly directory
unzip $WORKSPACE/Scheduler_Linux.zip -d .
# Copy config
cp config.ini Scheduler/

# Execute binary
cd Scheduler
BUILD_ID=dontKillMe nohup ./Anita >> output 2>&1 &
sleep 3s

# Check process
if [ -z $(pidof ./Anita) ];
then
    echo Scheduler is not running...
    exit 1
else
    echo Scheduler is running succesfully!
fi
