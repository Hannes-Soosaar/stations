london: 
	@echo "Starting the london map"
	@cd cmd && go run .  london-network.map waterloo st_pancras 1

terminus: 
	@echo "Starting the terminus map"
	@cd cmd && go run .  beginning-terminus.map beginning far 2