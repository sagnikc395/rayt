#include <SDL.h>
#include <stdio.h>

#define WIDTH 900
#define HEIGHT 600


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

    SDL_Rect rect = (SDL_Rect){200,200,200,200};
    Uint32 COLOR_WHITE = SDL_MapRGB(surface->format, 255, 255, 255);

    //draw something to the surface
    SDL_FillRect(surface, &rect, COLOR_WHITE);

    //update window surface
    SDL_UpdateWindowSurface(window);

    //small event loop to keep the window open and responsive
    SDL_Event e;
    int quit = 0 ;
    while(!quit) {
        while(SDL_PollEvent(&e)){
            if(e.type == SDL_QUIT) {
                quit = 1;
            }
        }
        SDL_Delay(20); //preventing usy looping
    }

    SDL_DestroyWindow(window);
    SDL_Quit();
    return 0;
}
