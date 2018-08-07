package test

import "github.com/golang/protobuf/proto"

/**
 *  This is a test comment on testobj 
 *  
 **/
type TestobjStub struct {
  Oid uint64 
  so Testobj 
}

func  CreateSapphireStubObject() (outParam  interface {}) {
  var stubObj TestobjStub
  return &stubObj
}

func  GetSapphireObjectStatus() (outParam  bool) {

  return true
}

/**
 *  This is a test comment on func1  
 **/
func (Obj TestobjStub) Func1(inParams  [] byte) (outParams  [] byte) {
  var err error
  InPutStruct := &Msg1{}
  err = proto.Unmarshal(inParams, InPutStruct)
  if err != nil  {
    return
  }
  OutPutStruct := &Msg2{}
  OutPutStruct.Liststring, OutPutStruct.Idtostringmap =  Obj.so.Func1(InPutStruct.Name , InPutStruct.Id , InPutStruct.Floatingvalue )
  outParams, err = proto.Marshal(OutPutStruct)

  return
}

/**
 *  This is a multiline test 
 *  Comment on func2 
 *  
 **/
func (Obj TestobjStub) Func2(inParams  [] byte) (outParams  [] byte) {
  var err error
  InPutStruct := &Msg2{}
  err = proto.Unmarshal(inParams, InPutStruct)
  if err != nil  {
    return
  }
  OutPutStruct := &Msg3{}
  OutPutStruct.Nested =  Obj.so.Func2(InPutStruct.Liststring , InPutStruct.Idtostringmap )
  outParams, err = proto.Marshal(OutPutStruct)

  return
}

/**
 *  This is a different test 
 * comment on func3  
 **/
func (Obj TestobjStub) Func3(inParams  [] byte) (outParams  [] byte) {
  var err error
  InPutStruct := &Msg3{}
  err = proto.Unmarshal(inParams, InPutStruct)
  if err != nil  {
    return
  }
  OutPutStruct := &Msg1{}
  OutPutStruct.Name, OutPutStruct.Id, OutPutStruct.Floatingvalue =  Obj.so.Func3(InPutStruct.Nested )
  outParams, err = proto.Marshal(OutPutStruct)

  return
}

/**
 *  This is a different test 
 * comment on func4  
 **/
func (Obj TestobjStub) Func4(inParams  [] byte) (outParams  [] byte) {
  var err error
  InPutStruct := &Msg4{}
  err = proto.Unmarshal(inParams, InPutStruct)
  if err != nil  {
    return
  }
  OutPutStruct := &Msg4{}
  OutPutStruct.Params =  Obj.so.Func4(InPutStruct.Params )
  outParams, err = proto.Marshal(OutPutStruct)

  return
}

