#!/bin/sh
#export $(cat ./image/test.json)
TOKEN=`cat image/access.json | jq -r .TOKEN`
URL=`cat image/access.json | jq -r .URL`
text="message=$1"

echo "curl --data-urlencode $text -H X-TYPETALK-TOKEN:$TOKEN $URL"
curl --data-urlencode $text -H X-TYPETALK-TOKEN:$TOKEN $URL
