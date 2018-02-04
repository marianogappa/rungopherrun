docker build -t rungopherrun:v1 .
docker run --rm -p 9000:9000 -v /Users/marianol/Code/go/src/github.com/marianogappa/rungopherrun/plugins/:/mnt/plugins/ --name rungopherrun rungopherrun:v1
