# scapi3

The Scratch 3 API, implemented in Golang.

This project is a Golang rewrite of
[this library](https://github.com/ErrorGamer2000/scratch3-api). Now featuring
even less callback hell or whatever.

This is still in progress.

## API coverage

### User sessions

- [x] User session login
- [ ] User session verify
- [ ] User session projects
- [ ] User session comment
- [ ] Cloud session creation
- [ ] Cloud session close
- [ ] Cloud session get variable
- [ ] Cloud session set variable
- [ ] Cloud session variable change event

### Rest API

Following [this article](https://en.scratch-wiki.info/wiki/Scratch_API):

- [x] API GET `/health`
- [x] API GET `/news`
- [X] API GET `/projects/count/all` (This is broken on Scratch's end)
- [X] API GET `/projects/<project_id>`
- [ ] API GET `/projects/<project_id>/remixes`
- [ ] API GET `/studios/<studio_id>`
- [ ] API GET `/studios/<studio_id>/projects`
- [ ] API GET `/studios/<studio_id>/managers`
- [ ] API GET `/studios/<studio_id>/curators`
- [ ] API GET `/studios/<studio_id>/activity?dateLimit=<date>`
- [ ] API GET `/studios/<studio_id>/comments`
- [ ] API GET `/studios/<studio_id>/comments/<comment_id>`
- [ ] API GET `/studios/<studio_id>/comments/<comment_id>/replies`
- [ ] API GET `/proxy/featured`
- [ ] API GET `/users/<username>`
- [ ] API GET `/users/<username>/favorites`
- [ ] API GET `/users/<username>/followers`
- [ ] API GET `/users/<username>/following`
- [ ] API GET `/users/<username>/messages/count`
- [ ] API GET `/users/<username>/projects`
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
