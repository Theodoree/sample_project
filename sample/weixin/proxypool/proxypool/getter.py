from proxypool.db import RedisClient
from proxypool.crawler import Crawler
import sys
from proxypool.setting import POOL_UPPER_THRESHOLD
class Getter():
    def __init__(self):
        self.redis = RedisClient()
        self.crawler = Crawler()
    
    def is_over_threshold(self):
        """
        判断是否达到了代理池限制
        """
        if self.redis.count() >= POOL_UPPER_THRESHOLD:
            return True
        else:
            return False
    
    def run(self):
        print('获取器开始执行')
        # self.redis.Clear_Cookie()
        if not self.is_over_threshold(): #判断是否达到代理池限制
            for callback_label in range(self.crawler.__CrawlFuncCount__): #这个__CrawlFuncCount__返回该列表中有多少个方法

                callback = self.crawler.__CrawlFunc__[callback_label] #通过列表下标获取爬虫方法
                # 获取代理
                proxies = self.crawler.get_proxies(callback)  #这里传入是方法名
                sys.stdout.flush()          #清理内存
                for proxy in proxies:       #遍历代理,添加到redis
                    self.redis.add(proxy)
