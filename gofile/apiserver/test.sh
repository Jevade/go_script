#! /bin/bash
for n in $(seq 1 1 100000)
do
    nohup curl -XGET -H "Content-Type: application/json"  http://testlog.log/v1/user/hanmeimei9  &>/dev/null
done