# 多分支开发

因为多处了越南版、台湾版等版本，会遇到多个分支代码管理问题。
在这里讨论多分支开发。

## 词汇表
* 发行分支：针对不同发行商定制的release分支，比如 越南分支(soha) 腾讯分支(tencent) 等

## 开发分支选择
1. 如果A功能是某*发行分支*独有的并且没有和其他功能产生关联，那么可以在这个发行分支基础上开新分支做开发
2. 如果A功能是多所有发行版本共有，那么在default分支开新分支做开发

## 一般流程

### 常规流程
一般来说 default 分支是会领先其他发行分支的。比如 default 分支现在T4版本已完成，正在开发T5版本。
而越南分支需要发布T3版本的内容，那么在 default 分支的 T3 这个 tag merge 到越南分支。

```
default branch 开发分支
T1 ----> T2 -----> T3 ------> T4 --- t5.feature 1, t5.feature 2 --->
                      \
                        \
                          \
                           |
T1 ----> T2 -----> T2.1 - T3 
soha branch 越南版发行分支
```

### 紧急任务 
假设但前正在开发 T5版本，  feature 1 和 feature 2 已经完成。但是由于某些特别原因，需要把 feature 2 发布到 越南版本。
那么就不能简单的合并，因为这样会把不能发布的 feature 1也合并到 soha 分支。
这种情况可以手动把 feature 2 这个 commit 作为 patch ，打到 soha 分支。
```
hg export -o t5.feature2.patch  -r t5.feature2
hg check soha
hg import t.feature2.patch
```

###  bug 修复
bug修复建议的做法是，版本回推到引入bug的一个commit，以这个comiit为基础建立hotfix-XXX分支。
受影响的分支 merge hotfix-XXX 分支。

