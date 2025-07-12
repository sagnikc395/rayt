#include <SDL.h>
#include <stdio.h>
#include <math.h>
#include <stdbool.h>

#define WIDTH 900
#define HEIGHT 600

struct Circle{
    double x;
    double y;
    double r;
};

void FillCircle(SDL_Surface* surface, struct Circle circle,Uint32 color){
    double radius_squared = pow(circle.r,2);
    for(double x = circle.x-circle.r;x<= circle.x+circle.r;x++){
        for(double y=circle.y-circle.r;y<=circle.x+circle.r;y++){
            double distance_squared = pow(x- circle.x,2.0) + pow(y-circle.y,2.0);
            if(distance_squared < radius_squared) {
                SDL_Rect pixel = (SDL_Rect){x,y,1,1};
                SDL_FillRect(surface, &pixel,color);
            }
        }
    }
}


int main() {

    // Initialize SDL
    if(SDL_Init(SDL_INIT_VIDEO) <0) {
        fprintf(stderr,"SDL_Init Error: %s\n", SDL_GetError());
        return 1;
    }

    // Create a window
    SDL_Window* window = SDL_CreateWindow("rayt", SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, WIDTH, HEIGHT, SDL_WINDOW_SHOWN);
    if (!window) {
           fprintf(stderr,"SDL_CreateWindow Error: %s\n", SDL_GetError());
           SDL_Quit();
           return 1;
    }
    SDL_Surface* surface = SDL_GetWindowSurface(window);
    if (!surface) {
            fprintf(stderr,"SDL_GetWindowSurface Error: %s\n", SDL_GetError());
            SDL_DestroyWindow(window);
            SDL_Quit();
            return 1;
    }

    //create a rectangle
    // SDL_Rect rect = (SDL_Rect){200,200,200,200};
    // define a white
    Uint32 COLOR_WHITE = SDL_MapRGB(surface->format, 255, 255, 255);
    //define color black
    Uint32 COLOR_BLACK = SDL_MapRGB(surface->format,0,0,0);

    //x,y 0,0 as we have to cover the top corners
    SDL_Rect erase_rect = (SDL_Rect){0,0,WIDTH,HEIGHT};

    //draw something to the surface
    // SDL_FillRect(surface, &rect, COLOR_WHITE);

    struct Circle circle= {200,200,80};
    FillCircle(surface,circle,COLOR_WHITE);

    //update window surface
    SDL_UpdateWindowSurface(window);

    //redraw the circle ?
    bool simulation_running = true;
    SDL_Event event;
    while(simulation_running){
        // poll for the current independent events
        while(SDL_PollEvent(&event)) {
            if(event.type == SDL_QUIT) {
                simulation_running = false;
            }
            if(event.type == SDL_MOUSEMOTION && event.motion.state != 0) {
                //mouse is clicked and event motion,
                // can get the x coordinate and send the x coordinate
                circle.x = event.motion.x;
                circle.y = event.motion.y;
            }
        }
        SDL_FillRect(surface,&erase_rect,COLOR_BLACK);
        FillCircle(surface,circle,COLOR_WHITE);
        //update window surface
        SDL_UpdateWindowSurface(window);
        SDL_Delay(100);
    }


    //small event loop to keep the window open and responsive
    // SDL_Event e;
    // int quit = 0 ;
    // while(!quit) {
    //     while(SDL_PollEvent(&e)){
    //         if(e.type == SDL_QUIT) {
    //             quit = 1;
    //         }
    //     }
    //     SDL_Delay(20); //preventing usy looping
    // }

    SDL_DestroyWindow(window);
    SDL_Quit();
    return 0;
}
