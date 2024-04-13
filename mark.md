目录结构思路
（不要再纠结了）

service -- 服务入口，区分不同领域
    seat 座位
    packagecard 套餐
    member 会员
    shop 商店
    

biz --  不同领域的逻辑，实现具体的 【现实中】 的操作
            🌰 座位服务（seat），订座操作（booking），里面包含 对不同仓库对操作

data -- 仓库，存储不同数据类型