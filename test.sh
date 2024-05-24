
#!/bin/bash
# Test 1 - number of trains is not a valid positive integer (0 is not an integer) london.map
project_root_dir=$(pwd)

cd cmd
echo
echo START TEST 1
echo
go run .  ../assets/tests/input/london-network.map waterloo st_pancras -4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
echo  displays "Error" on stderr when the number of trains is not a valid positive integer.
echo "END OF TEST"
echo 

# Test 2 - prints the train movements with the correct format (0 is not an integer) london.map
cd $project_root_dir
cd cmd
echo
echo START TEST 2
echo
go run .  ../assets/tests/input/london-network.map waterloo st_pancras 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test2.txt
echo 
echo "END OF TEST"
echo 
# Test 3 - connection is made with a station which does not exist) london.map
cd $project_root_dir
cd cmd
echo
echo START TEST 3
echo
go run . ../assets/tests/input/err1.txt  waterloo st_pancras 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
echo displays "Error" on stderr when a connection is made with a station which does not exist.
echo 
echo "END OF TEST"
echo 
# Test 4 - completes the movements in no more than 6 turns for 4 trains between bond_square and space_port
cd $project_root_dir
cd cmd
echo
echo START TEST 4
echo
go run . ../assets/tests/input/space-port.map bond_square space_port 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test4.txt
echo 
echo "END OF TEST"
echo 



## We can just run the progamm and the files under assets/test/input
## 


