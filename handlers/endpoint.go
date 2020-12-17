package handlers

import (
	"context"
	"encoding/json"
	. "history-graph-notes-server/model"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

//figure handler
func MakeFigureSingleProperEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(FigureSingleProperReq)
		v, err := svc.ReadFigureSingleProper(req.Name, req.Proper)
		if err != nil || v == nil {
			return FigureSingleProperResp{-1, err.Error(), ""}, nil
		}
		return FigureSingleProperResp{0, "", v.(string)}, nil
	}
}

func MakeFigureNodeProperEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(FigureNodeProperReq)
		v, err := svc.ReadFigureNodeProper(req.Name)
		if err != nil || v == nil {
			return FigureNodeProperResp{-1, err.Error(), nil}, nil
		}
		return FigureNodeProperResp{0, "", v.(*HistoryFigure)}, nil
	}
}

func decodeFigureSingleProperRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request FigureSingleProperReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeFigureNodeProperRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request FigureNodeProperReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//event handler
func MakeEventSingleProperEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(EventSingleProperReq)
		v, err := svc.ReadEventSingleProper(req.Name, req.Proper)
		if err != nil || v == nil {
			return EventSingleProperResp{-1, err.Error(), ""}, nil
		}
		return EventSingleProperResp{0, "", v.(string)}, nil
	}
}

func MakeEventNodeProperEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(EventNodeProperReq)
		v, err := svc.ReadEventNodeProper(req.Name)
		if err != nil || v == nil {
			return EventNodeProperResp{-1, err.Error(), nil}, nil
		}
		return EventNodeProperResp{0, "", v.(*HistoryEvent)}, nil
	}
}

func decodeEventSingleProperRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request EventSingleProperReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeEventNodeProperRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request EventNodeProperReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//relation handler
func MakeRelationLineEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(RelationLineReq)
		v, err := svc.ReadLineRelation(req.NameA, req.LableA, req.NameB, req.LableB)
		if err != nil || v == nil {
			return RelationLineResp{-1, err.Error(), ""}, nil
		}
		return RelationLineResp{0, "", v.(string)}, nil
	}
}

func decodeRelationLineRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request RelationLineReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func MakeRelationNodeEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(NodeRelationReq)
		v, err := svc.ReadRelationNode(req.Name, req.Lable)
		if err != nil || v == nil {
			return NodeRelationResp{-1, err.Error(), nil}, nil
		}
		return NodeRelationResp{0, "", v.(map[string]string)}, nil
	}
}

func decodeRelationNodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request NodeRelationReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func MakeRelationPathEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(RelationPathReq)
		v, err := svc.ReadAllRelationPath(req.NameA, req.LableA, req.NameB, req.LableB)
		if err != nil || v == nil {
			return RelationPathResp{-1, err.Error(), nil}, nil
		}
		return RelationPathResp{0, "", v.([]string)}, nil
	}
}

func decodeRelationPathRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request RelationPathReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
