#include "elev.h"
#include "statemachine.h"
#include <stdio.h>


int main() {
    // Initialize hardware
    if (!elev_init()) {
        printf("Unable to initialize elevator hardware!\n");
        return 1;
    }

	initialize();

	while(1) {
		eventManager();
		driveElevator();
	}

    return 0;
}
