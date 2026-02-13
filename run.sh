#!/bin/bash
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
CYAN='\033[0;36m'
NC='\033[0m'
DIR=$(pwd)


run_server ()
{
	echo -e "${GREEN}Building...${NC}"
	go build -o go-app . && echo -e "${GREEN}Running...${NC}" && ./go-app &
}

watch_changes() {
	pkill -f go-app
	run_server 
	last_event=0
	echo -e "${YELLOW}Watching for changes...${NC}"
	inotifywait -m -r \
		--exclude '(^|/)build(/|$)' \
		-e create,modify,delete "$DIR" |
		while read -r _; do
			now=$(date +%s)
			if (( now - last_event >= 3 )); then
				echo -e "${CYAN}Reloading...${NC}"
				last_event=$now
				pkill -f go-app
				run_server 
			fi
		done
	}
alias run="watch_changes"
