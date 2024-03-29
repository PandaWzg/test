package response

import (
	"strconv"
)

var ErrorMap map[string]string

type Error struct {
	Code    int
	Message string
}

type Errors struct {
	Error map[string]Error
}

func init() {
	ErrorMap = map[string]string{
		"101": "深切哀悼在抗击新冠肺炎疫情斗争的烈士和逝世同胞，4月4日停服一天",
		"10000": "无效的API",
		"10001": "未定义错误",
		"10002": "已知错误",
		"10003": "系统错误",
		"10004": "数据错误",
		"10005": "内测版本，仅对邀请用户开放",
		"10006": "操作频繁，请稍后重试",
		"10007": "数据已存在",
		"10008": "网关服务异常",
		"10009": "网关服务异常",
		"10010": "服务器升级，请稍后重试",
		"10011": "配置信息异常",
		"0":     "没有找到记录",
		"20000": "登录已过期",
		"20001": "数据格式错误",
		"20002": "文章标题不能为空",
		"20003": "客户端类型不能为空",
		"20004": "上传失败，请稍后重试",
		"20005": "不支持该类型文件",
		"20006": "文件上传过大，请压缩后上传",


		"30001": "请求次数超过限制",
		"30002": "您已成功注册，预计3月18日开始公测，敬请期待",
		"30003": "首次登录，请使用微信方式登录",
		"30004": "未开通渠道合伙人",
		"40001": "sign校验失败",
		"40002": "支付网关异常，请稍后尝试",
		"40003": "非法请求，被拒绝",
		"40004": "Token错误",
		"40005": "time参数异常",

		"50001": "两次输入密码不一致",
		"50002": "红包领取失败，已经领取过",
		"50003": "红包类型错误",


		//dog
		"60001": "金币不足",
		"60002": "位置已满",
		"60003": "位置信息异常",
		"60004": "无法购买该等级",
		"60005": "无法执行该操作",
		"60006": "合成失败，五大洲犬条件不足",
		"60007": "招财犬无法存入仓库",
		"60008": "无法回收，至少要保留一只狗狗",
		"60009": "今日领取次数不足",
		"60010": "领取时间未到，请稍后领取",
		"60011": "请升级狗狗解锁",

		//receive gold
		"60108": "金币领取次数不足",
		"60109": "无法领取，时间未到",
		"60110": "金币领取失败",
		"60111": "金币翻倍领取失败",

		"60200": "宝箱领取次数不足",


		"61000": "存入仓库异常，请稍后操作",
		"61005": "存入仓库失败，请稍后操作",
		"61006": "仓库已满",
		"61007": "仓库背包容量已满",

		//user
		"62000": "用户无效",
		"62001": "手机号格式不正确",
		"62002": "验证码不正确",
		"62003": "验证码请求过于频繁，请2分钟后重试",
		"62004": "账号异常",
		"62005": "已经绑定过师父，请勿重复绑定",
		"62006": "师父ID不存在",
		"62007": "没有邀请人",
		"62008": "已完成实名认证",
		"62009": "余额不足，无法提现",
		"62010": "提现金额无效",
		"62011": "未完成实名认证，无法提现",
		"62012": "手机号码未绑定",
		"62013": "验证码已过期",
		"62014": "恭喜您提交成功，我们将会在3个工作日内尽快审核",
		"62015": "实名认证通过，微信提现成功",
		"62016": "实名认证失败，请确认与微信实名一致",
		"62017": "无效的绑定关系",
		"62018": "升级到34级后可进行5元提现",
		"62019": "手机号已绑定其它账号",
		"62020": "昵称含有敏感词",
		"62021": "昵称长度超出限制",
		"62022": "昵称已存在",
		"62023": "每天只能修改一次昵称",
		"62024": "支付账号信息填写不完整",
		"62025": "该身份证号已签署过其他账号",
		"62026": "该身份证号未申请过电签",
		"62027": "电签未完成，请重新验证",
		"62028": "验证码超出发送上限",
		"62029": "申请失败",
		"62030": "您已提交申请",
		"62031": "无法修改昵称",
		"62032": "签到次数不足，请到探险大厅签到",
		"62033": "识别到该账户为作弊或虚拟用户，提现失败",
		"62034": "请先绑定手机号",
		"62035": "微信提现额本月已超限，请选择支付宝",



		//lottery
		"63000": "抽奖券已领取",
		"63001": "抽奖券不足，请完成任务领取",
		"63002": "狗狗位置满了，请至少预留一个位置来抽奖",
		"63003": "狗狗位置满了，请至少预留一个位置来兑换单身犬",
		"63004": "碎片不足，无法合成",
		"63005": "请至少保留一只狗狗在格子中",

		//task
		"64000": "今日已签到",
		"64001": "已经领取过",
		"64002": "领取失败，条件未达成",
		"64003": "今日任务已达上限",
		"64004": "抽奖次数不足",
		"64005": "任务无效，请认真完成",
		"64006": "该任务今日已完成",

		//兑换
		"65000" : "狗粮不足，无法兑换，再继续攒狗粮吧！",
		"65001" : "兑换失败",
		"65002" : "您已兑换过当前金额",

		//外部合作api
		"66001" : "您已经在其它渠道试玩过该游戏",
	}
}

func GetError(code int) string {
	key := strconv.Itoa(code)

	if msg, ok := ErrorMap[key]; ok {
		return msg
	}

	return "未知错误"
}
