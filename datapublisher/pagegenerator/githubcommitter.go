package main

import (
	"database/sql"
	"fmt"
	"os/exec"
	"time"
)

func githubcommitter(db *sql.DB) {
	duration := time.Duration(config.CommitDuration)
	for {
		logs.Debugf("=========waiting for generatePageWG")
		generatePageWG.Wait()
		logs.Debugf("========= githubPushWG is starting")
		githubPushWG.Add(1)
		if err := commit(); err != nil {
			logs.Critical(err)
		}
		//log pushed files to github
		if err := logGitPushedDB(db); err != nil {
			logs.Critical(err)
		}
		logs.Debugf("========= githubPushWG Done")
		githubPushWG.Done()
		<-time.After(time.Minute * duration)
	}
}

func commit() error {
	//the path should end with "/"
	_, err := exec.Command(config.CommitScriptPath, fmt.Sprintf("%v/", config.GeneratedDataPath)).Output()
	if err != nil {
		logs.Criticalf("execute Command failed:%v:%v output:%v|%s", config.CommitScriptPath, config.GeneratedDataPath)
		return err
	}
	logs.Infof("Githup COMMIT Done")
	return nil
}

//git --git-dir=/home/tarek/Projects/Elwizara/LiveApps/Elwizara.com/.git/ --work-tree=/home/tarek/Projects/Elwizara/LiveApps/Elwizara.com/ add .
