
# This how we want to name the binary output
MAIN="cmd/command/main.go"
BINARY="cmd/bin/simpleblockchain"
VER_FILE="version.txt"

# These are the values we want to pass for VERSION and BUILD
BUILD_TS := $(shell date +%FT%T%z)

# Parse file
PREV_VERSION := `[ -f ${VER_FILE} ] && cat ${VER_FILE} || echo 00.000.000`
MAJOR := $(shell echo $(PREV_VERSION) | cut -f1 -d. | sed 's/^0*//')
MINOR := $(shell echo $(PREV_VERSION) | cut -f2 -d. | sed 's/^0*//')
BUILD := $(shell echo $(PREV_VERSION) | cut -f3 -d. | sed 's/^0*//')

FIXNULL= $(if $(1),$(1),0)
INCVER = `echo $1+1 | bc`

ADD_MAJOR = $(shell printf "%02d.000.000\n" $(call INCVER,$(call FIXNULL,$(MAJOR))))
ADD_MINOR = $(shell printf "%02d.%03d.000\n" $(call FIXNULL,$(MAJOR)) $(call INCVER,$(call FIXNULL,$(MINOR))))
ADD_BUILD = $(shell printf "%02d.%03d.%03d\n" $(call FIXNULL,$(MAJOR)) $(call FIXNULL,$(MINOR)) $(call INCVER,$(call FIXNULL,$(BUILD))))

.PHONY: clean major minor build

update_ver_file = $(shell echo $(1) > $(VER_FILE))

# Builds the project
major:
	$(MAKE) make_linux LDFLAGS="-X main.Version=$(ADD_MAJOR) -X main.Build=$(BUILD_TS)"
	$(call update_ver_file,$(ADD_MAJOR))
minor:
	$(MAKE) make_linux LDFLAGS="-X main.Version=$(ADD_MINOR) -X main.Build=$(BUILD_TS)"
	$(call update_ver_file,$(ADD_MINOR))

build:
	$(MAKE) make_linux LDFLAGS="-X main.Version=$(ADD_BUILD) -X main.Build=$(BUILD_TS)"
	$(call update_ver_file,$(ADD_BUILD))

make_linux: clean
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BINARY) $(MAIN)

make_native: clean
	go build ${LDFLAGS} -o ${BINARY}"_native" ${MAIN}

# Remove temporary files
clean:
	go clean