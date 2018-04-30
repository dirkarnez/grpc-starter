// package: mentor
// file: mentor.proto

import * as jspb from "google-protobuf";

export class LoginRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginRequest;
  static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class LoginResponse extends jspb.Message {
  getIssuccessful(): boolean;
  setIssuccessful(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoginResponse.AsObject;
  static toObject(includeInstance: boolean, msg: LoginResponse): LoginResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoginResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoginResponse;
  static deserializeBinaryFromReader(message: LoginResponse, reader: jspb.BinaryReader): LoginResponse;
}

export namespace LoginResponse {
  export type AsObject = {
    issuccessful: boolean,
  }
}

export class Lesson extends jspb.Message {
  getId(): number;
  setId(value: number): void;

  getScore(): number;
  setScore(value: number): void;

  getTitle(): string;
  setTitle(value: string): void;

  getBy(): string;
  setBy(value: string): void;

  getTime(): number;
  setTime(value: number): void;

  getUrl(): string;
  setUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Lesson.AsObject;
  static toObject(includeInstance: boolean, msg: Lesson): Lesson.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Lesson, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Lesson;
  static deserializeBinaryFromReader(message: Lesson, reader: jspb.BinaryReader): Lesson;
}

export namespace Lesson {
  export type AsObject = {
    id: number,
    score: number,
    title: string,
    by: string,
    time: number,
    url: string,
  }
}

export class ListLessonsResponse extends jspb.Message {
  hasLesson(): boolean;
  clearLesson(): void;
  getLesson(): Lesson | undefined;
  setLesson(value?: Lesson): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListLessonsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListLessonsResponse): ListLessonsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListLessonsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListLessonsResponse;
  static deserializeBinaryFromReader(message: ListLessonsResponse, reader: jspb.BinaryReader): ListLessonsResponse;
}

export namespace ListLessonsResponse {
  export type AsObject = {
    lesson?: Lesson.AsObject,
  }
}

export class ListLessonsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListLessonsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListLessonsRequest): ListLessonsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListLessonsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListLessonsRequest;
  static deserializeBinaryFromReader(message: ListLessonsRequest, reader: jspb.BinaryReader): ListLessonsRequest;
}

export namespace ListLessonsRequest {
  export type AsObject = {
  }
}

