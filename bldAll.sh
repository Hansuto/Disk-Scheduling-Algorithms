rm *.base #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

echo "run all FCFS cases"
./ex fcfs01.txt>fcfs01.base
diff fcfs01.base fcfs01.base
./ex fcfs20.txt>fcfs20.base
diff fcfs20.base fcfs20.base
./ex fcfsPA.txt>fcfsPA.base
diff fcfsPA.base fcfsPA.base

echo "run all SSTF cases"
./ex sstf01.txt>sstf01.base
diff sstf01.base sstf01.base
./ex sstf20.txt>sstf20.base
diff sstf20.base sstf20.base
./ex sstfPA.txt>sstfPA.base
diff sstfPA.base sstfPA.base

echo "run all SCAN cases"
./ex scan01.txt>scan01.base
diff scan01.base scan01.base
./ex scan20.txt>scan20.base
diff scan20.base scan20.base
./ex scanPA.txt>scanPA.base
diff scanPA.base scanPA.base

echo "run all C-SCAN cases"
./ex c-scan01.txt>c-scan01.base
diff c-scan01.base c-scan01.base
./ex c-scan20.txt>c-scan20.base
diff c-scan20.base c-scan20.base
./ex c-scanPA.txt>c-scanPA.base
diff c-scanPA.base c-scanPA.base

echo "run all LOOK cases"
./ex look01.txt>look01.base
diff look01.base look01.base
./ex look20.txt>look20.base
diff look20.base look20.base
./ex lookPA.txt>lookPA.base
diff lookPA.base lookPA.base

echo "run all C-LOOK cases"
./ex c-look01.txt>c-look01.base
diff c-look01.base c-look01.base
./ex c-look20.txt>c-look20.base
diff c-look20.base c-look20.base
./ex c-lookPA.txt>c-lookPA.base
diff c-lookPA.base c-lookPA.base
