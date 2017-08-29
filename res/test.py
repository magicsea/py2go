class BattleCard(object):
    def checkBuffLimit(self, buffId, bigRound):
        buffInfo = d_buff_info.getInfo(buffId)
        if self.rub_BigRound != bigRound:
            self.rub_BigRound = bigRound
            self.roundUsedBuffDict = {}
        if buffId not in self.roundUsedBuffDict:
            self.roundUsedBuffDict[buffId] = 1
            return True
        elif buffInfo.RoundLimit > self.roundUsedBuffDict[buffId]:
            self.roundUsedBuffDict[buffId] += 1
            return True
        if cardId not in self.ownCardDict:
            self.ownCardDict[cardId] = cardNum
        else:
            self.ownCardDict[cardId] += cardNum
        return False
    
    def checkBuffLimit2(self, buffId, bigRound):
        buffInfo2 = d_buff_info.getInfo(buffId)
        if self.rub_BigRound != bigRound:
            self.rub_BigRound = bigRound
            self.roundUsedBuffDict = {}
        a = 3
        if a==b and value in [1, 2] and c==d:
            a = 4
        if X in [-233, -232, -247]:
            print 'a'
        cardInfo.equipList.append(1)
        _list.append((cardId, equipList))
        if buffId not in self.roundUsedBuffDict:
            self.roundUsedBuffDict[buffId] = 1
            return True
        elif buffInfo.RoundLimit > self.roundUsedBuffDict[buffId]:
            self.roundUsedBuffDict[buffId] += 1
            return True
        return True
    def testFor(self):
        for _guildId in lists:
            print _guildId
        for uid in guild.guildMember.iterkeys():  
            print uid
        for guild in self.guildDict.itervalues():   
            print guild
        for uid, member in guild.guildMember.iteritems():
            print uid, member 
        for i in xrange(100): 		
            print i

