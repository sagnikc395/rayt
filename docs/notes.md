### PPM Image Format:

- Whenever you start a renderer, you need a way to see an image. The most straightforward way is to write it to a file. The catch is, there are so many formats. Many of those are complex. I always start with a plain text ppm file.
- There are some things to note in this code:
  1. The pixels are written out in rows.
  2. Every row of pixels is written out left to right.
  3. These rows are written out from top to bottom.
  4. By convention, each of the red/green/blue components are represented internally by real-valued variables that range from 0.0 to 1.0. These must be scaled to integer values between 0 and 255 before we print them out.
  5. Red goes from fully off (black) to fully on (bright red) from left to right, and green goes from fully off at the top (black) to fully on at the bottom (bright green). Adding red and green light together make yellow so we should expect the bottom right corner to be yellow.

### Creating a Image File:

- Build the target target/release/rayt and redirect the output to image.ppm
- and convert the image to jpg using the script (convert.py)

### Adding a Progress Indicator:

- useuful to track the progress of a long render, and also to possibly identify a run that's stalled out due to an infinite loop or other problems.
- Our program outputs the image of the standard output stream , so we leave that and instead write to the logging output stream.

### vec3 class:

- TBD

## Ray Class

- All ray tracers have is the ray class and a computation of what color is seen along a ray.
- Think of ray as a function P(t) = A + tB.
- P is a 3d posn along a line in 3D.
- A is the ray origin, b is the ray direction.
- t -> real number.

### Sending Ray Into the Scene:

- Ray Tracer in its core sends rays thorugh pixels and computes the color seen in the direction of those rays. The involved steps are:

  - Calculate the ray from the "eye" thorugh the pixel.
  - Determine which objcts the ray intersects.
  - Compute a color for the closest intersection point.

- Aspect Ratio -> ratio of image width / ratio of image height.
- for this we are taking a 16/9 aspect ratio -> 1.7778

- Seen we have the aspect ratio , it's easier to set the image's width and the aspect ratio and then using this to calculate for its height.

  - easier to scale up or down the image by changing the image width, and it won;t throw off our desired aspect ratio.
  - have to make sure that when we solve for the iamge height the resulting image is at-least 1.

- Need to also set up a virtual viewport through which to pass our scene rays.
- The viewport is a virtual rectangle in the 3d world, that contains the grid of image pixel locations.
- If pixels are spaced the same distance horizontally as they are vertically, the viewport than bounds them will have the same aspect ratio as the rendered image.
- Pixel Spacing -> distance between 2 adjacent pixels and square pixels are standard.

### Define the Camera Center:

- a point in 3D space from which all scene rays will originate (also commonly referred to as the eye point).
- the vector from the camera center to the viewport center will be orthogonal to the viewport.
- Initially set the distance between the viewport and the camera center point to be 1 unit. This distance is often referred to as the focal length.
- For simplicity, start at (0,0,0).
- also have the y-axis to go up, the x-axis to the right and negative z-axis pointing in the viewing direction.
- while we have 3D space that has the convetions above, this will conflict with our image coordinates, where we ant have the zeroth pixel in the top-left and work our way down to the last pixel at the bottom right. This means that our image cooridinate Y-axis is inverted -> Y will increase going down the image.

- First start at the upper left pixel(pixel ,0,0) ,scan left-to-right across each row and then scan row-by-row ,top-to-bottom. To help navigate the pixel grid, we will use a vector from the left edge to the right-edge(Vu) and a vector from the upper edge to the lower edge(Vv).

## CPP Features:

## Ray Class

- All ray tracers have is the ray class and a computation of what color is seen along a ray.
- Think of ray as a function P(t) = A + tB.
- P is a 3d posn along a line in 3D.
- A is the ray origin, b is the ray direction.
- t -> real number.

### Sending Ray Into the Scene:

- Ray Tracer in its core sends rays thorugh pixels and computes the color seen in the direction of those rays. The involved steps are:

  - Calculate the ray from the "eye" thorugh the pixel.
  - Determine which objcts the ray intersects.
  - Compute a color for the closest intersection point.

- Aspect Ratio -> ratio of image width / ratio of image height.
- for this we are taking a 16/9 aspect ratio -> 1.7778

- Seen we have the aspect ratio , it's easier to set the image's width and the aspect ratio and then using this to calculate for its height.

  - easier to scale up or down the image by changing the image width, and it won;t throw off our desired aspect ratio.
  - have to make sure that when we solve for the iamge height the resulting image is at-least 1.

- Need to also set up a virtual viewport through which to pass our scene rays.
- The viewport is a virtual rectangle in the 3d world, that contains the grid of image pixel locations.
- If pixels are spaced the same distance horizontally as they are vertically, the viewport than bounds them will have the same aspect ratio as the rendered image.
- Pixel Spacing -> distance between 2 adjacent pixels and square pixels are standard.

### Define the Camera Center:

- a point in 3D space from which all scene rays will originate (also commonly referred to as the eye point).
- the vector from the camera center to the viewport center will be orthogonal to the viewport.
- Initially set the distance between the viewport and the camera center point to be 1 unit. This distance is often referred to as the focal length.
- For simplicity, start at (0,0,0).
- also have the y-axis to go up, the x-axis to the right and negative z-axis pointing in the viewing direction.
- while we have 3D space that has the convetions above, this will conflict with our image coordinates, where we ant have the zeroth pixel in the top-left and work our way down to the last pixel at the bottom right. This means that our image cooridinate Y-axis is inverted -> Y will increase going down the image.

- First start at the upper left pixel(pixel ,0,0) ,scan left-to-right across each row and then scan row-by-row ,top-to-bottom. To help navigate the pixel grid, we will use a vector from the left edge to the right-edge(Vu) and a vector from the upper edge to the lower edge(Vv).

-
