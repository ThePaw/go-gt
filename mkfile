TAR=sa

SRC=sa.go

OBJ=${SRC:%.go=%.6}

$TAR:	$OBJ
	6l -o $TAR $OBJ

$OBJ:	$SRC
	6g $prereq

default:	$TAR

clean:
	rm $OBJ $TAR
