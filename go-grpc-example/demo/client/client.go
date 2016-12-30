package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "../proto"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

type calcClient struct {
	client pb.StackCalcClient
}

func (c *calcClient) createCalc(stackDepth uint32) {
	resp, err := c.client.CreateCalc(context.Background(), &pb.CreateCalcRequest{StackDepth: uint32(stackDepth)})
	if err != nil {
		fmt.Println("Error creating calc", err)
		return
	}
	fmt.Println("Calc ID:", resp.CalcId)
}

func (c *calcClient) destroyCalc(calcId string) {
	_, err := c.client.DestroyCalc(context.Background(), &pb.DestroyCalcRequest{CalcId: calcId})
	if err != nil {
		fmt.Println("Error destroying calc", err)
		return
	}
	fmt.Println("Destroyed calc", calcId)
}

func (c *calcClient) listCalcs(limit uint32) {
	resp, err := c.client.ListCalcs(context.Background(), &pb.ListCalcsRequest{Limit: limit})
	if err != nil {
		fmt.Println("Error listing calcs", err)
		return
	}
	for _, calcId := range resp.CalcIds {
		fmt.Println("Calc ID:", calcId)
	}
}

func parseToks(toks []string) ([]*pb.Statement_Token, error) {
	var ops []*pb.Statement_Token
	for _, tok := range toks {
		switch tok {
		case "+":
			ops = append(ops, &pb.Statement_Token{Op: pb.Statement_Token_ADD})
		case "-":
			ops = append(ops, &pb.Statement_Token{Op: pb.Statement_Token_SUBTRACT})
		case "*":
			ops = append(ops, &pb.Statement_Token{Op: pb.Statement_Token_MULTIPLY})
		case "/":
			ops = append(ops, &pb.Statement_Token{Op: pb.Statement_Token_DIVIDE})
		case ".":
			ops = append(ops, &pb.Statement_Token{Op: pb.Statement_Token_DROP})
		default:
			val, err := strconv.ParseInt(tok, 10, 64)
			if err != nil {
				fmt.Println("Error parsing argument '", tok, "'", err)
				return nil, err
			}
			ops = append(ops, &pb.Statement_Token{Op: pb.Statement_Token_VAL, Val: val})
		}
	}
	return ops, nil
}

func (c *calcClient) evaluateStatement(calcId string, toks []string) {
	ops, err := parseToks(toks)
	if err != nil {
		return
	}
	req := pb.EvaluateStatementRequest{CalcId: calcId, Statement: &pb.Statement{Tokens: ops}}
	resp, err := c.client.EvaluateStatement(context.Background(), &req)
	if err != nil {
		fmt.Println("Error evaluating statement:", err)
		return
	}
	fmt.Println("Stack Depth:", resp.State.StackDepth)
	fmt.Println("Vals:", resp.State.Vals)
	fmt.Println("Error:", resp.Err)
}

func (c *calcClient) getState(calcId string) {
	resp, err := c.client.GetState(context.Background(), &pb.GetStateRequest{CalcId: calcId})
	if err != nil {
		fmt.Println("Error getting state:", err)
		return
	}
	fmt.Println("Stack Depth:", resp.State.StackDepth)
	fmt.Println("Vals:", resp.State.Vals)
}

func (c *calcClient) interact(calcId string) {
	stream, err := c.client.Interact(context.Background())
	if err != nil {
		fmt.Println("Error interacting with calc", err)
		return
	}
	defer fmt.Println()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("grpcdemo [%s]> ", calcId)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, " ")
		if strings.ToLower(parts[0]) == "stop" {
			return
		}
		ops, err := parseToks(parts)
		if err != nil {
			return
		}
		req := &pb.EvaluateStatementRequest{CalcId: calcId, Statement: &pb.Statement{Tokens: ops}}
		err = stream.Send(req)
		if err != nil {
			fmt.Println("Error interacting", err)
			return
		}
		resp, err := stream.Recv()
		if err != nil {
			fmt.Println("Error receiving interaction response", err)
			return
		}
		fmt.Println("Stack Depth:", resp.State.StackDepth)
		fmt.Println("Vals:", resp.State.Vals)
		fmt.Println("Error:", resp.Err)
		fmt.Printf("grpcdemo [%s]> ", calcId)
	}
}

func runRepl(client pb.StackCalcClient) {
	calc := new(calcClient)
	calc.client = client
	fmt.Print("grpcdemo> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, " ")
		switch command := strings.ToLower(parts[0]); command {
		case "create":
			stackDepth, err := strconv.ParseUint(parts[1], 10, 32)
			if err != nil {
				fmt.Println("Error parsing argument", err)
				continue
			}
			calc.createCalc(uint32(stackDepth))
		case "destroy":
			calcId := parts[1]
			calc.destroyCalc(calcId)
		case "list":
			var limit uint32
			if len(parts) == 2 {
				parsedLimit, err := strconv.ParseUint(parts[1], 10, 32)
				if err != nil {
					fmt.Println("Error parsing argument", err)
					continue
				}
				limit = uint32(parsedLimit)
			} else {
				limit = 10
			}
			calc.listCalcs(limit)
		case "eval":
			calc.evaluateStatement(parts[1], parts[2:])
		case "get":
			calc.getState(parts[1])
		case "interact":
			calc.interact(parts[1])
		default:
			fmt.Println("Invalid command", command)
		}
		fmt.Print("grpcdemo> ")
	}
	fmt.Println()
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)

	}
	defer conn.Close()
	client := pb.NewStackCalcClient(conn)
	runRepl(client)
}
