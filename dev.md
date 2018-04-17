
# 生成 api model controller
bee api world -conn="root:Lawuyou@tcp(127.0.0.1:8889)/study-dev"

# 生成数据
bee generate migration course -driver=mysql -fields="id:auto"

# 生成model
bee generate model base -fields="name:string,age:int"
bee generate model package/base -fields="name:string,age:int"

