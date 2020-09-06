package proxy

import "testing"

/*

 tip:
  使用场景:
     洋葱模型,一般使用在web中间件中
*/
func TestProxy(t *testing.T) {

    var chain handlerChain
    chain = append(chain,recordExecutionTime)
    chain = append(chain,PrintEnv1)
    var c context
    c.chain =chain
    c.Next()


}
