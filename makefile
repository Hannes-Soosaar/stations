london: 
	@echo "Starting the london map"
	@cd cmd && go run .  london-network.map waterloo st_pancras 10
terminus: 
	@echo "Starting the terminus map"
	@cd cmd && go run .  beginning-terminus.map beginning far 99
jungle: 
	@echo "Starting the jungle map"
	@cd cmd && go run .  jungle-desert.map jungle treetop 50
