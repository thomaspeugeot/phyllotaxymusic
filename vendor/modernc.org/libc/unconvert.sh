set -evx
until unconvert -fastmath ./...
do
	unconvert -fastmath -apply ./...
done
git checkout -f pthread_musl.go
