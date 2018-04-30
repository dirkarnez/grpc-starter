cd /d "proto" 

"C:\protoc-3.5.1-win32\bin\protoc.exe" --go_out=plugins=grpc:../server/src/proto/ mentor.proto
"C:\protoc-3.5.1-win32\bin\protoc.exe" --plugin=protoc-gen-ts=..\web\node_modules\.bin\protoc-gen-ts.cmd --ts_out=service=true:..\web\src\proto --js_out=import_style=commonjs,binary:..\web\src\proto mentor.proto

pause