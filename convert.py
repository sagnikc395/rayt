from PIL import Image
import argparse
import sys
from pathlib import Path
from tqdm import tqdm


def convert_ppm_to_jpg(input_path: Path, output_path: Path = None) -> None:
    """
    Convert a PPM image to JPG format with progress indication.

    Args:
        input_path: Path to the input PPM file
        output_path: Optional path for the output JPG file
    """
    try:
        # If output path is not specified, create one from input path
        if output_path is None:
            output_path = input_path.with_suffix(".jpg")

        print("Reading PPM file...")
        with tqdm(total=2, desc="Converting") as pbar:
            # Open and convert the image
            with Image.open(input_path) as img:
                if img.mode != "RGB":
                    img = img.convert("RGB")
                pbar.update(1)

                # Save as JPG
                img.save(output_path, "JPEG", quality=95)
                pbar.update(1)

        print(f"\nSuccessfully converted {input_path} to {output_path}")

    except FileNotFoundError:
        print(f"Error: Input file '{input_path}' not found", file=sys.stderr)
        sys.exit(1)
    except Exception as e:
        print(f"Error during conversion: {e}", file=sys.stderr)
        sys.exit(1)


def main():
    parser = argparse.ArgumentParser(
        description="Convert PPM image to JPG format",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )

    parser.add_argument("input", type=Path, help="Path to input PPM file")
    parser.add_argument(
        "-o",
        "--output",
        type=Path,
        help="Path to output JPG file (optional)",
        default=None,
    )

    args = parser.parse_args()
    convert_ppm_to_jpg(args.input, args.output)


if __name__ == "__main__":
    main()
