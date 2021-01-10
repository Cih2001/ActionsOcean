FROM debian:latest

ADD main .

ENTRYPOINT exec ./main :1234 $DBUSERNAME $DBPASSWORD
