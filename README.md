## rayt

a raytracer written in Rust , taken from [writing a raytracer in a weekend](https://raytracing.github.io/books/RayTracingInOneWeekend.html)

~~current progress is that we are able to render a blue-to-whiote gradinet image using a camera focussed on ray Y coordinate.~~

~~we have generated the sphere colored according to its normal vectors.~~

we have a resulting render of normals-colored sphere with ground.

> ![Image](./image.jpg)

### Workflow

```sh
cargo run > etc/image.ppm
```

```python
python convert.py etc/image.ppm > image.jpg
```
