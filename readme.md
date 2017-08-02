# Docker-Compose Examples

There are different examples of processing variables when using Docker-Compose in this repository.

* The `.env` file in the docker-compose directory
* The named variable file in a separate directory
* An override `-compose.yml` file
* Extending a service

> The app is a simple Go command line app that does nothing more than print the values of the envirnoment variables. The "work" piece of the app is in the `cmd` directory. Nothing special to look at there. Built with [Cobra](https://github.com/spf13/cobra) if you're interested.

## The `.env` file
The `env_file` property of the compose file is in relation to the compose file being used to build the services.

You'll notice in the directories `qa` and `prod`, there are separate docker-compose.yml files (the names are irrelevant, they could be named anything). Both compose files use the same `env_file` property of `.env`. This is because the `.env` file is in the same directory as the compose file being used. 

This is good to keep the compose and variables together in the same space to avoid selecting the incorrect env file, or adjusting the compose file to point to an incorrect env file as it is relative

You may also notice that the `build` `context` is pointed to the parent directory by way of `../`. This is because they are all using the same Dockerfile to build the image.

## A file in a separate directory
The `docker-compose.stage.yml` file uses a specific folder/file for the `env_file` property pointing to a single directory for all environment files with specially named files. These files are nothing more than moving the assignment to a separate file rather than defining them in the compose file. No special syntaxm or file name needed. Docker-Compose simply reads the file line by line as a "Key=Value" list _separated by new lines_.

## An override `compose.yml` file
The `docker-compose` cli command takes multiple `-f` parameters. Doing this allows you to override, or supplement, the values in the first compose file. 

`docker-compose -f docker-compose.yml -f prod/docker-compose.prod.yml up` 

Doing this, the docker-compose.yml file is read, _then_ the docker-compose.prod.yml file is read. The values defined in the second compose file override the ones in the first if there are any conflicts. Also, if you define _additional_ variables, or services, or anything, those will be _added_ to the first compose file.

As such, if you take a look at the `docker-compose.prod.yml` file in the `prod` folder, you'll notice that there is a service defined that is named different than the one defined in the `docker-compose.yml` file. 
> If they are not named different, and run as show above, you will get an error since compose appends the 2 services, but cannot as they are named the same.

Because you are defining a service in the override file, both the base, and the override services are created and run by calling the `up` command (as shown above)

You may also notice that the `context` used in `prod` is `.` and not `../` as it is in the `docker-compose.qa.yml` file (which is also located in a subdirectory). This is because the context is based off the location of the initial compose file being used, and thus, at the root directory. Notice the value of the `ANOTHER_ENV_ITEM` in the output of the program. It is "this is the root env". Again, because the context is based off the initial compose file, and the `env_file` property only indicates `.env` without a directory. Compare this to the `docker-compose.qa.yml` file to see the difference.

## Extending a service
Another option is to **extend** the service(s). This is similar to the override, but instead of appending to the service, you redefine the values you want changed.

For an example, have a look at the `docker-compose.uat.yml` file. Comparing it to the `docker-compose` file, you'll notice that not everything is defined again. Instead, I've only redefined the values/properties that needed to change. The values that aren't defined in the extends compose, get pulled in from the original.

You'll also notice, that I've added another service that extends the other service. Note that the `enxtends` property only needs the name of the service if they're defined in the same compose file. If you're extending from a different file, you'll need to list the `file` and `service` that you're extending.

> Note: extends **does not** allow extending volumes or links. Those must always be defined in whole at the place the service being run is defined.

## Some Notes
If you look at the `env/stage` and `docker-compose.stage.yml` files, you'll see they both define the variable `TEST_ENV_ITEM`.

The value in the compose file will override the value in the file no matter what order you place the `env_file` and `environment` properties. This is similar to defining a variable in a Dockerfile and overriding it from the command line during the `docker run -e 'TEST_ENV_ITEM=something'...`


Huge thanks out to the team at Runnable for putting together a [blog post](https://runnable.com/docker/advanced-docker-compose-configuration) that inspired these examples.