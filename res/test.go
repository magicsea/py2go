type  BattleCard struct {}

func (self *BattleCard)  checkBuffLimit(self, buffId, bigRound) {
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
    return False
}
func (self *BattleCard)  checkBuffLimit2(self, buffId, bigRound) {
    buffInfo2 = d_buff_info.getInfo(buffId)
    if  self.rub_BigRound != bigRound {
        self.rub_BigRound = bigRound
        self.roundUsedBuffDict = {}
    }
    a = 3
    a = 4
    if  buffId !in self.roundUsedBuffDict {
        self.roundUsedBuffDict[buffId] = 1
        return true
    } else if  buffInfo.RoundLimit > self.roundUsedBuffDict[buffId] {
        self.roundUsedBuffDict[buffId] += 1
        return true
    }
    return true
}
func (self *BattleCard)  testFor(self) {
    for  _guildId :=range lists {
        print _guildId
    }
    for  uid,_ :=range guild.guildMember {
        print uid
    }
    for _, guild :=range self.guildDict {
        print guild
    }
    for  uid, member in guild.guildMember.iteritems() {
        print uid, member
    }
    for  i:=0; i<100; i++ {
        print i
    }
}