#!/bin/bash   

# Print the usage message
function printHelp() {
  echo "Usage: "
  echo "  run.sh <mode> [-p <host path>] "
  echo "    <mode> - one of 'up', 'down' or 'restart'"
  echo "      - 'up' - bring up the network "
  echo "      - 'down' - clear the network"
  echo "      - 'restart' - restart the network" 
  echo "    -p <host path> - directory has files to apache server" 
  echo "  run.sh -h (print this message)"
  echo 
  echo "	run.sh up -c mychannel -s couchdb" 
  echo "	run.sh down -c mychannel" 
  echo
  echo "Taking all defaults:" 
  echo "	run.sh up"
  echo "	run.sh down"
  echo "	run.sh restart"
}

HOSTDIR="/home/tarek/Projects/Elwizara/Code/Elwizara/datapublisher/test/frontendhtml"

function networkUp() {
    #docker run -d -p 8051:8050 --net host --name testpageSplash scrapinghub/splash  
    docker run -d -p 9091:80 -v $HOSTDIR:/usr/local/apache2/htdocs/ --name testpageHOST httpd 
    (cd Generator;./testpage &) 
}


function networkDown() {
    kill $(ps aux | grep './testpage' | awk '{print $2}') 
    docker rm -f $(docker ps -a --filter "name=testpage" -q)
} 
 
MODE=$1 

if [ "${MODE}" == "up" ]; then 
    networkUp
elif [ "${MODE}" == "down" ]; then 
    networkDown
elif [ "${MODE}" == "restart" ]; then 
    networkDown 
    networkUp 
else  
    printHelp
    exit 1
fi
