package handle

import (
	"context"
	"ginServer/logger"
	pb "ginServer/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type VideoReportClientServer struct {
	Engine *gin.Engine
	Client pb.VideoReportClient
}

func InitClient(addr string) pb.VideoReportClient {
	logger.Log.Infof("get addr:%s", addr)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second))
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr,
		[]grpc.DialOption{
			grpc.WithInsecure(),
		}...)
	if err != nil {
		logger.Log.Errorf("error, conn:%v", err)
	}
	return pb.NewVideoReportClient(conn)
}

// 调用rpc 接口
func (s *VideoReportClientServer) GetVideoReportData(c *gin.Context) {
	start := c.Request.FormValue("starttime")
	end := c.Request.FormValue("endtime")
	req := &pb.DailyReportReq{
		StartTime: start,
		EndTime:   end,
	}
	logger.Log.Infof("get req:%+v", req)
	rsp, err := s.Client.QueryDailyReport(context.Background(), req)
	if err != nil {
		logger.Log.Errorf("err rpc: %v", err)
	}
	logger.Log.Debugf("rsp:%+v", rsp)
	c.JSON(http.StatusOK, rsp)

}
