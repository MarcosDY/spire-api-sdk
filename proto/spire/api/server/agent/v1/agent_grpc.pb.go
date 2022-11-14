// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package agentv1

import (
	context "context"
	types "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	// Count agents.
	//
	// The caller must be local or present an admin X509-SVID.
	CountAgents(ctx context.Context, in *CountAgentsRequest, opts ...grpc.CallOption) (*CountAgentsResponse, error)
	// Lists agents.
	//
	// The caller must be local or present an admin X509-SVID.
	ListAgents(ctx context.Context, in *ListAgentsRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error)
	// Gets an agent.
	//
	// The caller must be local or present an admin X509-SVID.
	GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*types.Agent, error)
	// Deletes an agent. The agent can come back into the trust domain through
	// the Issuer AttestAgent RPC.
	//
	// The caller must be local or present an admin X509-SVID.
	DeleteAgent(ctx context.Context, in *DeleteAgentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Bans an agent. This evicts the agent and prevents it from rejoining the
	// trust domain through attestation until the ban is lifted via a call to
	// DeleteAgent.
	//
	// The caller must be local or present an admin X509-SVID.
	BanAgent(ctx context.Context, in *BanAgentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Attests the agent via node attestation, using a bidirectional stream to
	// faciliate attestation methods that require challenge/response.
	//
	// The caller is not authenticated.
	AttestAgent(ctx context.Context, opts ...grpc.CallOption) (Agent_AttestAgentClient, error)
	// Renews the agent and returns a new X509-SVID. The new SVID is not enabled
	// on the server side until its first use.
	//
	// The caller must present an active agent X509-SVID, i.e. the X509-SVID
	// returned by the AttestAgent or the most recent RenewAgent call.
	RenewAgent(ctx context.Context, in *RenewAgentRequest, opts ...grpc.CallOption) (*RenewAgentResponse, error)
	// Creates an agent join token. The token can be used with `join_token`
	// attestation to join the trust domain.
	//
	// The caller must be local or present an admin X509-SVID.
	CreateJoinToken(ctx context.Context, in *CreateJoinTokenRequest, opts ...grpc.CallOption) (*types.JoinToken, error)
	// Pushes Agent status, relevant information is returned to agent to keep it
	// updated on server changes.
	//
	// The caller must present an active agent X509-SVID, i.e. the X509-SVID
	// returned by the AttestAgent or the most recent RenewAgent call.
	PushStatus(ctx context.Context, in *PushStatusRequest, opts ...grpc.CallOption) (*PushStatusResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) CountAgents(ctx context.Context, in *CountAgentsRequest, opts ...grpc.CallOption) (*CountAgentsResponse, error) {
	out := new(CountAgentsResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.agent.v1.Agent/CountAgents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ListAgents(ctx context.Context, in *ListAgentsRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error) {
	out := new(ListAgentsResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.agent.v1.Agent/ListAgents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*types.Agent, error) {
	out := new(types.Agent)
	err := c.cc.Invoke(ctx, "/spire.api.server.agent.v1.Agent/GetAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeleteAgent(ctx context.Context, in *DeleteAgentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/spire.api.server.agent.v1.Agent/DeleteAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) BanAgent(ctx context.Context, in *BanAgentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/spire.api.server.agent.v1.Agent/BanAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) AttestAgent(ctx context.Context, opts ...grpc.CallOption) (Agent_AttestAgentClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Agent_serviceDesc.Streams[0], "/spire.api.server.agent.v1.Agent/AttestAgent", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentAttestAgentClient{stream}
	return x, nil
}

type Agent_AttestAgentClient interface {
	Send(*AttestAgentRequest) error
	Recv() (*AttestAgentResponse, error)
	grpc.ClientStream
}

type agentAttestAgentClient struct {
	grpc.ClientStream
}

func (x *agentAttestAgentClient) Send(m *AttestAgentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentAttestAgentClient) Recv() (*AttestAgentResponse, error) {
	m := new(AttestAgentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentClient) RenewAgent(ctx context.Context, in *RenewAgentRequest, opts ...grpc.CallOption) (*RenewAgentResponse, error) {
	out := new(RenewAgentResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.agent.v1.Agent/RenewAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CreateJoinToken(ctx context.Context, in *CreateJoinTokenRequest, opts ...grpc.CallOption) (*types.JoinToken, error) {
	out := new(types.JoinToken)
	err := c.cc.Invoke(ctx, "/spire.api.server.agent.v1.Agent/CreateJoinToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) PushStatus(ctx context.Context, in *PushStatusRequest, opts ...grpc.CallOption) (*PushStatusResponse, error) {
	out := new(PushStatusResponse)
	err := c.cc.Invoke(ctx, "/spire.api.server.agent.v1.Agent/PushStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	// Count agents.
	//
	// The caller must be local or present an admin X509-SVID.
	CountAgents(context.Context, *CountAgentsRequest) (*CountAgentsResponse, error)
	// Lists agents.
	//
	// The caller must be local or present an admin X509-SVID.
	ListAgents(context.Context, *ListAgentsRequest) (*ListAgentsResponse, error)
	// Gets an agent.
	//
	// The caller must be local or present an admin X509-SVID.
	GetAgent(context.Context, *GetAgentRequest) (*types.Agent, error)
	// Deletes an agent. The agent can come back into the trust domain through
	// the Issuer AttestAgent RPC.
	//
	// The caller must be local or present an admin X509-SVID.
	DeleteAgent(context.Context, *DeleteAgentRequest) (*emptypb.Empty, error)
	// Bans an agent. This evicts the agent and prevents it from rejoining the
	// trust domain through attestation until the ban is lifted via a call to
	// DeleteAgent.
	//
	// The caller must be local or present an admin X509-SVID.
	BanAgent(context.Context, *BanAgentRequest) (*emptypb.Empty, error)
	// Attests the agent via node attestation, using a bidirectional stream to
	// faciliate attestation methods that require challenge/response.
	//
	// The caller is not authenticated.
	AttestAgent(Agent_AttestAgentServer) error
	// Renews the agent and returns a new X509-SVID. The new SVID is not enabled
	// on the server side until its first use.
	//
	// The caller must present an active agent X509-SVID, i.e. the X509-SVID
	// returned by the AttestAgent or the most recent RenewAgent call.
	RenewAgent(context.Context, *RenewAgentRequest) (*RenewAgentResponse, error)
	// Creates an agent join token. The token can be used with `join_token`
	// attestation to join the trust domain.
	//
	// The caller must be local or present an admin X509-SVID.
	CreateJoinToken(context.Context, *CreateJoinTokenRequest) (*types.JoinToken, error)
	// Pushes Agent status, relevant information is returned to agent to keep it
	// updated on server changes.
	//
	// The caller must present an active agent X509-SVID, i.e. the X509-SVID
	// returned by the AttestAgent or the most recent RenewAgent call.
	PushStatus(context.Context, *PushStatusRequest) (*PushStatusResponse, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) CountAgents(context.Context, *CountAgentsRequest) (*CountAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAgents not implemented")
}
func (UnimplementedAgentServer) ListAgents(context.Context, *ListAgentsRequest) (*ListAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgents not implemented")
}
func (UnimplementedAgentServer) GetAgent(context.Context, *GetAgentRequest) (*types.Agent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgent not implemented")
}
func (UnimplementedAgentServer) DeleteAgent(context.Context, *DeleteAgentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgent not implemented")
}
func (UnimplementedAgentServer) BanAgent(context.Context, *BanAgentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanAgent not implemented")
}
func (UnimplementedAgentServer) AttestAgent(Agent_AttestAgentServer) error {
	return status.Errorf(codes.Unimplemented, "method AttestAgent not implemented")
}
func (UnimplementedAgentServer) RenewAgent(context.Context, *RenewAgentRequest) (*RenewAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewAgent not implemented")
}
func (UnimplementedAgentServer) CreateJoinToken(context.Context, *CreateJoinTokenRequest) (*types.JoinToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJoinToken not implemented")
}
func (UnimplementedAgentServer) PushStatus(context.Context, *PushStatusRequest) (*PushStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushStatus not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_CountAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CountAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.agent.v1.Agent/CountAgents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CountAgents(ctx, req.(*CountAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ListAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ListAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.agent.v1.Agent/ListAgents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ListAgents(ctx, req.(*ListAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.agent.v1.Agent/GetAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetAgent(ctx, req.(*GetAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeleteAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeleteAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.agent.v1.Agent/DeleteAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeleteAgent(ctx, req.(*DeleteAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_BanAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).BanAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.agent.v1.Agent/BanAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).BanAgent(ctx, req.(*BanAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_AttestAgent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).AttestAgent(&agentAttestAgentServer{stream})
}

type Agent_AttestAgentServer interface {
	Send(*AttestAgentResponse) error
	Recv() (*AttestAgentRequest, error)
	grpc.ServerStream
}

type agentAttestAgentServer struct {
	grpc.ServerStream
}

func (x *agentAttestAgentServer) Send(m *AttestAgentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentAttestAgentServer) Recv() (*AttestAgentRequest, error) {
	m := new(AttestAgentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Agent_RenewAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RenewAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.agent.v1.Agent/RenewAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RenewAgent(ctx, req.(*RenewAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CreateJoinToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJoinTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateJoinToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.agent.v1.Agent/CreateJoinToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateJoinToken(ctx, req.(*CreateJoinTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_PushStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).PushStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.server.agent.v1.Agent/PushStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).PushStatus(ctx, req.(*PushStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.server.agent.v1.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountAgents",
			Handler:    _Agent_CountAgents_Handler,
		},
		{
			MethodName: "ListAgents",
			Handler:    _Agent_ListAgents_Handler,
		},
		{
			MethodName: "GetAgent",
			Handler:    _Agent_GetAgent_Handler,
		},
		{
			MethodName: "DeleteAgent",
			Handler:    _Agent_DeleteAgent_Handler,
		},
		{
			MethodName: "BanAgent",
			Handler:    _Agent_BanAgent_Handler,
		},
		{
			MethodName: "RenewAgent",
			Handler:    _Agent_RenewAgent_Handler,
		},
		{
			MethodName: "CreateJoinToken",
			Handler:    _Agent_CreateJoinToken_Handler,
		},
		{
			MethodName: "PushStatus",
			Handler:    _Agent_PushStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AttestAgent",
			Handler:       _Agent_AttestAgent_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "spire/api/server/agent/v1/agent.proto",
}
