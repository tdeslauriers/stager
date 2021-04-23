## Stager

**Support application written in Go to automate adding photos to the family gallery site.**

1. It will run a cron job to look for new files in a given directory where the user (me) has dumped the photos they want to upload to the site.

1. It will read the files and gather their exif data.

1. It will use the exif data to determine which album to map the picture to.

* If no correct album exists, it will create one in the database.

1. It will rename the .jpgs with UUIDs (since that will be the URL identifier): [uuid].jpg.

1. I will extract each images thumbnail and name it with the matching UUID: [uuid]_thumb.jpg.

1. It will load the database with the images names, album id, and other metadata.

1. It will sftp the files to the server directory that hosts the images for the site.

*Note: files will not be published/visible on the site without adminsitrative action on the site itself to prevent mistakes and mischeif.*