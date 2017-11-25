# How to generate the htpasswd file

docker run --rm -ti xmartlabs/htpasswd test_user test_password > htpasswd

# How to start the server

go run main.go -static-dir=. -password-file=htpasswd -p=:5000
