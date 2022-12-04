test:
	go test ./...

run:
	go run github.com/tomasruud/advent/2022/day$(day)

gen:
	mkdir day$(day) && \
	(cd day$(day) && \
	echo "package main\n\nconst input = \`\`" >> input.go && \
	echo "package main\n\nfunc main() {\n}" >> main.go && \
	echo "package main" >> main_test.go && \
	git add .)