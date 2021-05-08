startContainer(){
    echo "Starting the MySQL Contianer"
    cd ./mysql
    echo "Executing the run script to start the contianer"
    ./run.sh
    count=1
    echo "Starting the monitor of the Container status"
    while [ $count -lt $1 ]; do
        dockerStatus=$2
        dockerPS=$(docker ps)
        if [[ $dockerPS == *"${dockerStatus[0]}"* ]];then
            echo "\nDocker container is now healthy and ready to use"
            break
        fi
        if [[ $dockerPS == *"${dockerStatus[1]}"* ]];then
            echo "\nDocker container is in an unhealthy state. Please check the logs for further details"
            break
        fi
        if [[ $dockerPS == *"${dockerStatus[2]}"* ]];then
            if [[ $count ==  1 ]];then
                printf "Container Starting Up ."
            else
                printf "."
            fi
        fi
        count=$((count+1))
        sleep 1 # sleep for 1 second
    done
}

runningContainers=$(docker ps)
# echo $runningContainers
containerName='mysql-cluster'
dockerStatus[0]="healthy"
dockerStatus[1]="unhealthy"
dockerStatus[2]="starting"
timeout=120

if [[ "$runningContainers" == *"$containerName"* ]];then
    echo "Mysql Container is running"
else
    startContainer $timeout $dockerStatus
fi