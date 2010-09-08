# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=cairo
CGOFILES=\
	cairo.go

#CGO_CFLAGS= -I/opt/local/include
#CGO_LDFLAGS+= -L/opt/local/lib -lcairo
CGO_LDFLAGS+= -lcairo

CLEANFILES+=draw

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
