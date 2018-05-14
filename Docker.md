# addservermessage Docker Helper

The addservermessage Docker Helper is a simple tool to help build and run addservermessage as a container

## Prerequisites

You must have at least Docker v17 installed on your system.

Having this repo cloned helps too 😇

## Build
From wherever you have cloned this repo, run

`docker build -t addservermessage_container .`

(yes, you can replace **addservermessage_container** with whatever you want to call the container.  e.g. **addservermessage**, **foo**, etc.)

#### Cross-Compile
To cross-compile for a different target, you can pass in a `build-arg` as so

`docker build -t addservermessage_container --build-arg GOOS=darwin .`

## Run
`docker run --rm addservermessage_container <arguments>`
  
* This will start up **addservermessage** with `<arguments>`.
* **Note** - In the above, replace **addservermessage_container** with whatever you called it when you built it - e.g. **addservermessage**, **foo**, etc.

## Copy
So yeah, you want to get your binary _out_ of the container. To do so, you basically mount your target into the container, and copy the binary over, like so


`docker run --rm --entrypoint='' -v <FULLY_QUALIFIED_PATH_TO_TARGET_DIRECTORY>:/destination addservermessage_container /bin/cp /go/bin/addservermessage /destination`

e.g.

`docker run --rm --entrypoint='' -v /tmp:/destination addservermessage_container /bin/cp /go/bin/addservermessage /destination`

which will copy the binary to `/tmp/addservermessage`

**Note** : You should replace ** addservermessage_container** with whatever you called it in the **build** section above  e.g. **addservermessage**, **foo**, etc.

#### Cross-Compile
If you cross-compiled to a different target, your binary will be in `/go/bin/<target>/addservermessage`.  e.g. If you built with `--build-arg GOOS=darwin`, then you can copy out the binary with

`docker run --rm --entrypoint='' -v <FULLY_QUALIFIED_PATH_TO_TARGET_DIRECTORY>:/destination addservermessage_container /bin/cp /go/bin/darwin_amd64/addservermessage /destination`

e.g.

`docker run --rm --entrypoint='' -v /tmp:/destination addservermessage_container /bin/cp /go/bin/darwin_amd64/addservermessage /destination` 

which will copy the darwin_amd64 version of the binary to `/tmp/addservermessage`

**Note** : You should replace ** addservermessage_container** with whatever you called it in the **build** section above  e.g. **addservermessage**, **foo**, etc.
