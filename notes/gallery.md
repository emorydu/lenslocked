# Gallery
1. Create a new gallery with a title.
2. Upload images to a gallery.
3. Delete images from a gallery.
4. Update the title of a gallery.
5. View a gallery (so we can share it with others).
6. Delete a gallery.
7. View a list of galleries we are allowed to edit.


This might give us the following views we need to create:
- Create a new gallery
- Edit a gallery
- View an existing gallery
- View a list of all of our galleries

new, edit, show, index

CRUD - Create, Read, Update, Delete

We will also need controllers (aka HTTP handlers) to support these views:
- New and Create to render and process a new gallery form.
- Edit and Update to render and process a form to edit a gallery.
- Show to render a gallery.
- Delete to delete a gallery.

- An HTTP handler to process image uploads.
- An HTTP handler to remove images from a gallery.


And finally, we need a way to persist data in our models package, and this will
need to support the following:
- Creating a gallery
- Updating a gallery
- Querying for a gallery by ID
- Querying for all galleries with a user ID
- Deleting a gallery
- Creating an image for a gallery.
- Deleting an image from a gallery.
