echo "Generating files..."
mkdir -p ./v2/storiespb
# protoc -I ../proto/ --proto_path=../proto/third_party ../proto/stories.proto --go_out=plugins=grpc:./v2
protoc -I=../proto/ --go_opt=paths=source_relative --go_out=v2/storiespb ../proto/stories.proto
