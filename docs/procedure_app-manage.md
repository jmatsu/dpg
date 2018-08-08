# app-manage

`app-manage` command has these subcommands

- `expose` [help](procedure.app-manage.expose.txt)
- `on-feature-branch` [help](procedure.app-manage.on-feature-branch.txt)
- `on-deploy-branch` [help](procedure.app-manage.on-deploy-branch.txt)

*expose command*
 
This infers some required values for `app-manage` from the current working directory and environment variables. This is a kind of helper so don't need to run itself in some cases.

Output format is like

```
key1='value1'
key2='value2'
```

So it can be used as env file.  
Also, `--prefix` option can add the prefix to each lines. So `dpg expose --prefix 'export '` can help shell scripts.

*on-feature-branch command*

This uploads your application and create a distribution. 

By default, the distribution name, the upload message and the release note are automatically generated from the branch and git logs.  

*on-deploy-branch command*

This destroys a distribution which has been associated with a branch is recently merged. By default, the target distribution is determined by git log.

## CI example

These example are based on CircleCI.

## Download the single binary

```.circleci/config.yml

jobs:
  build_apk:
    <<: android_env
    steps:
      - checkout
      - run:
          name: Enable dpg command in this job
          command: |
            curl -sL "https://raw.githubusercontent.com/jmatsu/dpg/master/install.bash" | bash
            echo "export PATH=$PWD:$PATH" >> $BASH_ENV
      - run:
          name: Assemble and deploy
          command: |
            ./gradlew :app:assembleDebug

            source <(dpg procedure app-manage expose --prefix "export " --feature-branch --token <your api token> --app-owner <your app's owner name>) 
            dpg procedure app-manage on-feature-branch --app app/build/outputs/apk/debug/app-debug.apk 

            # Or

            dpg procedure app-manage on-feature-branch --app app/build/outputs/apk/debug/app-debug.apk --token <your api token> --app-owner <your app's owner name>

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

              # Or

              dpg procedure app-manage on-feature-branch --app /tmp/app-debug.apk --token <your api token> --app-owner <your app's owner name>
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

              # Or

              dpg procedure app-manage on-deploy-branch --token <your api token> --app-owner <your app's owner name> --android --app-id <your app id>

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
The following commands would be your help. 

```
dpg procedure app-manage expose --feature-branch > .env
env_opts=$(dpg procedure app-manage expose --feature-branch --prefix "-e " | xargs)

dpg procedure app-manage expose > .env
env_opts=$(dpg procedure app-manage expose --prefix "-e " | xargs)
```