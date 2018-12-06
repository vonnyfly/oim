# Copyright 2017 The Kubernetes Authors.
# Copyright 2018 Intel Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

IMPORT_PATH=github.com/intel/oim

REGISTRY_NAME=localhost:5000
IMAGE_VERSION_oim-csi-driver=canary
IMAGE_TAG=$(REGISTRY_NAME)/$*:$(IMAGE_VERSION_$*)

REV=$(shell git describe --long --tags --match='v*' --dirty)

OIM_CMDS=oim-controller oim-csi-driver oim-registry oimctl

# Need bash for coproc in test/test.make.
SHELL=bash
# Ensures that (among others) _work/clear-kvm.img gets deleted when configuring it fails.
.DELETE_ON_ERROR:

include doc/doc.make
include test/test.make

# Build main set of components.
.PHONY: all
all: $(OIM_CMDS)

# Build all binaries, including tests.
# Must use the workaround from https://github.com/golang/go/issues/15513
build: $(OIM_CMDS)
	go test -run none $(TEST_ALL)

# Run operations only developers should need after making code changes.
update:
.PHONY: update


.PHONY: $(OIM_CMDS)
$(OIM_CMDS):
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-X main.version=$(REV) -extldflags "-static"' -o _output/$@ ./cmd/$@

# _output is used as the build context. All files inside it are sent
# to the Docker daemon when building images.
%-container: %
	cp cmd/$*/Dockerfile _output/Dockerfile.$*
	cd _output && \
	docker build \
		--build-arg HTTP_PROXY \
		--build-arg HTTPS_PROXY \
		--build-arg NO_PROXY \
		-t $(IMAGE_TAG) -f Dockerfile.$* .

push-%: %-container
	docker push $(IMAGE_TAG)

.PHONY: clean
clean:
	go clean -r -x
	-rm -rf _output _work

SPDK_SOURCE = vendor/github.com/spdk/spdk
_work/vhost:
	mkdir -p _work
	cd $(SPDK_SOURCE) && ./configure --with-rbd && make -j
	cp -a $(SPDK_SOURCE)/app/vhost/vhost $@

# protobuf API handling
OIM_SPEC := spec.md
OIM_PROTO := pkg/spec/oim.proto

# This is the target for building the temporary OIM protobuf file.
#
# The temporary file is not versioned, and thus will always be
# built on Travis-CI.
$(OIM_PROTO).tmp: $(OIM_SPEC) Makefile
	echo "// Code generated by make; DO NOT EDIT." > "$@"
	cat $< | sed -n -e '/```protobuf$$/,/^```$$/ p' | sed '/^```/d' >> "$@"

# This is the target for building the OIM protobuf file.
#
# This target depends on its temp file, which is not versioned.
# Therefore when built on Travis-CI the temp file will always
# be built and trigger this target. On Travis-CI the temp file
# is compared with the real file, and if they differ the build
# will fail.
#
# Locally the temp file is simply copied over the real file.
$(OIM_PROTO): $(OIM_PROTO).tmp
ifeq (true,$(TRAVIS))
	diff "$@" "$?"
else
	diff "$@" "$?" > /dev/null 2>&1 || cp -f "$?" "$@"
endif

# If this is not running on Travis-CI then for sake of convenience
# go ahead and update the language bindings as well.
ifneq (true,$(TRAVIS))
#build:
#	$(MAKE) -C lib/go
#	$(MAKE) -C lib/cxx
endif

update: update_spec
update_spec: $(OIM_PROTO)
	$(MAKE) -C pkg/spec

# check generated files for violation of standards
test: test_proto
test_proto: $(OIM_PROTO)
	awk '{ if (length > 72) print NR, $$0 }' $? | diff - /dev/null

update: update_dep
update_dep:
	dep ensure -v
