添加工具
db.role.updateMany({"type":{"$in":["admin", "manage"]}}, {"$addToSet":{"tool_limits":"tool_mk_lottery"}})
db.default_role_2b.updateMany({"type":{"$in":["admin", "manage"]}}, {"$addToSet":{"tool_limits":"tool_mk_lottery"}})
db.default_role_2c.updateMany({"type":{"$in":["admin", "manage"]}}, {"$addToSet":{"tool_limits":"tool_mk_lottery"}})

db.getCollection("tool_master4").insert({
_id: "tool_mk_lottery",
categoryId: "0513a4da-1801-11eb-bcc3-4e36b81d7ae2",
courses: [],
createTime: "2021-01-15 10:19:00",
desc: "抽奖大转盘",
enable: true,
examples: [],
homePage: "/material-bank/app-center/official-account-management",
icon: "https://cf-markting-test.oss-cn-hangzhou.aliyuncs.com/system/material/template/eaaeed47-2784-4ed7-aedd-704cc74614ce/0j5p6o7lauv1603767601333.png",
name: "抽奖大转盘",
platforms: [
NumberInt("1")
],
ruleGroups: [],
updateTime: "2021-01-18 10:19:00"
})

在tool_category中加入对应的tool