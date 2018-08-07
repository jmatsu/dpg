# app-manage

`app-manage` command has these subcommands

- `expose` [help](procedure.app-manage.expose.txt)
- `on-feature-branch` [help](procedure.app-manage.on-feature-branch.txt)
- `on-deploy-branch` [help](procedure.app-manage.on-deploy-branch.txt)

*expose command*
 
This infers some required values for `app-manage` from the current working directory and environment variables. This is a kind of helper so don't need to run itself in some cases.

*on-feature-branch command*

This uploads your application and create a distribution. 

By default, the distribution name, the upload message and the release note are automatically generated from the branch and git logs.  

*on-deploy-branch command*

This destroys a distribution which has been associated with a branch is recently merged. By default, the target distribution is determined by git log.

## CI example

These example are based on CircleCI.

## Without Remote Docker

`.circleci/config.yml` is like below. 

```.circleci/config.yml
jobs:
  build_apk:
    <<: android_env
    steps:
      - checkout
      - run:
          name: Assemble and save it
          command: |
            ./gradlew :app:assembleDebug
            mkdir -p /tmp
            cp app/build/outputs/apk/debug/app-debug.apk /tmp/app-debug.apk
      - persist_to_workspace:
          root: /tmp
          paths:
            - app-debug.apk
  on_feature_branch:
    docker:
      - image: jmatsu/dpg:{{ version or latest }}
    working_directory: ~/{{ anywhere you want }}
    steps:
      - checkout
      - attach_workspace:
          at: /tmp
      - run:
          name: Upload an apk and create a distribution by app-manage procedure.
          command: |
              source <(dpg procedure app-manage expose --prefix "export " --feature-branch --token <your api token> --app-owner <your app's owner name>) 
              dpg procedure app-manage on-feature-branch --app /tmp/app-debug.apk
  on_deploy_branch:
    docker:
      - image: jmatsu/dpg:{{ version or latest }}
    working_directory: ~/{{ anywhere you want }}
    steps:
      - checkout
      - run:
          name: Destroy the associated distribution by app-manage procedure.
          command: |
              source <(dpg procedure app-manage expose --prefix "export " --token <your api token> --app-owner <your app's owner name> --android --app-id <your app id>)
              dpg procedure app-manage on-deploy-branch

workflows:
  version: 2
  every_branch:
    jobs:
      - build_apk
      - on_feature_branch:
          requires:
            - build_apk
      - on_deploy_branch:
          filters:
            branches:
              only: /(master|develop)/
```

## With Remote Docker

In this case, you cannot use `expose` command. Some of exposed values are not usable because it will see the environment in the docker container.

```
/repo/to/root
   |-- .dpg
   |    └-- Dockerfile
   └-- app <android app module>
```

And then, `.dpg/Dockerfile` should be like the following. 

```dpg/Dockerfile
FROM jmatsu/dpg:latest

ENV DPG_APP_FILE_PATH android.apk

ARG APK_PATH

COPY $APK_PATH android.apk
```

`.circleci/config.yml` is like below. 

```.circleci/config.yml
jobs:
  on_feature_branch:
    <<: *android_env
    steps:
      - checkout
      - run: ./gradlew assembleDebug
      - setup_remote_docker:
          docker_layer_caching: true
      - run:
          name: Copy your build artifact to under Docker daemon's space
          command: cp app/build/outputs/apk/debug/app-debug.apk dpg/app-debug.apk
      - run:
          name: Build a docker image for this build
          command: docker build --build-arg APK_PATH=app-debug.apk -t $CIRCLE_BRANCH:dpg .dpg
      - run: |
          name: Upload an apk and create a distribution by app-manage procedure.
          command: |
              docker run --rm $CIRCLE_BRANCH:{{ version or latest }} dpg procedure app-manage on-feature-branch --token <your api token> --app-owner <your app's owner name> --android --distribution-name "$CIRCLE_BRANCH"
  on_deploy_branch:
    <<: *any_env
    steps:
      - checkout
      - setup_remote_docker:
          docker_layer_caching: true
      - run:
          name: Destroy the associated distribution by app-manage procedure.
          command: |
              branch_name="$(git log --format=%s --merges -1 | sed 's/^.* from [^\/]*\/\(.*\)$/\1/')"
              docker run $env_opts --rm jmatsu/dpg:{{ version or latest }} dpg procedure app-manage on-deploy-branch --token <your api token> --app-owner <your app's owner name> --android --app-id <your app id> --distribution-name "$branch_name"
  
workflows:
  version: 2
  every_branch:
    jobs:
      - on_feature_branch
      - on_deploy_branch:
          filters:
            branches:
              only: /(master|develop)/
```
