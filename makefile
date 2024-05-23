london: 
	@echo "Starting the london map"
	@cd cmd && go run .  london-network.map waterloo st_pancras 10
jungle: 
	@echo "Starting the jungle map"
	@cd cmd && go run .  jungle-desert.map jungle treetop 5
music: 
	@echo "Starting the beethoven-part map"
	@cd cmd && go run .  beethoven-part.map beethoven part 5
num: 
	@echo "Starting the jungle map"
	@cd cmd && go run .  two-four.map two four 5

terminus:
	@echo "Starting the terminus map"
	@cd cmd && go run .  beginning-terminus.map beginning terminus 5

test:
	@echo "Start tests"
	@./test.sh

init:
	@echo "Giving presmmison to the test script"
	@chmod +x test.sh


