# Configuration variables
src_dir := "src"
build_dir := "build/debug"
cc := "gcc"
obj_name := "rayt"

include_paths := "-I include/SDL2 -I include/SDL2_image -I include/SDL2_ttf"
library_paths := "-L lib/SDL2 -L lib/SDL2_image -L lib/SDL2_ttf"
compiler_flags := "-std=c11 -Wall -O0 -g -lm"
linker_flags := "-lSDL2 -lSDL2_image -lSDL2_ttf"

# Build rule
all:
    mkdir -p {{build_dir}}
    {{cc}} {{compiler_flags}} {{linker_flags}} {{include_paths}} {{library_paths}} $(find {{src_dir}} -name '*.c') -o {{build_dir}}/{{obj_name}}
default:
    clang src/main.c -Iinclude -Iinclude/SDL2 -Llib/SDL2 -lSDL2 -lSDL2_image -lSDL2_ttf -std=c11 -o build/debug/play
