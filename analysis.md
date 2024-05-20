# PART 1: Url Scraper
The main purpose of this program is to get an exhaustive list of the paths of a url.
To achieve this we have two approachs: 
- Relying on the sitemap: 
    <br> With this approach it is straighforward to find the paths if we suppose that the sitemap.xml is well documented.
    If not we might miss some paths.
- Scraping every pages to find links:
   <br> With this approach we can read every html page of the website following links to have a list of its paths.
    But if there is some paths that are not referenced anywhere, we will miss them.

What I choose to implement, is a mix of both method to maximize the coverage.

# PART 2: 
## Docker 

- Here I have wrotten a Dockerfile in which there is a non-root user who own the copied project.
- I have used the "ENTRYPOINT" statement in the Dockerfile to make the docker container take the script arguments.
- I had run a vulnerability scan and it report is in vulnerabilityScan/report.txt file.

## Kubernetes
- I have created a deployment manifest in file kubernetes/deployment.yaml.

# PART 3: CI/CD
In this part I have created a pipeline which build, scan, push and deploy the app to a kubernetes cluster.

I kept it very basic but for improvement we can add logic for dev and prod build, and also block deploy on prod if security scan have a number of HIGH vulnerabilities which exceed a define max.  

# PART 4: Domains Sanitizer
In order to sanitize the url, we have three main steps:
- Remove scheme, trailing dot and paths:
   <br> To achieve this I used regex with sed and gsub(of awk)
- Ensure having lowercase letters: 
    <br> To achieve this I used tolower with awk and tr
- Take only the domain Second Level Domain and the Top Level Domain:
    <br> To achieve this I used split with awk/cut and choose the relevents fields.

# Bonus 
I can related this case study to:
- Algorithm 
- Regular expression 