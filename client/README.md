# GitHub API

## ■ref

https://developer.github.com/v3/

https://developer.github.com/v3/repos/#list-user-repositories

## ■curl
<pre>
curl -i https://api.github.com/users/sky0621/repos?per_page=200 > /tmp/sky0621repos.txt
cat /tmp/sky0621repos.txt | grep clone_url
    "clone_url": "https://github.com/sky0621/aws-describe-prj.git",
    "clone_url": "https://github.com/sky0621/dockerfile-gowithdep.git",
    "clone_url": "https://github.com/sky0621/dockerfile-gowithglide.git",
    "clone_url": "https://github.com/sky0621/dockerfile-gowithgrpc.git",
    "clone_url": "https://github.com/sky0621/document.git",
    "clone_url": "https://github.com/sky0621/fons.git",
    "clone_url": "https://github.com/sky0621/fun-grpcgo.git",
    "clone_url": "https://github.com/sky0621/gitlab-seek-expert.git",
    "clone_url": "https://github.com/sky0621/go-algo.git",
　　　　・
　　　　・
　　　　・
    "clone_url": "https://github.com/sky0621/study-spring-rest.git",
    "clone_url": "https://github.com/sky0621/study-springboot.git",
    "clone_url": "https://github.com/sky0621/study-springboot-cli.git",
    "clone_url": "https://github.com/sky0621/study-springboot-maven.git",
    "clone_url": "https://github.com/sky0621/study-terraform.git",
</pre>

<pre>
curl -u sky0621 -i https://api.github.com/user/repos?per_page=200
</pre>
