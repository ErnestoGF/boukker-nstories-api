echo "Generating files..."
mkdir -p ./v2/storiespb
protoc -I=../proto/ \
--go_opt=paths=source_relative --go_out=v2/storiespb \
--go-grpc_opt=paths=source_relative --go-grpc_out=v2/storiespb \
../proto/commons.proto \
../proto/chapters.proto \
../proto/bookmarks.proto \
../proto/readings.proto \
../proto/stories.proto
