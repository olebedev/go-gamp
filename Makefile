SOURCE = ./swagger.yaml
LIB    = ./.lib

build: $(LIB)

$(LIB): $(SOURCE)
	swagger generate client --skip-validation -f $<
	touch $@
