
#build
mkdir -p bin
g++ -Isrc src/generator_base.cc src/generator.cc src/plugin.cc -std=c++11 `pkg-config --cflags --libs protobuf` -lprotoc -o bin/protoc-gen-java_srv && \

#install
chmod +x bin/protoc-gen-java_srv && \
sudo cp bin/protoc-gen-java_srv /usr/local/bin/ && \


#run code generator
echo "generating code" && \
mkdir -p javaoutput
protoc -I proto/ --java_out=javaoutput proto/*.proto && \
protoc -I proto/ --java_srv_out=javaoutput proto/*.proto && \

echo "done"