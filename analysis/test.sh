set -e

export target=go

DIR=$(pwd)/$(date +%Y-%m-%d_%H-%M-%S)
mkdir $DIR

cd ../dev
# : > mysql/log/general.log
# : > mysql/log/slow.log
# : > nginx/log/access.log
# : > nginx/log/error.log


docker compose -f docker-compose-go.yaml exec frontend bash -c ': > /var/log/nginx/access.log'
docker compose -f docker-compose-go.yaml exec mysql bash -c ': > /var/log/mysql/slow.log'
docker compose -f docker-compose-go.yaml exec mysql bash -c ': > /var/log/mysql/general.log'

pushd ../benchmarker
./bin/benchmarker 1> $DIR/bench.log 2>&1
popd

# docker compose -f docker-compose-go.yaml cp frontend:/var/log/nginx/access.log $DIR/nginx.log
# docker compose -f docker-compose-go.yaml cp mysql:/var/log/mysql/slow.log $DIR/slowquery.log
# docker compose -f docker-compose-go.yaml cp mysql:/var/log/mysql/general.log $DIR/general.log


cat nginx/log/access.log | alp ltsv --sort sum -r -m  '^/api/announcements$','^/api/announcements/[^/]+','^/api/courses/[^/]+','/api/courses/.+/status$','^/api/courses/.+/classes$','^/api/courses/.+/classes/.+/assignments$','^/api/courses/.+/classes/.+/assignments/export$','^/api/courses/.+/classes/.+/assignments/export$','^/api/courses/.+/classes/.+/assignments/scores$' > $DIR/alp_result.log

pt-query-digest mysql/log/slow.log > $DIR/pt_query.log