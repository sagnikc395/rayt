build:
  mkdir -p dist/
  go build -o ./dist/rayt

run:
  ./dist/rayt > image.ppm 
  open ./image.ppm
