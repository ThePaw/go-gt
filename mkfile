PROG=sa

O=`(uname -m -p 2>/dev/null || uname -m) | sed '
	s;.*i[3-6]86.*;8;;
	s;.*i86pc.*;8;;
	s;.*amd64.*;6;;
	s;.*x86_64.*;6;;
	s;.*armv.*;5;;
'`

TAR=${O}.$PROG

GO=${O}g
LN=${O}l

SRC=sa.go

OBJ=${SRC:%.go=%.$O}

$TAR:	$OBJ
	$LN -o $TAR $OBJ

$OBJ:	$SRC
	$GO $prereq

default:	$TAR

clean:
	rm $OBJ $TAR

