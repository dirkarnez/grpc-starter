import { Action } from 'redux';
import { ListLessonsRequest, ListLessonsResponse, Lesson } from '../proto/mentor_pb';
import { GrpcAction, grpcRequest } from '../middleware/grpc';
import { Code, Metadata } from 'grpc-web-client';
import { MentorService } from '../proto/mentor_pb_service';

export const LESSONS_INIT = 'LESSONS_INIT';
export const ADD_LESSON = 'ADD_LESSON';
export const SELECT_LESSON = 'SELECT_LESSON';

type AddLesson = {
  type: typeof ADD_LESSON,
  payload: Lesson,
};
export const addLesson = (lesson: Lesson) => ({ type: ADD_LESSON, payload: lesson });

type ListLessonsInit = {
  type: typeof LESSONS_INIT,
};

export const listLessonsInit = (): ListLessonsInit => ({type: LESSONS_INIT});

export const listLessons = () => {
  return grpcRequest<ListLessonsRequest, ListLessonsResponse>({
    request: new ListLessonsRequest(),
    onStart: () => listLessonsInit(),
    onEnd: (code: Code, message: string | undefined, trailers: Metadata): Action | void => {
      console.log(code, message, trailers);
      return;
    },
    host: 'http://localhost:8900',
    methodDescriptor: MentorService.ListLessons,
    onMessage: message => {
      const lesson = message.getLesson();
      if (lesson) {
        return addLesson(lesson);
      }
      return;
    },
  });
};

type SelectLesson = {
  type: typeof SELECT_LESSON,
  payload: number,
};
export const selectLesson = (lessonId: number): SelectLesson => ({ type: SELECT_LESSON, payload: lessonId });

export type LessonActionTypes =
  | ListLessonsInit
  | AddLesson
  | SelectLesson
  | GrpcAction<ListLessonsRequest, ListLessonsResponse>;
