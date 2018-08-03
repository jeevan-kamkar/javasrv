# Building dependencies

Need to have protobuf installed to build (note that protobuf does not have an uninstaller)

```
git clone git@github.com:google/protobuf.git
cd protobuf
git submodule update --init --recursive
cmake cmake/
cd cmake
make
sudo make install
```



## Java

You'll need to install jdk
Installing java protoc runtime requires using maven, which I didn't set up, so I built it from source and included the jar in the repository.



```
cd java
protoc --java_out=core/src/main/java -I../src ../src/google/protobuf/descriptor.proto
cd core/src/main/java/com/google/protobuf
javac -d ./build *.java
cd build
jar cvf javaprotobuf.jar *
```

