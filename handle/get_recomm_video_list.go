package handle

import (
	"context"
	"ginServer/logger"
	pb "ginServer/proto/user_video_list"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type RecommClientServer struct {
	Engine *gin.Engine
	Client pb.UserVideServiceClient
}

func InitRecommClient(addr string) pb.UserVideServiceClient {
	logger.Log.Infof("get rpc addr:%v", addr)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second))
	defer cancel()
	conn, err := grpc.DialContext(ctx, addr,
		[]grpc.DialOption{
			grpc.WithInsecure(),
		}...)
	if err != nil {
		logger.Log.Errorf("init recomm client err:%v", err)
	}
	return pb.NewUserVideServiceClient(conn)
}

func (s *RecommClientServer) GetRecommVideoList(c *gin.Context) {
	uid := c.Request.FormValue("uid")
	num := c.Request.FormValue("num")
	req := &pb.SearchUserVideoReq{
		Uid:              cast.ToInt64(uid),
		RequestVideoSize: cast.ToInt64(num),
	}
	logger.Log.Infof("get recomm req:%v", req)
	resp, err := s.Client.SearchUserVideoList(context.Background(), req)
	if err != nil {
		logger.Log.Errorf("call recomm video fail:%v", err)
	}
	logger.Log.Debugf("get Rsp:%v", resp)
	c.JSON(http.StatusOK, resp)
}
