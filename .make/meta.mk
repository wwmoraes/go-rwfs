empty :=
comma := ,
space := $(empty) $(empty)

%/:
	@test -d $@ || mkdir -p $@
