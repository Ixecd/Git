## 基于gitflow的项目管理

1. 基于当前master分支创建一个新的分支 develop
2. 基于develop分支，创建一个功能分支 feature/hello-world
3. 在当前feature/hello-world分支上，进行相对应的功能开发
    - 如果这时候线上出现了Bug，需要立即处理Bug
    1. 开发只完成一半，但是不想提交，临时保存修改到堆栈区 git stash
    2. 从 master 建立 hotfix 分支
    3. 在 hotfix 分支上修复 Bug