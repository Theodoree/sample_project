# Redis数据库地址
REDIS_HOST = '127.0.0.1'

# Redis端口
REDIS_PORT = 6379

# Redis密码，如无填None
REDIS_PASSWORD = 'Wanxinkun2012'

REDIS_KEY = 'http'
### 搜狗
Redis_Cookie =  'SouGou_Cookie'
Cookie_id = 'SUID'
# TEST_URL = 'https://weixin.sogou.com/'

### 豆瓣
TEST_URL = 'https://www.zhipin.com'



# 代理分数
MAX_SCORE = 100
MIN_SCORE = 0
INITIAL_SCORE = 10
TIMEOUT_SCORE = 0

# VALID_STATUS_CODES = [200, 302]
VALID_STATUS_CODES = [200]

# 代理池数量界限
POOL_UPPER_THRESHOLD = 50000

# 检查周期
TESTER_CYCLE = 20
# 获取周期
GETTER_CYCLE = 300


# API配置
API_HOST = '0.0.0.0'
API_PORT = 9527

# 开关
TESTER_ENABLED = True
GETTER_ENABLED = True
API_ENABLED = True

# 最大批测试量
BATCH_TEST_SIZE = 20
