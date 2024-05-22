
# Test london.map
cd cmd
echo
echo START TEST
echo
go run .  london-network.map waterloo st_pancras 4 | tee myfile.txt
echo 
echo END TEST
cd ..
cd assets/tests/expected
echo 
echo "EXPECTED RESULT"
echo 
cat correct_format.txt
echo 
echo
echo "END OF TEST"
echo 


## We can just run the progamm and the files under assets/test/input
## 


