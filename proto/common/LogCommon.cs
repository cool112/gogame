// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: log_common.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Common {

  /// <summary>Holder for reflection information generated from log_common.proto</summary>
  public static partial class LogCommonReflection {

    #region Descriptor
    /// <summary>File descriptor for log_common.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static LogCommonReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChBsb2dfY29tbW9uLnByb3RvEgZjb21tb24i4QEKCUxvZ0NvbW1vbhILCgNV",
            "aWQYASABKAkSEAoIRGV2aWNlSWQYAiABKAkSJwoEVHlwZRgDIAEoDjIZLmNv",
            "bW1vbi5Mb2dDb21tb24uTG9nVHlwZRILCgNNc2cYBCABKAkSDAoEVGltZRgF",
            "IAEoAxILCgNBcHAYBiABKAkiZAoHTG9nVHlwZRINCglTVEFSVF9BUFAQABIQ",
            "CgxVUERBVEVfU1RBUlQQARIOCgpVUERBVEVfRU5EEAISCQoFTE9HSU4QAxIN",
            "CglMT0dJTl9TVUMQBBIOCgpMT0dJTl9GQUlMEAViBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Common.LogCommon), global::Common.LogCommon.Parser, new[]{ "Uid", "DeviceId", "Type", "Msg", "Time", "App" }, null, new[]{ typeof(global::Common.LogCommon.Types.LogType) }, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class LogCommon : pb::IMessage<LogCommon> {
    private static readonly pb::MessageParser<LogCommon> _parser = new pb::MessageParser<LogCommon>(() => new LogCommon());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<LogCommon> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Common.LogCommonReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public LogCommon() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public LogCommon(LogCommon other) : this() {
      uid_ = other.uid_;
      deviceId_ = other.deviceId_;
      type_ = other.type_;
      msg_ = other.msg_;
      time_ = other.time_;
      app_ = other.app_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public LogCommon Clone() {
      return new LogCommon(this);
    }

    /// <summary>Field number for the "Uid" field.</summary>
    public const int UidFieldNumber = 1;
    private string uid_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Uid {
      get { return uid_; }
      set {
        uid_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "DeviceId" field.</summary>
    public const int DeviceIdFieldNumber = 2;
    private string deviceId_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string DeviceId {
      get { return deviceId_; }
      set {
        deviceId_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "Type" field.</summary>
    public const int TypeFieldNumber = 3;
    private global::Common.LogCommon.Types.LogType type_ = 0;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Common.LogCommon.Types.LogType Type {
      get { return type_; }
      set {
        type_ = value;
      }
    }

    /// <summary>Field number for the "Msg" field.</summary>
    public const int MsgFieldNumber = 4;
    private string msg_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Msg {
      get { return msg_; }
      set {
        msg_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "Time" field.</summary>
    public const int TimeFieldNumber = 5;
    private long time_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public long Time {
      get { return time_; }
      set {
        time_ = value;
      }
    }

    /// <summary>Field number for the "App" field.</summary>
    public const int AppFieldNumber = 6;
    private string app_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string App {
      get { return app_; }
      set {
        app_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as LogCommon);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(LogCommon other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Uid != other.Uid) return false;
      if (DeviceId != other.DeviceId) return false;
      if (Type != other.Type) return false;
      if (Msg != other.Msg) return false;
      if (Time != other.Time) return false;
      if (App != other.App) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Uid.Length != 0) hash ^= Uid.GetHashCode();
      if (DeviceId.Length != 0) hash ^= DeviceId.GetHashCode();
      if (Type != 0) hash ^= Type.GetHashCode();
      if (Msg.Length != 0) hash ^= Msg.GetHashCode();
      if (Time != 0L) hash ^= Time.GetHashCode();
      if (App.Length != 0) hash ^= App.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Uid.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Uid);
      }
      if (DeviceId.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(DeviceId);
      }
      if (Type != 0) {
        output.WriteRawTag(24);
        output.WriteEnum((int) Type);
      }
      if (Msg.Length != 0) {
        output.WriteRawTag(34);
        output.WriteString(Msg);
      }
      if (Time != 0L) {
        output.WriteRawTag(40);
        output.WriteInt64(Time);
      }
      if (App.Length != 0) {
        output.WriteRawTag(50);
        output.WriteString(App);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Uid.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Uid);
      }
      if (DeviceId.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(DeviceId);
      }
      if (Type != 0) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Type);
      }
      if (Msg.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Msg);
      }
      if (Time != 0L) {
        size += 1 + pb::CodedOutputStream.ComputeInt64Size(Time);
      }
      if (App.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(App);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(LogCommon other) {
      if (other == null) {
        return;
      }
      if (other.Uid.Length != 0) {
        Uid = other.Uid;
      }
      if (other.DeviceId.Length != 0) {
        DeviceId = other.DeviceId;
      }
      if (other.Type != 0) {
        Type = other.Type;
      }
      if (other.Msg.Length != 0) {
        Msg = other.Msg;
      }
      if (other.Time != 0L) {
        Time = other.Time;
      }
      if (other.App.Length != 0) {
        App = other.App;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Uid = input.ReadString();
            break;
          }
          case 18: {
            DeviceId = input.ReadString();
            break;
          }
          case 24: {
            type_ = (global::Common.LogCommon.Types.LogType) input.ReadEnum();
            break;
          }
          case 34: {
            Msg = input.ReadString();
            break;
          }
          case 40: {
            Time = input.ReadInt64();
            break;
          }
          case 50: {
            App = input.ReadString();
            break;
          }
        }
      }
    }

    #region Nested types
    /// <summary>Container for nested types declared in the LogCommon message type.</summary>
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static partial class Types {
      public enum LogType {
        [pbr::OriginalName("START_APP")] StartApp = 0,
        [pbr::OriginalName("UPDATE_START")] UpdateStart = 1,
        [pbr::OriginalName("UPDATE_END")] UpdateEnd = 2,
        [pbr::OriginalName("LOGIN")] Login = 3,
        [pbr::OriginalName("LOGIN_SUC")] LoginSuc = 4,
        [pbr::OriginalName("LOGIN_FAIL")] LoginFail = 5,
      }

    }
    #endregion

  }

  #endregion

}

#endregion Designer generated code
