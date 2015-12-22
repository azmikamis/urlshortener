#Start the service and send data through curl
#
#To shorten:
#
#curl -sX POST -H 'Content-Type:application/json' 'http://localhost:8080/shorten' -d '{"url":"http://a-very-long-url"}'
#
#Response:
#
#{"short":"http://localhost/1"}
#
#To expand:
#
#curl -sX POST -H 'Content-Type:application/json' 'http://localhost:8080/expand' -d '{"short":"http://localhost/1"}'
#
#Response:
#
#{"expand":"http://a-very-long-url"}

curl -sX POST -H 'Content-Type:application/json' 'http://localhost:8080/shorten' -d '{"url":"http://a-very-long-url"}'
curl -sX POST -H 'Content-Type:application/json' 'http://localhost:8080/expand' -d '{"short":"http://localhost/1"}'
