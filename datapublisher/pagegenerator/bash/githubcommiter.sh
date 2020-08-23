#!/bin/bash   
#./bash/githubcommiter.sh "/home/tarek/Projects/Elwizara/LiveApps/Elwizara.com/"


git --git-dir=$1.git/ --work-tree=$1 add .

git --git-dir=$1.git/ --work-tree=$1 commit -m "auto commit"

git --git-dir=$1.git/ --work-tree=$1 push