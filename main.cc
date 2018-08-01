#include <stdio.h>



int main(){
    if (argc > 2){

        for (int i = 0; i < argc; i++ )
        {
            printf("%s", argv[i]);
        }
    }
}
