//lv 0 1 2
package main

import "fmt"

type bildetail struct {
	topside					//上边栏
	vedio					//视频
	commends				//评论
	author			  		 //作者
	recomved			   //推荐视频
}
type topside struct {
	liplink [16]string
}
type vedio struct{
	titlem string			//标题
	vv	  string			//内容是啥
	likes int				//点赞
	hoards int				//收藏
	coins int 				//投币
	aname string			//作者名
}
type author struct {
	Name string             //名字
	VIP bool                //是否是高贵的带会员
	Icon string             //头像
	Signature string        //签名
	Focus int				//粉丝
}
type commends struct {
	num int
}
type recomved struct {
	titlet []string
}							//推荐视频


//方法来了
//点赞
func (onevd *vedio)llike(){
		onevd.likes++;
}
//收藏
func (onevd *vedio)ho(){
	onevd.hoards++;
}
//投币
func (onevd *vedio)coinss(){
	onevd.coins++;
}
//三连
func (onevd *vedio)comboo(){
	onevd.likes++;
	onevd.hoards++;
	onevd.coins++;
}

//发视频
func upvd(anamee string,vtitle string)vedio{
	return vedio{aname: anamee,titlem: vtitle}
}
func main() {
	bilv := make([]bildetail, 1)
	//点开一个视频
	bilv[0].topside.liplink=[16]string{
			"主站",
			"番剧",
			"游戏中心",
			"直播",
			"会员购",
			"漫画",
			"赛事",
			"下载app",
			"搜索",
			"我的主页",
			"大会员",
			"消息",
			"动态",
			"收藏",
			"历史记录",
			"创作中心",
	}//topside模块
	bilv[0].vedio.titlem="海南热带雨林第三集：在夜晚的雨林，找到了目标物种"
	bilv[0].vedio.vv="变色树蜥咬藏狐"
	//作者
	bilv[0].author.Name="无穷小亮的科普日常"
	bilv[0].author.VIP=true
	bilv[0].author.Icon="man"
	bilv[0].author.Signature="发一些科普工作中的有趣小事"
	bilv[0].author.Focus=5775000
	bilv[0].recomved.titlet=[]string{"狐主任现场鉴定海南热带雨林里的特有物种！","和国家地理一起，探究人类文明的地理密码！","【亮记生物鉴定】抖音巨鱼是什么？","【亮记生物鉴定】近期网络热传生物合集","《网络热门科普工作者鉴定》"}
	bilv[0].commends.num=2684
	//too vegetable to be graceful


	fmt.Println(bilv)

}