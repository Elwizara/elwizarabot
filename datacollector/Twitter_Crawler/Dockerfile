# Use an official Python runtime as a parent image
FROM python:3.6

# Set the working directory to /app
WORKDIR /app

#Add requirements.txt to image
ADD requirements.txt /app

# Install any needed packages specified in requirements.txt
RUN pip install --trusted-host pypi.python.org -r requirements.txt

# Copy the current directory contents into the container at /app
ADD . /app 

CMD ["python", "__init__.py"]

#how to run  
#docker build -t tarekbadr/twittercrawler:1.0 .  
#docker run --rm -ti  --net=host -v $(pwd)/conf.json:/app/Twitter_Crawler/conf.json tarekbadr/twittercrawler:1.0 /bin/bash
#docker push tarekbadr/twittercrawler:1.0

