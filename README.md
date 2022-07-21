# scapi3

The Scratch 3 API, implemented in Golang.

This project is a Golang rewrite of
[this library](https://github.com/ErrorGamer2000/scratch3-api). Now featuring
even less callback hell or whatever.

This is still in progress.

## API coverage

### User sessions

- [X] User session login
- [X] User session verify
- [ ] User session projects
- [X] User session comment (Broken)
- [ ] Cloud session creation
- [ ] Cloud session close
- [ ] Cloud session get variable
- [ ] Cloud session set variable
- [ ] Cloud session variable change event

### Rest API

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
- [ ] API GET `/users/<username>/projects/<project_id>`
- [ ] API GET `/users/<username>/projects/<project_id>/studios`
- [ ] API GET `/users/<username>/projects/<project_id>/comments`
- [ ] API GET `/users/<username>/projects/<project_id>/comments/<comment_id>`
- [ ] API GET `/users/<username>/projects/<project_id>/comments/<comment_id>/replies`
- [ ] API GET `/users/<username>/studios/curate`
- [ ] API GET `/accounts/checkusername/<username>`
- [ ] API GET `/explore/projects?q=<query>&mode=<mode>&language=<language_code>`
- [ ] API GET `/explore/studios?q=<query>&mode=<mode>&language=<language_code>`
- [ ] API GET `/search/projects?q=<query>&mode=<mode>&language=<language_code>`
- [ ] API GET `/search/studios?q=<query>&mode=<mode>&language=<language_code>`
