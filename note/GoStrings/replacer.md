### replacer

#### compare.go
> 只有一个compare方法,但是不建议用这个方法，有这个方法主要为了跟bytes包对称
> ![compare](compare.png)

#### search.go 
> 使用Boyer-Moore string search algorithm(BM算法)匹配字符串，这里查找的时候只匹配一次，即找到相应的字符串就返回，如果要继续查找，则继续传入参数```s[match:]```即可
> ![stringFinder](stringFinder.png)
> 需要构造坏字符表和好后缀表， 好后缀构造代码如下
> ![goodSuffixSkip](goodSuffixSkip.png)

#### replace.go 字符串替换
- 定义
> ![replacer](replacer.png)
> 接下来定义了四个结构体并实现了replace接口
>
> singleStringReplacer: 
> 单个字符串替换，即只有一个旧字符串需要替换为新字符串，使用了前面的stringFinder来查找相应的旧字符串。 使用了append生成新字符串
>
> genericReplacer: 
> 字符串组的替换，所有旧字符串都为字符串(而不是byte)，具体代码还没看 todo
>
> byteReplacer:
> 多组byte的替换
>
> byteStringReplacer：
> 旧byte替换为新string或byte，使用了copy生成新字符串，应该是考虑到新的byte用append不太合适(append会先复制，而copy则直接写入)

- NewReplacer(oldnew ...string) *Replacer 
> 首先 传入的字符串数据格式为 ```[]stirng{old1, new1, old2,new2,...}```
>
> 通过传入的字符串数据的具体情况，决定要生成的是上面的哪一个replacer,具体如下：
> * oldnew 长度为2，且旧字符串不为byte --> singleStringReplacer
> * 多个替换组，且所有旧字符串不为byte --> genericReplacer
> * 所有旧的为byte，所有新的也为byte --> byteReplacer
> * 所有旧的为byte，所有新的不全为byte --> byteStringReplacer
>
- (r *Replacer) Replace(s string) string
- (r *Replacer) WriteString(w io.Writer, s string)
> ![Replace&WriteString](Replace&WriteString.png)

##### genericReplacer todo