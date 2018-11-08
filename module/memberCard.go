package module

import (
	"fmt"
	"sctek.com/typhoon/th-platform-gateway/common"
	"sctek.com/typhoon/th-platform-gateway/manageMq"
	"time"
)

type MemberCard struct {
	Id                int       `xorm:"not null pk autoincr INT(11)"`
	CorpId            int       `xorm:"not null default 0 comment('机构ID') INT(11)"`
	Level             int       `xorm:"not null default 0 comment('会员卡等级') TINYINT(4)"`
	LevelName         string    `xorm:"not null default '' comment('等级名称') VARCHAR(255)"`
	CoverImg          string    `xorm:"not null default '' comment('会员卡卡面') VARCHAR(255)"`
	CoverImgWxUrl     string    `xorm:"default '' comment('微信cdn地址') VARCHAR(255)"`
	UpCondition       int64     `xorm:"not null default 0 comment('升至本等级需要消费的累计金额') BIGINT(20)"`
	DownConditionTime int       `xorm:"not null default 0 comment('自开卡时间起降级的时间纬度') INT(11)"`
	DownCondition     int64     `xorm:"not null default 0 comment('降级条件') BIGINT(20)"`
	IsPrimary         int       `xorm:"not null default 0 comment('主要的，不能删除的') TINYINT(4)"`
	Created           time.Time `xorm:"not null DATETIME"`
	Updated           time.Time `xorm:"not null DATETIME"`
	Deleted           int       `xorm:"not null default 0 comment('删除状态（0：正常，1：删除）') TINYINT(1)"`
	IsTimeDelete      int       `xorm:"not null default 0 comment('是否定时删除（0：否，1：是）') TINYINT(1)"`
}

func (a MemberCard) TableName() string {
	return "member_card"
}

//根据会员等级发送
func (m *MemberCard) SendMessageForGrade(grade, message string) error {
	common.Log.Infoln("根据会员等级把消息压入mq队列")
	engine := common.DB
	list := make([]MemberCard, 0)
	err := engine.Where("level=?", grade).Find(&list)
	if err != nil {
		common.Log.Infoln(err)
		manageMq.ExampleLoggerOutput(err.Error())
		return err
	}
	//test
	if len(list) <= 0 {
		for i := 0; i < 10000; i++ {
			list = append(list, MemberCard{})
		}
	}
	for _, v := range list {
		_ = v
		phone := "15920038315"
		message := "are you sure??"
		message = fmt.Sprintf("{\"phone\":\"%q\",\"message\":\"%q\"}", phone, message)
		manageMq.GlobalMq.Publish("fanout", "")
	}
	return nil
}
