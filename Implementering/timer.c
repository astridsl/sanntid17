#include <time.h>
#include "timer.h"


static time_t startTime = -1;


void timer_start(void){
	startTime = time(0);
}


int evIsTimeout(void){
	if (startTime < 0){
		return 0;
	}
	
	timer_t now = time(0);
	
	if (now - startTime >= 3 ){
		startTime = -1;
		return 1;
	}
	else{
		return 0;
	}
}


//* CLOCKS_PER_SEC
