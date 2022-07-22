<img src="assets/logo.svg" alt="I am so sorry for making this." width="96"/>

# scapi3

The Scratch 3 API, implemented in Golang.

This project is a Golang rewrite of
[this library](https://github.com/ErrorGamer2000/scratch3-api) (plus more
stuff). Now featuring even less callback hell or whatever.

This is still in progress.

# What's not working

- Commenting returns 401 unauthorized
- Cloud sessions are created successfully, but the server seems to shut down the
  websocket connection right after the handshake is sent.

## API coverage

### User sessions

- [X] User session login
- [X] User session verify
- [X] User session comment (Broken)
- [X] Cloud session creation (Broken)
- [X] Cloud session close
- [X] Cloud session get variable
- [ ] Cloud session set variable
- [X] Cloud session variable change event (Broken)

### Rest API (Complete!)

Following [this article](https://en.scratch-wiki.info/wiki/Scratch_API):

- [X] API GET `/health`
- [X] API GET `/news`
- [X] API GET `/projects/count/all` (This is broken on Scratch's end)
- [X] API GET `/projects/<project_id>`
- [X] API GET `/projects/<project_id>/remixes`
- [X] API GET `/studios/<studio_id>`
- [X] API GET `/studios/<studio_id>/projects`
- [X] API GET `/studios/<studio_id>/managers`
- [X] API GET `/studios/<studio_id>/curators`
- [X] API GET `/studios/<studio_id>/activity?dateLimit=<date>`
- [X] API GET `/studios/<studio_id>/comments`
- [X] API GET `/studios/<studio_id>/comments/<comment_id>`
- [X] API GET `/studios/<studio_id>/comments/<comment_id>/replies`
- [X] API GET `/proxy/featured`
- [X] API GET `/users/<username>`
- [X] API GET `/users/<username>/favorites`
- [X] API GET `/users/<username>/followers`
- [X] API GET `/users/<username>/following`
- [X] API GET `/users/<username>/messages/count`
- [X] API GET `/users/<username>/projects`
- [X] API GET `/users/<username>/projects/<project_id>`
- [X] API GET `/users/<username>/projects/<project_id>/studios`
- [X] API GET `/users/<username>/projects/<project_id>/comments`
- [X] API GET `/users/<username>/projects/<project_id>/comments/<comment_id>`
- [X] API GET `/users/<username>/projects/<project_id>/comments/<comment_id>/replies`
- [X] API GET `/users/<username>/studios/curate`
- [X] API GET `/accounts/checkusername/<username>`
- [X] API GET `/explore/projects?q=<query>&mode=<mode>&language=<language_code>`
- [X] API GET `/explore/studios?q=<query>&mode=<mode>&language=<language_code>`
- [X] API GET `/search/projects?q=<query>&mode=<mode>&language=<language_code>`
- [X] API GET `/search/studios?q=<query>&mode=<mode>&language=<language_code>`
