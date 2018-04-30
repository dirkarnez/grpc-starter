// package: mentor
// file: mentor.proto

import * as mentor_pb from "./mentor_pb";
export class MentorService {
  static serviceName = "mentor.MentorService";
}
export namespace MentorService {
  export class Login {
    static readonly methodName = "Login";
    static readonly service = MentorService;
    static readonly requestStream = false;
    static readonly responseStream = false;
    static readonly requestType = mentor_pb.LoginRequest;
    static readonly responseType = mentor_pb.LoginResponse;
  }
  export class ListLessons {
    static readonly methodName = "ListLessons";
    static readonly service = MentorService;
    static readonly requestStream = false;
    static readonly responseStream = true;
    static readonly requestType = mentor_pb.ListLessonsRequest;
    static readonly responseType = mentor_pb.ListLessonsResponse;
  }
}
