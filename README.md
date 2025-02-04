GoDIP (Go Digital Image Processing) is a simple image-processing library written in Go.
It currently supports:
- Loading and saving PNG, JPG, GIF and PPM images
- Multi-channel, single-channel and float images
- Simple math operations
- Transformations: resizing, rotation, flipping
- Blurring: Gaussian and mean
- Convolutions (and generating some kernels like Gaussian or circular)
- Correction: changing brightness, contrast, gamma correction
- Edge detection using gradients
- Morphology: Min filter, max filter, dilatation, erosion, opening and closing
- Thresholding: using a single threshold or a range, gaussian adaptive thresholding
- Image types with convenient API allowing for simple creation of more filters
