setup:
	cd app; yarn;
	@echo "SETUP DONE!"
build:
	cd app; yarn build-go; cd ..; bazel build :webview-wasm;