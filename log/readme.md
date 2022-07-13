### 总结

| 包 | 平均运行时间 | 平均消耗内存 | 每次操作申请内存数 |   
| --- | --- | --- | --- | 
| zap | 605.3 ns/op | 4 B/op | 0 allocs/op | 
| zap（sugar） | 615.4 ns/op | 4 B/op | 0 allocs/op | 
| zerolog | 48447 ns/op |176 B/op | 1 allocs/op| 
| logrus | 50959 ns/op |  641 B/op|  17 allocs/op