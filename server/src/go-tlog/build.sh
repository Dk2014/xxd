echo "change director to"
pushd src/parser

echo "building"
go build
mv parser ../../bin/

echo "change director to"
popd
