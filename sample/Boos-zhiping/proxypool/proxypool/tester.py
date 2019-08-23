import asyncio
import aiohttp
import time
import sys
import  requests
from proxypool import setting
try:
    from aiohttp import ClientError
except:
    from aiohttp import ClientProxyConnectionError as ProxyConnectionError
from proxypool.db import RedisClient


class Tester(object):
    def __init__(self):
        self.redis = RedisClient()
    
    async def test_single_proxy(self, proxy):
        """
        测试单个代理
        :param proxy:
        :return:
        """
        conn = aiohttp.TCPConnector(verify_ssl=False) #不验证ssl证书
        async with aiohttp.ClientSession(connector=conn) as session:
            try:
                if isinstance(proxy, bytes):  #如果是bytes类型就做utf-8编码
                    proxy = proxy.decode('utf-8')
                real_proxy = 'http://' + proxy #组合成url
                print('正在测试', proxy)
                # 测试代理是否有效      setting.TEST_URL是测试网站     超时时间为十五秒 不允许重定向(改变网页)
                async with session.get(setting.TEST_URL, proxy=real_proxy, timeout=15, allow_redirects=False) as response:
                    if response.status in setting.VALID_STATUS_CODES:  #如果是200 302 则代理可用
                        self.redis.max(proxy)
                        print('代理可用', proxy)
                    else:
                        self.redis.decrease(proxy) #响应码不合法则减分
                        print('请求响应码不合法 ', response.status, 'IP', proxy)
            except (ClientError, aiohttp.client_exceptions.ClientConnectorError, asyncio.TimeoutError, AttributeError):
                self.redis.decrease(proxy)  #捕捉错误  客户端错误 连接器错误   单线程调度器超时 属性错误
                print('代理请求失败', proxy)

    def run(self):
        """
        测试主函数
        :return:
        """
        print('测试器开始运行')
        try:
            count = self.redis.count()
            print('当前剩余', count, '个代理')
            for i in range(0, count, setting.BATCH_TEST_SIZE):
                start = i
                stop = min(i + setting.BATCH_TEST_SIZE, count)  #取最小值
                print('正在测试第', start + 1, '-', stop, '个代理')
                test_proxies = self.redis.batch(start, stop)    #取ip列表
                loop = asyncio.get_event_loop()                 #单线程并发
                #新建一个测试代理函数列表
                tasks = [self.test_single_proxy(proxy) for proxy in test_proxies]
                loop.run_until_complete(asyncio.wait(tasks)) #把线程列表加入并发池
                sys.stdout.flush() #等待并发池内的线程执行完毕,清空内存
                time.sleep(5)      #定时睡眠
        except Exception as e:
            print('测试器发生错误', e.args)
