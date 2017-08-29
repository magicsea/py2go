
type  BattleCard struct {}

func (self *BattleCard)  checkBuffLimit( buffId, bigRound) {
    buffInfo = d_buff_info.getInfo(buffId)
    if  self.rub_BigRound != bigRound {
        self.rub_BigRound = bigRound
        self.roundUsedBuffDict = {}
    }
    if  buffId !in self.roundUsedBuffDict {
        self.roundUsedBuffDict[buffId] = 1
        return true
    } else if  buffInfo.RoundLimit > self.roundUsedBuffDict[buffId] {
        self.roundUsedBuffDict[buffId] += 1
        return true
    }
    if  cardId !in self.ownCardDict {
        self.ownCardDict[cardId] = cardNum
    } else {
        self.ownCardDict[cardId] += cardNum
    }
    return false
}
func (self *BattleCard)  checkBuffLimit2( buffId, bigRound) {
    buffInfo2 = d_buff_info.getInfo(buffId)
    if  self.rub_BigRound != bigRound {
        self.rub_BigRound = bigRound
        self.roundUsedBuffDict = {}
    }
    a = 3
    if  a==b&&(value==1||value== 2)&&c==d {
        a = 4
    }
    if ( X==-233|| X== -232|| X== -247) {
        fmt.Println('a')
    }
    cardInfo.equipList=append(cardInfo.equipList,1)
    _list=append(_list,(cardId, equipList))
    if  buffId !in self.roundUsedBuffDict {
        self.roundUsedBuffDict[buffId] = 1
        return true
    } else if  buffInfo.RoundLimit > self.roundUsedBuffDict[buffId] {
        self.roundUsedBuffDict[buffId] += 1
        return true
    }
    return true
}
func (self *BattleCard)  testFor() {
    for _, _guildId :=range lists {
        fmt.Println(_guildId)
    }
    for  uid,_ :=range guild.guildMember {
        fmt.Println(uid)
    }
    for _, guild :=range self.guildDict {
        fmt.Println(guild)
    }
    for  uid, member :=range guild.guildMember {
        fmt.Println(uid, member)
    }
    for  i:=0; i<100; i++ {
        fmt.Println(i)
    }
}