package sapphire.userApp.sapphireObject.Algo;
import java.util.*;
import com.google.protobuf.InvalidProtocolBufferException;
import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.util.logging.Logger;
import algo_proto.AlgoProto;
import sapphire.userApp.sapphireObject.Algo
;class Algo_Stub extends Algo {
  Object sObj;
  public Algo_Stub(){
  super();
  }

  byte[] Fibbonaci (byte[] reqData ) {
    AlgoProto.FibbonaciRequest request;
    try{
    request = AlgoProto.FibbonaciRequest.parseFrom(reqData); 


    }catch (InvalidProtocolBufferException e) {
     System.out.println("Fibbonaci parameter deserialization failed: " + e.getMessage());
     }
    String  Ret =super.Fibbonaci(request.getNum());

    AlgoProto.FibbonaciReply.Builder retmsg = AlgoProto.FibbonaciReply.newBuilder();

    retmsg.setRet(Ret);
    AlgoProto.FibbonaciReply msgbuffer = retmsg.build();

    return msgbuffer.toByteArray();
  }

  byte[] Search (byte[] reqData ) {
    AlgoProto.SearchRequest request;
    try{
    request = AlgoProto.SearchRequest.parseFrom(reqData); 


    }catch (InvalidProtocolBufferException e) {
     System.out.println("Search parameter deserialization failed: " + e.getMessage());
     }
    int  Ret =super.Search(request.getNumList(), request.getKey());

    AlgoProto.SearchReply.Builder retmsg = AlgoProto.SearchReply.newBuilder();

    retmsg.setRet(Ret);
    AlgoProto.SearchReply msgbuffer = retmsg.build();

    return msgbuffer.toByteArray();
  }

  byte[] Sort (byte[] reqData ) {
    AlgoProto.SortRequest request;
    try{
    request = AlgoProto.SortRequest.parseFrom(reqData); 


    }catch (InvalidProtocolBufferException e) {
     System.out.println("Sort parameter deserialization failed: " + e.getMessage());
     }
    List<Integer>  Arr =super.Sort(request.getArrList());

    AlgoProto.SortReply.Builder retmsg = AlgoProto.SortReply.newBuilder();

    retmsg.addAllArr(Arr);
    AlgoProto.SortReply msgbuffer = retmsg.build();

    return msgbuffer.toByteArray();
  }

}
