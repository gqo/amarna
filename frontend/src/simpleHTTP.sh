#!/bin/sh

while true 
do
    cat <<EOF |
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *

{
    "letters": [
        {
            "body": "hello world",
            "timestamp": "199001010201122323",
            "from": "user1"
        },
        {
            "body": "hi there.",
            "timestamp": "199001010201122323",
            "from": "user2"
        },
        {
            "body": "yes, me too!",
            "timestamp": "199001010201122323",
            "from": "user1"
        },
        {
            "body": "test message\nwith a newline",
            "timestamp": "199001010201122323",
            "from": "user2"
        }
    ],
    "lesson": "lorem ipsum"
}
EOF
    nc -l -p $1
done
