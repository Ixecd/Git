## 基于gitflow的项目管理

1. 基于当前master分支创建一个新的分支 develop
2. 基于develop分支，创建一个功能分支 feature/hello-world
3. 在当前feature/hello-world分支上，进行相对应的功能开发
    - 如果这时候线上出现了Bug，需要立即处理Bug
    1. 开发只完成一半，但是不想提交，临时保存修改到堆栈区 git stash
    2. 从 master 建立 hotfix 分支
    3. 在 hotfix 分支上修复 Bug
    4. 修复完成后，合并到 master 分支，并且删除 hotfix 分支
    5. 切换到 develop 分支，合并 hotfix 分支的修改
    6. 切换到 feature/hello-world 分支，合并 develop 分支的修改
    7. git stash pop 从堆栈区恢复修改，继续开发
    8. 开发完毕之后，push origin feature/hello-world分支
    9. 创建Pull Request，合并到develop分支
    10. 合并完成之后基于develop分支创建release/1.0.0分支
    11. go build -v . 构建后，部署二进制文件，并测试
    12. 如果测试有问题，直接在当前release/v1.0.0分支上修复问题
    13. 测试通过之后，将release/v1.0.0分支合并到develop和master分支
    14. 合并完成之后，删除 feature/hello-world 分支

```bash
# 单元测试指令
# -race 检测是否有数据竞态
# -cover 检测代码覆盖率
# -coverprofile=./coverage.out 生成覆盖率报告
# -timeout=10m 超时时间10分钟
# -short 只运行快速测试（跳过耗时长的测试）
# -v  verbose 详细输出
# ./... 递归测试当前目录及所有子目录下的包
go test -race -cover  -coverprofile=./coverage.out -timeout=10m -short -v ./...
```

```bash
# 生成测试覆盖率报告
# -func 以函数为单位展示每个函数的覆盖率
# -html 生成html格式的覆盖率报告
# ./coverage.out 指定覆盖率数据文件（通常由go test-coverprofile=./coverage.out生成
go tool cover -func ./coverage.out
```