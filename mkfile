PROG=sa

SRC=sa.go matrix.go

INC=.
LIB=.

O=`(uname -m -p 2>/dev/null || uname -m) | sed '
	s;.*i[3-6]86.*;8;;
	s;.*i86pc.*;8;;
	s;.*amd64.*;6;;
	s;.*x86_64.*;6;;
	s;.*armv.*;5;;
'`

GO=${O}g
LN=${O}l

OBJ=${SRC:%.go=%.$O}

$O.$PROG:	$PROG.$O
	$LN -L$LIB -o $O.$PROG $PROG.$O

%.$O:	%.go
	$GO -I$INC $stem.go

sa.$O:	matrix.$O

default:	$TAR

clean:
	rm $OBJ $TAR

