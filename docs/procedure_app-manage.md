# app-manage

`app-manage` command has these subcommands

- `expose` [help](procedure.app-manage.expose.txt)
- `on-feature-branch` [help](procedure.app-manage.on-feature-branch.txt)
- `on-deploy-branch` [help](procedure.app-manage.on-deploy-branch.txt)

`expose` command infers some required values for `app-manage` from the current working directory and environment variables. This is kinda helper so don't need to run always.

## CI example

These example are based on CircleCI.

## Without Remote Docker

`.circleci/config.yml` is like below. 

```.circleci/config.ylm
jobs:
  build_apk:
    <<: android_env
    steps:
      - checkout
      - run: ./gradlew assembleDebug
      - persist_to_workspace:
          root: /tmp/workspace
          paths:
            - app/build/outputs/apk/debug/app-debug.apk
  on_feature_branch:
    docker:
      - image: jmatsu/dpg:latest
    working_directory: ~/{{ ORG_NAME }}/{{ REPO_NAME }}
    steps:
      - checkout
      - attach_workspace:
          at: /tmp/workspace
      - run:
          name: Upload an apk and create a distribution by app-manage procedure.
          command: |
              cp /tmp/workspace/app/build/outputs/apk/debug/app-debug.apk app.apk
              source <(dpg procedure app-manage expose --prefix "export " --token <your api token> --app-owner <your app's owner name>) 
              dpg procedure app-manage on-feature-branch --app app.apk
  on_deploy_branch:
    docker:
      - image: jmatsu/dpg:latest
    working_directory: ~/{{ ORG_NAME }}/{{ REPO_NAME }}
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

COPY <application file> android.apk
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
          command: cp app/build/outputs/apk/debug/app-debug.apk dpg/android.apk
      - run:
          name: Build a docker image for this build
          command: docker build -t $CIRCLE_BRANCH:dpg .dpg
      - run: |
          name: Upload an apk and create a distribution by app-manage procedure.
          command: |
              env_opts=$(dpg procedure app-manage expose --prefix "export " --token <your api token> --app-owner <your app's owner name>)
              docker run $env_opts --rm $CIRCLE_BRANCH:dpg dpg procedure app-manage on-feature-branch
  on_deploy_branch:
    <<: *any_env
    steps:
      - checkout
      - setup_remote_docker:
          docker_layer_caching: true
      - run:
          name: Destroy the associated distribution by app-manage procedure.
          command: |
              env_opts=$(dpg procedure app-manage expose --prefix "-e " --token <your api token> --app-owner <your app's owner name> --android --app-id <your app id> | tr "\n" " ")
              docker run $env_opts --rm jmatsu/dpg:latest dpg procedure app-manage on-deploy-branch
  
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