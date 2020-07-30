#!/bin/sh
set -e
echo '=================== Install Fig ==================='
DOWNLOAD_FIG_VERSION=${HUGO_VERSION:-0.0.2}
GITHUB_DEPLOY_REPOSITORY=${GITHUB_REMOTE_REPOSITORY:-$GITHUB_REPOSITORY}
GITHUB_DEPLOY_BRANCH=${GITHUB_BRANCH:-"gh-pages"}
echo "Installing Fig $DOWNLOAD_FIG_VERSION"
wget -O /usr/local/bin/fig https://github.com/hobochild/fig/releases/download/${DOWNLOAD_FIG_VERSION}/fig &&\
chmod +x /usr/local/bin/fig
echo '=================== Create deploy key to push ==================='
mkdir /root/.ssh
ssh-keyscan -t rsa github.com > /root/.ssh/known_hosts && \
echo "${GIT_DEPLOY_KEY}" > /root/.ssh/id_rsa && \
chmod 400 /root/.ssh/id_rsa
echo '=================== Update all submodules ==================='
git submodule init
git submodule update --recursive --remote
echo '=================== Build site ==================='
fig
echo '=================== Publish to GitHub Pages ==================='
cd dist
remote_repo="git@github.com:${GITHUB_DEPLOY_REPOSITORY}.git" && \
remote_branch=${GITHUB_DEPLOY_BRANCH} && \
echo "Pushing Builds to $remote_repo:$remote_branch" && \
git init && \
git remote add deploy $remote_repo && \
git checkout $remote_branch || git checkout --orphan $remote_branch && \
git config user.name "${GITHUB_ACTOR}" && \
git config user.email "${GITHUB_ACTOR}@users.noreply.github.com" && \
git add . && \
echo -n 'Files to Commit:' && ls -l | wc -l && \
timestamp=$(date +%s%3N) && \
git commit -m "Automated deployment to GitHub Pages on $timestamp" > /dev/null 2>&1 && \
git push deploy $remote_branch --force && \
rm -fr .git && \
cd ../
echo '=================== Done  ==================='
