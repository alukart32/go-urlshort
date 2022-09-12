The goal is to create an http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for /dogs to https://www.somesite.com/a-story-about-dogs we would look for any incoming web requests with the path /dogs and redirect them.

Redirects rules can be stored in
- YAML (default: config.yaml)
- JSON

Original task: https://courses.calhoun.io/lessons/les_goph_04