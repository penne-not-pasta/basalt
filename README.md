# Basalt
Basalt is a software that intents to make managing cloud services easier.

## Basalt Cloud Storage
Basalt is a Software that can implement multiple Cloud Providers into one App for easy access to your Cloud services. Basalt is developed by the oreon project, for the Oreon Linux distribution.

## Support For Cloud Providers
Basalt Currently supports only **Google Drive** since this project is just starting but we are going to add more very soon

## Legal.
Copyright (C) 2025 oreonproject
This program comes with ABSOLUTELY NO WARRANTY;
This is free software, and you are welcome to redistribute it under certain conditions;
Everything in this project is licensed under the GPLv3 license, unless specified otherwise

# Notice.

If you manage to run this,  it will ask you to go to a URL and authorize from there. Do so and then when it redirects you to that ``localhost`` URL, that URL will have a `code` query parameter in it which you're gonna copy the value of and paste it in the program (its gonna be hanging there and waiting for the code), if everything goes well (hopefully,  Please Google) it will list the files in your Google Drive.

Now this is only temporary because if we are going to actually use this we're gonna have to make small HTTP server which captures that ``code`` parameter automatically.
