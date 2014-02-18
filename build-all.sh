# usage: source build-all.sh
# Assumes #2 in this article has already been done: http://dave.cheney.net/2013/07/09/an-introduction-to-cross-compilation-with-go-1-1
rm -rf bin/*
go-build-all gemmer.go
mv gemmer-* bin