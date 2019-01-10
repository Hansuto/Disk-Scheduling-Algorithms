rm ex # remove previous student executable
go build -o ex $1 # build new student executable
echo "error 01 req > upper - produces error"
./ex error01.txt
echo "error 02 init < lower - produces ABORT" 
./ex error02.txt
echo "error 03 init > upper - produces ABORT"
./ex error03.txt
echo "error 04 upper < lower - produces ABORT"
./ex error04.txt
