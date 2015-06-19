
MAKEFLAGS = -s

sql.go: sql.y
	go tool yacc -o sql.go sql.y
	gofmt -w sql.go

clean:
	rm -f y.output sql.go
