
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
echo SART TEST 3
echo
go run . ../assets/tests/input/err1.txt  waterloo st_pancras 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo "EXPECTED RESULT"
echo 
echo displays "Error" on stderr when a connection is made with a station which does not exist.
echo 
echo "END OF TEST"
echo 
# Test 4 - completes the movements in no more than 6 turns for 4 trains between bond_square and space_port
cd $project_root_dir
cd cmd
echo
echo START TEST 4
go run . ../assets/tests/input/space-port.map bond_square space_port 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test4.txt
echo 
echo 
echo "END OF TEST"
Test 5 - displays "Error" on stderr when no path exists between the start and end stations. london.map
cd $project_root_dir
cd cmd
echo
echo SART TEST 5
echo
go run . ../assets/tests/input/err2.txt  waterloo hannes 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo "EXPECTED RESULT"
echo 
echo displays "Error" on stderr when no path exists between the start and end stations.
echo 
echo "END OF TEST"
echo 
Test 6 - displays "Error" on stderr when station names are duplicated. london.map
cd $project_root_dir
cd cmd
echo
echo SART TEST 6
echo
go run . ../assets/tests/input/err3.txt  waterloo st_pancras 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo "EXPECTED RESULT"
echo 
echo displays "Error" on stderr when station names are duplicated.
echo 
echo "END OF TEST"
echo 
Test 7 - err4    -displays "Error" on stderr when station names are invalid. london.map
cd $project_root_dir
cd cmd
echo
echo SART TEST 7
echo
go run . ../assets/tests/input/err4.txt  waterloo st_pancras 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo "EXPECTED RESULT"
echo 
echo displays "Error" on stderr when station names are invalid.
echo 
echo "END OF TEST"
echo 
# Test 8 - err5   -displays "Error" on stderr when the map does not contain a "stations:" section. london.map
cd $project_root_dir
cd cmd
echo
echo SART TEST 8
echo
go run . ../assets/tests/input/err5.txt  waterloo st_pancras 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo "EXPECTED RESULT"
echo 
echo displays "Error" on stderr when the map does not contain a "stations:"
echo 
echo "END OF TEST"
echo 
Test 9 - err6   -displays "Error" on stderr when a map contains more than 10000 stations.
cd $project_root_dir
cd cmd
echo
echo SART TEST 9
echo
go run . ../assets/tests/input/err6.txt  waterloo st_pancras 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo "EXPECTED RESULT"
echo 
echo displays "Error" on stderr when a map contains more than 10000 stations.
echo 
echo "END OF TEST"
echo 
Test 10 - It finds more than one valid route for 100 trains between waterloo and st_pancras in the London Network Map
cd $project_root_dir
cd cmd
echo
echo START TEST 10
go run . ../assets/tests/input/london-network.map waterloo st_pancras 100 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test10.txt
echo 
echo 
echo "END OF TEST"
# Test 11 - It finds more than one valid route for 100 trains between waterloo and st_pancras in the London Network Map
cd $project_root_dir
cd cmd
echo
echo START TEST 11
go run . ../assets/tests/input/beethoven-part.map beethoven part 9 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test11.txt
echo 
echo 
echo "END OF TEST"
# Test 12 - test12- It completes the movements in no more than 11 turns for 20 trains between beginning and terminus
cd $project_root_dir
cd cmd
echo
echo START TEST 12
go run . ../assets/tests/input/beginning-terminus.map beginning terminus 20 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test12.txt
echo 
echo 
echo "END OF TEST"
# Test 13 - test13- It completes the movements in no more than 6 turns for 4 trains between two and four
cd $project_root_dir
cd cmd
echo
echo START TEST 13
go run . ../assets/tests/input/two-four.map two four 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test13.txt
echo 
echo 
echo "END OF TEST"
# Test 14 - It can find more than one route for 2 trains between waterloo and st_pancras for the London Network Map
cd $project_root_dir
cd cmd
echo
echo START TEST 14
go run . ../assets/tests/input/london-network.map waterloo st_pancras 2 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test14.txt
echo 
echo 
echo "END OF TEST"
# Test 15 - It can find more than one route for 2 trains between waterloo and st_pancras for the London Network Map
cd $project_root_dir
cd cmd
echo
echo START TEST 15
go run . ../assets/tests/input/london-network.map waterloo st_pancras 3 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test15.txt
echo 
echo 
echo "END OF TEST"
# Test 16 - It can find more than one route for 2 trains between waterloo and st_pancras for the London Network Map
cd $project_root_dir
cd cmd
echo
echo START TEST 16
go run . ../assets/tests/input/london-network.map waterloo st_pancras 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test16.txt
echo 
echo 
echo "END OF TEST"
# Test 17 - It finds only a single valid route for 1 train between waterloo and st_pancras in the London Network Map 
cd $project_root_dir
cd cmd
echo
echo START TEST 17
go run . ../assets/tests/input/london-network.map waterloo st_pancras 1 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test17.txt
echo 
echo 
echo "END OF TEST"
# Test 18 - displays "Error" on stderr when too few command line arguments are used
cd $project_root_dir
cd cmd
echo
echo START TEST 18
go run . ../assets/tests/input/london-network.map waterloo st_pancras  | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test18.txt
echo 
echo 
echo "END OF TEST"
# Test 19 - It displays "Error" on stderr when too many command line arguments are used.
cd $project_root_dir
cd cmd
echo
echo START TEST 19
go run . ../assets/tests/input/london-network.map waterloo st_pancras 4  fdsag 41  | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test19.txt
echo 
echo 
echo "END OF TEST"
# Test 20 - It works with additional tricky cases.
cd $project_root_dir
cd cmd
echo
echo START TEST 20
# go run . ../assets/tests/input/london-network.map waterloo st_pancras 1  | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test20.txt
echo 
echo 
echo "END OF TEST"
# Test 21 - displays "Error" on stderr when the start station does not exist. #PASS need to add details!.
cd $project_root_dir
cd cmd
echo
echo START TEST 21
go run . ../assets/tests/input/london-network.map MadeUpStation st_pancras 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test21.txt
echo 
echo 
echo "END OF TEST"
Test 22  -displays "Error" on stderr when the end station does not exist.
cd $project_root_dir
cd cmd
echo
echo START TEST 22
go run . ../assets/tests/input/london-network.map waterloo MadeUpStation 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test22.txt
echo 
echo 
echo "END OF TEST"
Test 23  -displays "Error" on stderr when the start and end station are the same.
cd $project_root_dir
cd cmd
echo
echo START TEST 23
go run . ../assets/tests/input/london-network.map waterloo waterloo 4 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test23.txt
echo 
echo 
echo "END OF TEST"
# Test 24 - displays "Error" on stderr when any of the coordinates are not valid positive integers.
cd $project_root_dir
cd cmd
echo
echo START TEST 24
go run . ../assets/tests/input/err7.txt waterloo st_pancras 1  | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test24.txt
echo 
echo 
echo "END OF TEST"
Test 25 - displays "Error" on stderr when two stations exist at the same coordinates.
cd $project_root_dir
cd cmd
echo
echo START TEST 25
go run . ../assets/tests/input/err8.txt waterloo st_pancras 1  | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test25.txt
echo 
echo 
echo "END OF TEST"
# Test 26 - displays "Error" on stderr when the map does not contain a "connections:" section.
cd $project_root_dir
cd cmd
echo
echo START TEST 26
go run . ../assets/tests/input/err9.txt waterloo st_pancras 1  | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test26.txt
echo 
echo 
echo "END OF TEST"
# Test 27 - It completes the movements in no more than 8 turns for 10 trains between jungle and desert
cd $project_root_dir
cd cmd
echo
echo START TEST 27
go run . ../assets/tests/input/jungle-desert.map jungle desert 10 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test27.txt
echo 
echo 
echo "END OF TEST"
# Test 28 - It completes the movements in no more than 8 turns for 9 trains between small and large
cd $project_root_dir
cd cmd
echo
echo START TEST 28
go run . ../assets/tests/input/small-large.map small large 9 | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test28.txt
echo 
echo 
echo "END OF TEST"
# Test 29 - displays "Error" on stderr when duplicate routes exist between two stations, including in reverse.
cd $project_root_dir
cd cmd
echo
echo START TEST 29
go run . ../assets/tests/input/err10.txt waterloo st_pancras 1  | tee myfile.txt
echo 
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
cat test29.txt
echo 
echo 
echo "END OF TEST"
