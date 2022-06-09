## Stager

**Support application written in Go to automate adding photos to the family gallery site.**

1. It will run a cron job to look for new files in a given directory where the user (me) has dumped the photos they want to upload to the site.

1. It will read the files and gather their exif data.

1. It will compare hash of image to db image hashes to prevent duplication

1. It will use the exif data to determine which album to map the picture to.

    * If no correct album exists, it will create one in the database.

1. It will rename the .jpgs with UUIDs (since that will be the URL identifier): [uuid].jpg.

1. ~~It will extract each image's thumbnail from exif data.~~
    
    * Not all pictures have exif data; not all exif has thumbnails/sometimes it is corrupted; not all Thumbnails preserve correct image rotation. **Too manual.**

    * Makes more sense to *this usecase* to generate thumbnails by resizing the original image.

1. It will load the database with the image , thumbnail, name, and other metadata.

### Notes:
* *Files will not be published/visible on the site without adminsitrative action on the site itself to prevent mistakes and mischeif.*

* *There are easier/faster ways to prevent duplication than hashing, but I wanted to give it a whirl.*

* *go-exiftool is a wrapper for a perl program exiftool which must be installed on the machine.*

* *golang's image/jpeg decoding/encoding does **NOT** preserve exif data.*