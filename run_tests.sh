echo "Python"
/usr/bin/time python3.7 python/main.py

echo "Go"
echo "v1/"
/usr/bin/time go run v1/main.go
echo "v2/"
/usr/bin/time go run v2/main.go
echo "v3/"
/usr/bin/time go run v3/main.go
echo "v4/"
/usr/bin/time go run v4/main.go
echo "v5/"
/usr/bin/time go run v5/main.go
echo "v6/"
/usr/bin/time go run v6/main.go

