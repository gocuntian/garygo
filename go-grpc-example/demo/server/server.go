package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"strconv"
    "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"

	pb "../proto"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type Stack struct {
	stack []int64
}

func (s Stack) Empty() bool {
	return len(s.stack) == 0
}

func (s Stack) Len() uint32 {
	return uint32(len(s.stack))
}

func (s *Stack) Push(i int64) {
	s.stack = append(s.stack, i)
}

func (s *Stack) Pop() int64 {
	d := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return d
}

type calc struct {
	stack      Stack
	stackDepth uint32
}

type stackCalcServer struct {
	calcs    map[string]*calc
	nextCalc uint64
}

func newCalc(stackDepth uint32) *calc {
	c := new(calc)
	c.stackDepth = stackDepth
	return c
}

func (c *calc) PushVal(val int64) pb.CalculationError {
	if c.stack.Len() >= c.stackDepth {
		return pb.CalculationError_STACK_OVERFLOW
	}
	c.stack.Push(val)
	return pb.CalculationError_NO_ERROR
}

func (c *calc) Add() pb.CalculationError {
	if c.stack.Len() < 2 {
		return pb.CalculationError_STACK_UNDERFLOW
	}
	a := c.stack.Pop()
	b := c.stack.Pop()
	c.stack.Push(a + b)
	return pb.CalculationError_NO_ERROR
}

func (c *calc) Subtract() pb.CalculationError {
	if c.stack.Len() < 2 {
		return pb.CalculationError_STACK_UNDERFLOW
	}
	a := c.stack.Pop()
	b := c.stack.Pop()
	c.stack.Push(b - a)
	return pb.CalculationError_NO_ERROR
}

func (c *calc) Multiply() pb.CalculationError {
	if c.stack.Len() < 2 {
		return pb.CalculationError_STACK_UNDERFLOW
	}
	a := c.stack.Pop()
	b := c.stack.Pop()
	c.stack.Push(a * b)
	return pb.CalculationError_NO_ERROR
}

func (c *calc) Divide() pb.CalculationError {
	if c.stack.Len() < 2 {
		return pb.CalculationError_STACK_UNDERFLOW
	}
	a := c.stack.Pop()
	b := c.stack.Pop()
	if a == 0 {
		c.stack.Push(b)
		c.stack.Push(a)
		return pb.CalculationError_DIVIDE_BY_ZERO
	}
	c.stack.Push(b / a)
	return pb.CalculationError_NO_ERROR
}

func (c *calc) Drop() pb.CalculationError {
	if c.stack.Len() < 1 {
		return pb.CalculationError_STACK_UNDERFLOW
	}
	c.stack.Pop()
	return pb.CalculationError_NO_ERROR
}

func (s *stackCalcServer) CreateCalc(ctx context.Context,create *pb.CreateCalcRequest) (*pb.CreateCalcResponse, error) {
	fmt.Println("Received CreateCalc", create)
	calcIndex := strconv.FormatUint(s.nextCalc, 36)
	s.nextCalc += 1
	s.calcs[calcIndex] = newCalc(create.StackDepth)
	return &pb.CreateCalcResponse{CalcId: calcIndex}, nil
}

func (s *stackCalcServer) DestroyCalc(ctx context.Context,destroy *pb.DestroyCalcRequest) (*pb.DestroyCalcResponse, error) {
	fmt.Println("Received DestroyCalc", destroy)
	delete(s.calcs, destroy.CalcId)
	return &pb.DestroyCalcResponse{}, nil
}

func (s *stackCalcServer) ListCalcs(ctx context.Context,list *pb.ListCalcsRequest) (*pb.ListCalcsResponse, error) {
	fmt.Println("Received ListCalcs", list)
	var calcs []string
	for k, _ := range s.calcs {
		calcs = append(calcs, k)
	}
	return &pb.ListCalcsResponse{CalcIds: calcs}, nil
}

func (s *stackCalcServer) eval(eval *pb.EvaluateStatementRequest) (*pb.EvaluateStatementResponse, error) {
	calc, ok := s.calcs[eval.CalcId]
	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "Calculator %s not found", eval.CalcId)
	}

	state := pb.State{StackDepth: calc.stackDepth}
	resp := pb.EvaluateStatementResponse{}
	for _, token := range eval.Statement.Tokens {
		switch token.Op {
		case pb.Statement_Token_VAL:
			resp.Err = calc.PushVal(token.Val)
		case pb.Statement_Token_ADD:
			resp.Err = calc.Add()
		case pb.Statement_Token_SUBTRACT:
			resp.Err = calc.Subtract()
		case pb.Statement_Token_MULTIPLY:
			resp.Err = calc.Multiply()
		case pb.Statement_Token_DIVIDE:
			resp.Err = calc.Divide()
		case pb.Statement_Token_DROP:
			resp.Err = calc.Drop()
		}
		if resp.Err != pb.CalculationError_NO_ERROR {
			break
		}
	}
	state.Vals = calc.stack.stack
	resp.State = &state
	return &resp, nil
}

func (s *stackCalcServer) EvaluateStatement(ctx context.Context,eval *pb.EvaluateStatementRequest) (*pb.EvaluateStatementResponse, error) {
	fmt.Println("Received EvaluateStatement", eval)
	resp, err := s.eval(eval)
	return resp, err
}

func (s *stackCalcServer) Interact(stream pb.StackCalc_InteractServer) error {
	fmt.Println("Started bidirectional streaming")
	for {
		req, err := stream.Recv()
		fmt.Println("Stream received", req)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		resp, err := s.eval(req)
		if err != nil {
			return err
		}
		err = stream.Send(resp)
		fmt.Println("Stream sent", resp)
		if err != nil {
			return err
		}
	}
}

func (s *stackCalcServer) GetState(ctx context.Context,get *pb.GetStateRequest) (*pb.GetStateResponse, error) {
	fmt.Println("Received GetState", get)
	calc, ok := s.calcs[get.CalcId]
	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "Calculator %s not found", get.CalcId)
	}
	return &pb.GetStateResponse{State: &pb.State{StackDepth: calc.stackDepth, Vals: calc.stack.stack}}, nil
}

func newServer() *stackCalcServer {
	s := new(stackCalcServer)
	s.calcs = make(map[string]*calc)
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)

	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterStackCalcServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
