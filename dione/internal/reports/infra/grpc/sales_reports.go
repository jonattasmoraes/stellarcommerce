package grpc

import (
	"context"

	pb "github.com/jonattasmoraes/dione/internal/reports/infra/gen"
	"github.com/jonattasmoraes/dione/internal/reports/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type salesReportsGrpcServer struct {
	pb.UnimplementedReportsServiceServer
	reportService  *usecase.GetSalesById
}

func NewSalesReportsGrpcServer(getSalesById *usecase.GetSalesById) *salesReportsGrpcServer {
	return &salesReportsGrpcServer{
		reportService:  getSalesById,
	}
}

func (s *salesReportsGrpcServer) GetSalesByID(ctx context.Context, req *pb.GetSalesByIdRequest) (*pb.GetSalesByIdResponse, error) {
	report, err := s.reportService.Execute(req.SaleId,)
	if err != nil {
		return nil, err
	}

	if report == nil {
		return nil, status.Error(codes.NotFound, "sale not found")
	}

	response := &pb.GetSalesByIdResponse{
		SaleId:      report.SaleID,
		UserId:      report.UserID,
		Total:       int32(report.Total),
		PaymentType: report.PaymentType,
		Products:    []*pb.Product{},
	}

	for _, product := range report.Products {
		response.Products = append(response.Products, &pb.Product{
			ProductId: product.ProductID,
			Name:      product.Name,
			Unit:      product.Unit,
			Category:  product.Category,
			Quantity:  int32(product.Quantity),
			Price:     int32(product.Price),
		})
	}

	return response, nil
}
