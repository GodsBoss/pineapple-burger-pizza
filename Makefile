GOROOT ?= /usr/local/go
SERVE_PORT ?= 8080

all: dist/main.wasm dist/wasm_exec.js dist/index.html dist/gfx.png

dist/main.wasm: | dist
	GOOS=js GOARCH=wasm go build -o $@ ./main

dist/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js | dist
	cp $< $@

dist/index.html: static/index.html | dist
	cp $< $@

dist/gfx.png: gfx/gfx.xcf gfx/gfx.sh | dist
	gfx/gfx.sh $< $@

dist:
	mkdir -p dist

clean:
	rm -rf dist

serve:
	docker run --rm -it \
		-v ${PWD}/static/usr/local/apache2/conf/httpd.conf:/usr/local/apache2/conf/httpd.conf:ro \
		-v ${PWD}/dist:/usr/local/apache2/htdocs/:ro \
		-p ${SERVE_PORT}:80 \
		httpd:2.4.46

.PHONY: all clean dist/main.wasm serve
