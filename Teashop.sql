create database webapp;
use webapp;
drop table if exists users;
create table users(
id int primary key auto_increment,
username varchar(100) not null unique,
password varchar(100) not null,
money int not null
);
drop table if exists shop;
create table shop(
shopname varchar(100) not null unique,
address varchar(100)not null 
);
drop table if exists milktea;
create table milktea(
teaname varchar(100) not null,
basicprice int not null,
profile varchar(200) not null,
introduction varchar(1000) not null,
category varchar(100) not null
);
drop table if exists orderform;
create table orderform(
shopname varchar(100) not null,
teaname varchar(100) not null,
sweet varchar(100) not null,
cond varchar(100) not null,
addmaterial varchar(100) not null,
num int not null,
totalprice int not null
);
create table orderform(
    username varchar(100) not null ,
    shopname varchar(100) not null,
    teaname varchar(100) not null,
    sweet varchar(100) not null,
    cond varchar(100) not null,
    addmaterial varchar(100) not null,
    num int not null,
    totalprice int null
);
insert into shop(shopname,address) values("武汉荟聚中心店","长宜路1号荟聚中心");
insert into shop(shopname,address) values("武汉泛海城市广场GO店","云杉路199号泛海城市广场购物中心");
insert into shop(shopname,address) values("武汉新世界百货国贸店","建设大道566号新世界百货国贸大楼");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("浓暴柠",21,"原创粤红茶底，超浓出涩","冷500ml铁臂阿童木附体！老饕级享受，广东柠檬茶，特浓偏酸偏涩，不喜慎点。原创粤红茶底，经典红茶的浓醇饱满强强联手香水柠檬的香浓酸爽，更香更浓更上头。","暴柠家族");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("嫩暴柠",21,"全新高温烘炒嫩芽茶，清香爆发","冷500ml铁臂阿童木附体！入门级必点，不喜太酸涩口感的朋友可选。全新嫩芽茶底，经过高温烘炒工艺，茶色由绿转黄，绿茶的清香中延伸出熟火豆香余韵，与暴打35次的香水柠檬结合，多重口感，清新释放","暴柠家族");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("岩暴柠",19,"定制岩骨香茶底，香到入骨","冷500ml请尽快饮用，放置过久会影响口感。源自广东柠檬茶，口感偏酸、苦、涩、冰，我们建议选择“标准热 甜”，不喜慎点。独家定制岩骨香茶底，岩茶风味香厚浓重，介意慎点。新鲜暴打香水柠檬和暴风冰球劲爽加持，酸涩相宜，冰爽回甘。不涩不香不柠茶。","暴柠家族");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("皮暴柠",21,"皮一下真的很暴柠","冷500ml请尽快饮用，放置过久会影响口感。源自广东柠檬茶，口感偏酸、苦、涩、冰，我们建议选择“标准喜 甜”，不喜慎点。超燃暴柠茶与黄皮的清奇碰撞。颗颗手剥黄皮鲜果，去皮去核去果芯，黄皮特有的辛香搭配新鲜捶打香水柠檬片，果 入口瞬间上头。特调海盐风味柔和微咸，清奇甘爽，皮一下真的很暴柠。","暴柠家族");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("芝芝桃桃",30,"当季水蜜桃邂逅全新琥珀兰茶底","冷热皆宜冷650ml热500ml当季水蜜桃邂逅大师监制全新琥珀兰茶底，源自福建的奇兰乌龙茶种，兰香清幽，新鲜蜜桃果肉与经典芝士绵密交织，一口沁心。","果茶家族");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("多肉葡萄",29,"巨峰葡萄颗颗手剥","冷热皆宜 冷650ml热500ml当季巨峰葡萄回归，颗颗手剥，保留果肉完整肉感。搭配清雅绿妍茶底与醇香芝士，鲜爽可口。如选择原创0糖0卡糖，遇酸性水果会产生泡沫，属正常现象。波 。","果茶家族");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("雪山思乡龙眼",28,"新鲜龙眼手剥去核，蜜香清甜","冷500ml优选肉足多汁的新鲜龙眼，颗颗手剥去核。完全鲜果，不经腌制，保留龙眼原始的蜜香清甜。将顺滑冷藏热 牛乳打制成细腻冰沙，鲜甜果肉佐以嫩滑的无香精桂花冻，花香果甜一口满足。顶部为动物奶油制作的雪山奶油顶，颗颗饱满龙眼肉如夜明珠。","果茶家族");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("多肉芒芒甘露",27,"当季台芒新升级","冷热皆宜 冷/热标准杯500ml冷加大MAX杯650ml多肉芒芒甘露全新升级，甄选当季台芒，增加生打椰乳，搭配喜茶热 原创胶原弹力波波脆，更加弹润更好吸收。沿用清幽绿妍茶底，爆汁白柚果粒与新鲜台芒果肉相互映衬，酸甜中进发热带清香。热饮默认为芋圆波波，不含胶原弹力波波纯 脆。","果茶家族");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("烤黑糖波波牛乳",21,"黑糖波波融入顺滑牛乳，香甜浓郁","冷热皆宜 冷480ml热500ml黑波波搭配顺滑冷藏牛乳，波波系列奶味较为浓郁，不喜欢浓厚口感的朋友慎点。","波波家族");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("烤黑糖波波牛乳茶",22,"定制嫣红茶底与冷藏牛乳灵感特调","冷热皆宜冷480ml热500ml升级为混波波牛乳茶王冷饮650ml热饮500ml选用定制嫣红茶底与冷藏牛乳灵感特热 调，花香高扬，入口丝滑。搭配现熬黑糖波波与黑糖布蕾，带来满分波波味觉体验。","波波家族");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("芋泥波波牛乳",27,"新鲜芋泥融入冷藏牛乳，绵密润泽","冷热皆宜冷/热500ml默认冷饮，可做热。本品含有芋头，过敏人士请谨慎选择。因芋泥容易氧化，为保持最佳体热 验，请务必于一小时内饮用完毕。茶底可选冷藏牛乳/椰奶。新鲜芋泥融入顺滑冷藏牛乳，再加入颗颗Q弹的芋泥波波，绵密与润泽，尽在这一杯。","波波家族");
insert into milktea(teaname,basicprice,profile,introduction,category)  values("芋泥波波牛乳茶",28,"定制嫣红茶底，花香芋香叠加","冷热皆宜冷/热500ml升级为混波波芋泥牛乳茶王冷饮650ml热饮500ml默认冷饮，可做热。本品含有芋头，过热 敏人士请谨慎选择。因芋泥容易氧化，为保持最佳体验，请务必于一小时内饮用完毕。选用定制嫣红茶底与冷藏牛乳灵感特调，花香高扬，入口丝滑。再加入颗颗Q弹黑糖波纯 波与绵密芋泥，丰盈芋香萦绕齿间。","波波家族");