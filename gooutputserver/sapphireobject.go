package test

import "github.com/golang/protobuf/proto"
import "fmt"

type Testobj struct {
	data int32
	Name string
}
/**
 *  This is a test comment on func1
 **/
func (o Testobj) Func1(Name string, id int32, floatingvalue float64) ([]string, map[int32]string) {

	msg := &Msg1{
		Name: Name,
		Id: id,
		Floatingvalue: floatingvalue,
	}
	fmt.Println("Func1 Invocked",msg)

	var abc Msg2
	abc.Liststring = []string {"abc","cde"}
	abc.Idtostringmap = make(map[int32]string)
	abc.Idtostringmap[1] = "test"


	return abc.Liststring, abc.Idtostringmap
}

/**
 *  This is a multiline test
 *  Comment on func2
 *
 **/
func (o Testobj) Func2(Liststring []string, Idtostringmap map[int32]string) (*Msg2) {
	fmt.Println("Func2***** Invocked")
	msg := &Msg2{
		Liststring: Liststring,
		Idtostringmap: Idtostringmap,
	}
	msgbuffer, _ := proto.Marshal(msg)

	_ = msgbuffer //suppresses warning
	retmsgbuffer := make([]byte, 0)

	retmsg := &Msg3{}
	_ = proto.Unmarshal(retmsgbuffer, retmsg)
	return retmsg.Nested
}

/**
 *  This is a different test
 * comment on func3
 **/
func (o Testobj) Func3(Nested *Msg2) (string, int32, float64) {
	msg := &Msg3{
		Nested: Nested,
	}
	msgbuffer, _ := proto.Marshal(msg)

	_ = msgbuffer //suppresses warning
	retmsgbuffer := make([]byte, 0)

	retmsg := &Msg1{}
	_ = proto.Unmarshal(retmsgbuffer, retmsg)
	return retmsg.Name, retmsg.Id, retmsg.Floatingvalue
}

/**
 *  This is a different test
 * comment on func4
 **/
func (o Testobj) Func4(params []byte) ([]byte) {
	msg := &Msg4{
		Params: params,
	}
	msgbuffer, _ := proto.Marshal(msg)

	_ = msgbuffer //suppresses warning
	retmsgbuffer := make([]byte, 0)

	retmsg := &Msg4{}
	_ = proto.Unmarshal(retmsgbuffer, retmsg)
	return retmsg.Params
}