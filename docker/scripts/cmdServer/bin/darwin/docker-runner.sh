#!/usr/bin/env bash

while read line
do
    isGenerate=`echo "$line" | grep 'deploy!'`

    if [[ ! -z "$isGenerate" ]]
     then
         GIT_URL=`echo "$line" | jq -r '.repository.ssh_url'`
         GIT_EMAIL=`echo "$line" | jq -r '.head_commit.committer.email'`
         GIT_NAME=`echo "$line" | jq -r '.head_commit.committer.name'`

         echo `date` "is commit contain text \"deploy!\"? TRUE"  >> docker-runner.log
         echo `date` "request body" "$line" >> docker-runner.log
         echo `date` "data for docker environment variables GIT_URL, GIT_EMAIL, GIT_NAME: "${GIT_URL} ${GIT_EMAIL} ${GIT_NAME} >> docker-runner.log

        docker run -t --rm -v ${HOME}/.ssh:/root/.ssh -e GIT_URL=${GIT_URL} -e GIT_EMAIL=${GIT_EMAIL} -e GIT_NAME='${GIT_NAME}' kupolua/site-builder
     else
        echo `date` "commit did not contain text deploy!"
    fi

done < "${1:-/dev/stdin}"


