#ifndef QUEUE_H_
#define QUEUE_H_

//Setter alle elementene i matrisa lik 0.
void queue_del_queue(void);

//Setter alle elementene på en rad lik 0.
void queue_del_one_floor(void);

//Setter valgt etasje (rad) og knapp (kolonne) lik 1 i matrisa.
void queue_add_to_queue(int floor, int button);

//Finner neste etasjedestinasjon.
int queue_get_next_floor(void);

//Sjekker om valgt etasje (rad) og knapp (kolonne) er bestilt, altså lik 1 i matrisa.
int get_button_pressed(int floor, int button);



#endif /*QUEUE_H_*/
