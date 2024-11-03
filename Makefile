default:
	# Generate packaged executables for Linux & Windows.
	# In the case of Linux, a Makefile is created
	# to facilitate its installation in the operating system.
	# 
	# Run "create-bundles" to generate the bundles

create-bundles:
	wails build -clean -o wails-todoapp
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc wails build -skipbindings -s -platform windows/amd64 -o wails-todoapp.exe
	sh scripts/packaging-linux.sh 'Wails Todoapp' wails-todoapp
	sh scripts/packaging-windows.sh wails-todoapp