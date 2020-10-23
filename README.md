# linkerd-slo
SLOs with prometheus and linkerd from https://buoyant.io/2020/10/21/kubernetes-slos-with-prometheus-linkerd/

# usage
```bash
# install linkerd
https://github.com/linkerd/linkerd2/releases/
# check if cluster can handle linkerd
$ linkerd check --pre
# install linkerd server
$ linkerd install | kubectl apply -f -
# check again
$ linkerd check
# run demo app
$ curl -sL https://run.linkerd.io/emojivoto.yml \
| linkerd inject - \
| kubectl apply -f -
...
```

# todos
- [x] install all the things
- [x] run checks
- [ ] fix bug in demo app (not meshed) or deploy something else
- [ ] graph SLOs
